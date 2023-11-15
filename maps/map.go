package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 15
	m["k2"] = 18

	fmt.Println(m)
}