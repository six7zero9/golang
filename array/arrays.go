package main

import "fmt"

func main() {
	// manually assign values to array
	var favNums [6]float64

	favNums[0] = 123
	favNums[1] = 1234
	favNums[2] = 3.1415
	favNums[3] = 2.71828182
	favNums[4] = 23
	favNums[5] = 35.01

	fmt.Println(favNums[2])

	// declare and assign values in one line
	favNums2 := [4]float64{1, 2, 3, 4}
	for value := range favNums2 {
		fmt.Println(value)
	}

	// grabbing index and value from array
	for index, value := range favNums {
		fmt.Printf("Index: %d\tValue: %.3f\n", index, value)
	}

}
