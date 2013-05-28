package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := float64(x)
    a := float64(x * 3 - 1)
    for i := float64(0); i < a; i++ {
        if z != z - (((z*z) - x) / (2*z)) {
            z = z - (((z*z) - x) / (2*z))
        } else {
            break
        }
        fmt.Println(i, ": ", z)
    }
    
    return z
}

func main() {
    n := float64(7)
    fmt.Println("Sqrt: ", Sqrt(n))
    fmt.Println("math.Sqrt: ", math.Sqrt(n))
}

