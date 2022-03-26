package stringfunctions

import (
	"fmt"
	s "strings" // alias tanımladım
)

// case sensitive
// ascii
func StringFunctionTest() {
	var exampleValue string = "Hi developer! This function is an example."

	fmt.Println("Bulunan Karakter sayısı :", s.Count(exampleValue, "e"))
	fmt.Println("İçerik içerisinde bulunuyor mu ? :", s.Contains(exampleValue, "k"))
	indexValue := s.Index(exampleValue, "d") // Harfin string içerisindeki index numarasını dönderir
	if indexValue == -1 {
		fmt.Println("Verilen değer string içerisinde yok. İndex değeri :", indexValue)
	} else {
		fmt.Println("Verilen değer string içerisinde var. İndex değeri :", indexValue)
	}

	fmt.Println("Tüm harfleri büyük yap :", s.ToUpper(exampleValue))
	fmt.Println("Tüm harfleri küçük yap :", s.ToLower(exampleValue))

	fmt.Println("Ön ek değerinde istenilen var mı ? :", s.HasPrefix(exampleValue, "Hi"))
	fmt.Println("Son ek değerinde istenilen var mı ? :", s.HasSuffix(exampleValue, "god"))

	letters := []string{"b", "l", "o", "c", "k", "c", "h", "a", "i", "n"}
	fmt.Println("Verilen stringlerin birleşmiş hali :", s.Join(letters, "-"))
	fmt.Println("Verilen stringlerin birleşmiş hali :", s.Join(letters, ""))

	joinValue := s.Join(letters, "*")
	fmt.Println("Birleşmiş hali :", joinValue)

	fmt.Println("'*' Karakterini '+' ile değiştir :", s.Replace(joinValue, "*", "+", 3))
	fmt.Println("Verilen stringi parçala :", s.Split(joinValue, "*"))
	fmt.Println("Tekrar et :", s.Repeat(joinValue, 5))
}
