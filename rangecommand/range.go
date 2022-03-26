package rangecommand

import "fmt"

func RangeTestFunc() {
	cities := []string{"İstanbul", "İzmir", "Antalya"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for i, city := range cities {
		fmt.Println(i)
		fmt.Println(city)
	}
}
