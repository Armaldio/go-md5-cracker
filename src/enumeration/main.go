package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"time"
)

func combinationForLengthRec(alphabet []byte, hash [16]byte, generatedString [] byte, maxLength int, c chan string) {
	// if maxLength = 0, our generatedString is match our desired length
	if maxLength == 0 {

		// If the generated string match our current hash, set found to true
		// so we can exit

		if md5.Sum(generatedString) == hash {
			c <- string(generatedString)
		}
		return
	}

	// ----
	// Generate a new branch with the current generated string + each letter from the alphabet
	// also decrease length because we added a char
	for i := range alphabet {
		combinationForLengthRec(alphabet, hash, append(generatedString, alphabet[i]), maxLength-1, c)
	}
}

func hack(rawHash string, length int) (float64, bool) {
	hash := [16]byte{}
	if _, err := hex.Decode(hash[:], bytes.TrimSpace([]byte(rawHash))); err != nil {
		return 0, false // error
	}

	start := time.Now()

	// var alphabet = [] byte("abcdefghijklmnopqrstuvwxyz0123456789!@#$%&*")
	var alphabet = [] byte("e3a4@is$5nrt7o0ludcmpgbvhfqyxjkwz12689!#%&*")

	resultChan := make(chan string, length)

	for _, char := range alphabet {
		char := char
		go func() {
			combinationForLengthRec(alphabet, hash, []byte{char}, length-1, resultChan)
			resultChan <- ""
		}()
	}

	found := false
	password := ""
	for i := 0; i < len(alphabet); i++ {
		val := <-resultChan
		if val != "" {
			password = val
			found = true
			fmt.Printf("Password found: --> %s <--\n", password)
			break
		}

	}

	elapsed := time.Since(start).Seconds()

	return elapsed, found
}

func main() {
	fmt.Println("Running on", runtime.GOMAXPROCS(0), "cores")

	if len(os.Args) == 2 { // there is the correct amount of arguments

		for i := 1; i < 8; i++ {
			t, found := hack(os.Args[1], i) // main function
			fmt.Println("Done for length =", i, "in", t, "s")

			if found {
				break
			}
		}
	} else {
		fmt.Println("Invalid usage:\n\n\t", os.Args[0], "hash")
	}

}
