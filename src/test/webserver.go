package main

import (
    "net/http"
    "fmt"
    "flag"
    "os"
    "io/ioutil"
    "log"
    "code.google.com/p/go.net/websocket"   
)

type String string

type Struct struct {
    begrüßung string
}

type Datei struct {
	index string
	pfad string
}

func (h Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.begrüßung)
}

func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(h))
}

func Meldung(text string) string {
	return "<html><p style='color: purple;'>" + text + "</p></html>"
}

func getdatei(pfad string) string {
	var inhalt string

	inhalt = ""
	
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Fehler beim Ermitteln des aktuellen Pfads")
	} else {
		log.Println("Zugriff:", pfad)
		data, err := ioutil.ReadFile(pwd + pfad)
		if err != nil {
			log.Println("Datei nicht gefunden:", pfad, err.Error())
			inhalt += Meldung("<b>404</b>, Datei nicht gefunden. <i>*hmpf*</i>")
		}
		inhalt += fmt.Sprintf("%s", string(data))
	}
	
	return inhalt
}

func (datei Datei) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pfad := r.RequestURI
	inhalt := getdatei(pfad)
	datei.index = inhalt

	fmt.Fprint(w, datei.index)
}

func webserver(adresse string) {
	var h Struct
    var h2 String
    var datei Datei
    
    h.begrüßung = "<html>Jemand zuhause? - <a href='/nein'>Antwort</a></html>"
    h2 = "<html>Nein! - <a href='/hallo'>Was war die Frage?</a></html>"
    datei.index = "?"
    http.Handle("/hallo", h)
    http.Handle("/nein", h2)
    http.Handle("/", datei)
    
    http.Handle("/ws", websocket.Handler(wsHandler))

	http.ListenAndServe(adresse, nil)
}

var portflag int
const DEFAULTPORTFLAG int = 8080

func getflags() {
	flag.IntVar(&portflag, "port", DEFAULTPORTFLAG, "Port")
	flag.Parse()
	
	// check wether Port is 80 or > 1024
	if portflag != 80 {
		if portflag < 1024 {
			portflag = DEFAULTPORTFLAG
		}
	}
}

func main() {
    getflags()
    
    InitSteine()
    
    go h.run()
    
    adresse := "localhost:" + fmt.Sprint(portflag)
    log.Println("Webserver ist online auf:", adresse)
    
    webserver(adresse)
}
