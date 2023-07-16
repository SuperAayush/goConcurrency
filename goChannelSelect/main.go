package main

import "fmt"

func main() {
	doc1, doc2 := make(chan string), make(chan string)

	go senior(doc1, "Senior Doctor")
	go senior(doc2, "Senior Doctor")

	equality()
}

func senior(doctor chan string, message string) {
	doctor <- message
}

func equality() {
	doc1 := make(chan interface{})
	close(doc1)
	doc2 := make(chan interface{})
	close(doc2)

	var doc1Chance, doc2Chance int

	for i := 0; i < 100; i++ {
		select {
		case <-doc1:
			doc1Chance++
		case <-doc2:
			doc2Chance++
		}
	}

	fmt.Printf("doc1Chance: %d, doc2Chance: %d,\n", doc1Chance, doc2Chance)
}
