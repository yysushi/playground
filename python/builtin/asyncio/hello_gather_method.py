import asyncio


async def hoge_call(args):
    await getattr(args[0], 'hoge')(args[1])


class Me:
    async def hoge(self, i):
        await asyncio.sleep(i)
        print(i)


async def main():
    objects = [Me(), Me()]
    all_args = zip(objects, range(2))
    await asyncio.gather(*map(hoge_call, all_args))


asyncio.run(main())
