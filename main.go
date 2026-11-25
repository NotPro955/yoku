package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	fmt.Println(args)

	short := sha256.Sum256([]byte(args[1]))
	fmt.Println(string(short[:]))
}
