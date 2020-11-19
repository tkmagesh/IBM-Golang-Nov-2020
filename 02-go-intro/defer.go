package main

import "fmt"

func main(){
	f1()
}

func f1(){
	defer f2()
	fmt.Println("f1 is invoked")
}

func f2(){
	fmt.Println("f2 is invoked")
}


