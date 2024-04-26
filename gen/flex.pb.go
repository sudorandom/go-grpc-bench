// Protobuf that flexes a lot of the types

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: flex.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the user's name.
type FlexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *ComplexType `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *FlexRequest) Reset() {
	*x = FlexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlexRequest) ProtoMessage() {}

func (x *FlexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlexRequest.ProtoReflect.Descriptor instead.
func (*FlexRequest) Descriptor() ([]byte, []int) {
	return file_flex_proto_rawDescGZIP(), []int{0}
}

func (x *FlexRequest) GetMsg() *ComplexType {
	if x != nil {
		return x.Msg
	}
	return nil
}

// The response message containing the greetings
type FlexReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *ComplexType `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *FlexReply) Reset() {
	*x = FlexReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlexReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlexReply) ProtoMessage() {}

func (x *FlexReply) ProtoReflect() protoreflect.Message {
	mi := &file_flex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlexReply.ProtoReflect.Descriptor instead.
func (*FlexReply) Descriptor() ([]byte, []int) {
	return file_flex_proto_rawDescGZIP(), []int{1}
}

func (x *FlexReply) GetMsg() *ComplexType {
	if x != nil {
		return x.Msg
	}
	return nil
}

// Type that has a bunch of different types
type ComplexType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoubleField      float64  `protobuf:"fixed64,1,opt,name=doubleField,proto3" json:"doubleField,omitempty"`
	FloatField       float32  `protobuf:"fixed32,2,opt,name=floatField,proto3" json:"floatField,omitempty"`
	Int32Field       int32    `protobuf:"varint,3,opt,name=int32Field,proto3" json:"int32Field,omitempty"`
	Int64Field       int64    `protobuf:"varint,4,opt,name=int64Field,proto3" json:"int64Field,omitempty"`
	Uint32Field      uint32   `protobuf:"varint,5,opt,name=uint32Field,proto3" json:"uint32Field,omitempty"`
	Uint64Field      uint64   `protobuf:"varint,6,opt,name=uint64Field,proto3" json:"uint64Field,omitempty"`
	Sint32Field      int32    `protobuf:"zigzag32,7,opt,name=sint32Field,proto3" json:"sint32Field,omitempty"`
	Sint64Field      int64    `protobuf:"zigzag64,8,opt,name=sint64Field,proto3" json:"sint64Field,omitempty"`
	Fixed32Field     uint32   `protobuf:"fixed32,9,opt,name=fixed32Field,proto3" json:"fixed32Field,omitempty"`
	Fixed64Field     uint64   `protobuf:"fixed64,10,opt,name=fixed64Field,proto3" json:"fixed64Field,omitempty"`
	Sfixed32Field    int32    `protobuf:"fixed32,11,opt,name=sfixed32Field,proto3" json:"sfixed32Field,omitempty"`
	Sfixed64Field    int64    `protobuf:"fixed64,12,opt,name=sfixed64Field,proto3" json:"sfixed64Field,omitempty"`
	BoolField        bool     `protobuf:"varint,13,opt,name=boolField,proto3" json:"boolField,omitempty"`
	StringField      string   `protobuf:"bytes,14,opt,name=stringField,proto3" json:"stringField,omitempty"`
	MsgField         *Other   `protobuf:"bytes,15,opt,name=msgField,proto3" json:"msgField,omitempty"`
	RepeatedMsgField []*Other `protobuf:"bytes,16,rep,name=repeatedMsgField,proto3" json:"repeatedMsgField,omitempty"`
	OptionalMsgField *Other   `protobuf:"bytes,17,opt,name=optionalMsgField,proto3,oneof" json:"optionalMsgField,omitempty"`
}

func (x *ComplexType) Reset() {
	*x = ComplexType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexType) ProtoMessage() {}

