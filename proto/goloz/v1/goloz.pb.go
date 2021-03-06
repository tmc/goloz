// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/goloz/v1/goloz.proto

package v1

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

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goloz_v1_goloz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goloz_v1_goloz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_proto_goloz_v1_goloz_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Character *Character `protobuf:"bytes,1,opt,name=character,proto3" json:"character,omitempty"`
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goloz_v1_goloz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goloz_v1_goloz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_proto_goloz_v1_goloz_proto_rawDescGZIP(), []int{1}
}

func (x *SyncRequest) GetCharacter() *Character {
	if x != nil {
		return x.Character
	}
	return nil
}

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos         *Position `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`
	SpriteIndex int32     `protobuf:"varint,2,opt,name=sprite_index,json=spriteIndex,proto3" json:"sprite_index,omitempty"` // TODO: consider an enum
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goloz_v1_goloz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goloz_v1_goloz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_proto_goloz_v1_goloz_proto_rawDescGZIP(), []int{2}
}

func (x *Character) GetPos() *Position {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *Character) GetSpriteIndex() int32 {
	if x != nil {
		return x.SpriteIndex
	}
	return 0
}

type SyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Characters map[string]*Character `protobuf:"bytes,1,rep,name=characters,proto3" json:"characters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SyncResponse) Reset() {
	*x = SyncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goloz_v1_goloz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncResponse) ProtoMessage() {}

func (x *SyncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goloz_v1_goloz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncResponse.ProtoReflect.Descriptor instead.
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return file_proto_goloz_v1_goloz_proto_rawDescGZIP(), []int{3}
}

func (x *SyncResponse) GetCharacters() map[string]*Character {
	if x != nil {
		return x.Characters
	}
	return nil
}

var File_proto_goloz_v1_goloz_proto protoreflect.FileDescriptor

var file_proto_goloz_v1_goloz_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x6f, 0x7a, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x6f, 0x6c, 0x6f, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x6f,
	0x6c, 0x6f, 0x7a, 0x2e, 0x76, 0x31, 0x22, 0x26, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x40,
	0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a,
	0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6c, 0x6f, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x22, 0x54, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x03, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6c,
	0x6f, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03,
	0x70, 0x6f, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x70, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x70, 0x72, 0x69, 0x74,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xaa, 0x01, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f,
	0x6c, 0x6f, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x1a,
	0x52, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6c, 0x6f, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0x50, 0x0a, 0x11, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63,
	0x12, 0x15, 0x2e, 0x67, 0x6f, 0x6c, 0x6f, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6c, 0x6f, 0x7a, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6d, 0x63, 0x2f, 0x67, 0x6f, 0x6c, 0x6f, 0x7a, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x6f, 0x7a, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_goloz_v1_goloz_proto_rawDescOnce sync.Once
	file_proto_goloz_v1_goloz_proto_rawDescData = file_proto_goloz_v1_goloz_proto_rawDesc
)

func file_proto_goloz_v1_goloz_proto_rawDescGZIP() []byte {
	file_proto_goloz_v1_goloz_proto_rawDescOnce.Do(func() {
		file_proto_goloz_v1_goloz_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_goloz_v1_goloz_proto_rawDescData)
	})
	return file_proto_goloz_v1_goloz_proto_rawDescData
}

var file_proto_goloz_v1_goloz_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_goloz_v1_goloz_proto_goTypes = []interface{}{
	(*Position)(nil),     // 0: goloz.v1.Position
	(*SyncRequest)(nil),  // 1: goloz.v1.SyncRequest
	(*Character)(nil),    // 2: goloz.v1.Character
	(*SyncResponse)(nil), // 3: goloz.v1.SyncResponse
	nil,                  // 4: goloz.v1.SyncResponse.CharactersEntry
}
var file_proto_goloz_v1_goloz_proto_depIdxs = []int32{
	2, // 0: goloz.v1.SyncRequest.character:type_name -> goloz.v1.Character
	0, // 1: goloz.v1.Character.pos:type_name -> goloz.v1.Position
	4, // 2: goloz.v1.SyncResponse.characters:type_name -> goloz.v1.SyncResponse.CharactersEntry
	2, // 3: goloz.v1.SyncResponse.CharactersEntry.value:type_name -> goloz.v1.Character
	1, // 4: goloz.v1.GameServerService.Sync:input_type -> goloz.v1.SyncRequest
	3, // 5: goloz.v1.GameServerService.Sync:output_type -> goloz.v1.SyncResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_goloz_v1_goloz_proto_init() }
func file_proto_goloz_v1_goloz_proto_init() {
	if File_proto_goloz_v1_goloz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_goloz_v1_goloz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_proto_goloz_v1_goloz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
		file_proto_goloz_v1_goloz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
		file_proto_goloz_v1_goloz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncResponse); i {
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
			RawDescriptor: file_proto_goloz_v1_goloz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_goloz_v1_goloz_proto_goTypes,
		DependencyIndexes: file_proto_goloz_v1_goloz_proto_depIdxs,
		MessageInfos:      file_proto_goloz_v1_goloz_proto_msgTypes,
	}.Build()
	File_proto_goloz_v1_goloz_proto = out.File
	file_proto_goloz_v1_goloz_proto_rawDesc = nil
	file_proto_goloz_v1_goloz_proto_goTypes = nil
	file_proto_goloz_v1_goloz_proto_depIdxs = nil
}
