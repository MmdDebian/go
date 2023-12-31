package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	for i := 0 ; i < len(s) ; i++ {
		fmt.Printf("%x" , s[i])
	}

	fmt.Println()

	fmt.Println("Rune count : " , utf8.RuneCountInString(s))
}