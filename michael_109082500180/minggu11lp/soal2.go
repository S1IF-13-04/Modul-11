package main

import "fmt"


func main() {
	var harga,waktu int64
	var jenis string
	fmt.Scan(&jenis,&waktu)

	switch jenis{
	case "motor" :
		harga = 2000 * waktu
		fmt.Println(harga)

	case "mobil":
		harga = 5000 * waktu
		fmt.Println(harga)

	case "truk":
		harga = 8000 * waktu
		fmt.Println(harga)
		
	default:
		fmt.Printf("tidak falid")
	} 
}
