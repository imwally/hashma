package main

import (
	"testing"
	"fmt"
	"sync"
	"io/ioutil"
)

var (
	algorithms = []string{
		"md5",
		"sha1",
		"sha256",
		"sha512",
	}
	
	data = []byte{
		0x54, 0x68, 0x65, 0x20, 0x73, 0x6B, 0x79, 0x20,
		0x61, 0x62, 0x6F, 0x76, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70,
		0x6F, 0x72, 0x74, 0x20, 0x77, 0x61, 0x73, 0x20, 0x74, 0x68, 0x65,
		0x20, 0x63, 0x6F, 0x6C, 0x6F, 0x72, 0x20, 0x6F, 0x66, 0x20, 0x74,
		0x65, 0x6C, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6F, 0x6E, 0x2C, 0x20,
		0x74, 0x75, 0x6E, 0x65, 0x64, 0x20, 0x74, 0x6F, 0x20, 0x61, 0x20,
		0x64, 0x65, 0x61, 0x64, 0x20, 0x63, 0x68, 0x61, 0x6E, 0x6E, 0x65,
		0x6C, 0x2E,
	}

	MD5hash = "c736b509feee0cb9a4f8d606d5f6050d"
	SHA1hash = "49ceb59827e0c76fcf94d5da80b8f4bc0dc94b15"
	SHA256hash = "5041821981ec48d8db280ff293c35de17ef5dbfbac25adc81ff272d0fc22b2ae"
	SHA512hash = "53164b83c1e27d6f0f4175fb3f0a8c668ad7afe8e9f423ed8fe5311c80556e7be33449369947bb0c29bd6702c90ab08a004a767802d309fb04726f2a799ee1b3"

	// Load a lager external file to test benchmarks.
	file = "debian-live-8.2.0-amd64-gnome-desktop.iso"
)

func TestMD5Hasher(t *testing.T) {
	md5 := hasher(data, "md5")
	if md5 != MD5hash {
		t.Error(`MD5 failed.`)
	}
}

func TestSHA1Hasher(t *testing.T) {
	sha1 := hasher(data, "sha1")
	if sha1 != SHA1hash {
		t.Error(`SHA1 failed.`)
	}
}

func TestSHA256Hasher(t *testing.T) {
	sha256 := hasher(data, "sha256")
	if sha256 != SHA256hash {
		t.Error(`SHA256 failed.`)
	}
}

func TestSHA512Hasher(t *testing.T) {
	sha512 := hasher(data, "sha512")
	if sha512 != SHA512hash {
		t.Error(`SHA512 failed.`)
	}
}

func TestFindHashesWithNoCheckSum(t *testing.T) {
	hashchan := make(chan string)

	var wg sync.WaitGroup
	for _, algo := range algorithms {
		wg.Add(1)
		go func(algo string) {
			defer wg.Done()
			hashchan <- hasher(data, algo)
		}(algo)
	}

	go func() {
		wg.Wait()
		close(hashchan)
	}()
	
	for hash := range hashchan {
		if findHash("test", hash) {
			fmt.Printf("%s", hash)
		}
	}
}


func BenchmarkSequentialHasher(t *testing.B) {
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	fmt.Println()
	for _, algo := range algorithms {
		hash := hasher(fileBytes, algo)
		fmt.Printf("%s\n", hash)
	}
}
		
func BenchmarkConcurrentHasher(t *testing.B) {
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		t.Error(err)
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

	fmt.Println()
	for hash := range hashchan {
		fmt.Printf("%s\n", hash)
	}
}
