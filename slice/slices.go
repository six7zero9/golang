package main

import "fmt"

func main() {
	// make a slice and take a slice of it
	numSlice := []int{1, 2, 3, 4, 5}
	numSlice2 := numSlice[2:5]
	fmt.Println("numSlice2[2] = ", numSlice2[2])

	// appending to a slice and printing values
	numSlice2 = append(numSlice2, 12, 13, 14, 15)
	for index, value := range numSlice2 {
		fmt.Printf("Index: %d\tValue: %d\n", index, value)
	}
}
