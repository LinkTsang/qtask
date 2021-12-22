package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	pb "qtask/api/proto/v1"
	"qtask/pkg/model"
	"qtask/pkg/service"
)

type (
	Endpoints struct {
		RunTaskEndpoint endpoint.Endpoint
	}
)

func MakeServerEndpoints(s service.ExecutorService, logger log.Logger) Endpoints {
	var runTaskEndpoint endpoint.Endpoint
	{
		runTaskEndpoint = MakeRunTask(s)
		runTaskEndpoint = LoggingMiddleware(log.With(logger, "method", "RunTask"))(runTaskEndpoint)
	}

	return Endpoints{
		RunTaskEndpoint: runTaskEndpoint,
	}
}

func MakeRunTask(s service.ExecutorService) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		proto := request.(*pb.RunTaskRequest)
		taskDetail := model.TaskDetailFromProto(proto)
		err := s.RunTask(context, taskDetail)
		return nil, err
	}
}

func (e *Endpoints) RunTask(ctx context.Context, taskDetail *model.TaskDetail) error {
	panic("implement me")
}

func (e *Endpoints) WatchTasks(ctx context.Context) (service.TaskMonitor, error) {
	panic("implement me")
}

func (e *Endpoints) KillTask(ctx context.Context, id model.TaskId) error {
	panic("implement me")
}

func (e *Endpoints) StopTask(ctx context.Context, id model.TaskId) error {
	panic("implement me")
}

func (e *Endpoints) PauseTask(ctx context.Context, id model.TaskId) error {
	panic("implement me")
}

func (e *Endpoints) ResumeTask(ctx context.Context, id model.TaskId) error {
	panic("implement me")
}
