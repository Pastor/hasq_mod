package hashq_mod

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
	if sc.Connection != nil {
		_ = sc.Connection.Close()
	}
}

func (sc *SimpleClient) CreateHash(hash *CanonicalHash) bool {
	if sc.Connection == nil {
		return false
	}
	w := bufio.NewWriter(sc.Connection)
	_, _ = w.WriteString(hash.StringifyWithDigest() + "\n")
	_ = w.Flush()
	r := bufio.NewReader(sc.Connection)
	scanner := bufio.NewScanner(r)
	scanned := scanner.Scan()
	if !scanned {
		if err := scanner.Err(); err != nil {
			return false
		}
	}
	line := scanner.Text()
	return line == "true"
}
