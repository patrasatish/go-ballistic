package main

import "fmt"

type person struct {
	name string
	age  int
}

// go is a garbase collected language
// if any pointer is not actively referenced, will be cleaned up
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// creates a new struct
	fmt.Println(person{"Bob", 20})

	// name the fields
	fmt.Println(person{name: "Alice", age: 30})

	// Ommitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// & prefix -> yields a ptr to struct
	fmt.Println(&person{name: "Linda", age: 40})
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sonny", age: 18}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	// can use the pointers
	sp.age = 51
	fmt.Println(sp)
	fmt.Println(s)

	// struct used for a single value
	// no need to give it a name
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
