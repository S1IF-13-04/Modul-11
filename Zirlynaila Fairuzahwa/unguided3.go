package main
import "fmt"

func main(){
	var bilangan, hasil int
	var kategori string
	fmt.Scan(&bilangan)
	switch {
	case bilangan % 10 == 0:
		kategori = "Bilangan Kelipatan 10"
		hasil = bilangan / 10
		fmt.Println("Kategori:", kategori)
		fmt.Printf("Hasil pembagian antara %d / 10 = %d", bilangan, hasil)
	case bilangan % 5 == 0:
		if bilangan == 5 {
			kategori = "Bilangan Ganjil"
			hasil = bilangan + (bilangan + 1)
			fmt.Println("Kategori:", kategori)
			fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d", bilangan, bilangan + 1, hasil)
		} else {
			kategori = "Bilangan Kelipatan 5"
			hasil = bilangan * bilangan
			fmt.Println("Kategori:", kategori)
			fmt.Printf("Hasil kuadrat dari %d ^ 2 = %d", bilangan, hasil)
		}
	case bilangan % 2 == 1:
		kategori = "Bilangan Ganjil"
		hasil = bilangan + (bilangan + 1)
		fmt.Println("Kategori:", kategori)
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d", bilangan, bilangan + 1, hasil)
	case bilangan % 2 == 0:
		kategori = "Bilangan Genap"
		hasil = bilangan * (bilangan + 1)
		fmt.Println("Kategori:", kategori)
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d", bilangan, bilangan + 1, hasil)
	}
}
