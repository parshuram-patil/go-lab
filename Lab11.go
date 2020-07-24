package main

import (
	"fmt"
)

func main() {
	var marks map[string]int
	fmt.Println("Map=", marks)
	// marks["s1"] = 100 //this will give "assignment to entry in nil map"
	marks = make(map[string]int)
	fmt.Println(marks)
	marks["s1"] = 100
	marks["s2"] = 90
	marks["s3"] = 70
	fmt.Println("Map = ", marks, "Len = ", len(marks))
	x, avail := marks["s4"]
	fmt.Println("x = ", x, " avail= ", avail)
	x, avail = marks["s2"]
	fmt.Println("x = ", x, " avail = ", avail)

	name := make(map[string]string)
	name["s1"] = "s1"
	fmt.Println("Map = ", name, "Len = ", len(name))
	y, avail := name["s1"]
	fmt.Println("y = ", y, " avail = ", avail)
	y, avail = name["s2"]
	fmt.Println("y = ", y, " avail = ", avail)

}
