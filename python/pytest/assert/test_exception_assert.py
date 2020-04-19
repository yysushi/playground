import pytest

from myfunc import f2


def test_exception_assert():
    with pytest.raises(SystemExit):
        f2(3)
