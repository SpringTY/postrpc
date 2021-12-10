package main

import post_model_manage "post_model_manage/idl"
import "context"
// server is used to implement helloworld.GreeterServer.
type server struct {
	post_model_manage.UnsafePostModelManageServer
}

func (s* server) SayHello(ctx context.Context, request *post_model_manage.HelloRequest) (*post_model_manage.HelloReply, error) {
	input := request.Hello
	resp := &post_model_manage.HelloReply{
		Message: input,
	}
	return resp,nil
}

