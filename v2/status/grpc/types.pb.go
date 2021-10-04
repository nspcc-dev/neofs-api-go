// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: v2/status/grpc/types.proto

package status

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

// Section identifiers.
type Section int32

const (
	// Successful return codes.
	Section_SECTION_SUCCESS Section = 0
	// Failure codes regardless of the operation.
	Section_SECTION_FAILURE_COMMON Section = 1
)

// Enum value maps for Section.
var (
	Section_name = map[int32]string{
		0: "SECTION_SUCCESS",
		1: "SECTION_FAILURE_COMMON",
	}
	Section_value = map[string]int32{
		"SECTION_SUCCESS":        0,
		"SECTION_FAILURE_COMMON": 1,
	}
)

func (x Section) Enum() *Section {
	p := new(Section)
	*p = x
	return p
}

func (x Section) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Section) Descriptor() protoreflect.EnumDescriptor {
	return file_v2_status_grpc_types_proto_enumTypes[0].Descriptor()
}

func (Section) Type() protoreflect.EnumType {
	return &file_v2_status_grpc_types_proto_enumTypes[0]
}

func (x Section) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Section.Descriptor instead.
func (Section) EnumDescriptor() ([]byte, []int) {
	return file_v2_status_grpc_types_proto_rawDescGZIP(), []int{0}
}

// Section of NeoFS successful return codes.
type Success int32

const (
	// [**0**] Default success. Not detailed.
	// If the server cannot match successful outcome to the code, it should
	// use this code.
	Success_OK Success = 0
)

// Enum value maps for Success.
var (
	Success_name = map[int32]string{
		0: "OK",
	}
	Success_value = map[string]int32{
		"OK": 0,
	}
)

func (x Success) Enum() *Success {
	p := new(Success)
	*p = x
	return p
}

func (x Success) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Success) Descriptor() protoreflect.EnumDescriptor {
	return file_v2_status_grpc_types_proto_enumTypes[1].Descriptor()
}

func (Success) Type() protoreflect.EnumType {
	return &file_v2_status_grpc_types_proto_enumTypes[1]
}

func (x Success) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Success.Descriptor instead.
func (Success) EnumDescriptor() ([]byte, []int) {
	return file_v2_status_grpc_types_proto_rawDescGZIP(), []int{1}
}

// Section of failed statuses independent of the operation.
type CommonFail int32

const (
	// [**1024**] Internal server error, default failure. Not detailed.
	// If the server cannot match failed outcome to the code, it should
	// use this code.
	CommonFail_INTERNAL CommonFail = 0
)

// Enum value maps for CommonFail.
var (
	CommonFail_name = map[int32]string{
		0: "INTERNAL",
	}
	CommonFail_value = map[string]int32{
		"INTERNAL": 0,
	}
)

func (x CommonFail) Enum() *CommonFail {
	p := new(CommonFail)
	*p = x
	return p
}

func (x CommonFail) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonFail) Descriptor() protoreflect.EnumDescriptor {
	return file_v2_status_grpc_types_proto_enumTypes[2].Descriptor()
}

func (CommonFail) Type() protoreflect.EnumType {
	return &file_v2_status_grpc_types_proto_enumTypes[2]
}

func (x CommonFail) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonFail.Descriptor instead.
func (CommonFail) EnumDescriptor() ([]byte, []int) {
	return file_v2_status_grpc_types_proto_rawDescGZIP(), []int{2}
}

