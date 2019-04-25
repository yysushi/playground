import gevent
import gevent.monkey
gevent.monkey.patch_all()

from gevent import Greenlet

import time
import random


class Task(object):

    def __init__(self, id, out):
        self.id = id
        self.out = out

    def __str__(self):
        return "Task(id:{},out:{})".format(self.id, self.out)


class Worker(Greenlet):

    def __init__(self, queue):
        Greenlet.__init__(self)
        self.queue = queue

    def _run(self):
        while True:
            task = self.queue.get()
            print(task.id, "is going")
            task.out = random.randint(0, 1)
            time.sleep(random.randint(0, 2))
            print(task.id, "is done")
            self.queue.task_done()


class Client(Greenlet):

    def __init__(self, queue, num_of_workers):
        Greenlet.__init__(self)
        self.queue = queue
        self.num_of_workers = num_of_workers

    def _run(self):
        start_id = 0
        while True:
            print("start from", start_id)
            tasks = []
            for i in range(self.num_of_workers):
                task = Task(start_id + i, 0)
                tasks.append(task)
                print("put task", task)
                self.queue.put_nowait(task)
            self.queue.join()
            if self.check(tasks[start_id:start_id+self.num_of_workers]):
                break
            start_id += self.num_of_workers
        return tasks

    @classmethod
    def check(cls, tasks):
        for t in tasks:
            if t.out == 0:
                return False
        return True


queue = gevent.queue.JoinableQueue()

num_of_workers = 4
for i in range(num_of_workers):
    Worker(queue).start()

client = Client(queue, num_of_workers)
client.start()
client.join()
