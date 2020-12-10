import sys
import inspect


def second(key):
    frame = inspect.currentframe()
    frames = inspect.getouterframes(frame)
    infos = []
    for f in frames:
        argvalues = inspect.getargvalues(f[0])
        args = inspect.formatargvalues(*argvalues)
        info = "{} {} {} {}".format(f[1], f[2], f[3], args)
        infos.append(info)
    print("\n".join(infos))
    raise Exception("some error")


def first(key):
    second(key+1)


first(1)
