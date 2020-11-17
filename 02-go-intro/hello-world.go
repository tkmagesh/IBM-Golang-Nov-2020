package main

/* 
import fm "fmt"
import "os" 
*/

import (
	fm "fmt"
	//"os"
)
//

func init(){
	fm.Println("init triggered")
}

func main(){
	fm.Println("Hello World!")
}

//go run hello-world.go 