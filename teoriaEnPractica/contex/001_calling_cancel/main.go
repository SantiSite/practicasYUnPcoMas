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
			return
		}
	}
	fmt.Println("Operation completed successfully")
}

func main() {
	contextChild, cancelFunc := context.WithCancel(context.Background())

	go simulateLongOperation(contextChild)

	time.Sleep(3 * time.Second)

	cancelFunc()

	// Ejemplo haciendo conexiones a base de datos:
	//($mongo.Client{}).Database("ba").Collection("c").FindOne(contextChild, nil)
}
