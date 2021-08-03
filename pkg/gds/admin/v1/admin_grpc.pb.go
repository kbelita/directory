// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package admin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DirectoryAdministrationClient is the client API for DirectoryAdministration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectoryAdministrationClient interface {
	Review(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewReply, error)
	Resend(ctx context.Context, in *ResendRequest, opts ...grpc.CallOption) (*ResendReply, error)
}

type directoryAdministrationClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectoryAdministrationClient(cc grpc.ClientConnInterface) DirectoryAdministrationClient {
	return &directoryAdministrationClient{cc}
}

func (c *directoryAdministrationClient) Review(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewReply, error) {
	out := new(ReviewReply)
	err := c.cc.Invoke(ctx, "/gds.admin.v1.DirectoryAdministration/Review", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryAdministrationClient) Resend(ctx context.Context, in *ResendRequest, opts ...grpc.CallOption) (*ResendReply, error) {
	out := new(ResendReply)
	err := c.cc.Invoke(ctx, "/gds.admin.v1.DirectoryAdministration/Resend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectoryAdministrationServer is the server API for DirectoryAdministration service.
// All implementations must embed UnimplementedDirectoryAdministrationServer
// for forward compatibility
type DirectoryAdministrationServer interface {
	Review(context.Context, *ReviewRequest) (*ReviewReply, error)
	Resend(context.Context, *ResendRequest) (*ResendReply, error)
	mustEmbedUnimplementedDirectoryAdministrationServer()
}

// UnimplementedDirectoryAdministrationServer must be embedded to have forward compatible implementations.
type UnimplementedDirectoryAdministrationServer struct {
}

func (UnimplementedDirectoryAdministrationServer) Review(context.Context, *ReviewRequest) (*ReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Review not implemented")
}
func (UnimplementedDirectoryAdministrationServer) Resend(context.Context, *ResendRequest) (*ResendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resend not implemented")
}
func (UnimplementedDirectoryAdministrationServer) mustEmbedUnimplementedDirectoryAdministrationServer() {
}

// UnsafeDirectoryAdministrationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectoryAdministrationServer will
// result in compilation errors.
type UnsafeDirectoryAdministrationServer interface {
	mustEmbedUnimplementedDirectoryAdministrationServer()
}

func RegisterDirectoryAdministrationServer(s grpc.ServiceRegistrar, srv DirectoryAdministrationServer) {
	s.RegisterService(&_DirectoryAdministration_serviceDesc, srv)
}

func _DirectoryAdministration_Review_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryAdministrationServer).Review(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gds.admin.v1.DirectoryAdministration/Review",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryAdministrationServer).Review(ctx, req.(*ReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryAdministration_Resend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryAdministrationServer).Resend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gds.admin.v1.DirectoryAdministration/Resend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryAdministrationServer).Resend(ctx, req.(*ResendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DirectoryAdministration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gds.admin.v1.DirectoryAdministration",
	HandlerType: (*DirectoryAdministrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Review",
			Handler:    _DirectoryAdministration_Review_Handler,
		},
		{
			MethodName: "Resend",
			Handler:    _DirectoryAdministration_Resend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gds/admin/v1/admin.proto",
}
