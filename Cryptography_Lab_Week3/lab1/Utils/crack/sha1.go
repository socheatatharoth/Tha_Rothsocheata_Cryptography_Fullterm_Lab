package crack

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

// SHA1Hash returns lowercase hex-encoded SHA1 of s.
func SHA1Hash(s string) string {
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// CrackSHA1 scans the wordlist and returns the plaintext password if found.
// It prints a verbose line for every attempt via the standard logger (log.Println)
// so main.go's log.SetOutput(...) will pick it up.
func CrackSHA1(target, wordlistPath string) (string, error) {
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
		h := SHA1Hash(word)
		// Verbose: use fmt.Printf or log.Printf from caller; here use fmt so caller's MultiWriter still captures it
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
