// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.0
// source: hashq.proto

package gen

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Hash_Add_FullMethodName = "/hashq_grpc.Hash/Add"
)

// HashClient is the client API for Hash service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HashClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
}

type hashClient struct {
	cc grpc.ClientConnInterface
}

func NewHashClient(cc grpc.ClientConnInterface) HashClient {
	return &hashClient{cc}
}

func (c *hashClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddReply)
	err := c.cc.Invoke(ctx, Hash_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HashServer is the server API for Hash service.
// All implementations must embed UnimplementedHashServer
// for forward compatibility.
type HashServer interface {
	Add(context.Context, *AddRequest) (*AddReply, error)
	mustEmbedUnimplementedHashServer()
}

// UnimplementedHashServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHashServer struct{}

func (UnimplementedHashServer) Add(context.Context, *AddRequest) (*AddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedHashServer) mustEmbedUnimplementedHashServer() {}
func (UnimplementedHashServer) testEmbeddedByValue()              {}

// UnsafeHashServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HashServer will
// result in compilation errors.
type UnsafeHashServer interface {
	mustEmbedUnimplementedHashServer()
}

func RegisterHashServer(s grpc.ServiceRegistrar, srv HashServer) {
	// If the following call pancis, it indicates UnimplementedHashServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Hash_ServiceDesc, srv)
}

func _Hash_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hash_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hash_ServiceDesc is the grpc.ServiceDesc for Hash service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hash_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hashq_grpc.Hash",
	HandlerType: (*HashServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Hash_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hashq.proto",
}
