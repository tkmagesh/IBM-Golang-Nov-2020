package main

import "fmt"

/* func adder() func(int, int) int {
	add := func(x int, y int) int {
		return x + y;
	}
	return add
} */

/* func adder(x int) func(int) int {
	add := func(y int) int {
		return x + y;
	}
	return add
}
 */

 func evenNoGenerator() func() int {
	 no := int(0)
	 return func()(int){
		 no += 2
		 return no
	 }
 }
func main(){
	/* addWith100 := adder(100)
	fmt.Println(addWith100(200)) */

	nextEvenNumber := evenNoGenerator()
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
}