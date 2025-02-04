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
	close(doneChan) //line only makes sense if we know this function will take the longest
}
func main () {
	// dones := make([]chan bool, 4)
	// dones[0] = make(chan bool)
	// dones[1] = make(chan bool)
	// dones[2] = make(chan bool)
	// dones[3] = make(chan bool)

	done := make(chan bool)
	go greet("hi", done)
	go greet("hi hi",  done)
	go greet("hi hi hi",  done)
	go slowGreet("Hello", done)

	for doneChan := range done {
		fmt.Println(doneChan)
	}
	//as is, this results in a race condition since whatever function finishes 
	//first will send data to the channel and exit the program
	// for _, done := range dones {
	// 	<- done
	// }

}