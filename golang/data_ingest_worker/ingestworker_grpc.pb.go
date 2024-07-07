// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: data_ingest_worker/ingestworker.proto

package data_ingest_worker

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

const (
	Payment_GetData_FullMethodName = "/Payment/GetData"
)

// PaymentClient is the client API for Payment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentClient interface {
	GetData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetDataResponse, error)
}

type paymentClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentClient(cc grpc.ClientConnInterface) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) GetData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetDataResponse, error) {
	out := new(GetDataResponse)
	err := c.cc.Invoke(ctx, Payment_GetData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServer is the server API for Payment service.
// All implementations must embed UnimplementedPaymentServer
// for forward compatibility
type PaymentServer interface {
	GetData(context.Context, *Empty) (*GetDataResponse, error)
	mustEmbedUnimplementedPaymentServer()
}

// UnimplementedPaymentServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServer struct {
}

func (UnimplementedPaymentServer) GetData(context.Context, *Empty) (*GetDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedPaymentServer) mustEmbedUnimplementedPaymentServer() {}

// UnsafePaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServer will
// result in compilation errors.
type UnsafePaymentServer interface {
	mustEmbedUnimplementedPaymentServer()
}

func RegisterPaymentServer(s grpc.ServiceRegistrar, srv PaymentServer) {
	s.RegisterService(&Payment_ServiceDesc, srv)
}

func _Payment_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_GetData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).GetData(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Payment_ServiceDesc is the grpc.ServiceDesc for Payment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Payment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Payment",
	HandlerType: (*PaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _Payment_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data_ingest_worker/ingestworker.proto",
}