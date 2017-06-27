package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web   = simSearch("web")
	image = simSearch("image")
	video = simSearch("video")
)

type Result string
type Search func(query string) Result

func simSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	results = append(results, web(query))
	results = append(results, image(query))
	results = append(results, video(query))
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Printf("About %d results (%v)", len(results), elapsed)
}
