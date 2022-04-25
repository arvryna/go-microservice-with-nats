// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/transervice.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TransactionManagerClient is the client API for TransactionManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionManagerClient interface {
	TransactionUp(ctx context.Context, in *NewTransaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	TransactionDown(ctx context.Context, in *NewTransaction, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type transactionManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionManagerClient(cc grpc.ClientConnInterface) TransactionManagerClient {
	return &transactionManagerClient{cc}
}

func (c *transactionManagerClient) TransactionUp(ctx context.Context, in *NewTransaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/transervice.TransactionManager/TransactionUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionManagerClient) TransactionDown(ctx context.Context, in *NewTransaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/transervice.TransactionManager/TransactionDown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionManagerServer is the server API for TransactionManager service.
// All implementations must embed UnimplementedTransactionManagerServer
// for forward compatibility
type TransactionManagerServer interface {
	TransactionUp(context.Context, *NewTransaction) (*TransactionResponse, error)
	TransactionDown(context.Context, *NewTransaction) (*TransactionResponse, error)
	mustEmbedUnimplementedTransactionManagerServer()
}

// UnimplementedTransactionManagerServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionManagerServer struct {
}

func (UnimplementedTransactionManagerServer) TransactionUp(context.Context, *NewTransaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactionUp not implemented")
}
func (UnimplementedTransactionManagerServer) TransactionDown(context.Context, *NewTransaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactionDown not implemented")
}
func (UnimplementedTransactionManagerServer) mustEmbedUnimplementedTransactionManagerServer() {}

// UnsafeTransactionManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionManagerServer will
// result in compilation errors.
type UnsafeTransactionManagerServer interface {
	mustEmbedUnimplementedTransactionManagerServer()
}

func RegisterTransactionManagerServer(s grpc.ServiceRegistrar, srv TransactionManagerServer) {
	s.RegisterService(&TransactionManager_ServiceDesc, srv)
}

func _TransactionManager_TransactionUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionManagerServer).TransactionUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transervice.TransactionManager/TransactionUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionManagerServer).TransactionUp(ctx, req.(*NewTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionManager_TransactionDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionManagerServer).TransactionDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transervice.TransactionManager/TransactionDown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionManagerServer).TransactionDown(ctx, req.(*NewTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionManager_ServiceDesc is the grpc.ServiceDesc for TransactionManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transervice.TransactionManager",
	HandlerType: (*TransactionManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransactionUp",
			Handler:    _TransactionManager_TransactionUp_Handler,
		},
		{
			MethodName: "TransactionDown",
			Handler:    _TransactionManager_TransactionDown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transervice.proto",
}