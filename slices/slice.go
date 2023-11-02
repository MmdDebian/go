package main

import "fmt"

func main() {
	var s []string

	fmt.Println("uninit:" , s , s == nil , len(s) == 0)

	s = make([]string , 5)

	fmt.Println("emt:" , s , "len:" , len(s) , "cap:" , cap(s))


	s[0] = "A"
	s[1] = "B"
	s[2] = "C"


	fmt.Println("set:" ,s)
	fmt.Println("get:", s[2])

	fmt.Println("len" , len(s))


	s = append(s, "D")
	s = append(s, "e" , "f")

	fmt.Println("apd :" , s)


	c := make([]string , len(s))

	copy(c , s)

	fmt.Println(c)

	l := s[2:5]

	fmt.Println("sl1:" , l)
}