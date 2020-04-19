"""
koketani:pytest (master %>)$ pytest --capture=no fixture/autouse
========================================================================================================= test session starts ==========================================================================================================
platform darwin -- Python 3.7.7, pytest-5.4.1, py-1.8.1, pluggy-0.13.1
rootdir: /Users/y-tsuji/Developments/git/github.com/koketani/playground/python/pytest, inifile: pytest.ini
collected 2 items

fixture/autouse/test_autouse.py here
.here
here2
.

========================================================================================================== 2 passed in 0.01s ===========================================================================================================
"""


def test_autouse1():
    pass


def test_autouse2(print_in_advance2):
    pass
