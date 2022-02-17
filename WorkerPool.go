package workerpool

import "context"

type inFunc[T any] func(id int) T

type WorkerPool[T any] struct {
	jobs   chan Job[T] // Refer to job.go file!
	result chan T
}

func NewWorkerPool[T any]() *WorkerPool[T] {
	workerpool := new(WorkerPool[T])
	workerpool.jobs = make(chan Job[T])
	workerpool.result = make(chan T)
	return workerpool
}

func (w *WorkerPool[T]) AddJob(f inFunc[T]) {
	j := SimpJob[T]{id: 0, fn: f}
	w.jobs <- j
}

func (w *WorkerPool[T]) AddJobPtr(ptr *T, f inFunc[T]) {
	j := PtrJob[T]{id: 0, ptr: ptr, fn: f}
	w.jobs <- j
}

func (w *WorkerPool[T]) worker(ctx context.Context) {
	for {
		select {
		case job := <-w.jobs:
			if _, ok := job.(PtrJob[T]); !ok {
				w.result <- job.Execute()
			} else {
				job.Execute()
			}
		case <-ctx.Done():
			return
		}
	}
}

func (w *WorkerPool[T]) Start(numWorkers int, ctx context.Context) chan T {
	for i := 0; i < numWorkers; i++ {
		go w.worker(ctx)
	}
	return w.result
}
