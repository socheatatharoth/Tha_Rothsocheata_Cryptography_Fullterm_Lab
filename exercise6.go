package main

import (
	"fmt"
)

// xorEncrypt decrypts or encrypts a string using a single byte key
func xorEncrypt(text string, key byte) string {
	result := make([]byte, len(text)) // create a byte slice to store result

	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key // XOR each byte with the key
	}

	return string(result) // convert byte slice back to string
}

func main() {
	var plaintext string
	var keyInput int

	fmt.Print("Enter plaintext message: ")
	fmt.Scanln(&plaintext)

	fmt.Print("Enter key (0-255): ")
	fmt.Scanln(&keyInput)

	key := byte(keyInput) // convert int to byte

	// Encrypt
	ciphertext := xorEncrypt(plaintext, key)
	fmt.Println("Ciphertext:", ciphertext)

	// Decrypt
	decrypted := xorEncrypt(ciphertext, key)
	fmt.Println("Decrypted message:", decrypted)
}
