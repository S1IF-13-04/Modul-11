package main

import "fmt"

func main() {
	var jam int
	var jamFormat int
	var perihal string

	fmt.Print("Masukkan jam (0-23): ")
	fmt.Scan(&jam)

	switch jam {
	case 0:
		jamFormat = 12
		perihal = "AM"
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		jamFormat = jam
		perihal = "AM"
	case 12:
		jamFormat = 12
		perihal = "PM"
	case 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
		jamFormat = jam - 12
		perihal = "PM"
	default:
		fmt.Println("Jam tidak valid! Masukkan jam 0-23")
		return
	}

	fmt.Printf("%d:00 = %d %s\n", jam, jamFormat, perihal)
}
