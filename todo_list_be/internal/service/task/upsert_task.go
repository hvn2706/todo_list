package task

import (
	"context"
	"errors"
	"time"
	"todo_list/internal/model"
	"todo_list/logger"
	"todo_list/pkg/common"
	"todo_list/pkg/constants"
)

func (s *ServiceImpl) UpsertTask(ctx context.Context, request UpsertTaskRequest) (UpsertTaskResponse, error) {
	if request.TaskID == nil {
		return s.insertTask(ctx, request)
	}

	return s.updateTask(ctx, request)
}

func (s *ServiceImpl) insertTask(ctx context.Context, request UpsertTaskRequest) (UpsertTaskResponse, error) {
	task := model.Task{
		Title:    request.Title,
		Subtitle: request.SubTitle,
		Status:   common.GetPointer(constants.TaskStatusOpen),
		Priority: request.Priority,
	}

	err := s.gdb.DB().Table("task").Create(&task).Error
	if err != nil {
		logger.Errorf("insert task failed: %+v", err.Error())
		return UpsertTaskResponse{}, err
	}

	return UpsertTaskResponse{
		TaskID: task.ID,
	}, nil
}

func (s *ServiceImpl) updateTask(ctx context.Context, request UpsertTaskRequest) (UpsertTaskResponse, error) {
	if request.TaskID == nil {
		return UpsertTaskResponse{}, errors.New("task id is nil")
	}

	var dueDateTime *time.Time
	if request.DueDate != nil {
		dueDateTime = common.GetPointer(time.Unix(*request.DueDate, 0).UTC())
	}

	var completedAt *time.Time
	if request.Status != nil && *request.Status == constants.TaskStatusDone {
		completedAt = common.GetPointer(time.Now().UTC())
	}

	var task model.Task
	err := s.gdb.DB().Table("task").Where("id = ?", *request.TaskID).First(&task).Error
	if err != nil {
		logger.Errorf("get task failed: %+v", err.Error())
		return UpsertTaskResponse{}, err
	}

	if task.Status != nil && *task.Status == constants.TaskStatusDone {
		return UpsertTaskResponse{}, errors.New("task is already done")
	}

	if task.Status != nil && request.Status != nil && *task.Status == constants.TaskStatusOpen && *request.Status == constants.TaskStatusDone {
		return UpsertTaskResponse{}, errors.New("task is in open can't be done")
	}

	err = s.gdb.DB().Table("task").Where("id = ?", *request.TaskID).Updates(model.Task{
		Title:       request.Title,
		Subtitle:    request.SubTitle,
		DueDate:     dueDateTime,
		Status:      request.Status,
		CompletedAt: completedAt,
		Priority:    request.Priority,
	}).Error
	if err != nil {
		logger.Errorf("update task failed: %+v", err.Error())
		return UpsertTaskResponse{}, err
	}

	return UpsertTaskResponse{
		TaskID: task.ID,
	}, nil
}
