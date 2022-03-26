package pointers

import "fmt"

func PointerBasicTestFunc(val *int) {
	fmt.Println("Value in parameter :", *val)
	*val = 150
	fmt.Println("Value in function :", *val)
}

func PointerAddressFunc(val *int)  {
	fmt.Println("Variable address in function :", val)
}
