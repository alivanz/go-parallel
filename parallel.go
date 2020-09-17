package parallel

import (
	"context"
	"sync"
)

type Parallel struct {
	ctx  context.Context
	lock sync.Locker
	wg   sync.WaitGroup
}

type Func func(ctx context.Context)

func NewParallel(ctx context.Context, lock sync.Locker) *Parallel {
	return &Parallel{
		ctx:  ctx,
		lock: lock,
	}
}

// Do call function in different go routine
func (parallel *Parallel) Do(f Func) {
	parallel.wg.Add(1)
	go func() {
		defer parallel.wg.Done()
		parallel.lock.Lock()
		defer parallel.lock.Unlock()
		f(parallel.ctx)
	}()
}

// DoBlock lock and call function in different go routine
func (parallel *Parallel) DoBlock(f Func) {
	parallel.wg.Add(1)
	parallel.lock.Lock()
	go func() {
		defer parallel.wg.Done()
		defer parallel.lock.Unlock()
		f(parallel.ctx)
	}()
}

// Wait wait all goroutine
func (parallel *Parallel) Wait() {
	parallel.wg.Wait()
}
