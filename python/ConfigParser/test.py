import ConfigParser

default = {
    'dir': 'whao',
    'newdir': 'newwhao',
}

config = ConfigParser.RawConfigParser(default, allow_no_value=True)
config.read("sample.ini")
config.set('My Section', 'name', 'hmmm')

print(config._sections)

# config.get()
print(config.items("My Section"))
