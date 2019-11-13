import re

# pattern = r'def\s+([a-zA-Z_][a-zA-Z_0-9\n]*)\s*\(\s*\):'
# orig = 'def myf\nunc():'
# # orig = 'def myfunc():'
# 
# a = re.sub(pattern,
#            r'static PyObject*\npy_\1(void)\n{',
#            orig,
#            re.MULTILINE)
# print(orig)
# print(a)

pattern = r'b\nb'
orig = """ab\nbc
ab\nbbc"""

a = re.sub(pattern, r'b', orig, re.MULTILINE)
print(a)
