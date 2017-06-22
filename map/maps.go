package main

import "fmt"

func main() {

	presAge := make(map[string]int)

	presAge["Teddy"] = 42

	fmt.Println(presAge["Teddy"])

	presAge["JFK"] = 43

	fmt.Println(presAge["JFK"])

}
