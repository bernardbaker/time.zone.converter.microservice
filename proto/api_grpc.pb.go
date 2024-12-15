// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: api.proto

package proto

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
	TimeZoneConverter_ConvertTime_FullMethodName = "/timezone.TimeZoneConverter/ConvertTime"
)

// TimeZoneConverterClient is the client API for TimeZoneConverter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeZoneConverterClient interface {
	ConvertTime(ctx context.Context, in *ConvertTimeRequest, opts ...grpc.CallOption) (*ConvertTimeResponse, error)
}

type timeZoneConverterClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeZoneConverterClient(cc grpc.ClientConnInterface) TimeZoneConverterClient {
	return &timeZoneConverterClient{cc}
}

func (c *timeZoneConverterClient) ConvertTime(ctx context.Context, in *ConvertTimeRequest, opts ...grpc.CallOption) (*ConvertTimeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConvertTimeResponse)
	err := c.cc.Invoke(ctx, TimeZoneConverter_ConvertTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeZoneConverterServer is the server API for TimeZoneConverter service.
// All implementations must embed UnimplementedTimeZoneConverterServer
// for forward compatibility.
type TimeZoneConverterServer interface {
	ConvertTime(context.Context, *ConvertTimeRequest) (*ConvertTimeResponse, error)
	mustEmbedUnimplementedTimeZoneConverterServer()
}

// UnimplementedTimeZoneConverterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTimeZoneConverterServer struct{}

func (UnimplementedTimeZoneConverterServer) ConvertTime(context.Context, *ConvertTimeRequest) (*ConvertTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertTime not implemented")
}
func (UnimplementedTimeZoneConverterServer) mustEmbedUnimplementedTimeZoneConverterServer() {}
func (UnimplementedTimeZoneConverterServer) testEmbeddedByValue()                           {}

// UnsafeTimeZoneConverterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeZoneConverterServer will
// result in compilation errors.
type UnsafeTimeZoneConverterServer interface {
	mustEmbedUnimplementedTimeZoneConverterServer()
}

func RegisterTimeZoneConverterServer(s grpc.ServiceRegistrar, srv TimeZoneConverterServer) {
	// If the following call pancis, it indicates UnimplementedTimeZoneConverterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TimeZoneConverter_ServiceDesc, srv)
}

func _TimeZoneConverter_ConvertTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeZoneConverterServer).ConvertTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeZoneConverter_ConvertTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeZoneConverterServer).ConvertTime(ctx, req.(*ConvertTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeZoneConverter_ServiceDesc is the grpc.ServiceDesc for TimeZoneConverter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeZoneConverter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "timezone.TimeZoneConverter",
	HandlerType: (*TimeZoneConverterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertTime",
			Handler:    _TimeZoneConverter_ConvertTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}