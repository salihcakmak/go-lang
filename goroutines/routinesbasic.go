package goroutines

import (
	"fmt"
	"time"
)

func WriteOddNumbers() {
	for i := 1; i <= 9; i += 2 {
		fmt.Println("Odd :", i)
		time.Sleep(time.Second)
	}
}

func WriteEvenNumbers() {
	for i := 0; i <= 12; i += 2 {
		fmt.Println("Even :", i)
		time.Sleep(time.Second)
	}
}
