package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

type SimpleClient struct {
	Connection *grpc.ClientConn
	Client     HashServiceClient
	Context    context.Context
	Cancel     context.CancelFunc
}

func NewSimpleClient(address string) SimpleClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := NewHashServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	return SimpleClient{Connection: conn, Client: c, Context: ctx, Cancel: cancel}
}

func (sc *SimpleClient) Close() {
	_ = sc.Connection.Close()
	sc.Cancel()
}

func (sc *SimpleClient) CreateHash(sequence int32, token string, key string, gen string, owner string) bool {
	r, err := sc.Client.Create(sc.Context,
		&HasqHash{Sequence: sequence, Token: token, Key: key, Gen: gen, Owner: owner})
	if err != nil {
		log.Fatalf("could not call: %v", err)
	}
	return r.GetVerified()
}

func (sc *SimpleClient) LatestHash(token string) *CanonicalHash {
	r, err := sc.Client.Latest(sc.Context, &HasqRequest{Id: token})
	if err != nil {
		log.Fatalf("could not call: %v", err)
	}
	if r == nil {
		return nil
	}
	return &CanonicalHash{
		Sequence: r.Sequence,
		Token:    r.Token,
		Key:      r.Key,
		Gen:      r.Gen,
		Owner:    r.Owner,
	}
}
