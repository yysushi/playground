import asyncio
import time


async def say(i):
    print(f"{i} at {time.strftime('%X')}: I'm just in")
    await asyncio.sleep(3)
    print(f"{i} at {time.strftime('%X')}: I'm being out of function")
    return i


async def main():
    values = await asyncio.gather(*map(say, range(100)))
    print(values)


asyncio.run(main())
