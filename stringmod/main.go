package main

import (
	"fmt"

	"github.com/jdeng2857/stringmod/quotes"
	"github.com/jdeng2857/stringmod/strings"
)

func main() {
	o, e := strings.CountOddEven("12345")
	fmt.Println(o, e)
	fmt.Println(quotes.GetEmoji("turtle"))
}
