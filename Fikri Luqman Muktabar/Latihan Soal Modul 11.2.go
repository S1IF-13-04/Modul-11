package main

import "fmt"

func main() {
	var jenis_kendaraan string
	var jumlah_jam, total_biaya int
	
	fmt.Scan(&jenis_kendaraan, &jumlah_jam)

	switch jenis_kendaraan {
	case "motor":
		total_biaya = jumlah_jam * 2000
	case "mobil":
		total_biaya = jumlah_jam * 5000
	case "truk":
		total_biaya = jumlah_jam * 8000 
	}
	fmt.Print("Rp ", total_biaya)
}