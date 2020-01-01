# Notes for pytest

## with unittest

- run with unittest module

```shell-session
koketani:pytest (master *%=)$ python -m unittest -v
test_get_ip (test_get_ip.TestGetIP) ... ok
test_get_ip_failure (test_get_ip.TestGetIP) ... ok

----------------------------------------------------------------------
Ran 2 tests in 0.001s

OK
```

- run with pytest module

```shell-session
koketani:pytest (master *%=)$ pytest -v
======================================== test session starts ========================================
platform darwin -- Python 3.7.5, pytest-5.3.2, py-1.8.1, pluggy-0.13.1 -- /Users/koketani/.local/share/virtualenvs/pytest-lIbIOdJU/bin/python3
cachedir: .pytest_cache
rootdir: /Users/koketani/Developments/go/src/github.com/koketani/playground/python/pytest
collected 2 items

test_get_ip.py::TestGetIP::test_get_ip PASSED                                                 [ 50%]
test_get_ip.py::TestGetIP::test_get_ip_failure PASSED                                         [100%]

========================================= 2 passed in 0.09s =========================================
```

## with pytest fixture

- <https://docs.pytest.org/en/latest/monkeypatch.html#monkeypatching-returned-objects-building-mock-classes>

```shell-session

koketani:pytest (master *%=)$ pytest -v
======================================== test session starts ========================================
platform darwin -- Python 3.7.5, pytest-5.3.2, py-1.8.1, pluggy-0.13.1 -- /Users/koketani/.local/share/virtualenvs/pytest-lIbIOdJU/bin/python3
cachedir: .pytest_cache
rootdir: /Users/koketani/Developments/go/src/github.com/koketani/playground/python/pytest
collected 3 items

test_get_ip.py::TestGetIP::test_get_ip PASSED                                                 [ 33%]
test_get_ip.py::TestGetIP::test_get_ip_failure PASSED                                         [ 66%]
test_get_ip2.py::test_get_ip PASSED                                                           [100%]

========================================= 3 passed in 0.09s =========================================
```
