package test

import (
	workerpool "WorkerPool"
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	//p := fmt.Printf
	const numJobs = 6
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	pool := workerpool.NewWorkerPool[int]()
	res := pool.Start(6, ctx)
	for i := 0; i < numJobs; i++ {
		pool.AddJob(func(id int) int {
			time.Sleep(50 * time.Millisecond)
			return 2 * 2
		})
	}
	for i := 0; i < numJobs; i++ {
		r := <-res
		fmt.Println(r)
		if r != 4 {
			t.Errorf("Failed")
		}
		if i > 3 {
			break
		}
	}
}

func TestPtrWorkerPool(t *testing.T) {
	//p := fmt.Printf
	const numJobs = 6
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	pool := workerpool.NewPtrWorkerPool[int]()
	done := pool.Start(numJobs, ctx)
	res := make([]int, numJobs)
	for i := 0; i < numJobs; i++ {
		pool.AddJob(&res[i], func(id int) int {
			time.Sleep(50 * time.Millisecond)
			return 2 * 2
		})
	}
	<-done
	for i := 0; i < numJobs; i++ {
		r := res[i]
		fmt.Println(r)
		if reflect.TypeOf(r) == nil {
			t.Errorf("Failed")
		}
		if i > 3 {
			break
		}
	}
}
