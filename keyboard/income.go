package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var legend string
var err error

func EnterNumbers() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter number 1 : ")
	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("The data entered is incorrect " + err.Error())
		}
	}

	fmt.Print("Enter number 2 : ")
	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("The data entered is incorrect " + err.Error())
		}
	}

	fmt.Print("Enter legend : ")
	if scanner.Scan() {
		legend = scanner.Text()
		if err != nil {
			panic("The data entered is incorrect " + err.Error())
		}
	}

	fmt.Println(legend, number1*number2)
}
