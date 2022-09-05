// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: routeguide/route_guide.proto

package routeguide

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

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteGuideClient interface {
	ProcessData(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_ProcessDataClient, error)
	GetLatestCatDB(ctx context.Context, in *C2SData, opts ...grpc.CallOption) (*S2CData, error)
}

type routeGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuideClient(cc grpc.ClientConnInterface) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) ProcessData(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_ProcessDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[0], "/routeguide.RouteGuide/ProcessData", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideProcessDataClient{stream}
	return x, nil
}

type RouteGuide_ProcessDataClient interface {
	Send(*C2SData) error
	Recv() (*S2CData, error)
	grpc.ClientStream
}

type routeGuideProcessDataClient struct {
	grpc.ClientStream
}

func (x *routeGuideProcessDataClient) Send(m *C2SData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideProcessDataClient) Recv() (*S2CData, error) {
	m := new(S2CData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) GetLatestCatDB(ctx context.Context, in *C2SData, opts ...grpc.CallOption) (*S2CData, error) {
	out := new(S2CData)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/GetLatestCatDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteGuideServer is the server API for RouteGuide service.
// All implementations must embed UnimplementedRouteGuideServer
// for forward compatibility
type RouteGuideServer interface {
	ProcessData(RouteGuide_ProcessDataServer) error
	GetLatestCatDB(context.Context, *C2SData) (*S2CData, error)
	mustEmbedUnimplementedRouteGuideServer()
}

// UnimplementedRouteGuideServer must be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServer struct {
}

func (UnimplementedRouteGuideServer) ProcessData(RouteGuide_ProcessDataServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessData not implemented")
}
func (UnimplementedRouteGuideServer) GetLatestCatDB(context.Context, *C2SData) (*S2CData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestCatDB not implemented")
}
func (UnimplementedRouteGuideServer) mustEmbedUnimplementedRouteGuideServer() {}

// UnsafeRouteGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteGuideServer will
// result in compilation errors.
type UnsafeRouteGuideServer interface {
	mustEmbedUnimplementedRouteGuideServer()
}

func RegisterRouteGuideServer(s grpc.ServiceRegistrar, srv RouteGuideServer) {
	s.RegisterService(&RouteGuide_ServiceDesc, srv)
}

func _RouteGuide_ProcessData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).ProcessData(&routeGuideProcessDataServer{stream})
}

type RouteGuide_ProcessDataServer interface {
	Send(*S2CData) error
	Recv() (*C2SData, error)
	grpc.ServerStream
}

type routeGuideProcessDataServer struct {
	grpc.ServerStream
}

func (x *routeGuideProcessDataServer) Send(m *S2CData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideProcessDataServer) Recv() (*C2SData, error) {
	m := new(C2SData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_GetLatestCatDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2SData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).GetLatestCatDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/GetLatestCatDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).GetLatestCatDB(ctx, req.(*C2SData))
	}
	return interceptor(ctx, in, info, handler)
}

// RouteGuide_ServiceDesc is the grpc.ServiceDesc for RouteGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatestCatDB",
			Handler:    _RouteGuide_GetLatestCatDB_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessData",
			Handler:       _RouteGuide_ProcessData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "routeguide/route_guide.proto",
}
