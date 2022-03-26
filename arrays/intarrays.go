package arrays

import "fmt"

func IntArrayTestFunc() {
	var numbers [5]int

	fmt.Println(numbers)
	numbers[2] = 50
	fmt.Println(numbers)
	fmt.Println(numbers[2])
}