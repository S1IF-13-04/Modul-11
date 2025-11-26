package main

import "fmt"

func main() {
	var pH float64
	fmt.Scan(&pH)
	switch {
	case pH >= 6.5 && pH <= 8.6:
		fmt.Print("Air layak minum")
	case (pH >= 0 && pH < 6.5) || (pH > 8.6 && pH <= 14):
		fmt.Print("Air tidak layak minum")
	default:
		fmt.Print("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14")
	}
}
