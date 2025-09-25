package main

import (
	"fmt"
	"sync"
	"time"
)

func buffCounter(countCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 15; i++ {
		fmt.Println("Writing data to channel")
		countCh <- i

	}
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan int, 5)

	wg.Add(1)
	go buffCounter(ch, &wg)

	for i := 0; i < 15; i++ {
		time.Sleep(time.Second * 1)
		fmt.Printf("Reading %d\n", <-ch)
	}
	wg.Wait()
}
