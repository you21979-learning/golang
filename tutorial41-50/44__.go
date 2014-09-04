package main

import "fmt"
import "math"


var SQRT_5 = math.Sqrt(5);
var PHY = (1 + SQRT_5) / 2;

func round(f float64) int {
    return  int(math.Floor(f+0.5))
}
func fib(n float64) int {
  return round(math.Pow(PHY, n) / SQRT_5);
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    n := 0.0
    return func() int {
        n = n + 1
        return fib(n)
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
