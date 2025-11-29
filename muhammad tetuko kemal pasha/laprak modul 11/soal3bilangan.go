package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilangan)

	// Tentukan kategori bilangan
	var kategori int
	// 1 = Kelipatan 10, 2 = Kelipatan 5, 3 = Ganjil, 4 = Genap
	switch {
	case bilangan%10 == 0:
		kategori = 1
	case bilangan%5 == 0:
		kategori = 2
	case bilangan%2 == 1 || bilangan%2 == -1:
		kategori = 3
	case bilangan%2 == 0:
		kategori = 4
	}

	// Switch case untuk output berdasarkan kategori
	switch kategori {
	case 1:
		hasil := bilangan / 10
		fmt.Printf("Kategori: Bilangan Kelipatan 10\n")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", bilangan, hasil)
	case 2:
		hasil := bilangan * bilangan
		fmt.Printf("Kategori: Bilangan Kelipatan 5\n")
		fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n", bilangan, hasil)
	case 3:
		hasil := bilangan + (bilangan + 1)
		fmt.Printf("Kategori: Bilangan Ganjil\n")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, hasil)
	case 4:
		hasil := bilangan * (bilangan + 1)
		fmt.Printf("Kategori: Bilangan Genap\n")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, hasil)
	}
}
