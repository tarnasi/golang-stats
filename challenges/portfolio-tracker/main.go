package main

import (
	"fmt"
)


type Asset struct {
	Symbol string
	Quantity  float64
	Price  float64
}

func main() {
	var numberOfAssets int
	portfolio := map[string]Asset{}

	fmt.Println("--- Welcome To Portfolio Tracker ---")

	// Initial
	fmt.Print("\nHow many assets? ")
	fmt.Scan(&numberOfAssets)

	for step := range numberOfAssets {
		asset :=readAsset(step+1)
		portfolio[asset.Symbol] = asset
	}

	printAssets(portfolio)
}

func readAsset(step int) Asset {
	var asset Asset

	fmt.Printf("Asset %d Symbol: ", step)
	fmt.Scan(&asset.Symbol)

	fmt.Print("Quantity: ")
	fmt.Scan(&asset.Quantity)

	fmt.Print("Price: ")
	fmt.Scan(&asset.Price)

	return asset
}

func printAssets(portfolio map[string]Asset) {
	for Symbol, asset := range portfolio {
		fmt.Printf("\n%v: %.2f\n", Symbol, asset.Price * asset.Quantity)
	}
}
