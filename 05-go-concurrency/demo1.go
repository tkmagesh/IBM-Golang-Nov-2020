package main

import (
	"fmt"
	"time"
)

func say(str string){
	for i:=0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str)
	}
}

func main(){
	go say("World")
	say("Hello")
}

