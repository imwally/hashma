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

func findHash(hash, digest string) bool {
	return strings.Contains(digest, hash)
}

func Hashes(file []byte) map[string]string {
	hashes := make(map[string]string)

	md5 := md5.Sum(file)
	hashes["md5"] = hex.EncodeToString(md5[:])

	sha1 := sha1.Sum(file)
	hashes["sha1"] = hex.EncodeToString(sha1[:])

	sha256 := sha256.Sum256(file)
	hashes["sha256"] = hex.EncodeToString(sha256[:])

	sha512 := sha512.Sum512(file)
	hashes["sha512"] = hex.EncodeToString(sha512[:])

	return hashes
}

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "hashma: expected two arguments\n")
		return
	}

	file := os.Args[1]
	digest := os.Args[2]

	fb, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := ioutil.ReadFile(digest)
	if err != nil {
		fmt.Println(err)
		return
	}

	hashes := Hashes(fb)

	for algo, hash := range hashes {
		if findHash(hash, string(db)) {
			fmt.Printf("Found %s: %s\n", algo, hash)
		}
	}
}
