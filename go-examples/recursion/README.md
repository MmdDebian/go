# Go : Recursion


This fact function calls itself until it reaches the base case of fact(0).




Closures can also be recursive, but this requires the closure to be declared with a typed var explicitly before it’s defined.





Since fib was previously declared in main, Go knows which function to call with fib here.




```

package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))

    var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n
        }

        return fib(n-1) + fib(n-2)
    }

    fmt.Println(fib(7))
}
```


output : 


```

$ go run recursion.go 
5040
13

```