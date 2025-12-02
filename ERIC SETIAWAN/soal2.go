package main

import "fmt"

func main() {
	var Kendaraan string
	var durasi, tarifPerjam int
	fmt.Scan(&Kendaraan, &durasi)

	if durasi < 1 {
		durasi = 1
	}
	switch Kendaraan {
	case "motor":
		tarifPerjam = 2000
	case "mobil":
		tarifPerjam = 5000
	case "truk":
		tarifPerjam = 8000
	}
	fmt.Printf("Rp %d\n", tarifPerjam*durasi)
}
