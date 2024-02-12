// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: trustmatrix.proto

package trustmatrixpb

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
	Service_Create_FullMethodName = "/trustmatrix.Service/Create"
	Service_Get_FullMethodName    = "/trustmatrix.Service/Get"
	Service_Update_FullMethodName = "/trustmatrix.Service/Update"
	Service_Flush_FullMethodName  = "/trustmatrix.Service/Flush"
	Service_Delete_FullMethodName = "/trustmatrix.Service/Delete"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// Create a new trust matrix (for local trust), return its ID.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get the full contents of a trust matrix.
	// The header is sent first (includes metadata such as timestamp),
	// followed by non-zero entries.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (Service_GetClient, error)
	// Update a trust matrix.
	// Entries for the same matrix with the same timestamp can be batched
	// in the same request.
	// Updates cannot go back in time:
	// An update older than the target matrix is rejected.
	// Updates can be empty (have zero entries):
	// This can be used to force-update the timestamp of the trust matrix
	// and trigger periodic re-compute of a compute job.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Flush a trust matrix, i.e. remove (zero out) all its entries.
	Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error)
	// Delete a trust matrix altogether.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Service_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (Service_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[0], Service_Get_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_GetClient interface {
	Recv() (*GetResponse, error)
	grpc.ClientStream
}

type serviceGetClient struct {
	grpc.ClientStream
}

func (x *serviceGetClient) Recv() (*GetResponse, error) {
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, Service_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error) {
	out := new(FlushResponse)
	err := c.cc.Invoke(ctx, Service_Flush_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, Service_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// Create a new trust matrix (for local trust), return its ID.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get the full contents of a trust matrix.
	// The header is sent first (includes metadata such as timestamp),
	// followed by non-zero entries.
	Get(*GetRequest, Service_GetServer) error
	// Update a trust matrix.
	// Entries for the same matrix with the same timestamp can be batched
	// in the same request.
	// Updates cannot go back in time:
	// An update older than the target matrix is rejected.
	// Updates can be empty (have zero entries):
	// This can be used to force-update the timestamp of the trust matrix
	// and trigger periodic re-compute of a compute job.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Flush a trust matrix, i.e. remove (zero out) all its entries.
	Flush(context.Context, *FlushRequest) (*FlushResponse, error)
	// Delete a trust matrix altogether.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedServiceServer) Get(*GetRequest, Service_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedServiceServer) Flush(context.Context, *FlushRequest) (*FlushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
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

func _Service_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).Get(m, &serviceGetServer{stream})
}

type Service_GetServer interface {
	Send(*GetResponse) error
	grpc.ServerStream
}

type serviceGetServer struct {
	grpc.ServerStream
}

func (x *serviceGetServer) Send(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_Flush_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Flush(ctx, req.(*FlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trustmatrix.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Service_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Service_Update_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _Service_Flush_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Service_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _Service_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "trustmatrix.proto",
}