func (x *ComplexType) ProtoReflect() protoreflect.Message {
	mi := &file_flex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexType.ProtoReflect.Descriptor instead.
func (*ComplexType) Descriptor() ([]byte, []int) {
	return file_flex_proto_rawDescGZIP(), []int{2}
}

func (x *ComplexType) GetDoubleField() float64 {
	if x != nil {
		return x.DoubleField
	}
	return 0
}

func (x *ComplexType) GetFloatField() float32 {
	if x != nil {
		return x.FloatField
	}
	return 0
}

func (x *ComplexType) GetInt32Field() int32 {
	if x != nil {
		return x.Int32Field
	}
	return 0
}

func (x *ComplexType) GetInt64Field() int64 {
	if x != nil {
		return x.Int64Field
	}
	return 0
}

func (x *ComplexType) GetUint32Field() uint32 {
	if x != nil {
		return x.Uint32Field
	}
	return 0
}

func (x *ComplexType) GetUint64Field() uint64 {
	if x != nil {
		return x.Uint64Field
	}
	return 0
}

func (x *ComplexType) GetSint32Field() int32 {
	if x != nil {
		return x.Sint32Field
	}
	return 0
}

func (x *ComplexType) GetSint64Field() int64 {
	if x != nil {
		return x.Sint64Field
	}
	return 0
}

func (x *ComplexType) GetFixed32Field() uint32 {
	if x != nil {
		return x.Fixed32Field
	}
	return 0
}

func (x *ComplexType) GetFixed64Field() uint64 {
	if x != nil {
		return x.Fixed64Field
	}
	return 0
}

func (x *ComplexType) GetSfixed32Field() int32 {
	if x != nil {
		return x.Sfixed32Field
	}
	return 0
}

func (x *ComplexType) GetSfixed64Field() int64 {
	if x != nil {
		return x.Sfixed64Field
	}
	return 0
}

func (x *ComplexType) GetBoolField() bool {
	if x != nil {
		return x.BoolField
	}
	return false
}

func (x *ComplexType) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *ComplexType) GetMsgField() *Other {
	if x != nil {
		return x.MsgField
	}
	return nil
}

func (x *ComplexType) GetRepeatedMsgField() []*Other {
	if x != nil {
		return x.RepeatedMsgField
	}
	return nil
}

func (x *ComplexType) GetOptionalMsgField() *Other {
	if x != nil {
		return x.OptionalMsgField
	}
	return nil
}

type Other struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Other) Reset() {
	*x = Other{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flex_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Other) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Other) ProtoMessage() {}

func (x *Other) ProtoReflect() protoreflect.Message {
	mi := &file_flex_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Other.ProtoReflect.Descriptor instead.
func (*Other) Descriptor() ([]byte, []int) {
	return file_flex_proto_rawDescGZIP(), []int{3}
}

var File_flex_proto protoreflect.FileDescriptor

var file_flex_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6c, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6c,
	0x65, 0x78, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x32, 0x0a, 0x0b, 0x46, 0x6c, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x6c,
	0x65, 0x78, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x30, 0x0a, 0x09, 0x46, 0x6c, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x23, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x66, 0x6c, 0x65, 0x78, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xa0, 0x05, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x0b, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x12, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x0d,
	0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x10, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x27, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x6c, 0x65, 0x78, 0x2e, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x37, 0x0a, 0x10,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x6c, 0x65, 0x78, 0x2e, 0x4f, 0x74,
	0x68, 0x65, 0x72, 0x52, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x73, 0x67,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3c, 0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x4d, 0x73, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x66, 0x6c, 0x65, 0x78, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x10,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x73, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x4d, 0x73, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x32, 0xaf, 0x02, 0x0a, 0x0b, 0x46, 0x6c, 0x65, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x31, 0x0a, 0x09, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x52, 0x50, 0x43, 0x12, 0x11,
	0x2e, 0x66, 0x6c, 0x65, 0x78, 0x2e, 0x46, 0x6c, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x66, 0x6c, 0x65, 0x78, 0x2e, 0x46, 0x6c, 0x65, 0x78, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x66, 0x6c, 0x65, 0x78, 0x2e, 0x46, 0x6c, 0x65, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x66, 0x6c, 0x65, 0x78, 0x2e, 0x46,
	0x6c, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x36, 0x0a, 0x0c,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x66,
	0x6c, 0x65, 0x78, 0x2e, 0x46, 0x6c, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x66, 0x6c, 0x65, 0x78, 0x2e, 0x46, 0x6c, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x15, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e,
	0x66, 0x6c, 0x65, 0x78, 0x2e, 0x46, 0x6c, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x66, 0x6c, 0x65, 0x78, 0x2e, 0x46, 0x6c, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x50,
	0x43, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x6e, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x65, 0x78, 0x42,
	0x09, 0x46, 0x6c, 0x65, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x64, 0x6f, 0x72, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x2f, 0x67, 0x65, 0x6e, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x46, 0x6c,
	0x65, 0x78, 0xca, 0x02, 0x04, 0x46, 0x6c, 0x65, 0x78, 0xe2, 0x02, 0x10, 0x46, 0x6c, 0x65, 0x78,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x46,
	0x6c, 0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flex_proto_rawDescOnce sync.Once
	file_flex_proto_rawDescData = file_flex_proto_rawDesc
)

