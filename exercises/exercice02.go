package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// exercice02 - Multiplication table

func MultiplicationTable() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a number to multiply: ")

	if scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			MultiplicationTable()
		} else {
			fmt.Printf("*** Table of %d ***\n", value)
			for i := 1; i <= 10; i++ {
				fmt.Printf("%d x %d = %d \n", value, i, i*value)
			}
		}
	}
}
