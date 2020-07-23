package main

import "fmt"

func main() {
	x, y := 4, 3
	fmt.Println("Before swap -> X:", x, " Y:", y)

	add, sub := calc(x, y)
	fmt.Println("Add:", add, " SUB:", sub)

}

func calc(x, y int) (add, sub int) {
	add = x + y
	sub = x - y

	return
}
