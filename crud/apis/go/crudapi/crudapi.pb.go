// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: apis/crudapi.proto

package crudapi

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// TodoStatus enum defines the possible states of a Todo object.
type TodoStatus int32

const (
	TodoStatus_UNKNOWN     TodoStatus = 0
	TodoStatus_PENDING     TodoStatus = 1
	TodoStatus_IN_PROGRESS TodoStatus = 2
	TodoStatus_DONE        TodoStatus = 3
	TodoStatus_DELETED     TodoStatus = 4
	TodoStatus_ABANDONED   TodoStatus = 5
)

// Enum value maps for TodoStatus.
var (
	TodoStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "PENDING",
		2: "IN_PROGRESS",
		3: "DONE",
		4: "DELETED",
		5: "ABANDONED",
	}
	TodoStatus_value = map[string]int32{
		"UNKNOWN":     0,
		"PENDING":     1,
		"IN_PROGRESS": 2,
		"DONE":        3,
		"DELETED":     4,
		"ABANDONED":   5,
	}
)

func (x TodoStatus) Enum() *TodoStatus {
	p := new(TodoStatus)
	*p = x
	return p
}

func (x TodoStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TodoStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_apis_crudapi_proto_enumTypes[0].Descriptor()
}

func (TodoStatus) Type() protoreflect.EnumType {
	return &file_apis_crudapi_proto_enumTypes[0]
}

func (x TodoStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TodoStatus.Descriptor instead.
func (TodoStatus) EnumDescriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{0}
}

// Todo message represents a Todo object.
type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// uuid uniquely represents any Todo object in the storage. The value of
	// id will be generated by the system while the todo is being crated.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// User provided display name of the todo.
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// User defined priority of the TODO, can only be set a string value of P(0-4).
	Priority string `protobuf:"bytes,4,opt,name=priority,proto3" json:"priority,omitempty"`
	// Created at timestamp, represents when the entry is being created by the user.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// User defined deadline specified in duration, indicates the time by when the task
	// should be completed.
	Deadline *durationpb.Duration `protobuf:"bytes,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// Status of the todo.
	Status TodoStatus `protobuf:"varint,7,opt,name=status,proto3,enum=crudapi.v1.TodoStatus" json:"status,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_crudapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_apis_crudapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Todo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Todo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Todo) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *Todo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Todo) GetDeadline() *durationpb.Duration {
	if x != nil {
		return x.Deadline
	}
	return nil
}

func (x *Todo) GetStatus() TodoStatus {
	if x != nil {
		return x.Status
	}
	return TodoStatus_UNKNOWN
}

type CreateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *CreateTodoRequest) Reset() {
	*x = CreateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_crudapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequest) ProtoMessage() {}

func (x *CreateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_crudapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequest.ProtoReflect.Descriptor instead.
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTodoRequest) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type CreateTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *CreateTodoResponse) Reset() {
	*x = CreateTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_crudapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoResponse) ProtoMessage() {}

func (x *CreateTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_crudapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoResponse.ProtoReflect.Descriptor instead.
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTodoResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type ListTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit     int64      `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset    int64      `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	SortOrder string     `protobuf:"bytes,3,opt,name=sort_order,json=sortOrder,proto3" json:"sort_order,omitempty"`
	Status    TodoStatus `protobuf:"varint,4,opt,name=status,proto3,enum=crudapi.v1.TodoStatus" json:"status,omitempty"`
	Priority  string     `protobuf:"bytes,5,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *ListTodoRequest) Reset() {
	*x = ListTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_crudapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodoRequest) ProtoMessage() {}

func (x *ListTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_crudapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodoRequest.ProtoReflect.Descriptor instead.
func (*ListTodoRequest) Descriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{3}
}

func (x *ListTodoRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTodoRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListTodoRequest) GetSortOrder() string {
	if x != nil {
		return x.SortOrder
	}
	return ""
}

func (x *ListTodoRequest) GetStatus() TodoStatus {
	if x != nil {
		return x.Status
	}
	return TodoStatus_UNKNOWN
}

func (x *ListTodoRequest) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

type ListTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *ListTodoResponse) Reset() {
	*x = ListTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_crudapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodoResponse) ProtoMessage() {}

func (x *ListTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_crudapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodoResponse.ProtoReflect.Descriptor instead.
func (*ListTodoResponse) Descriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{4}
}

func (x *ListTodoResponse) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

type GetTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTodoRequest) Reset() {
	*x = GetTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_crudapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoRequest) ProtoMessage() {}

func (x *GetTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_crudapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoRequest.ProtoReflect.Descriptor instead.
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{5}
}

func (x *GetTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *GetTodoResponse) Reset() {
	*x = GetTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_crudapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoResponse) ProtoMessage() {}

func (x *GetTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_crudapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoResponse.ProtoReflect.Descriptor instead.
func (*GetTodoResponse) Descriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{6}
}

func (x *GetTodoResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type UpdateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload *Todo  `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *UpdateTodoRequest) Reset() {
	*x = UpdateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_crudapi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoRequest) ProtoMessage() {}

