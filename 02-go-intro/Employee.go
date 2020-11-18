package main

import "fmt"

type Employee struct {
	id int
	name string
	salary float64
}

//update as a function
/* func update(e *Employee) {
	e.id = 200
	fmt.Println(e)
}
 */
//update as method
func (e *Employee) update() {
	e.id = 200
	fmt.Println(e)
}