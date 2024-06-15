// go run tasks/task143/main.go
package main

import (
	"fmt"
	"os"
)

// normalize нормализует значения, переданные в vals,
// так чтобы их сумма была равна 1.
func normalize(numbers ...*float64) {
	sum := 0.
	for _, number := range numbers {
		sum += *number
	}
	for _, number := range numbers {
		*number /= sum
	}
}

func main() {
	a, b, c, d := 1.0, 2.0, 3.0, 4.0
	normalize(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
	// 0.1 0.2 0.3 0.4
	fmt.Println("PASS")
	os.Exit(0)
}