func file_flex_proto_rawDescGZIP() []byte {
	file_flex_proto_rawDescOnce.Do(func() {
		file_flex_proto_rawDescData = protoimpl.X.CompressGZIP(file_flex_proto_rawDescData)
	})
	return file_flex_proto_rawDescData
}

var file_flex_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_flex_proto_goTypes = []interface{}{
	(*FlexRequest)(nil),   // 0: flex.FlexRequest
	(*FlexReply)(nil),     // 1: flex.FlexReply
	(*ComplexType)(nil),   // 2: flex.ComplexType
	(*Other)(nil),         // 3: flex.Other
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_flex_proto_depIdxs = []int32{
	2,  // 0: flex.FlexRequest.msg:type_name -> flex.ComplexType
	2,  // 1: flex.FlexReply.msg:type_name -> flex.ComplexType
	3,  // 2: flex.ComplexType.msgField:type_name -> flex.Other
	3,  // 3: flex.ComplexType.repeatedMsgField:type_name -> flex.Other
	3,  // 4: flex.ComplexType.optionalMsgField:type_name -> flex.Other
	0,  // 5: flex.FlexService.NormalRPC:input_type -> flex.FlexRequest
	0,  // 6: flex.FlexService.ClientStream:input_type -> flex.FlexRequest
	0,  // 7: flex.FlexService.ServerStream:input_type -> flex.FlexRequest
	0,  // 8: flex.FlexService.BiDirectorionalStream:input_type -> flex.FlexRequest
	4,  // 9: flex.FlexService.EmptyRPC:input_type -> google.protobuf.Empty
	1,  // 10: flex.FlexService.NormalRPC:output_type -> flex.FlexReply
	1,  // 11: flex.FlexService.ClientStream:output_type -> flex.FlexReply
	1,  // 12: flex.FlexService.ServerStream:output_type -> flex.FlexReply
	1,  // 13: flex.FlexService.BiDirectorionalStream:output_type -> flex.FlexReply
	4,  // 14: flex.FlexService.EmptyRPC:output_type -> google.protobuf.Empty
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_flex_proto_init() }
func file_flex_proto_init() {
	if File_flex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlexRequest); i {
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
		file_flex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlexReply); i {
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
		file_flex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexType); i {
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
		file_flex_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Other); i {
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
	file_flex_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flex_proto_goTypes,
		DependencyIndexes: file_flex_proto_depIdxs,
		MessageInfos:      file_flex_proto_msgTypes,
	}.Build()
	File_flex_proto = out.File
	file_flex_proto_rawDesc = nil
	file_flex_proto_goTypes = nil
	file_flex_proto_depIdxs = nil
}
