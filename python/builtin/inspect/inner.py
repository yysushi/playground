import sys
import inspect


def second(key):
    raise Exception("some error")


def first(key):
    second(key+1)


def main():
    try:
        first(1)
    except:
        _, _, tb = sys.exc_info()
        infos = []
        for f in inspect.getinnerframes(tb):
            argvalues = inspect.getargvalues(f[0])
            args = inspect.formatargvalues(*argvalues)
            info = "{} {} {} {}".format(f[1], f[2], f[3], args)
            infos.append(info)
        print("\n".join(infos))


main()
