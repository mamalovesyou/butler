// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package octopus

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

// ConnectorsServiceClient is the client API for ConnectorsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectorsServiceClient interface {
	ListConnectors(ctx context.Context, in *ListConnectorsRequest, opts ...grpc.CallOption) (*ConnectorList, error)
	CreateConnector(ctx context.Context, in *CreateConnectorRequest, opts ...grpc.CallOption) (*Connector, error)
	MutateConnector(ctx context.Context, in *MutateConnectorRequest, opts ...grpc.CallOption) (*Connector, error)
	GetConnector(ctx context.Context, in *GetConnectorRequest, opts ...grpc.CallOption) (*Connector, error)
	AuthenticateOAuthConnector(ctx context.Context, in *AuthenticateConnectorRequest, opts ...grpc.CallOption) (*Connector, error)
	TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error)
}

type connectorsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorsServiceClient(cc grpc.ClientConnInterface) ConnectorsServiceClient {
	return &connectorsServiceClient{cc}
}

func (c *connectorsServiceClient) ListConnectors(ctx context.Context, in *ListConnectorsRequest, opts ...grpc.CallOption) (*ConnectorList, error) {
	out := new(ConnectorList)
	err := c.cc.Invoke(ctx, "/v1.ConnectorsService/ListConnectors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) CreateConnector(ctx context.Context, in *CreateConnectorRequest, opts ...grpc.CallOption) (*Connector, error) {
	out := new(Connector)
	err := c.cc.Invoke(ctx, "/v1.ConnectorsService/CreateConnector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) MutateConnector(ctx context.Context, in *MutateConnectorRequest, opts ...grpc.CallOption) (*Connector, error) {
	out := new(Connector)
	err := c.cc.Invoke(ctx, "/v1.ConnectorsService/MutateConnector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) GetConnector(ctx context.Context, in *GetConnectorRequest, opts ...grpc.CallOption) (*Connector, error) {
	out := new(Connector)
	err := c.cc.Invoke(ctx, "/v1.ConnectorsService/GetConnector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) AuthenticateOAuthConnector(ctx context.Context, in *AuthenticateConnectorRequest, opts ...grpc.CallOption) (*Connector, error) {
	out := new(Connector)
	err := c.cc.Invoke(ctx, "/v1.ConnectorsService/AuthenticateOAuthConnector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error) {
	out := new(TestConnectionResponse)
	err := c.cc.Invoke(ctx, "/v1.ConnectorsService/TestConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorsServiceServer is the server API for ConnectorsService service.
// All implementations must embed UnimplementedConnectorsServiceServer
// for forward compatibility
type ConnectorsServiceServer interface {
	ListConnectors(context.Context, *ListConnectorsRequest) (*ConnectorList, error)
	CreateConnector(context.Context, *CreateConnectorRequest) (*Connector, error)
	MutateConnector(context.Context, *MutateConnectorRequest) (*Connector, error)
	GetConnector(context.Context, *GetConnectorRequest) (*Connector, error)
	AuthenticateOAuthConnector(context.Context, *AuthenticateConnectorRequest) (*Connector, error)
	TestConnection(context.Context, *TestConnectionRequest) (*TestConnectionResponse, error)
	mustEmbedUnimplementedConnectorsServiceServer()
}

// UnimplementedConnectorsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectorsServiceServer struct {
}

func (UnimplementedConnectorsServiceServer) ListConnectors(context.Context, *ListConnectorsRequest) (*ConnectorList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnectors not implemented")
}
func (UnimplementedConnectorsServiceServer) CreateConnector(context.Context, *CreateConnectorRequest) (*Connector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnector not implemented")
}
func (UnimplementedConnectorsServiceServer) MutateConnector(context.Context, *MutateConnectorRequest) (*Connector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateConnector not implemented")
}
func (UnimplementedConnectorsServiceServer) GetConnector(context.Context, *GetConnectorRequest) (*Connector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnector not implemented")
}
func (UnimplementedConnectorsServiceServer) AuthenticateOAuthConnector(context.Context, *AuthenticateConnectorRequest) (*Connector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateOAuthConnector not implemented")
}
func (UnimplementedConnectorsServiceServer) TestConnection(context.Context, *TestConnectionRequest) (*TestConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestConnection not implemented")
}
func (UnimplementedConnectorsServiceServer) mustEmbedUnimplementedConnectorsServiceServer() {}

// UnsafeConnectorsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectorsServiceServer will
// result in compilation errors.
type UnsafeConnectorsServiceServer interface {
	mustEmbedUnimplementedConnectorsServiceServer()
}

func RegisterConnectorsServiceServer(s grpc.ServiceRegistrar, srv ConnectorsServiceServer) {
	s.RegisterService(&ConnectorsService_ServiceDesc, srv)
}

func _ConnectorsService_ListConnectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).ListConnectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ConnectorsService/ListConnectors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).ListConnectors(ctx, req.(*ListConnectorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorsService_CreateConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).CreateConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ConnectorsService/CreateConnector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).CreateConnector(ctx, req.(*CreateConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorsService_MutateConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).MutateConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ConnectorsService/MutateConnector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).MutateConnector(ctx, req.(*MutateConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorsService_GetConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).GetConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ConnectorsService/GetConnector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).GetConnector(ctx, req.(*GetConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorsService_AuthenticateOAuthConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).AuthenticateOAuthConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ConnectorsService/AuthenticateOAuthConnector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).AuthenticateOAuthConnector(ctx, req.(*AuthenticateConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorsService_TestConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).TestConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ConnectorsService/TestConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).TestConnection(ctx, req.(*TestConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectorsService_ServiceDesc is the grpc.ServiceDesc for ConnectorsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectorsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ConnectorsService",
	HandlerType: (*ConnectorsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListConnectors",
			Handler:    _ConnectorsService_ListConnectors_Handler,
		},
		{
			MethodName: "CreateConnector",
			Handler:    _ConnectorsService_CreateConnector_Handler,
		},
		{
			MethodName: "MutateConnector",
			Handler:    _ConnectorsService_MutateConnector_Handler,
		},
		{
			MethodName: "GetConnector",
			Handler:    _ConnectorsService_GetConnector_Handler,
		},
		{
			MethodName: "AuthenticateOAuthConnector",
			Handler:    _ConnectorsService_AuthenticateOAuthConnector_Handler,
		},
		{
			MethodName: "TestConnection",
			Handler:    _ConnectorsService_TestConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/octopus/v1/connectors.proto",
}