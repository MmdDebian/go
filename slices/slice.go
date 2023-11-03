package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

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

	l = s[:5]
	fmt.Println("sl2:" , l)

	l = s[2:]
    fmt.Println("sl3:", l)

	t := []string{"g","h","i"}
    fmt.Println("dcl:" , t)

	t2 := []string{"g","h","i"}

	equal := slices.Equal(t , t2) 

	fmt.Println(equal)
	
	if equal{
		fmt.Println("t == t2")
	}

	twoD := make([][]int , 3)

	for i := 0; i < 3 ; i++ {
		innerLen := i + 1 

		twoD[i] = make([]int , innerLen)

		for j := 0 ; j < innerLen ; j++ {
			twoD[i][j] = i + j 
		}
	}

	fmt.Println("2d :" , twoD)
}