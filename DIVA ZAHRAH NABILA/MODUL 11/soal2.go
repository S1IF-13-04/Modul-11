package main

import "fmt"

func main() {
	var kendaraan string
	var jam int
	fmt.Scan(&kendaraan, &jam)
	if jam < 1 {
		jam = 1
	}
	tarif := 0
	switch kendaraan {
	case "Motor":
		tarif = 2000
	case "Mobil":
		tarif = 5000
	case "Truk":
		tarif = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid!")
		return
	}
	total := tarif * jam
	fmt.Print("Rp", total)
}
