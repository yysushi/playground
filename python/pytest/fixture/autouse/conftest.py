import pytest


@pytest.fixture(autouse=True)
def print_in_advance():
    print("here")


@pytest.fixture
def print_in_advance2():
    print("here2")
