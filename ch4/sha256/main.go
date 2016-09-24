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
	"crypto/sha512"
	"hash"
)

//!+

func main() {
	var hash hash.Hash
	if len(os.Args) == 2 {
		if os.Args[1] == "sha384" {
			hash = sha512.New384()
		} else if os.Args[1] == "sha512" {
			hash = sha512.New()
		} else if os.Args[1] == "sha256" {
			hash = sha256.New()
		} else {
			log.Fatal(fmt.Sprintf("Unknown hash '%s'", os.Args[1]))
		}
	}

	if hash == nil {
		hash = sha256.New()
	}

	n, err := io.Copy(hash, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes hashed. Hash = \n%x\n", n, hash.Sum(nil))
}

//!-
