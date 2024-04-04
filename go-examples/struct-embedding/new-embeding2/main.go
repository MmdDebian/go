package main

import "fmt"

type calc interface {
	add() int
	minus() int
}

type values struct {
	num1, num2 int
}

func (v *values) add() int {
	return v.num1 + v.num2

}

func (v *values) minus() int {
	return v.num1 - v.num2
}

func res(c calc) {
	fmt.Println(c)
	fmt.Println(c.add())
	fmt.Println(c.minus())
}

func main() {
	value := values{
		num1: 1,
		num2: 3,
	}
	res(&value)
}