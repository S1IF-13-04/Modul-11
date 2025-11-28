package main
import "fmt"

func main(){
	var jenis string
	var durasi, tarif float64
	fmt.Scan(&jenis, &durasi)
	if durasi < 1 {
		durasi = 1
	}
	switch jenis {
	case "motor":
		tarif = durasi * 2000
	case "mobil":
		tarif = durasi * 5000
	case "truk":
		tarif = durasi * 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid")
	}
	fmt.Print("Rp ", tarif)
}