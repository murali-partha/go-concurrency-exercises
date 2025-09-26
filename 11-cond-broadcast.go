package main

import (
	"fmt"
	"sync"
	"time"
)

type CondBroadcastCounter struct {
	c     sync.Cond
	count int
}

func main() {
	var cCounter CondBroadcastCounter
	cCounter.c = *sync.NewCond(&sync.Mutex{})

	increment := func() {
		for i := 0; i < 5; i++ {
			cCounter.c.L.Lock()
			cCounter.count++
			cCounter.c.Broadcast()
			fmt.Println("Incremented:", cCounter.count)
			time.Sleep(time.Millisecond * 500)
			cCounter.c.L.Unlock()
		}
	}

	checker := func() {
		for i := 0; i < 5; i++ {
			cCounter.c.L.Lock()
			for cCounter.count < 1 {
				cCounter.c.Wait()
			}
			fmt.Println("Greater than one! :", cCounter.count)

			cCounter.c.L.Unlock()
		}
	}

	cCounter.count++

	go increment()
	go checker()

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
