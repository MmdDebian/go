package main

import "fmt"

type calc interface {
	add() float64
}

type input struct {
	oneNumber float64
	twoNumber float64
}


func (i input) add() float64 {
	return i.oneNumber + i.twoNumber
}

func testClac(c calc) {
	fmt.Println(c)
	fmt.Println(c.add())
}
