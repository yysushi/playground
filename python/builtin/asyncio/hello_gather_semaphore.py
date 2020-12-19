import asyncio
import time
import functools


async def say(s, i):
    print(f"{i} at {time.strftime('%X')}: I'm just in")
    async with s:
        await asyncio.sleep(3)
    print(f"{i} at {time.strftime('%X')}: I'm being out of function")
    return i


async def main():
    semaphore = asyncio.Semaphore(33)
    _say = functools.partial(say, semaphore)
    values = await asyncio.gather(*map(_say, range(100)))
    print(values)


asyncio.run(main())
