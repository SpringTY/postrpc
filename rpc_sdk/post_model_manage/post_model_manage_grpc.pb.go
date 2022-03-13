// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package post_model_manage

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

// PostModelManageClient is the client API for PostModelManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostModelManageClient interface {
	GetModels(ctx context.Context, in *GetModelsRequest, opts ...grpc.CallOption) (*GetModelsResponse, error)
	RemoveModel(ctx context.Context, in *RemoveModelRequest, opts ...grpc.CallOption) (*RemoveModelResponse, error)
	UpdateModelConfig(ctx context.Context, in *UpdateModelConfigRequest, opts ...grpc.CallOption) (*UpdateModelConfigResponse, error)
	GetModelStates(ctx context.Context, in *GetModelStatesRequest, opts ...grpc.CallOption) (*GetModelStatesResponse, error)
	RegisterModel(ctx context.Context, in *RegisterModelRequest, opts ...grpc.CallOption) (*RegisterModelResponse, error)
}

type postModelManageClient struct {
	cc grpc.ClientConnInterface
}

func NewPostModelManageClient(cc grpc.ClientConnInterface) PostModelManageClient {
	return &postModelManageClient{cc}
}

func (c *postModelManageClient) GetModels(ctx context.Context, in *GetModelsRequest, opts ...grpc.CallOption) (*GetModelsResponse, error) {
	out := new(GetModelsResponse)
	err := c.cc.Invoke(ctx, "/post_model_manage.PostModelManage/GetModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postModelManageClient) RemoveModel(ctx context.Context, in *RemoveModelRequest, opts ...grpc.CallOption) (*RemoveModelResponse, error) {
	out := new(RemoveModelResponse)
	err := c.cc.Invoke(ctx, "/post_model_manage.PostModelManage/RemoveModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postModelManageClient) UpdateModelConfig(ctx context.Context, in *UpdateModelConfigRequest, opts ...grpc.CallOption) (*UpdateModelConfigResponse, error) {
	out := new(UpdateModelConfigResponse)
	err := c.cc.Invoke(ctx, "/post_model_manage.PostModelManage/UpdateModelConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postModelManageClient) GetModelStates(ctx context.Context, in *GetModelStatesRequest, opts ...grpc.CallOption) (*GetModelStatesResponse, error) {
	out := new(GetModelStatesResponse)
	err := c.cc.Invoke(ctx, "/post_model_manage.PostModelManage/GetModelStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postModelManageClient) RegisterModel(ctx context.Context, in *RegisterModelRequest, opts ...grpc.CallOption) (*RegisterModelResponse, error) {
	out := new(RegisterModelResponse)
	err := c.cc.Invoke(ctx, "/post_model_manage.PostModelManage/RegisterModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostModelManageServer is the server API for PostModelManage service.
// All implementations must embed UnimplementedPostModelManageServer
// for forward compatibility
type PostModelManageServer interface {
	GetModels(context.Context, *GetModelsRequest) (*GetModelsResponse, error)
	RemoveModel(context.Context, *RemoveModelRequest) (*RemoveModelResponse, error)
	UpdateModelConfig(context.Context, *UpdateModelConfigRequest) (*UpdateModelConfigResponse, error)
	GetModelStates(context.Context, *GetModelStatesRequest) (*GetModelStatesResponse, error)
	RegisterModel(context.Context, *RegisterModelRequest) (*RegisterModelResponse, error)
	mustEmbedUnimplementedPostModelManageServer()
}

// UnimplementedPostModelManageServer must be embedded to have forward compatible implementations.
type UnimplementedPostModelManageServer struct {
}

func (UnimplementedPostModelManageServer) GetModels(context.Context, *GetModelsRequest) (*GetModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModels not implemented")
}
func (UnimplementedPostModelManageServer) RemoveModel(context.Context, *RemoveModelRequest) (*RemoveModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveModel not implemented")
}
func (UnimplementedPostModelManageServer) UpdateModelConfig(context.Context, *UpdateModelConfigRequest) (*UpdateModelConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateModelConfig not implemented")
}
func (UnimplementedPostModelManageServer) GetModelStates(context.Context, *GetModelStatesRequest) (*GetModelStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelStates not implemented")
}
func (UnimplementedPostModelManageServer) RegisterModel(context.Context, *RegisterModelRequest) (*RegisterModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterModel not implemented")
}
func (UnimplementedPostModelManageServer) mustEmbedUnimplementedPostModelManageServer() {}

// UnsafePostModelManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostModelManageServer will
// result in compilation errors.
type UnsafePostModelManageServer interface {
	mustEmbedUnimplementedPostModelManageServer()
}

func RegisterPostModelManageServer(s grpc.ServiceRegistrar, srv PostModelManageServer) {
	s.RegisterService(&PostModelManage_ServiceDesc, srv)
}

func _PostModelManage_GetModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostModelManageServer).GetModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_model_manage.PostModelManage/GetModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostModelManageServer).GetModels(ctx, req.(*GetModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostModelManage_RemoveModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostModelManageServer).RemoveModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_model_manage.PostModelManage/RemoveModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostModelManageServer).RemoveModel(ctx, req.(*RemoveModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostModelManage_UpdateModelConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateModelConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostModelManageServer).UpdateModelConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_model_manage.PostModelManage/UpdateModelConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostModelManageServer).UpdateModelConfig(ctx, req.(*UpdateModelConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostModelManage_GetModelStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostModelManageServer).GetModelStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_model_manage.PostModelManage/GetModelStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostModelManageServer).GetModelStates(ctx, req.(*GetModelStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostModelManage_RegisterModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostModelManageServer).RegisterModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_model_manage.PostModelManage/RegisterModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostModelManageServer).RegisterModel(ctx, req.(*RegisterModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostModelManage_ServiceDesc is the grpc.ServiceDesc for PostModelManage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostModelManage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post_model_manage.PostModelManage",
	HandlerType: (*PostModelManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModels",
			Handler:    _PostModelManage_GetModels_Handler,
		},
		{
			MethodName: "RemoveModel",
			Handler:    _PostModelManage_RemoveModel_Handler,
		},
		{
			MethodName: "UpdateModelConfig",
			Handler:    _PostModelManage_UpdateModelConfig_Handler,
		},
		{
			MethodName: "GetModelStates",
			Handler:    _PostModelManage_GetModelStates_Handler,
		},
		{
			MethodName: "RegisterModel",
			Handler:    _PostModelManage_RegisterModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post_model_manage.proto",
}
