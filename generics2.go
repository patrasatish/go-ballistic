package main

import "fmt"

type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Add[T Num](a T, b T) T {
	return a + b
}

func main() {
	result := Add(1, 2)
	fmt.Printf("result: %+v\n", result)
	result2 := Add(1.1, 2.1)
	fmt.Printf("result: %+v\n", result2)
}
