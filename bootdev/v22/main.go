package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name     string
	hourlyPy int
	hourYrly int
}

type fullTime struct {
	name   string
	salary int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPy * c.hourYrly
}

func (ft fullTime) getName() string {
	return ft.name
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("========================================")
}

func main() {
	test(fullTime{
		name:   "0x",
		salary: 50000,
	})
	test(contractor{
		name:     "atanda",
		hourlyPy: 100,
		hourYrly: 73,
	})
}
