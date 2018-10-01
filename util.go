package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func EncodeToString(data []byte) string {
	return strings.ToUpper(hex.EncodeToString(data))
}

func DecodeFromString(data string) []byte {
	ret, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}
	return ret
}

func (dk DeviceCrypto) Print(f *os.File) {
	f.WriteString("PublicKey : " + dk.PublicKey)
	f.Write([]byte{0x0A, 0x0D})
	f.WriteString("PrivateKey: " + dk.PrivateKey)
	f.Write([]byte{0x0A, 0x0D})
	f.Sync()
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
		fmt.Sprintf("%05d", hash.Sequence),
		" ", hash.Token,
		" ", key,
		" ", hash.Gen,
		" ", hash.Owner,
		" ", verified)
}

func (tok *Token) Print() {
	for i := len(tok.List) - 1; i >= 0; i-- {
		tok.List[i].Print()
	}
}

func (store *HashStore) Print(hash string) {
	fmt.Println("Token: ", hash)
	hashes := store.IndexToken[hash]
	if hashes != nil {
		for temp := hashes.Back(); temp != nil; temp = temp.Prev() {
			temp.Value.(*CanonicalHash).Print()
		}
	}
}
