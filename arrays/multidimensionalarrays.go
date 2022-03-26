package arrays

import "fmt"

func MultidimensionalArrayTest() {
	var numbers [2][3]int

	numbers[0][0] = 7
	numbers[0][1] = 44
	numbers[0][2] = 3
	numbers[1][0] = 56
	numbers[1][1] = 435
	numbers[1][2] = 234
	fmt.Println(numbers)
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			fmt.Println(numbers[i][j])
		}
	}

}
