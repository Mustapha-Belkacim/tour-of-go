package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }
    z := 1.0
    s := 0.0
    for i := 0; i < 10; i++ {
        z -= (z*z - x) / (2*z)
        if math.Abs(z-s) < 1e-10 {
            break
        }
        s = z
    }
    return z, nil
}

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}


