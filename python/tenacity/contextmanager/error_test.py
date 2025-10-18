import contextlib
import typing as t

import tenacity as tc


class TargetException(Exception): ...


def target(error: str = ""):
    if error:
        raise TargetException(error)
    else:
        pass


def attempts[T](
    max_attempt: int = 2,
    wait_fixed: float = 0.2,
    error: type[T] | None = None,
) -> t.Iterator[t.ContextManager[None]]:
    def attempter(retry: tc.AttemptManager) -> t.ContextManager[None]:
        @contextlib.contextmanager
        def asserter() -> t.Iterator[None]:
            with retry:
                try:
                    yield
                except Exception as exc:
                    if error is None or not isinstance(exc, error):
                        raise AssertionError("unexpected error")
                else:
                    if error:
                        raise AssertionError("unexpected success")

        return asserter()

    for retry in tc.Retrying(
        stop=tc.stop_after_attempt(max_attempt),
        wait=tc.wait_fixed(wait_fixed),
    ):
        yield attempter(retry)


def test_expected_error():
    for attempt in attempts(error=TargetException):
        with attempt:
            target("hoge")


def test_expected_no_error():
    for attempt in attempts():
        with attempt:
            target()
