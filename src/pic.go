package main

import (
	"code.google.com/p/go-tour/pic"
	// "fmt"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
		for ii := range pic[i] {
			pic[i][ii] = uint8((dy ^ i) - (dx ^ ii))
		}
	}

	// fmt.Println(pic)

	return pic
}

func main() {
	pic.Show(Pic)
}

