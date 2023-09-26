package arrangementsslices

import "fmt"

var table [10]int = [10]int{10, 0, 59}
var matrix [20][30]int

func ShowArrangements() {
	table[7] = 33
	table[2] = 54
	fmt.Println(table)

	tableTwo := [10]string{"Pablo", "Juan"}
	fmt.Println(tableTwo)

	for i := 0; i < len(table); i++ {
		fmt.Printf("table[%d] = %d \n", i, table[i])
	}

	matrix[7][24] = 15
	fmt.Println(matrix)
}
