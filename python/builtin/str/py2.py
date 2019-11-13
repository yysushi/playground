print(type("aaa"))
# <type 'str'>

print(type(b"aaa"))
# <type 'bytes'>

print(type(u"aaa"))
# <type 'unicode'>

print(type(bytearray("aaa")))
# <type 'bytearray'>

print(type(bytearray(b"aaa")))
# <type 'bytearray'>

try:
    print(type(bytearray(u"aaa")))
except TypeError as e:
    print(e)
    # unicode argument without an encoding
