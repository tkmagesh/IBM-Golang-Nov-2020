package main

import "fmt"

func main(){
	var choice int
	fmt.Println("Enter your choice ")
	fmt.Println("1. F to C")
	fmt.Println("2. F to meter")
	fmt.Scanf("%d", &choice)
	/* 
	if choice == 1 {
		convert_f_2_c()		
	} else {
		convert_f_2_m()
	} 
	*/
	fmt.Println(add(100,200))
	switch choice {
		case 1:
			convert_f_2_c()
		case 2:
			convert_f_2_m()
		default:
			fmt.Println("Invalid choice")
	}
}

func convert_f_2_c(){
	var value float64
	fmt.Println("Enter the farenheit for conversion")
	fmt.Scanf("%f", &value)
	fmt.Printf("Celcius=%v\n", (value - 32.0)*(float64(5)/float64(9)))
}

func convert_f_2_m(){
	var value float64
	fmt.Println("Enter the feet for conversion")
	fmt.Scanf("%f", &value)
	fmt.Printf("Meter=%v\n", value * 0.3048)
}

