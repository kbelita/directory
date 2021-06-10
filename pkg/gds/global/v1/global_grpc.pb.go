// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package global

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ReplicationClient is the client API for Replication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplicationClient interface {
	// During gossip, the initiating replica sends a randomly selected remote peer the
	// version vectors of all objects it currently stores. The remote peer should
	// respond with updates that correspond to more recent versions of the objects. The
	// remote peer can than also make a reciprocal request for updates by sending the
	// set of versions requested that were more recent on the initiating replica, and
	// use a partial flag to indicate that it is requesting specific versions. This
	// mechanism implements bilateral anti-entropy: a push and pull gossip.
	Gossip(ctx context.Context, in *VersionVectors, opts ...grpc.CallOption) (*Updates, error)
}

type replicationClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicationClient(cc grpc.ClientConnInterface) ReplicationClient {
	return &replicationClient{cc}
}

func (c *replicationClient) Gossip(ctx context.Context, in *VersionVectors, opts ...grpc.CallOption) (*Updates, error) {
	out := new(Updates)
	err := c.cc.Invoke(ctx, "/gds.global.v1.Replication/Gossip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicationServer is the server API for Replication service.
// All implementations must embed UnimplementedReplicationServer
// for forward compatibility
type ReplicationServer interface {
	// During gossip, the initiating replica sends a randomly selected remote peer the
	// version vectors of all objects it currently stores. The remote peer should
	// respond with updates that correspond to more recent versions of the objects. The
	// remote peer can than also make a reciprocal request for updates by sending the
	// set of versions requested that were more recent on the initiating replica, and
	// use a partial flag to indicate that it is requesting specific versions. This
	// mechanism implements bilateral anti-entropy: a push and pull gossip.
	Gossip(context.Context, *VersionVectors) (*Updates, error)
	mustEmbedUnimplementedReplicationServer()
}

// UnimplementedReplicationServer must be embedded to have forward compatible implementations.
type UnimplementedReplicationServer struct {
}

func (UnimplementedReplicationServer) Gossip(context.Context, *VersionVectors) (*Updates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gossip not implemented")
}
func (UnimplementedReplicationServer) mustEmbedUnimplementedReplicationServer() {}

// UnsafeReplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicationServer will
// result in compilation errors.
type UnsafeReplicationServer interface {
	mustEmbedUnimplementedReplicationServer()
}

func RegisterReplicationServer(s grpc.ServiceRegistrar, srv ReplicationServer) {
	s.RegisterService(&_Replication_serviceDesc, srv)
}

func _Replication_Gossip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionVectors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).Gossip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gds.global.v1.Replication/Gossip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).Gossip(ctx, req.(*VersionVectors))
	}
	return interceptor(ctx, in, info, handler)
}

var _Replication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gds.global.v1.Replication",
	HandlerType: (*ReplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Gossip",
			Handler:    _Replication_Gossip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gds/global/v1/global.proto",
}
