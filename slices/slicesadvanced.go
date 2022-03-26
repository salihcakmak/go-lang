package slices

import "fmt"

func SlicesAdvancedTestFunc() {
	cities := []string{"Ankara", "İstanbul", "İzmir"}
	fmt.Println(cities)

	copyObj := make([]string, len(cities))
	fmt.Println(copyObj)

	copy(copyObj, cities)
	fmt.Println(copyObj)

	cities = append(cities, "Adana") // Burada yeni bir şehir dizisi oluşturu ve bu 2 parametreyi birleştirir.
	fmt.Println(cities)
	fmt.Println(copyObj)

	fmt.Println(cities[1:3])
	fmt.Println(cities[1:])
	fmt.Println(cities[:2])

}
