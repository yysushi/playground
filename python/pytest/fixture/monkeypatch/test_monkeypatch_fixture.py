from myfunc import getssh, get_json

from pathlib import Path
import requests


def test_getssh(monkeypatch):
    monkeypatch.setattr(Path, "home", lambda: Path("/monkey"))
    assert getssh() == Path("/monkey/.ssh")


class MockResonse:
    @staticmethod
    def json():
        return {"mock_key": "mock_response"}


def test_get_json(monkeypatch):
    monkeypatch.setattr(requests, "get", lambda x: MockResonse())
    json = get_json("https://httpbin.org/ip")
    assert json == {"mock_key": "mock_response"}
