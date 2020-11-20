//Reading the contents of an entire file in a string:
// read_write_file.go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "emp.dat"
	outputFile := "emp_2.dat"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0x777)
	if err != nil {
		panic(err.Error())
	}
}
