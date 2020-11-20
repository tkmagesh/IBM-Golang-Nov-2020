
/*
os.O_RDONLY : the read flag for read-only access
os.WRONLY : the write flag for write-only access
os.O_CREATE : the create flag: create the file if it doesn’t exist 
os.O_TRUNC : the truncate flag: truncate to size 0 if the file already exists
*/

package main

import (
	"os"
	"bufio"
	"fmt"
)

func main () {
	// var outputWriter *bufio.Writer
	// var outputFile *os.File
	// var outputError os.Error
	// var outputString string
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return  
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	
	for i:=0; i<10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}