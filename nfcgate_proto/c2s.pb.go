// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: nfcgate_proto/c2s.proto

package nfcgate_proto

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

type ServerData_Opcode int32

const (
	ServerData_OP_PSH ServerData_Opcode = 0
	ServerData_OP_SYN ServerData_Opcode = 1
	ServerData_OP_ACK ServerData_Opcode = 2
	ServerData_OP_FIN ServerData_Opcode = 3
)

// Enum value maps for ServerData_Opcode.
var (
	ServerData_Opcode_name = map[int32]string{
		0: "OP_PSH",
		1: "OP_SYN",
		2: "OP_ACK",
		3: "OP_FIN",
	}
	ServerData_Opcode_value = map[string]int32{
		"OP_PSH": 0,
		"OP_SYN": 1,
		"OP_ACK": 2,
		"OP_FIN": 3,
	}
)

func (x ServerData_Opcode) Enum() *ServerData_Opcode {
	p := new(ServerData_Opcode)
	*p = x
	return p
}

func (x ServerData_Opcode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerData_Opcode) Descriptor() protoreflect.EnumDescriptor {
	return file_nfcgate_proto_c2s_proto_enumTypes[0].Descriptor()
}

func (ServerData_Opcode) Type() protoreflect.EnumType {
	return &file_nfcgate_proto_c2s_proto_enumTypes[0]
}

func (x ServerData_Opcode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerData_Opcode.Descriptor instead.
func (ServerData_Opcode) EnumDescriptor() ([]byte, []int) {
	return file_nfcgate_proto_c2s_proto_rawDescGZIP(), []int{0, 0}
}

// Session management messages
type ServerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describes intent of the sender
	Opcode ServerData_Opcode `protobuf:"varint,1,opt,name=opcode,proto3,enum=de.tu_darmstadt.seemoo.nfcgate.network.c2s.ServerData_Opcode" json:"opcode,omitempty"`
	// Binary blob containing the message content
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ServerData) Reset() {
	*x = ServerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nfcgate_proto_c2s_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerData) ProtoMessage() {}

func (x *ServerData) ProtoReflect() protoreflect.Message {
	mi := &file_nfcgate_proto_c2s_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerData.ProtoReflect.Descriptor instead.
func (*ServerData) Descriptor() ([]byte, []int) {
	return file_nfcgate_proto_c2s_proto_rawDescGZIP(), []int{0}
}

func (x *ServerData) GetOpcode() ServerData_Opcode {
	if x != nil {
		return x.Opcode
	}
	return ServerData_OP_PSH
}

func (x *ServerData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_nfcgate_proto_c2s_proto protoreflect.FileDescriptor

var file_nfcgate_proto_c2s_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6e, 0x66, 0x63, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x32, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x64, 0x65, 0x2e, 0x74, 0x75,
	0x5f, 0x64, 0x61, 0x72, 0x6d, 0x73, 0x74, 0x61, 0x64, 0x74, 0x2e, 0x73, 0x65, 0x65, 0x6d, 0x6f,
	0x6f, 0x2e, 0x6e, 0x66, 0x63, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x63, 0x32, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x55, 0x0a, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x64, 0x65, 0x2e, 0x74, 0x75, 0x5f, 0x64, 0x61, 0x72,
	0x6d, 0x73, 0x74, 0x61, 0x64, 0x74, 0x2e, 0x73, 0x65, 0x65, 0x6d, 0x6f, 0x6f, 0x2e, 0x6e, 0x66,
	0x63, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x32,
	0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x70, 0x63,
	0x6f, 0x64, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x38, 0x0a, 0x06, 0x4f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x50, 0x5f,
	0x50, 0x53, 0x48, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x50, 0x5f, 0x53, 0x59, 0x4e, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x50, 0x5f, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x0a, 0x0a,
	0x06, 0x4f, 0x50, 0x5f, 0x46, 0x49, 0x4e, 0x10, 0x03, 0x42, 0x10, 0x5a, 0x0e, 0x6e, 0x66, 0x63,
	0x67, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_nfcgate_proto_c2s_proto_rawDescOnce sync.Once
	file_nfcgate_proto_c2s_proto_rawDescData = file_nfcgate_proto_c2s_proto_rawDesc
)

func file_nfcgate_proto_c2s_proto_rawDescGZIP() []byte {
	file_nfcgate_proto_c2s_proto_rawDescOnce.Do(func() {
		file_nfcgate_proto_c2s_proto_rawDescData = protoimpl.X.CompressGZIP(file_nfcgate_proto_c2s_proto_rawDescData)
	})
	return file_nfcgate_proto_c2s_proto_rawDescData
}

var file_nfcgate_proto_c2s_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nfcgate_proto_c2s_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nfcgate_proto_c2s_proto_goTypes = []interface{}{
	(ServerData_Opcode)(0), // 0: de.tu_darmstadt.seemoo.nfcgate.network.c2s.ServerData.Opcode
	(*ServerData)(nil),     // 1: de.tu_darmstadt.seemoo.nfcgate.network.c2s.ServerData
}
var file_nfcgate_proto_c2s_proto_depIdxs = []int32{
	0, // 0: de.tu_darmstadt.seemoo.nfcgate.network.c2s.ServerData.opcode:type_name -> de.tu_darmstadt.seemoo.nfcgate.network.c2s.ServerData.Opcode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nfcgate_proto_c2s_proto_init() }
func file_nfcgate_proto_c2s_proto_init() {
	if File_nfcgate_proto_c2s_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nfcgate_proto_c2s_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerData); i {
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
			RawDescriptor: file_nfcgate_proto_c2s_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nfcgate_proto_c2s_proto_goTypes,
		DependencyIndexes: file_nfcgate_proto_c2s_proto_depIdxs,
		EnumInfos:         file_nfcgate_proto_c2s_proto_enumTypes,
		MessageInfos:      file_nfcgate_proto_c2s_proto_msgTypes,
	}.Build()
	File_nfcgate_proto_c2s_proto = out.File
	file_nfcgate_proto_c2s_proto_rawDesc = nil
	file_nfcgate_proto_c2s_proto_goTypes = nil
	file_nfcgate_proto_c2s_proto_depIdxs = nil
}
