package task

import (
	"context"
	"todo_list/internal/model"
	"todo_list/logger"
)

func (s *ServiceImpl) GetListTasks(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task
	err := s.gdb.DB().Table("task").Scan(&tasks).Error
	if err != nil {
		logger.Errorf("get list tasks failed: %+v", err.Error())
		return nil, err
	}
	return tasks, nil
}
