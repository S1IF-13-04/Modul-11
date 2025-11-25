package main

import "fmt"

func main() {
	var time12, time24 int
	var lable string
	fmt.Scan(&time24)
	switch {
	case time24 == 0:
		time12 = 12
		lable = "AM"
	case time24 < 12:
		time12 = time24
		lable = "AM"
	case time24 == 12:
		time12 = 12
		lable = "PM"
	case time24 > 12:
		time12 = time24 - 12
		lable = "PM"
	}

	fmt.Print(time12, lable)
}
