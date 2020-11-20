package main

import (
	//"bufio"
	"fmt"
	"io"
	"os"
	"encoding/csv"
	"strconv"
)

var (
	id int
	name string
	salary float64
)

func main() {
	inputFile, inputError := os.Open("emp.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" + 
		           "Does the file exist?\n" + 
			       "Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := csv.NewReader(inputFile)
	sum := float64(0)
	count := 0
	for {
		data, readerError := inputReader.Read()
		if readerError == io.EOF {
			break // error or EOF
		}
		sal, _ := strconv.ParseFloat(data[2],64)
		sum += sal
		count++
	}
	fmt.Println(sum/float64(count))
}
