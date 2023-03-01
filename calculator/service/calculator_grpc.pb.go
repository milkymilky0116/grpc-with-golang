// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: calculator.proto

package proto

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

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorClient interface {
	Add(ctx context.Context, in *AddInput, opts ...grpc.CallOption) (*Result, error)
	Prime(ctx context.Context, in *PrimeInput, opts ...grpc.CallOption) (Calculator_PrimeClient, error)
	Average(ctx context.Context, opts ...grpc.CallOption) (Calculator_AverageClient, error)
	Max(ctx context.Context, opts ...grpc.CallOption) (Calculator_MaxClient, error)
	Sqrt(ctx context.Context, in *SqrtInput, opts ...grpc.CallOption) (*SqrtOutput, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *AddInput, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Prime(ctx context.Context, in *PrimeInput, opts ...grpc.CallOption) (Calculator_PrimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[0], "/calculator.Calculator/Prime", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorPrimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Calculator_PrimeClient interface {
	Recv() (*PrimeResult, error)
	grpc.ClientStream
}

type calculatorPrimeClient struct {
	grpc.ClientStream
}

func (x *calculatorPrimeClient) Recv() (*PrimeResult, error) {
	m := new(PrimeResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) Average(ctx context.Context, opts ...grpc.CallOption) (Calculator_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[1], "/calculator.Calculator/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorAverageClient{stream}
	return x, nil
}

type Calculator_AverageClient interface {
	Send(*AvgInput) error
	CloseAndRecv() (*AvgResult, error)
	grpc.ClientStream
}

type calculatorAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorAverageClient) Send(m *AvgInput) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorAverageClient) CloseAndRecv() (*AvgResult, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AvgResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) Max(ctx context.Context, opts ...grpc.CallOption) (Calculator_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[2], "/calculator.Calculator/Max", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorMaxClient{stream}
	return x, nil
}

type Calculator_MaxClient interface {
	Send(*MaxInput) error
	Recv() (*MaxResult, error)
	grpc.ClientStream
}

type calculatorMaxClient struct {
	grpc.ClientStream
}

func (x *calculatorMaxClient) Send(m *MaxInput) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorMaxClient) Recv() (*MaxResult, error) {
	m := new(MaxResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) Sqrt(ctx context.Context, in *SqrtInput, opts ...grpc.CallOption) (*SqrtOutput, error) {
	out := new(SqrtOutput)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Sqrt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
// All implementations must embed UnimplementedCalculatorServer
// for forward compatibility
type CalculatorServer interface {
	Add(context.Context, *AddInput) (*Result, error)
	Prime(*PrimeInput, Calculator_PrimeServer) error
	Average(Calculator_AverageServer) error
	Max(Calculator_MaxServer) error
	Sqrt(context.Context, *SqrtInput) (*SqrtOutput, error)
	mustEmbedUnimplementedCalculatorServer()
}

// UnimplementedCalculatorServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (UnimplementedCalculatorServer) Add(context.Context, *AddInput) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalculatorServer) Prime(*PrimeInput, Calculator_PrimeServer) error {
	return status.Errorf(codes.Unimplemented, "method Prime not implemented")
}
func (UnimplementedCalculatorServer) Average(Calculator_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (UnimplementedCalculatorServer) Max(Calculator_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}
func (UnimplementedCalculatorServer) Sqrt(context.Context, *SqrtInput) (*SqrtOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}
func (UnimplementedCalculatorServer) mustEmbedUnimplementedCalculatorServer() {}

// UnsafeCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServer will
// result in compilation errors.
type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {
	s.RegisterService(&Calculator_ServiceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*AddInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Prime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServer).Prime(m, &calculatorPrimeServer{stream})
}

type Calculator_PrimeServer interface {
	Send(*PrimeResult) error
	grpc.ServerStream
}

type calculatorPrimeServer struct {
	grpc.ServerStream
}

func (x *calculatorPrimeServer) Send(m *PrimeResult) error {
	return x.ServerStream.SendMsg(m)
}

func _Calculator_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).Average(&calculatorAverageServer{stream})
}

type Calculator_AverageServer interface {
	SendAndClose(*AvgResult) error
	Recv() (*AvgInput, error)
	grpc.ServerStream
}

type calculatorAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorAverageServer) SendAndClose(m *AvgResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorAverageServer) Recv() (*AvgInput, error) {
	m := new(AvgInput)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calculator_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).Max(&calculatorMaxServer{stream})
}

type Calculator_MaxServer interface {
	Send(*MaxResult) error
	Recv() (*MaxInput, error)
	grpc.ServerStream
}

type calculatorMaxServer struct {
	grpc.ServerStream
}

func (x *calculatorMaxServer) Send(m *MaxResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorMaxServer) Recv() (*MaxInput, error) {
	m := new(MaxInput)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calculator_Sqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Sqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Sqrt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Sqrt(ctx, req.(*SqrtInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Calculator_ServiceDesc is the grpc.ServiceDesc for Calculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculator_Add_Handler,
		},
		{
			MethodName: "Sqrt",
			Handler:    _Calculator_Sqrt_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Prime",
			Handler:       _Calculator_Prime_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _Calculator_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Max",
			Handler:       _Calculator_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
