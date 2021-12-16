package service

import (
	"context"
	"qtask/pkg/model"
)

type ExecutorService interface {
	Service
	RunTask(ctx context.Context, taskDetail *model.TaskDetail, taskDetailUpdated chan model.TaskDetail) error
	KillTask(id model.TaskId) error
	StopTask(id model.TaskId) error
	PauseTask(id model.TaskId) error
	ResumeTask(id model.TaskId) error
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

func (s *executorService) KillTask(id model.TaskId) error {
	panic("implement me")
}

func (s *executorService) StopTask(id model.TaskId) error {
	panic("implement me")
}

func (s *executorService) PauseTask(id model.TaskId) error {
	panic("implement me")
}

func (s *executorService) ResumeTask(id model.TaskId) error {
	panic("implement me")
}
