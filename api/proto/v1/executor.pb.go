// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: api/proto/v1/executor.proto

package executor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskStatus int32

const (
	TaskStatus_UNKNOWN   TaskStatus = 0
	TaskStatus_READY     TaskStatus = 1
	TaskStatus_RUNNING   TaskStatus = 2
	TaskStatus_PAUSED    TaskStatus = 3
	TaskStatus_PENDING   TaskStatus = 4
	TaskStatus_CANCELED  TaskStatus = 5
	TaskStatus_COMPLETED TaskStatus = 6
	TaskStatus_DETACHED  TaskStatus = 7
	TaskStatus_ERROR     TaskStatus = 8
	TaskStatus_NOT_EXIST TaskStatus = 9
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "READY",
		2: "RUNNING",
		3: "PAUSED",
		4: "PENDING",
		5: "CANCELED",
		6: "COMPLETED",
		7: "DETACHED",
		8: "ERROR",
		9: "NOT_EXIST",
	}
	TaskStatus_value = map[string]int32{
		"UNKNOWN":   0,
		"READY":     1,
		"RUNNING":   2,
		"PAUSED":    3,
		"PENDING":   4,
		"CANCELED":  5,
		"COMPLETED": 6,
		"DETACHED":  7,
		"ERROR":     8,
		"NOT_EXIST": 9,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_v1_executor_proto_enumTypes[0].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_api_proto_v1_executor_proto_enumTypes[0]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_v1_executor_proto_rawDescGZIP(), []int{0}
}

type RunTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId         string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Status         TaskStatus             `protobuf:"varint,2,opt,name=status,proto3,enum=pb.v1.TaskStatus" json:"status,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	StartedAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	PausedAt       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=paused_at,json=pausedAt,proto3" json:"paused_at,omitempty"`
	TerminatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=terminated_at,json=terminatedAt,proto3" json:"terminated_at,omitempty"`
	Name           string                 `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Description    string                 `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	WorkingDir     string                 `protobuf:"bytes,9,opt,name=working_dir,json=workingDir,proto3" json:"working_dir,omitempty"`
	Path           string                 `protobuf:"bytes,10,opt,name=path,proto3" json:"path,omitempty"`
	Args           []string               `protobuf:"bytes,11,rep,name=args,proto3" json:"args,omitempty"`
	OutputFilePath string                 `protobuf:"bytes,12,opt,name=output_file_path,json=outputFilePath,proto3" json:"output_file_path,omitempty"`
	ExitCode       int32                  `protobuf:"varint,13,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	ExitMessage    string                 `protobuf:"bytes,14,opt,name=exit_message,json=exitMessage,proto3" json:"exit_message,omitempty"`
}

func (x *RunTaskRequest) Reset() {
	*x = RunTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_executor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunTaskRequest) ProtoMessage() {}

func (x *RunTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_executor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunTaskRequest.ProtoReflect.Descriptor instead.
func (*RunTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_executor_proto_rawDescGZIP(), []int{0}
}

func (x *RunTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *RunTaskRequest) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_UNKNOWN
}

func (x *RunTaskRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RunTaskRequest) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *RunTaskRequest) GetPausedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PausedAt
	}
	return nil
}

func (x *RunTaskRequest) GetTerminatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.TerminatedAt
	}
	return nil
}

func (x *RunTaskRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RunTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RunTaskRequest) GetWorkingDir() string {
	if x != nil {
		return x.WorkingDir
	}
	return ""
}

func (x *RunTaskRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RunTaskRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *RunTaskRequest) GetOutputFilePath() string {
	if x != nil {
		return x.OutputFilePath
	}
	return ""
}

func (x *RunTaskRequest) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *RunTaskRequest) GetExitMessage() string {
	if x != nil {
		return x.ExitMessage
	}
	return ""
}

type WatchTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId      string     `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Status      TaskStatus `protobuf:"varint,2,opt,name=status,proto3,enum=pb.v1.TaskStatus" json:"status,omitempty"`
	ExitCode    int32      `protobuf:"varint,3,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	ExitMessage string     `protobuf:"bytes,4,opt,name=exit_message,json=exitMessage,proto3" json:"exit_message,omitempty"`
}

func (x *WatchTasksResponse) Reset() {
	*x = WatchTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_executor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchTasksResponse) ProtoMessage() {}

func (x *WatchTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_executor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchTasksResponse.ProtoReflect.Descriptor instead.
func (*WatchTasksResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_executor_proto_rawDescGZIP(), []int{1}
}

func (x *WatchTasksResponse) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *WatchTasksResponse) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_UNKNOWN
}

func (x *WatchTasksResponse) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *WatchTasksResponse) GetExitMessage() string {
	if x != nil {
		return x.ExitMessage
	}
	return ""
}

type KillTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *KillTaskRequest) Reset() {
	*x = KillTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_executor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillTaskRequest) ProtoMessage() {}

func (x *KillTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_executor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillTaskRequest.ProtoReflect.Descriptor instead.
func (*KillTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_executor_proto_rawDescGZIP(), []int{2}
}

func (x *KillTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type StopTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *StopTaskRequest) Reset() {
	*x = StopTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_executor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTaskRequest) ProtoMessage() {}

func (x *StopTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_executor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTaskRequest.ProtoReflect.Descriptor instead.
func (*StopTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_executor_proto_rawDescGZIP(), []int{3}
}

func (x *StopTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type PauseTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *PauseTaskRequest) Reset() {
	*x = PauseTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_executor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PauseTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PauseTaskRequest) ProtoMessage() {}

func (x *PauseTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_executor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PauseTaskRequest.ProtoReflect.Descriptor instead.
func (*PauseTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_executor_proto_rawDescGZIP(), []int{4}
}

func (x *PauseTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type ResumeTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *ResumeTaskRequest) Reset() {
	*x = ResumeTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_executor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeTaskRequest) ProtoMessage() {}

func (x *ResumeTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_executor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeTaskRequest.ProtoReflect.Descriptor instead.
func (*ResumeTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_executor_proto_rawDescGZIP(), []int{5}
}

func (x *ResumeTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

var File_api_proto_v1_executor_proto protoreflect.FileDescriptor

var file_api_proto_v1_executor_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xad, 0x04, 0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x29,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x37, 0x0a, 0x09, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3f, 0x0a, 0x0d, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x12, 0x57, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78,
	0x69, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x65, 0x78, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a,
	0x0f, 0x4b, 0x69, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0f, 0x53, 0x74, 0x6f,
	0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x10, 0x50, 0x61, 0x75, 0x73, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x22, 0x2c, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x2a, 0x8f, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45,
	0x54, 0x41, 0x43, 0x48, 0x45, 0x44, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54,
	0x10, 0x09, 0x32, 0xfd, 0x02, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x12,
	0x38, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x0a, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x19, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x08,
	0x4b, 0x69, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x70,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f,
	0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x09, 0x50, 0x61, 0x75, 0x73, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x75, 0x73, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x3e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x2c, 0x5a, 0x2a, 0x6c, 0x69, 0x6e, 0x6b, 0x74, 0x73, 0x61, 0x6e, 0x67, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x74, 0x61, 0x73, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_executor_proto_rawDescOnce sync.Once
	file_api_proto_v1_executor_proto_rawDescData = file_api_proto_v1_executor_proto_rawDesc
)

func file_api_proto_v1_executor_proto_rawDescGZIP() []byte {
	file_api_proto_v1_executor_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_executor_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_executor_proto_rawDescData)
	})
	return file_api_proto_v1_executor_proto_rawDescData
}

var file_api_proto_v1_executor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_v1_executor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_v1_executor_proto_goTypes = []interface{}{
	(TaskStatus)(0),               // 0: pb.v1.TaskStatus
	(*RunTaskRequest)(nil),        // 1: pb.v1.RunTaskRequest
	(*WatchTasksResponse)(nil),    // 2: pb.v1.WatchTasksResponse
	(*KillTaskRequest)(nil),       // 3: pb.v1.KillTaskRequest
	(*StopTaskRequest)(nil),       // 4: pb.v1.StopTaskRequest
	(*PauseTaskRequest)(nil),      // 5: pb.v1.PauseTaskRequest
	(*ResumeTaskRequest)(nil),     // 6: pb.v1.ResumeTaskRequest
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_api_proto_v1_executor_proto_depIdxs = []int32{
	0,  // 0: pb.v1.RunTaskRequest.status:type_name -> pb.v1.TaskStatus
	7,  // 1: pb.v1.RunTaskRequest.created_at:type_name -> google.protobuf.Timestamp
	7,  // 2: pb.v1.RunTaskRequest.started_at:type_name -> google.protobuf.Timestamp
	7,  // 3: pb.v1.RunTaskRequest.paused_at:type_name -> google.protobuf.Timestamp
	7,  // 4: pb.v1.RunTaskRequest.terminated_at:type_name -> google.protobuf.Timestamp
	0,  // 5: pb.v1.WatchTasksResponse.status:type_name -> pb.v1.TaskStatus
	1,  // 6: pb.v1.Executor.RunTask:input_type -> pb.v1.RunTaskRequest
	8,  // 7: pb.v1.Executor.WatchTasks:input_type -> google.protobuf.Empty
	3,  // 8: pb.v1.Executor.KillTask:input_type -> pb.v1.KillTaskRequest
	4,  // 9: pb.v1.Executor.StopTask:input_type -> pb.v1.StopTaskRequest
	5,  // 10: pb.v1.Executor.PauseTask:input_type -> pb.v1.PauseTaskRequest
	6,  // 11: pb.v1.Executor.ResumeTask:input_type -> pb.v1.ResumeTaskRequest
	8,  // 12: pb.v1.Executor.RunTask:output_type -> google.protobuf.Empty
	2,  // 13: pb.v1.Executor.WatchTasks:output_type -> pb.v1.WatchTasksResponse
	8,  // 14: pb.v1.Executor.KillTask:output_type -> google.protobuf.Empty
	8,  // 15: pb.v1.Executor.StopTask:output_type -> google.protobuf.Empty
	8,  // 16: pb.v1.Executor.PauseTask:output_type -> google.protobuf.Empty
	8,  // 17: pb.v1.Executor.ResumeTask:output_type -> google.protobuf.Empty
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_proto_v1_executor_proto_init() }
func file_api_proto_v1_executor_proto_init() {
	if File_api_proto_v1_executor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_executor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunTaskRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_executor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchTasksResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_executor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillTaskRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_executor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopTaskRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_executor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PauseTaskRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_executor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeTaskRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_v1_executor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_v1_executor_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_executor_proto_depIdxs,
		EnumInfos:         file_api_proto_v1_executor_proto_enumTypes,
		MessageInfos:      file_api_proto_v1_executor_proto_msgTypes,
	}.Build()
	File_api_proto_v1_executor_proto = out.File
	file_api_proto_v1_executor_proto_rawDesc = nil
	file_api_proto_v1_executor_proto_goTypes = nil
	file_api_proto_v1_executor_proto_depIdxs = nil
}
