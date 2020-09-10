package pld_sync

import "sync"

type LimitedGo struct {
	n  int
	c  chan struct{}
	wg sync.WaitGroup
}

func NewLimitedGo(n int) *LimitedGo {
	return &LimitedGo{
		n:  n,
		c:  make(chan struct{}, n),
		wg: sync.WaitGroup{},
	}
}

func (g *LimitedGo) Add(f func()) {
	g.wg.Add(1)
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}
func (g *LimitedGo) Done() {
	g.wg.Done()
}
func (g *LimitedGo) Wait() {
	g.wg.Wait()
}
