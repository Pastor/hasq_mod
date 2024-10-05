package hashq_mod

import (
	"os"
	"testing"
)

func TestMemoryHashSequence_Add(t *testing.T) {
	sequence := NewMemoryHashSequence()
	tok := NewToken("DATA")
	sequence.Add(tok.Next())
	sequence.Add(tok.Next())
	sequence.Add(tok.Next())
	if sequence.Length() != 3 {
		t.Fatal("Sequence not have three elements")
	}
}

func TestServer_Create(t *testing.T) {
	address := "127.0.0.1:59090"
	sc := NewSimpleClient(address)
	defer sc.Close()
	c := NewClient()
	tokenHash := c.NewToken("SIMPLE_TOKEN")
	hash := c.AddHash(tokenHash)
	verified := sc.CreateHash(hash)
	if !verified {
		t.Error("Error create first hash")
	}
	hash = c.AddHash(tokenHash)
	verified = sc.CreateHash(hash)
	if !verified {
		t.Error("Error create second hash")
	}
}

func TestMain(m *testing.M) {
	store := NewStore()
	go StartService("127.0.0.1:59090", &store)
	os.Exit(m.Run())
}
