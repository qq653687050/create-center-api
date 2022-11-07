// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: api/create/v1/poll_record.proto

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

// PollRecordClient is the client API for PollRecord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PollRecordClient interface {
	CreatePollRecord(ctx context.Context, in *PollRecordDTO, opts ...grpc.CallOption) (*PollRecordVO, error)
	DeletePollRecord(ctx context.Context, in *PollRecordDTO, opts ...grpc.CallOption) (*PollRecordVO, error)
	GetPollRecord(ctx context.Context, in *PollRecordDTO, opts ...grpc.CallOption) (*PollRecordInfo, error)
	ListPollRecord(ctx context.Context, in *ListPollRecordDTO, opts ...grpc.CallOption) (*ListPollRecordVO, error)
}

type pollRecordClient struct {
	cc grpc.ClientConnInterface
}

func NewPollRecordClient(cc grpc.ClientConnInterface) PollRecordClient {
	return &pollRecordClient{cc}
}

func (c *pollRecordClient) CreatePollRecord(ctx context.Context, in *PollRecordDTO, opts ...grpc.CallOption) (*PollRecordVO, error) {
	out := new(PollRecordVO)
	err := c.cc.Invoke(ctx, "/api.create.v1.PollRecord/CreatePollRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pollRecordClient) DeletePollRecord(ctx context.Context, in *PollRecordDTO, opts ...grpc.CallOption) (*PollRecordVO, error) {
	out := new(PollRecordVO)
	err := c.cc.Invoke(ctx, "/api.create.v1.PollRecord/DeletePollRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pollRecordClient) GetPollRecord(ctx context.Context, in *PollRecordDTO, opts ...grpc.CallOption) (*PollRecordInfo, error) {
	out := new(PollRecordInfo)
	err := c.cc.Invoke(ctx, "/api.create.v1.PollRecord/GetPollRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pollRecordClient) ListPollRecord(ctx context.Context, in *ListPollRecordDTO, opts ...grpc.CallOption) (*ListPollRecordVO, error) {
	out := new(ListPollRecordVO)
	err := c.cc.Invoke(ctx, "/api.create.v1.PollRecord/ListPollRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PollRecordServer is the server API for PollRecord service.
// All implementations must embed UnimplementedPollRecordServer
// for forward compatibility
type PollRecordServer interface {
	CreatePollRecord(context.Context, *PollRecordDTO) (*PollRecordVO, error)
	DeletePollRecord(context.Context, *PollRecordDTO) (*PollRecordVO, error)
	GetPollRecord(context.Context, *PollRecordDTO) (*PollRecordInfo, error)
	ListPollRecord(context.Context, *ListPollRecordDTO) (*ListPollRecordVO, error)
	mustEmbedUnimplementedPollRecordServer()
}

// UnimplementedPollRecordServer must be embedded to have forward compatible implementations.
type UnimplementedPollRecordServer struct {
}

func (UnimplementedPollRecordServer) CreatePollRecord(context.Context, *PollRecordDTO) (*PollRecordVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePollRecord not implemented")
}
func (UnimplementedPollRecordServer) DeletePollRecord(context.Context, *PollRecordDTO) (*PollRecordVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePollRecord not implemented")
}
func (UnimplementedPollRecordServer) GetPollRecord(context.Context, *PollRecordDTO) (*PollRecordInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPollRecord not implemented")
}
func (UnimplementedPollRecordServer) ListPollRecord(context.Context, *ListPollRecordDTO) (*ListPollRecordVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPollRecord not implemented")
}
func (UnimplementedPollRecordServer) mustEmbedUnimplementedPollRecordServer() {}

// UnsafePollRecordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PollRecordServer will
// result in compilation errors.
type UnsafePollRecordServer interface {
	mustEmbedUnimplementedPollRecordServer()
}

func RegisterPollRecordServer(s grpc.ServiceRegistrar, srv PollRecordServer) {
	s.RegisterService(&PollRecord_ServiceDesc, srv)
}

func _PollRecord_CreatePollRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollRecordDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollRecordServer).CreatePollRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.PollRecord/CreatePollRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollRecordServer).CreatePollRecord(ctx, req.(*PollRecordDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _PollRecord_DeletePollRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollRecordDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollRecordServer).DeletePollRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.PollRecord/DeletePollRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollRecordServer).DeletePollRecord(ctx, req.(*PollRecordDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _PollRecord_GetPollRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollRecordDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollRecordServer).GetPollRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.PollRecord/GetPollRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollRecordServer).GetPollRecord(ctx, req.(*PollRecordDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _PollRecord_ListPollRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPollRecordDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollRecordServer).ListPollRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.create.v1.PollRecord/ListPollRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollRecordServer).ListPollRecord(ctx, req.(*ListPollRecordDTO))
	}
	return interceptor(ctx, in, info, handler)
}

// PollRecord_ServiceDesc is the grpc.ServiceDesc for PollRecord service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PollRecord_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.create.v1.PollRecord",
	HandlerType: (*PollRecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePollRecord",
			Handler:    _PollRecord_CreatePollRecord_Handler,
		},
		{
			MethodName: "DeletePollRecord",
			Handler:    _PollRecord_DeletePollRecord_Handler,
		},
		{
			MethodName: "GetPollRecord",
			Handler:    _PollRecord_GetPollRecord_Handler,
		},
		{
			MethodName: "ListPollRecord",
			Handler:    _PollRecord_ListPollRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/create/v1/poll_record.proto",
}
