package main

import (
    "net/http"
    "fmt"
)

type String string

type Struct struct {
    Greeting string
    Punct string
    Who string
}

func (h Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.Greeting)
	// fmt.Println(h.Greeting)
}

func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(h))
}

func main() {
    var h Struct
    var h2 String
    
    h.Greeting = "Jemand zuhause?"
    h2 = "Nein!"
    http.Handle("/hallo", h)
    http.Handle("/nein", h2)
    http.ListenAndServe("localhost:4000", nil)
}
