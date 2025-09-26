package main

import (
	"fmt"
	"sync"
)

type OnceCounter struct {
	counter int
	once    sync.Once
}

func main() {
	var oCounter OnceCounter

	var increments sync.WaitGroup
	increments.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer increments.Done()
			oCounter.once.Do(func() {
				oCounter.counter++
				fmt.Println("Counter incremented:", oCounter.counter)
			})
		}()
	}

	increments.Wait()

	fmt.Println("Total Count:", oCounter.counter)
}
