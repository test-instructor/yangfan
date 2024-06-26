// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0--rc2
// source: install.proto

package tools

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

type Operate int32

const (
	Operate_INSTALL Operate = 0
	Operate_REMOVE  Operate = 1
	Operate_UPDATE  Operate = 2
)

// Enum value maps for Operate.
var (
	Operate_name = map[int32]string{
		0: "INSTALL",
		1: "REMOVE",
		2: "UPDATE",
	}
	Operate_value = map[string]int32{
		"INSTALL": 0,
		"REMOVE":  1,
		"UPDATE":  2,
	}
)

func (x Operate) Enum() *Operate {
	p := new(Operate)
	*p = x
	return p
}

func (x Operate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operate) Descriptor() protoreflect.EnumDescriptor {
	return file_install_proto_enumTypes[0].Descriptor()
}

func (Operate) Type() protoreflect.EnumType {
	return &file_install_proto_enumTypes[0]
}

func (x Operate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operate.Descriptor instead.
func (Operate) EnumDescriptor() ([]byte, []int) {
	return file_install_proto_rawDescGZIP(), []int{0}
}

type TimerStatusOperate int32

const (
	TimerStatusOperate_ADD    TimerStatusOperate = 0
	TimerStatusOperate_DELETE TimerStatusOperate = 1
	TimerStatusOperate_RESET  TimerStatusOperate = 2
)

// Enum value maps for TimerStatusOperate.
var (
	TimerStatusOperate_name = map[int32]string{
		0: "ADD",
		1: "DELETE",
		2: "RESET",
	}
	TimerStatusOperate_value = map[string]int32{
		"ADD":    0,
		"DELETE": 1,
		"RESET":  2,
	}
)

func (x TimerStatusOperate) Enum() *TimerStatusOperate {
	p := new(TimerStatusOperate)
	*p = x
	return p
}

func (x TimerStatusOperate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimerStatusOperate) Descriptor() protoreflect.EnumDescriptor {
	return file_install_proto_enumTypes[1].Descriptor()
}

func (TimerStatusOperate) Type() protoreflect.EnumType {
	return &file_install_proto_enumTypes[1]
}

func (x TimerStatusOperate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimerStatusOperate.Descriptor instead.
func (TimerStatusOperate) EnumDescriptor() ([]byte, []int) {
	return file_install_proto_rawDescGZIP(), []int{1}
}

type InstallPackageRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Operate Operate `protobuf:"varint,3,opt,name=operate,proto3,enum=tools.Operate" json:"operate,omitempty"`
}

func (x *InstallPackageRes) Reset() {
	*x = InstallPackageRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_install_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallPackageRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallPackageRes) ProtoMessage() {}

func (x *InstallPackageRes) ProtoReflect() protoreflect.Message {
	mi := &file_install_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallPackageRes.ProtoReflect.Descriptor instead.
func (*InstallPackageRes) Descriptor() ([]byte, []int) {
	return file_install_proto_rawDescGZIP(), []int{0}
}

func (x *InstallPackageRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InstallPackageRes) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *InstallPackageRes) GetOperate() Operate {
	if x != nil {
		return x.Operate
	}
	return Operate_INSTALL
}

type InstallPackageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InstallPackageReq) Reset() {
	*x = InstallPackageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_install_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallPackageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallPackageReq) ProtoMessage() {}

