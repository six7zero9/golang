package main

import "fmt"

func main() {

	const me = "michael jackson"

	var foo = 1 + 2

	fmt.Println("foo bar")
	fmt.Println("practicing go sytnax!")
	fmt.Printf("math: 1 + 2 = %d\n", foo)

	fmt.Println(me)

	fmt.Println("for loop with manual increment")
	var i = 0

	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("for loop with one line initialize, condition, and increment")
	for j := 0; j < 10; j += 2 {
		fmt.Printf("%d\n", j)
	}

}
