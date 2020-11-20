package main

import (
	"fmt"
	"sync"
	"time"
)

var data = 10

var mutex = &sync.Mutex{}

func main(){
	go func(){
		mutex.Lock()
		defer mutex.Unlock()
		data += 10;
		//mutex.Unlock()
		fmt.Println(data)
		time.Sleep(time.Millisecond * 500)
	}()
	go func(){
		mutex.Lock()
		defer mutex.Unlock()
		data += 10;
		//mutex.Unlock()
		fmt.Println(data)
		time.Sleep(time.Millisecond * 500)
	}()
	time.Sleep(time.Millisecond * 1000)
}
