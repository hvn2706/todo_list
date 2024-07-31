package task

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
	"todo_list/pkg/constants"
)

type GetListTasksResponse struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	TaskID      int32   `json:"taskId"`
	Title       string  `json:"title"`
	Subtitle    string  `json:"subtitle"`
	DueDate     *int64  `json:"dueDate"`
	Status      *string `json:"status"`
	CompletedAt *int64  `json:"completed"`
	Priority    *int32  `json:"priority"`
}

type UpsertTaskRequest struct {
	TaskID   *int32  `json:"taskId"`
	Title    string  `json:"title"`
	SubTitle string  `json:"subTitle"`
	DueDate  *int64  `json:"dueDate"`
	Status   *string `json:"status"`
	Priority *int32  `json:"priority"`
}

func (r *UpsertTaskRequest) Bind(_ *http.Request) error {
	if r.Status != nil {
		*r.Status = strings.TrimSpace(*r.Status)
		*r.Status = strings.ToUpper(*r.Status)

		if !slices.Contains(constants.TaskStatuses, *r.Status) {
			return fmt.Errorf("invalid status, must be one of %v", constants.TaskStatuses)
		}

		if *r.Status == constants.TaskStatusDoing && r.DueDate == nil {
			return fmt.Errorf("due date is required when status is doing")
		}
	}

	if r.DueDate != nil && *r.DueDate < 0 {
		return fmt.Errorf("invalid due date, must be unix timestamp")
	}

	if r.TaskID != nil && *r.TaskID < 0 {
		return fmt.Errorf("invalid task id")
	}

	return nil
}
