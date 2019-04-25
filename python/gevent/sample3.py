import gevent
import gevent.monkey
gevent.monkey.patch_all()

import time  # noqa
import random  # noqa


def run(idx):
    out = random.randint(0, 3)
    time.sleep(random.randint(0, 2))
    if out == 0:
        return True
    elif out == 1:
        raise Exception("some error")
    else:
        return False


def loop_until_success():
    while True:
        print("continue...")
        greenlets = []
        for i in range(4):
            greenlets.append(gevent.spawn(run, i))
        gevent.joinall(greenlets)
        go_out = True
        for g in greenlets:
            go_out = go_out and (g.successful() and g.value)
            if g.successful():
                print(g.value)
            else:
                print(g.exception)
        if go_out:
            break


loop_until_success()
