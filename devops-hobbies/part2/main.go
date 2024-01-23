package main

import "fmt"

func printArrays(){
	// array
	var a [5]int32

	// short declartion 
	

	b := [10]int8{
		1,2,3,4,
	}

	fmt.Println(b)

	a[3] = 3
	fmt.Println(a)


	var twoD [2][3]int 

	for index1 := 0; index1 < 2; index1++ {
		for index2 := 0; index2 < 3; index2++ {
			twoD[index1][index2] = index1 + index2 	
		}
	}

	fmt.Println(twoD)
}



func printSlice(){
	s := make([]int , 3, 4)

	s[0] = 0 
	s[1] = 1 
	s[2] = 2 

	
	fmt.Printf("%p\n" , s)	
	fmt.Println(len(s))
	fmt.Println(cap(s))



	s = append(s, 3)
	fmt.Printf("%p\n" , s)	
	fmt.Println(len(s))
	fmt.Println(cap(s))


	s = append(s, 4)
	fmt.Printf("%p\n" , s)	
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func printMap(){
	m := make(map[string]int)

	m["first"] = 1 
	m["second"] = 2 

	m["first"] = 112 

	delete(m , "second")

	fmt.Println(m)
}


func printLoopSlice(){
	names := []string{"mohamad","milad" , "erfan"}

	for index, value := range names {
		fmt.Println(index)
		fmt.Println(value)
	}
}



func printloopMap(){
	message := map[string]string{
		"ok" : "This is ok" ,
		"error" : "This is error",
	}


	for key, value := range message {
		fmt.Println(key , value)
	}
}




// One input 
func myfunc(num1 int , message  string ) int {
	return 1 
}


// two and unlimited type input 
func myfunc1(num1 int , message  string ) (int , string ) {
	return 1 , "test"
}


// plus with 2 
func plusWith2(num int)(int , string ){
	return num + 2 , "successfully"
}


// closure
func closureFunc()func() int {
	index := 5 

	return func () int {
		index++ 
		return index 
	}
}


// fac
func fact(num int )int {
	if num == 0 {
		return 1  
	}

	return num *
	fact(num-1)
}

func passByValue(val int){
	val = 0 
}


func passByPointer(val *int){
	*val = 0 
}

func main() {
	value := 1 

	fmt.Printf("initial: %d\n" , value)

	passByPointer(&value)
	fmt.Printf("third: %d" , value)

}