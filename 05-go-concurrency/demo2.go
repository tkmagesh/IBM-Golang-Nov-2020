package main

import "fmt"

func sum(nos []int, name string, c chan int){
	fmt.Printf("%s is triggered\n", name)
	sum := 0
	for _, n := range nos{
		sum += n
	}
	c <- sum
}

func main(){
	a := []int{ 8,3,1,6,4,-5 }
	c := make(chan int)
	nos1 := a[:len(a)/2]
	nos2 := a[len(a)/2 :] 
	go sum(nos1, "ins-1", c)
	go sum(nos2, "ins-2", c)
	fmt.Println("Attempting to read channel for x ")
	x := <- c
	fmt.Println("Attempting to read channel for y ")
	y := <- c
	fmt.Println(nos1, nos2, x + y)
	/* 
	a := []int{ 8,3,1,6,4,-5 }
	c1 := make(chan int)
	c2 := make(chan int)
	nos1 := a[:len(a)/2]
	nos2 := a[len(a)/2 :] 
	go sum(nos1, c1)
	go sum(nos2, c2)
	x := <- c1
	y := <- c2
	fmt.Println(nos1, nos2, x + y) 
	*/
}