package main

import (
	"fmt"
)

func main() {

	// slices
	nums := []int{2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum: ", sum)

	// index in slices
	// range on array, lices -> both index and value
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	// maps - key, value
	// range iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// maps - key only
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	// string - index, rune value
	// range iterates over unicode code points
	// i, c -> starting byte index of rune, rune itself (in int32)
	for i, c := range "go" {
		// fmt.Printf("%T\n", c) -> int32
		fmt.Println(i, c)
	}
}
