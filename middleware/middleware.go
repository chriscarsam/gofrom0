package middleware

import "fmt"

func sum(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func MyMiddleware() {
	fmt.Println("Home")

	result := operationsMidd(sum)(2, 3)
	fmt.Println(result)
	result = operationsMidd(subtract)(10, 6)
	fmt.Println(result)
	result = operationsMidd(multiply)(2, 4)
	fmt.Println(result)
}

func operationsMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Start of operation")
		return f(a, b)
	}
}
