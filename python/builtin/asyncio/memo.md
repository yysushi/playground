# Note for asyncio

<https://docs.python.org/ja/3/library/asyncio.html>

<https://pymotw.com/3/asyncio/concepts.html>

## Reference

- high level API
- low level API
  - eventloop
  - future

## Coroutine and Task

### Coroutine

To run coroutine, asyncio provides three ways.

1. To run a top level entrypoint main(), provides an event loop by `asyncio.run(main())`

2. Await a coroutine
3. Run a coroutine as Tasks by `asyncio.create_task(some_coroutine())` and await it

```
koketani: ~/g/g/k/p/p/asyncio (master ?)$ python hello.py
started at 16:07:14
hello
world
finished at 16:07:17
koketani: ~/g/g/k/p/p/asyncio (master ?)$ python hello2.py
started at 16:07:20
hello
world
finished at 16:07:22
```

### Awaitable objects

objects to call with `await xxx()` are

- coroutine
- task
- future

we have two closely concepts to coroutine

- coroutine funciton: `async def xxxx()`
- coroutine object: return object when calling coroutine function

### References

- run async program `asyncio.run(main())`
- create task `asyncio.create_task(some_coroutine())`

according to [create_task vs ensure_feature](https://stackoverflow.com/a/36415477),
```
Starting from Python 3.7 asyncio.create_task(coro) high-level function was added for this purpose.

You should use it instead other ways of creating tasks from coroutimes. However if you need to create task from arbitrary awaitable, you should use asyncio.ensure_future(obj).
```

- run tasks concurrently `await asyncio.gather(some_coroutine())`
- preserve from cacellation `await asyncio.shield(some_coroutine())`
- timeout `await asyncio.wait_for(some_coroutine(), 1)`
- wait finish for element `await asyncio.wait()`
