package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var pizzaNum = 0
var pizzaName = ""
var wg sync.WaitGroup

func makeDough(stringChan chan string) {
	pizzaNum++

	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make Dough and Send for Sauce")

	stringChan <- pizzaName

}

func addSauce(stringChan chan string) {
	pizza := <-stringChan

	fmt.Println("Add Sauce and Send", pizza, "for Toppings")

	stringChan <- pizzaName

}

func addToppings(stringChan chan string) {
	pizza := <-stringChan

	fmt.Println("Add Toppings and Send", pizza, "for Baking")

	stringChan <- pizzaName

}

func bakePie(stringChan chan string) {
	defer wg.Done()
	pizza := <-stringChan

	fmt.Println("Bake", pizza, "and ship")

}

func main() {

	stringChan := make(chan string)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go makeDough(stringChan)
		time.Sleep(time.Millisecond * 10)
		go addSauce(stringChan)
		time.Sleep(time.Millisecond * 10)
		go addToppings(stringChan)
		time.Sleep(time.Millisecond * 10)
		go bakePie(stringChan)
		wg.Wait()
	}
	close(stringChan)
}
