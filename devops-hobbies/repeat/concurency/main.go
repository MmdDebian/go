package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	file := "users.txt"
	var wg sync.WaitGroup 
	jobs := make(chan string, 1)
	const workers = 4

	for worker := 0; worker < workers; worker++ {
		go findAdmin(jobs , &wg , worker)
	}

	reader , _ := os.Open(file)
	scanner := bufio.NewScanner(reader)
	
	for scanner.Scan() {
		jobs <- scanner.Text() 
		wg.Add(1)
	}

	wg.Wait()
}

func findAdmin(props chan string , wg *sync.WaitGroup , worker int){
	for prop := range props {
		if prop == "ADMIN" {
			fmt.Printf("this prop is admin processed by worker %d\n" , worker)
		}else {
			fmt.Printf("this user can be %s processed by %d \n" , prop , worker)
		}
		wg.Done()
	}
}