package main

import "fmt"

func main() {
	var bilangan, hasil, bilanganBerikutnya int
	fmt.Scan(&bilangan)

	switch {
	case bilangan%10 == 0:
		hasil = bilangan / 10
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Println("Hasil pembagian antara", bilangan, "/ 10 =", hasil)

	case bilangan%2 == 0:
		bilanganBerikutnya = bilangan + 1
		hasil = bilangan * bilanganBerikutnya
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Println("Hasil perkalian dengan bilangan berikutnya", bilangan, "*", bilanganBerikutnya, "=", hasil)

	case bilangan%5 == 0 && bilangan != 5:
		hasil = bilangan * bilangan
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Println("Hasil kuadrat dari", bilangan, "^2 =", hasil)

	default:
		bilanganBerikutnya = bilangan + 1
		hasil = bilangan + bilanganBerikutnya
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya", bilangan, "+", bilanganBerikutnya, "=", hasil)
	}
}
