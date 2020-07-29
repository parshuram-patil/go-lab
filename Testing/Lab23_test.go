package main

import (
	"fmt"
	"testing"
)

/* func TestMain(t *testing.T) {
	main()
} */

func TestDivideSuccess(t *testing.T) {
	x := divide("10", "5")
	if x != 2 {
		t.Error("Invalid Sum")
	}
}

func TestDivideAtoi(t *testing.T) {
	x := divide("10a", "5")
	if x != 0 {
		t.Error("Invalid Sum")
	}
}

func TestDivideByZero(t *testing.T) {
	defer func() {
		fmt.Println("end of main")
		r := recover()
		if r != nil {

			fmt.Println("Recover status ", r)
		} else {
			t.Error("Invalid Divide - divide by zero exception not raiased")
		}
	}()
	x := divide("40", "0")
	if x != 0 {
		t.Error("Invalid Divide divide by zero")
	}
}
