package main

import (
	"fmt"
	"golang.org/x/exp/maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 15
	m["k2"] = 18

	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println("v1 :" , v1)

	v2 := m["k2"]
	fmt.Println("v2 :" , v2)

	
	// Create and delete map : 


	// create 
	test := make(map[int]bool)

	test[15] = false 
	test[16] = true
	test[17] = true
	test[18] = true

	fmt.Println(test[15])


	// delete 
	delete(test , 15)

	fmt.Println(test)


	// check equal maps : 
	userList := map[string]int{"foo":1 , "bar" : 2}

	userList2 := map[string]int{"foo":1 , "bar" : 2}


	if maps.Equal(userList , userList2) {
		fmt.Println("this array can be equal .. .. ")
	}else {
		fmt.Println("this array can be NOT equal .. .. ")
	}
}