package task

type UpsertTaskRequest struct {
	TaskID   *int32
	Title    string
	SubTitle string
	DueDate  *int64
	Status   *string
	Priority *int32
}

type UpsertTaskResponse struct {
	TaskID int32 `json:"taskId"`
}
