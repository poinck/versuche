package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int { 
    n := 0
	m := 1
    return func() int {
        fmt.Println("n, m =", n, m)
        
        // gleichzeitiges Berechnen zweier Variablen um zu verhindern,
        // dass die Berechnungen voneinander abhängen (alternativ 
        // könnte eine dritte Variable deklariert werden)
        n, m = m, n + m
        
        return n
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 400; i++ {
        fmt.Println(i, ":", f())
    }
}

