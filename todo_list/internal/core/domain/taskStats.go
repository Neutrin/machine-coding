package domain

type TaskStats struct {
	CompletedCount int64
	ActiveCount    int64
	SpilledCount   int64
}

func NewTaskStats() *TaskStats {
	return &TaskStats{}
}
