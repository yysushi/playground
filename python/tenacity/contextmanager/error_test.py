import contextlib
import typing as t

import tenacity as tc


class TargetException(BaseException): ...


class OtherException(BaseException): ...


def target(error: BaseException | None = None):
    if error:
        raise error


def attempts[T](
    max_attempt: int = 2,
    wait_fixed: float = 0.2,
    error: type[T] | T | None = None,
) -> t.Iterator[t.ContextManager[None]]:
    def attempter(retry: tc.AttemptManager) -> t.ContextManager[None]:
        @contextlib.contextmanager
        def asserter() -> t.Iterator[None]:
            with retry:
                try:
                    yield
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
                        raise AssertionError(f"unexpected error {exc}")
                else:
                    if error:
                        raise AssertionError("unexpected success")

        return asserter()

    for retry in tc.Retrying(
        stop=tc.stop_after_attempt(max_attempt),
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
