// version of protocol buffer used

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: resize.proto

package proto

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

// argument
type ResizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// data type and position of data
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Chunk  []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Width  uint32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height uint32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *ResizeRequest) Reset() {
	*x = ResizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resize_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResizeRequest) ProtoMessage() {}

func (x *ResizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resize_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResizeRequest.ProtoReflect.Descriptor instead.
func (*ResizeRequest) Descriptor() ([]byte, []int) {
	return file_resize_proto_rawDescGZIP(), []int{0}
}

func (x *ResizeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResizeRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *ResizeRequest) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ResizeRequest) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

// return value
type ResizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// data type and position of data
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Original []byte `protobuf:"bytes,2,opt,name=original,proto3" json:"original,omitempty"`
	Chunk    []byte `protobuf:"bytes,3,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Width    uint32 `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
	Height   uint32 `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *ResizeResponse) Reset() {
	*x = ResizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resize_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResizeResponse) ProtoMessage() {}

func (x *ResizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resize_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResizeResponse.ProtoReflect.Descriptor instead.
func (*ResizeResponse) Descriptor() ([]byte, []int) {
	return file_resize_proto_rawDescGZIP(), []int{1}
}

func (x *ResizeResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResizeResponse) GetOriginal() []byte {
	if x != nil {
		return x.Original
	}
	return nil
}

func (x *ResizeResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *ResizeResponse) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ResizeResponse) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_resize_proto protoreflect.FileDescriptor

var file_resize_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x67, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22,
	0x84, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0x3f, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x35, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x7a,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resize_proto_rawDescOnce sync.Once
	file_resize_proto_rawDescData = file_resize_proto_rawDesc
)

func file_resize_proto_rawDescGZIP() []byte {
	file_resize_proto_rawDescOnce.Do(func() {
		file_resize_proto_rawDescData = protoimpl.X.CompressGZIP(file_resize_proto_rawDescData)
	})
	return file_resize_proto_rawDescData
}

var file_resize_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resize_proto_goTypes = []interface{}{
	(*ResizeRequest)(nil),  // 0: resize.ResizeRequest
	(*ResizeResponse)(nil), // 1: resize.ResizeResponse
}
var file_resize_proto_depIdxs = []int32{
	0, // 0: resize.Resize.Send:input_type -> resize.ResizeRequest
	1, // 1: resize.Resize.Send:output_type -> resize.ResizeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resize_proto_init() }
func file_resize_proto_init() {
	if File_resize_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resize_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResizeRequest); i {
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
		file_resize_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResizeResponse); i {
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
			RawDescriptor: file_resize_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resize_proto_goTypes,
		DependencyIndexes: file_resize_proto_depIdxs,
		MessageInfos:      file_resize_proto_msgTypes,
	}.Build()
	File_resize_proto = out.File
	file_resize_proto_rawDesc = nil
	file_resize_proto_goTypes = nil
	file_resize_proto_depIdxs = nil
}
