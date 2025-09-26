package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	message string
}

func RunMutexUnsafe() {
	var sMessage Message

	carMessage := func() {
		sMessage.message = "Car"
		time.Sleep(time.Second * 2)
		fmt.Printf("Audi is a %s\n", sMessage.message)
	}

	bikeMessage := func() {
		sMessage.message = "Bike"
		time.Sleep(time.Second * 2)
		fmt.Printf("Ducati is a %s\n", sMessage.message)
	}

	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			carMessage()
		}()
	}

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bikeMessage()
		}()
	}

	wg.Wait()
}

func main() {
	RunMutexUnsafe()
	fmt.Println("Messages complete.")
}
