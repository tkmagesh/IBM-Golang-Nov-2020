package main

import (
	"fmt"
	"time"
	"sync"
)

var wg = &sync.WaitGroup{}

func main(){
	wg.Add(3)
	go f1()
	go f2()
	go f3()
	wg.Wait()
	fmt.Println("All job done!")
}
func f1(){
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("f1 completed")
	wg.Done()
}

func f2(){
	time.Sleep(time.Millisecond * 3000)
	fmt.Println("f2 completed")
	wg.Done()
}

func f3(){
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("f3 completed")
	wg.Done()
}
