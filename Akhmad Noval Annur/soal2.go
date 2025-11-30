package main

import "fmt"

func main() {
	var kendaraan string
	var durasi, total int

	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	if durasi < 1 {
		durasi = 1
	}

	switch kendaraan {
	case "motor":
		total = 2000 * durasi
	case "mobil":
		total = 5000 * durasi
	case "truk":
		total = 8000 * durasi
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		
	}
	fmt.Printf("Tarif Parkir: Rp %d\n", total)
}