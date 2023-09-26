package functions

import "fmt"

func Calculations() {
	var number3 int = 32
	var number4 int = 243

	calculation := func(number1, number2 int) int {
		return number1 + number2 + number3 + number4
	}
	fmt.Println(calculation(10, 25))

	calculation = func(number1, number2 int) int {
		return number1 * number2 * number3
	}

	fmt.Println(calculation(10, 25))
}