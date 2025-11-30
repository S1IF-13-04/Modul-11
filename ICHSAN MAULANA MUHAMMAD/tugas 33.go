package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	switch {
	case n%2 != 0:
		fmt.Printf("- Bilangan Ganjil (hasil %d + %d = %d)\n", n, n+1, n+(n+1))
	case n%2 == 0:
		fmt.Printf("- Bilangan Genap (hasil %d * %d = %d)\n", n, n+1, n*(n+1))
	}
	switch {
	case n%5 == 0:
		fmt.Printf("- Bilangan Kelipatan 5 (hasil %d^2 = %d)\n", n, n*n)
	}
	switch {
	case n%10 == 0:
		fmt.Printf("- Bilangan Kelipatan 10 (hasil %d / 10 = %d)\n", n, n/10)
	}
}
