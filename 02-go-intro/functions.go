package main

import "fmt"

func main(){
	fmt.Println(add(100,200))
	var result, _ = divide(100, 7)
	/* fmt.Printf("%v, %v\n", result, remainder) */
	fmt.Printf("%v\n", result)
}

func add(x int, y int) (result int) {
	return x + y;
}

/* func divide(x int, y int) (result int, remainder int) {
	result = x / y
	remainder = x % y
	return result, remainder
} */

func divide(x int, y int) (int, int) {
	result := x / y
	remainder := x % y
	return result, remainder
}