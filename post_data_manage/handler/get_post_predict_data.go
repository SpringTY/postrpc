package handler

import (
	"context"
	"log"
	"post_data_manage/idl/post_data_manage"
	"post_data_manage/model"
	"post_data_manage/mq"
	"post_data_manage/util"
)

type GeneratePostPredictDataHandler struct {
	Request  *post_data_manage.GeneratePostPredictDataRequest
	Response *post_data_manage.GeneratePostPredictDataResponse
	Ctx      context.Context
}

func NewGeneratePostPredictDataHandler(Ctx context.Context, Request *post_data_manage.GeneratePostPredictDataRequest) *GeneratePostPredictDataHandler {
	return &GeneratePostPredictDataHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &post_data_manage.GeneratePostPredictDataResponse{
			Message: "success",
			Status:  0,
			TaskId:  "",
		},
	}
}
func (h *GeneratePostPredictDataHandler) Run() {
	task := h.BuidTask()
	err := h.SendTask(task)
	if err != nil {
		log.Printf("[GeneratePostPredictDataHandler] error call SendTask,err: %v", err)
		h.Response.Status = 1
		h.Response.Message = "error"
		return
	}
	h.Response.TaskId = task.TaskID
}

func (h *GeneratePostPredictDataHandler) CheckParam() error {
	return nil
}
func (h *GeneratePostPredictDataHandler) BuidTask() (task *model.GeneratePredictDataTask) {
	return &model.GeneratePredictDataTask{
		TaskID:  util.GetUUID(),
		TaskTag: h.Request.Tag,
	}
}
func (h *GeneratePostPredictDataHandler) SendTask(task *model.GeneratePredictDataTask) error {
	return mq.SendGeneratePredictDataTask(h.Ctx, task)
}
