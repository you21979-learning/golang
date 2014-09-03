package main

import "code.google.com/p/go-tour/pic"

func f(x,y int) uint8 {
    return uint8(x^y)
    //return uint8(x*y)
    //return uint8((x+y)/2)
}

func Pic(dx, dy int) [][]uint8 {
    screen := make([][]uint8, dx)
    for x := 0; x < dx; x++ {
    	screen[x] = make([]uint8, dy)
        for y := 0; y < dy; y++ {
            screen[x][y] = f(x, y)
        }
    }
    return screen
}

func main() {
    pic.Show(Pic)
}
