import types
import time


class A(types.SimpleNamespace):
    @property
    def a(self):
        return int(time.time())


a = A(**{
    'one': 1,
    'two': 2,
})

b = a.a
