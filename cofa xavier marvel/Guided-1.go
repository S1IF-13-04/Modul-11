package main

import "fmt"

func main() {
	var time12, time24 int
	var lable string
	fmt.Scan(&time24)
	switch {
	case time24 > 12:
		time12 = time24 - 12
		lable = " PM"
	default:
		time12 = time24
		lable = " AM"
	}

	fmt.Print(time12, lable)
}
