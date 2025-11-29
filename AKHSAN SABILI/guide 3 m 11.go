package main
import "fmt"
func main(){
	var kendaraan string
	var waktu, harga int
	fmt.Print("jenis kendaraan :")
	fmt.Scan(&kendaraan)
	fmt.Print("waktu parkir :")
	fmt.Scan(&waktu)
	switch {
	case kendaraan == "motor" && waktu <= 2:
		harga = 7000
	case kendaraan == "motor" && waktu > 2:
		harga = 9000
	case kendaraan == "mobil" && waktu <= 2:
		harga = 15000
	case kendaraan == "mobil" && waktu > 2:
		harga = 20000
	case kendaraan == "truk" && waktu <= 2:
		harga = 25000
	case kendaraan == "truk" && waktu > 2:
		harga = 35000
	default:
		harga = 0
	}
	fmt.Println("Tarif parkir :", harga)
}