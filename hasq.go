package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)

type CanonicalHash struct {
	Number int32
	Token  string
	Key    string
	Gen    string
	Owner  string

	//Help     string
	Verified bool
}

func (hash CanonicalHash) Print() {
	key := hash.Key
	if len(key) == 0 {
		key = "00000000000000000000000000000000"
	}
	verified := "NotVerified"
	if hash.Verified {
		verified = "Verified"
	}
	fmt.Println(
		fmt.Sprintf("%05d", hash.Number),
		" ", hash.Token,
		" ", key,
		" ", hash.Gen,
		" ", hash.Owner,
		" ", verified)
}

type Token struct {
	Data    string
	Digest  string
	Key1    string
	Key2    string
	LastGen string
	List    []CanonicalHash
}

func Hash(params ...interface{}) string {
	var h string
	for _, p := range params {
		h += fmt.Sprint(p)
	}
	digest := md5.Sum([]byte(h))
	return EncodeToString(digest[:])
}

func NextKey() string {
	r := make([]byte, 16)
	rand.Read(r)
	return EncodeToString(r)
}

func NewToken(data string) Token {
	digest := md5.Sum([]byte(data))

	return Token{
		Data:   data,
		Digest: EncodeToString(digest[:]),
		Key1:   "",
		Key2:   NextKey(),
		List:   make([]CanonicalHash, 0),
	}
}

func (tok *Token) Next() CanonicalHash {
	var hash = CanonicalHash{}
	hash.Number = int32(len(tok.List) + 1)
	hash.Token = tok.Digest
	hash.Key = tok.Key1
	hash.Gen = Hash(hash.Number, tok.Digest, tok.Key2)
	hash.Owner = Hash(hash.Number, tok.Digest, tok.LastGen)
	//hash.Help = fmt.Sprint(hash.Number) + "_" + tok.Digest + "_" + tok.Key2
	hash.Verified = false
	tok.Key1 = tok.Key2
	tok.Key2 = NextKey()
	tok.LastGen = hash.Gen

	tok.List = append(tok.List, hash)
	return hash
}
func (tok *Token) Print() {
	for i := len(tok.List) - 1; i >= 0; i-- {
		tok.List[i].Print()
	}
}
func (tok *Token) Validate() bool {
	var ph *CanonicalHash
	for i := len(tok.List) - 1; i >= 0; i-- {
		if ph == nil {
			ph = &tok.List[i]
			continue
		}
		var ch = &tok.List[i]
		n := ch.Number
		key := ph.Key
		digest := tok.Digest
		hash := Hash(n, digest, key)
		if ch.Gen != hash {
			return false
		}
		ch.Verified = true
		ph = ch
	}
	return true
}
