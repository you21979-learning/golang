package main

import (
    "fmt"
)
func improve (guess, x float64) float64 {
    return average(guess, x / guess)
}
func average (x, y float64) float64 {
    return (x + y) / 2
}
func good_enough(guess, x float64) bool {
    if (abs((square(guess)) - x) < 0.001) {
        return true
    } else {
        return false
    }
}
func abs (x float64) float64 {
    if x < 0 {
        return x * -1
    } else {
        return x
    }
}
func square (x float64) float64 {
    return x * x;
}
func sqrt_iter (guess, x float64) float64 {
    if good_enough(guess, x) {
        return guess
    } else {
        return sqrt_iter(improve(guess, x), x)
    }
}
func Sqrt(x float64) float64 {
    return sqrt_iter(1.0, x)
}
func main() {
    fmt.Println(Sqrt(2))
}
