package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return (r.width + r.height) * 2
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius // Pi * r^2
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return (r.width + r.height) * 2
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	r2 := rectangle{width: 2, height: 2}

	measure(r)
	measure(c)
	measure(r2)

	myPrinter("Hola", 88, true)

	printType("Hola")
	printType(88)
	printType(true)
}

func measure(g geometry) {
	fmt.Printf("%T\n", g)
	fmt.Printf("%.4v\n", g.area())
	fmt.Printf("%.4v\n", g.perimeter())
}

func myPrinter(i ...interface{}) {
	for _, s := range i {
		fmt.Println(s)
	}
}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown")
	}
}
