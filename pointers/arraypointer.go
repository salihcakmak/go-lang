package pointers

import "fmt"

func ArrayPointerFunc(numbers []int) {
	numbers[0] = 55
	numbers[3] = 999
	fmt.Println("Array in function :", numbers)
}