func (x *UpdateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_crudapi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoRequest.ProtoReflect.Descriptor instead.
func (*UpdateTodoRequest) Descriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTodoRequest) GetPayload() *Todo {
	if x != nil {
		return x.Payload
	}
	return nil
}

type UpdateTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *UpdateTodoResponse) Reset() {
	*x = UpdateTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_crudapi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoResponse) ProtoMessage() {}

func (x *UpdateTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_crudapi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoResponse.ProtoReflect.Descriptor instead.
func (*UpdateTodoResponse) Descriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTodoResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type DeleteTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTodoRequest) Reset() {
	*x = DeleteTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_crudapi_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoRequest) ProtoMessage() {}

func (x *DeleteTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_crudapi_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoRequest.ProtoReflect.Descriptor instead.
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_apis_crudapi_proto protoreflect.FileDescriptor

var file_apis_crudapi_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x02, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x64, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xf4, 0x03, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xfa, 0x42, 0x16,
	0x72, 0x14, 0x52, 0x02, 0x50, 0x30, 0x52, 0x02, 0x50, 0x31, 0x52, 0x02, 0x50, 0x32, 0x52, 0x02,
	0x50, 0x33, 0x52, 0x02, 0x50, 0x34, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x64,
	0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x38, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63,
	0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x43, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x72, 0x75, 0x64,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x3a, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x72,
	0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x22, 0x3a, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04,
	0x74, 0x6f, 0x64, 0x6f, 0x22, 0x62, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x3a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04,
	0x74, 0x6f, 0x64, 0x6f, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x2a, 0x5d, 0x0a, 0x0a, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e,
	0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x4f, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x42, 0x41, 0x4e, 0x44, 0x4f, 0x4e, 0x45, 0x44, 0x10,
	0x05, 0x32, 0xfa, 0x03, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x64, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12,
	0x1d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x6f, 0x64, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x5b, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x6f, 0x64, 0x6f, 0x12, 0x1b, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x6f, 0x64, 0x6f, 0x12, 0x5d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x12,
	0x1a, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x72,
	0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x69, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64,
	0x6f, 0x12, 0x1d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x1a, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x5e,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x1d, 0x2e, 0x63,
	0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xb1,
	0x01, 0x5a, 0x09, 0x2e, 0x2f, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x92, 0x41, 0xa2, 0x01,
	0x12, 0x2e, 0x0a, 0x10, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x43, 0x52, 0x55, 0x44, 0x20,
	0x41, 0x50, 0x49, 0x73, 0x22, 0x12, 0x1a, 0x10, 0x6b, 0x7a, 0x69, 0x72, 0x74, 0x6d, 0x40, 0x67,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x06, 0x76, 0x30, 0x2e, 0x31, 0x2e, 0x30,
	0x1a, 0x15, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f,
	0x73, 0x74, 0x3a, 0x36, 0x39, 0x39, 0x39, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x32,
	0x12, 0x30, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x64, 0x6c, 0x69, 0x6c, 0x2f, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x63, 0x72,
	0x75, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_crudapi_proto_rawDescOnce sync.Once
	file_apis_crudapi_proto_rawDescData = file_apis_crudapi_proto_rawDesc
)

func file_apis_crudapi_proto_rawDescGZIP() []byte {
	file_apis_crudapi_proto_rawDescOnce.Do(func() {
		file_apis_crudapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_crudapi_proto_rawDescData)
	})
	return file_apis_crudapi_proto_rawDescData
}

