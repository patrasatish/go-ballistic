package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// to implement an interface in Go, just implement all methods in the interface
func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// can call all methods in the named interface - geometry by all implementors
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// can detect the runtime type of an interface value
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
		return
	}
	fmt.Println("Not a circle")
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// rect, circle both implement geometry, so can use instances of these structs
	// as arguments to measure
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
