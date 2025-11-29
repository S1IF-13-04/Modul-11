package main

import "fmt"

func main() {
	var pH float64

	fmt.Print("Masukkan kadar pH: ")
	fmt.Scan(&pH)

	if pH < 0 || pH > 14 {
		fmt.Println("Input tidak valid, rentang pH 0 - 14")
		return
	}

	var kategori int
	switch {
	case pH < 6.5:
		kategori = 0
	case pH >= 6.5 && pH <= 8.6:
		kategori = 1
	case pH > 8.6:
		kategori = 2
	}

	switch kategori {
	case 0:
		fmt.Println("Air tidak layak minum")
	case 1:
		fmt.Println("Air layak minum")
	case 2:
		fmt.Println("Air tidak layak minum")
	}
}