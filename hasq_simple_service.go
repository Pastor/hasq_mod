package hashq_mod

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct{}

func StartService(address string, store *HashStore) error {
	log.Printf("starting server on %v\n", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting connection %v", err)
			continue
		}
		log.Printf("accepted connection from %v", conn.RemoteAddr())
		HandleClient(store, conn) //TODO: Implement me
	}
}

func HandleClient(store *HashStore, conn net.Conn) error {
	defer func() {
		log.Printf("closing connection from %v", conn.RemoteAddr())
		conn.Close()
	}()
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	scanner := bufio.NewScanner(r)
	for {
		scanned := scanner.Scan()
		if !scanned {
			if err := scanner.Err(); err != nil {
				log.Printf("%v(%v)", err, conn.RemoteAddr())
				return err
			}
			continue
		}
		line := scanner.Text()
		log.Println("Receive: ", line)
		parts := strings.Fields(line)
		if len(parts) < 4 {
			log.Println("Error parse line \"", line, "\"")
			w.WriteString("false\n")
			w.Flush()
			continue
		}
		n, _ := strconv.ParseInt(parts[0], 10, 32)
		hash := CanonicalHash{
			Sequence: int32(n),
			Token:    parts[1],
			Key:      parts[2],
			Gen:      parts[3],
			Owner:    parts[4],
			Verified: false,
		}

		hash.Verified = store.Add(&hash)
		if hash.Verified {
			w.WriteString("true")
		} else {
			w.WriteString("false")
		}
		w.WriteString("\n")
		w.Flush()
	}
	return nil
}