var file_apis_crudapi_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apis_crudapi_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_apis_crudapi_proto_goTypes = []interface{}{
	(TodoStatus)(0),               // 0: crudapi.v1.TodoStatus
	(*Todo)(nil),                  // 1: crudapi.v1.Todo
	(*CreateTodoRequest)(nil),     // 2: crudapi.v1.CreateTodoRequest
	(*CreateTodoResponse)(nil),    // 3: crudapi.v1.CreateTodoResponse
	(*ListTodoRequest)(nil),       // 4: crudapi.v1.ListTodoRequest
	(*ListTodoResponse)(nil),      // 5: crudapi.v1.ListTodoResponse
	(*GetTodoRequest)(nil),        // 6: crudapi.v1.GetTodoRequest
	(*GetTodoResponse)(nil),       // 7: crudapi.v1.GetTodoResponse
	(*UpdateTodoRequest)(nil),     // 8: crudapi.v1.UpdateTodoRequest
	(*UpdateTodoResponse)(nil),    // 9: crudapi.v1.UpdateTodoResponse
	(*DeleteTodoRequest)(nil),     // 10: crudapi.v1.DeleteTodoRequest
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 12: google.protobuf.Duration
	(*emptypb.Empty)(nil),         // 13: google.protobuf.Empty
}
var file_apis_crudapi_proto_depIdxs = []int32{
	11, // 0: crudapi.v1.Todo.created_at:type_name -> google.protobuf.Timestamp
	12, // 1: crudapi.v1.Todo.deadline:type_name -> google.protobuf.Duration
	0,  // 2: crudapi.v1.Todo.status:type_name -> crudapi.v1.TodoStatus
	1,  // 3: crudapi.v1.CreateTodoRequest.todo:type_name -> crudapi.v1.Todo
	1,  // 4: crudapi.v1.CreateTodoResponse.todo:type_name -> crudapi.v1.Todo
	0,  // 5: crudapi.v1.ListTodoRequest.status:type_name -> crudapi.v1.TodoStatus
	1,  // 6: crudapi.v1.ListTodoResponse.todos:type_name -> crudapi.v1.Todo
	1,  // 7: crudapi.v1.GetTodoResponse.todo:type_name -> crudapi.v1.Todo
	1,  // 8: crudapi.v1.UpdateTodoRequest.payload:type_name -> crudapi.v1.Todo
	1,  // 9: crudapi.v1.UpdateTodoResponse.todo:type_name -> crudapi.v1.Todo
	2,  // 10: crudapi.v1.TodoService.CreateTodo:input_type -> crudapi.v1.CreateTodoRequest
	4,  // 11: crudapi.v1.TodoService.ListTodo:input_type -> crudapi.v1.ListTodoRequest
	6,  // 12: crudapi.v1.TodoService.GetTodo:input_type -> crudapi.v1.GetTodoRequest
	8,  // 13: crudapi.v1.TodoService.UpdateTodo:input_type -> crudapi.v1.UpdateTodoRequest
	10, // 14: crudapi.v1.TodoService.DeleteTodo:input_type -> crudapi.v1.DeleteTodoRequest
	3,  // 15: crudapi.v1.TodoService.CreateTodo:output_type -> crudapi.v1.CreateTodoResponse
	5,  // 16: crudapi.v1.TodoService.ListTodo:output_type -> crudapi.v1.ListTodoResponse
	7,  // 17: crudapi.v1.TodoService.GetTodo:output_type -> crudapi.v1.GetTodoResponse
	9,  // 18: crudapi.v1.TodoService.UpdateTodo:output_type -> crudapi.v1.UpdateTodoResponse
	13, // 19: crudapi.v1.TodoService.DeleteTodo:output_type -> google.protobuf.Empty
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_apis_crudapi_proto_init() }
func file_apis_crudapi_proto_init() {
	if File_apis_crudapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_crudapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_apis_crudapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoRequest); i {
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
		file_apis_crudapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoResponse); i {
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
		file_apis_crudapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTodoRequest); i {
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
		file_apis_crudapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTodoResponse); i {
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
		file_apis_crudapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoRequest); i {
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
		file_apis_crudapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoResponse); i {
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
		file_apis_crudapi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoRequest); i {
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
		file_apis_crudapi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoResponse); i {
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
		file_apis_crudapi_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodoRequest); i {
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
			RawDescriptor: file_apis_crudapi_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apis_crudapi_proto_goTypes,
		DependencyIndexes: file_apis_crudapi_proto_depIdxs,
		EnumInfos:         file_apis_crudapi_proto_enumTypes,
		MessageInfos:      file_apis_crudapi_proto_msgTypes,
	}.Build()
	File_apis_crudapi_proto = out.File
	file_apis_crudapi_proto_rawDesc = nil
	file_apis_crudapi_proto_goTypes = nil
	file_apis_crudapi_proto_depIdxs = nil
}
