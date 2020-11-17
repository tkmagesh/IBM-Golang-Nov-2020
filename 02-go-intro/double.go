package main

import "fmt"

func main(){
	fmt.Print("Enter a number : ")
	var input int
	fmt.Scanf("%d", &input)
	output := input * 2
	fmt.Println(output)
	if input % 2 == 0 {
		fmt.Println("it is even")
	} else {
		fmt.Println("it is odd")
	}
}