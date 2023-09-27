package exer_interfaces

import (
	"fmt"

	"github.com/chriscarsam/gofrom0/interfaces"
)

func HumanBreathing(hu interfaces.Human) {
	hu.Breathe()
	fmt.Printf("I am a %s, and I am breathing \n", hu.Sex())
}
