package main

import "fmt"

func main() {
	x, y := 4, 5
	fmt.Println("Before swap -> X:", x, " Y:", y)

	x, y = swap(x, y)
	fmt.Println("After swap -> X:", x, " Y:", y)

}

func swap(x, y int) (int, int) {
	return y, x
}
