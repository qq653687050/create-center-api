// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: api/create/v1/words.proto

package v1

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

// WordsClient is the client API for Words service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WordsClient interface {
	CreateWords(ctx context.Context, in *CreateWordsRequest, opts ...grpc.CallOption) (*CreateWordsReply, error)
	UpdateWordsPoll(ctx context.Context, in *UpdateWordsPollDTO, opts ...grpc.CallOption) (*UpdateWordsPollVO, error)
	DeleteWords(ctx context.Context, in *DeleteWordsRequest, opts ...grpc.CallOption) (*DeleteWordsReply, error)
	GetWords(ctx context.Context, in *GetWordsDTO, opts ...grpc.CallOption) (*GetWordsVO, error)
	ListWords(ctx context.Context, in *ListWordsDTO, opts ...grpc.CallOption) (*ListWordsVO, error)
	ListWordsByIds(ctx context.Context, in *ListWordsByIdsDTO, opts ...grpc.CallOption) (*ListWordsVO, error)
}

type wordsClient struct {
	cc grpc.ClientConnInterface
}

func NewWordsClient(cc grpc.ClientConnInterface) WordsClient {
	return &wordsClient{cc}
}

func (c *wordsClient) CreateWords(ctx context.Context, in *CreateWordsRequest, opts ...grpc.CallOption) (*CreateWordsReply, error) {
	out := new(CreateWordsReply)
	err := c.cc.Invoke(ctx, "/api.create.v1.Words/CreateWords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordsClient) UpdateWordsPoll(ctx context.Context, in *UpdateWordsPollDTO, opts ...grpc.CallOption) (*UpdateWordsPollVO, error) {
	out := new(UpdateWordsPollVO)
	err := c.cc.Invoke(ctx, "/api.create.v1.Words/UpdateWordsPoll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordsClient) DeleteWords(ctx context.Context, in *DeleteWordsRequest, opts ...grpc.CallOption) (*DeleteWordsReply, error) {
	out := new(DeleteWordsReply)
	err := c.cc.Invoke(ctx, "/api.create.v1.Words/DeleteWords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordsClient) GetWords(ctx context.Context, in *GetWordsDTO, opts ...grpc.CallOption) (*GetWordsVO, error) {
	out := new(GetWordsVO)
	err := c.cc.Invoke(ctx, "/api.create.v1.Words/GetWords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordsClient) ListWords(ctx context.Context, in *ListWordsDTO, opts ...grpc.CallOption) (*ListWordsVO, error) {
	out := new(ListWordsVO)
	err := c.cc.Invoke(ctx, "/api.create.v1.Words/ListWords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordsClient) ListWordsByIds(ctx context.Context, in *ListWordsByIdsDTO, opts ...grpc.CallOption) (*ListWordsVO, error) {
	out := new(ListWordsVO)
	err := c.cc.Invoke(ctx, "/api.create.v1.Words/ListWordsByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WordsServer is the server API for Words service.
// All implementations must embed UnimplementedWordsServer
// for forward compatibility
type WordsServer interface {
	CreateWords(context.Context, *CreateWordsRequest) (*CreateWordsReply, error)
	UpdateWordsPoll(context.Context, *UpdateWordsPollDTO) (*UpdateWordsPollVO, error)
	DeleteWords(context.Context, *DeleteWordsRequest) (*DeleteWordsReply, error)
	GetWords(context.Context, *GetWordsDTO) (*GetWordsVO, error)
	ListWords(context.Context, *ListWordsDTO) (*ListWordsVO, error)
	ListWordsByIds(context.Context, *ListWordsByIdsDTO) (*ListWordsVO, error)
	mustEmbedUnimplementedWordsServer()
}

// UnimplementedWordsServer must be embedded to have forward compatible implementations.
type UnimplementedWordsServer struct {
}

func (UnimplementedWordsServer) CreateWords(context.Context, *CreateWordsRequest) (*CreateWordsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWords not implemented")
}
func (UnimplementedWordsServer) UpdateWordsPoll(context.Context, *UpdateWordsPollDTO) (*UpdateWordsPollVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWordsPoll not implemented")
}
func (UnimplementedWordsServer) DeleteWords(context.Context, *DeleteWordsRequest) (*DeleteWordsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWords not implemented")
}
func (UnimplementedWordsServer) GetWords(context.Context, *GetWordsDTO) (*GetWordsVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWords not implemented")
}
func (UnimplementedWordsServer) ListWords(context.Context, *ListWordsDTO) (*ListWordsVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWords not implemented")
}
func (UnimplementedWordsServer) ListWordsByIds(context.Context, *ListWordsByIdsDTO) (*ListWordsVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWordsByIds not implemented")
}
func (UnimplementedWordsServer) mustEmbedUnimplementedWordsServer() {}

// UnsafeWordsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WordsServer will
// result in compilation errors.
type UnsafeWordsServer interface {
	mustEmbedUnimplementedWordsServer()
}

func RegisterWordsServer(s grpc.ServiceRegistrar, srv WordsServer) {
	s.RegisterService(&Words_ServiceDesc, srv)
}

func _Words_CreateWords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordsServer).CreateWords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.Words/CreateWords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordsServer).CreateWords(ctx, req.(*CreateWordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Words_UpdateWordsPoll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWordsPollDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordsServer).UpdateWordsPoll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.Words/UpdateWordsPoll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordsServer).UpdateWordsPoll(ctx, req.(*UpdateWordsPollDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _Words_DeleteWords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordsServer).DeleteWords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.Words/DeleteWords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordsServer).DeleteWords(ctx, req.(*DeleteWordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Words_GetWords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWordsDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordsServer).GetWords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.Words/GetWords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordsServer).GetWords(ctx, req.(*GetWordsDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _Words_ListWords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWordsDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordsServer).ListWords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.Words/ListWords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordsServer).ListWords(ctx, req.(*ListWordsDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _Words_ListWordsByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWordsByIdsDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordsServer).ListWordsByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.Words/ListWordsByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordsServer).ListWordsByIds(ctx, req.(*ListWordsByIdsDTO))
	}
	return interceptor(ctx, in, info, handler)
}

// Words_ServiceDesc is the grpc.ServiceDesc for Words service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Words_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.create.v1.Words",
	HandlerType: (*WordsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWords",
			Handler:    _Words_CreateWords_Handler,
		},
		{
			MethodName: "UpdateWordsPoll",
			Handler:    _Words_UpdateWordsPoll_Handler,
		},
		{
			MethodName: "DeleteWords",
			Handler:    _Words_DeleteWords_Handler,
		},
		{
			MethodName: "GetWords",
			Handler:    _Words_GetWords_Handler,
		},
		{
			MethodName: "ListWords",
			Handler:    _Words_ListWords_Handler,
		},
		{
			MethodName: "ListWordsByIds",
			Handler:    _Words_ListWordsByIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/create/v1/words.proto",
}
