package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var State bool
var Salary float32
var Date time.Time

func RestVariables() {
	Name = "Charly"
	State = true
	Salary = 1577.66
	Date = time.Now()

	fmt.Println("Name = ", Name)
	fmt.Println("State = ", State)
	fmt.Println("Salary = ", Salary)
	fmt.Println("Date = ", Date)
}

func ConvertText(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text
}
