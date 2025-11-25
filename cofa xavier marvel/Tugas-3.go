package main

import "fmt"

func main() {
	var num, num1 int
	var Category string
	fmt.Scan(&num)

	switch {
	case num%10 == 0 && num > 10:
		Category = "Multiple of 10"
		num1 = num / 10
	case num%5 == 0 && num > 10:
		Category = "Multiple of 5"
		num1 = num * num
	}
	switch {
	case num%2 != 0 && num < 10:
		Category = "Even Number"
		num1 = num + (num + 1)
	case num%2 == 0 && num < 10:
		Category = "Odd Number"
		num1 = num * (num + 1)
	}
	fmt.Printf("Category: %s \nResult: %d", Category, num1)
}
