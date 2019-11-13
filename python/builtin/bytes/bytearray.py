# immutable
# a = b'ab'
# a = b'あab'
# SyntaxError: bytes can only contain ASCII literal characters.
a = bytes('あab', 'utf-8')
# a = 'あab'.encode('utf-8')
try:
    a[0] = 98
except TypeError as e:
    print(e)
    # TypeError: 'bytes' object does not support item assignment
else:
    exit('no exception')
print(type(a), a)
# <class 'bytes'> b'\xe3\x81\x82ab'
print(type(a[1]), a[1])
# <class 'int'> 129

# mutable
a = bytearray('あab', 'utf-8')
a[0] = 98
print(a)
# bytearray(b'b\x81\x82ab')
print(type(a[1]), a[1])
# <class 'int'> 129
