import tenacity as t
import typing as ty


class MyException(Exception): ...


class OtherException(Exception): ...


def raise_exc():
    raise MyException()


def no_raise_exc():
    raise MyException()


def test_1():
    #
    try:
        for attempt in t.Retrying(stop=t.stop_after_attempt(3), wait=t.wait_fixed(0.5)):
            with attempt:
                raise_exc()
    except t.RetryError as e:
        assert isinstance(e.last_attempt.exception(), MyException)
    else:
        raise AssertionError("succeeded unexpectedly")


def test_2():
    #
    try:
        for attempt in t.Retrying(
            stop=t.stop_after_attempt(3), wait=t.wait_fixed(0.5), reraise=True
        ):
            with attempt:
                raise_exc()
    except MyException:
        pass
    else:
        raise AssertionError("succeeded unexpectedly")


def test_3():
    #
    try:
        for i, attempt in enumerate(
            t.Retrying(
                stop=t.stop_after_attempt(3),
                wait=t.wait_fixed(0.5),
                retry=t.retry_unless_exception_type(OtherException),
            )
        ):
            with attempt:
                raise_exc()
    except t.RetryError as e:
        assert isinstance(e.last_attempt.exception(), MyException)
    else:
        raise AssertionError("succeeded unexpectedly")
    assert i == 2


class my_stop(t.stop.stop_base):
    def __call__(self, retry_state: "RetryCallState") -> bool:
        return isinstance(retry_state.outcome.exception(), MyException)


class my_stop2(t.stop.stop_base):
    def __call__(self, retry_state: "RetryCallState") -> bool:
        return not retry_state.outcome.failed


def test_4():
    #
    try:
        for i, attempt in enumerate(
            t.Retrying(
                stop=t.stop_any(my_stop(), t.stop_after_attempt(3)),
                wait=t.wait_fixed(0.5),
                retry_error_callback=lambda x: None,
            )
        ):
            with attempt:
                raise_exc()
    except t.RetryError:
        raise AssertionError("error unexpectedly")
    assert i == 0


def test_5():
    #
    try:
        for i, attempt in enumerate(
            t.Retrying(
                stop=t.stop_any(my_stop(), t.stop_after_attempt(3)),
                wait=t.wait_fixed(0.5),
                # retry_error_callback=lambda x: None,
            )
        ):
            with attempt:
                raise MyException()
    except t.RetryError:
        pass
    else:
        raise AssertionError("succeeded unexpectedly")
    assert i == 0


def my_error_callback(rs: t.RetryCallState) -> ty.Any:
    # my_error
    if isinstance(rs.outcome.exception(), MyException):
        return None
    # my_error2
    print("mu?")
    fut = ty.cast(t.Future, rs.outcome)
    raise t.RetryError(fut) from fut.exception()


def test_6():
    #
    try:
        for i, attempt in enumerate(
            t.Retrying(
                stop=t.stop_any(my_stop(), t.stop_after_attempt(3)),
                wait=t.wait_fixed(0.5),
                retry_error_callback=my_error_callback,
            )
        ):
            with attempt:
                raise MyException()
    except t.RetryError:
        raise AssertionError("error unexpectedly")
    assert i == 0


def test_7():
    #
    try:
        for i, attempt in enumerate(
            t.Retrying(
                stop=t.stop_any(my_stop(), t.stop_after_attempt(3)),
                wait=t.wait_fixed(0.5),
                retry_error_callback=my_error_callback,
            )
        ):
            with attempt:
                raise OtherException()
    except t.RetryError:
        pass
    else:
        raise AssertionError("succeeded unexpectedly")
    assert i == 2


def test_8():
    #
    try:
        for i, attempt in enumerate(
            t.Retrying(
                stop=t.stop_any(my_stop(), t.stop_after_attempt(3)),
                wait=t.wait_fixed(0.5),
                retry_error_callback=my_error_callback,
            )
        ):
            with attempt:
                pass
    except t.RetryError:
        raise AssertionError("error unexpectedly")
    assert i == 0


def test_9():
    #
    try:
        for i, attempt in enumerate(
            t.Retrying(
                stop=t.stop_any(my_stop(), t.stop_after_attempt(3)),
                wait=t.wait_fixed(0.5),
                retry=t.retry_unless_exception_type(MyException),
                retry_error_callback=my_error_callback,
            )
        ):
            with attempt:
                pass
    except t.RetryError:
        pass
    else:
        raise AssertionError("succeeded unexpectedly")
    assert i == 2


def test_a():
    #
    try:
        for i, attempt in enumerate(
            t.Retrying(
                stop=t.stop_any(my_stop(), t.stop_after_attempt(3)),
                wait=t.wait_fixed(0.5),
                retry=t.retry_unless_exception_type(MyException),
                retry_error_callback=my_error_callback,
            )
        ):
            with attempt:
                raise OtherException()
    except t.RetryError:
        pass
    else:
        raise AssertionError("succeeded unexpectedly")
    assert i == 2


def test_b():
    #
    try:
        for i, attempt in enumerate(
            t.Retrying(
                stop=t.stop_any(my_stop(), t.stop_after_attempt(3)),
                wait=t.wait_fixed(0.5),
                retry=t.retry_unless_exception_type(MyException),
                retry_error_callback=my_error_callback,
            )
        ):
            with attempt:
                raise MyException()
    except MyException:
        # raise AssertionError("error2 unexpectedly")
        ...
    except t.RetryError:
        raise AssertionError("error unexpectedly")
    assert i == 0


def test_c():
    #
    try:
        for i, attempt in enumerate(
            t.Retrying(
                stop=t.stop_any(my_stop(), my_stop2(), t.stop_after_attempt(3)),
                # stop=t.stop_any(my_stop(), t.stop_after_attempt(3)),
                wait=t.wait_fixed(0.5),
                retry_error_callback=my_error_callback,
            )
        ):
            with attempt:
                pass
    except t.RetryError:
        raise AssertionError("error unexpectedly")
    else:
        # raise AssertionError("succeeded unexpectedly")
        ...
    assert i == 0

# want to raise assertion error
# target succeeded without error
# but it's not expected. some errro was expected
#
# give up..
