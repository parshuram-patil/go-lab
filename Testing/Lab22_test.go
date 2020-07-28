package main

import "testing"

/* func TestMain(t *testing.T) {
	main()
} */

func TestAdd1(t *testing.T) {
	x := add(10, 30)
	if x != 40 {
		t.Error("Invalid Sum")
	}
}

func TestAdd2(t *testing.T) {
	x := add(20, 30)
	if x == 40 {
		t.Error("Invalid Sum")
	}
}
