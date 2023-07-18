package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var missionCompleted bool

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	var once sync.Once
	for i := 0; i < 100; i++ {
		go func() {
			if foundTreasure() {
				once.Do(markMissionCompleted)
			}
			wg.Done()
		}()

	}
	wg.Wait()
	checkMissionCompletion()
}

func checkMissionCompletion() {
	if missionCompleted {
		fmt.Println("Mission is now completed")
	} else {
		fmt.Println("Mission was a failure")
	}
}

func markMissionCompleted() {
	missionCompleted = true
	fmt.Println("Marking mission completed")
}

func foundTreasure() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}
