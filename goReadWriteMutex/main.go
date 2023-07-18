package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwLock sync.RWMutex
)

func main() {
	readAndWrite()
}

func readAndWrite() {
	go read()
	go write()

	time.Sleep(5 * time.Second)
	fmt.Println("DONE")
}

func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("Read Locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Read Unlocking")
}

func write() {
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("Write Locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Write Unlocking")
}
