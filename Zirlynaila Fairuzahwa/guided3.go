package main
import "fmt"

func main(){
	var jenis string
	var jam, tarif int
	fmt.Print("Masukkan jenis kendaraan: ")
	fmt.Scan(&jenis)
	fmt.Print("Masukkan durasi parkir: ")
	fmt.Scan(&jam)
	switch jenis {
	case "Motor":
		if jam > 2 {
			tarif = 9000
		} else {
			tarif = 7000
		}
	case "Mobil":
		if jam > 2 {
			tarif = 20000
		} else {
			tarif = 15000
		}
	case "Truk":
		if jam > 2 {
			tarif = 35000
		} else {
			tarif = 25000
		}
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Print("Tarif Parkir: Rp ", tarif)
}