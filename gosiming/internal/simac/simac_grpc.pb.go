// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package simac

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SiMacClient is the client API for SiMac service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SiMacClient interface {
	// Get Mac Details
	GetMacDetails(ctx context.Context, opts ...grpc.CallOption) (SiMac_GetMacDetailsClient, error)
	// Create a MAC
	Create(ctx context.Context, in *Commissioning, opts ...grpc.CallOption) (*MacReply, error)
	// Delete a Mac
	Delete(ctx context.Context, in *DevEui, opts ...grpc.CallOption) (*MacReply, error)
	// Send stream Join Requests and returns stream the results
	Join(ctx context.Context, opts ...grpc.CallOption) (SiMac_JoinClient, error)
	// Accepts a stream of Mac configurations while receiving MacState updates
	Configure(ctx context.Context, opts ...grpc.CallOption) (SiMac_ConfigureClient, error)
	// Accepts a stream of uplink requests while receiving completed uplink status
	SendUplink(ctx context.Context, opts ...grpc.CallOption) (SiMac_SendUplinkClient, error)
}

type siMacClient struct {
	cc grpc.ClientConnInterface
}

func NewSiMacClient(cc grpc.ClientConnInterface) SiMacClient {
	return &siMacClient{cc}
}

func (c *siMacClient) GetMacDetails(ctx context.Context, opts ...grpc.CallOption) (SiMac_GetMacDetailsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SiMac_serviceDesc.Streams[0], "/siming.SiMac/GetMacDetails", opts...)
	if err != nil {
		return nil, err
	}
	x := &siMacGetMacDetailsClient{stream}
	return x, nil
}

type SiMac_GetMacDetailsClient interface {
	Send(*DevEui) error
	Recv() (*MacDetails, error)
	grpc.ClientStream
}

type siMacGetMacDetailsClient struct {
	grpc.ClientStream
}

func (x *siMacGetMacDetailsClient) Send(m *DevEui) error {
	return x.ClientStream.SendMsg(m)
}

