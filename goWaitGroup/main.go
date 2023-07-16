package main

import (
	"fmt"
	"sync"
)

func main() {
	var keepRunning sync.WaitGroup
	sickPatients := []string{"John", "Carl", "Jad", "Tony"}
	keepRunning.Add(len(sickPatients))
	for _, patient := range sickPatients {
		go heal(patient, &keepRunning)
	}

	keepRunning.Wait()
	fmt.Println("All patients healed")
}

func heal(patient string, keepRunning *sync.WaitGroup) {
	fmt.Println("Healed the patient: ", patient)
	keepRunning.Done()
}
