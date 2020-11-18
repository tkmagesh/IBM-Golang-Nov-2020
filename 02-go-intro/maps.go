package main

import "fmt"

func main(){
	/* map[keyType]valueType */
	states := map[string]string {
		"KA" : "Bengaluru",
		"TN" : "Chennai",
		"KL" : "Trivandrum",
		"AP" : "Amaravathi",
		"TS" : "Hyderabad",
	}

	fmt.Println(states)

	//accessing by key
	fmt.Println(states["KA"])

	//add a new item
	states["OD"] = "Bhubhaneswar"
	fmt.Println(states)

	//remove an item
	delete(states, "KA")
	fmt.Println(states)

	//check if a key exists
	//if ok is true the item exists, false otherwise
	cap, ok := states["OD"]
	fmt.Println(cap, ok)
}