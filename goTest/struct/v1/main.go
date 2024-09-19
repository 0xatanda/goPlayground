package main

import "math"

type Rectange struct {
	height float64
	wight  float64
}

type Circle struct {
	Radius float64
}

func (r Rectange) Area() float64 {
	return (r.height * r.wight)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rec Rectange) float64 {
	return 2 * (rec.height + rec.wight)
}
