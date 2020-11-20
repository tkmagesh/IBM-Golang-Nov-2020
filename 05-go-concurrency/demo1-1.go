package main

import (
	"fmt"
	"time"
	"sync"
)

var data string 
var mutex = &sync.Mutex{}

func say(str string){ 
	for { 
		if data != str { 
			mutex.Lock() 
			fmt.Println(str) 
			data = str 
			mutex.Unlock() 
		} else { 
			time.Sleep(500 * time.Millisecond) 
		} 
	} 
} 


func main(){
	go say("Hello")
	go say("World")
	
	var input string
    fmt.Scanln(&input)
}

