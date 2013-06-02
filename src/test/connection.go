package main

import (
    "code.google.com/p/go.net/websocket"
    "log"
    "fmt"
    // "encoding/json"
)

type connection struct {
    // Zeiger zur Websocket connection
    ws *websocket.Conn

    // Buffered channel of outbound messages.
    send chan string
}

func (c *connection) reader() {
    for {
        var message string
        err := websocket.Message.Receive(c.ws, &message)
        if err != nil {
        	log.Println("Fehler", err)
        	break
        }
        log.Println("message =", message)
        if message == "Hey Server" {
        	err := websocket.Message.Send(c.ws, "Ja, was gibts?")
        	if err != nil {
	        	log.Println("Fehler", err)
	        }
	    } else if message == "Ping" {
	    	log.Println("anzahl der verbindungen =", len(h.connections))
	    	message = "Pong"
	    	h.broadcast <- message
	    } else if message == "gib" {
	    	stein := Getstein(c)
	    	/*
	    	b, err := json.Marshal(stein)
	    	fmt.Printf("jsonzeugs = %s\n", b)
	    	*/
	    	err = websocket.Message.Send(c.ws, fmt.Sprintf("ich %d %d %d", stein.aussen, stein.mitte, stein.innen))
	    	message = fmt.Sprintf("feld %d %d %d", stein.aussen, stein.mitte, stein.innen)
	    	h.broadcast <- message
	    	if err != nil {
	        	log.Println("Fehler", err)
	        }
        } else {
        	h.broadcast <- message
        }
    }
    c.ws.Close()
}

func (c *connection) writer() {
    for message := range c.send {
        err := websocket.Message.Send(c.ws, message)
        if err != nil {
        	log.Println("message =", message)
        
            break
        }
        
    }
    c.ws.Close()
}

func disconnect(c *connection) {
	h.unregister <- c
	h.broadcast <- "Jemand verließ das sinkende Schiff"
}

func wsHandler(ws *websocket.Conn) {
    err := websocket.Message.Send(ws, "Hier Server an Client, du kannst tatsächlich <b>Websockets</b>")
    if err != nil {
        log.Println("Versenden lief schief")
    }
    
    c := &connection{send: make(chan string, 256), ws: ws}
    h.register <- c
    defer disconnect(c)
    go c.writer()
    c.reader()
}