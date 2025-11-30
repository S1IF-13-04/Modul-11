package main

import "fmt"

func main(){
	var jenis string
	var durasi, harga int
	fmt.Scan(&jenis, &durasi)
	
	switch jenis {
		case "motor":
			harga = durasi*2000
		case "mobil":
			harga = durasi*5000
		case "truk":
			harga = durasi*8000
	}
	fmt.Println(harga)
}