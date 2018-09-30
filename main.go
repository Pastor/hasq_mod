package main

func main() {
	token := NewToken("TOKEN1")
	token.Next()
	token.Next()
	token.Next()
	token.Next()
	token.Next()
	token.Next()
	token.Next()
	token.Next()
	token.Validate()
	token.Print()
}
