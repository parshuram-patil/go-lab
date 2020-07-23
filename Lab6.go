package main

import (
	"fmt"
	"second"
)

func main() {
	x, y := 4, 3
	fmt.Println("Before swap -> X:", x, " Y:", y)

	add, sub := second.Calc(x, y)
	fmt.Println("Add:", add, " SUB:", sub)

}
