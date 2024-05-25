// echo Eyjafjallajokull 6 | go run tasks/task131/main.go
package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	var res string
	if len(text) <= width {
		res = text
	} else {
		res = text[:width] + "..."
	}

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...

	fmt.Println(res)
}
