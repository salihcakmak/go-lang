package interfaces

import (
	"fmt"
)

// Bu parametre her türden değer alabilir demek. Go da herşey interface üzerinedir.
func isTrue(i interface{}) {
	number, ok := i.(int)
	fmt.Println(number, ok)
}

func TypeAssertionTestFunc() {

	var val interface{} = "Salih"
	isTrue(val)

	var val2 interface{} = 5
	isTrue(val2)
}
