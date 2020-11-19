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

 type Predicate func(p *Product) bool

 func filterProducts(products []*Product, predicate Predicate) []*Product {
	var result = []*Product{} 
	for _, product := range products {
		if predicate(product) { 
			result = append(result, product) 
		} 
	}
	return result; 
 }

 

func filterCostlyProducts(products []*Product) []*Product { 
	//product.cost >= 50 
	var result = []*Product{} 
	for _, product := range products {
		if product.cost >= 50 { 
			result = append(result, product) 
		} 
	}
	return result; 
} 

func filterUnderStockedProducts(products []*Product) []*Product { 
	//product.units < 40 
	var result = []*Product{} 
	for _, product := range products {
		if product.units < 40 { 
			result = append(result, product) 
		} 
	} 
	return result;
}

func print(products []*Product){
	for _, product := range products {
		fmt.Println(product.id, product.name, product.cost, product.units)
	}
	return
}

func negate(predicate Predicate) Predicate {
	return func(p *Product) bool {
		return !predicate(p)
	}
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
	fmt.Println("Products")
	print(products)
	fmt.Println();

	fmt.Println("Under stocked products [units < 40]")
	underStockedProductsPredicate := func(p *Product) bool{
		return p.units < 40
	}
	underStockedProducts := filterProducts(products, underStockedProductsPredicate)
	print(underStockedProducts)
	fmt.Println()

	fmt.Println("Well stocked products [!understocked]")
	/* 
	wellStockedProductsPredicate := func(p *Product) bool{
		return !underStockedProductsPredicate(p);
	} 
	*/
	wellStockedProductsPredicate := negate(underStockedProductsPredicate)
	wellStockedProducts := filterProducts(products, wellStockedProductsPredicate)
	print(wellStockedProducts)
	fmt.Println()


	fmt.Println("Costly products [cost >= 50]")
	costlyProductPredicate := func(p *Product) bool{
		return p.cost >= 50
	}
	costlyProducts := filterProducts(products, costlyProductPredicate)
	print(costlyProducts)
	
	fmt.Println()
	fmt.Println("Cheaper Products [!costly]")
	/* 
	cheaperProductPredicate := func(p *Product) bool{
		return !costlyProductPredicate(p)
	}; 
	*/
	cheaperProductPredicate := negate(costlyProductPredicate)
	cheaperProducts := filterProducts(products, cheaperProductPredicate)
	print(cheaperProducts)
	
	
	
}