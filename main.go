package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/curve25519"
  "flag"
)

func main() {
        in:=flag.String("in","","in base64")
       flag.Parse() 
	//str := "sCecN+FH+pIPL4j2yheiN9n4IJbXg9gRUofTPGf/EXA="
        str:=*in
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	var src, dst [32]byte
	copy(src[:], data)
	curve25519.ScalarBaseMult(&dst, &src)
	str1 := base64.StdEncoding.EncodeToString(dst[:])
	fmt.Println(str1)

}
