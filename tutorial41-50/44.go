package main

import "fmt"

func f(n int) int {
    if n < 2 {
        return n
    } else {
        return f(n - 2) + f(n - 1)
    }
}
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    n := 0
    return func() int {
        n = n + 1
        return f(n)
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
