package main

import "fmt"

func main() {

	fmt.Println("Loop-1: Only condition")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("Loop-2: Initial/Condition/After")
	for j := 0; j <= 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("Loop-3")
	for i := range 3 {
		fmt.Println(i)
	}
	// i := range 3 => i := 0; i <= 3; i++

	fmt.Println("Infinite Loop and Break")
	for {
		fmt.Println("In loop, breaking away from it")
		break
	}

	fmt.Println("Only odds")
	for nn := range 6 {
		if nn%2 == 0 {
			continue
		}
		fmt.Println(nn)
	}
}
