from retrying import Retrying


def three_and_error():
    a = [1, 2, 3]
    for idx, aa in enumerate(a):
        print(idx)
        yield aa


def result_check(result):
    print(dir(result))
    print(result.get())
    return True


def exception_check(exc):
    print(dir(exc))
    print(exc.__class__ == StopIteration)
    # return False
    return exc.__class__ != StopIteration


b = three_and_error()
# retry = Retrying(retry_on_result=result_check, wrap_exception=True).call(next, b)
# retry = Retrying(retry_on_result=result_check).call(next, b)
# retry = Retrying(retry_on_exception=exception_check, wrap_exception=True).call(next, b)
# retry = Retrying(retry_on_result=lambda x: True, retry_on_exception=exception_check, wrap_exception=True).call(next, b)
retry = Retrying(retry_on_result=lambda x: True, retry_on_exception=exception_check).call(next, b)
