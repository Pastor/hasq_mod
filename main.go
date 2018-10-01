package main

import (
	"fmt"
	"os"
)

func main() {
	go StartService("127.0.0.1:59090")
	c := NewClient()
	tokenHash := c.NewToken("SIMPLE_TOKEN")
	hash := c.AddHash(tokenHash)
	address := "127.0.0.1:59090"
	verified := CreateHash(address, *hash)
	fmt.Println("Verified: ", verified)
	hash = c.AddHash(tokenHash)
	verified = CreateHash(address, *hash)
	fmt.Println("Verified: ", verified)
	os.Exit(0)
}
