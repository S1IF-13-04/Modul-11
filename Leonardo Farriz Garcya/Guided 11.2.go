package main 
import "fmt" 
func main() { 

var tanaman string 
fmt.Scan(&tanaman) 

switch tanaman { 
	case "nepenthes", "drosera": 
	fmt.Println("Termasuk Tanaman Karnivora.") 
	fmt.Println("Asli Indonesia.") 
	
	case "venus", "sarracenia": 
	fmt.Println("Termasuk Tanaman Karnivora.") 
	fmt.Println("Tidak Asli Indonesia.") 

	default: 
	fmt.Println("Tidak termasuk Tanaman Karnivora.") 
	} 
} 