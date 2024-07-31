package task

import (
	"context"
	"github.com/go-chi/render"
	"net/http"
	"todo_list/internal/service/task"
	"todo_list/logger"
	"todo_list/pkg/common"
	response "todo_list/server/common"
)

type (
	IServer interface {
		UpsertTaskAPI() func(http.ResponseWriter, *http.Request)

		GetListTasksAPI() func(http.ResponseWriter, *http.Request)
	}

	ServerImpl struct {
		service task.IService
	}
)

var _ IServer = &ServerImpl{}

func InitServer() IServer {
	service := task.InitService()

	return &ServerImpl{
		service: service,
	}
}

func (s *ServerImpl) GetListTasksAPI() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := s.GetListTasks(r.Context())
		if err != nil {
			logger.Errorf("get list tasks failed: %+v", err.Error())
			_ = render.Render(w, r, &response.BaseResponse[any]{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		err = render.Render(w, r, res)
		if err != nil {
			logger.Errorf("render failed: %+v", err.Error())
			_ = render.Render(w, r, &response.BaseResponse[any]{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
	}
}

func (s *ServerImpl) GetListTasks(ctx context.Context) (*response.BaseResponse[GetListTasksResponse], error) {
	logger.Infof("===== GetListTasks")

	tasks, err := s.service.GetListTasks(ctx)
	if err != nil {
		logger.Errorf("get list tasks failed: %+v", err.Error())
		return nil, err
	}

	tasksResp := make([]Task, len(tasks))
	for i, taskDB := range tasks {
		// 1. Convert time.Time to unix timestamp
		var dueDate *int64
		if taskDB.DueDate != nil {
			dueDate = common.GetPointer(taskDB.DueDate.UTC().Unix())
		}
		var completedAt *int64
		if taskDB.CompletedAt != nil {
			completedAt = common.GetPointer(taskDB.CompletedAt.UTC().Unix())
		}

		tasksResp[i] = Task{
			TaskID:      taskDB.ID,
			Title:       taskDB.Title,
			Subtitle:    taskDB.Subtitle,
			DueDate:     dueDate,
			Status:      taskDB.Status,
			CompletedAt: completedAt,
			Priority:    taskDB.Priority,
		}
	}

	logger.Infof("===== GetListTasks resp: %+v", common.LogStruct(tasksResp))

	return &response.BaseResponse[GetListTasksResponse]{
		Code:    http.StatusOK,
		Message: "success",
		Data: &GetListTasksResponse{
			Tasks: tasksResp,
		},
	}, nil
}

func (s *ServerImpl) UpsertTaskAPI() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Bind DTO
		data := &UpsertTaskRequest{}
		if err := render.Bind(r, data); err != nil {
			logger.Errorf("===== UpsertTaskAPI Bind DTO failed: %+v", err.Error())
			_ = render.Render(w, r, &response.BaseResponse[any]{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		logger.Infof("===== UpsertTaskAPI resp: %+v", common.LogStruct(data))

		// 2. Handle request and map response
		res, err := s.UpsertTask(r.Context(), *data)
		if err != nil {
			logger.Errorf("===== UpsertTaskAPI failed: %+v", err.Error())
			_ = render.Render(w, r, &response.BaseResponse[any]{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
		logger.Infof("===== UpsertTaskAPI ok: %+v", res)

		// 3. Render
		err = render.Render(w, r, res)
		if err != nil {
			_ = render.Render(w, r, &response.BaseResponse[any]{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			logger.Errorf("===== UpsertTaskAPI Render failed: %+v", err.Error())
			return
		}
	}
}

func (s *ServerImpl) UpsertTask(ctx context.Context, request UpsertTaskRequest) (*response.BaseResponse[task.UpsertTaskResponse], error) {
	res, err := s.service.UpsertTask(ctx, task.UpsertTaskRequest{
		TaskID:   request.TaskID,
		Title:    request.Title,
		SubTitle: request.SubTitle,
		DueDate:  request.DueDate,
		Status:   request.Status,
		Priority: request.Priority,
	})
	if err != nil {
		return &response.BaseResponse[task.UpsertTaskResponse]{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, err
	}
	return &response.BaseResponse[task.UpsertTaskResponse]{
		Code:    http.StatusOK,
		Message: "Success",
		Data: &task.UpsertTaskResponse{
			TaskID: res.TaskID,
		},
	}, nil
}
