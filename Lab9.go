package main

import (
	"fmt"
)

func main() {
	primeArray := [10]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primeArray len = ", len(primeArray), ", cap = ", cap(primeArray))

	var primeSlice []int = primeArray[:]
	fmt.Println("\n\nprimeSlice with [:]", primeSlice, "\n len: ", len(primeSlice), ", cap = ", cap(primeSlice))

	primeSlice = primeArray[5:7]
	fmt.Println("\n\nprimeSlice with [5:7]", primeSlice, "\n len: ", len(primeSlice), ", cap = ", cap(primeSlice))

	primeArray[6] = 99
	//primeSlice[1] = 88
	fmt.Println("\n\nprimeSlice [5:7] after modifiying array ", primeSlice, "\n len: ", len(primeSlice), ", cap = ", cap(primeSlice))

	primeArray[1] = 10
	primeSlice[1] = 88

	fmt.Println("\n\nAfter modifying bothe Array and Slice")
	fmt.Println("\nprimeSlice [5:7] after modifiying array and slice ", primeSlice, "\n len: ", len(primeSlice), ", cap = ", cap(primeSlice))
	fmt.Println("\nprimeArray after modifiying array and slice ", primeArray, "\n len: ", len(primeArray), ", cap = ", cap(primeArray))

}
