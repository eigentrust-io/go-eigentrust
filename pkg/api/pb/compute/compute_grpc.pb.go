// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: compute.proto

package computepb

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

const (
	Service_BasicCompute_FullMethodName = "/compute.Service/BasicCompute"
	Service_CreateJob_FullMethodName    = "/compute.Service/CreateJob"
	Service_DeleteJob_FullMethodName    = "/compute.Service/DeleteJob"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// Perform a basic EigenTrust compute.
	BasicCompute(ctx context.Context, in *BasicComputeRequest, opts ...grpc.CallOption) (*BasicComputeResponse, error)
	// Create a compute job.
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error)
	// Delete/decommission a compute job.
	DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...grpc.CallOption) (*DeleteJobResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) BasicCompute(ctx context.Context, in *BasicComputeRequest, opts ...grpc.CallOption) (*BasicComputeResponse, error) {
	out := new(BasicComputeResponse)
	err := c.cc.Invoke(ctx, Service_BasicCompute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error) {
	out := new(CreateJobResponse)
	err := c.cc.Invoke(ctx, Service_CreateJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...grpc.CallOption) (*DeleteJobResponse, error) {
	out := new(DeleteJobResponse)
	err := c.cc.Invoke(ctx, Service_DeleteJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// Perform a basic EigenTrust compute.
	BasicCompute(context.Context, *BasicComputeRequest) (*BasicComputeResponse, error)
	// Create a compute job.
	CreateJob(context.Context, *CreateJobRequest) (*CreateJobResponse, error)
	// Delete/decommission a compute job.
	DeleteJob(context.Context, *DeleteJobRequest) (*DeleteJobResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) BasicCompute(context.Context, *BasicComputeRequest) (*BasicComputeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BasicCompute not implemented")
}
func (UnimplementedServiceServer) CreateJob(context.Context, *CreateJobRequest) (*CreateJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (UnimplementedServiceServer) DeleteJob(context.Context, *DeleteJobRequest) (*DeleteJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJob not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_BasicCompute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicComputeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).BasicCompute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_BasicCompute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).BasicCompute(ctx, req.(*BasicComputeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateJob(ctx, req.(*CreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteJob(ctx, req.(*DeleteJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "compute.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BasicCompute",
			Handler:    _Service_BasicCompute_Handler,
		},
		{
			MethodName: "CreateJob",
			Handler:    _Service_CreateJob_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _Service_DeleteJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "compute.proto",
}
