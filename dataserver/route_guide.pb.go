// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: routeguide/route_guide.proto

package routeguide

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

// A c2sData is a message from the client to the server.
type C2SData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind      int32  `protobuf:"varint,1,opt,name=kind,proto3" json:"kind,omitempty"`
	StrData   string `protobuf:"bytes,2,opt,name=str_data,json=strData,proto3" json:"str_data,omitempty"`
	BytesData []byte `protobuf:"bytes,3,opt,name=bytes_data,json=bytesData,proto3" json:"bytes_data,omitempty"`
}

func (x *C2SData) Reset() {
	*x = C2SData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SData) ProtoMessage() {}

func (x *C2SData) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SData.ProtoReflect.Descriptor instead.
func (*C2SData) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{0}
}

func (x *C2SData) GetKind() int32 {
	if x != nil {
		return x.Kind
	}
	return 0
}

func (x *C2SData) GetStrData() string {
	if x != nil {
		return x.StrData
	}
	return ""
}

func (x *C2SData) GetBytesData() []byte {
	if x != nil {
		return x.BytesData
	}
	return nil
}

// A s2cData is a message from the server to the client.
type S2CData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind      int32  `protobuf:"varint,1,opt,name=kind,proto3" json:"kind,omitempty"`
	StrData   string `protobuf:"bytes,2,opt,name=str_data,json=strData,proto3" json:"str_data,omitempty"`
	BytesData []byte `protobuf:"bytes,3,opt,name=bytes_data,json=bytesData,proto3" json:"bytes_data,omitempty"`
}

func (x *S2CData) Reset() {
	*x = S2CData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CData) ProtoMessage() {}

func (x *S2CData) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CData.ProtoReflect.Descriptor instead.
func (*S2CData) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{1}
}

func (x *S2CData) GetKind() int32 {
	if x != nil {
		return x.Kind
	}
	return 0
}

func (x *S2CData) GetStrData() string {
	if x != nil {
		return x.StrData
	}
	return ""
}

func (x *S2CData) GetBytesData() []byte {
	if x != nil {
		return x.BytesData
	}
	return nil
}

var File_routeguide_route_guide_proto protoreflect.FileDescriptor

var file_routeguide_route_guide_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x22, 0x57, 0x0a, 0x07, 0x63, 0x32,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x72,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x57, 0x0a, 0x07, 0x73, 0x32, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x32, 0x89, 0x01, 0x0a,
	0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x13, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x63, 0x32, 0x73, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x13, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x32, 0x63,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x61, 0x74, 0x44, 0x42, 0x12, 0x13, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x63, 0x32, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x13, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x73,
	0x32, 0x63, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x42, 0x21, 0x42, 0x0f, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0c, 0x2e,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_routeguide_route_guide_proto_rawDescOnce sync.Once
	file_routeguide_route_guide_proto_rawDescData = file_routeguide_route_guide_proto_rawDesc
)

func file_routeguide_route_guide_proto_rawDescGZIP() []byte {
	file_routeguide_route_guide_proto_rawDescOnce.Do(func() {
		file_routeguide_route_guide_proto_rawDescData = protoimpl.X.CompressGZIP(file_routeguide_route_guide_proto_rawDescData)
	})
	return file_routeguide_route_guide_proto_rawDescData
}

var file_routeguide_route_guide_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_routeguide_route_guide_proto_goTypes = []interface{}{
	(*C2SData)(nil), // 0: routeguide.c2sData
	(*S2CData)(nil), // 1: routeguide.s2cData
}
var file_routeguide_route_guide_proto_depIdxs = []int32{
	0, // 0: routeguide.RouteGuide.ProcessData:input_type -> routeguide.c2sData
	0, // 1: routeguide.RouteGuide.GetLatestCatDB:input_type -> routeguide.c2sData
	1, // 2: routeguide.RouteGuide.ProcessData:output_type -> routeguide.s2cData
	1, // 3: routeguide.RouteGuide.GetLatestCatDB:output_type -> routeguide.s2cData
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_routeguide_route_guide_proto_init() }
func file_routeguide_route_guide_proto_init() {
	if File_routeguide_route_guide_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routeguide_route_guide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SData); i {
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
		file_routeguide_route_guide_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2CData); i {
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
			RawDescriptor: file_routeguide_route_guide_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routeguide_route_guide_proto_goTypes,
		DependencyIndexes: file_routeguide_route_guide_proto_depIdxs,
		MessageInfos:      file_routeguide_route_guide_proto_msgTypes,
	}.Build()
	File_routeguide_route_guide_proto = out.File
	file_routeguide_route_guide_proto_rawDesc = nil
	file_routeguide_route_guide_proto_goTypes = nil
	file_routeguide_route_guide_proto_depIdxs = nil
}
