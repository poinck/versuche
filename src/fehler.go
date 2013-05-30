package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    // e muss vor Ausgabe nach float64 gecastet werden, weil *print*()
    // den Typ "ErrNegativeSqrt" nicht kennt
    return fmt.Sprint(float64(e), " ist eine netagive Zahl")
}

func Sqrt(f float64) (float64, error, int) {
    if f < 0 {
  		return 0, ErrNegativeSqrt(f), 1
    }
    
    return math.Sqrt(f), nil, 0
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
    
}

