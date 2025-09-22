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
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2 * (r.width * r.height)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type square struct {
	sideLength float64
}

func (sq square) area() float64 {
	return sq.sideLength * sq.sideLength
}

func (sq square) perimeter() float64 {
	return 4 * sq.sideLength
}

// Create method to use interface
func multiplyObj(g geometry) {
	fmt.Printf("Doubled %T Area: %.2f\n", g, 2*g.area())
}

func main() {
	rect := rect{width: 10, height: 15}
	square := square{sideLength: 100}
	circle := circle{radius: 54}

	multiplyObj(rect)
	multiplyObj(square)
	multiplyObj(circle)
	// renamed
}
