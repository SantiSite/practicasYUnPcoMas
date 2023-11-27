package main

import (
	"context"
	"fmt"
)

func main() {
	ctxParent := context.Background()
	auth(ctxParent)
}

func auth(ctxParent context.Context) {
	ctxChild := context.WithValue(ctxParent, "env", "sandbox")
	func1(ctxChild)
}

func func1(ctxChild context.Context) {
	fmt.Println("Hola desde func1")
	func2(ctxChild)
}

func func2(ctxChild context.Context) {
	fmt.Println("Hola desde func2")

	env := "live"
	if v := ctxChild.Value("env"); v != nil {
		envToString, ok := v.(string)
		if ok {
			env = envToString
		}
	}
	fmt.Println("env:", env)
}
