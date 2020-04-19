from pathlib import Path
import requests


def getssh():
    return Path.home() / ".ssh"


def get_json(url):
    return requests.get(url).json()
