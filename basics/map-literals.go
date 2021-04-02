package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}


var m = map[string]Vertex{
    "Bell labs": Vertex{
        40.34, -71.34,
    },
    "Google": Vertex{
        34.34, 12.34,
    },
}

func main() {
    fmt.Println(m)
}
