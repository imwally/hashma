package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "hashma: expected two arguments\n")
		return
	}

	file := os.Args[1]
	hashFile := os.Args[2]

	fb, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	hb, err := ioutil.ReadFile(hashFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	md5 := md5.Sum(fb)
	sha1 := sha1.Sum(fb)
	sha256 := sha256.Sum256(fb)
	sha512 := sha512.Sum512(fb)

	fmt.Println(string(hb))

	fmt.Printf("%x\n", md5)
	fmt.Printf("%x\n", sha1)
	fmt.Printf("%x\n", sha256)
	fmt.Printf("%x\n", sha512)
}
