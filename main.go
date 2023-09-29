package main

import (
	"fmt"

	"github.com/chriscarsam/gofrom0/goroutines"
)

func main() {
	// fmt.Println(exercises.ConverValue("500"))
	//keyboard.EnterNumbers()
	//iterations.Iterate()
	//fmt.Println(exercises.MultiplicationTable())
	//files.SaveTable()
	//files.AddTable()
	//files.ReadFile()
	//functions.Calculations()
	//functions.CallClosure()
	//functions.Exponence(2)
	//arrangementsslices.ShowSlice()
	//arrangementsslices.Capacity()
	//maps.ShowMaps()
	//users.UserRegistration()
	/* Charly := new(models.Man)
	exer_interfaces.HumanBreathing(Charly)

	Isabella := new(models.Woman)
	exer_interfaces.HumanBreathing(Isabella) */
	//defer_panic.ExamplePanic()
	go goroutines.MySlowName("Charly Test")
	fmt.Println("I am here")
	var x string
	fmt.Scanln(&x)
}
