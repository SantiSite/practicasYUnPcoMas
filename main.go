package main

import (
	"fmt"

	"github.com/SantiSite/practicasYUnPocoMas/retos"
)

func main() {
	fmt.Println(retos.ReverseString("abcd"))
	fmt.Println(retos.ReverseString("santi"))
	fmt.Println(retos.ReverseString("santi res"))
	fmt.Println("----------------------------------")
	fmt.Println(retos.ReverseStringPro("abcd"))
	fmt.Println(retos.ReverseStringPro("santi"))
	fmt.Println(retos.ReverseStringPro("santi res"))
	fmt.Println("----------------------------------")
	fmt.Println(retos.ReverseStringUltra("abcd"))
	fmt.Println(retos.ReverseStringUltra("santi"))
	fmt.Println(retos.ReverseStringUltra("santi res"))
}
