package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

// http://goinbigdata.com/golang-wait-for-all-goroutines-to-finish/
// https://codereview.stackexchange.com/questions/114376/bruteforce-md5-password-cracker

/*
var hashes = [15] string{
	"060453b490e5d87744c3703195df2f1a",
	"21ad598175add22e981d56073e4b0ffd",
	"6bbb51b3c4c56d20ed3b8a8629dae0a4",
	"423f92cba4341e7064f9906db9d56469",
	// "be2d9e79c322f7a3f2fe3dd6faba4fc3", // 31d3$
}
*/

// var hash = "e2fc714c4727ee9395f324cd2e7f331f" // abcd
// var hash = "cd088ce6eab814a28a558ed1906f1053" // !1q*h

func verifyHash(str [] byte, checkSum [16]byte) bool {
	// 0.60s  / L4
	// 25.76s / L5
	return md5.Sum(str) == checkSum
}

func combinationForLength(hash [16]byte, generatedString [] byte, maxLength int) {
	var alphabet = [] byte("abcdefghijklmnopqrstuvwxyz0123456789!@#$%&*")
	var n = len(alphabet)

	// if maxLength = 0, our generatedString is match our desired length
	if maxLength == 0 {

		// If the generated string match our current hash, set found to true
		// so we can exit

		if verifyHash(generatedString, hash) {
			// fmt.Printf("Password found: '%s'\n", generatedString)
		}

		return
	}

	// Generate a new branch with the current generated string + each letter from the alphabet
	// also decrease length because we added a char
	for i := 0; i < n; i++ {

		generatedString := append(generatedString, alphabet[i])

		combinationForLength(hash, generatedString, maxLength-1)

	}

}

func hack(rawHash string, length int) float64 {
	hash := [16]byte{}
	if _, err := hex.Decode(hash[:], bytes.TrimSpace([]byte(rawHash))); err != nil {
		return 0 // error
	}

	start := time.Now()

	combinationForLength(hash, []byte{}, length)

	elapsed := time.Since(start).Seconds()

	return elapsed
}

func main() {
	fmt.Printf("Running on %d cores\n\n", runtime.GOMAXPROCS(0))

	if len(os.Args) == 3 { // there is the correct amount of arguments
		length, _ := strconv.Atoi(os.Args[2])

		t := hack(os.Args[1], length) // main function

		fmt.Printf("Done for length = %d in %fs\n", length, t)

		input := bufio.NewScanner(os.Stdin)
		input.Scan()
	} else {
		fmt.Printf("Invalid usage: %s hash len", os.Args[0])
	}

}
