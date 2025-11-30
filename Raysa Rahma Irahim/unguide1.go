package main

import "fmt"

func main() {
	var ph float64
	fmt.Scan(&ph)

	switch {
	case ph < 0 && ph < 14:
		fmt.Println("Input tidak valid, rentang pH 0 - 14")

	case ph > 14 && ph > 0:
		fmt.Println("Input tidak valid, rentang pH 0 - 14")

	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("Air layak minum")

	case ph < 6.5 && ph >= 0:
		fmt.Println("Air tidak layak minum")

	case ph > 8.6 && ph <= 14:
		fmt.Println("Air tidak layak minum")
	}
}
