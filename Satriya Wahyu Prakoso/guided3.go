package main
import "fmt"
func main() {
	var kendaraan string
	var tarif, durasi int
	fmt.Scan(&kendaraan, &durasi)
	switch {
		case kendaraan == "Motor" && durasi <=2:
			tarif = 7000
		case kendaraan == "Motor" && durasi > 2:
			tarif = 9000
		case kendaraan == "Mobil" && durasi <= 2:
			tarif = 15000
		case kendaraan == "Mobil" && durasi > 2:
			tarif = 20000
		case kendaraan == "Truk" && durasi <=2:
			tarif = 25000
		case kendaraan == "Truk" && durasi > 2:
			tarif = 35000
		default:
			tarif = 0
		}
		fmt.Println("Masukkan jenis kendaraan (Motor/Mobil/Truk):", kendaraan)
		fmt.Println("Masukkan durasi parkir (dalam jam):", durasi)
		fmt.Println("Tarif parkir: Rp.", tarif)
}