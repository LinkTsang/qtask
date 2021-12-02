package service

import (
	"qtask/pkg/model"
	"testing"
)

func TestNewTaskDetail(t *testing.T) {
	taskDetail := model.NewTaskDetail("go", "run", "./examples/dummy/dummy.go")
	if taskDetail.Status != model.TaskStatusPending {
		t.Errorf("task.Status = %v; want %v", taskDetail.Status, model.TaskStatusPending)
	}

}
