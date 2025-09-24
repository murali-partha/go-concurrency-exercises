package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("hello", 15)
		wg.Done()
	}()
	count("world", 3)
	wg.Wait()
}

func count(thing string, count int) {
	for i := 1; i <= count; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second * 1)
	}
}
