// go run tasks/task184/main.go
package main

import (
	"fmt"
)

type Map[K comparable, V any] struct {
	impl map[K]V
}

func (m *Map[K, V]) Set(k K, v V) {
	if m.impl == nil {
		m.impl = make(map[K]V, 1)
	}
	m.impl[k] = v
}

func (m *Map[K, V]) Get(k K) V {

	return m.impl[k]
}

func (m *Map[K, V]) Keys() []K {
	result := make([]K, 0, len(m.impl))
	for k, _ := range m.impl {
		result = append(result, k)
	}
	return result
}

func (m *Map[K, V]) Values() []V {
	result := make([]V, 0, len(m.impl))
	for _, v := range m.impl {
		result = append(result, v)
	}
	return result
}

func main() {
	m := Map[string, int]{}
	m.Set("one", 1)
	m.Set("two", 2)

	fmt.Println(m.Get("one")) // 1
	fmt.Println(m.Get("two")) // 2

	fmt.Println(m.Keys())   // [one two]
	fmt.Println(m.Values()) // [1 2]
}
