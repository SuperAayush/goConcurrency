// Using the buffered channels in order to avoid deadlocking

package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "Channel 1"
	channel <- "Channel 2"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
