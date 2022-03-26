package arrays

import "fmt"

func BiggestArrayNumberFunc() {
	numbers := [5]int{10, 20, 70, 40, 50}
	fmt.Println(numbers)

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	biggestNumber := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > biggestNumber {
			biggestNumber = numbers[i]
		}
	}
	fmt.Printf("En büyük sayı : %v\n", biggestNumber)
}
