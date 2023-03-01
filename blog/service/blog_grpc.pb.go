// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: blog.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BlogClient is the client API for Blog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogClient interface {
	CreateBlog(ctx context.Context, in *Post, opts ...grpc.CallOption) (*PostId, error)
	ReadBlog(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*Post, error)
	UpdateBlog(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Status, error)
	DeleteBlog(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*Status, error)
	ListBlog(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Blog_ListBlogClient, error)
}

type blogClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogClient(cc grpc.ClientConnInterface) BlogClient {
	return &blogClient{cc}
}

func (c *blogClient) CreateBlog(ctx context.Context, in *Post, opts ...grpc.CallOption) (*PostId, error) {
	out := new(PostId)
	err := c.cc.Invoke(ctx, "/blog.Blog/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) ReadBlog(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.Blog/ReadBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) UpdateBlog(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/blog.Blog/UpdateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) DeleteBlog(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/blog.Blog/DeleteBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) ListBlog(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Blog_ListBlogClient, error) {
	stream, err := c.cc.NewStream(ctx, &Blog_ServiceDesc.Streams[0], "/blog.Blog/ListBlog", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogListBlogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Blog_ListBlogClient interface {
	Recv() (*Post, error)
	grpc.ClientStream
}

type blogListBlogClient struct {
	grpc.ClientStream
}

func (x *blogListBlogClient) Recv() (*Post, error) {
	m := new(Post)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServer is the server API for Blog service.
// All implementations must embed UnimplementedBlogServer
// for forward compatibility
type BlogServer interface {
	CreateBlog(context.Context, *Post) (*PostId, error)
	ReadBlog(context.Context, *PostId) (*Post, error)
	UpdateBlog(context.Context, *Post) (*Status, error)
	DeleteBlog(context.Context, *PostId) (*Status, error)
	ListBlog(*emptypb.Empty, Blog_ListBlogServer) error
	mustEmbedUnimplementedBlogServer()
}

// UnimplementedBlogServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServer struct {
}

func (UnimplementedBlogServer) CreateBlog(context.Context, *Post) (*PostId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (UnimplementedBlogServer) ReadBlog(context.Context, *PostId) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlog not implemented")
}
func (UnimplementedBlogServer) UpdateBlog(context.Context, *Post) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (UnimplementedBlogServer) DeleteBlog(context.Context, *PostId) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
func (UnimplementedBlogServer) ListBlog(*emptypb.Empty, Blog_ListBlogServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBlog not implemented")
}
func (UnimplementedBlogServer) mustEmbedUnimplementedBlogServer() {}

// UnsafeBlogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServer will
// result in compilation errors.
type UnsafeBlogServer interface {
	mustEmbedUnimplementedBlogServer()
}

func RegisterBlogServer(s grpc.ServiceRegistrar, srv BlogServer) {
	s.RegisterService(&Blog_ServiceDesc, srv)
}

func _Blog_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).CreateBlog(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).ReadBlog(ctx, req.(*PostId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).UpdateBlog(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Blog/DeleteBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).DeleteBlog(ctx, req.(*PostId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_ListBlog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServer).ListBlog(m, &blogListBlogServer{stream})
}

type Blog_ListBlogServer interface {
	Send(*Post) error
	grpc.ServerStream
}

type blogListBlogServer struct {
	grpc.ServerStream
}

func (x *blogListBlogServer) Send(m *Post) error {
	return x.ServerStream.SendMsg(m)
}

// Blog_ServiceDesc is the grpc.ServiceDesc for Blog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.Blog",
	HandlerType: (*BlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _Blog_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _Blog_ReadBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _Blog_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _Blog_DeleteBlog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBlog",
			Handler:       _Blog_ListBlog_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blog.proto",
}