package main

import (
	"fmt"
)


func main()  {
	// i := 30


	for i := 0 ; i <= 50 ; i++ {
		switch i {
		case 1:
			fmt.Println("your number is 1")
		case 2:
			fmt.Println("your number is 2")
		case 50 : 
			fmt.Println("you number is 50")
	}
	}

	// printMyType(i)
	printMyType("hello")
	printMyType(false)
}


func printMyType(i interface{}){
	switch i.(type) {
		case bool:
			fmt.Println("you are boolean")
		case string :
			fmt.Println("you are string")
		case int : 
			fmt.Println("you are integer")
		default : 
			fmt.Println("non above type")			
	}
}