package main

import "fmt"
//import "./Employee"

/* type Employee struct {
	id int
	name string
	salary float64
}
 */
//update as a function
/* func update(e *Employee) {
	e.id = 200
	fmt.Println(e)
}
 */
//update as method
/* func (e *Employee) update() {
	e.id = 200
	fmt.Println(e)
} */

func main(){
	//emp := Employee{id : 100, name: "Magesh", salary: 10000}
	
	/* 
	emp := new(Employee)
	emp.id = 100
	emp.name = "Magesh"
	emp.salary = 10000 
	*/

	//This is the equivalent of above
	emp := &Employee{id : 100, name: "Magesh", salary: 10000}

	fmt.Println(emp)
	fmt.Println(emp.id) 
	

	//update(emp)
	emp.update()
	fmt.Println("After updating")
	fmt.Println(emp)
}



/* 
	& - address
	* - dereference
*/