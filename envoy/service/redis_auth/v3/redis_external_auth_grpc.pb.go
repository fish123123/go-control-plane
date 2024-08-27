// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: envoy/service/redis_auth/v3/redis_external_auth.proto

package redis_authv3

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
	RedisProxyExternalAuth_Authenticate_FullMethodName = "/envoy.service.redis_auth.v3.RedisProxyExternalAuth/Authenticate"
)

// RedisProxyExternalAuthClient is the client API for RedisProxyExternalAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RedisProxyExternalAuthClient interface {
	// Performs authentication check based on the data sent with the AUTH request.
	// Returns either an OK status or an error status.
	Authenticate(ctx context.Context, in *RedisProxyExternalAuthRequest, opts ...grpc.CallOption) (*RedisProxyExternalAuthResponse, error)
}

type redisProxyExternalAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewRedisProxyExternalAuthClient(cc grpc.ClientConnInterface) RedisProxyExternalAuthClient {
	return &redisProxyExternalAuthClient{cc}
}

func (c *redisProxyExternalAuthClient) Authenticate(ctx context.Context, in *RedisProxyExternalAuthRequest, opts ...grpc.CallOption) (*RedisProxyExternalAuthResponse, error) {
	out := new(RedisProxyExternalAuthResponse)
	err := c.cc.Invoke(ctx, RedisProxyExternalAuth_Authenticate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedisProxyExternalAuthServer is the server API for RedisProxyExternalAuth service.
// All implementations should embed UnimplementedRedisProxyExternalAuthServer
// for forward compatibility
type RedisProxyExternalAuthServer interface {
	// Performs authentication check based on the data sent with the AUTH request.
	// Returns either an OK status or an error status.
	Authenticate(context.Context, *RedisProxyExternalAuthRequest) (*RedisProxyExternalAuthResponse, error)
}

// UnimplementedRedisProxyExternalAuthServer should be embedded to have forward compatible implementations.
type UnimplementedRedisProxyExternalAuthServer struct {
}

func (UnimplementedRedisProxyExternalAuthServer) Authenticate(context.Context, *RedisProxyExternalAuthRequest) (*RedisProxyExternalAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

// UnsafeRedisProxyExternalAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedisProxyExternalAuthServer will
// result in compilation errors.
type UnsafeRedisProxyExternalAuthServer interface {
	mustEmbedUnimplementedRedisProxyExternalAuthServer()
}

func RegisterRedisProxyExternalAuthServer(s grpc.ServiceRegistrar, srv RedisProxyExternalAuthServer) {
	s.RegisterService(&RedisProxyExternalAuth_ServiceDesc, srv)
}

func _RedisProxyExternalAuth_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedisProxyExternalAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisProxyExternalAuthServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisProxyExternalAuth_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisProxyExternalAuthServer).Authenticate(ctx, req.(*RedisProxyExternalAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RedisProxyExternalAuth_ServiceDesc is the grpc.ServiceDesc for RedisProxyExternalAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedisProxyExternalAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.redis_auth.v3.RedisProxyExternalAuth",
	HandlerType: (*RedisProxyExternalAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _RedisProxyExternalAuth_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/redis_auth/v3/redis_external_auth.proto",
}
