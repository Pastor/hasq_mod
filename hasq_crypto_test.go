package hashq_mod

import (
	"testing"
)

func (dk DeviceCrypto) PrintTest(t *testing.T) {
	t.Log("PUBLIC_KEY: ", dk.PublicKey)
	t.Log("PRIVATE_KEY: ", dk.PrivateKey)
}

func TestGenerateKey(t *testing.T) {
	key := GenerateKey()
	if key.Original == nil {
		t.Error("Internal key was null")
	}
	if len(key.PublicKey) == 0 {
		t.Error("Public key was null")
	}
	if len(key.PrivateKey) == 0 {
		t.Error("Private key was null")
	}
	key.PrintTest(t)
}

func TestDeviceCrypto_Sign(t *testing.T) {
	key := GenerateKey()
	origin := []byte("HELLO")
	sign := key.Sign(origin)
	t.Log("HELLO: ", sign)
	t.Log("DIGEST: ", EncodeToString(key.Digest(origin)))
	key.PrintTest(t)
}

func TestDeviceCrypto_Verify(t *testing.T) {
	key := GenerateKey()
	origin := []byte("HELLO")
	sign := key.Sign(origin)
	ret := key.Verify(nil, origin, DecodeFromString(sign))
	if ret != true {
		t.Error("Signature verify error")
	}
	t.Log("HELLO: ", sign)
	t.Log("DIGEST: ", EncodeToString(key.Digest(origin)))
	key.PrintTest(t)
}

func TestDeviceCrypto_Public(t *testing.T) {
	key1 := GenerateKey()
	key2 := GenerateKey()
	origin := []byte("HELLO")
	sign := key1.Sign(origin)
	ret := key1.Verify(nil, origin, DecodeFromString(sign))
	if ret != true {
		t.Error("Signature verify error")
	}
	t.Log("HELLO: ", sign)
	t.Log("DIGEST: ", EncodeToString(key1.Digest(origin)))
	key1.PrintTest(t)
	public := key1.Public()
	ret = key2.Verify(public, origin, DecodeFromString(sign))
	if ret != true {
		t.Error("Signature verify error")
	}
	t.Log("DIGEST: ", EncodeToString(key2.Digest(origin)))
	key2.PrintTest(t)
}

func TestHashStore_Add(t *testing.T) {
	store := NewStore()
	c1 := NewClient()
	tokenHash := c1.NewToken("TOKEN_CLIENT1")
	ch := c1.AddHash(tokenHash)
	if ch == nil {
		t.Fatal("Can't add hash to client")
	}
	check := store.Add(ch)
	if !check {
		t.Error("Can't add hash to store")
	}
	check = store.Add(c1.AddHash(tokenHash))
	if !check {
		t.Error("Can't add hash to store")
	}
}

func TestHashStore_Length(t *testing.T) {
	store := NewStore()
	c := NewClient()
	hash := c.NewToken("T")
	store.Add(c.AddHash(hash))
	store.Add(c.AddHash(hash))
	store.Add(c.AddHash(hash))
	if store.Length(hash) != 3 {
		t.Error("Store not contains 3 hashes")
	}
}

func TestHashStore_Validate(t *testing.T) {
	store := NewStore()
	c := NewClient()
	hash := c.NewToken("T")
	store.Add(c.AddHash(hash))
	store.Add(c.AddHash(hash))
	store.Add(c.AddHash(hash))
	if !store.Validate(hash) {
		t.Error("Validate hash sequence exception")
	}
}
