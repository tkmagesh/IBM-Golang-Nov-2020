package main

import (
	"fmt"
	"time"
)

func say(str string, x chan int, y chan int){
	for i:=0; i < 5; i++ {
		<- y
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str)
		x <- i
	}
}

func main(){
	c1 := make(chan int)
	c2 := make(chan int)
	//c1 <- -1
	c2 <- -1
	go say("World", c1, c2)
	//go say("Hello", c1, c2)
	<- c1
	fmt.Println("End of main")
	time.Sleep(10000 * time.Millisecond)
}

