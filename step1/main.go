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

func main() {
	// MAKE A HASH SHA256(HEXSTRING) - (1)
	// input := "In the name of GOD"
	// fmt.Println("Sha256: ", make_hash(input))


	// Lucky Game ;) - (2)
	// lucky_game()

	// A Simple APP
	
}
