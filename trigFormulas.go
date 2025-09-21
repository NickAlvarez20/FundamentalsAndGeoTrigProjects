package main

import (
	"fmt"
	"math"
)

type Triangle struct {
	Sides  [3]float64 // a,b,c
	Angles [3]float64 // A, B, C in radians
	Name   string
}

func (t Triangle) cosB() float64 {
	return t.Sides[1] / t.Sides[2]
}

func (t Triangle) sinB() float64 {
	return t.Sides[0] / t.Sides[2]
}

func (t Triangle) tanB() float64 {
	return t.Sides[0] / t.Sides[1]
}

func (t Triangle) lawCosines() float64 {
	a := t.Sides[0]
	b := t.Sides[1]

	//Convert degrees to radians: radians = degrees * (Pi/180)
	angleC := t.Angles[2] * math.Pi / 180

	return math.Sqrt(a*a + b*b - 2*a*b*math.Cos(angleC))
}

func (t Triangle) lawSinesFindSideA() float64 {
	sinA := math.Sin(t.Angles[0] * math.Pi / 180)
	sinB := math.Sin(t.Angles[1] * math.Pi / 180)
	b := t.Sides[1]

	// Finds side a
	return b * (sinA / sinB)
}

func (t Triangle) lawTangents() float64 {
	// Sides b
	b := t.Sides[1]
	// Angles A & B in radians
	angleA := t.Angles[0] * math.Pi / 180
	angleB := t.Angles[1] * math.Pi / 180

	tanDiff := math.Tan((angleA - angleB) / 2)
	tanSum := math.Tan((angleA + angleB) / 2)

	// Avoid division by zero
	if math.Abs(tanSum) < 1e-10 {
		return math.NaN()
	}

	ratio := tanDiff / tanSum

	// Formula for law of Tangents (a-b)/(a+b) = tan[(angle(A)-angle(B))/2] / tan[(angle(A)+angle(B))/2]
	return b * (1 + ratio) / (1 - ratio)

}

func main() {
	rightTriangle := Triangle{Sides: [3]float64{3, 4, 5},
		Angles: [3]float64{36.87, 53.13, 90}, Name: "right"}

	fmt.Println("Cosine: ", rightTriangle.cosB())
	fmt.Println("Sine: ", rightTriangle.sinB())
	fmt.Println("Tangent: ", rightTriangle.tanB())
	fmt.Println("Law of Cosines: ", rightTriangle.lawCosines())
	fmt.Println("Law of Sines-Finding Side A: ", rightTriangle.lawSinesFindSideA())
	fmt.Println("Law of Tangent: ", rightTriangle.lawTangents())
}
