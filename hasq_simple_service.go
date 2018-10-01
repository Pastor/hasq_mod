package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	Store HashStore
}

func (s *Server) Create(ctx context.Context, in *HasqCreate) (*HasqResult, error) {
	hash := CanonicalHash{Sequence: in.Sequence, Token: in.Token, Key: in.Key, Gen: in.Gen, Owner: in.Owner}
	verified := s.Store.Add(&hash)
	return &HasqResult{Verified: verified}, nil
}

func StartService(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	store := NewStore()
	RegisterHashServiceServer(s, &Server{store})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
