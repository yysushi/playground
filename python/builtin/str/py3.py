print(type("aaa"))
# <class 'str'>

print(type(b"aaa"))
# <class 'bytes'>

print(type(u"aaa"))
# <class 'str'>

try:
    print(type(bytearray("aaa")))
except TypeError as e:
    print(e)
    # TypeError: string argument without an encoding

print(type(bytearray(b"aaa")))
# <class 'bytearray'>
