package main

import "fmt"


func main(){
	if true {
		fmt.Println("chacking one thing right (here the executed .)")
	}

	if 1 == 5 {
		fmt.Println("checking two number (here not executed .)")
	} else {
		fmt.Println("Yes . this runned")	
	}


	if false {
		print("next 1")
	}else if true {
		print("next 1")
	} else if 1 != 2 {
		print("next")
	} else {
		print("next else")
	}
}


func scanComment(){
	
}