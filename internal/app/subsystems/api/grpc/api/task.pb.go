// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: internal/app/subsystems/api/grpc/api/task.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClaimTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProcessId string `protobuf:"bytes,2,opt,name=processId,proto3" json:"processId,omitempty"`
	Counter   int32  `protobuf:"varint,3,opt,name=counter,proto3" json:"counter,omitempty"`
	Frequency int32  `protobuf:"varint,4,opt,name=frequency,proto3" json:"frequency,omitempty"`
	RequestId string `protobuf:"bytes,5,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *ClaimTaskRequest) Reset() {
	*x = ClaimTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimTaskRequest) ProtoMessage() {}

func (x *ClaimTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimTaskRequest.ProtoReflect.Descriptor instead.
func (*ClaimTaskRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_api_task_proto_rawDescGZIP(), []int{0}
}

func (x *ClaimTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClaimTaskRequest) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

func (x *ClaimTaskRequest) GetCounter() int32 {
	if x != nil {
		return x.Counter
	}
	return 0
}

func (x *ClaimTaskRequest) GetFrequency() int32 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *ClaimTaskRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type ClaimTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ClaimTaskResponse) Reset() {
	*x = ClaimTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimTaskResponse) ProtoMessage() {}

func (x *ClaimTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimTaskResponse.ProtoReflect.Descriptor instead.
func (*ClaimTaskResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_api_task_proto_rawDescGZIP(), []int{1}
}

func (x *ClaimTaskResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CompleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Counter   int32  `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
	RequestId string `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *CompleteTaskRequest) Reset() {
	*x = CompleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTaskRequest) ProtoMessage() {}

func (x *CompleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTaskRequest.ProtoReflect.Descriptor instead.
func (*CompleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_api_task_proto_rawDescGZIP(), []int{2}
}

func (x *CompleteTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompleteTaskRequest) GetCounter() int32 {
	if x != nil {
		return x.Counter
	}
	return 0
}

func (x *CompleteTaskRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type CompleteTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompleteTaskResponse) Reset() {
	*x = CompleteTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTaskResponse) ProtoMessage() {}

func (x *CompleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTaskResponse.ProtoReflect.Descriptor instead.
func (*CompleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_api_task_proto_rawDescGZIP(), []int{3}
}

type HeartbeatTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessId string `protobuf:"bytes,1,opt,name=processId,proto3" json:"processId,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *HeartbeatTasksRequest) Reset() {
	*x = HeartbeatTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatTasksRequest) ProtoMessage() {}

func (x *HeartbeatTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatTasksRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatTasksRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_api_task_proto_rawDescGZIP(), []int{4}
}

func (x *HeartbeatTasksRequest) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

func (x *HeartbeatTasksRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type HeartbeatTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TasksAffected int32 `protobuf:"varint,1,opt,name=tasksAffected,proto3" json:"tasksAffected,omitempty"`
}

func (x *HeartbeatTasksResponse) Reset() {
	*x = HeartbeatTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatTasksResponse) ProtoMessage() {}

func (x *HeartbeatTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatTasksResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatTasksResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_api_task_proto_rawDescGZIP(), []int{5}
}

func (x *HeartbeatTasksResponse) GetTasksAffected() int32 {
	if x != nil {
		return x.TasksAffected
	}
	return 0
}

var File_internal_app_subsystems_api_grpc_api_task_proto protoreflect.FileDescriptor

var file_internal_app_subsystems_api_grpc_api_task_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x96, 0x01, 0x0a, 0x10, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x22, 0x27, 0x0a, 0x11, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5d, 0x0a, 0x13, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x53, 0x0a, 0x15, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x16, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x41, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x32, 0xdf, 0x01, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12,
	0x3e, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x47, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x19, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x1b, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x68, 0x71,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_subsystems_api_grpc_api_task_proto_rawDescOnce sync.Once
	file_internal_app_subsystems_api_grpc_api_task_proto_rawDescData = file_internal_app_subsystems_api_grpc_api_task_proto_rawDesc
)

func file_internal_app_subsystems_api_grpc_api_task_proto_rawDescGZIP() []byte {
	file_internal_app_subsystems_api_grpc_api_task_proto_rawDescOnce.Do(func() {
		file_internal_app_subsystems_api_grpc_api_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_subsystems_api_grpc_api_task_proto_rawDescData)
	})
	return file_internal_app_subsystems_api_grpc_api_task_proto_rawDescData
}

var file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_app_subsystems_api_grpc_api_task_proto_goTypes = []interface{}{
	(*ClaimTaskRequest)(nil),       // 0: task.ClaimTaskRequest
	(*ClaimTaskResponse)(nil),      // 1: task.ClaimTaskResponse
	(*CompleteTaskRequest)(nil),    // 2: task.CompleteTaskRequest
	(*CompleteTaskResponse)(nil),   // 3: task.CompleteTaskResponse
	(*HeartbeatTasksRequest)(nil),  // 4: task.HeartbeatTasksRequest
	(*HeartbeatTasksResponse)(nil), // 5: task.HeartbeatTasksResponse
}
var file_internal_app_subsystems_api_grpc_api_task_proto_depIdxs = []int32{
	0, // 0: task.Tasks.ClaimTask:input_type -> task.ClaimTaskRequest
	2, // 1: task.Tasks.CompleteTask:input_type -> task.CompleteTaskRequest
	4, // 2: task.Tasks.HeartbeatTasks:input_type -> task.HeartbeatTasksRequest
	1, // 3: task.Tasks.ClaimTask:output_type -> task.ClaimTaskResponse
	3, // 4: task.Tasks.CompleteTask:output_type -> task.CompleteTaskResponse
	5, // 5: task.Tasks.HeartbeatTasks:output_type -> task.HeartbeatTasksResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_app_subsystems_api_grpc_api_task_proto_init() }
func file_internal_app_subsystems_api_grpc_api_task_proto_init() {
	if File_internal_app_subsystems_api_grpc_api_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimTaskRequest); i {
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
		file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimTaskResponse); i {
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
		file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteTaskRequest); i {
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
		file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteTaskResponse); i {
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
		file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatTasksRequest); i {
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
		file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatTasksResponse); i {
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
			RawDescriptor: file_internal_app_subsystems_api_grpc_api_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_subsystems_api_grpc_api_task_proto_goTypes,
		DependencyIndexes: file_internal_app_subsystems_api_grpc_api_task_proto_depIdxs,
		MessageInfos:      file_internal_app_subsystems_api_grpc_api_task_proto_msgTypes,
	}.Build()
	File_internal_app_subsystems_api_grpc_api_task_proto = out.File
	file_internal_app_subsystems_api_grpc_api_task_proto_rawDesc = nil
	file_internal_app_subsystems_api_grpc_api_task_proto_goTypes = nil
	file_internal_app_subsystems_api_grpc_api_task_proto_depIdxs = nil
}
