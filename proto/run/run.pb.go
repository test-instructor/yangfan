// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.2
// source: run.proto

package run

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

type RunningType int32

const (
	RunningType_RunningTypeDefault   RunningType = 0
	RunningType_RunningTypeRun       RunningType = 1
	RunningType_RunningTypeRebalance RunningType = 2
	RunningType_RunningTypeStop      RunningType = 3
)

// Enum value maps for RunningType.
var (
	RunningType_name = map[int32]string{
		0: "RunningTypeDefault",
		1: "RunningTypeRun",
		2: "RunningTypeRebalance",
		3: "RunningTypeStop",
	}
	RunningType_value = map[string]int32{
		"RunningTypeDefault":   0,
		"RunningTypeRun":       1,
		"RunningTypeRebalance": 2,
		"RunningTypeStop":      3,
	}
)

func (x RunningType) Enum() *RunningType {
	p := new(RunningType)
	*p = x
	return p
}

func (x RunningType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunningType) Descriptor() protoreflect.EnumDescriptor {
	return file_run_proto_enumTypes[0].Descriptor()
}

func (RunningType) Type() protoreflect.EnumType {
	return &file_run_proto_enumTypes[0]
}

func (x RunningType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunningType.Descriptor instead.
func (RunningType) EnumDescriptor() ([]byte, []int) {
	return file_run_proto_rawDescGZIP(), []int{0}
}

type CaseType int32

const (
	CaseType_CaseTypeDefault     CaseType = 0 // 默认
	CaseType_CaseTypeApi         CaseType = 1 // api 运行
	CaseType_CaseTypeStep        CaseType = 2 // 步骤运行
	CaseType_CaseTypeCases       CaseType = 3 // 用例运行
	CaseType_CaseTypeTask        CaseType = 4 // 定时任务运行
	CaseType_CaseTypeBoomerDebug CaseType = 5 // 性能任务调试
	CaseType_CaseTypeBoomer      CaseType = 6 // 性能任务运行
	CaseType_CaseTypeTag         CaseType = 7 // 性能任务运行
)

// Enum value maps for CaseType.
var (
	CaseType_name = map[int32]string{
		0: "CaseTypeDefault",
		1: "CaseTypeApi",
		2: "CaseTypeStep",
		3: "CaseTypeCases",
		4: "CaseTypeTask",
		5: "CaseTypeBoomerDebug",
		6: "CaseTypeBoomer",
		7: "CaseTypeTag",
	}
	CaseType_value = map[string]int32{
		"CaseTypeDefault":     0,
		"CaseTypeApi":         1,
		"CaseTypeStep":        2,
		"CaseTypeCases":       3,
		"CaseTypeTask":        4,
		"CaseTypeBoomerDebug": 5,
		"CaseTypeBoomer":      6,
		"CaseTypeTag":         7,
	}
)

func (x CaseType) Enum() *CaseType {
	p := new(CaseType)
	*p = x
	return p
}

func (x CaseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CaseType) Descriptor() protoreflect.EnumDescriptor {
	return file_run_proto_enumTypes[1].Descriptor()
}

func (CaseType) Type() protoreflect.EnumType {
	return &file_run_proto_enumTypes[1]
}

func (x CaseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CaseType.Descriptor instead.
func (CaseType) EnumDescriptor() ([]byte, []int) {
	return file_run_proto_rawDescGZIP(), []int{1}
}

type RunType int32

const (
	RunType_RunTypeDefault  RunType = 0 // 默认
	RunType_RunTypeDebug    RunType = 1 // 调试模式
	RunType_RunTypeRuning   RunType = 2 // 立即运行
	RunType_RunTypeRunBack  RunType = 3 // 后台运行
	RunType_RunTypeRunTimer RunType = 4 // 定时执行
	RunType_RunTypeRunSave  RunType = 5 // 调试并保存
)

// Enum value maps for RunType.
var (
	RunType_name = map[int32]string{
		0: "RunTypeDefault",
		1: "RunTypeDebug",
		2: "RunTypeRuning",
		3: "RunTypeRunBack",
		4: "RunTypeRunTimer",
		5: "RunTypeRunSave",
	}
	RunType_value = map[string]int32{
		"RunTypeDefault":  0,
		"RunTypeDebug":    1,
		"RunTypeRuning":   2,
		"RunTypeRunBack":  3,
		"RunTypeRunTimer": 4,
		"RunTypeRunSave":  5,
	}
)

func (x RunType) Enum() *RunType {
	p := new(RunType)
	*p = x
	return p
}

func (x RunType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunType) Descriptor() protoreflect.EnumDescriptor {
	return file_run_proto_enumTypes[2].Descriptor()
}

func (RunType) Type() protoreflect.EnumType {
	return &file_run_proto_enumTypes[2]
}

func (x RunType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunType.Descriptor instead.
func (RunType) EnumDescriptor() ([]byte, []int) {
	return file_run_proto_rawDescGZIP(), []int{2}
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Running    RunningType `protobuf:"varint,1,opt,name=running,proto3,enum=run.RunningType" json:"running,omitempty"`
	SpawnCount int64       `protobuf:"varint,2,opt,name=spawnCount,proto3" json:"spawnCount,omitempty"`
	SpawnRate  float32     `protobuf:"fixed32,3,opt,name=spawnRate,proto3" json:"spawnRate,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_run_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_run_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_run_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetRunning() RunningType {
	if x != nil {
		return x.Running
	}
	return RunningType_RunningTypeDefault
}

func (x *Operation) GetSpawnCount() int64 {
	if x != nil {
		return x.SpawnCount
	}
	return 0
}

func (x *Operation) GetSpawnRate() float32 {
	if x != nil {
		return x.SpawnRate
	}
	return 0
}

type RunCaseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiID     uint32     `protobuf:"varint,1,opt,name=apiID,proto3" json:"apiID,omitempty"`
	ConfigID  uint32     `protobuf:"varint,2,opt,name=configID,proto3" json:"configID,omitempty"`
	CaseID    uint32     `protobuf:"varint,3,opt,name=caseID,proto3" json:"caseID,omitempty"`
	RunType   RunType    `protobuf:"varint,4,opt,name=run_type,json=runType,proto3,enum=run.RunType" json:"run_type,omitempty"`
	TaskID    uint32     `protobuf:"varint,5,opt,name=taskID,proto3" json:"taskID,omitempty"`
	Operation *Operation `protobuf:"bytes,6,opt,name=operation,proto3" json:"operation,omitempty"`
	TagID     uint32     `protobuf:"varint,7,opt,name=tagID,proto3" json:"tagID,omitempty"`
	ProjectID uint32     `protobuf:"varint,8,opt,name=projectID,proto3" json:"projectID,omitempty"`
	Env       uint32     `protobuf:"varint,9,opt,name=env,proto3" json:"env,omitempty"`
	ReportID  uint32     `protobuf:"varint,10,opt,name=reportID,proto3" json:"reportID,omitempty"`
	CaseType  CaseType   `protobuf:"varint,11,opt,name=case_type,json=caseType,proto3,enum=run.CaseType" json:"case_type,omitempty"`
}

func (x *RunCaseReq) Reset() {
	*x = RunCaseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_run_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCaseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCaseReq) ProtoMessage() {}

func (x *RunCaseReq) ProtoReflect() protoreflect.Message {
	mi := &file_run_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCaseReq.ProtoReflect.Descriptor instead.
func (*RunCaseReq) Descriptor() ([]byte, []int) {
	return file_run_proto_rawDescGZIP(), []int{1}
}

func (x *RunCaseReq) GetApiID() uint32 {
	if x != nil {
		return x.ApiID
	}
	return 0
}

func (x *RunCaseReq) GetConfigID() uint32 {
	if x != nil {
		return x.ConfigID
	}
	return 0
}

func (x *RunCaseReq) GetCaseID() uint32 {
	if x != nil {
		return x.CaseID
	}
	return 0
}

func (x *RunCaseReq) GetRunType() RunType {
	if x != nil {
		return x.RunType
	}
	return RunType_RunTypeDefault
}

func (x *RunCaseReq) GetTaskID() uint32 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *RunCaseReq) GetOperation() *Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *RunCaseReq) GetTagID() uint32 {
	if x != nil {
		return x.TagID
	}
	return 0
}

func (x *RunCaseReq) GetProjectID() uint32 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

func (x *RunCaseReq) GetEnv() uint32 {
	if x != nil {
		return x.Env
	}
	return 0
}

func (x *RunCaseReq) GetReportID() uint32 {
	if x != nil {
		return x.ReportID
	}
	return 0
}

func (x *RunCaseReq) GetCaseType() CaseType {
	if x != nil {
		return x.CaseType
	}
	return CaseType_CaseTypeDefault
}

type RunCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportID uint32 `protobuf:"varint,1,opt,name=reportID,proto3" json:"reportID,omitempty"`
}

func (x *RunCaseResponse) Reset() {
	*x = RunCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_run_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCaseResponse) ProtoMessage() {}

func (x *RunCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_run_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCaseResponse.ProtoReflect.Descriptor instead.
func (*RunCaseResponse) Descriptor() ([]byte, []int) {
	return file_run_proto_rawDescGZIP(), []int{2}
}

func (x *RunCaseResponse) GetReportID() uint32 {
	if x != nil {
		return x.ReportID
	}
	return 0
}

var File_run_proto protoreflect.FileDescriptor

var file_run_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x75, 0x6e,
	0x22, 0x75, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a,
	0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x61,
	0x77, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73,
	0x70, 0x61, 0x77, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x70, 0x61,
	0x77, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x73, 0x70,
	0x61, 0x77, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x22, 0xd3, 0x02, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x43,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x69, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x69, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x73, 0x65,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x61, 0x73, 0x65, 0x49, 0x44,
	0x12, 0x27, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x07, 0x72, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x44, 0x12, 0x2c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x61, 0x67, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x74, 0x61, 0x67, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49,
	0x44, 0x12, 0x2a, 0x0a, 0x09, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2d, 0x0a,
	0x0f, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x2a, 0x68, 0x0a, 0x0b,
	0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x52,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x75, 0x6e, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x10,
	0x02, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x74, 0x6f, 0x70, 0x10, 0x03, 0x2a, 0xa5, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x73, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x61, 0x73, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x41, 0x70, 0x69, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x61, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x65, 0x70, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x61, 0x73, 0x65, 0x73, 0x10, 0x03, 0x12, 0x10,
	0x0a, 0x0c, 0x43, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x10, 0x04,
	0x12, 0x17, 0x0a, 0x13, 0x43, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x6f, 0x6f, 0x6d,
	0x65, 0x72, 0x44, 0x65, 0x62, 0x75, 0x67, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x61, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x6f, 0x6f, 0x6d, 0x65, 0x72, 0x10, 0x06, 0x12, 0x0f, 0x0a,
	0x0b, 0x43, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x54, 0x61, 0x67, 0x10, 0x07, 0x2a, 0x7f,
	0x0a, 0x07, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x75, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x62, 0x75, 0x67, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x75, 0x6e, 0x69, 0x6e, 0x67,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x75, 0x6e,
	0x42, 0x61, 0x63, 0x6b, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x52,
	0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x75, 0x6e, 0x53, 0x61, 0x76, 0x65, 0x10, 0x05, 0x32,
	0xd0, 0x02, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x52,
	0x75, 0x6e, 0x41, 0x70, 0x69, 0x12, 0x0f, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x43,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x52, 0x75, 0x6e,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x07, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x12, 0x0f, 0x2e, 0x72, 0x75, 0x6e, 0x2e,
	0x52, 0x75, 0x6e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x72, 0x75, 0x6e,
	0x2e, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x73, 0x65, 0x12, 0x0f, 0x2e,
	0x72, 0x75, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x42, 0x6f, 0x6f,
	0x6d, 0x65, 0x72, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x0f, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x52,
	0x75, 0x6e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x72, 0x75, 0x6e, 0x2e,
	0x52, 0x75, 0x6e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x0f, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x14, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0b, 0x52, 0x75,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12, 0x0f, 0x2e, 0x72, 0x75, 0x6e, 0x2e,
	0x52, 0x75, 0x6e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x72, 0x75, 0x6e,
	0x2e, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x72, 0x75, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_run_proto_rawDescOnce sync.Once
	file_run_proto_rawDescData = file_run_proto_rawDesc
)

func file_run_proto_rawDescGZIP() []byte {
	file_run_proto_rawDescOnce.Do(func() {
		file_run_proto_rawDescData = protoimpl.X.CompressGZIP(file_run_proto_rawDescData)
	})
	return file_run_proto_rawDescData
}

var file_run_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_run_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_run_proto_goTypes = []interface{}{
	(RunningType)(0),        // 0: run.RunningType
	(CaseType)(0),           // 1: run.CaseType
	(RunType)(0),            // 2: run.RunType
	(*Operation)(nil),       // 3: run.Operation
	(*RunCaseReq)(nil),      // 4: run.RunCaseReq
	(*RunCaseResponse)(nil), // 5: run.RunCaseResponse
}
var file_run_proto_depIdxs = []int32{
	0,  // 0: run.Operation.running:type_name -> run.RunningType
	2,  // 1: run.RunCaseReq.run_type:type_name -> run.RunType
	3,  // 2: run.RunCaseReq.operation:type_name -> run.Operation
	1,  // 3: run.RunCaseReq.case_type:type_name -> run.CaseType
	4,  // 4: run.RunCase.RunApi:input_type -> run.RunCaseReq
	4,  // 5: run.RunCase.RunStep:input_type -> run.RunCaseReq
	4,  // 6: run.RunCase.RunCase:input_type -> run.RunCaseReq
	4,  // 7: run.RunCase.RunBoomerDebug:input_type -> run.RunCaseReq
	4,  // 8: run.RunCase.RunTimerTask:input_type -> run.RunCaseReq
	4,  // 9: run.RunCase.RunTimerTag:input_type -> run.RunCaseReq
	5,  // 10: run.RunCase.RunApi:output_type -> run.RunCaseResponse
	5,  // 11: run.RunCase.RunStep:output_type -> run.RunCaseResponse
	5,  // 12: run.RunCase.RunCase:output_type -> run.RunCaseResponse
	5,  // 13: run.RunCase.RunBoomerDebug:output_type -> run.RunCaseResponse
	5,  // 14: run.RunCase.RunTimerTask:output_type -> run.RunCaseResponse
	5,  // 15: run.RunCase.RunTimerTag:output_type -> run.RunCaseResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_run_proto_init() }
func file_run_proto_init() {
	if File_run_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_run_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
		file_run_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCaseReq); i {
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
		file_run_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCaseResponse); i {
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
			RawDescriptor: file_run_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_run_proto_goTypes,
		DependencyIndexes: file_run_proto_depIdxs,
		EnumInfos:         file_run_proto_enumTypes,
		MessageInfos:      file_run_proto_msgTypes,
	}.Build()
	File_run_proto = out.File
	file_run_proto_rawDesc = nil
	file_run_proto_goTypes = nil
	file_run_proto_depIdxs = nil
}
