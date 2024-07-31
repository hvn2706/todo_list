package task

import (
	"context"
	"todo_list/database"
	"todo_list/internal/model"
)

type IService interface {
	UpsertTask(ctx context.Context, request UpsertTaskRequest) (UpsertTaskResponse, error)

	GetListTasks(ctx context.Context) ([]model.Task, error)
}

type ServiceImpl struct {
	gdb database.DBAdapter
}

var _ IService = &ServiceImpl{}

func InitService() IService {
	return &ServiceImpl{
		gdb: database.GetDBInstance(),
	}
}
