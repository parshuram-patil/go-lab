package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sum: ", add(10, 30))
}

func add(x int, y int) int {
	return x + y
}
