package parallel

import (
	"sync"
)

type Limit struct {
	c chan struct{}
}

// NewLimit this locker can be Lock() n times in parallel. n = 0 means no limit
func NewLimit(n int) sync.Locker {
	if n < 1 {
		return NoLimit{}
	}
	return &Limit{
		c: make(chan struct{}, n),
	}
}

func (limit *Limit) Lock() {
	limit.c <- struct{}{}
}
func (limit *Limit) Unlock() {
	<-limit.c
}
