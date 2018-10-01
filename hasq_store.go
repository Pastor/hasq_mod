package main

import (
	"container/list"
)

type HashStore struct {
	IndexToken map[string]*list.List
}

func NewStore() HashStore {
	return HashStore{IndexToken: make(map[string]*list.List)}
}

func (store *HashStore) Add(ch *CanonicalHash) bool {
	hashes := store.IndexToken[ch.Token]
	if hashes == nil {
		if len(ch.Key) != 0 {
			return false
		}
		store.IndexToken[ch.Token] = list.New()
		store.IndexToken[ch.Token].PushBack(ch)
		return true
	}
	lastHash := hashes.Back()
	if lastHash == nil {
		return false
	}
	ph := lastHash.Value.(*CanonicalHash)
	if ph.Sequence+1 != ch.Sequence {
		return false
	}
	verified := ValidateHash(*ph, *ch)
	if verified {
		hashes.PushBack(ch)
		return true
	}
	return false
}

func (store *HashStore) Validate(hash string) bool {
	hashes := store.IndexToken[hash]
	if hashes != nil {
		return ValidateList(hashes)
	}
	return false
}
func (store *HashStore) Length(hash string) int {
	hashes := store.IndexToken[hash]
	if hashes != nil {
		return hashes.Len()
	}
	return 0
}
