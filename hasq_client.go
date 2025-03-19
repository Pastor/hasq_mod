package hashq_mod

import (
	"container/list"
	"os"
	"strings"
)

type Client struct {
	Tokens map[string]*Token
}

func NewClient() Client {
	return Client{Tokens: make(map[string]*Token)}
}

func (c *Client) LoadTokens() bool {
	files, err := os.ReadDir(".")
	if err != nil {
		return false
	}
	for _, f := range files {
		name := f.Name()
		index := strings.Index(name, ".tok")
		if index <= -1 {
			continue
		}
		w := name[0:index]
		token := LoadToken(w)
		if token == nil {
			continue
		}
		c.Tokens[w] = token
	}
	return true
}

func (c *Client) StoreTokens() bool {
	for _, v := range c.Tokens {
		StoreToken(v)
	}
	return true
}

func (c *Client) NewToken(data string) string {
	hash := Hash(data)
	token := c.Tokens[hash]
	if token != nil {
		return token.Digest
	}
	newToken := NewToken(data)
	c.Tokens[hash] = &newToken
	return hash
}

func (c *Client) RegisterToken(hash string, key1 string, key2 string, gen string, owner string) {
	c.Tokens[hash] = &Token{
		List:    list.New(),
		Digest:  hash,
		Key1:    key1,
		Key2:    key2,
		LastGen: gen,
	}
}

func (c *Client) AddHash(hash string) *CanonicalHash {
	token := c.Tokens[hash]
	if token == nil {
		return nil
	}
	ch := token.Next()
	return &ch
}

func (c *Client) RegisterHash(sequence int32, token string, key string, gen string, owner string) *CanonicalHash {
	t := c.Tokens[token]
	if t == nil {
		return nil
	}
	return t.Add(sequence, key, gen, owner)
}
