// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: sso/customer/customer_office.proto

package customer_office

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

// OfficeServiceClient is the client API for OfficeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OfficeServiceClient interface {
	CreateOffice(ctx context.Context, in *CreateOfficeRequest, opts ...grpc.CallOption) (*CreateOfficeResponse, error)
	GetOfficeList(ctx context.Context, in *GetOfficeListRequest, opts ...grpc.CallOption) (*GetOfficeListResponse, error)
}

type officeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOfficeServiceClient(cc grpc.ClientConnInterface) OfficeServiceClient {
	return &officeServiceClient{cc}
}

func (c *officeServiceClient) CreateOffice(ctx context.Context, in *CreateOfficeRequest, opts ...grpc.CallOption) (*CreateOfficeResponse, error) {
	out := new(CreateOfficeResponse)
	err := c.cc.Invoke(ctx, "/customer.OfficeService/CreateOffice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *officeServiceClient) GetOfficeList(ctx context.Context, in *GetOfficeListRequest, opts ...grpc.CallOption) (*GetOfficeListResponse, error) {
	out := new(GetOfficeListResponse)
	err := c.cc.Invoke(ctx, "/customer.OfficeService/GetOfficeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OfficeServiceServer is the server API for OfficeService service.
// All implementations must embed UnimplementedOfficeServiceServer
// for forward compatibility
type OfficeServiceServer interface {
	CreateOffice(context.Context, *CreateOfficeRequest) (*CreateOfficeResponse, error)
	GetOfficeList(context.Context, *GetOfficeListRequest) (*GetOfficeListResponse, error)
	mustEmbedUnimplementedOfficeServiceServer()
}

// UnimplementedOfficeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOfficeServiceServer struct {
}

func (UnimplementedOfficeServiceServer) CreateOffice(context.Context, *CreateOfficeRequest) (*CreateOfficeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOffice not implemented")
}
func (UnimplementedOfficeServiceServer) GetOfficeList(context.Context, *GetOfficeListRequest) (*GetOfficeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOfficeList not implemented")
}
func (UnimplementedOfficeServiceServer) mustEmbedUnimplementedOfficeServiceServer() {}

// UnsafeOfficeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OfficeServiceServer will
// result in compilation errors.
type UnsafeOfficeServiceServer interface {
	mustEmbedUnimplementedOfficeServiceServer()
}

func RegisterOfficeServiceServer(s grpc.ServiceRegistrar, srv OfficeServiceServer) {
	s.RegisterService(&OfficeService_ServiceDesc, srv)
}

func _OfficeService_CreateOffice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOfficeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfficeServiceServer).CreateOffice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.OfficeService/CreateOffice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfficeServiceServer).CreateOffice(ctx, req.(*CreateOfficeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfficeService_GetOfficeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOfficeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfficeServiceServer).GetOfficeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.OfficeService/GetOfficeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfficeServiceServer).GetOfficeList(ctx, req.(*GetOfficeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OfficeService_ServiceDesc is the grpc.ServiceDesc for OfficeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OfficeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customer.OfficeService",
	HandlerType: (*OfficeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOffice",
			Handler:    _OfficeService_CreateOffice_Handler,
		},
		{
			MethodName: "GetOfficeList",
			Handler:    _OfficeService_GetOfficeList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/customer/customer_office.proto",
}
