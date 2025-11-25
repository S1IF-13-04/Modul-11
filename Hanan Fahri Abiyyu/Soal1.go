package main

import "fmt"

func main() {
	var ph float64
	var teks string
	fmt.Scan(&ph)

	switch {
		case ph >= 6.5 && ph <= 8.6:
			teks = "Air layak minum"
		case ph >= 0 && ph < 6.5 || ph > 8.6 && ph <= 14 :
			teks = "Air tidak layak minum"
		default:
			teks = "Nilai pH tidak valid. Nilai pH harus antara 0 dan 14"
	}
	fmt.Println(teks)
}
