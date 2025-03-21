// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// P4RuntimeClient is the client API for P4Runtime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type P4RuntimeClient interface {
	// Update one or more P4 entities on the target.
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	// Read one or more P4 entities from the target.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (P4Runtime_ReadClient, error)
	// Sets the P4 forwarding-pipeline config.
	SetForwardingPipelineConfig(ctx context.Context, in *SetForwardingPipelineConfigRequest, opts ...grpc.CallOption) (*SetForwardingPipelineConfigResponse, error)
	// Gets the current P4 forwarding-pipeline config.
	GetForwardingPipelineConfig(ctx context.Context, in *GetForwardingPipelineConfigRequest, opts ...grpc.CallOption) (*GetForwardingPipelineConfigResponse, error)
	// Represents the bidirectional stream between the controller and the
	// switch (initiated by the controller), and is managed for the following
	// purposes:
	// - connection initiation through client arbitration
	// - indicating switch session liveness: the session is live when switch
	//   sends a positive client arbitration update to the controller, and is
	//   considered dead when either the stream breaks or the switch sends a
	//   negative update for client arbitration
	// - the controller sending/receiving packets to/from the switch
	// - streaming of notifications from the switch
	StreamChannel(ctx context.Context, opts ...grpc.CallOption) (P4Runtime_StreamChannelClient, error)
	Capabilities(ctx context.Context, in *CapabilitiesRequest, opts ...grpc.CallOption) (*CapabilitiesResponse, error)
}

type p4RuntimeClient struct {
	cc grpc.ClientConnInterface
}

func NewP4RuntimeClient(cc grpc.ClientConnInterface) P4RuntimeClient {
	return &p4RuntimeClient{cc}
}

