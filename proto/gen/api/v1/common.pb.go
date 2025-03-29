// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1/common.proto

package apiv1

import (
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

type State int32

const (
	State_STATE_UNSPECIFIED State = 0
	State_NORMAL            State = 1
	State_ARCHIVED          State = 2
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "NORMAL",
		2: "ARCHIVED",
	}
	State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"NORMAL":            1,
		"ARCHIVED":          2,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_common_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_api_v1_common_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_common_proto_rawDescGZIP(), []int{0}
}

type Direction int32

const (
	Direction_DIRECTION_UNSPECIFIED Direction = 0
	Direction_ASC                   Direction = 1
	Direction_DESC                  Direction = 2
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "DIRECTION_UNSPECIFIED",
		1: "ASC",
		2: "DESC",
	}
	Direction_value = map[string]int32{
		"DIRECTION_UNSPECIFIED": 0,
		"ASC":                   1,
		"DESC":                  2,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_common_proto_enumTypes[1].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_api_v1_common_proto_enumTypes[1]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_common_proto_rawDescGZIP(), []int{1}
}

// Used internally for obfuscating the page token.
type PageToken struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int32                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PageToken) Reset() {
	*x = PageToken{}
	mi := &file_api_v1_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PageToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageToken) ProtoMessage() {}

func (x *PageToken) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageToken.ProtoReflect.Descriptor instead.
func (*PageToken) Descriptor() ([]byte, []int) {
	return file_api_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *PageToken) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PageToken) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_api_v1_common_proto protoreflect.FileDescriptor

const file_api_v1_common_proto_rawDesc = "" +
	"\n" +
	"\x13api/v1/common.proto\x12\fmemos.api.v1\"9\n" +
	"\tPageToken\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\x05R\x06offset*8\n" +
	"\x05State\x12\x15\n" +
	"\x11STATE_UNSPECIFIED\x10\x00\x12\n" +
	"\n" +
	"\x06NORMAL\x10\x01\x12\f\n" +
	"\bARCHIVED\x10\x02*9\n" +
	"\tDirection\x12\x19\n" +
	"\x15DIRECTION_UNSPECIFIED\x10\x00\x12\a\n" +
	"\x03ASC\x10\x01\x12\b\n" +
	"\x04DESC\x10\x02B\xa3\x01\n" +
	"\x10com.memos.api.v1B\vCommonProtoP\x01Z0github.com/usememos/memos/proto/gen/api/v1;apiv1\xa2\x02\x03MAX\xaa\x02\fMemos.Api.V1\xca\x02\fMemos\\Api\\V1\xe2\x02\x18Memos\\Api\\V1\\GPBMetadata\xea\x02\x0eMemos::Api::V1b\x06proto3"

var (
	file_api_v1_common_proto_rawDescOnce sync.Once
	file_api_v1_common_proto_rawDescData []byte
)

func file_api_v1_common_proto_rawDescGZIP() []byte {
	file_api_v1_common_proto_rawDescOnce.Do(func() {
		file_api_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_common_proto_rawDesc), len(file_api_v1_common_proto_rawDesc)))
	})
	return file_api_v1_common_proto_rawDescData
}

var file_api_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_common_proto_goTypes = []any{
	(State)(0),        // 0: memos.api.v1.State
	(Direction)(0),    // 1: memos.api.v1.Direction
	(*PageToken)(nil), // 2: memos.api.v1.PageToken
}
var file_api_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_common_proto_init() }
func file_api_v1_common_proto_init() {
	if File_api_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_common_proto_rawDesc), len(file_api_v1_common_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_common_proto_goTypes,
		DependencyIndexes: file_api_v1_common_proto_depIdxs,
		EnumInfos:         file_api_v1_common_proto_enumTypes,
		MessageInfos:      file_api_v1_common_proto_msgTypes,
	}.Build()
	File_api_v1_common_proto = out.File
	file_api_v1_common_proto_goTypes = nil
	file_api_v1_common_proto_depIdxs = nil
}
