import importlib
import time

import sample
# from sample import b


print(sample.b)
time.sleep(2)
print(sample.b)
importlib.reload(sample)
time.sleep(2)
print(sample.b)
