import tenacity  as t

class MyException(Exception): ...

def raise_exc():
    raise MyException()

def no_raise_exc():
    raise MyException()

def test_raise_exc():
    for attempt in t.Retrying():
        with attempt:
            raise_exc()
