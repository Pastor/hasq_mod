package main

import (
	"bufio"
	"net"
)

type SimpleClient struct {
	Connection net.Conn
}

func NewSimpleClient(address string) SimpleClient {
	conn, _ := net.Dial("tcp", address)
	return SimpleClient{Connection: conn}
}

func (sc *SimpleClient) Close() {
	_ = sc.Connection.Close()
}

func (sc *SimpleClient) CreateHash(hash *CanonicalHash) bool {
	w := bufio.NewWriter(sc.Connection)
	w.WriteString(hash.StringifyWithDigest() + "\n")
	w.Flush()
	r := bufio.NewReader(sc.Connection)
	scanr := bufio.NewScanner(r)
	scanned := scanr.Scan()
	if !scanned {
		if err := scanr.Err(); err != nil {
			return false
		}
	}
	line := scanr.Text()
	return line == "true"
}
