package main

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	ss := "23230"
	code64 := base64.StdEncoding.EncodeToString([]byte(ss))
	data64, _ := base64.StdEncoding.DecodeString(code64)
	fmt.Println("base64", code64, "data64", string(data64))

	code32 := base32.StdEncoding.EncodeToString([]byte(ss))
	data32, _ := base32.StdEncoding.DecodeString(code32)
	fmt.Println("base32", code32, "data32", string(data32))

	code16 := hex.EncodeToString([]byte(ss))
	data16, _ := hex.DecodeString(code16)
	fmt.Println("base16", code16, "data16", string(data16))
}
