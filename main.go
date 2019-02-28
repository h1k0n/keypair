package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/curve25519"
)

func main() {

	var b [32]byte
	_, err := rand.Read(b[:])
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(PrintPriKey(b))
	var dst [32]byte
	curve25519.ScalarBaseMult(&dst, &b)
	str1 := base64.StdEncoding.EncodeToString(dst[:])
	fmt.Println(str1)

}

func PrintPriKey(e [32]byte) string {
	e[0] &= 248
	e[31] &= 127
	e[31] |= 64
	str1 := base64.StdEncoding.EncodeToString(e[:])
	return str1
}
