import requests


class GetIPError(Exception):
    pass


def get_ip():
    try:
        ip = requests.get("http://httpbin.org/ip").json()
    except requests.exceptions.RequestException as e:
        raise GetIPError(e)
    return ip['origin']


if __name__ == '__main__':
    print(get_ip())
