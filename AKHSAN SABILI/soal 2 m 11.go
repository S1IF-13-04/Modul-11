package main
import "fmt"
func main(){
	var kendaraan,jam string
	var waktu, parkir int
	fmt.Scan(&kendaraan, &waktu, &jam)
	switch {
	case kendaraan == "motor":
		parkir = waktu * 2000
	case kendaraan == "mobil":
		parkir = waktu * 5000
	case kendaraan == "truk":
		parkir = waktu * 8000
	}
	fmt.Println(parkir)
}