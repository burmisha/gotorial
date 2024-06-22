// go run tasks/task182/main.go
package main

import (
	"fmt"
)

// Produce возвращает срез из n значений val.
func ZipMap[K comparable, V any](keys []K, vals []V) map[K]V {
	keysLength := len(keys)
	valuesLength := len(vals)

	length := keysLength
	if valuesLength < keysLength {
		length = valuesLength
	}

	result := make(map[K]V, length)
	for i := 0; i < length; i++ {
		result[keys[i]] = vals[i]
	}
	return result
}

// конец решения

func main() {
	keys := []string{"one", "two", "thr"}
	vals := []int{11, 22, 33}

	m := ZipMap(keys, vals)
	fmt.Println(m)
}
