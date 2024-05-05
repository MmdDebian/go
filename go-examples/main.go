package main

import (
	"encoding/json"
	"fmt"
)

// type mytypes bool

// var isAdmin mytypes = true

type Person struct {
	Name string
	Age  int
}

func createJson() ([]byte, error) {
	person := Person{Name: "Alice", Age: 30}
	jsonByte, err := json.Marshal(person)

	if err != nil {
		return nil, err
	}

	return jsonByte, nil
}

func logger(data byte) {
	fmt.Println(data)
}

func main() {
	for index := 0; index <= 10; index++ {
		print(index, "\n")
	}

	if false {
		print("is admin")
	} else if 1 != 1 {
		print("True")
	} else {
		print("not admin")
	}

	jsonByte, err := createJson()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonByte))
}
