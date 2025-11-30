package main
import "fmt"
func main() {
	var tanaman string
	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthes":
		fmt.Println("termasuk tanaman karnivora")
		fmt.Println("asli indonesia")
	case "venus":
		fmt.Println("termasuk tanaman karnivora")
		fmt.Println("bukan asli indonesia")
	default:
		fmt.Println("tidak termasuk tanaman karnivora")				
	}
}