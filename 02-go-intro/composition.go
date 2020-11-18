package main

import "fmt"

type Person struct {
	id int
	name string
}

type Employee struct {
	Person
	
	salary float64 
	location string
}

func main() {
	emp := Employee{ Person : Person{ id : 100, name :"Magesh" } , salary : 10000 }
	//emp := Employee{}
	/* 
	emp.Person.id = 100
	emp.Person.name = "Magesh"  
	*/

	/* 
	emp.id = 100
	emp.name = "Magesh" 
	*/
	
	//emp.salary = 10000

	fmt.Println(emp.Person.id, emp.name, emp.salary)
}