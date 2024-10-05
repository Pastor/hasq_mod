package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	mod "hashq_mod"
	"hashq_mod/grpc/gen"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	gen.UnimplementedHashServer

	store *mod.HashStore
}

func (s *server) Add(_ context.Context, in *gen.AddRequest) (*gen.AddReply, error) {
	hash := mod.CanonicalHash{
		Sequence: in.Sequence,
		Token:    in.Token,
		Key:      in.Key,
		Gen:      in.Gen,
		Owner:    in.Owner,
		Verified: false,
	}

	verified := s.store.Add(&hash)
	return &gen.AddReply{Verified: verified}, nil
}

func main() {
	store := mod.NewStore()
	store.LoadAll()
	defer func() {
		if err := recover(); err != nil {
			store.StoreAll()
		}
	}()
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gen.RegisterHashServer(s, &server{store: &store})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
