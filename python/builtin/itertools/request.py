# import requests
import asyncio


async def get(i):
    # response = await requests.get(f'https://httpbin.org/anything/{i}', verify=False)
    print(f"{i} is starting")
    response = await asyncio.sleep(i*2)
    # return response.json()['url']
    print(f"{i} is done")
    return response


async def main():
    responses = await asyncio.gather(*map(get, range(10)))
    print(responses)


asyncio.run(main())
