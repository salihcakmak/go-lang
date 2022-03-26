package deferstatement

import "fmt"

func CheckNumber(sayi int) string {
	defer DeferFunc()

	if sayi%2 == 0 {
		return "Çift sayıdır."
	} else {
		return "Tek sayıdır."
	}
}

func DeferFunc() {
	fmt.Println("Bitti.")
}

func DeferDemoTestFunc()  {
	result := CheckNumber(55)
	fmt.Println(result)
}