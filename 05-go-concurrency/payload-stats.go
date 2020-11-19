package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type PageSize struct {
	Url string
	Size int
}

func getPayload(url string, results chan PageSize){
	fmt.Println("initiating call for ", url)
	res, err := http.Get(url)
	if  err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	result := PageSize{
		Url : url,
		Size : len(bs),
	}
	fmt.Println(result)
	results <- result
}

func main(){
	urls := []string{
		"http://www.microsoft.com",
		"http://www.golang.org",
		"http://www.google.com",
		"http://www.apple.com",
	}

	results := make(chan PageSize)

	for _, url := range urls {
		go getPayload(url, results)
	}

	var biggest PageSize

	for range urls {
		fmt.Println("reading from results")
		result := <- results
		if (result.Size > biggest.Size) {
			biggest = result
		}
	}

	fmt.Println("The biggest home page is ", biggest)
}