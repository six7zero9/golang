package main

import "fmt"

func main() {
	x := 0

	fmt.Println("Not using a pointer to reference x (value of x remains 0)")
	changeXVal(x)
	fmt.Println("x =", x)

	fmt.Println("Using a pointer to change the value of x (uses '*')")
	changePointerXVal(&x)
	fmt.Println("x =", x)
	fmt.Println("Memory address for x =", &x)

	fmt.Println("Another way to user pointers (similar to example above), but uses *y and new(int)")
	y := new(int)
	changePointerYVal(y)
	fmt.Println("y =", *y)
	fmt.Println("Memory address for y=", &y)

	fmt.Println("\nRunning pointer example from golang.org")
	example()
}

// stupid warning ineffectual assigment to x (ineffassign)
func changeXVal(x int) {
	x = 2
}

func changePointerXVal(x *int) {
	*x = 2
}

func changePointerYVal(y *int) {
	*y = 200
}

func example() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
