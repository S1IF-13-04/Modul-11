package main

import "fmt"

func main() {
	var bil int
	fmt.Scan(&bil)

	switch bil {
	case 0:
		fmt.Print("12 AM")
	case 12:
		fmt.Print("12 PM")
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		fmt.Print(bil, "AM")
	case 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
		bil = bil - 12
		fmt.Print(bil, "PM")
	default:
		fmt.Print("tidak falid")
		return
	}
}
