package main

import (
	"context"
	"fmt"
	"time"
)

func simulateLongOperation(ctx context.Context) {
	for i := 0; i < 5; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Doing some work...")
		case <-ctx.Done():
			fmt.Println("Operation canceled")
			break
		}
	}
	fmt.Println("Operation completed successfully!")
}

func main() {
	contextChild, _ := context.WithTimeout(context.Background(), 5*time.Second)

	go simulateLongOperation(contextChild)

	time.Sleep(6 * time.Second) // Give the operation time to complete or be canceled
}
