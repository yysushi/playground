import itertools


a = zip(iter([iter([1]),iter([2])]), iter([3,4]))
print(a)

aa, bb = next(a)
print(next(aa))
# next()
# print(next(a))
# print(any(a))
# print(any(a))
# print(next(a))
# print(any(a))
