from b import B
from c import C, C2
from c.c2 import C2 as C3
# from d.d import D

print(B(1, 2))
print(C(1), C2(2), C3(3))
# print(D(1, 2, 3))
