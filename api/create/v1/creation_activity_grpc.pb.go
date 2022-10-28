// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: api/create/v1/creation_activity.proto

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

// CreationActivityClient is the client API for CreationActivity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreationActivityClient interface {
	CreateCreationActivity(ctx context.Context, in *CreateCreationActivityDTO, opts ...grpc.CallOption) (*CreateCreationActivityVO, error)
	UpdateCreationActivity(ctx context.Context, in *UpdateCreationActivityDTO, opts ...grpc.CallOption) (*UpdateCreationActivityVO, error)
	DeleteCreationActivity(ctx context.Context, in *DeleteCreationActivityRequest, opts ...grpc.CallOption) (*DeleteCreationActivityReply, error)
	GetCreationActivity(ctx context.Context, in *GetCreationActivityDTO, opts ...grpc.CallOption) (*GetCreationActivityVO, error)
	ListCreationActivity(ctx context.Context, in *ListCreationActivityDTO, opts ...grpc.CallOption) (*ListCreationActivityVO, error)
}

type creationActivityClient struct {
	cc grpc.ClientConnInterface
}

func NewCreationActivityClient(cc grpc.ClientConnInterface) CreationActivityClient {
	return &creationActivityClient{cc}
}

func (c *creationActivityClient) CreateCreationActivity(ctx context.Context, in *CreateCreationActivityDTO, opts ...grpc.CallOption) (*CreateCreationActivityVO, error) {
	out := new(CreateCreationActivityVO)
	err := c.cc.Invoke(ctx, "/api.create.v1.CreationActivity/CreateCreationActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationActivityClient) UpdateCreationActivity(ctx context.Context, in *UpdateCreationActivityDTO, opts ...grpc.CallOption) (*UpdateCreationActivityVO, error) {
	out := new(UpdateCreationActivityVO)
	err := c.cc.Invoke(ctx, "/api.create.v1.CreationActivity/UpdateCreationActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationActivityClient) DeleteCreationActivity(ctx context.Context, in *DeleteCreationActivityRequest, opts ...grpc.CallOption) (*DeleteCreationActivityReply, error) {
	out := new(DeleteCreationActivityReply)
	err := c.cc.Invoke(ctx, "/api.create.v1.CreationActivity/DeleteCreationActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationActivityClient) GetCreationActivity(ctx context.Context, in *GetCreationActivityDTO, opts ...grpc.CallOption) (*GetCreationActivityVO, error) {
	out := new(GetCreationActivityVO)
	err := c.cc.Invoke(ctx, "/api.create.v1.CreationActivity/GetCreationActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationActivityClient) ListCreationActivity(ctx context.Context, in *ListCreationActivityDTO, opts ...grpc.CallOption) (*ListCreationActivityVO, error) {
	out := new(ListCreationActivityVO)
	err := c.cc.Invoke(ctx, "/api.create.v1.CreationActivity/ListCreationActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreationActivityServer is the server API for CreationActivity service.
// All implementations must embed UnimplementedCreationActivityServer
// for forward compatibility
type CreationActivityServer interface {
	CreateCreationActivity(context.Context, *CreateCreationActivityDTO) (*CreateCreationActivityVO, error)
	UpdateCreationActivity(context.Context, *UpdateCreationActivityDTO) (*UpdateCreationActivityVO, error)
	DeleteCreationActivity(context.Context, *DeleteCreationActivityRequest) (*DeleteCreationActivityReply, error)
	GetCreationActivity(context.Context, *GetCreationActivityDTO) (*GetCreationActivityVO, error)
	ListCreationActivity(context.Context, *ListCreationActivityDTO) (*ListCreationActivityVO, error)
	mustEmbedUnimplementedCreationActivityServer()
}

// UnimplementedCreationActivityServer must be embedded to have forward compatible implementations.
type UnimplementedCreationActivityServer struct {
}

func (UnimplementedCreationActivityServer) CreateCreationActivity(context.Context, *CreateCreationActivityDTO) (*CreateCreationActivityVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCreationActivity not implemented")
}
func (UnimplementedCreationActivityServer) UpdateCreationActivity(context.Context, *UpdateCreationActivityDTO) (*UpdateCreationActivityVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCreationActivity not implemented")
}
func (UnimplementedCreationActivityServer) DeleteCreationActivity(context.Context, *DeleteCreationActivityRequest) (*DeleteCreationActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCreationActivity not implemented")
}
func (UnimplementedCreationActivityServer) GetCreationActivity(context.Context, *GetCreationActivityDTO) (*GetCreationActivityVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreationActivity not implemented")
}
func (UnimplementedCreationActivityServer) ListCreationActivity(context.Context, *ListCreationActivityDTO) (*ListCreationActivityVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCreationActivity not implemented")
}
func (UnimplementedCreationActivityServer) mustEmbedUnimplementedCreationActivityServer() {}

// UnsafeCreationActivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreationActivityServer will
// result in compilation errors.
type UnsafeCreationActivityServer interface {
	mustEmbedUnimplementedCreationActivityServer()
}

func RegisterCreationActivityServer(s grpc.ServiceRegistrar, srv CreationActivityServer) {
	s.RegisterService(&CreationActivity_ServiceDesc, srv)
}

func _CreationActivity_CreateCreationActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCreationActivityDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationActivityServer).CreateCreationActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.CreationActivity/CreateCreationActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationActivityServer).CreateCreationActivity(ctx, req.(*CreateCreationActivityDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationActivity_UpdateCreationActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCreationActivityDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationActivityServer).UpdateCreationActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.CreationActivity/UpdateCreationActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationActivityServer).UpdateCreationActivity(ctx, req.(*UpdateCreationActivityDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationActivity_DeleteCreationActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCreationActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationActivityServer).DeleteCreationActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.CreationActivity/DeleteCreationActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationActivityServer).DeleteCreationActivity(ctx, req.(*DeleteCreationActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationActivity_GetCreationActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCreationActivityDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationActivityServer).GetCreationActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.CreationActivity/GetCreationActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationActivityServer).GetCreationActivity(ctx, req.(*GetCreationActivityDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreationActivity_ListCreationActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCreationActivityDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationActivityServer).ListCreationActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.CreationActivity/ListCreationActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationActivityServer).ListCreationActivity(ctx, req.(*ListCreationActivityDTO))
	}
	return interceptor(ctx, in, info, handler)
}

// CreationActivity_ServiceDesc is the grpc.ServiceDesc for CreationActivity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreationActivity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.create.v1.CreationActivity",
	HandlerType: (*CreationActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCreationActivity",
			Handler:    _CreationActivity_CreateCreationActivity_Handler,
		},
		{
			MethodName: "UpdateCreationActivity",
			Handler:    _CreationActivity_UpdateCreationActivity_Handler,
		},
		{
			MethodName: "DeleteCreationActivity",
			Handler:    _CreationActivity_DeleteCreationActivity_Handler,
		},
		{
			MethodName: "GetCreationActivity",
			Handler:    _CreationActivity_GetCreationActivity_Handler,
		},
		{
			MethodName: "ListCreationActivity",
			Handler:    _CreationActivity_ListCreationActivity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/create/v1/creation_activity.proto",
}
