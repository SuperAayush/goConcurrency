package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock  sync.Mutex
	count int
)

func main() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Resulted Count is :", count)
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}
