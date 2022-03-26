package functions

func BasicOperations(numberOne int, numberTwo int) (int, int, int, int) {
	addition := numberOne + numberTwo
	subtraction := numberOne - numberTwo
	multipliaction := numberOne * numberTwo
	division := numberOne / numberTwo
	return addition, subtraction, multipliaction, division
}
