package main

import (
	"fmt"
	"time"
)

func main() {
	go count("hello", 5)
	count("world", 3)
}

func count(thing string, count int) {
	for i := 1; i <= count; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second * 1)
	}
}
