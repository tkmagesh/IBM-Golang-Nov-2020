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

 type Product struct{
	 id int
	 name string
	 cost float64
	 units int
 }

func filterCostlyProducts(products []*Product) []*Product {
	/* fill in the blanks */
}


func main(){
	/* addWith100 := adder(100)
	fmt.Println(addWith100(200)) */

	/* nextEvenNumber := evenNoGenerator()
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())

	fn := evenNoGenerator() // fn will have its own "no" instance */

	products := []*Product{
		&Product{ id:5, name: "Pen", cost:10, units:100},
		&Product{ id:7, name: "Pencil", cost:5, units:50},
		&Product{ id:2, name: "Marker", cost:50, units:30},
		&Product{ id:6, name: "Scribble Pad", cost:30, units:40},
		&Product{ id:9, name: "Mouse", cost:100, units:10},
	}

	for _, product := range products {
		fmt.Println(product.id, product.name, product.cost, product.units)
	}

	
}