package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	generateRandomState()

	fmt.Printf("\n\n\n")
	generateRandomState()

	fmt.Printf("\n\n\n")
	generateRandomState()
}

func generateRandomState() {
	b := make([]byte, 16)
	fmt.Println("Before  :", b)

	rand.Read(b)
	fmt.Println("After   :", b)

	state := base64.URLEncoding.EncodeToString(b)
	fmt.Println("Encoded :", state)
}
