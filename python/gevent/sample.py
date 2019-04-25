import gevent
import gevent.monkey
gevent.monkey.patch_all()

import time
import requests


def send(i):
    print("start", i)
    response = requests.get("http://httpbin.org/ip")
    time.sleep(2-i)
    print(response.text)


gevent.joinall([
    gevent.spawn(send, 0),
    gevent.spawn(send, 1),
])
