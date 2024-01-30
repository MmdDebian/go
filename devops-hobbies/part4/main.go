package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type element struct {
	data byte 
	prev *element
	next *element 
}

type command string 

const (
	INSERT 		command = "INSERT"
	LEFT 		command = "LEFT"
	RIGHT 		command = "RIGHT"
	BACKSPACE 	command = "BACKSPACE"
)

func main() {

	var count int 

	fmt.Scan(&count)

	reader := bufio.NewReader(os.Stdin) 


	curser := &element{data: '|'}

	for index := 0; index < count; index++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line , "\n")
		fields := strings.Fields(line)

		switch command(fields[0]){
			case INSERT :
				element := element{data: []byte(fields[1])[0]}

				curser.prev

			case LEFT :
				if curser.prev != nil {
					curser.prev.data,curser.data = curser.data , curser.prev.data
					curser = curser.prev
				}
			case RIGHT :
					if curser.next != nil {
					curser.next.data,curser.data = curser.data , curser.next.data
					curser = curser.next
				}

			case BACKSPACE :
				if curser.prev.prev != nil{
					curser.prev.prev.next = curser
					curser.prev = curser.prev.prev 
				}else {
					curser.prev = nil 
				}
			default :
				panic("invalid command given")
		}
	}
}