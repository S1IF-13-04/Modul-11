package main

import "fmt"

func main() {
	var jenis string
	var durasi, tarifperjam int

	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&jenis)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	if durasi < 1 {
		durasi = 1
	}

	switch jenis {
	case "motor":
		tarifperjam = 2000
	case "mobil":
		tarifperjam = 5000
	case "truk":
		tarifperjam = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}

	total := tarifperjam * durasi
	fmt.Println("Total biaya parkir:", "Rp", total)
}
