package main

import "fmt"

type Vertex struct{
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": {34.34, 32.12},
    "Google": {65.12, -78.32},
}

func main() {
    fmt.Println(m)
}