func (x *siMacGetMacDetailsClient) Recv() (*MacDetails, error) {
	m := new(MacDetails)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *siMacClient) Create(ctx context.Context, in *Commissioning, opts ...grpc.CallOption) (*MacReply, error) {
	out := new(MacReply)
	err := c.cc.Invoke(ctx, "/siming.SiMac/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siMacClient) Delete(ctx context.Context, in *DevEui, opts ...grpc.CallOption) (*MacReply, error) {
	out := new(MacReply)
	err := c.cc.Invoke(ctx, "/siming.SiMac/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siMacClient) Join(ctx context.Context, opts ...grpc.CallOption) (SiMac_JoinClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SiMac_serviceDesc.Streams[1], "/siming.SiMac/Join", opts...)
	if err != nil {
		return nil, err
	}
	x := &siMacJoinClient{stream}
	return x, nil
}

type SiMac_JoinClient interface {
	Send(*JoinRequest) error
	Recv() (*JoinReply, error)
	grpc.ClientStream
}

type siMacJoinClient struct {
	grpc.ClientStream
}

func (x *siMacJoinClient) Send(m *JoinRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *siMacJoinClient) Recv() (*JoinReply, error) {
	m := new(JoinReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *siMacClient) Configure(ctx context.Context, opts ...grpc.CallOption) (SiMac_ConfigureClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SiMac_serviceDesc.Streams[2], "/siming.SiMac/Configure", opts...)
	if err != nil {
		return nil, err
	}
	x := &siMacConfigureClient{stream}
	return x, nil
}

type SiMac_ConfigureClient interface {
	Send(*ConfigRequest) error
	Recv() (*MacReply, error)
	grpc.ClientStream
}

type siMacConfigureClient struct {
	grpc.ClientStream
}

func (x *siMacConfigureClient) Send(m *ConfigRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *siMacConfigureClient) Recv() (*MacReply, error) {
	m := new(MacReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *siMacClient) SendUplink(ctx context.Context, opts ...grpc.CallOption) (SiMac_SendUplinkClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SiMac_serviceDesc.Streams[3], "/siming.SiMac/SendUplink", opts...)
	if err != nil {
		return nil, err
	}
	x := &siMacSendUplinkClient{stream}
	return x, nil
}

type SiMac_SendUplinkClient interface {
	Send(*UplinkRequest) error
	Recv() (*UplinkStatus, error)
	grpc.ClientStream
}

type siMacSendUplinkClient struct {
	grpc.ClientStream
}

func (x *siMacSendUplinkClient) Send(m *UplinkRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *siMacSendUplinkClient) Recv() (*UplinkStatus, error) {
	m := new(UplinkStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SiMacServer is the server API for SiMac service.
// All implementations must embed UnimplementedSiMacServer
// for forward compatibility
type SiMacServer interface {
	// Get Mac Details
	GetMacDetails(SiMac_GetMacDetailsServer) error
	// Create a MAC
	Create(context.Context, *Commissioning) (*MacReply, error)
	// Delete a Mac
	Delete(context.Context, *DevEui) (*MacReply, error)
	// Send stream Join Requests and returns stream the results
	Join(SiMac_JoinServer) error
	// Accepts a stream of Mac configurations while receiving MacState updates
	Configure(SiMac_ConfigureServer) error
	// Accepts a stream of uplink requests while receiving completed uplink status
	SendUplink(SiMac_SendUplinkServer) error
	mustEmbedUnimplementedSiMacServer()
}

// UnimplementedSiMacServer must be embedded to have forward compatible implementations.
type UnimplementedSiMacServer struct {
}

func (UnimplementedSiMacServer) GetMacDetails(SiMac_GetMacDetailsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMacDetails not implemented")
}
func (UnimplementedSiMacServer) Create(context.Context, *Commissioning) (*MacReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSiMacServer) Delete(context.Context, *DevEui) (*MacReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSiMacServer) Join(SiMac_JoinServer) error {
	return status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedSiMacServer) Configure(SiMac_ConfigureServer) error {
	return status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedSiMacServer) SendUplink(SiMac_SendUplinkServer) error {
	return status.Errorf(codes.Unimplemented, "method SendUplink not implemented")
}
func (UnimplementedSiMacServer) mustEmbedUnimplementedSiMacServer() {}

// UnsafeSiMacServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SiMacServer will
// result in compilation errors.
type UnsafeSiMacServer interface {
	mustEmbedUnimplementedSiMacServer()
}

func RegisterSiMacServer(s *grpc.Server, srv SiMacServer) {
	s.RegisterService(&_SiMac_serviceDesc, srv)
}

func _SiMac_GetMacDetails_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SiMacServer).GetMacDetails(&siMacGetMacDetailsServer{stream})
}

type SiMac_GetMacDetailsServer interface {
	Send(*MacDetails) error
	Recv() (*DevEui, error)
	grpc.ServerStream
}

type siMacGetMacDetailsServer struct {
	grpc.ServerStream
}

func (x *siMacGetMacDetailsServer) Send(m *MacDetails) error {
	return x.ServerStream.SendMsg(m)
}

func (x *siMacGetMacDetailsServer) Recv() (*DevEui, error) {
	m := new(DevEui)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SiMac_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Commissioning)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiMacServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/siming.SiMac/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiMacServer).Create(ctx, req.(*Commissioning))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiMac_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DevEui)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiMacServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/siming.SiMac/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiMacServer).Delete(ctx, req.(*DevEui))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiMac_Join_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SiMacServer).Join(&siMacJoinServer{stream})
}

type SiMac_JoinServer interface {
	Send(*JoinReply) error
	Recv() (*JoinRequest, error)
	grpc.ServerStream
}

type siMacJoinServer struct {
	grpc.ServerStream
}

func (x *siMacJoinServer) Send(m *JoinReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *siMacJoinServer) Recv() (*JoinRequest, error) {
	m := new(JoinRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SiMac_Configure_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SiMacServer).Configure(&siMacConfigureServer{stream})
}

type SiMac_ConfigureServer interface {
	Send(*MacReply) error
	Recv() (*ConfigRequest, error)
	grpc.ServerStream
}

type siMacConfigureServer struct {
	grpc.ServerStream
}

func (x *siMacConfigureServer) Send(m *MacReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *siMacConfigureServer) Recv() (*ConfigRequest, error) {
	m := new(ConfigRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SiMac_SendUplink_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SiMacServer).SendUplink(&siMacSendUplinkServer{stream})
}

type SiMac_SendUplinkServer interface {
	Send(*UplinkStatus) error
	Recv() (*UplinkRequest, error)
	grpc.ServerStream
}

type siMacSendUplinkServer struct {
	grpc.ServerStream
}

func (x *siMacSendUplinkServer) Send(m *UplinkStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *siMacSendUplinkServer) Recv() (*UplinkRequest, error) {
	m := new(UplinkRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SiMac_serviceDesc = grpc.ServiceDesc{
	ServiceName: "siming.SiMac",
	HandlerType: (*SiMacServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SiMac_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SiMac_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMacDetails",
			Handler:       _SiMac_GetMacDetails_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Join",
			Handler:       _SiMac_Join_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Configure",
			Handler:       _SiMac_Configure_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SendUplink",
			Handler:       _SiMac_SendUplink_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "simac.proto",
}