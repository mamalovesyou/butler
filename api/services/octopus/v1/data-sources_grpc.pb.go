// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package octopus

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

// DataSourcesServiceClient is the client API for DataSourcesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataSourcesServiceClient interface {
	ListAvailableSources(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DataSourceList, error)
}

type dataSourcesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataSourcesServiceClient(cc grpc.ClientConnInterface) DataSourcesServiceClient {
	return &dataSourcesServiceClient{cc}
}

func (c *dataSourcesServiceClient) ListAvailableSources(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DataSourceList, error) {
	out := new(DataSourceList)
	err := c.cc.Invoke(ctx, "/v1.DataSourcesService/ListAvailableSources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataSourcesServiceServer is the server API for DataSourcesService service.
// All implementations must embed UnimplementedDataSourcesServiceServer
// for forward compatibility
type DataSourcesServiceServer interface {
	ListAvailableSources(context.Context, *emptypb.Empty) (*DataSourceList, error)
	mustEmbedUnimplementedDataSourcesServiceServer()
}

// UnimplementedDataSourcesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataSourcesServiceServer struct {
}

func (UnimplementedDataSourcesServiceServer) ListAvailableSources(context.Context, *emptypb.Empty) (*DataSourceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAvailableSources not implemented")
}
func (UnimplementedDataSourcesServiceServer) mustEmbedUnimplementedDataSourcesServiceServer() {}

// UnsafeDataSourcesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataSourcesServiceServer will
// result in compilation errors.
type UnsafeDataSourcesServiceServer interface {
	mustEmbedUnimplementedDataSourcesServiceServer()
}

func RegisterDataSourcesServiceServer(s grpc.ServiceRegistrar, srv DataSourcesServiceServer) {
	s.RegisterService(&DataSourcesService_ServiceDesc, srv)
}

func _DataSourcesService_ListAvailableSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourcesServiceServer).ListAvailableSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DataSourcesService/ListAvailableSources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourcesServiceServer).ListAvailableSources(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DataSourcesService_ServiceDesc is the grpc.ServiceDesc for DataSourcesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataSourcesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DataSourcesService",
	HandlerType: (*DataSourcesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAvailableSources",
			Handler:    _DataSourcesService_ListAvailableSources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/octopus/v1/data-sources.proto",
}
