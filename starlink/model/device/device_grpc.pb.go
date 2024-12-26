// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: spacex/api/device/device.proto

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
	Device_Stream_FullMethodName = "/SpaceX.API.Device.Device/Stream"
	Device_Handle_FullMethodName = "/SpaceX.API.Device.Device/Handle"
)

// DeviceClient is the client API for Device service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ToDevice, FromDevice], error)
	Handle(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type deviceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceClient(cc grpc.ClientConnInterface) DeviceClient {
	return &deviceClient{cc}
}

func (c *deviceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ToDevice, FromDevice], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Device_ServiceDesc.Streams[0], Device_Stream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ToDevice, FromDevice]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Device_StreamClient = grpc.BidiStreamingClient[ToDevice, FromDevice]

func (c *deviceClient) Handle(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Device_Handle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServer is the server API for Device service.
// All implementations must embed UnimplementedDeviceServer
// for forward compatibility.
type DeviceServer interface {
	Stream(grpc.BidiStreamingServer[ToDevice, FromDevice]) error
	Handle(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedDeviceServer()
}

// UnimplementedDeviceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeviceServer struct{}

func (UnimplementedDeviceServer) Stream(grpc.BidiStreamingServer[ToDevice, FromDevice]) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedDeviceServer) Handle(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}
func (UnimplementedDeviceServer) mustEmbedUnimplementedDeviceServer() {}
func (UnimplementedDeviceServer) testEmbeddedByValue()                {}

// UnsafeDeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServer will
// result in compilation errors.
type UnsafeDeviceServer interface {
	mustEmbedUnimplementedDeviceServer()
}

func RegisterDeviceServer(s grpc.ServiceRegistrar, srv DeviceServer) {
	// If the following call pancis, it indicates UnimplementedDeviceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Device_ServiceDesc, srv)
}

func _Device_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeviceServer).Stream(&grpc.GenericServerStream[ToDevice, FromDevice]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Device_StreamServer = grpc.BidiStreamingServer[ToDevice, FromDevice]

func _Device_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Device_Handle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).Handle(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Device_ServiceDesc is the grpc.ServiceDesc for Device service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Device_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SpaceX.API.Device.Device",
	HandlerType: (*DeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Device_Handle_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Device_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "spacex/api/device/device.proto",
}
