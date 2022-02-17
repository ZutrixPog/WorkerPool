package workerpool

import "context"

type PtrWorkerPool[T any] struct {
	jobs chan Job[T] // Refer to job.go file!
	done chan bool
}

func NewPtrWorkerPool[T any]() *PtrWorkerPool[T] {
	workerpool := new(PtrWorkerPool[T])
	workerpool.jobs = make(chan Job[T])
	workerpool.done = make(chan bool, 1)
	return workerpool
}

func (w *PtrWorkerPool[T]) AddJob(ptr *T, f inFunc[T]) {
	j := PtrJob[T]{id: 0, ptr: ptr, fn: f}
	w.jobs <- j
}

func (w *PtrWorkerPool[T]) worker(ctx context.Context) {
	for {
		select {
		case job := <-w.jobs:
			job.Execute()
		case <-ctx.Done():
			return
		default:
			w.done <- true
		}
	}
}

func (w *PtrWorkerPool[T]) Start(numWorkers int, ctx context.Context) chan bool {
	for i := 0; i < numWorkers; i++ {
		go w.worker(ctx)
	}
	return w.done
}
