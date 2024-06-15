package main

import (
	"fmt"

	"github.com/gothanks/more/text"
)

func main() {
	digits := text.AsDigits(42513)
	fmt.Println("42513 →", digits)
}
