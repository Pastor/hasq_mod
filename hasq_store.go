package hashq_mod

import (
	"bufio"
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type HashStore struct {
	IndexToken map[string]*list.List
}

func NewStore() HashStore {
	return HashStore{IndexToken: make(map[string]*list.List)}
}

func (store *HashStore) StoreAll() bool {
	for k := range store.IndexToken {
		store.StoreHash(k)
	}
	return true
}

func (store *HashStore) StoreHash(hash string) bool {
	hashes := store.IndexToken[hash]
	if hashes == nil {
		return false
	}
	file, err := os.Create(hash + ".hasq")
	if err != nil {
		log.Println(err)
		return false
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for temp := hashes.Front(); temp != nil; temp = temp.Next() {
		ch := temp.Value.(*CanonicalHash)
		_, _ = fmt.Fprintf(writer, "%s\n", ch.Stringify())
	}
	_ = writer.Flush()
	return true
}

func (store *HashStore) LoadAll() bool {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return false
	}
	for _, f := range files {
		name := f.Name()
		index := strings.Index(name, ".hasq")
		if index <= -1 {
			continue
		}
		w := string(name[0:index])
		store.LoadHash(w)
	}
	return true
}

func (store *HashStore) LoadHash(hash string) bool {
	hashes := store.IndexToken[hash]
	if hashes != nil {
		log.Println("Hash already exists")
		return false
	}
	LoadHash(hash, func(hash CanonicalHash) {
		store.Add(&hash)
	})
	return store.Validate(hash)
}

func (store *HashStore) Add(ch *CanonicalHash) bool {
	hashes := store.IndexToken[ch.Token]
	if hashes == nil {
		//if !ch.IsEmpty() || len(ch.Key) != 0 {
		//	return false
		//}
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
