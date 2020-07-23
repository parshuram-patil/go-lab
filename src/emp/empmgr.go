package emp

import "fmt"

type EmpList struct {
	EmpArr [10]Employee
	cnt    int
}

func (emplist *EmpList) Print() {
	/* for _, emp := range EmpList.EmpArr {
		fmt.Println(emp)
	} */

	fmt.Println("\nEmployee List --->")

	for _, emp := range emplist.EmpArr {
		fmt.Println(emp)
	}
}

func (emplist *EmpList) Create(emp Employee) {
	emplist.EmpArr[emplist.cnt] = emp
	emplist.cnt++

	fmt.Println("\nEmplyee Created---> \n", emp)
}
