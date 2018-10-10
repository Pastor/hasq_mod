package main

import (
	"container/list"
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/sha3"
)

type HashSequence interface {
	Add(ch CanonicalHash) bool
	Validate() bool
	Length() int
}

type MemoryHashSequence struct {
	Hashes list.List
}

func NewMemoryHashSequence() HashSequence {
	return &MemoryHashSequence{}
}

func (m *MemoryHashSequence) Add(ch CanonicalHash) bool {
	lastHash := m.Hashes.Back()
	if lastHash == nil {
		m.Hashes.PushBack(&ch)
		return true
	}
	ph := lastHash.Value.(*CanonicalHash)
	if ph.Sequence+1 != ch.Sequence {
		return false
	}
	verified := ValidateHash(*ph, ch)
	if verified {
		m.Hashes.PushBack(&ch)
		return true
	}
	return false
}

func (m *MemoryHashSequence) Validate() bool {
	return ValidateList(&m.Hashes)
}

func (m *MemoryHashSequence) Length() int {
	return m.Hashes.Len()
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
	digest := sha3.Sum256([]byte(h))
	return digest[:]
}

func Hash(params ...interface{}) string {
	return EncodeToString(Digest(params))
}

func NextKey() string {
	r := make([]byte, 16)
	_, _ = rand.Read(r)
	return EncodeToString(r)
}
