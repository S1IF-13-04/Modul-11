package main
import "fmt"

func main (){
	var kendaraan string
	var jam, harga float64
	fmt.Scan(&kendaraan,&jam)

	switch kendaraan {
	case "Motor":
		if jam<=2 {
			harga = 7000
		}else {
			harga= 9000
		}
	
	case "Mobil" :
		if jam<=2 {
			harga = 15000
		} else {
			harga = 20000
		}
	case "Truk":
		if jam<=2 {
			harga =25000
		} else {
			harga=35000
		}
	default:
		fmt.Println("jam atau kendaraan tidak valid")
		harga=0
	}
	fmt.Println(harga)
}