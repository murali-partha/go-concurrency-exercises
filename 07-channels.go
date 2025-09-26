package main

import (
	"fmt"
	"sync"
	"time"
)

func counter(countCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 15; i++ {
		fmt.Println("Writing data to channel")
		countCh <- i

	}
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go counter(ch, &wg)

	for i := 0; i < 15; i++ {
		time.Sleep(time.Second * 1)
		fmt.Printf("Reading %d\n", <-ch)
	}
	close(ch)
	wg.Wait()
}
