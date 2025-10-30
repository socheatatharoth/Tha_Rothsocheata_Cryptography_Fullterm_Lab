package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/socheatatharoth/Cryptography/Cryptography_Lab_Week3/lab1/utils/crack"
)

// createVerboseFile tries to create the primary verbose file and falls back to
// a timestamped filename if the primary cannot be created.
func createVerboseFile(base string) (*os.File, string, error) {
	base = filepath.Clean(base)

	// Try primary filename first
	f, err := os.Create(base)
	if err == nil {
		return f, base, nil
	}

	// If cannot create (e.g., locked), fall back to a timestamped file
	timestamp := time.Now().Format("20060102_150405")
	fallback := fmt.Sprintf("%s_%s.log", base[:len(base)-len(filepath.Ext(base))], timestamp)
	f2, err2 := os.Create(fallback)
	if err2 != nil {
		// return original error and name attempt
		return nil, base, fmt.Errorf("create %s: %v (fallback %s failed: %v)", base, fallback, err2)
	}
	return f2, fallback, nil
}

func main() {
	// Flags
	mode := flag.String("mode", "md5", "cracking mode: md5 | sha1 | sha512")
	target := flag.String("target", "", "target hash to crack")
	wordlist := flag.String("wordlist", "wordlist.txt", "path to wordlist file")
	verboseName := flag.String("verbose", "verbose.txt", "verbose output filename (optional)")

	flag.Parse()

	// DEBUG: print parsed flags (temporary — helpful when running)
	fmt.Printf("DEBUG: parsed flags -> mode=%q, target=%q, wordlist=%q, verbose=%q\n", *mode, *target, *wordlist, *verboseName)

	if *target == "" {
		fmt.Println("Usage: go run main.go -mode=[md5|sha1|sha512] -target=<hash> -wordlist=<wordlist.txt> [-verbose=verbose.txt]")
		os.Exit(1)
	}

	// Normalize inputs
	*mode = strings.ToLower(strings.TrimSpace(*mode))
	*target = strings.ToLower(strings.TrimSpace(*target))
	*wordlist = filepath.Clean(strings.TrimSpace(*wordlist))
	*verboseName = filepath.Clean(strings.TrimSpace(*verboseName))

	// Create verbose file (with fallback if locked)
	f, fname, err := createVerboseFile(*verboseName)
	if err != nil {
		log.Fatalf("Failed to create verbose file: %v", err)
	}
	defer f.Close()

	// Make a writer that writes to stdout AND verbose file
	mw := io.MultiWriter(os.Stdout, f)
	// ensure the logger writes to the same writer
	log.SetOutput(mw)

	// Header
	fmt.Fprintf(mw, "=== Password cracking started (mode=%s, target=%s, wordlist=%s) ===\n", *mode, *target, *wordlist)
	fmt.Fprintf(mw, "Verbose log file: %s\n", fname)

	// Print the resolved wordlist path so failures are easier to diagnose.
	absWordlist, _ := filepath.Abs(*wordlist)
	log.Printf("Resolved wordlist path: %s", absWordlist)

	var password string
	var crackErr error

	switch *mode {
	case "md5":
		log.Println("🔍 Starting MD5 password cracking...")
		password, crackErr = crack.CrackMD5(*target, *wordlist)
	case "sha1":
		log.Println("🔍 Starting SHA1 password cracking...")
		password, crackErr = crack.CrackSHA1(*target, *wordlist)
	case "sha512":
		log.Println("🔍 Starting SHA512 password cracking...")
		password, crackErr = crack.CrackSHA512(*target, *wordlist)
	default:
		log.Fatalf("Unknown mode: %s. Use 'md5', 'sha1' or 'sha512'.", *mode)
	}

	// Record and print result (also goes to verbose file)
	if crackErr != nil {
		log.Println("❌ Error:", crackErr)
		fmt.Fprintf(mw, "❌ Error: %v\n", crackErr)
	} else if password != "" {
		log.Println("✅ Password found:", password)
		fmt.Fprintf(mw, "✅ Password found: %s\n", password)
	} else {
		log.Println("⚠️  Finished scanning wordlist: password not found.")
		fmt.Fprintf(mw, "⚠️  Finished scanning wordlist: password not found.\n")
	}

	fmt.Fprintln(mw, "=== Done ===")
}
