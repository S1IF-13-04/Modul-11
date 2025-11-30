package main

import "fmt"

func main() {
	var jenisKendaraan string
	var durasi int

	fmt.Print("Masukkan jenis kendaraan dan durasi parkir: ")
	fmt.Scan(&jenisKendaraan, &durasi)

	if durasi < 1 {
		durasi = 1
	}

	var tarifPerJam int
	var totalBiaya int

	switch jenisKendaraan {
	case "motor":
		tarifPerJam = 2000
	case "mobil":
		tarifPerJam = 5000
	case "truk":
		tarifPerJam = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}

	totalBiaya = tarifPerJam * durasi
	fmt.Printf("Rp %d\n", totalBiaya)
}
