package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func main() {
	file, err := os.Create("foo.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Foo Bar Baz Qux Quux Quuux")
	file.Close()

	stream, err := ioutil.ReadFile("foo.txt")
	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)
	fmt.Println("String from foo.txt", readString)
	fmt.Println("Stream reflect: ", reflect.TypeOf(stream))
	fmt.Println("ReadString reflect: ", reflect.TypeOf(readString))

}
