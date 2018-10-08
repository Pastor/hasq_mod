package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var service string
	var address string
	var mode string
	var data string
	var count int

	var clientKey string
	var clientGen string
	var clientToken string

	flag.StringVar(&service, "service_type", "simple", "Only [simple | mod] service type")
	flag.StringVar(&address, "address", "127.0.0.1:59090", "Bind address")
	flag.StringVar(&mode, "mode", "testing", "Only [client | service | testing | selftest]")
	flag.StringVar(&data, "data", "Simple_Token_Data", "Token data")
	flag.IntVar(&count, "count", 100, "Count")
	flag.StringVar(&clientKey, "c_key", "empty", "Client private key")
	flag.StringVar(&clientGen, "c_gen", "empty", "Client generation")
	flag.StringVar(&clientToken, "c_tok", "empty", "Client token")
	flag.Parse()

	store := NewStore()
	store.LoadAll()
	if mode == "selftest" {
		c := NewClient()
		c.LoadTokens()
		tokenHash := c.NewToken(data)
		for i := 0; i < count; i++ {
			hash := c.AddHash(tokenHash)
			verified := store.Add(hash)
			log.Println("Verified: ", verified)
		}
		c.StoreTokens()
	} else if mode == "testing" {
		go StartService(address, &store)
		sc := NewSimpleClient(address)
		defer sc.Close()
		c := NewClient()
		c.LoadTokens()
		tokenHash := c.NewToken(data)
		for i := 0; i < count; i++ {
			hash := c.AddHash(tokenHash)
			verified := sc.CreateHash(hash.Sequence, hash.Token, hash.Key, hash.Gen, hash.Owner)
			log.Println("Verified: ", verified)
		}
		c.StoreTokens()
	} else if mode == "client" {
		sc := NewSimpleClient(address)
		defer sc.Close()
		c := NewClient()
		c.LoadTokens()
		latestHash := sc.LatestHash(clientToken)
		if latestHash == nil {
			log.Fatal("Token not found")
		}
		if clientKey == "empty" {
			log.Fatal("Key not defined")
		}
		if clientGen == "empty" {
			log.Fatal("Gen not defined")
		}

		tokenHash := c.NewToken(data)
		hash := c.AddHash(tokenHash)
		verified := sc.CreateHash(hash.Sequence, hash.Token, hash.Key, hash.Gen, hash.Owner)
		log.Println("Verified: ", verified)
		c.StoreTokens()
	} else if mode == "service" {
		log.Println("Starting ", address, " ...")
		StartService(address, &store)
	} else {
		log.Fatal("Unknown mode ", mode)
		os.Exit(-1)
	}
	store.StoreAll()
	os.Exit(0)
}
