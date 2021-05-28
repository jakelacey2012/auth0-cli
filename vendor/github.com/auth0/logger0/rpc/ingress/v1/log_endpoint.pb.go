// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: rpc/ingress/v1/log_endpoint.proto

package ingressv1

import (
	logger0 "github.com/auth0/logger0"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PostLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant   string                 `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Type     logger0.LogRecord_Type `protobuf:"varint,2,opt,name=type,proto3,enum=auth0.logger0.LogRecord_Type" json:"type,omitempty"`
	Metadata map[string]string      `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Messages [][]byte               `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *PostLogRequest) Reset() {
	*x = PostLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_ingress_v1_log_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostLogRequest) ProtoMessage() {}

func (x *PostLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_v1_log_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostLogRequest.ProtoReflect.Descriptor instead.
func (*PostLogRequest) Descriptor() ([]byte, []int) {
	return file_rpc_ingress_v1_log_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *PostLogRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *PostLogRequest) GetType() logger0.LogRecord_Type {
	if x != nil {
		return x.Type
	}
	return logger0.LogRecord_TYPE_UNSPECIFIED
}

func (x *PostLogRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PostLogRequest) GetMessages() [][]byte {
	if x != nil {
		return x.Messages
	}
	return nil
}

type PostLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostLogResponse) Reset() {
	*x = PostLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_ingress_v1_log_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostLogResponse) ProtoMessage() {}

func (x *PostLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ingress_v1_log_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostLogResponse.ProtoReflect.Descriptor instead.
func (*PostLogResponse) Descriptor() ([]byte, []int) {
	return file_rpc_ingress_v1_log_endpoint_proto_rawDescGZIP(), []int{1}
}

var File_rpc_ingress_v1_log_endpoint_proto protoreflect.FileDescriptor

var file_rpc_ingress_v1_log_endpoint_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x30, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x70, 0x69,
	0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88,
	0x02, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2e,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x52, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x3b, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x11, 0x0a, 0x0f, 0x50, 0x6f, 0x73,
	0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x7f, 0x0a, 0x0b,
	0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x70, 0x0a, 0x04, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x28, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x30, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x30, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x30, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_ingress_v1_log_endpoint_proto_rawDescOnce sync.Once
	file_rpc_ingress_v1_log_endpoint_proto_rawDescData = file_rpc_ingress_v1_log_endpoint_proto_rawDesc
)

func file_rpc_ingress_v1_log_endpoint_proto_rawDescGZIP() []byte {
	file_rpc_ingress_v1_log_endpoint_proto_rawDescOnce.Do(func() {
		file_rpc_ingress_v1_log_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_ingress_v1_log_endpoint_proto_rawDescData)
	})
	return file_rpc_ingress_v1_log_endpoint_proto_rawDescData
}

var file_rpc_ingress_v1_log_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_ingress_v1_log_endpoint_proto_goTypes = []interface{}{
	(*PostLogRequest)(nil),      // 0: auth0.logger0.ingress.v1.PostLogRequest
	(*PostLogResponse)(nil),     // 1: auth0.logger0.ingress.v1.PostLogResponse
	nil,                         // 2: auth0.logger0.ingress.v1.PostLogRequest.MetadataEntry
	(logger0.LogRecord_Type)(0), // 3: auth0.logger0.LogRecord.Type
}
var file_rpc_ingress_v1_log_endpoint_proto_depIdxs = []int32{
	3, // 0: auth0.logger0.ingress.v1.PostLogRequest.type:type_name -> auth0.logger0.LogRecord.Type
	2, // 1: auth0.logger0.ingress.v1.PostLogRequest.metadata:type_name -> auth0.logger0.ingress.v1.PostLogRequest.MetadataEntry
	0, // 2: auth0.logger0.ingress.v1.LogEndpoint.Post:input_type -> auth0.logger0.ingress.v1.PostLogRequest
	1, // 3: auth0.logger0.ingress.v1.LogEndpoint.Post:output_type -> auth0.logger0.ingress.v1.PostLogResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_ingress_v1_log_endpoint_proto_init() }
func file_rpc_ingress_v1_log_endpoint_proto_init() {
	if File_rpc_ingress_v1_log_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_ingress_v1_log_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostLogRequest); i {
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
		file_rpc_ingress_v1_log_endpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostLogResponse); i {
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
			RawDescriptor: file_rpc_ingress_v1_log_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_ingress_v1_log_endpoint_proto_goTypes,
		DependencyIndexes: file_rpc_ingress_v1_log_endpoint_proto_depIdxs,
		MessageInfos:      file_rpc_ingress_v1_log_endpoint_proto_msgTypes,
	}.Build()
	File_rpc_ingress_v1_log_endpoint_proto = out.File
	file_rpc_ingress_v1_log_endpoint_proto_rawDesc = nil
	file_rpc_ingress_v1_log_endpoint_proto_goTypes = nil
	file_rpc_ingress_v1_log_endpoint_proto_depIdxs = nil
}