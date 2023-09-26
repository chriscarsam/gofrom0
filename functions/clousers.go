package functions

import "fmt"

func table(value int) func() int {
	number := value
	sequence := 0
	return func() int {
		sequence++
		return number * sequence
	}
}

func CallClosure() {
	tableOf := 2
	MyTable := table(tableOf)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", i, tableOf, MyTable())
	}
}
