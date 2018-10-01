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

func (c *Client) AddHash(hash string) *CanonicalHash {
	token := c.Tokens[hash]
	if token == nil {
		return nil
	}
	ch := token.Next()
	return &ch
}
