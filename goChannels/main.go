package main

import (
	"fmt"
	"time"
)

func main() {

	timeSpend := time.Now()
	defer func() {
		fmt.Println(time.Since(timeSpend))
	}()

	channel := make(chan bool)
	sickPatient := "John"
	go heal(sickPatient, channel)
	fmt.Println(<-channel)
}

func heal(cure string, healed chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Attending the patient", cure)
	healed <- true
}
