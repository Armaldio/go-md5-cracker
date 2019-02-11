package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	hashes := [15] string{
		"58047859b0e1218acd754f569baf9e33", // dilatat
		"94bf87e03cd7dd9f4b826b6f200b98f4", // gateront
		"aae81cc29985fe2462ffee9a63371a70", // poutsais
		"6bc8d7c479ed8ebac94c763766a8f514", // strippas
		"99ae3a8efc9bf7fd17bc947706644c91", // abricots
		"f2246fbd2e2e3f93c3c50922bd16cbbd", // percets
		"9735f6cc8bce4a82d77ea74b8fe2f994", // carias
		"1efa33adb7f6a92e69a3b6cd3bf532ab", // suspens
		"17c58fad14ecb9953c652b6517ee2022", // orpheons
		"c8af88b1d7a7b3fbe39f3c6de35364ca", // sursoies
		"060453b490e5d87744c3703195df2f1a",
		"21ad598175add22e981d56073e4b0ffd",
		"6bbb51b3c4c56d20ed3b8a8629dae0a4",
		"423f92cba4341e7064f9906db9d56469",
		"be2d9e79c322f7a3f2fe3dd6faba4fc3",
	}

	// Making the path relative to the execution
	absPath, _ := filepath.Abs("bin/mots-8-et-moins.txt")

	// Opening the file
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// starting counter
	start := time.Now()

	// for each line in the file
	for scanner.Scan() {
		str := scanner.Text()
		for _, hash := range hashes {
			calculatedHash := md5.Sum([]byte(str))
			checkSum := bytes.TrimSpace([]byte(hash))

			dst := [16]byte{}
			if _, err := hex.Decode(dst[:], checkSum); err != nil {

			} else {
				if calculatedHash == dst {
					fmt.Printf("Password found for hash '%s': '%s'\n", hash, str)
				}
			}

		}
	}

	elapsed := time.Since(start).Seconds()

	fmt.Printf("\nTook %fs", elapsed)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
