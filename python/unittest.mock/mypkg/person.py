import json
import time
import urllib.request

from .external import SOURCE, get_active_source, Discovery


class Person:

    def __init__(self, name):
        self.name = name
        self.d = Discovery()

    def greet(self):
        return "hello %s" % self.name

    def sleep(self, seconds):
        # print(time.sleep)
        time.sleep(seconds)

    def is_where_born(self):
        print(SOURCE)
        with urllib.request.urlopen(SOURCE) as f:
            data = f.read()
        return json.loads(data)['origin'].split(',')[0]

    def is_where_living(self):
        source = get_active_source()
        print(source)
        with urllib.request.urlopen(source) as f:
            data = f.read()
        return json.loads(data)['origin'].split(',')[0]

    def is_where_living2(self):
        source = self.d.active_source()
        print(source)
        with urllib.request.urlopen(source) as f:
            data = f.read()
        return json.loads(data)['origin'].split(',')[0]
