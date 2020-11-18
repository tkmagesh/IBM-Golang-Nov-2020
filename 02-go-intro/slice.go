package main

import "fmt"

func main(){
	//nos := []int{3, 1, 4, 2, 5, 7} //using the slice literal syntax
	
	/* 
	nos := make([]int, 6)
	nos[0] = 3
	nos[1] = 1
	nos[2] = 4
	nos[3] = 2
	nos[4] = 5
	nos[5] = 7 
	*/

	nos := []int{}
	
	//nos = append(nos, 3, 1, 4, 2, 5, 7)

	/* 
	nos = append(nos, 3, 1, 4)
	nos = append(nos, 2, 5, 7) 
	*/

	nos1 := []int { 3,1,4 }
	nos2 := []int { 2,5,7 }

	 
	nos = append(nos, nos1...)
	nos = append(nos, nos2...) 
	
	fmt.Printf("%v\n",nos)

	newNos := nos[1 : 4]
	fmt.Printf("%v\n",newNos)

	oneNo := nos[2:3]
	fmt.Printf("%v\n",oneNo)

	fmt.Printf("%v\n", nos[2:])
	fmt.Printf("%v\n", nos[:4])

	for idx, value := range nos {
		fmt.Println(value)
		if idx == 4 {
			continue
		}
	}

	var dummy []int
	dummy = append(dummy, 10,20,30 )

	//dummy := []int{}
	if dummy == nil {
		fmt.Println("dummy is empty")
	}
}