func (x *InstallPackageReq) ProtoReflect() protoreflect.Message {
	mi := &file_install_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallPackageReq.ProtoReflect.Descriptor instead.
func (*InstallPackageReq) Descriptor() ([]byte, []int) {
	return file_install_proto_rawDescGZIP(), []int{1}
}

type SetTaskRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint64             `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	EntryID     int64              `protobuf:"varint,2,opt,name=EntryID,proto3" json:"EntryID,omitempty"`
	TimerStatus TimerStatusOperate `protobuf:"varint,3,opt,name=TimerStatus,proto3,enum=tools.TimerStatusOperate" json:"TimerStatus,omitempty"`
}

func (x *SetTaskRes) Reset() {
	*x = SetTaskRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_install_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTaskRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTaskRes) ProtoMessage() {}

func (x *SetTaskRes) ProtoReflect() protoreflect.Message {
	mi := &file_install_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTaskRes.ProtoReflect.Descriptor instead.
func (*SetTaskRes) Descriptor() ([]byte, []int) {
	return file_install_proto_rawDescGZIP(), []int{2}
}

func (x *SetTaskRes) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *SetTaskRes) GetEntryID() int64 {
	if x != nil {
		return x.EntryID
	}
	return 0
}

func (x *SetTaskRes) GetTimerStatus() TimerStatusOperate {
	if x != nil {
		return x.TimerStatus
	}
	return TimerStatusOperate_ADD
}

type SetTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetTaskReq) Reset() {
	*x = SetTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_install_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTaskReq) ProtoMessage() {}

func (x *SetTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_install_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTaskReq.ProtoReflect.Descriptor instead.
func (*SetTaskReq) Descriptor() ([]byte, []int) {
	return file_install_proto_rawDescGZIP(), []int{3}
}

var File_install_proto protoreflect.FileDescriptor

var file_install_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x22, 0x6b, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x07, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x22, 0x73, 0x0a, 0x0a, 0x73, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x44,
	0x12, 0x3b, 0x0a, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0c, 0x0a,
	0x0a, 0x73, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x2a, 0x2e, 0x0a, 0x07, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c,
	0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x2a, 0x34, 0x0a, 0x12, 0x54,
	0x69, 0x6d, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x53, 0x45, 0x54, 0x10,
	0x02, 0x32, 0xac, 0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x58, 0x0a, 0x1e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x17, 0x53,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x11, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x73,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x73, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_install_proto_rawDescOnce sync.Once
	file_install_proto_rawDescData = file_install_proto_rawDesc
)

func file_install_proto_rawDescGZIP() []byte {
	file_install_proto_rawDescOnce.Do(func() {
		file_install_proto_rawDescData = protoimpl.X.CompressGZIP(file_install_proto_rawDescData)
	})
	return file_install_proto_rawDescData
}

var file_install_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_install_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_install_proto_goTypes = []interface{}{
	(Operate)(0),              // 0: tools.Operate
	(TimerStatusOperate)(0),   // 1: tools.TimerStatusOperate
	(*InstallPackageRes)(nil), // 2: tools.installPackageRes
	(*InstallPackageReq)(nil), // 3: tools.installPackageReq
	(*SetTaskRes)(nil),        // 4: tools.setTaskRes
	(*SetTaskReq)(nil),        // 5: tools.setTaskReq
}
var file_install_proto_depIdxs = []int32{
	0, // 0: tools.installPackageRes.operate:type_name -> tools.Operate
	1, // 1: tools.setTaskRes.TimerStatus:type_name -> tools.TimerStatusOperate
	3, // 2: tools.ToolsServer.InstallPackageStreamingMessage:input_type -> tools.installPackageReq
	5, // 3: tools.ToolsServer.SetTaskStreamingMessage:input_type -> tools.setTaskReq
	2, // 4: tools.ToolsServer.InstallPackageStreamingMessage:output_type -> tools.installPackageRes
	4, // 5: tools.ToolsServer.SetTaskStreamingMessage:output_type -> tools.setTaskRes
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_install_proto_init() }
func file_install_proto_init() {
	if File_install_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_install_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallPackageRes); i {
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
		file_install_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallPackageReq); i {
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
		file_install_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTaskRes); i {
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
		file_install_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTaskReq); i {
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
			RawDescriptor: file_install_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_install_proto_goTypes,
		DependencyIndexes: file_install_proto_depIdxs,
		EnumInfos:         file_install_proto_enumTypes,
		MessageInfos:      file_install_proto_msgTypes,
	}.Build()
	File_install_proto = out.File
	file_install_proto_rawDesc = nil
	file_install_proto_goTypes = nil
	file_install_proto_depIdxs = nil
}
