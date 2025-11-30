package main

import "fmt"

func main() {
	var jenis string
	var durasi int
	var tarifPerJam int
	var total int

	fmt.Scan(&jenis)
	fmt.Scan(&durasi)

	if durasi < 1 {
		durasi = 1
	}
	switch jenis {
	case "motor":
		tarifPerJam = 2000
	case "mobil":
		tarifPerJam = 5000
	case "truk":
		tarifPerJam = 8000
	default:
		tarifPerJam = 0
	}

	total = tarifPerJam * durasi
	fmt.Println("Rp", total)
}
