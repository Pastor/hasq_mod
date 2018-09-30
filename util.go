package main

import (
	"encoding/hex"
	"strings"
)

func EncodeToString(data []byte) string {
	return strings.ToUpper(hex.EncodeToString(data))
}

func DecodeFromString(data string) []byte {
	ret, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}
	return ret
}
