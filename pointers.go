package main

import "fmt"

// ival = a distinct copy of the one in the calling function
func passbyval(ival int) {
	ival = 0
}

// *iptr -> dereferences the pointer from its memory address
// to current value at that address
func passbyptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial: ", i)

	passbyval(i)
	fmt.Println("pass by val: ", i)

	passbyptr(&i)
	fmt.Println("pass by ptr: ", i)

	// &i -> gives memory address of i
	fmt.Println("pointer: ", &i)
}
