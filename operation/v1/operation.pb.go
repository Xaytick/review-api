// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: api/operation/v1/operation.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 审核评论请求
type AuditReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewID      int64                  `protobuf:"varint,1,opt,name=reviewID,proto3" json:"reviewID,omitempty"`
	Status        int32                  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	OpUser        string                 `protobuf:"bytes,3,opt,name=opUser,proto3" json:"opUser,omitempty"`
	OpReason      string                 `protobuf:"bytes,4,opt,name=opReason,proto3" json:"opReason,omitempty"`
	OpRemarks     *string                `protobuf:"bytes,5,opt,name=opRemarks,proto3,oneof" json:"opRemarks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuditReviewRequest) Reset() {
	*x = AuditReviewRequest{}
	mi := &file_api_operation_v1_operation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditReviewRequest) ProtoMessage() {}

func (x *AuditReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_operation_v1_operation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditReviewRequest.ProtoReflect.Descriptor instead.
func (*AuditReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_operation_v1_operation_proto_rawDescGZIP(), []int{0}
}

func (x *AuditReviewRequest) GetReviewID() int64 {
	if x != nil {
		return x.ReviewID
	}
	return 0
}

func (x *AuditReviewRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AuditReviewRequest) GetOpUser() string {
	if x != nil {
		return x.OpUser
	}
	return ""
}

func (x *AuditReviewRequest) GetOpReason() string {
	if x != nil {
		return x.OpReason
	}
	return ""
}

func (x *AuditReviewRequest) GetOpRemarks() string {
	if x != nil && x.OpRemarks != nil {
		return *x.OpRemarks
	}
	return ""
}

// 审核评论回复
type AuditReviewReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewID      int64                  `protobuf:"varint,1,opt,name=reviewID,proto3" json:"reviewID,omitempty"`
	Status        int32                  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuditReviewReply) Reset() {
	*x = AuditReviewReply{}
	mi := &file_api_operation_v1_operation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditReviewReply) ProtoMessage() {}

func (x *AuditReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_operation_v1_operation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditReviewReply.ProtoReflect.Descriptor instead.
func (*AuditReviewReply) Descriptor() ([]byte, []int) {
	return file_api_operation_v1_operation_proto_rawDescGZIP(), []int{1}
}

func (x *AuditReviewReply) GetReviewID() int64 {
	if x != nil {
		return x.ReviewID
	}
	return 0
}

func (x *AuditReviewReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

// 审核申诉请求
type AuditAppealRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppealID      int64                  `protobuf:"varint,1,opt,name=appealID,proto3" json:"appealID,omitempty"`
	Status        int32                  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	OpUser        string                 `protobuf:"bytes,3,opt,name=opUser,proto3" json:"opUser,omitempty"`
	OpReason      string                 `protobuf:"bytes,4,opt,name=opReason,proto3" json:"opReason,omitempty"`
	OpRemarks     *string                `protobuf:"bytes,5,opt,name=opRemarks,proto3,oneof" json:"opRemarks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuditAppealRequest) Reset() {
	*x = AuditAppealRequest{}
	mi := &file_api_operation_v1_operation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditAppealRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditAppealRequest) ProtoMessage() {}

func (x *AuditAppealRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_operation_v1_operation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditAppealRequest.ProtoReflect.Descriptor instead.
func (*AuditAppealRequest) Descriptor() ([]byte, []int) {
	return file_api_operation_v1_operation_proto_rawDescGZIP(), []int{2}
}

func (x *AuditAppealRequest) GetAppealID() int64 {
	if x != nil {
		return x.AppealID
	}
	return 0
}

func (x *AuditAppealRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AuditAppealRequest) GetOpUser() string {
	if x != nil {
		return x.OpUser
	}
	return ""
}

func (x *AuditAppealRequest) GetOpReason() string {
	if x != nil {
		return x.OpReason
	}
	return ""
}

func (x *AuditAppealRequest) GetOpRemarks() string {
	if x != nil && x.OpRemarks != nil {
		return *x.OpRemarks
	}
	return ""
}

// 审核申诉回复
type AuditAppealReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppealID      int64                  `protobuf:"varint,1,opt,name=appealID,proto3" json:"appealID,omitempty"`
	Status        int32                  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuditAppealReply) Reset() {
	*x = AuditAppealReply{}
	mi := &file_api_operation_v1_operation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditAppealReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditAppealReply) ProtoMessage() {}

