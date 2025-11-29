package main
import "fmt"
func main(){
	var a, indeks int
	fmt.Scan(&a)
	switch {
	case a % 10 == 0:
		fmt.Println("Kategori : kelipatan 10")
		indeks = a / 10
		fmt.Println("Hasil pembagian antara", a, "/", 10, "=", indeks)
	case a % 5 == 0 && a != 5:
		fmt.Println("Kategori : bilangan kelipatan 5")
		indeks = a * a
		fmt.Println("Hasil kuadrat dari", a, "^ 2 =", indeks)
	case a % 2 != 0:
		fmt.Println("Kategori : Bilangan ganjil")
		indeks = a + (a + 1)
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya", a, "+", (a + 1), "=", indeks)
	default:
		fmt.Println("Kategori : Bilangan genap")
		indeks = a * (a + 1)
		fmt.Println("Hasil perkalian dengan bilangan berikutnya", a, "*", (a + 1), "=", indeks)
	}

}