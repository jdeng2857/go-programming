package main

import (
	"fmt"
	"geometry/coordinate"

	"github.com/hackebrot/turtle"
)

func main() {
	pt1 := coordinate.Point{X: 2, Y: 3}
	fmt.Println(pt1)
	fmt.Println(pt1.Length())

	emoji, ok := turtle.Emojis["smiley"]
	if !ok {
		fmt.Println("No emoji found.")
	} else {
		fmt.Println(emoji.Char)
	}
}
