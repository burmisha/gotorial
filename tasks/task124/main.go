package main

import (
	"fmt"
)

// echo en | go run tasks/task124/main.go
func main() {
	var code string
	fmt.Scan(&code)

	// определите полное название языка по его коду
	// и запишите его в переменную `lang`
	// ...
	var lang string
	switch code {
	case "ru", "rus":
		lang = "Russian"
	case "en":
		lang = "English"
	case "fr":
		lang = "French"
	default:
		lang = "Unknown"
	}

	fmt.Println(lang)
}
