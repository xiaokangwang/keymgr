package def

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
	TPMKeyService_GetPrimaryKeyHash_FullMethodName = "/org.kkdev.keymgr.def.TPMKeyService/GetPrimaryKeyHash"
	TPMKeyService_CreateECDSAKey_FullMethodName    = "/org.kkdev.keymgr.def.TPMKeyService/CreateECDSAKey"
	TPMKeyService_SignWithECDSAKey_FullMethodName  = "/org.kkdev.keymgr.def.TPMKeyService/SignWithECDSAKey"
)

// TPMKeyServiceClient is the client API for TPMKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TPMKeyServiceClient interface {
	GetPrimaryKeyHash(ctx context.Context, in *GetPrimaryKeyHashReq, opts ...grpc.CallOption) (*GetPrimaryKeyHashResp, error)
	CreateECDSAKey(ctx context.Context, in *CreateECDSAKeyReq, opts ...grpc.CallOption) (*CreateECDSAKeyResp, error)
	SignWithECDSAKey(ctx context.Context, in *SignWithECDSAKeyReq, opts ...grpc.CallOption) (*SignWithECDSAKeyResp, error)
}

type tPMKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTPMKeyServiceClient(cc grpc.ClientConnInterface) TPMKeyServiceClient {
	return &tPMKeyServiceClient{cc}
}

func (c *tPMKeyServiceClient) GetPrimaryKeyHash(ctx context.Context, in *GetPrimaryKeyHashReq, opts ...grpc.CallOption) (*GetPrimaryKeyHashResp, error) {
	out := new(GetPrimaryKeyHashResp)
	err := c.cc.Invoke(ctx, TPMKeyService_GetPrimaryKeyHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tPMKeyServiceClient) CreateECDSAKey(ctx context.Context, in *CreateECDSAKeyReq, opts ...grpc.CallOption) (*CreateECDSAKeyResp, error) {
	out := new(CreateECDSAKeyResp)
	err := c.cc.Invoke(ctx, TPMKeyService_CreateECDSAKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tPMKeyServiceClient) SignWithECDSAKey(ctx context.Context, in *SignWithECDSAKeyReq, opts ...grpc.CallOption) (*SignWithECDSAKeyResp, error) {
	out := new(SignWithECDSAKeyResp)
	err := c.cc.Invoke(ctx, TPMKeyService_SignWithECDSAKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TPMKeyServiceServer is the server API for TPMKeyService service.
// All implementations must embed UnimplementedTPMKeyServiceServer
// for forward compatibility
type TPMKeyServiceServer interface {
	GetPrimaryKeyHash(context.Context, *GetPrimaryKeyHashReq) (*GetPrimaryKeyHashResp, error)
	CreateECDSAKey(context.Context, *CreateECDSAKeyReq) (*CreateECDSAKeyResp, error)
	SignWithECDSAKey(context.Context, *SignWithECDSAKeyReq) (*SignWithECDSAKeyResp, error)
	mustEmbedUnimplementedTPMKeyServiceServer()
}

// UnimplementedTPMKeyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTPMKeyServiceServer struct {
}

func (UnimplementedTPMKeyServiceServer) GetPrimaryKeyHash(context.Context, *GetPrimaryKeyHashReq) (*GetPrimaryKeyHashResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrimaryKeyHash not implemented")
}
func (UnimplementedTPMKeyServiceServer) CreateECDSAKey(context.Context, *CreateECDSAKeyReq) (*CreateECDSAKeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateECDSAKey not implemented")
}
func (UnimplementedTPMKeyServiceServer) SignWithECDSAKey(context.Context, *SignWithECDSAKeyReq) (*SignWithECDSAKeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignWithECDSAKey not implemented")
}
func (UnimplementedTPMKeyServiceServer) mustEmbedUnimplementedTPMKeyServiceServer() {}

// UnsafeTPMKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TPMKeyServiceServer will
// result in compilation errors.
type UnsafeTPMKeyServiceServer interface {
	mustEmbedUnimplementedTPMKeyServiceServer()
}

func RegisterTPMKeyServiceServer(s grpc.ServiceRegistrar, srv TPMKeyServiceServer) {
	s.RegisterService(&TPMKeyService_ServiceDesc, srv)
}

func _TPMKeyService_GetPrimaryKeyHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrimaryKeyHashReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TPMKeyServiceServer).GetPrimaryKeyHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TPMKeyService_GetPrimaryKeyHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TPMKeyServiceServer).GetPrimaryKeyHash(ctx, req.(*GetPrimaryKeyHashReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TPMKeyService_CreateECDSAKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateECDSAKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TPMKeyServiceServer).CreateECDSAKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TPMKeyService_CreateECDSAKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TPMKeyServiceServer).CreateECDSAKey(ctx, req.(*CreateECDSAKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TPMKeyService_SignWithECDSAKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignWithECDSAKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TPMKeyServiceServer).SignWithECDSAKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TPMKeyService_SignWithECDSAKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TPMKeyServiceServer).SignWithECDSAKey(ctx, req.(*SignWithECDSAKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TPMKeyService_ServiceDesc is the grpc.ServiceDesc for TPMKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TPMKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.kkdev.keymgr.def.TPMKeyService",
	HandlerType: (*TPMKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrimaryKeyHash",
			Handler:    _TPMKeyService_GetPrimaryKeyHash_Handler,
		},
		{
			MethodName: "CreateECDSAKey",
			Handler:    _TPMKeyService_CreateECDSAKey_Handler,
		},
		{
			MethodName: "SignWithECDSAKey",
			Handler:    _TPMKeyService_SignWithECDSAKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "def/rpc.proto",
}
