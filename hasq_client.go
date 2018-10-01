package main

type Client struct {
	Tokens map[string]*Token
}

func NewClient() Client {
	return Client{Tokens: make(map[string]*Token)}
}

func (c *Client) NewToken(data string) string {
	hash := Hash(data)
	token := c.Tokens[hash]
	if token != nil {
		return ""
	}
	newToken := NewToken(data)
	c.Tokens[hash] = &newToken
	return hash
}

func (c *Client) RegisterToken(hash string, key1 string, key2 string, gen string, owner string) {
	c.Tokens[hash] = &Token{
		List:    make([]CanonicalHash, 0),
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
