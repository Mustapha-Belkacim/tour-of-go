package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}


var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.343, -74.34,
    }

    fmt.Println(m["Bell Labs"])
}
