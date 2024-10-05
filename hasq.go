package hashq_mod

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	Sequence int32
	Data     string
	Digest   string
	Key1     string
	Key2     string
	LastGen  string
	List     *list.List
}

func NewToken(data string) Token {
	return Token{
		Data:   data,
		Digest: Hash(data),
		Key1:   "",
		Key2:   NextKey(),
		List:   list.New(),
	}
}

func EmptyKey() string {
	return "00000000000000000000000000000000"
}

func (hash CanonicalHash) IsEmpty() bool {
	return hash.Key == EmptyKey()
}

func (tok *Token) NextSequence() int32 {
	sequence := int32(tok.List.Len() + 1)
	tok.Sequence = sequence
	return sequence
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
	tok.List.PushBack(&hash)
	return hash
}

func (tok *Token) Validate() bool {
	var ph *CanonicalHash
	for temp := tok.List.Back(); temp != nil; temp = temp.Prev() {
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

func StoreToken(tok *Token) bool {
	file, err := os.Create(tok.Digest + ".tok")
	if err != nil {
		log.Println(err)
		return false
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, _ = fmt.Fprintf(writer, "%05d %s %s %s %s\n", tok.Sequence, tok.Key1, tok.Key2, tok.LastGen, tok.Data)
	_ = writer.Flush()
	return true
}

func LoadToken(hash string) *Token {
	file, err := os.Open(hash + ".tok")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 4 {
			log.Println("Error parse line \"", line, "\"")
			continue
		}
		n, _ := strconv.ParseInt(parts[0], 10, 32)
		tok := Token{
			Sequence: int32(n),
			Digest:   hash,
			Key1:     parts[1],
			Key2:     parts[2],
			LastGen:  parts[3],
			Data:     parts[4],
			List:     list.New(),
		}
		if tok.Key1 == EmptyKey() {
			tok.Key1 = ""
		}
		LoadHash(hash, func(hash CanonicalHash) {
			tok.List.PushBack(&hash)
		})
		tok.Validate()
		return &tok
	}
	return nil
}

func (tok *Token) LastHash() *CanonicalHash {
	return tok.List.Back().Value.(*CanonicalHash)
}
func (tok *Token) Add(sequence int32, key string, gen string, owner string) *CanonicalHash {
	var hash = CanonicalHash{Sequence: sequence, Token: tok.Digest, Key: key, Gen: gen, Owner: owner, Verified: false}
	tok.List.PushBack(&hash)
	return &hash
}

func LoadHash(hash string, appender func(CanonicalHash)) {
	file, err := os.Open(hash + ".hasq")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 4 {
			log.Println("Error parse line \"", line, "\"")
			continue
		}
		n, _ := strconv.ParseInt(parts[0], 10, 32)
		hash := CanonicalHash{
			Sequence: int32(n),
			Token:    hash,
			Key:      parts[1],
			Gen:      parts[2],
			Owner:    parts[3],
			Verified: false,
		}
		if hash.IsEmpty() {
			hash.Key = ""
		}
		appender(hash)
	}
}
