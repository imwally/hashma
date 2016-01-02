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
	"strings"
)

var (
	algorithms = []string{
		"md5",
		"sha1",
		"sha256",
		"sha512",
	}

	hashchan = make(chan map[string]string)
	hashes   = make(map[string]string)
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
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "hashma: expected two arguments\n")
		return
	}

	file := os.Args[1]
	sums := os.Args[2]

	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("hashma:", err)
		return
	}

	sumsBytes, err := ioutil.ReadFile(sums)
	if err != nil {
		fmt.Println("hashma:", err)
		return
	}

	for _, algo := range algorithms {
		go func(algo string) {
			hashes[algo] = hasher(fileBytes, algo)
			hashchan <- hashes
		}(algo)
	}

	for {
		for algo, hash := range <-hashchan {
			if findHash(string(sumsBytes), hash) {
				fmt.Printf("%s: %s\n", algo, hash)
				close(hashchan)
				return
			}
		}
	}
}
