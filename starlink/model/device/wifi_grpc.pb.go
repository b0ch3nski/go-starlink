// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: spacex/api/device/wifi.proto

package device

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Mesh_MeshStream_FullMethodName = "/SpaceX.API.Device.Mesh/MeshStream"
)

// MeshClient is the client API for Mesh service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeshClient interface {
	MeshStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ToController, FromController], error)
}

type meshClient struct {
	cc grpc.ClientConnInterface
}

func NewMeshClient(cc grpc.ClientConnInterface) MeshClient {
	return &meshClient{cc}
}

func (c *meshClient) MeshStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ToController, FromController], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Mesh_ServiceDesc.Streams[0], Mesh_MeshStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ToController, FromController]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Mesh_MeshStreamClient = grpc.BidiStreamingClient[ToController, FromController]

// MeshServer is the server API for Mesh service.
// All implementations must embed UnimplementedMeshServer
// for forward compatibility.
type MeshServer interface {
	MeshStream(grpc.BidiStreamingServer[ToController, FromController]) error
	mustEmbedUnimplementedMeshServer()
}

// UnimplementedMeshServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMeshServer struct{}

func (UnimplementedMeshServer) MeshStream(grpc.BidiStreamingServer[ToController, FromController]) error {
	return status.Errorf(codes.Unimplemented, "method MeshStream not implemented")
}
func (UnimplementedMeshServer) mustEmbedUnimplementedMeshServer() {}
func (UnimplementedMeshServer) testEmbeddedByValue()              {}

// UnsafeMeshServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeshServer will
// result in compilation errors.
type UnsafeMeshServer interface {
	mustEmbedUnimplementedMeshServer()
}

func RegisterMeshServer(s grpc.ServiceRegistrar, srv MeshServer) {
	// If the following call pancis, it indicates UnimplementedMeshServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Mesh_ServiceDesc, srv)
}

func _Mesh_MeshStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MeshServer).MeshStream(&grpc.GenericServerStream[ToController, FromController]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Mesh_MeshStreamServer = grpc.BidiStreamingServer[ToController, FromController]

// Mesh_ServiceDesc is the grpc.ServiceDesc for Mesh service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mesh_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SpaceX.API.Device.Mesh",
	HandlerType: (*MeshServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MeshStream",
			Handler:       _Mesh_MeshStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "spacex/api/device/wifi.proto",
}
