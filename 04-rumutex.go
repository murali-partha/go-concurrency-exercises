package main

import (
	"fmt"
	"sync"
	"time"
)

type RWSafeMessage struct {
	mu      sync.RWMutex
	message string
}

func main() {
	var sMessage RWSafeMessage

	sMessage.message = ""

	audiMessage := func() {
		sMessage.mu.RLock()
		defer sMessage.mu.RUnlock()
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("Audi is a%s car\n", sMessage.message)
	}

	ducatiMessage := func() {
		sMessage.mu.RLock()
		defer sMessage.mu.RUnlock()
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("Ducati is a%s car\n", sMessage.message)
	}

	updateMessage := func() {
		sMessage.mu.Lock()
		defer sMessage.mu.Unlock()
		fmt.Printf("Updating message to Fast\n")
		time.Sleep(time.Second * 2)
		sMessage.message = " Fast"
	}

	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			audiMessage()
		}()
	}

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ducatiMessage()
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		updateMessage()
	}()

	wg.Wait()
	fmt.Println("Messages complete.")
}
