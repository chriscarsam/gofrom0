package exercises

import "strconv"

func ConverValue(value string) (int, string) {
	if num, err := strconv.Atoi(value); err != nil {
		return 0, "Error converting string to integer" + err.Error()
	} else {
		switch {
		case num > 100:
			return num, "is greater than 100"
		case num < 100:
			return num, "is less than 100"
		default:
			return num, "equals 100"
		}
	}
}
