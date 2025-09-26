package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMessage struct {
	mu      sync.Mutex
	message string
}

func main() {
	var sMessage SafeMessage

	carMessage := func() {
		sMessage.mu.Lock()
		defer sMessage.mu.Unlock()
		sMessage.message = "Car"
		time.Sleep(time.Second * 1)
		fmt.Printf("Audi is a %s\n", sMessage.message)
	}

	bikeMessage := func() {
		sMessage.mu.Lock()
		defer sMessage.mu.Unlock()
		sMessage.message = "Bike"
		time.Sleep(time.Second * 1)
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
	fmt.Println("Messages complete.")
}
