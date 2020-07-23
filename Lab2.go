package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Current Executable Name: ", strings.Split(os.Args[0], "\\"))
	//fmt.Println("Current Executable Name: ", strings.Split(os.Args[0], "\")[len(os.Args[0]) - 1]))

	totalStrLength := 0
	for i := range os.Args {
		//fmt.Println(os.Args[i], " ")
		totalStrLength += len(os.Args[i])
	}

	fmt.Println("Total Lenght: ", totalStrLength-len(os.Args[0]))

}
