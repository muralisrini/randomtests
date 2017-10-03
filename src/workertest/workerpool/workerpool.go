package workerpool

import (
	"fmt"
	"github.com/pkg/errors"
)

// WorkerPool represents a pool of pre-started goroutines
// that wait for work, execute it, and produce results
type WorkerPool interface {
	// AssignWorkAsync is used to provide the worker pool
	// with items of work that will be processed by the
	// first free goroutine; AssignWorkAsync is always
	// a non-blocking call
	AssignWorkAsync(...func())

	// AssignWorkSync is used to provide the worker pool
	// with items of work that will be processed by the
	// first free goroutine; AssignWorkSync may block
	AssignWorkSync(...func())
}

type genericWorkerPool struct {
	work chan func()
}

// NewGenericWorkerPool returns a WorkerPool with `nWorkers` goroutines,
// This function returns an error if the supplied arguments are invalid
func NewGenericWorkerPool(nWorkers int) (WorkerPool, error) {
	if nWorkers <= 0 {
		return nil, errors.New("invalid arguments to NewGenericWorkerPool")
	}

	w := &genericWorkerPool{make(chan func())}

	//fmt.Printf("Starting %d workers\n", nWorkers)
	for i := 0; i < nWorkers; i++ {
		go w.doWork(i)
	}

	return w, nil
}

// NewGenericWorkerPoolOrPanic returns a WorkerPool with the same
// characteristics as that returned by NewGenericWorkerPool. This
// function panics instead of returning an error if the supplied
// arguments are invalid
func NewGenericWorkerPoolOrPanic(nWorkers int) WorkerPool {
	wp, err := NewGenericWorkerPool(nWorkers)
	if err != nil {
		panic(fmt.Sprintf("NewGenericWorkerPool failed, err %s", err))
	}

	return wp
}

func (gw *genericWorkerPool) AssignWorkSync(work ...func()) {
	for _, workItem := range work {
		gw.work <- workItem
	}
}

func (gw *genericWorkerPool) AssignWorkAsync(work ...func()) {
	go gw.AssignWorkSync(work...)
}

func (gw *genericWorkerPool) doWork(i int) {
	//fmt.Printf("Worker %d starts\n", i)

	for {
		work, ok := <-gw.work
		if ok {
			//fmt.Printf("Worker %d starts work\n", i)
			work()
			//fmt.Printf("Worker %d completes work\n", i)
		} else {
			//fmt.Printf("Worker %d returns\n", i)
			return
		}
	}
}

