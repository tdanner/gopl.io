// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"crypto/sha256"
)

//!+

func main() {
	hash := sha256.New()
	n, err := io.Copy(hash, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes hashed. Hash = \n%x\n", n, hash.Sum(nil))
}

//!-
