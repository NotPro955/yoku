package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	_ "os"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const short_len = 8

func short_gen() (string, error) {
	short := make([]byte, short_len)

	for i := range short {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		short[i] = chars[n.Int64()]
	}
	return string(short), nil
}

func main() {
	short, err := short_gen()
	if err != nil {
		fmt.Println(err)
	}

}
