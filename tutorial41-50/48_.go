package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) complex128 {
    return cmplx.Pow(x, 1.0/3)
}

func main() {
    fmt.Println(Cbrt(2))
}
