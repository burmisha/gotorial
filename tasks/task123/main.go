package main

import (
	"fmt"
)

// echo a 5 | go run tasks/task123/main.go
func main() {
	var source string
	var times int
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	result := ""
	for i := 0; i < times; i++ {
		result += source
	}
	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
	// ...

	fmt.Println(result)
}