func (x *AuditAppealReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_operation_v1_operation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditAppealReply.ProtoReflect.Descriptor instead.
func (*AuditAppealReply) Descriptor() ([]byte, []int) {
	return file_api_operation_v1_operation_proto_rawDescGZIP(), []int{3}
}

func (x *AuditAppealReply) GetAppealID() int64 {
	if x != nil {
		return x.AppealID
	}
	return 0
}

func (x *AuditAppealReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_api_operation_v1_operation_proto protoreflect.FileDescriptor

const file_api_operation_v1_operation_proto_rawDesc = "" +
	"\n" +
	" api/operation/v1/operation.proto\x12\x10api.operation.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x17validate/validate.proto\"\xd1\x01\n" +
	"\x12AuditReviewRequest\x12#\n" +
	"\breviewID\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\breviewID\x12\x1f\n" +
	"\x06status\x18\x02 \x01(\x05B\a\xfaB\x04\x1a\x02 \x00R\x06status\x12\x1f\n" +
	"\x06opUser\x18\x03 \x01(\tB\a\xfaB\x04r\x02\x10\x02R\x06opUser\x12#\n" +
	"\bopReason\x18\x04 \x01(\tB\a\xfaB\x04r\x02\x10\x02R\bopReason\x12!\n" +
	"\topRemarks\x18\x05 \x01(\tH\x00R\topRemarks\x88\x01\x01B\f\n" +
	"\n" +
	"_opRemarks\"F\n" +
	"\x10AuditReviewReply\x12\x1a\n" +
	"\breviewID\x18\x01 \x01(\x03R\breviewID\x12\x16\n" +
	"\x06status\x18\x02 \x01(\x05R\x06status\"\xd1\x01\n" +
	"\x12AuditAppealRequest\x12#\n" +
	"\bappealID\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\bappealID\x12\x1f\n" +
	"\x06status\x18\x02 \x01(\x05B\a\xfaB\x04\x1a\x02 \x00R\x06status\x12\x1f\n" +
	"\x06opUser\x18\x03 \x01(\tB\a\xfaB\x04r\x02\x10\x02R\x06opUser\x12#\n" +
	"\bopReason\x18\x04 \x01(\tB\a\xfaB\x04r\x02\x10\x02R\bopReason\x12!\n" +
	"\topRemarks\x18\x05 \x01(\tH\x00R\topRemarks\x88\x01\x01B\f\n" +
	"\n" +
	"_opRemarks\"F\n" +
	"\x10AuditAppealReply\x12\x1a\n" +
	"\bappealID\x18\x01 \x01(\x03R\bappealID\x12\x16\n" +
	"\x06status\x18\x02 \x01(\x05R\x06status2\xfb\x01\n" +
	"\tOperation\x12v\n" +
	"\vAuditReview\x12$.api.operation.v1.AuditReviewRequest\x1a\".api.operation.v1.AuditReviewReply\"\x1d\x82\xd3\xe4\x93\x02\x17:\x01*\"\x12/o/v1/review/audit\x12v\n" +
	"\vAuditAppeal\x12$.api.operation.v1.AuditAppealRequest\x1a\".api.operation.v1.AuditAppealReply\"\x1d\x82\xd3\xe4\x93\x02\x17:\x01*\"\x12/o/v1/appeal/auditB3\n" +
	"\x10api.operation.v1P\x01Z\x1doperation/api/operation/v1;v1b\x06proto3"

var (
	file_api_operation_v1_operation_proto_rawDescOnce sync.Once
	file_api_operation_v1_operation_proto_rawDescData []byte
)

func file_api_operation_v1_operation_proto_rawDescGZIP() []byte {
	file_api_operation_v1_operation_proto_rawDescOnce.Do(func() {
		file_api_operation_v1_operation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_operation_v1_operation_proto_rawDesc), len(file_api_operation_v1_operation_proto_rawDesc)))
	})
	return file_api_operation_v1_operation_proto_rawDescData
}

var file_api_operation_v1_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_operation_v1_operation_proto_goTypes = []any{
	(*AuditReviewRequest)(nil), // 0: api.operation.v1.AuditReviewRequest
	(*AuditReviewReply)(nil),   // 1: api.operation.v1.AuditReviewReply
	(*AuditAppealRequest)(nil), // 2: api.operation.v1.AuditAppealRequest
	(*AuditAppealReply)(nil),   // 3: api.operation.v1.AuditAppealReply
}
var file_api_operation_v1_operation_proto_depIdxs = []int32{
	0, // 0: api.operation.v1.Operation.AuditReview:input_type -> api.operation.v1.AuditReviewRequest
	2, // 1: api.operation.v1.Operation.AuditAppeal:input_type -> api.operation.v1.AuditAppealRequest
	1, // 2: api.operation.v1.Operation.AuditReview:output_type -> api.operation.v1.AuditReviewReply
	3, // 3: api.operation.v1.Operation.AuditAppeal:output_type -> api.operation.v1.AuditAppealReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_operation_v1_operation_proto_init() }
func file_api_operation_v1_operation_proto_init() {
	if File_api_operation_v1_operation_proto != nil {
		return
	}
	file_api_operation_v1_operation_proto_msgTypes[0].OneofWrappers = []any{}
	file_api_operation_v1_operation_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_operation_v1_operation_proto_rawDesc), len(file_api_operation_v1_operation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_operation_v1_operation_proto_goTypes,
		DependencyIndexes: file_api_operation_v1_operation_proto_depIdxs,
		MessageInfos:      file_api_operation_v1_operation_proto_msgTypes,
	}.Build()
	File_api_operation_v1_operation_proto = out.File
	file_api_operation_v1_operation_proto_goTypes = nil
	file_api_operation_v1_operation_proto_depIdxs = nil
}
