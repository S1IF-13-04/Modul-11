package main
import "fmt"
func main() {
	var Tanaman string
	fmt.Scan(&Tanaman)
	switch Tanaman {
	case "nepenthes":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Asli Indonesia")
	case "venus":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Bukan Asli Indonesia")
	default:
		fmt.Println("Tidak termasuk Tanaman Karnivora")
	}
}