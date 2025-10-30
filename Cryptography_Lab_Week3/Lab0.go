package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/sha3"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("======== Name + Hashing Program ========")
	fmt.Print("Please input value 1: ")
	input1, _ := reader.ReadString('\n')
	fmt.Print("Please input value 2: ")
	input2, _ := reader.ReadString('\n')

	// Trim newline characters
	input1 = strings.TrimSpace(input1)
	input2 = strings.TrimSpace(input2)

	proofMe(input1, input2)
}

func proofMe(txt1, txt2 string) {
	md5a := md5.Sum([]byte(txt1))
	md5b := md5.Sum([]byte(txt2))
	hashCompare("MD5", md5a[:], md5b[:])

	sha1a := sha1.Sum([]byte(txt1))
	sha1b := sha1.Sum([]byte(txt2))
	hashCompare("SHA1", sha1a[:], sha1b[:])

	sha256a := sha256.Sum256([]byte(txt1))
	sha256b := sha256.Sum256([]byte(txt2))
	hashCompare("SHA256", sha256a[:], sha256b[:])

	sha512a := sha512.Sum512([]byte(txt1))
	sha512b := sha512.Sum512([]byte(txt2))
	hashCompare("SHA512", sha512a[:], sha512b[:])

	sha3a := sha3.Sum256([]byte(txt1))
	sha3b := sha3.Sum256([]byte(txt2))
	hashCompare("SHA3-256", sha3a[:], sha3b[:])
}

func hashCompare(name string, hash1, hash2 []byte) {
	// Convert to hex string
	hashStr1 := hex.EncodeToString(hash1)
	hashStr2 := hex.EncodeToString(hash2)

	match := "No Match!"
	if hashStr1 == hashStr2 {
		match = "Match!"
	}

	fmt.Printf("Hash (%s):\n  Output A: %s\n  Output B: %s\n  => %s\n\n",
		name, hashStr1, hashStr2, match)
}
