package main

import "fmt"

func fibonacci() func() int {
    i_0, i_1 := 0, 1
    return func() int {
        ret := i_0
        i_0, i_1 = i_1, i_0 + i_1
        return ret
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
