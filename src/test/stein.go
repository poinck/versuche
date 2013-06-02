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

func Getstein(c *connection) steinmuster {
	var muster steinmuster
	
	fmt.Println("anzahl der steine = ", len(steine))
	dieses := rand.Intn(len(steine) - 1)
	fmt.Println("zufällige zahl =", dieses)
	
	muster = steine[dieses]
	steine[dieses].verwendet = true
	steine[dieses].verbindung = c
	
	fmt.Println(steine[dieses])

	return muster
}