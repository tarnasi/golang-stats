package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func make_hash(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashByte := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashByte)
	fmt.Println("Orginal Text: ", input)
	return hashString
}

func main() {

}
