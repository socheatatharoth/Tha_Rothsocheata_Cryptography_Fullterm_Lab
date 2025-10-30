package crack

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
)

// SHA512Hash returns the lowercase hex-encoded SHA512 of the input string.
func SHA512Hash(s string) string {
	sum := sha512.Sum512([]byte(s))
	return hex.EncodeToString(sum[:])
}

// CrackSHA512 scans the wordlist and returns the plaintext if the SHA512 hash matches.
// It prints a verbose line for every attempt using fmt.Printf (so main.go's MultiWriter will capture it).
func CrackSHA512(target, wordlistPath string) (string, error) {
	f, err := os.Open(wordlistPath)
	if err != nil {
		return "", fmt.Errorf("open wordlist: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		word := scanner.Text()
		h := SHA512Hash(word)
		fmt.Printf("Line %d: trying '%s' -> %s\n", lineNo, word, h)
		if h == target {
			return word, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scan wordlist: %w", err)
	}
	return "", nil
}
