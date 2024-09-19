package main

import "math"

type Rectange struct {
	height float64
	weight float64
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (r Rectange) Area() float64 {
	return r.height * r.weight
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
