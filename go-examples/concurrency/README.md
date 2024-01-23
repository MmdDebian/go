# Concurrency


Large programs are often made up of many smaller sub-programs .

For example a web server handles requests made from web browsers up HTML web pages in response , Each request is handles like a small program . 


It world be idea for programs like these to be able to run their smaller components at the same time (in the case of the web server to handle multiple requests) . Making progress on more than one task simultaneously is know as concurrency . Go has rich support of concurrency using gotoutines and channels . 


## Goroutines 

A goroutines is a function that is capable of running concurency with other function . to create a goroutine we use the kewyword **go** followed by a function invocation :


```
pcakage main 


import "fmt" 

func f (n int) {
    for i := 0 ; i < 10 ; i++ {
        fmt.Println(i , ":" , i)
    }
}


func main(){
    go f(0)
    var input string 
    fmt.Scanln(&input)
}
```

This program consists of two goroutines , The first goroutine is implicit is the main function itself. The second goroutine is created when we call **go f(0)** . 

Normally when we invoke a function our program will execute all the statements in a funcation and then return to the next line following the invocation . 


Goroutines are lightweight and we can easily create thousands of them. We can modify our program to run 10 goroutines by doing this:


```
func main() {
  for i := 0; i < 10; i++ {
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
}
```


You may have noticed that when you run this program it seems to run the goroutines in order rather than simultaneously. Let's add some delay to the function using time.Sleep and rand.Intn:


```
package main

import (
  "fmt"
  "time"
  "math/rand"
)

func f(n int) {
  for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
  }
}

func main() {
  for i := 0; i < 10; i++ {
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
}
```
**f** prints out the numbers from 0 to 10, waiting between 0 and 250 ms after each one. The goroutines should now run simultaneously .
