package main

import (
	"fmt"
	"math/rand"
	"time"
)

var tester = test_func("foo")

type Result string
type Foobar func(thing string) Result

func test_func(bar string) Foobar {
	return func(thing string) Result {
		start := time.Now()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		elapsed := time.Since(start)
		return Result(fmt.Sprintf("%s time: %v", thing, elapsed))
	}
}

func main() {
	rs := tester("thingy")
	fmt.Println(rs)
}
