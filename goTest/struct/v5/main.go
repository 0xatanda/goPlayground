package main

import "math"

type Rectangle struct {
	Height float64
	Weight float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area()
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Weight
}
