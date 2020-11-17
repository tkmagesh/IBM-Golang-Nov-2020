package main

import "fmt"

func main(){
	var nos [5]int
	nos[0] = 10
	nos[1] = 40
	nos[2] = 50
	nos[3] = 70
	nos[4] = 100

	/* for x :=0; x < len(nos); x++ {
		fmt.Println(nos[x])
	} */

	for _, value := range nos {
		fmt.Printf("%v\n", value)
	}
}