package arrangementsslices

import "fmt"

var tableSlice []int = []int{2, 5, 4}
var arrangement [10]int = [10]int{6, 98, 21, 674, 18, 35, 78, 9, 5, 7}

func ShowSlice() {
	fmt.Println(tableSlice)

	portion := arrangement[3:]       // Slice created from a vector, from position 3
	portionTwo := arrangement[:5]    // Slice created from a vector, from position 0 to position 5
	portionThree := arrangement[6:7] // Slice created from a vector, from position 6 to 7

	fmt.Println(portion)
	fmt.Println(portionTwo)
	fmt.Println(portionThree)
}

func Capacity() {
	elements := make([]int, 5, 20)
	fmt.Printf("Lengh %d, Capacity %d \n", len(elements), cap(elements))

	numbers := make([]int, 0)
	for i := 0; i < 100; i++ {
		numbers = append(numbers, i)
	}
	fmt.Printf("Lengh %d, Capacity %d \n", len(numbers), cap(numbers))
}
