package functions

import (
	"fmt"
)

/* Temel Fonksiyon tanımlama prototipi zorunlu : name | opsiyonel : params, returnType
func name(params) returnType {
	..
	..
}
*/
// Bu fonksiyon geri dönüş değeri olmayan iş yapan bir fonksiyondur.
func SayHello(name string) {
	fmt.Println("Hello dear", name)
}

// Bu fonksiyon geriye integer değer döndüren bir fonksiyondur.
func AdditionOperation(numberOne int, numberTwo int) int {
	sum := numberOne + numberTwo
	return sum
}
