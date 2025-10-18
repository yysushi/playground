from contextlib import contextmanager

import tenacity as tc

for retry in tc.Retrying():
    with retry:
        print("aiu")
