package main

import "fmt"

func main(){
	var pH float64
	fmt.Scan(&pH)

	switch {
	case pH < 0 || pH > 14:
		fmt.Println("Tidak Valid")
	
	case pH >= 6.5 && pH <= 8.6:
		fmt.Println("Air Layak Minum")
	
	default:
		fmt.Println ("Air Tidak Layak Minum")
	}
	
}