package service

import (
	"context"
	"qtask/pkg/model"
)

type ExecutorService interface {
	Service
	RunTask(ctx context.Context, taskDetail *model.TaskDetail, taskDetailUpdated chan model.TaskDetail) error
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

func (s *executorService) Health(ctx context.Context) (bool, error) {
	return true, nil
}

func (s *executorService) RunTask(ctx context.Context, taskDetail *model.TaskDetail, taskDetailUpdated chan model.TaskDetail) error {
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
