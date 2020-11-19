package main

/* 
import fm "fmt"
import "os" 
*/

import (
	fm "fmt"
	"modularity/utils"
	//"os"
)
//

func init(){
	fm.Println("init triggered")
}

func main(){
	fm.Println(utils.Greet())
}

//go run hello-world.go 