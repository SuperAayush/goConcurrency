// This file contains basic case of concurrency where in a hospital every patient is attended by doctors in very less time.
package main

import (
	"fmt"
	"time"
)

func main() {

	// To execute the process in minimum time we need to be aware of the total time spent in the completion of the process
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	sickPatients := []string{"John", "Charles", "Doe", "Brat"}

	// After creating the heal function we will heal each of the patient
	for _, patient := range sickPatients {

		/* If we don't use gorountine here it will always take 4 secs to execute but after adding `go` in front of the below line of code it will execute very fast that sometimes it may even
		skip some process so to overcome this we add an extra time of 2 seconds after the heal function so that each patient is healed*/
		go heal(patient)
	}

	time.Sleep(time.Second * 2)
}

func heal(cure string) {
	fmt.Println("Attending the patient", cure)
	// We will make this process last for 1 sec to indictate that the curing process takes 1 sec
	time.Sleep(time.Second)

}
