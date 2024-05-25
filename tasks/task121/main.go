package main

import (
	"fmt"
	"time"
)

// echo 1h30m | go run tasks/task121/main.go

func main() {
	// считываем временной отрезок из os.Stdin
	// гарантируется, что значение корректное
	// не меняйте этот блок
	var s string
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)
	minutes := int64(d.Minutes())
	// выведите исходное значение
	// и количество минут в нем
	// в формате "исходное = X min"
	// используйте метод .Minutes() объекта d
	result := fmt.Sprintf("%s = %d min", s, minutes)
	fmt.Println(result)
}
