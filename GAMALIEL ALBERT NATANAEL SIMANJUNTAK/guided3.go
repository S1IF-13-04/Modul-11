package main
import "fmt"
func main() {
	var kendaraan string
	var jam, tarif int
	fmt.Scan(&kendaraan, &jam)
	switch kendaraan {
	case "Motor":
		switch jam {
		case 1, 2:
			tarif = 7000
		default:
			tarif = 9000
		}
	case "Mobil":
		switch jam {
		case 1, 2:
			tarif = 15000
		default:
			tarif = 20000
		}
	case "Truk":
		switch jam {
		case 1, 2:
			tarif = 25000
		default:
			tarif = 35000
		}
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Println("Tarif Parkir: Rp", tarif)
}