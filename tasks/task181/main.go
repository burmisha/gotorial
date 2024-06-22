// go run tasks/task181/main.go
package main

import (
	"fmt"
)

// Produce возвращает срез из n значений val.
func Produce[T any](val T, n int) []T {
	vals := make([]T, n)
	for i := range n {
		vals[i] = val
	}
	return vals
}

// конец решения

func main() {
	fmt.Printf("%v %v", Produce("abc", 5), Produce(21, 5))
}
