package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func CreateHash(address string, hash CanonicalHash) bool {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewHashServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Create(ctx, &HasqCreate{Sequence: hash.Sequence, Token: hash.Token, Key: hash.Key, Gen: hash.Gen, Owner: hash.Owner})
	if err != nil {
		log.Fatalf("could not call: %v", err)
	}
	return r.GetVerified()
}
