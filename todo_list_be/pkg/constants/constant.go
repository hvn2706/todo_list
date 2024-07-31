package constants

const (
	Yml          = "yml"
	ConstConfig  = "config"
	RootPath     = "."
	HeaderAPIKey = "X-Api-Key"

	TaskStatusOpen  = "OPEN"
	TaskStatusDone  = "DONE"
	TaskStatusDoing = "DOING"
)

var TaskStatuses = []string{TaskStatusOpen, TaskStatusDone, TaskStatusDoing}
