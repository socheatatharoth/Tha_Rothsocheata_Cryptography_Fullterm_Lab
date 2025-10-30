package crack

import (
    "bufio"
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "os"
    "strings"
    "path/filepath"
)

// CrackMD5 tries each word in the wordlist and returns plaintext if matched.
func CrackMD5(targetHash string, wordlistPath string) (string, error) {
    // Normalize path (helps with ./ and backslashes)
    wordlistPath = filepath.Clean(wordlistPath)

    file, err := os.Open(wordlistPath)
    if err != nil {
        return "", fmt.Errorf("failed to open wordlist: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := strings.TrimSpace(scanner.Text()) // remove \r, spaces, etc.
        if word == "" {
            continue
        }
        hash := md5.Sum([]byte(word))
        hashStr := hex.EncodeToString(hash[:])

        if hashStr == targetHash {
            return word, nil
        }
    }
    if err := scanner.Err(); err != nil {
        return "", fmt.Errorf("error reading wordlist: %v", err)
    }
    return "", fmt.Errorf("password not found")
}
