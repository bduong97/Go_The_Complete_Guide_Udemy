package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool ) {
	fmt.Println("Hello")
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep((3 * time.Second))
	fmt.Println("Slow greet", phrase)
	doneChan <- true //sending data to bool channel
}
func main () {
	done := make(chan bool)
	go greet("hi", done)
	go greet("hi hi",  done)
	go greet("hi hi hi",  done)
	go slowGreet("Hello", done)
	//as is, this results in a race condition since whatever function finishes 
	//first will send data to the channel and exit the program
	<- done //waiting for data to come out of the channel
}