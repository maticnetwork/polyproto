// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: heimdall/heimdall.proto

package heimdall

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HeimdallClient is the client API for Heimdall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeimdallClient interface {
	Span(ctx context.Context, in *SpanRequest, opts ...grpc.CallOption) (*SpanResponse, error)
	StateSyncEvents(ctx context.Context, in *StateSyncEventsRequest, opts ...grpc.CallOption) (Heimdall_StateSyncEventsClient, error)
	FetchCheckpoint(ctx context.Context, in *FetchCheckpointRequest, opts ...grpc.CallOption) (*FetchCheckpointResponse, error)
	FetchCheckpointCount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FetchCheckpointCountResponse, error)
}

type heimdallClient struct {
	cc grpc.ClientConnInterface
}

func NewHeimdallClient(cc grpc.ClientConnInterface) HeimdallClient {
	return &heimdallClient{cc}
}

func (c *heimdallClient) Span(ctx context.Context, in *SpanRequest, opts ...grpc.CallOption) (*SpanResponse, error) {
	out := new(SpanResponse)
	err := c.cc.Invoke(ctx, "/heimdall.Heimdall/Span", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallClient) StateSyncEvents(ctx context.Context, in *StateSyncEventsRequest, opts ...grpc.CallOption) (Heimdall_StateSyncEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Heimdall_ServiceDesc.Streams[0], "/heimdall.Heimdall/StateSyncEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &heimdallStateSyncEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Heimdall_StateSyncEventsClient interface {
	Recv() (*StateSyncEventsResponse, error)
	grpc.ClientStream
}

type heimdallStateSyncEventsClient struct {
	grpc.ClientStream
}

func (x *heimdallStateSyncEventsClient) Recv() (*StateSyncEventsResponse, error) {
	m := new(StateSyncEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *heimdallClient) FetchCheckpoint(ctx context.Context, in *FetchCheckpointRequest, opts ...grpc.CallOption) (*FetchCheckpointResponse, error) {
	out := new(FetchCheckpointResponse)
	err := c.cc.Invoke(ctx, "/heimdall.Heimdall/FetchCheckpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallClient) FetchCheckpointCount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FetchCheckpointCountResponse, error) {
	out := new(FetchCheckpointCountResponse)
	err := c.cc.Invoke(ctx, "/heimdall.Heimdall/FetchCheckpointCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeimdallServer is the server API for Heimdall service.
// All implementations must embed UnimplementedHeimdallServer
// for forward compatibility
type HeimdallServer interface {
	Span(context.Context, *SpanRequest) (*SpanResponse, error)
	StateSyncEvents(*StateSyncEventsRequest, Heimdall_StateSyncEventsServer) error
	FetchCheckpoint(context.Context, *FetchCheckpointRequest) (*FetchCheckpointResponse, error)
	FetchCheckpointCount(context.Context, *emptypb.Empty) (*FetchCheckpointCountResponse, error)
	mustEmbedUnimplementedHeimdallServer()
}

// UnimplementedHeimdallServer must be embedded to have forward compatible implementations.
type UnimplementedHeimdallServer struct {
}

func (UnimplementedHeimdallServer) Span(context.Context, *SpanRequest) (*SpanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Span not implemented")
}
func (UnimplementedHeimdallServer) StateSyncEvents(*StateSyncEventsRequest, Heimdall_StateSyncEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StateSyncEvents not implemented")
}
func (UnimplementedHeimdallServer) FetchCheckpoint(context.Context, *FetchCheckpointRequest) (*FetchCheckpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCheckpoint not implemented")
}
func (UnimplementedHeimdallServer) FetchCheckpointCount(context.Context, *emptypb.Empty) (*FetchCheckpointCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCheckpointCount not implemented")
}
func (UnimplementedHeimdallServer) mustEmbedUnimplementedHeimdallServer() {}

// UnsafeHeimdallServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeimdallServer will
// result in compilation errors.
type UnsafeHeimdallServer interface {
	mustEmbedUnimplementedHeimdallServer()
}

func RegisterHeimdallServer(s grpc.ServiceRegistrar, srv HeimdallServer) {
	s.RegisterService(&Heimdall_ServiceDesc, srv)
}

func _Heimdall_Span_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServer).Span(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.Heimdall/Span",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServer).Span(ctx, req.(*SpanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Heimdall_StateSyncEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StateSyncEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HeimdallServer).StateSyncEvents(m, &heimdallStateSyncEventsServer{stream})
}

type Heimdall_StateSyncEventsServer interface {
	Send(*StateSyncEventsResponse) error
	grpc.ServerStream
}

type heimdallStateSyncEventsServer struct {
	grpc.ServerStream
}

func (x *heimdallStateSyncEventsServer) Send(m *StateSyncEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Heimdall_FetchCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServer).FetchCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.Heimdall/FetchCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServer).FetchCheckpoint(ctx, req.(*FetchCheckpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Heimdall_FetchCheckpointCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServer).FetchCheckpointCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.Heimdall/FetchCheckpointCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServer).FetchCheckpointCount(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Heimdall_ServiceDesc is the grpc.ServiceDesc for Heimdall service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Heimdall_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "heimdall.Heimdall",
	HandlerType: (*HeimdallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Span",
			Handler:    _Heimdall_Span_Handler,
		},
		{
			MethodName: "FetchCheckpoint",
			Handler:    _Heimdall_FetchCheckpoint_Handler,
		},
		{
			MethodName: "FetchCheckpointCount",
			Handler:    _Heimdall_FetchCheckpointCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StateSyncEvents",
			Handler:       _Heimdall_StateSyncEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "heimdall/heimdall.proto",
}
