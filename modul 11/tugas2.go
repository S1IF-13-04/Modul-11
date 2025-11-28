package main
import "fmt"

func main (){
	var kendaraan string
	var harga, jam float64
	fmt.Scan(&kendaraan, &jam)
	
	harga=0
	if jam < 1{
		jam = 1
	}

	switch kendaraan{
	case "motor":
		harga = jam * 2000
	case "mobil":
		harga = 5000 * jam
	case "truk":
		harga= 8000 * jam
	default:
		fmt.Print("kendaraan tidak valid")
	}
	fmt.Println("Rp.", harga)
}