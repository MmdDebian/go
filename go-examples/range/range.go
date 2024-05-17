package main

import (
	"fmt"
)

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum : ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index : ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		// use range for key and value
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	usernames := map[string]int{"username": 15}

	for k, v := range usernames {
		print(k, v)
	}
}
