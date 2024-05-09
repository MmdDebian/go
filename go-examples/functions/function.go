package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)

	fmt.Println("1 + 2 = ", res)

	res = plusplus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)
	compileString("dedede")
}

func compileString(txt string ) {
	var arr [50]string

	for i := range txt {
		fmt.Println(i)
		arr[i] = "test"
	}

	fmt.Println(arr)
}
