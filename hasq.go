package main

import (
	"container/list"
	"crypto/md5"
	"fmt"
	"math/rand"
)

type CanonicalHash struct {
	Sequence int32
	Token    string
	Key      string
	Gen      string
	Owner    string

	//Help     string
	Verified bool
}

type Token struct {
	Data    string
	Digest  string
	Key1    string
	Key2    string
	LastGen string
	List    []CanonicalHash
}

func ValidateHash(ch CanonicalHash, ph CanonicalHash) bool {
	if ph.Token != ch.Token {
		return false
	}
	n := ch.Sequence
	key := ph.Key
	hash := Hash(n, ph.Token, key)
	if ch.Gen != hash {
		return false
	}
	return true
}

func Digest(params ...interface{}) []byte {
	var h string
	for _, p := range params {
		h += fmt.Sprint(p)
	}
	digest := md5.Sum([]byte(h))
	return digest[:]
}

func Hash(params ...interface{}) string {
	return EncodeToString(Digest(params))
}

func NextKey() string {
	r := make([]byte, 16)
	rand.Read(r)
	return EncodeToString(r)
}

func NewToken(data string) Token {
	return Token{
		Data:   data,
		Digest: Hash(data),
		Key1:   "",
		Key2:   NextKey(),
		List:   make([]CanonicalHash, 0),
	}
}

func (tok *Token) NextSequence() int32 {
	return int32(len(tok.List) + 1)
}

func (tok *Token) Next() CanonicalHash {
	var hash = CanonicalHash{}
	hash.Sequence = tok.NextSequence()
	hash.Token = tok.Digest
	hash.Key = tok.Key1
	hash.Gen = Hash(hash.Sequence, tok.Digest, tok.Key2)
	hash.Owner = Hash(hash.Sequence, tok.Digest, tok.LastGen)
	//hash.Help = fmt.Sprint(hash.Sequence) + "_" + tok.Digest + "_" + tok.Key2
	hash.Verified = false
	tok.Key1 = tok.Key2
	tok.Key2 = NextKey()
	tok.LastGen = hash.Gen

	tok.List = append(tok.List, hash)
	return hash
}

func (tok *Token) Validate() bool {
	var ph *CanonicalHash
	for i := len(tok.List) - 1; i >= 0; i-- {
		if ph == nil {
			ph = &tok.List[i]
			continue
		}
		var ch = &tok.List[i]
		ch.Verified = ValidateHash(*ch, *ph)
		if ch.Verified != true {
			return false
		}
		ph = ch
	}
	return true
}

func ValidateList(hashes *list.List) bool {
	var ph *CanonicalHash
	for temp := hashes.Back(); temp != nil; temp = temp.Prev() {
		if ph == nil {
			ph = temp.Value.(*CanonicalHash)
			continue
		}
		var ch = temp.Value.(*CanonicalHash)
		ch.Verified = ValidateHash(*ch, *ph)
		if ch.Verified != true {
			return false
		}
		ph = ch
	}
	return true
}

func (tok *Token) LastHash() *CanonicalHash {
	if len(tok.List) == 0 {
		return nil
	}
	return &tok.List[len(tok.List)-1]
}
