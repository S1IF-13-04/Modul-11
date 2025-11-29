package main

import "fmt"

func main() {
	var jenisKendaraan string
	var durasi int
	var tarifPerJam, totalBiaya int

	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&jenisKendaraan)

	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scan(&durasi)

	if durasi < 1 {
		durasi = 1
	}

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

	fmt.Printf("Total biaya parkir: Rp %d\n", totalBiaya)
}
