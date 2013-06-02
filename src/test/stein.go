package main

import (
	"math/rand"
	"fmt"
	"time"
)

type steinmuster struct {
	aussen int
	mitte int
	innen int
	verwendet bool
	verbindung *connection
}

var steine []steinmuster

func InitSteine() {
	rand.Seed(time.Now().UTC().UnixNano())

	// steine = make([]steinmuster, 0)
	anzahl := 0
	for i := 1; i <= 3; i++ {
		for m := 1; m <= 3; m++ {
			for a := 1; a <= 2; a++ {
				var muster steinmuster
				muster.aussen = a
				muster.innen = i
				muster.mitte = m
				steine = append(steine, muster)
				anzahl++
			}
		}
	}
	
}

func Meinstein(c *connection) steinmuster {
	stein := steinmuster{0, 0, 0, false, c}

	for aktuellerstein := range steine {
		if steine[aktuellerstein].verbindung == c {
			stein = steine[aktuellerstein]
		}
	}
	
	return stein
}

// Stein wieder freigeben
func Freestein(c *connection) bool {
	erfolg := false
	
	for aktuellerstein := range steine {
		if steine[aktuellerstein].verbindung == c {
			steine[aktuellerstein].verwendet = false
		}
	}
	
	return erfolg
}

// Stein vergeben und merken
func Getstein(c *connection) steinmuster {
	var muster steinmuster
	alleverwendet := true
	
	fmt.Println("anzahl der steine = ", len(steine))
	dieses := rand.Intn(len(steine) - 1)
	fmt.Println("zufällige zahl =", dieses)
	if steine[dieses].verwendet == true {
		for dieser := range steine {
			if steine[dieser].verwendet != true {
				fmt.Println(dieser, "war der nächste noch nicht verwendete Stein")
				dieses = dieser
				alleverwendet = false
				break
			} 
			
			// debug
			fmt.Println(dieser, "verwendet ?", steine[dieser].verwendet)
		}
	} else {
		alleverwendet = false
	}
	
	if alleverwendet == false {
		muster = steine[dieses]
		steine[dieses].verwendet = true
		steine[dieses].verbindung = c
	} else {
		// keine Steine mehr zum Vergeben; deswegen Dummy-Stein
		muster = steinmuster{0, 0, 0, false, c}
	}
	
	// debug
	// fmt.Println(steine[dieses])

	return muster
}