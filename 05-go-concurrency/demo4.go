package main

import "fmt"

func fibonacci(n int, c chan int){
	x, y, t := 0, 1, 0
	for i := 0; i < n; i++{
		c <- x
		t = x
		x = y
		y = t + y
	}
	close(c)
}

func main(){
	c := make(chan int, 10)
	go fibonacci(10, c)
	for i := range c{
		fmt.Println(i)
	}
}