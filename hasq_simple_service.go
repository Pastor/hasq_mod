package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Store *HashStore
}

func (s *Server) Create(ctx context.Context, in *HasqHash) (*HasqResult, error) {
	hash := CanonicalHash{Sequence: in.Sequence, Token: in.Token, Key: in.Key, Gen: in.Gen, Owner: in.Owner}
	verified := s.Store.Add(&hash)
	return &HasqResult{Verified: verified}, nil
}

func (s *Server) Latest(ctx context.Context, request *HasqRequest) (*HasqHash, error) {
	hashes := s.Store.IndexToken[request.Id]
	if hashes != nil && hashes.Back() != nil {
		back := hashes.Back().Value.(*CanonicalHash)
		return &HasqHash{Key: back.Key, Gen: back.Gen, Token: back.Token, Owner: back.Owner, Sequence: back.Sequence}, nil
	}
	return nil, nil
}

func StartService(address string, store *HashStore) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterHashServiceServer(s, &Server{store})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
