package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Channel notes from https://www.youtube.com/watch?v=f6kdp27TYZs

// Declaring and initializing.
// var chan int
// c = make(chan int)
// or
// c:= make(chan int)

// Sending to a Channel
// c <- 1

// Receiving from a Channel
// The "arrow" indicates the direction of data flow.
// value = <-c

func main() {
	channel1 := make(chan string)
	channel2 := boring2("boring")
	go boring("Boring message", channel1)
	for i := 0; i < 15; i++ {
		fmt.Printf("Response from boring func: %q\n", <-channel1) // receiving data from channel
		fmt.Printf("Boring 2: %q\n", <-channel2)
	}
	fmt.Println("Main func is done receiving data from boring func!")
}

func boring(msg string, channel1 chan string) {
	for i := 0; ; i++ {
		channel1 <- fmt.Sprintf("%s %d", msg, i) // appending data (string in this case) to the channel
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

/*
Boring2 has performs the exact same job, but creates its own channel before calling an
anonymous function ("func"). This function feeds data to the channel, but uses a loop to control
the number of iterations. Boring above just runs until func main is exited. Boring2 has a better
architecture. Also notice how the anonymous func ends with "()". This is one way to create and
call a function at the same time.
*/
func boring2(msg2 string) <-chan string {
	channel2 := make(chan string)
	go func() {
		for i := 0; i < 15; i++ {
			channel2 <- fmt.Sprintf("%s %d", msg2, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return channel2
}
