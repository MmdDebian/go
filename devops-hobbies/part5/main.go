package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	jobs := make(chan string , 1)

	const workers = 3 

	for worker := 0; worker < workers; worker++ {
		go getWebsite(jobs , &wg ,worker)
	}

	file := "websites.txt"
	reader , _ := os.Open(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		jobs <- scanner.Text()
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("total time : ", time.Since(startTime))
}


func getWebsite(websites chan string , wg *sync.WaitGroup , worker int){
	for website := range websites {
		if res , err := http.Get(website); err != nil {
			fmt.Printf("%s is down processed by worker %d \n" , website , worker)
		}else{
			fmt.Printf("[%d] %s is up processed by worker %d \n" , res.StatusCode , website , worker)
		}
		wg.Done()
	}
}