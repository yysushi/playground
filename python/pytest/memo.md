# Notes for pytest

## How to Run

```shell-session
koketani:pytest (master *+%=)$ pytest -v
========================================================================================================= test session starts ==========================================================================================================
platform darwin -- Python 3.7.7, pytest-5.4.1, py-1.8.1, pluggy-0.13.1 -- /Users/y-tsuji/.local/share/virtualenvs/pytest-j3f5aT7B/bin/python3
cachedir: .pytest_cache
rootdir: /Users/y-tsuji/Developments/git/github.com/koketani/playground/python/pytest
collected 1 item

test_simple.py::test_answer FAILED                                                                                                                                                                                               [100%]

=============================================================================================================== FAILURES ===============================================================================================================
_____________________________________________________________________________________________________________ test_answer ______________________________________________________________________________________________________________

    def test_answer():
>       assert func(3) == 5
E       assert 4 == 5
E         +4
E         -5

test_simple.py:6: AssertionError
======================================================================================================= short test summary info ========================================================================================================
FAILED test_simple.py::test_answer - assert 4 == 5
========================================================================================================== 1 failed in 0.04s ===========================================================================================================
```

## Discover

```shell-session
koketani:pytest (master *%>)$ pytest -v
========================================================================================================= test session starts ==========================================================================================================
platform darwin -- Python 3.7.7, pytest-5.4.1, py-1.8.1, pluggy-0.13.1 -- /Users/y-tsuji/.local/share/virtualenvs/pytest-j3f5aT7B/bin/python3
cachedir: .pytest_cache
rootdir: /Users/y-tsuji/Developments/git/github.com/koketani/playground/python/pytest, inifile: pytest.ini
collected 4 items

assert/test_advanced_assert.py::test_answer FAILED                                                                                                                                                                               [ 25%]
discover/discover2_test.py::test_discover2 PASSED                                                                                                                                                                                [ 50%]
discover/test_discover1.py::test_discover1 PASSED                                                                                                                                                                                [ 75%]
discover/discover3/discover3.py::test_discover1 PASSED                                                                                                                                                                           [100%]

=============================================================================================================== FAILURES ===============================================================================================================
_____________________________________________________________________________________________________________ test_answer ______________________________________________________________________________________________________________

    def test_answer():
>       assert myfunc.f(3) == 5
E       assert 4 == 5
E         +4
E         -5

assert/test_advanced_assert.py:5: AssertionError
======================================================================================================= short test summary info ========================================================================================================
FAILED assert/test_advanced_assert.py::test_answer - assert 4 == 5
===================================================================================================== 1 failed, 3 passed in 0.06s ======================================================================================================
```

## Features

- assert
- discover
- class
- fixture
