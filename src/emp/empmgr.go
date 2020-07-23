package emp

import "fmt"

func Create(emp Employee) {
	printEmp(emp)
	//fmt.Println("EMP Create()")
}

func printEmp(emp Employee) {
	fmt.Println("Emplyee Details --> ", "\nEmpno: ", emp.Empno, " \nEname:", emp.Ename, "\nESalary", emp.ESalary)
}
