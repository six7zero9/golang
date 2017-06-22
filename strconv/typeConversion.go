package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	randoString := "hello world"

	// strings library funcs
	fmt.Println("\nWorking with 'strings' library\n")
	fmt.Println(strings.Contains(randoString, "lo"))
	fmt.Println(strings.Index(randoString, "lo"))
	fmt.Println(strings.Count(randoString, "l"))
	fmt.Println(strings.Replace(randoString, "l", "x", 4))

	listArray := []string{"c", "a", "b"}
	sort.Strings(listArray)
	fmt.Println("nSorted Letters: ", listArray)

	csvString := "1,2,3,4,5,6"
	fmt.Println("String to list with split: ", strings.Split(csvString, ","))

	numArray := []string{"2", "3", "1"}
	fmt.Println("List to string with join: ", strings.Join(numArray, ","))

	// data type conversion funcs
	fmt.Println("\n\nWorking with 'strconv' library")
	randoInt := 8
	intToFloat := float64(randoInt)
	fmt.Printf("\nInt to float: %f\ntype: %s\n\n", intToFloat, reflect.TypeOf(intToFloat))

	randoFloat := 23.23
	floatToInt := int(randoFloat)
	fmt.Printf("Float to int: %d\ntype: %s\n\n", floatToInt, reflect.TypeOf(floatToInt))

	randoString2 := "8011"
	stringToInt, err := strconv.ParseInt(randoString2, 0, 64)
	if err != nil {
		dontPanic()
	}
	fmt.Printf("String to int: %d\ntype: %s\n\n", stringToInt, reflect.TypeOf(stringToInt))

	randoString3 := "42.42"
	stringToFloat, err := strconv.ParseFloat(randoString3, 64)
	fmt.Printf("String to float: %f\ntype: %s\n\n", stringToFloat, reflect.TypeOf(stringToFloat))
	if err != nil {
		dontPanic()
	}

}

func dontPanic() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("PANIC")
}
