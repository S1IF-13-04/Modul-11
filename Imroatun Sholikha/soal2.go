package main

import "fmt"

func main() {
	var jenis string
	var durasi int
	var tarif int

	fmt.Scan(&jenis, &durasi)

	switch {
	case durasi < 1:
		durasi = 1
	default:
		durasi = durasi
	}

	switch jenis {
	case "motor":
		tarif = 2000
	case "mobil":
		tarif = 5000
	case "truk":
		tarif = 8000
	default:
		tarif = 0
	}

	total := tarif * durasi
	fmt.Println("Rp.", total)
}
