package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

func findHash(sums, hash string) bool {
	return strings.Contains(sums, hash)
}

func hasher(file []byte, algo string) string {
	switch algo {
	case "md5":
		md5 := md5.Sum(file)
		return hex.EncodeToString(md5[:])
	case "sha1":
		sha1 := sha1.Sum(file)
		return hex.EncodeToString(sha1[:])
	case "sha256":
		sha256 := sha256.Sum256(file)
		return hex.EncodeToString(sha256[:])
	case "sha512":
		sha512 := sha512.Sum512(file)
		return hex.EncodeToString(sha512[:])
	default:
		return ""
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: hashma [file to hash] [file containg hashes]\n")
		return
	}

	algorithms := []string{
		"md5",
		"sha1",
		"sha256",
		"sha512",
	}

	file := os.Args[1]
	sums := os.Args[2]

	sumsBytes, err := ioutil.ReadFile(sums)
	if err != nil {
		fmt.Fprintf(os.Stderr, "hashma: %s\n", err)
		return
	}

	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "hashma: %s\n", err)
		return
	}

	hashchan := make(chan string)
	var wg sync.WaitGroup
	for _, algo := range algorithms {
		wg.Add(1)
		go func(algo string) {
			defer wg.Done()
			hashchan <- hasher(fileBytes, algo)
		}(algo)
	}

	go func() {
		wg.Wait()
		close(hashchan)
	}()

	for hash := range hashchan {
		if findHash(string(sumsBytes), hash) {
			fmt.Printf("%s  %s\n", hash, path.Base(file))
			return
		}
	}
}
