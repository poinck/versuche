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
        var stein steinmuster
        
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
	    } else if message == "ping" {
	    	stein = Meinstein(c)
	    	if stein.aussen != 0 {
	    		message = fmt.Sprintf("pong %d %d %d", stein.aussen, stein.mitte, stein.innen)
	    		h.broadcast <- message
	    	} else {
	    		err = websocket.Message.Send(c.ws, "nö!")
	    	}
	    	
	    	// debug
	    	// log.Println("anzahl der verbindungen =", len(h.connections))
	    } else if message == "gib" {
	    	stein = Getstein(c)
	    	if stein.aussen != 0 {
		    	err = websocket.Message.Send(c.ws, fmt.Sprintf("ich %d %d %d", stein.aussen, stein.mitte, stein.innen))
		    	h.broadcast <- "resetfeld"
		    	for aktuellerstein := range steine {
		    		if steine[aktuellerstein].verwendet == true {
		    			message = fmt.Sprintf("feld %d %d %d", steine[aktuellerstein].aussen, steine[aktuellerstein].mitte, steine[aktuellerstein].innen)
		    			h.broadcast <- message
		    		}
		    	}
		    	if err != nil {
		        	log.Println("Fehler", err)
		        }
			} else {
				err = websocket.Message.Send(c.ws, "alle Steine vergeben")
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
	if Freestein(c) {
		h.broadcast <- "resetfeld"
		for aktuellerstein := range steine {
			if steine[aktuellerstein].verwendet == true {
				message := fmt.Sprintf("feld %d %d %d", steine[aktuellerstein].aussen, steine[aktuellerstein].mitte, steine[aktuellerstein].innen)
				h.broadcast <- message
			}
		}
	} else {
		log.Println("eine Verbindung von der ich nichts wusste wurde geschlossen")
	}
	
	h.unregister <- c
	// h.broadcast <- "Jemand verließ das sinkende Schiff"
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