package main

func main() {
	store := NewStore()
	c1 := NewClient()
	c2 := NewClient()

	token1 := c1.NewToken("TOKEN_CLIENT1")
	token2 := c2.NewToken("TOKEN_CLIENT2")
	hash := c1.AddHash(token1)
	store.Add(hash)
	hash = c1.AddHash(token1)
	store.Add(hash)
	hash = c2.AddHash(token2)
	store.Add(hash)

	store.Validate(token1)
	store.Validate(token2)
	store.Print(token1)
	store.Print(token2)
}
