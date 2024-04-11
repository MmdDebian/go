package main

import (
	"errors"
	"fmt"
	"math"
)

var Global int = 123
var AnotherGlobal = -5675

func inilizeValue(num int) (bool, error) {
	switch num {
	case 2:
		return false, nil
	case 3:
		return true, nil
	default:
		return false, errors.New("Error in number")
	}
}

func main() {
	var j int
	i := Global + AnotherGlobal
	fmt.Println("initial j value :", j)
	j = Global

	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global=%d , i=%d , j=%d , k=%.2f \n", Global, i, j, k)
	res, err := inilizeValue(1)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
