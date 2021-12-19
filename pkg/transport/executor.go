package transport

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "qtask/api/proto/v1"
	qtaskEndpoint "qtask/pkg/endpoint"
	"qtask/pkg/model"
	"qtask/pkg/service"
)

type executorServer struct {
	grpc_health_v1.UnimplementedHealthServer
	pb.UnimplementedExecutorServer
	runTask grpctransport.Handler
}

func NewGRPCServer(endpoints qtaskEndpoint.Endpoints, logger log.Logger) pb.ExecutorServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	return &executorServer{
		runTask: grpctransport.NewServer(endpoints.RunTaskEndpoint,
			decodeGRPCRunTaskRequest,
			encodeGRPCRunTaskResponse,
			options...),
	}
}

func decodeNullableProtoTime(pb *timestamppb.Timestamp) *time.Time {
	if pb != nil {
		t := pb.AsTime()
		return &t
	}
	return nil
}

func decodeGRPCRunTaskRequest(_ context.Context, grpcRequest interface{}) (request interface{}, err error) {
	pbTaskDetail := grpcRequest.(*pb.RunTaskRequest)
	if pbTaskDetail.CreatedAt == nil {
		return nil, status.Errorf(codes.Internal, "CreatedAt is nil")
	}
	taskDetail := model.TaskDetail{
		Id:             model.TaskId(pbTaskDetail.TaskId),
		Status:         model.TaskStatus(pbTaskDetail.Status),
		CreatedAt:      pbTaskDetail.CreatedAt.AsTime(),
		StartedAt:      decodeNullableProtoTime(pbTaskDetail.StartedAt),
		PausedAt:       decodeNullableProtoTime(pbTaskDetail.PausedAt),
		TerminatedAt:   decodeNullableProtoTime(pbTaskDetail.TerminatedAt),
		Name:           pbTaskDetail.Name,
		Description:    pbTaskDetail.Description,
		WorkingDir:     pbTaskDetail.WorkingDir,
		Path:           pbTaskDetail.Path,
		Args:           pbTaskDetail.Args,
		OutputFilePath: pbTaskDetail.OutputFilePath,
		ExitCode:       pbTaskDetail.ExitCode,
		ExitMessage:    pbTaskDetail.ExitMessage,
	}
	return taskDetail, nil
}

func encodeTime(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func encodeNullableTime(t *time.Time) *timestamppb.Timestamp {
	if t != nil {
		return timestamppb.New(*t)
	}
	return nil
}

func encodeGRPCRunTaskResponse(ctx context.Context, response interface{}) (grpcResponse interface{}, err error) {
	res := response.(*struct{ taskDetailUpdated chan *model.TaskDetail })
	grpcTaskDetailUpdated := make(chan *pb.RunTaskResponse)

	go func() {
		for t := range res.taskDetailUpdated {
			grpcTaskDetailUpdated <- &pb.RunTaskResponse{
				TaskId:      string(t.Id),
				Status:      pb.TaskStatus(pb.TaskStatus_value[string(t.Status)]),
				ExitCode:    t.ExitCode,
				ExitMessage: t.ExitMessage,
			}
		}
	}()
	return struct{ taskDetailUpdated chan *pb.RunTaskResponse }{taskDetailUpdated: grpcTaskDetailUpdated}, nil
}

func (s *executorServer) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (s *executorServer) RunTask(taskDetail *pb.RunTaskRequest, stream pb.Executor_RunTaskServer) error {
	ctx := stream.Context()
	_, response, err := s.runTask.ServeGRPC(ctx, taskDetail)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	res := response.(*struct{ taskDetailUpdated chan *pb.RunTaskResponse })
	for taskDetailUpdated := range res.taskDetailUpdated {
		err := stream.Send(taskDetailUpdated)
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

func (executorServer) KillTask(context.Context, *pb.KillTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KillTask not implemented")
}

func (executorServer) StopTask(context.Context, *pb.StopTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopTask not implemented")
}
func (executorServer) PauseTask(context.Context, *pb.PauseTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseTask not implemented")
}
func (executorServer) ResumeTask(context.Context, *pb.ResumeTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeTask not implemented")
}

func NewGRPCClient(conn *grpc.ClientConn, options []grpctransport.ClientOption) (service.ExecutorService, error) {
	serviceName := pb.Executor_ServiceDesc.ServiceName

	var healthEndpoint endpoint.Endpoint
	{
		healthEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"health",
			encodeGRPCHealthCheckRequest,
			decodeGRPCHealthCheckResponse,
			grpc_health_v1.HealthCheckResponse{}, options...).Endpoint()
	}

	var runTaskEndpoint endpoint.Endpoint
	{
		runTaskEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"runTask",
			encodeGRPCRunTaskRequest,
			decodeGRPCRunTaskResponse,
			pb.RunTaskResponse{}, options...).Endpoint()
	}

	return &qtaskEndpoint.Endpoints{
		HealthEndpoint:  healthEndpoint,
		RunTaskEndpoint: runTaskEndpoint,
	}, nil
}

func encodeGRPCHealthCheckRequest(_ context.Context, _ interface{}) (grpcRequest interface{}, err error) {
	return nil, nil
}

func decodeGRPCHealthCheckResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*grpc_health_v1.HealthCheckResponse)
	return qtaskEndpoint.HealthCheckResponse{
		Healthy: reply.Status == grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func encodeGRPCRunTaskRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*model.TaskDetail)
	return &pb.RunTaskRequest{
		TaskId:         string(req.Id),
		Status:         pb.TaskStatus(pb.TaskStatus_value[string(req.Status)]),
		CreatedAt:      encodeTime(req.CreatedAt),
		StartedAt:      encodeNullableTime(req.StartedAt),
		PausedAt:       encodeNullableTime(req.PausedAt),
		TerminatedAt:   encodeNullableTime(req.TerminatedAt),
		Name:           req.Name,
		Description:    req.Description,
		WorkingDir:     req.WorkingDir,
		Path:           req.Path,
		Args:           req.Args,
		OutputFilePath: req.OutputFilePath,
		ExitCode:       req.ExitCode,
		ExitMessage:    req.ExitMessage,
	}, nil
}

type RunTaskResponse struct {
	taskId      model.TaskId
	Status      model.TaskStatus
	ExitCode    int32
	ExitMessage string
}

func decodeGRPCRunTaskResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.RunTaskResponse)
	return RunTaskResponse{
		taskId:      model.TaskId(reply.TaskId),
		Status:      model.TaskStatus(reply.Status.String()),
		ExitCode:    reply.ExitCode,
		ExitMessage: reply.ExitMessage,
	}, nil
}
