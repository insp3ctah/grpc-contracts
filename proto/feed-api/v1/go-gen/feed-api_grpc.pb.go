// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: feed-api.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Feed_Like_FullMethodName      = "/feed_api.v1.Feed/Like"
	Feed_SuperLike_FullMethodName = "/feed_api.v1.Feed/SuperLike"
	Feed_Dislike_FullMethodName   = "/feed_api.v1.Feed/Dislike"
	Feed_Save_FullMethodName      = "/feed_api.v1.Feed/Save"
)

// FeedClient is the client API for Feed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedClient interface {
	Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error)
	SuperLike(ctx context.Context, in *SuperLikeRequest, opts ...grpc.CallOption) (*SuperLikeResponse, error)
	Dislike(ctx context.Context, in *DislikeRequest, opts ...grpc.CallOption) (*DislikeResponse, error)
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error)
}

type feedClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedClient(cc grpc.ClientConnInterface) FeedClient {
	return &feedClient{cc}
}

func (c *feedClient) Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LikeResponse)
	err := c.cc.Invoke(ctx, Feed_Like_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) SuperLike(ctx context.Context, in *SuperLikeRequest, opts ...grpc.CallOption) (*SuperLikeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SuperLikeResponse)
	err := c.cc.Invoke(ctx, Feed_SuperLike_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) Dislike(ctx context.Context, in *DislikeRequest, opts ...grpc.CallOption) (*DislikeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DislikeResponse)
	err := c.cc.Invoke(ctx, Feed_Dislike_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, Feed_Save_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServer is the server API for Feed service.
// All implementations must embed UnimplementedFeedServer
// for forward compatibility
type FeedServer interface {
	Like(context.Context, *LikeRequest) (*LikeResponse, error)
	SuperLike(context.Context, *SuperLikeRequest) (*SuperLikeResponse, error)
	Dislike(context.Context, *DislikeRequest) (*DislikeResponse, error)
	Save(context.Context, *SaveRequest) (*SaveResponse, error)
	mustEmbedUnimplementedFeedServer()
}

// UnimplementedFeedServer must be embedded to have forward compatible implementations.
type UnimplementedFeedServer struct {
}

func (UnimplementedFeedServer) Like(context.Context, *LikeRequest) (*LikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedFeedServer) SuperLike(context.Context, *SuperLikeRequest) (*SuperLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuperLike not implemented")
}
func (UnimplementedFeedServer) Dislike(context.Context, *DislikeRequest) (*DislikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dislike not implemented")
}
func (UnimplementedFeedServer) Save(context.Context, *SaveRequest) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedFeedServer) mustEmbedUnimplementedFeedServer() {}

// UnsafeFeedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServer will
// result in compilation errors.
type UnsafeFeedServer interface {
	mustEmbedUnimplementedFeedServer()
}

func RegisterFeedServer(s grpc.ServiceRegistrar, srv FeedServer) {
	s.RegisterService(&Feed_ServiceDesc, srv)
}

func _Feed_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_Like_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).Like(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_SuperLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuperLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).SuperLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_SuperLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).SuperLike(ctx, req.(*SuperLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_Dislike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DislikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).Dislike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_Dislike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).Dislike(ctx, req.(*DislikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feed_Save_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Feed_ServiceDesc is the grpc.ServiceDesc for Feed service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Feed_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feed_api.v1.Feed",
	HandlerType: (*FeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Like",
			Handler:    _Feed_Like_Handler,
		},
		{
			MethodName: "SuperLike",
			Handler:    _Feed_SuperLike_Handler,
		},
		{
			MethodName: "Dislike",
			Handler:    _Feed_Dislike_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _Feed_Save_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feed-api.proto",
}