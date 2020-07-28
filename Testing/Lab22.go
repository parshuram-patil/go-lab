package main

import "fmt"

func add(i int, j int) int {
	fmt.Println("i = ", i)
	return i + j
}
func main() {
	x := add(20, 30)
	fmt.Println("x = ", x)
}
