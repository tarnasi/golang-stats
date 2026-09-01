package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var luckyNumber string

func make_hash(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashByte := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashByte)
	fmt.Println("Orginal Text: ", input)
	return hashString
}

func lucky_game() {
	fmt.Println("Wlecome, Enter your lucky number: ")
	fmt.Scan(&luckyNumber)

	if luckyNumber == "8" {
		fmt.Println("Your lucky number is nothing you should rely on")
	} else {
		fmt.Println("Your are unlucky sorry choose another time ):")
	}
}

func fuel_live() {
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

func temperature_status() {
	var temperature float64

	fmt.Print("Temperature: ")
	fmt.Scan(&temperature)

	if temperature <= 60.0 {
		fmt.Println("STATUS: NORMAL")
	} else if temperature <= 80.0 {
		fmt.Println("STATUS: WARNING")
	} else {
		fmt.Println("STATUS: CRITICAL")
	}
}

func main() {
	var requestsNumber int
	var totalLatency float64

	fmt.Print("Number of requests: ")
	fmt.Scan(&requestsNumber)
	var latencyInput float64

	for rn := range requestsNumber {
		step := rn + 1
		fmt.Printf("Request %v latency: ", step)
		fmt.Scan(&latencyInput)
		totalLatency += latencyInput
	}

	avg := totalLatency / float64(requestsNumber)
	fmt.Printf("\nTotal latency: %.2f ms\n", totalLatency)
	fmt.Printf("Average latency: %.2f ms\n", avg)

	if avg <= 100 {
		fmt.Println("STATUS: FAST")
	} else if avg <= 300 {
		fmt.Println("STATUS: ACCEPTABLE")
	} else {
		fmt.Println("STATUS: SLOW")
	}
}
