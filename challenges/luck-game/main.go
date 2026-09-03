package main

import "fmt"

func main() {
	var luckyNumber string

	fmt.Println("Wlecome, Enter your lucky number: ")
	fmt.Scan(&luckyNumber)

	if luckyNumber == "8" {
		fmt.Println("Your lucky number is nothing you should rely on")
	} else {
		fmt.Println("Your are unlucky sorry choose another time ):")
	}
}