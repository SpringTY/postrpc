package main

import (
	"context"
	Handler "post_data_manage/handler"
	"post_data_manage/idl/post_data_manage"
)

type Server struct {
	post_data_manage.UnsafePostDataManageServer
}

func (server *Server) GeneratePostPredictData(ctx context.Context, req *post_data_manage.GeneratePostPredictDataRequest) (*post_data_manage.GeneratePostPredictDataResponse, error) {
	handler := Handler.NewGeneratePostPredictDataHandler(ctx, req)
	handler.Run()
	return handler.Response, nil
}

func (server *Server) GetPostPredictData(ctx context.Context, req *post_data_manage.GetPostPredictDataRequest) (*post_data_manage.GetPostPredictDataResponse, error) {
	handler := Handler.NewGetPostPredictDataHandler(ctx, req)
	handler.Run()
	return handler.Response, nil
}
