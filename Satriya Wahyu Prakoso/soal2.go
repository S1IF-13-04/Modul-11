package main
import "fmt"
func main() {
	var kendaraan string
	var tarif, durasi float64
	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)
	switch kendaraan {
	case "motor":
		tarif = durasi * 2000
	case "mobil":
		tarif = durasi * 5000
	case "truk":
		tarif = durasi * 8000
	}
	fmt.Printf("Tarif Parkir: Rp %.0f\n", tarif)
}