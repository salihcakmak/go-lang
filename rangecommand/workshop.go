package rangecommand

import "fmt"

// Görev çift sayıların toplamını bulmak ve yazmak
func TotalEvenNumbers() {
	var total int
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			total += number
		}
	}
	fmt.Println("Çift sayıların toplamı :", total)
}
