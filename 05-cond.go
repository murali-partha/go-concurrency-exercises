package main

import (
	"fmt"
	"sync"
	"time"
)

type CondCounter struct {
	c     sync.Cond
	count int
}

func main() {
	var cCounter CondCounter
	cCounter.c = *sync.NewCond(&sync.Mutex{})

	increment := func() {
		for i := 0; i < 5; i++ {
			cCounter.c.L.Lock()
			cCounter.count++
			fmt.Println("Incremented:", cCounter.count)
			cCounter.c.Signal()
			time.Sleep(time.Millisecond * 500)
			cCounter.c.L.Unlock()
		}
	}

	go increment()
	for i := 0; i < 5; i++ {
		cCounter.c.L.Lock()
		for cCounter.count == 0 {
			cCounter.c.Wait()
		}
		cCounter.count--
		fmt.Println("Decremented:", cCounter.count)

		cCounter.c.L.Unlock()
	}

}
