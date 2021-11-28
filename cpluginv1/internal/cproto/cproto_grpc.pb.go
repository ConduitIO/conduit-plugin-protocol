// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cproto

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

// SourcePluginClient is the client API for SourcePlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SourcePluginClient interface {
	Configure(ctx context.Context, in *Source_Configure_Request, opts ...grpc.CallOption) (*Source_Configure_Response, error)
	Start(ctx context.Context, in *Source_Start_Request, opts ...grpc.CallOption) (*Source_Start_Response, error)
	Run(ctx context.Context, opts ...grpc.CallOption) (SourcePlugin_RunClient, error)
	Stop(ctx context.Context, in *Source_Stop_Request, opts ...grpc.CallOption) (*Source_Stop_Response, error)
}

type sourcePluginClient struct {
	cc grpc.ClientConnInterface
}

func NewSourcePluginClient(cc grpc.ClientConnInterface) SourcePluginClient {
	return &sourcePluginClient{cc}
}

func (c *sourcePluginClient) Configure(ctx context.Context, in *Source_Configure_Request, opts ...grpc.CallOption) (*Source_Configure_Response, error) {
	out := new(Source_Configure_Response)
	err := c.cc.Invoke(ctx, "/cpluginv1.SourcePlugin/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcePluginClient) Start(ctx context.Context, in *Source_Start_Request, opts ...grpc.CallOption) (*Source_Start_Response, error) {
	out := new(Source_Start_Response)
	err := c.cc.Invoke(ctx, "/cpluginv1.SourcePlugin/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcePluginClient) Run(ctx context.Context, opts ...grpc.CallOption) (SourcePlugin_RunClient, error) {
	stream, err := c.cc.NewStream(ctx, &SourcePlugin_ServiceDesc.Streams[0], "/cpluginv1.SourcePlugin/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &sourcePluginRunClient{stream}
	return x, nil
}

type SourcePlugin_RunClient interface {
	Send(*Source_Run_Request) error
	Recv() (*Source_Run_Response, error)
	grpc.ClientStream
}

type sourcePluginRunClient struct {
	grpc.ClientStream
}

func (x *sourcePluginRunClient) Send(m *Source_Run_Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sourcePluginRunClient) Recv() (*Source_Run_Response, error) {
	m := new(Source_Run_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sourcePluginClient) Stop(ctx context.Context, in *Source_Stop_Request, opts ...grpc.CallOption) (*Source_Stop_Response, error) {
	out := new(Source_Stop_Response)
	err := c.cc.Invoke(ctx, "/cpluginv1.SourcePlugin/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SourcePluginServer is the server API for SourcePlugin service.
// All implementations must embed UnimplementedSourcePluginServer
// for forward compatibility
type SourcePluginServer interface {
	Configure(context.Context, *Source_Configure_Request) (*Source_Configure_Response, error)
	Start(context.Context, *Source_Start_Request) (*Source_Start_Response, error)
	Run(SourcePlugin_RunServer) error
	Stop(context.Context, *Source_Stop_Request) (*Source_Stop_Response, error)
	mustEmbedUnimplementedSourcePluginServer()
}

// UnimplementedSourcePluginServer must be embedded to have forward compatible implementations.
type UnimplementedSourcePluginServer struct {
}

func (UnimplementedSourcePluginServer) Configure(context.Context, *Source_Configure_Request) (*Source_Configure_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedSourcePluginServer) Start(context.Context, *Source_Start_Request) (*Source_Start_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedSourcePluginServer) Run(SourcePlugin_RunServer) error {
	return status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedSourcePluginServer) Stop(context.Context, *Source_Stop_Request) (*Source_Stop_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedSourcePluginServer) mustEmbedUnimplementedSourcePluginServer() {}

// UnsafeSourcePluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SourcePluginServer will
// result in compilation errors.
type UnsafeSourcePluginServer interface {
	mustEmbedUnimplementedSourcePluginServer()
}

func RegisterSourcePluginServer(s grpc.ServiceRegistrar, srv SourcePluginServer) {
	s.RegisterService(&SourcePlugin_ServiceDesc, srv)
}

func _SourcePlugin_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Source_Configure_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcePluginServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpluginv1.SourcePlugin/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcePluginServer).Configure(ctx, req.(*Source_Configure_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourcePlugin_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Source_Start_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcePluginServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpluginv1.SourcePlugin/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcePluginServer).Start(ctx, req.(*Source_Start_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourcePlugin_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SourcePluginServer).Run(&sourcePluginRunServer{stream})
}

type SourcePlugin_RunServer interface {
	Send(*Source_Run_Response) error
	Recv() (*Source_Run_Request, error)
	grpc.ServerStream
}

type sourcePluginRunServer struct {
	grpc.ServerStream
}

func (x *sourcePluginRunServer) Send(m *Source_Run_Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sourcePluginRunServer) Recv() (*Source_Run_Request, error) {
	m := new(Source_Run_Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SourcePlugin_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Source_Stop_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcePluginServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpluginv1.SourcePlugin/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcePluginServer).Stop(ctx, req.(*Source_Stop_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// SourcePlugin_ServiceDesc is the grpc.ServiceDesc for SourcePlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SourcePlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cpluginv1.SourcePlugin",
	HandlerType: (*SourcePluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _SourcePlugin_Configure_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _SourcePlugin_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _SourcePlugin_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _SourcePlugin_Run_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cproto.proto",
}

// SinkPluginClient is the client API for SinkPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SinkPluginClient interface {
	Configure(ctx context.Context, in *Sink_Configure_Request, opts ...grpc.CallOption) (*Sink_Configure_Response, error)
	Start(ctx context.Context, in *Sink_Start_Request, opts ...grpc.CallOption) (*Sink_Start_Response, error)
	Run(ctx context.Context, opts ...grpc.CallOption) (SinkPlugin_RunClient, error)
	Stop(ctx context.Context, in *Sink_Stop_Request, opts ...grpc.CallOption) (*Sink_Stop_Response, error)
}

type sinkPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewSinkPluginClient(cc grpc.ClientConnInterface) SinkPluginClient {
	return &sinkPluginClient{cc}
}

func (c *sinkPluginClient) Configure(ctx context.Context, in *Sink_Configure_Request, opts ...grpc.CallOption) (*Sink_Configure_Response, error) {
	out := new(Sink_Configure_Response)
	err := c.cc.Invoke(ctx, "/cpluginv1.SinkPlugin/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sinkPluginClient) Start(ctx context.Context, in *Sink_Start_Request, opts ...grpc.CallOption) (*Sink_Start_Response, error) {
	out := new(Sink_Start_Response)
	err := c.cc.Invoke(ctx, "/cpluginv1.SinkPlugin/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sinkPluginClient) Run(ctx context.Context, opts ...grpc.CallOption) (SinkPlugin_RunClient, error) {
	stream, err := c.cc.NewStream(ctx, &SinkPlugin_ServiceDesc.Streams[0], "/cpluginv1.SinkPlugin/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &sinkPluginRunClient{stream}
	return x, nil
}

type SinkPlugin_RunClient interface {
	Send(*Sink_Run_Request) error
	Recv() (*Sink_Run_Response, error)
	grpc.ClientStream
}

type sinkPluginRunClient struct {
	grpc.ClientStream
}

func (x *sinkPluginRunClient) Send(m *Sink_Run_Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sinkPluginRunClient) Recv() (*Sink_Run_Response, error) {
	m := new(Sink_Run_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sinkPluginClient) Stop(ctx context.Context, in *Sink_Stop_Request, opts ...grpc.CallOption) (*Sink_Stop_Response, error) {
	out := new(Sink_Stop_Response)
	err := c.cc.Invoke(ctx, "/cpluginv1.SinkPlugin/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SinkPluginServer is the server API for SinkPlugin service.
// All implementations must embed UnimplementedSinkPluginServer
// for forward compatibility
type SinkPluginServer interface {
	Configure(context.Context, *Sink_Configure_Request) (*Sink_Configure_Response, error)
	Start(context.Context, *Sink_Start_Request) (*Sink_Start_Response, error)
	Run(SinkPlugin_RunServer) error
	Stop(context.Context, *Sink_Stop_Request) (*Sink_Stop_Response, error)
	mustEmbedUnimplementedSinkPluginServer()
}

// UnimplementedSinkPluginServer must be embedded to have forward compatible implementations.
type UnimplementedSinkPluginServer struct {
}

func (UnimplementedSinkPluginServer) Configure(context.Context, *Sink_Configure_Request) (*Sink_Configure_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedSinkPluginServer) Start(context.Context, *Sink_Start_Request) (*Sink_Start_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedSinkPluginServer) Run(SinkPlugin_RunServer) error {
	return status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedSinkPluginServer) Stop(context.Context, *Sink_Stop_Request) (*Sink_Stop_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedSinkPluginServer) mustEmbedUnimplementedSinkPluginServer() {}

// UnsafeSinkPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SinkPluginServer will
// result in compilation errors.
type UnsafeSinkPluginServer interface {
	mustEmbedUnimplementedSinkPluginServer()
}

func RegisterSinkPluginServer(s grpc.ServiceRegistrar, srv SinkPluginServer) {
	s.RegisterService(&SinkPlugin_ServiceDesc, srv)
}

func _SinkPlugin_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sink_Configure_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinkPluginServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpluginv1.SinkPlugin/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinkPluginServer).Configure(ctx, req.(*Sink_Configure_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SinkPlugin_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sink_Start_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinkPluginServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpluginv1.SinkPlugin/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinkPluginServer).Start(ctx, req.(*Sink_Start_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SinkPlugin_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SinkPluginServer).Run(&sinkPluginRunServer{stream})
}

type SinkPlugin_RunServer interface {
	Send(*Sink_Run_Response) error
	Recv() (*Sink_Run_Request, error)
	grpc.ServerStream
}

type sinkPluginRunServer struct {
	grpc.ServerStream
}

func (x *sinkPluginRunServer) Send(m *Sink_Run_Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sinkPluginRunServer) Recv() (*Sink_Run_Request, error) {
	m := new(Sink_Run_Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SinkPlugin_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sink_Stop_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinkPluginServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpluginv1.SinkPlugin/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinkPluginServer).Stop(ctx, req.(*Sink_Stop_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// SinkPlugin_ServiceDesc is the grpc.ServiceDesc for SinkPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SinkPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cpluginv1.SinkPlugin",
	HandlerType: (*SinkPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _SinkPlugin_Configure_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _SinkPlugin_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _SinkPlugin_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _SinkPlugin_Run_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cproto.proto",
}

// SpecifierPluginClient is the client API for SpecifierPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpecifierPluginClient interface {
	Specify(ctx context.Context, in *Specifier_Specify_Request, opts ...grpc.CallOption) (*Specifier_Specify_Response, error)
}

type specifierPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewSpecifierPluginClient(cc grpc.ClientConnInterface) SpecifierPluginClient {
	return &specifierPluginClient{cc}
}

func (c *specifierPluginClient) Specify(ctx context.Context, in *Specifier_Specify_Request, opts ...grpc.CallOption) (*Specifier_Specify_Response, error) {
	out := new(Specifier_Specify_Response)
	err := c.cc.Invoke(ctx, "/cpluginv1.SpecifierPlugin/Specify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpecifierPluginServer is the server API for SpecifierPlugin service.
// All implementations must embed UnimplementedSpecifierPluginServer
// for forward compatibility
type SpecifierPluginServer interface {
	Specify(context.Context, *Specifier_Specify_Request) (*Specifier_Specify_Response, error)
	mustEmbedUnimplementedSpecifierPluginServer()
}

// UnimplementedSpecifierPluginServer must be embedded to have forward compatible implementations.
type UnimplementedSpecifierPluginServer struct {
}

func (UnimplementedSpecifierPluginServer) Specify(context.Context, *Specifier_Specify_Request) (*Specifier_Specify_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Specify not implemented")
}
func (UnimplementedSpecifierPluginServer) mustEmbedUnimplementedSpecifierPluginServer() {}

// UnsafeSpecifierPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpecifierPluginServer will
// result in compilation errors.
type UnsafeSpecifierPluginServer interface {
	mustEmbedUnimplementedSpecifierPluginServer()
}

func RegisterSpecifierPluginServer(s grpc.ServiceRegistrar, srv SpecifierPluginServer) {
	s.RegisterService(&SpecifierPlugin_ServiceDesc, srv)
}

func _SpecifierPlugin_Specify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Specifier_Specify_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecifierPluginServer).Specify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpluginv1.SpecifierPlugin/Specify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecifierPluginServer).Specify(ctx, req.(*Specifier_Specify_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// SpecifierPlugin_ServiceDesc is the grpc.ServiceDesc for SpecifierPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpecifierPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cpluginv1.SpecifierPlugin",
	HandlerType: (*SpecifierPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Specify",
			Handler:    _SpecifierPlugin_Specify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cproto.proto",
}
