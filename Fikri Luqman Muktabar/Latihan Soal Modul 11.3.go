package main

import "fmt"

func main() {
    var bilangan int
    fmt.Scan(&bilangan)

    switch {
    case bilangan % 10 == 0: 
        fmt.Println("Kategori: Bilangan Kelipatan 10")
        fmt.Print("Hasil pembagian antara ", bilangan, " / 10 = ", bilangan / 10)

    case bilangan % 5 == 0 && bilangan != 5: 
        fmt.Println("Kategori: Bilangan Kelipatan 5")
        fmt.Print("Hasil kuadrat dari ", bilangan, " ^2 = ", bilangan * bilangan)

    case bilangan % 2 == 0: 
        fmt.Println("Kategori: Bilangan Genap")
        fmt.Print("Hasil perkalian dengan bilangan berikutnya ", bilangan, " * ", bilangan + 1, " = ", bilangan * (bilangan + 1))
		
	case bilangan % 2 != 0: 
        fmt.Println("Kategori: Bilangan Ganjil")
        fmt.Print("Hasil penjumlahan dengan bilangan berikutnya ", bilangan, " + ", bilangan + 1, " = ", bilangan + (bilangan + 1))
    }
}