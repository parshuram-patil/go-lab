package main

import (
	"emp"
	"fmt"
)

func main() {
	fmt.Println("Start..")
	empData := emp.Employee{}
	empData.Empno = 1
	empData.Ename = "Parshuram"
	empData.ESalary = 1234.5
	emp.Create(empData)

}
