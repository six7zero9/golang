package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web1   = simSearch("web")
	web2   = simSearch("web")
	web3   = simSearch("web")
	image1 = simSearch("image")
	image2 = simSearch("image")
	image3 = simSearch("image")
	video1 = simSearch("video")
	video2 = simSearch("video")
	video3 = simSearch("video")
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
	c := make(chan Result)
	go func() { c <- First(query, web1, web2) }()
	go func() { c <- First(query, image1, image2) }()
	go func() { c <- First(query, video1, video2) }()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return results
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Printf("About %d results (%v)", len(results), elapsed)
}
