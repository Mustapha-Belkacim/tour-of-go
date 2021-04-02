package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    a := make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        pixels := make([]uint8, dx)
        for j := 0; j < dx; j++ {
            pixels[j] = uint8(i*j)
        }
        a[i] = pixels
    }

    return a
}

func main() {
    pic.Show(Pic)
}
