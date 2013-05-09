package main

import (
    "fmt"
    "net/http"
    "net/http/cgi"
    "html"
)

func versuch(w http.ResponseWriter, r *http.Request) {
    if (r.URL.Path == "/cgi-bin/versuch/anders") {
    	anders(w, r)
    } else {
    	fmt.Fprintf(w, "Hallo Du, willkommen auf: %s %q", r.URL.Path, html.EscapeString(r.URL.Path))
    }
}

func anders(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Das andere Hallo!")
}

func main() {
	// http.HandleFunc("/anders", anders)
    http.HandleFunc("/", versuch)
    cgi.Serve(nil)
}

