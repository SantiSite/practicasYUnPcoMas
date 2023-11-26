package retos

/**
Objetivo: crear una función que reciba una cade de texto y retorne la cade de texto al revés.

Ejemplo:
	- input: abcd
	- output: dcba
*/

func ReverseString(s string) string {
	l := len(s)
	r := []rune(s)
	rReverse := make([]rune, l)
	for i, code := range r {
		rReverse[(l-i)-1] = code
	}
	return string(rReverse)
}

func ReverseStringPro(s string) string {
	var result string
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func ReverseStringUltra(s string) string {
	r := []rune(s)

	for i, lastPosition := 0, len(r)-1; i < lastPosition; i, lastPosition = i+1, lastPosition-1 {
		r[i], r[lastPosition] = r[lastPosition], r[i]
	}

	return string(r)
}
