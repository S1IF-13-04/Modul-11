package main

import "fmt"

func main() {
	var kendaraan string
	var jam, tarif int

	fmt.Scan(&kendaraan, &jam)

	switch kendaraan {
	case "motor":
		tarif = jam * 2000
	case "mobil":
		tarif = jam * 5000
	case "truk":
		tarif = jam * 8000
	default:
		fmt.Println("Jenis kendaraan tidak dikenali")
	}
	fmt.Println("Rp", tarif)
}
