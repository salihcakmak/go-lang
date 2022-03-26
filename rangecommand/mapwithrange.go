package rangecommand

import "fmt"

func MapWithRangeTestFunc() {
	dictionary := map[string]string{"table": "masa", "book": "kitap", "phone": "telefon"}

	for key, value := range dictionary {
		fmt.Print("Anahtar :", key)
		fmt.Print(" ")
		fmt.Println("Değer :", value)
	}
}
