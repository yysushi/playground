import contextlib
import typing as t

import tenacity as tc


class TargetException(BaseException): ...


class OtherException(BaseException): ...


def target(error: BaseException | None = None):
    if error:
        raise error


class SuccessError(AssertionError): ...


class FailureError(AssertionError): ...


def attempts[T](
    max_attempt: int = 2,
    wait_fixed: float = 0.2,
    error: type[T] | T | None = None,
    stop_error_type: type[SuccessError] | type[FailureError] | None = None,
) -> t.Iterator[t.ContextManager[tc.RetryCallState]]:
    def attempter(retry: tc.AttemptManager) -> t.ContextManager[tc.RetryCallState]:
        @contextlib.contextmanager
        def asserter() -> t.Iterator[tc.RetryCallState]:
            with retry:
                try:
                    yield retry.retry_state
                except BaseException as exc:
                    if (
                        error is None
                        or (isinstance(error, type) and not isinstance(exc, error))
                        or (
                            not isinstance(error, type)
                            and not isinstance(exc, type(error))
                        )
                        or (
                            not isinstance(error, type)
                            and isinstance(exc, type(error))
                            and str(error) not in str(exc)
                        )
                    ):
                        raise FailureError
                else:
                    if error:
                        raise SuccessError

        return asserter()

    if stop_error_type:

        class my_stop(tc.stop.stop_base):
            def __call__(self, retry_state: tc.RetryCallState) -> bool:
                if retry_state.outcome is None or not retry_state.outcome.failed:
                    return False
                return isinstance(retry_state.outcome.exception(), stop_error_type)

        stop = tc.stop_any(tc.stop_after_attempt(max_attempt), my_stop())
    else:
        stop = tc.stop_after_attempt(max_attempt)

    for retry in tc.Retrying(
        stop=stop,
        wait=tc.wait_fixed(wait_fixed),
    ):
        yield attempter(retry)


def test_expected_error_type():
    for attempt in attempts(error=TargetException):
        with attempt:
            target(TargetException("hoge"))


def test_unexpected_error_type():
    try:
        for attempt in attempts(error=TargetException):
            with attempt:
                target(OtherException("hoge"))
    except tc.RetryError:
        pass
    else:
        raise AssertionError("succeeded unexpectedly")


def test_expected_error_instance():
    for attempt in attempts(error=TargetException("hoge")):
        with attempt:
            target(TargetException("hoge"))


def test_unexpected_error_instance():
    try:
        for attempt in attempts(error=TargetException("hoge2")):
            with attempt:
                target(TargetException("hoge"))
    except tc.RetryError:
        pass
    else:
        raise AssertionError("succeeded unexpectedly")


def test_expected_no_error():
    for attempt in attempts():
        with attempt:
            target()


def test_unexpected_no_error():
    try:
        for attempt in attempts(error=TargetException("hoge")):
            with attempt:
                target()
    except tc.RetryError:
        pass
    else:
        raise AssertionError("succeeded unexpectedly")


def test_unexpected_no_error_with_immediate_exit():
    try:
        for attempt in attempts(
            error=TargetException("hoge"), stop_error_type=SuccessError
        ):
            with attempt:
                target()
    except tc.RetryError as e:
        assert e.last_attempt.failed
        assert e.last_attempt.attempt_number == 1
        assert isinstance(e.last_attempt.exception(), SuccessError)
    else:
        raise AssertionError("succeeded unexpectedly")
