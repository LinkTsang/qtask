package model

import (
	"github.com/google/uuid"
	pb "qtask/api/proto/v1"
	"time"
)

type TaskId string

type TaskStatus string

const (
	TaskStatusReady     TaskStatus = "READY"
	TaskStatusRunning   TaskStatus = "RUNNING"
	TaskStatusPaused    TaskStatus = "PAUSED"
	TaskStatusPending   TaskStatus = "PENDING"
	TaskStatusCanceled  TaskStatus = "CANCELED"
	TaskStatusCompleted TaskStatus = "COMPLETED"
	TaskStatusDetached  TaskStatus = "DETACHED"
	TaskStatusError     TaskStatus = "ERROR"
	TaskStatusNotExist  TaskStatus = "NOT_EXIST"
)

type TaskDetail struct {
	Id     TaskId
	Status TaskStatus

	CreatedAt    time.Time
	StartedAt    *time.Time
	PausedAt     *time.Time
	TerminatedAt *time.Time

	Name           string
	Description    string
	WorkingDir     string
	Path           string
	Args           []string
	OutputFilePath string
	ExitCode       int32
	ExitMessage    string
}

func NewTaskDetail(path string, args ...string) *TaskDetail {
	id := TaskId(uuid.New().String())
	createdAt := time.Now()
	return &TaskDetail{
		Id:     id,
		Status: TaskStatusPending,

		CreatedAt:    createdAt,
		StartedAt:    nil,
		PausedAt:     nil,
		TerminatedAt: nil,

		Name:           "unnamed",
		Description:    "",
		WorkingDir:     ".",
		Path:           path,
		Args:           args,
		OutputFilePath: "output.log",
		ExitCode:       0,
		ExitMessage:    "",
	}
}

func TaskDetailFromProto(proto *pb.TaskDetail) *TaskDetail {
	var startedAt *time.Time = nil
	var pausedAt *time.Time = nil
	var terminatedAt *time.Time = nil
	if proto.StartedAt != nil {
		t := proto.StartedAt.AsTime()
		startedAt = &t
	}
	if proto.PausedAt != nil {
		t := proto.PausedAt.AsTime()
		pausedAt = &t
	}
	if proto.TerminatedAt != nil {
		t := proto.TerminatedAt.AsTime()
		terminatedAt = &t
	}
	taskDetail := TaskDetail{
		Id:             TaskId(proto.TaskId),
		Status:         TaskStatus(proto.Status),
		CreatedAt:      proto.CreatedAt.AsTime(),
		StartedAt:      startedAt,
		PausedAt:       pausedAt,
		TerminatedAt:   terminatedAt,
		Name:           proto.Name,
		Description:    proto.Description,
		WorkingDir:     proto.WorkingDir,
		Path:           proto.Path,
		Args:           proto.Args,
		OutputFilePath: proto.OutputFilePath,
		ExitCode:       proto.ExitCode,
		ExitMessage:    proto.ExitMessage,
	}
	return &taskDetail
}
