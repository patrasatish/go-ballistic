package main

import "fmt"

type rect struct {
	width, height int
}

// method area -> receiver type of *rect (ptr receiver)
func (r *rect) area() int {
	return r.width * r.height
}

// method perim -> receiver type of rect (value receiver)
func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func perim(r rect) int {
	return 2 * (r.width + r.height)
}

func area(r rect) int {
	return r.width * r.height
}

func main() {
	r := rect{width: 10, height: 20}
	fmt.Println("Area By Method", r.area())
	fmt.Println("Perimeter By Method", r.perim())

	r2 := &r
	fmt.Println("Area By Method", r2.area())
	fmt.Println("Perimeter By Method", r2.perim())

	fmt.Println("Area By Function", area(r))
	fmt.Println("Perimeter By Function", perim(r))
}
