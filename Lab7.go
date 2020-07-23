package main

import (
	"fmt"
)

func shift(v Vertex) Vertex {
	v.X += 10
	v.Y += 10
	return v
}

type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{}
	fmt.Println("Current Vertex: ", v)
	fmt.Println("Please enter Vertex co-ordinates")
	/* no, err := fmt.Scanf("%d%d", &v.X, &v.Y)
	fmt.Println(no, "  ", err) */
	fmt.Scanf("%d%d", &v.X, &v.Y)
	fmt.Println("Changed Vertex: ", v)
	v1 := shift(v)
	fmt.Println("Shifted Vertex: ", v1)
}
