package main

import "fmt"

func main() {
	defer printTwo()
	printOne()

	numArray := []float64{1, 2, 3, 4, 5, 6}
	fmt.Println("\nSum = ", sumThem(numArray))

	num1, num2 := returnTwoValues(10)
	fmt.Printf("\nTwo Returned Values: %d, %d\n", num1, num2)

	fmt.Printf("\nMultiple argument function results: %d\n", multiArgSubtract(1, 2, 3, 4, 5, 6))

	fmt.Printf("\nRecursive Factorial Function: %d\n", factorial(3))

	fmt.Println("Using Recover with defer to rescue errors")
	fmt.Printf("\n3 %% 0 = %d\n", safeDivision(3, 0))
	fmt.Printf("\n4 %% 2 = %d\n", safeDivision(4, 2))

	dontPanic()

}

func printOne() { fmt.Println("\n", 1) }
func printTwo() { fmt.Println("\n", 2) }

func sumThem(nums []float64) float64 {
	sum := 0.0
	for _, value := range nums {
		sum += value
	}
	return sum
}

func returnTwoValues(num int) (int, int) {
	return num + 1, num + 2
}

func multiArgSubtract(args ...int) int {
	result := 0
	for _, value := range args {
		result -= value
	}
	return result
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func safeDivision(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

func dontPanic() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("PANIC")
}
