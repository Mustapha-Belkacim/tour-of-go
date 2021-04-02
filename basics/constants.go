package main

import "fmt"


const Pi = 3.14

func main() {
    const World = "hello"
    fmt.Println("hello", World)
    fmt.Println("happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)
}
