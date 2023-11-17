package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 15
	m["k2"] = 18

	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println("v1 :" , v1)

	v2 := m["k2"]
	fmt.Println("v2 :" , v2)
}