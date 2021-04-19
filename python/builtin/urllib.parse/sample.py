import urllib.parse

url = 'http://example.com/index.html?true=true&true=True&true=1'

o = urllib.parse.urlparse(url)
qs = urllib.parse.parse_qs(o.query)
print(type(qs['true']))
print(qs)

# not 'true', but 'True'...
print(urllib.parse.urlencode({'true': True}))