// Declares the general format of the status returns of the NeoFS RPC protocol.
// Status is present in all response messages. Each RPC of NeoFS protocol
// describes the possible outcomes and details of the operation.
//
// Each status is assigned a one-to-one numeric code. Any unique result of an
// operation in NeoFS is unambiguously associated with the code value.
//
// Numerical set of codes is split into 1024-element sections. An enumeration
// is defined for each section. Values can be referred to in the following ways:
//
// * numerical value ranging from 0 to 4,294,967,295 (global code);
//
// * values from enumeration (local code). The formula for the ratio of the
//   local code (`L`) of a defined section (`S`) to the global one (`G`):
//   `G = 1024 * S + L`.
//
// All outcomes are divided into successful and failed, which corresponds
// to the success or failure of the operation. The definition of success
// follows from the semantics of RPC and the description of its purpose.
// The server must not attach code that is the opposite of outcome type.
//
// See the set of return codes in the description for calls.
//
// Each status can carry developer-facing error message. It should be human
// readable text in English. The server should not transmit (and the client
// should not expect) useful information in the message. Field `details`
// should make the return more detailed.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status code
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Developer-facing error message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Data detailing the outcome of the operation. Must be unique by ID.
	Details []*Status_Detail `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_status_grpc_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_v2_status_grpc_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_v2_status_grpc_types_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Status) GetDetails() []*Status_Detail {
	if x != nil {
		return x.Details
	}
	return nil
}

// Return detail. It contains additional information that can be used to
// analyze the response. Each code defines a set of details that can be
// attached to a status. Client should not handle details that are not
// covered by the code.
type Status_Detail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Detail ID. The identifier is required to determine the binary format
	// of the detail and how to decode it.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Binary status detail. Must follow the format associated with ID.
	// The possibility of missing a value must be explicitly allowed.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Status_Detail) Reset() {
	*x = Status_Detail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_status_grpc_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status_Detail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status_Detail) ProtoMessage() {}

func (x *Status_Detail) ProtoReflect() protoreflect.Message {
	mi := &file_v2_status_grpc_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status_Detail.ProtoReflect.Descriptor instead.
func (*Status_Detail) Descriptor() ([]byte, []int) {
	return file_v2_status_grpc_types_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Status_Detail) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Status_Detail) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_v2_status_grpc_types_proto protoreflect.FileDescriptor

var file_v2_status_grpc_types_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x65,
	0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa1,
	0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x1a, 0x2e, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2a, 0x3a, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a,
	0x0f, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x55, 0x52, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x10, 0x01, 0x2a, 0x11,
	0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10,
	0x00, 0x2a, 0x1a, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x12,
	0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x00, 0x42, 0x56, 0x5a,
	0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x73, 0x70, 0x63,
	0x63, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0xaa, 0x02, 0x1a, 0x4e, 0x65, 0x6f, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_status_grpc_types_proto_rawDescOnce sync.Once
	file_v2_status_grpc_types_proto_rawDescData = file_v2_status_grpc_types_proto_rawDesc
)

func file_v2_status_grpc_types_proto_rawDescGZIP() []byte {
	file_v2_status_grpc_types_proto_rawDescOnce.Do(func() {
		file_v2_status_grpc_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_status_grpc_types_proto_rawDescData)
	})
	return file_v2_status_grpc_types_proto_rawDescData
}

var file_v2_status_grpc_types_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_v2_status_grpc_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v2_status_grpc_types_proto_goTypes = []interface{}{
	(Section)(0),          // 0: neo.fs.v2.status.Section
	(Success)(0),          // 1: neo.fs.v2.status.Success
	(CommonFail)(0),       // 2: neo.fs.v2.status.CommonFail
	(*Status)(nil),        // 3: neo.fs.v2.status.Status
	(*Status_Detail)(nil), // 4: neo.fs.v2.status.Status.Detail
}
var file_v2_status_grpc_types_proto_depIdxs = []int32{
	4, // 0: neo.fs.v2.status.Status.details:type_name -> neo.fs.v2.status.Status.Detail
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_status_grpc_types_proto_init() }
func file_v2_status_grpc_types_proto_init() {
	if File_v2_status_grpc_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_status_grpc_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_v2_status_grpc_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status_Detail); i {
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
			RawDescriptor: file_v2_status_grpc_types_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_status_grpc_types_proto_goTypes,
		DependencyIndexes: file_v2_status_grpc_types_proto_depIdxs,
		EnumInfos:         file_v2_status_grpc_types_proto_enumTypes,
		MessageInfos:      file_v2_status_grpc_types_proto_msgTypes,
	}.Build()
	File_v2_status_grpc_types_proto = out.File
	file_v2_status_grpc_types_proto_rawDesc = nil
	file_v2_status_grpc_types_proto_goTypes = nil
	file_v2_status_grpc_types_proto_depIdxs = nil
}
