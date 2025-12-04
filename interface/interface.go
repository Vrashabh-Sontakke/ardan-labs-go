package main

import "fmt"

// interface
type Shape interface {
	area() float32
	perimeter() float32
}

// struct to implement interface
type Rectangle struct {
	length, breadth float32
}

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

// use struct to implement area() of interface
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

func (s Square) area() float32 {
	return s.side * s.side
}

func (c Circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

// perimeter method
func (r Rectangle) perimeter() float32 {
	return (r.length + r.breadth) * 2
}

func (s Square) perimeter() float32 {
	return s.side * 4
}

func (c Circle) perimeter() float32 {
	return 3.14 * c.radius * 2
}

// access method of the interface
func calculate(s Shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

// main function
func main() {

	// assigns value to struct members
	rect := Rectangle{7, 4}

	// call calculate() with struct variable rect
	calculate(rect)

	square := Square{5}

	calculate(square)

	circle := Circle{5}

	calculate(circle)
}
