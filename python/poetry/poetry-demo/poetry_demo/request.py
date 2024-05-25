import httpx

def google() -> int:
    resp = httpx.get("https://google.com")
    return resp.status_code
