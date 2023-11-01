package main

import "fmt"

func main() {
	var s []string

	fmt.Println("uninit:" , s , s == nil , len(s) == 0)
}