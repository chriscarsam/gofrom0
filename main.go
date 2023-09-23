package main

import (
	"fmt"
	"runtime"
)

func main() {
	// fmt.Println(variables.ConvertText(1544))

	if os := runtime.GOOS; os == "linux" || os == "OS X" {
		fmt.Println("This is", os)
	} else {
		fmt.Println("This is", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Printf("This is %s \n", os)
	case "darwin":
		fmt.Printf("This is %s \n", os)
	default:
		fmt.Printf("This is %s \n", os)
	}
}
