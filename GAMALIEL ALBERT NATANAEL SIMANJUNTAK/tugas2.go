package main
import (
	"fmt"
	"strings"
)
func main() {
	var jenisKendaraan string
	var durasiJam int
	fmt.Print("Masukan jenis kendaraan dan durasi: ")
	fmt.Scan(&jenisKendaraan, &durasiJam)
	jenisKendaraan = strings.ToLower(jenisKendaraan)
	if durasiJam < 1 {
		durasiJam = 1
	}
	var tarifPerJam int
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
	totalBiaya := tarifPerJam * durasiJam
	fmt.Printf("Rp %d\n", totalBiaya)
}