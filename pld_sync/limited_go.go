package pld_sync

import "sync"

type LimitedGo struct {
	n  int
	c  chan struct{}
	wg *sync.WaitGroup
}

func NewLimitedGo(n int) *LimitedGo {
	return &LimitedGo{
		n:  n,
		c:  make(chan struct{}, n),
		wg: &sync.WaitGroup{},
	}
}

func (g *LimitedGo) AddOne() {
	g.wg.Add(1)
}
func (g *LimitedGo) RunOne(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}

func (g *LimitedGo) DoneOne() {
	g.wg.Done()
}
func (g *LimitedGo) WaitAll() {
	g.wg.Wait()
}
