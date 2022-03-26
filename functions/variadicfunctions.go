package functions

// Variadic fonksiyonlar sınırsız sayıda parametre alabilir. Bu değerlere dizi gibi erişilebilir
func GetSum(numbers ...int) int {
	var sum int
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}
	return sum
}
