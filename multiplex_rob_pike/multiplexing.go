package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring2("Us"), boring2("Them"))
	for i := 0; i < 30; i++ {
		fmt.Println(<-c)
	}
}

func boring2(msg2 string) <-chan string {
	channel2 := make(chan string)
	go func() {
		for i := 0; i < 15; i++ {
			channel2 <- fmt.Sprintf("%s %d", msg2, i)
			time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		}
	}()
	return channel2
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
