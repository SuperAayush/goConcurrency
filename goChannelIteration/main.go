package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	go giveMedicine(channel)

	for {
		report, alive := <-channel
		if !alive {
			break
		}
		fmt.Println(report)
	}
}

func giveMedicine(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	totalRooms := rand.Intn(5)

	for i := 0; i < totalRooms; i++ {
		healed := rand.Intn(10)
		channel <- fmt.Sprint("Successful Heals: ", healed)
	}
	close(channel)
}
