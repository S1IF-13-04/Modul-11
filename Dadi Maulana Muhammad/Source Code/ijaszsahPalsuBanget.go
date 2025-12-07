package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	switch {
	case x%10 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		hasil := x / 10
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", x, hasil)

	case x%5 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		hasil := x * x
		fmt.Printf("Hasil kuadrat dari %d^2 = %d\n", x, hasil)

	case x%2 != 0:
		fmt.Println("Kategori: Bilangan Ganjil")
		hasil := x + (x + 1)
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", x, x+1, hasil)

	case x%2 == 0 && x%10 != 0 && x%5 != 0:
		fmt.Println("Kategori: Bilangan Genap")
		hasil := x * (x + 1)
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", x, x+1, hasil)
	}
}
