package surl

import (
	"math"
	"strings"
)

// hexConvert
func hexConvert(num string, currentRadix, targetRadix int) (result string) {
	ten := enten(num, currentRadix)
	result = deten(ten, targetRadix)
	return
}

// N to 10
func enten(num string, r int) (result int) {
	for k, v := range strings.Split(reverse(num), "") {
		result += strings.Index(chars, v) * int(math.Pow(float64(r), float64(k)))
	}
	return
}

// 10 to N
func deten(num int, r int) (result string) {
	for i := num; i > 0; i /= r {
		result = string(chars[i%r]) + result
	}
	return
}

// reverse
func reverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
