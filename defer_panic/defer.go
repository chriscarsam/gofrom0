package defer_panic

import (
	"fmt"
	"log"
)

func SeeDefer() {
	fmt.Println("This is a first message")
	defer fmt.Println("This is the final message")
	fmt.Println("This is the second message")
}

func ExamplePanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("An error occurred that generated Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Value 1 was found")
	}
}
