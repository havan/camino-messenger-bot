// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: types/traveller.proto

package types

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

// Traveller Type
type TravellerType int32

const (
	TravellerType_UNSPECIFIED_TRAVELLER_TYPE TravellerType = 0
	TravellerType_ADULT                      TravellerType = 1
	TravellerType_CHILD                      TravellerType = 2
	TravellerType_INFANT                     TravellerType = 3
	TravellerType_SENIOR                     TravellerType = 4
)

// Enum value maps for TravellerType.
var (
	TravellerType_name = map[int32]string{
		0: "UNSPECIFIED_TRAVELLER_TYPE",
		1: "ADULT",
		2: "CHILD",
		3: "INFANT",
		4: "SENIOR",
	}
	TravellerType_value = map[string]int32{
		"UNSPECIFIED_TRAVELLER_TYPE": 0,
		"ADULT":                      1,
		"CHILD":                      2,
		"INFANT":                     3,
		"SENIOR":                     4,
	}
)

func (x TravellerType) Enum() *TravellerType {
	p := new(TravellerType)
	*p = x
	return p
}

func (x TravellerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TravellerType) Descriptor() protoreflect.EnumDescriptor {
	return file_types_traveller_proto_enumTypes[0].Descriptor()
}

func (TravellerType) Type() protoreflect.EnumType {
	return &file_types_traveller_proto_enumTypes[0]
}

func (x TravellerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TravellerType.Descriptor instead.
func (TravellerType) EnumDescriptor() ([]byte, []int) {
	return file_types_traveller_proto_rawDescGZIP(), []int{0}
}

// Traveller
type Traveller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Traveller type from enum below
	Type TravellerType `protobuf:"varint,1,opt,name=type,proto3,enum=types.TravellerType" json:"type,omitempty"`
	// Age
	// FIXME: Do we need both age & birthdate? (Protobuf fields are optional, so
	// we can keep it also)
	Age int32 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	// Birthdate
	Birthdate *Date `protobuf:"bytes,3,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	// FIXME: An enum for this field will reduce the message size.
	Nationality string `protobuf:"bytes,4,opt,name=nationality,proto3" json:"nationality,omitempty"`
}

func (x *Traveller) Reset() {
	*x = Traveller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_traveller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Traveller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Traveller) ProtoMessage() {}

func (x *Traveller) ProtoReflect() protoreflect.Message {
	mi := &file_types_traveller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Traveller.ProtoReflect.Descriptor instead.
func (*Traveller) Descriptor() ([]byte, []int) {
	return file_types_traveller_proto_rawDescGZIP(), []int{0}
}

func (x *Traveller) GetType() TravellerType {
	if x != nil {
		return x.Type
	}
	return TravellerType_UNSPECIFIED_TRAVELLER_TYPE
}

func (x *Traveller) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Traveller) GetBirthdate() *Date {
	if x != nil {
		return x.Birthdate
	}
	return nil
}

func (x *Traveller) GetNationality() string {
	if x != nil {
		return x.Nationality
	}
	return ""
}

var File_types_traveller_proto protoreflect.FileDescriptor

var file_types_traveller_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x10,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x94, 0x01, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2a, 0x5d, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x56, 0x45, 0x4c, 0x4c, 0x45,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x55, 0x4c,
	0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x49, 0x4e, 0x46, 0x41, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45,
	0x4e, 0x49, 0x4f, 0x52, 0x10, 0x04, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x34, 0x74, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x2f, 0x63, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x2d, 0x62, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_traveller_proto_rawDescOnce sync.Once
	file_types_traveller_proto_rawDescData = file_types_traveller_proto_rawDesc
)

func file_types_traveller_proto_rawDescGZIP() []byte {
	file_types_traveller_proto_rawDescOnce.Do(func() {
		file_types_traveller_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_traveller_proto_rawDescData)
	})
	return file_types_traveller_proto_rawDescData
}

var file_types_traveller_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_traveller_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_traveller_proto_goTypes = []interface{}{
	(TravellerType)(0), // 0: types.TravellerType
	(*Traveller)(nil),  // 1: types.Traveller
	(*Date)(nil),       // 2: types.Date
}
var file_types_traveller_proto_depIdxs = []int32{
	0, // 0: types.Traveller.type:type_name -> types.TravellerType
	2, // 1: types.Traveller.birthdate:type_name -> types.Date
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_types_traveller_proto_init() }
func file_types_traveller_proto_init() {
	if File_types_traveller_proto != nil {
		return
	}
	file_types_date_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_types_traveller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Traveller); i {
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
			RawDescriptor: file_types_traveller_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_traveller_proto_goTypes,
		DependencyIndexes: file_types_traveller_proto_depIdxs,
		EnumInfos:         file_types_traveller_proto_enumTypes,
		MessageInfos:      file_types_traveller_proto_msgTypes,
	}.Build()
	File_types_traveller_proto = out.File
	file_types_traveller_proto_rawDesc = nil
	file_types_traveller_proto_goTypes = nil
	file_types_traveller_proto_depIdxs = nil
}