package main

import "fmt"


func main() {
	var distance float64
	var consumption float64
	var price float64

	fmt.Println("Welcome to Fuel Live - just enter inputs and we calculate for you")

	fmt.Print("Distance: ")
	fmt.Scan(&distance)

	fmt.Print("Consumption: ")
	fmt.Scan(&consumption)

	fmt.Print("Price: ")
	fmt.Scan(&price)

	fuelNeeded := distance * consumption / 100
	totalCost := fuelNeeded * price

	fmt.Printf("\nFuel Needed: %.2f Liters\nTotal Cost: %.2f AED", fuelNeeded, totalCost)
}
