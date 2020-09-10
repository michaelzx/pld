package pld_sync

import (
	"fmt"
	"testing"
	"time"
)

func TestOne(t *testing.T) {
	number := 10
	g := NewLimitedGo(2)
	for i := 0; i < number; i++ {
		g.AddOne()
		value := i
		goFunc := func() {
			// 做一些业务逻辑处理
			fmt.Printf("go func: %d\n", value)
			sleepSec := 1
			if value%2 == 0 {
				sleepSec = 2
			}
			// fmt.Printf("sleep: %d\n", sleepSec)
			time.Sleep(time.Second * time.Duration(sleepSec))
			g.DoneOne()
		}
		g.RunOne(goFunc)
	}
	g.WaitAll()
	fmt.Println("完成！")
}
