package text

import (
	"strconv"
	"strings"
)

// AsDigitString возвращает число в виде строки, состоящей из его цифр,
// разделенных дефисами. Например, 42513 → 4-2-5-1-3
func AsDigitString(n int) string {
	digits := AsDigits(n)
	parts := make([]string, len(digits))
	for idx, d := range digits {
		parts[idx] = strconv.Itoa(d)
	}
	return strings.Join(parts, "-")
}

// AsRunes возвращает возвращает срез с цифрами числа,
// представленными в виде рун.
func AsRunes(n int) []rune {
	return []rune(strconv.Itoa(n))
}

// AsDigits возвращает срез с цифрами числа.
func AsDigits(n int) []int {
	runes := AsRunes(n)
	count := len(runes)
	zero := int('0')
	digits := make([]int, count)
	for idx, char := range runes {
		digits[idx] = int(char) - zero
	}
	return digits
}
