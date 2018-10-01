package main

import (
	"os"
	"testing"
)

func TestServer_Create(t *testing.T) {
	c := NewClient()
	tokenHash := c.NewToken("SIMPLE_TOKEN")
	hash := c.AddHash(tokenHash)
	address := "127.0.0.1:59090"
	verified := CreateHash(address, *hash)
	if !verified {
		t.Error("Error create first hash")
	}
	hash = c.AddHash(tokenHash)
	verified = CreateHash(address, *hash)
	if !verified {
		t.Error("Error create second hash")
	}
}

func TestMain(m *testing.M) {
	go StartService("127.0.0.1:59090")
	os.Exit(m.Run())
}
