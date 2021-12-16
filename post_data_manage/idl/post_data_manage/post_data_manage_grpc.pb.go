// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package post_data_manage

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

// PostDataManageClient is the client API for PostDataManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostDataManageClient interface {
	GeneratePostPredictData(ctx context.Context, in *GeneratePostPredictDataRequest, opts ...grpc.CallOption) (*GeneratePostPredictDataResponse, error)
	GetPostPredictData(ctx context.Context, in *GetPostPredictDataRequest, opts ...grpc.CallOption) (*GetPostPredictDataResponse, error)
}

type postDataManageClient struct {
	cc grpc.ClientConnInterface
}

func NewPostDataManageClient(cc grpc.ClientConnInterface) PostDataManageClient {
	return &postDataManageClient{cc}
}

func (c *postDataManageClient) GeneratePostPredictData(ctx context.Context, in *GeneratePostPredictDataRequest, opts ...grpc.CallOption) (*GeneratePostPredictDataResponse, error) {
	out := new(GeneratePostPredictDataResponse)
	err := c.cc.Invoke(ctx, "/post_data_manage.PostDataManage/GeneratePostPredictData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postDataManageClient) GetPostPredictData(ctx context.Context, in *GetPostPredictDataRequest, opts ...grpc.CallOption) (*GetPostPredictDataResponse, error) {
	out := new(GetPostPredictDataResponse)
	err := c.cc.Invoke(ctx, "/post_data_manage.PostDataManage/GetPostPredictData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostDataManageServer is the server API for PostDataManage service.
// All implementations must embed UnimplementedPostDataManageServer
// for forward compatibility
type PostDataManageServer interface {
	GeneratePostPredictData(context.Context, *GeneratePostPredictDataRequest) (*GeneratePostPredictDataResponse, error)
	GetPostPredictData(context.Context, *GetPostPredictDataRequest) (*GetPostPredictDataResponse, error)
	mustEmbedUnimplementedPostDataManageServer()
}

// UnimplementedPostDataManageServer must be embedded to have forward compatible implementations.
type UnimplementedPostDataManageServer struct {
}

func (UnimplementedPostDataManageServer) GeneratePostPredictData(context.Context, *GeneratePostPredictDataRequest) (*GeneratePostPredictDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePostPredictData not implemented")
}
func (UnimplementedPostDataManageServer) GetPostPredictData(context.Context, *GetPostPredictDataRequest) (*GetPostPredictDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostPredictData not implemented")
}
func (UnimplementedPostDataManageServer) mustEmbedUnimplementedPostDataManageServer() {}

// UnsafePostDataManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostDataManageServer will
// result in compilation errors.
type UnsafePostDataManageServer interface {
	mustEmbedUnimplementedPostDataManageServer()
}

func RegisterPostDataManageServer(s grpc.ServiceRegistrar, srv PostDataManageServer) {
	s.RegisterService(&PostDataManage_ServiceDesc, srv)
}

func _PostDataManage_GeneratePostPredictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratePostPredictDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostDataManageServer).GeneratePostPredictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_data_manage.PostDataManage/GeneratePostPredictData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostDataManageServer).GeneratePostPredictData(ctx, req.(*GeneratePostPredictDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostDataManage_GetPostPredictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostPredictDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostDataManageServer).GetPostPredictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_data_manage.PostDataManage/GetPostPredictData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostDataManageServer).GetPostPredictData(ctx, req.(*GetPostPredictDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostDataManage_ServiceDesc is the grpc.ServiceDesc for PostDataManage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostDataManage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post_data_manage.PostDataManage",
	HandlerType: (*PostDataManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeneratePostPredictData",
			Handler:    _PostDataManage_GeneratePostPredictData_Handler,
		},
		{
			MethodName: "GetPostPredictData",
			Handler:    _PostDataManage_GetPostPredictData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post_data_manage.proto",
}
