import base64

# b = base64.b64encode('あ'.encode('utf-8'))
b = base64.b64encode(bytes('あ', 'utf-8'))
print(type(b), b)

# b = base64.b64encode('a'.encode('utf-8'))
# # b = base64.b64encode(b'a')
# print(type(b), b)

cc = base64.b64decode(b)
print(type(cc), cc)
c = cc.decode('utf-8')
# print(type(c), c.encode('utf-8'))
print(u'あ')
# print(type(c))

# print(base64.b64decode(b'44GC').decode('utf-8'))
# print(base64.b64decode(b))
