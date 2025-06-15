package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container embedding a base
// embedding -> a field without a name
type container struct {
	base
	str string
}

func main() {

	// explicit initialize embeddings
	// embedded type serves as field name
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// can access base's (embedded's) field / method directly
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// full path spelled out
	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// container implements  describer interface because it embeds base
	var d describer = co
	fmt.Println("describer:", d.describe())
}
