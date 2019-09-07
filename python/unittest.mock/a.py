import json
import time
import urllib.request


class Person(object):

    SOURCE = 'httpbin.org/ip'

    def __init__(self, name):
        self.name = name

    def greet(self):
        return "hello %s" % self.name

    def sleep(self, seconds):
        time.sleep(seconds)

    def is_where(self):
        print(self.SOURCE)
        with urllib.request.urlopen('http://' + self.SOURCE) as f:
            data = f.read()
        return json.loads(data)['origin'].split(',')[0]
