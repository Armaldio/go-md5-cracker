package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

// http://goinbigdata.com/golang-wait-for-all-goroutines-to-finish/
// https://codereview.stackexchange.com/questions/114376/bruteforce-md5-password-cracker

/*
var hashes = [15] string{
	// found "58047859b0e1218acd754f569baf9e33",
	// found "94bf87e03cd7dd9f4b826b6f200b98f4",
	// found "aae81cc29985fe2462ffee9a63371a70",
	// found "6bc8d7c479ed8ebac94c763766a8f514",
	// found "99ae3a8efc9bf7fd17bc947706644c91",
	// found "f2246fbd2e2e3f93c3c50922bd16cbbd",
	// found "9735f6cc8bce4a82d77ea74b8fe2f994",
	// found "1efa33adb7f6a92e69a3b6cd3bf532ab",
	// found "17c58fad14ecb9953c652b6517ee2022",
	// found "c8af88b1d7a7b3fbe39f3c6de35364ca",
	"060453b490e5d87744c3703195df2f1a",
	"21ad598175add22e981d56073e4b0ffd",
	"6bbb51b3c4c56d20ed3b8a8629dae0a4",
	"423f92cba4341e7064f9906db9d56469",
	// "be2d9e79c322f7a3f2fe3dd6faba4fc3", // 31d3$
}
*/

/*var hashes = [15] string{
	"e2fc714c4727ee9395f324cd2e7f331f", // abcd
}*/

var hash = "423f92cba4341e7064f9906db9d56469"
// var hash = "e2fc714c4727ee9395f324cd2e7f331f" // abcd
var checkSum = bytes.TrimSpace([]byte(hash))
var hashFound = false

var alphabet = [] string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "!", "@", "#", "$", "%", "&", "*"}
// alphabet := [] string{"a", "b", "c"}
// a
// b
// c
// aa
// ab
// ac
// ba
// bb
// bc
// ca
// cb
// cc
var n = len(alphabet)


func verifyHash(str string) bool {
	calculatedHash := md5.Sum([]byte(str))

	// 0.93s  / L4
	// 40.80s / L5
	dst := [16]byte{}
	if _, err := hex.Decode(dst[:], checkSum); err != nil {
		return false
	} else {
		return calculatedHash == dst
	}
}

func combinationsForLength(k int) {
	if hashFound {
		return
	}

	start := time.Now()

	combinationForLengthRec(alphabet, "", k)

	elapsed := time.Since(start).Seconds()

	fmt.Printf("Done for length = %d in %fs\n", k, elapsed)
}

// The main recursive method
// to print all possible
// strings of length k
func combinationForLengthRec(set [] string, prefix string, k int) {
	if hashFound {
		return
	}

	// Base case: k is 0,
	// print prefix
	if k == 0 {
		if verifyHash(prefix) {
			fmt.Printf("\n\tPassword found for '%s': '%s'\n\n", hash, prefix)
			hashFound = true
		}
		return
	}

	// One by one add all characters
	// from set and recursively
	// call for k equals to k-1
	for i := 0; i < n; i++ {
		// Next character of input added
		newPrefix := prefix + set[i]

		// k is decreased, because
		// we have added a new character
		combinationForLengthRec(set, newPrefix, k-1)
	}
}

func main() {

	// var wg sync.WaitGroup

	// max string = 8 chars
	for i := 1; i <= 3; i++ {
		// wg.Add(1)

		/*go func(alphabet []string, i int) {*/
		combinationsForLength(i)
		// wg.Done()
		/*}(alphabet, i)*/
	}
	/*wg.Wait()*/
}
