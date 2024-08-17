package main

type emplotee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name     string
	hourlyPy int
	hourYrly int
}

func (c contractor) getName() string {
	return ""
}
