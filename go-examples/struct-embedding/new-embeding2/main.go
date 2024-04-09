package main

import (
	"fmt"
)

type calc interface {
	add() int 
	minus() int
}

type values struct {
	num1, num2 int
}

type values1 struct {
	num3 int  
}

func (v values) add() int {
	return v.num1 + v.num2
}

func (v values) minus() int {
	return v.num1 - v.num2
}


func (v values1) add() int {
	return v.num3 + v.num3
}

func (v values1) minus() int {
	return v.num3 + 1 - v.num3
}


func res(c calc) {
	fmt.Println(c)
	fmt.Println(c.add())
	fmt.Println(c.minus())
}


func main() {
	a := values{num1: 1 ,num2: 3}
	b := values1{num3:  4}
	res(a)
	res(b)
}