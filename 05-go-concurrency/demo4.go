package main

import "fmt"

func fibonacci(n int, c chan int){
	x, y := 0, 1
	for i := 0; i < n; i++{
		c <- x
//		fmt.Printf("[fibonacci] Writing %d to channel\n", x)
		/* 
		t = x
		x = y
		y = t + y 
		*/
		x,y = y, x+y
	}
	close(c)
}

func main(){
	c := make(chan int, 10)
	go fibonacci(10, c)
	for i := range c{
//		fmt.Printf("[main] reading %d from channel\n", i)
		fmt.Println(i)
	}
}