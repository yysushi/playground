import get_ip
import pytest


# custom class to be the mock return value of requests.get()
class MockResponse:
    @staticmethod
    def json():
        return {"origin": "8.8.8.8"}


# monkeypatched requests.get moved to a fixture
@pytest.fixture
def mock_response(monkeypatch):

    def mock_get(*args, **kwargs):
        return MockResponse()

    monkeypatch.setattr(get_ip.requests, "get", mock_get)


# notice our test uses the custom fixture instead of monkeypatch directly
def test_get_ip(mock_response):
    assert get_ip.get_ip() == '8.8.8.8'
