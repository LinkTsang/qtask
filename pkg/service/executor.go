package service

import (
	"context"
	"qtask/pkg/model"
)

type TaskMonitor struct {
	ch chan struct {
		taskId  model.TaskId
		status  model.TaskStatus
		message string
	}
	err chan error
}

type ExecutorService interface {
	RunTask(ctx context.Context, taskDetail *model.TaskDetail, taskDetailUpdated chan model.TaskDetail) error
	WatchTasks(ctx context.Context) (TaskMonitor, error)
	KillTask(ctx context.Context, id model.TaskId) error
	StopTask(ctx context.Context, id model.TaskId) error
	PauseTask(ctx context.Context, id model.TaskId) error
	ResumeTask(ctx context.Context, id model.TaskId) error
}

type executorService struct {
}

func NewExecutorService() ExecutorService {
	return &executorService{}
}

func (s *executorService) RunTask(ctx context.Context, taskDetail *model.TaskDetail, taskDetailUpdated chan model.TaskDetail) error {
	panic("implement me")
}

func (s *executorService) WatchTasks(ctx context.Context) (TaskMonitor, error) {
	panic("implement me")
}

func (s *executorService) KillTask(ctx context.Context, id model.TaskId) error {
	panic("implement me")
}

func (s *executorService) StopTask(ctx context.Context, id model.TaskId) error {
	panic("implement me")
}

func (s *executorService) PauseTask(ctx context.Context, id model.TaskId) error {
	panic("implement me")
}

func (s *executorService) ResumeTask(ctx context.Context, id model.TaskId) error {
	panic("implement me")
}