func (c *p4RuntimeClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/p4.v1.P4Runtime/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p4RuntimeClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (P4Runtime_ReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &P4Runtime_ServiceDesc.Streams[0], "/p4.v1.P4Runtime/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &p4RuntimeReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type P4Runtime_ReadClient interface {
	Recv() (*ReadResponse, error)
	grpc.ClientStream
}

type p4RuntimeReadClient struct {
	grpc.ClientStream
}

func (x *p4RuntimeReadClient) Recv() (*ReadResponse, error) {
	m := new(ReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *p4RuntimeClient) SetForwardingPipelineConfig(ctx context.Context, in *SetForwardingPipelineConfigRequest, opts ...grpc.CallOption) (*SetForwardingPipelineConfigResponse, error) {
	out := new(SetForwardingPipelineConfigResponse)
	err := c.cc.Invoke(ctx, "/p4.v1.P4Runtime/SetForwardingPipelineConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p4RuntimeClient) GetForwardingPipelineConfig(ctx context.Context, in *GetForwardingPipelineConfigRequest, opts ...grpc.CallOption) (*GetForwardingPipelineConfigResponse, error) {
	out := new(GetForwardingPipelineConfigResponse)
	err := c.cc.Invoke(ctx, "/p4.v1.P4Runtime/GetForwardingPipelineConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p4RuntimeClient) StreamChannel(ctx context.Context, opts ...grpc.CallOption) (P4Runtime_StreamChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &P4Runtime_ServiceDesc.Streams[1], "/p4.v1.P4Runtime/StreamChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &p4RuntimeStreamChannelClient{stream}
	return x, nil
}

type P4Runtime_StreamChannelClient interface {
	Send(*StreamMessageRequest) error
	Recv() (*StreamMessageResponse, error)
	grpc.ClientStream
}

type p4RuntimeStreamChannelClient struct {
	grpc.ClientStream
}

func (x *p4RuntimeStreamChannelClient) Send(m *StreamMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *p4RuntimeStreamChannelClient) Recv() (*StreamMessageResponse, error) {
	m := new(StreamMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *p4RuntimeClient) Capabilities(ctx context.Context, in *CapabilitiesRequest, opts ...grpc.CallOption) (*CapabilitiesResponse, error) {
	out := new(CapabilitiesResponse)
	err := c.cc.Invoke(ctx, "/p4.v1.P4Runtime/Capabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// P4RuntimeServer is the server API for P4Runtime service.
// All implementations must embed UnimplementedP4RuntimeServer
// for forward compatibility
type P4RuntimeServer interface {
	// Update one or more P4 entities on the target.
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	// Read one or more P4 entities from the target.
	Read(*ReadRequest, P4Runtime_ReadServer) error
	// Sets the P4 forwarding-pipeline config.
	SetForwardingPipelineConfig(context.Context, *SetForwardingPipelineConfigRequest) (*SetForwardingPipelineConfigResponse, error)
	// Gets the current P4 forwarding-pipeline config.
	GetForwardingPipelineConfig(context.Context, *GetForwardingPipelineConfigRequest) (*GetForwardingPipelineConfigResponse, error)
	// Represents the bidirectional stream between the controller and the
	// switch (initiated by the controller), and is managed for the following
	// purposes:
	// - connection initiation through client arbitration
	// - indicating switch session liveness: the session is live when switch
	//   sends a positive client arbitration update to the controller, and is
	//   considered dead when either the stream breaks or the switch sends a
	//   negative update for client arbitration
	// - the controller sending/receiving packets to/from the switch
	// - streaming of notifications from the switch
	StreamChannel(P4Runtime_StreamChannelServer) error
	Capabilities(context.Context, *CapabilitiesRequest) (*CapabilitiesResponse, error)
	mustEmbedUnimplementedP4RuntimeServer()
}

// UnimplementedP4RuntimeServer must be embedded to have forward compatible implementations.
type UnimplementedP4RuntimeServer struct {
}

func (UnimplementedP4RuntimeServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedP4RuntimeServer) Read(*ReadRequest, P4Runtime_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedP4RuntimeServer) SetForwardingPipelineConfig(context.Context, *SetForwardingPipelineConfigRequest) (*SetForwardingPipelineConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetForwardingPipelineConfig not implemented")
}
func (UnimplementedP4RuntimeServer) GetForwardingPipelineConfig(context.Context, *GetForwardingPipelineConfigRequest) (*GetForwardingPipelineConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForwardingPipelineConfig not implemented")
}
func (UnimplementedP4RuntimeServer) StreamChannel(P4Runtime_StreamChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamChannel not implemented")
}
func (UnimplementedP4RuntimeServer) Capabilities(context.Context, *CapabilitiesRequest) (*CapabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Capabilities not implemented")
}
func (UnimplementedP4RuntimeServer) mustEmbedUnimplementedP4RuntimeServer() {}

// UnsafeP4RuntimeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to P4RuntimeServer will
// result in compilation errors.
type UnsafeP4RuntimeServer interface {
	mustEmbedUnimplementedP4RuntimeServer()
}

func RegisterP4RuntimeServer(s grpc.ServiceRegistrar, srv P4RuntimeServer) {
	s.RegisterService(&P4Runtime_ServiceDesc, srv)
}

func _P4Runtime_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P4RuntimeServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p4.v1.P4Runtime/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P4RuntimeServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P4Runtime_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(P4RuntimeServer).Read(m, &p4RuntimeReadServer{stream})
}

type P4Runtime_ReadServer interface {
	Send(*ReadResponse) error
	grpc.ServerStream
}

type p4RuntimeReadServer struct {
	grpc.ServerStream
}

func (x *p4RuntimeReadServer) Send(m *ReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _P4Runtime_SetForwardingPipelineConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetForwardingPipelineConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P4RuntimeServer).SetForwardingPipelineConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p4.v1.P4Runtime/SetForwardingPipelineConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P4RuntimeServer).SetForwardingPipelineConfig(ctx, req.(*SetForwardingPipelineConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P4Runtime_GetForwardingPipelineConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetForwardingPipelineConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P4RuntimeServer).GetForwardingPipelineConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p4.v1.P4Runtime/GetForwardingPipelineConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P4RuntimeServer).GetForwardingPipelineConfig(ctx, req.(*GetForwardingPipelineConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P4Runtime_StreamChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(P4RuntimeServer).StreamChannel(&p4RuntimeStreamChannelServer{stream})
}

type P4Runtime_StreamChannelServer interface {
	Send(*StreamMessageResponse) error
	Recv() (*StreamMessageRequest, error)
	grpc.ServerStream
}

type p4RuntimeStreamChannelServer struct {
	grpc.ServerStream
}

func (x *p4RuntimeStreamChannelServer) Send(m *StreamMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *p4RuntimeStreamChannelServer) Recv() (*StreamMessageRequest, error) {
	m := new(StreamMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _P4Runtime_Capabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P4RuntimeServer).Capabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p4.v1.P4Runtime/Capabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P4RuntimeServer).Capabilities(ctx, req.(*CapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// P4Runtime_ServiceDesc is the grpc.ServiceDesc for P4Runtime service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var P4Runtime_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "p4.v1.P4Runtime",
	HandlerType: (*P4RuntimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _P4Runtime_Write_Handler,
		},
		{
			MethodName: "SetForwardingPipelineConfig",
			Handler:    _P4Runtime_SetForwardingPipelineConfig_Handler,
		},
		{
			MethodName: "GetForwardingPipelineConfig",
			Handler:    _P4Runtime_GetForwardingPipelineConfig_Handler,
		},
		{
			MethodName: "Capabilities",
			Handler:    _P4Runtime_Capabilities_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _P4Runtime_Read_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamChannel",
			Handler:       _P4Runtime_StreamChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "p4/v1/p4runtime.proto",
}
