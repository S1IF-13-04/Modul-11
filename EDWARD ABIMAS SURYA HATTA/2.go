package main

import "fmt"

func main() {
	var kendaraan string
	var durasi, tarif int

	fmt.Scan(&kendaraan, &durasi)

	if durasi < 1 {
		durasi = 1
	}

	switch kendaraan {
	case "motor":
		tarif = 2000 * durasi
	case "mobil":
		tarif = 5000 * durasi
	case "truk":
		tarif = 8000 * durasi
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}

	fmt.Printf("Rp %d\n", tarif)
}