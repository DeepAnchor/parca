// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: debuginfo/debuginfo.proto

package debuginfo

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

type DebugInfoExistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *DebugInfoExistsRequest) Reset() {
	*x = DebugInfoExistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debuginfo_debuginfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugInfoExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugInfoExistsRequest) ProtoMessage() {}

func (x *DebugInfoExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_debuginfo_debuginfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugInfoExistsRequest.ProtoReflect.Descriptor instead.
func (*DebugInfoExistsRequest) Descriptor() ([]byte, []int) {
	return file_debuginfo_debuginfo_proto_rawDescGZIP(), []int{0}
}

func (x *DebugInfoExistsRequest) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

type DebugInfoExistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *DebugInfoExistsResponse) Reset() {
	*x = DebugInfoExistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debuginfo_debuginfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugInfoExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugInfoExistsResponse) ProtoMessage() {}

func (x *DebugInfoExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_debuginfo_debuginfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugInfoExistsResponse.ProtoReflect.Descriptor instead.
func (*DebugInfoExistsResponse) Descriptor() ([]byte, []int) {
	return file_debuginfo_debuginfo_proto_rawDescGZIP(), []int{1}
}

func (x *DebugInfoExistsResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type DebugInfoUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*DebugInfoUploadRequest_Info
	//	*DebugInfoUploadRequest_ChunkData
	Data isDebugInfoUploadRequest_Data `protobuf_oneof:"data"`
}

func (x *DebugInfoUploadRequest) Reset() {
	*x = DebugInfoUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debuginfo_debuginfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugInfoUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugInfoUploadRequest) ProtoMessage() {}

func (x *DebugInfoUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_debuginfo_debuginfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugInfoUploadRequest.ProtoReflect.Descriptor instead.
func (*DebugInfoUploadRequest) Descriptor() ([]byte, []int) {
	return file_debuginfo_debuginfo_proto_rawDescGZIP(), []int{2}
}

func (m *DebugInfoUploadRequest) GetData() isDebugInfoUploadRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *DebugInfoUploadRequest) GetInfo() *DebugInfoUploadInfo {
	if x, ok := x.GetData().(*DebugInfoUploadRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *DebugInfoUploadRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*DebugInfoUploadRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isDebugInfoUploadRequest_Data interface {
	isDebugInfoUploadRequest_Data()
}

type DebugInfoUploadRequest_Info struct {
	Info *DebugInfoUploadInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type DebugInfoUploadRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*DebugInfoUploadRequest_Info) isDebugInfoUploadRequest_Data() {}

func (*DebugInfoUploadRequest_ChunkData) isDebugInfoUploadRequest_Data() {}

type DebugInfoUploadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *DebugInfoUploadInfo) Reset() {
	*x = DebugInfoUploadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debuginfo_debuginfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugInfoUploadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugInfoUploadInfo) ProtoMessage() {}

func (x *DebugInfoUploadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_debuginfo_debuginfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugInfoUploadInfo.ProtoReflect.Descriptor instead.
func (*DebugInfoUploadInfo) Descriptor() ([]byte, []int) {
	return file_debuginfo_debuginfo_proto_rawDescGZIP(), []int{3}
}

func (x *DebugInfoUploadInfo) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

type DebugInfoUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	Size    uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *DebugInfoUploadResponse) Reset() {
	*x = DebugInfoUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debuginfo_debuginfo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugInfoUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugInfoUploadResponse) ProtoMessage() {}

func (x *DebugInfoUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_debuginfo_debuginfo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugInfoUploadResponse.ProtoReflect.Descriptor instead.
func (*DebugInfoUploadResponse) Descriptor() ([]byte, []int) {
	return file_debuginfo_debuginfo_proto_rawDescGZIP(), []int{4}
}

func (x *DebugInfoUploadResponse) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

func (x *DebugInfoUploadResponse) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_debuginfo_debuginfo_proto protoreflect.FileDescriptor

var file_debuginfo_debuginfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x65, 0x62, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x61, 0x72,
	0x63, 0x61, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x33, 0x0a, 0x16,
	0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x64, 0x22, 0x31, 0x0a, 0x17, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x22, 0x7d, 0x0a, 0x16, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70,
	0x61, 0x72, 0x63, 0x61, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x13, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x17, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x32,
	0xcb, 0x01, 0x0a, 0x09, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5d, 0x0a,
	0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x06,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x27, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x63,
	0x61, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x69, 0x6e,
	0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_debuginfo_debuginfo_proto_rawDescOnce sync.Once
	file_debuginfo_debuginfo_proto_rawDescData = file_debuginfo_debuginfo_proto_rawDesc
)

func file_debuginfo_debuginfo_proto_rawDescGZIP() []byte {
	file_debuginfo_debuginfo_proto_rawDescOnce.Do(func() {
		file_debuginfo_debuginfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_debuginfo_debuginfo_proto_rawDescData)
	})
	return file_debuginfo_debuginfo_proto_rawDescData
}

var file_debuginfo_debuginfo_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_debuginfo_debuginfo_proto_goTypes = []interface{}{
	(*DebugInfoExistsRequest)(nil),  // 0: parca.debuginfo.DebugInfoExistsRequest
	(*DebugInfoExistsResponse)(nil), // 1: parca.debuginfo.DebugInfoExistsResponse
	(*DebugInfoUploadRequest)(nil),  // 2: parca.debuginfo.DebugInfoUploadRequest
	(*DebugInfoUploadInfo)(nil),     // 3: parca.debuginfo.DebugInfoUploadInfo
	(*DebugInfoUploadResponse)(nil), // 4: parca.debuginfo.DebugInfoUploadResponse
}
var file_debuginfo_debuginfo_proto_depIdxs = []int32{
	3, // 0: parca.debuginfo.DebugInfoUploadRequest.info:type_name -> parca.debuginfo.DebugInfoUploadInfo
	0, // 1: parca.debuginfo.DebugInfo.Exists:input_type -> parca.debuginfo.DebugInfoExistsRequest
	2, // 2: parca.debuginfo.DebugInfo.Upload:input_type -> parca.debuginfo.DebugInfoUploadRequest
	1, // 3: parca.debuginfo.DebugInfo.Exists:output_type -> parca.debuginfo.DebugInfoExistsResponse
	4, // 4: parca.debuginfo.DebugInfo.Upload:output_type -> parca.debuginfo.DebugInfoUploadResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_debuginfo_debuginfo_proto_init() }
func file_debuginfo_debuginfo_proto_init() {
	if File_debuginfo_debuginfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_debuginfo_debuginfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugInfoExistsRequest); i {
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
		file_debuginfo_debuginfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugInfoExistsResponse); i {
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
		file_debuginfo_debuginfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugInfoUploadRequest); i {
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
		file_debuginfo_debuginfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugInfoUploadInfo); i {
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
		file_debuginfo_debuginfo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugInfoUploadResponse); i {
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
	file_debuginfo_debuginfo_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*DebugInfoUploadRequest_Info)(nil),
		(*DebugInfoUploadRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_debuginfo_debuginfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_debuginfo_debuginfo_proto_goTypes,
		DependencyIndexes: file_debuginfo_debuginfo_proto_depIdxs,
		MessageInfos:      file_debuginfo_debuginfo_proto_msgTypes,
	}.Build()
	File_debuginfo_debuginfo_proto = out.File
	file_debuginfo_debuginfo_proto_rawDesc = nil
	file_debuginfo_debuginfo_proto_goTypes = nil
	file_debuginfo_debuginfo_proto_depIdxs = nil
}