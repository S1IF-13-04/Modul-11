package main

import "fmt"

func main() {
	var hasil, angka int64
	fmt.Scan(&angka)

	switch{
	case angka % 10 == 0 && angka != 10:
		hasil = angka / 10
		fmt.Println("Bilangan Kelipatan 10")
		fmt.Println(hasil)

	case angka % 5 == 0 && angka != 5:
		hasil = angka * angka
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Println(hasil)


	case angka % 2 == 0:
		hasil = angka * (angka + 1)
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Println(hasil)

	default:
		hasil = angka + angka + 1
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Println(hasil)

	}
}
