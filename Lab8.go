package main

import (
	"emp"
)

func main() {
	empData := emp.Employee{}
	/* fmt.Println("Enter Employee Details")
	fmt.Scan(&empData.Empno)
	fmt.Scan(&empData.Ename)
	fmt.Scan(&empData.ESalary) */
	empData.Empno = 1
	empData.Ename = "Parshuram"
	empData.ESalary = 1234.5
	//emp.EmpList.Create(empData)
	empList := emp.EmpList{}
	empList.Create(empData)
	empData.Empno = 2
	empData.Ename = "Sidhant"
	empData.ESalary = 1111.5
	empList.Create(empData)
	empList.Print()
}
