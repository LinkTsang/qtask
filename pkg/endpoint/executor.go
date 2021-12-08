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
		HealthEndpoint  endpoint.Endpoint
		RunTaskEndpoint endpoint.Endpoint
	}

	HealthResponse struct {
		Healthy bool
	}
)

func MakeServerEndpoints(s service.ExecutorService, logger log.Logger) Endpoints {
	var healthEndpoint endpoint.Endpoint
	{
		healthEndpoint = makeHealthEndpoint(s)
		healthEndpoint = LoggingMiddleware(log.With(logger, "method", "Health"))(healthEndpoint)
	}

	var runTaskEndpoint endpoint.Endpoint
	{
		runTaskEndpoint = makeRunTask(s)
		runTaskEndpoint = LoggingMiddleware(log.With(logger, "method", "RunTask"))(runTaskEndpoint)
	}

	return Endpoints{
		HealthEndpoint:  healthEndpoint,
		RunTaskEndpoint: runTaskEndpoint,
	}
}

func makeHealthEndpoint(s service.ExecutorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		healthy := s.Health()
		return HealthResponse{Healthy: healthy}, nil
	}
}

func makeRunTask(s service.ExecutorService) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		proto := request.(*pb.RunTaskRequest)
		taskDetail := model.TaskDetailFromProto(proto)
		taskDetailUpdated := make(chan model.TaskDetail)
		err := s.RunTask(context, taskDetail, taskDetailUpdated)
		return nil, err
	}
}
