package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
    "fmt"
)

func WordCount(s string) map[string]int {
    wc := make(map[string]int)
    
    wörter := strings.Fields(s)
    fmt.Printf("%q\n", wörter)
    for n := range wörter {
        fmt.Println(n, ":", wörter[n])
        wc[wörter[n]] = wc[wörter[n]] + 1
    }
    
    return wc
}

func main() {
    wc.Test(WordCount)
}

