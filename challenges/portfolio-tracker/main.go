package main

import (
	"fmt"
)

type Asset struct {
	Symbol   string
	Quantity float64
	Price    float64
}

func main() {
	var numberOfAssets int
	portfolio := map[string]Asset{}

	fmt.Println("--- Welcome To Portfolio Tracker ---")

	// Initial
	fmt.Print("\nHow many assets? ")
	fmt.Scan(&numberOfAssets)

	for step := range numberOfAssets {
		asset := readAsset(step + 1)
		portfolio[asset.Symbol] = asset
	}

	printAssets(portfolio)
}

func readAsset(step int) Asset {
	var asset Asset

	fmt.Printf("\nAsset %d Symbol: ", step)
	fmt.Scan(&asset.Symbol)

	fmt.Print("Quantity: ")
	fmt.Scan(&asset.Quantity)

	fmt.Print("Price: ")
	fmt.Scan(&asset.Price)

	return asset
}

func printAssets(portfolio map[string]Asset) {
	fmt.Print("\n--- Portfolio ---\n\n")
	var total float64
	for Symbol, asset := range portfolio {
		portfolioValue := asset.Price * asset.Quantity
		fmt.Printf("%v: $%.2f\n", Symbol, portfolioValue)
		total += portfolioValue
	}
	fmt.Printf("\nTotal Portfolio Value: $%.2f\n", total)
}
