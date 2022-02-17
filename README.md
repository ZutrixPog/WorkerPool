# WorkerPool
Two Test Implementations of Worker Pools in go considering two diffrent Approaches.
You can use a context to control the life cycle of your worker pool.
you also have the chance to choose from two kinds of pools:
  1. WorkerPool: Uses a channel to return the result.
  2. PtrWorkerPool: You can provide your destination Pointer.
