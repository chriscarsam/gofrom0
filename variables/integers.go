package variables

import "fmt"

func ShowIntegers() {
	var intCommon int
	intFrom32 := int32(10)
	intFrom64 := int64(100)

	fmt.Println("intCommon = ", intCommon)
	fmt.Println("intFrom32 = ", intFrom32)
	fmt.Println("intFrom64 = ", intFrom64)

}
