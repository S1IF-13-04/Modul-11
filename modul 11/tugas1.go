package main
import "fmt"


func main (){
var ph float64
fmt.Scan(&ph)

switch{
case ph >= 6.5 && ph <= 8.6:
	fmt.Println("Air layak minum")
case ph <0 || ph>14:
	fmt.Println("Nilai ph tidak valid")
case ph<6.5 || ph > 8.6:
	fmt.Println("Air tidak layak diminum")
} 
}