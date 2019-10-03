// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package messages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Used to seralise messages for storaging and transfering.
type DomainMessage struct {
	// Identitifer for message.
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// The name of the message.
	MessageName string `protobuf:"bytes,2,opt,name=message_name,json=messageName,proto3" json:"message_name,omitempty"`
	// The message payload.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// Metadata associated to the message.
	Metadata map[string][]byte `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The version of the domain message.
	Version uint64 `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	// When the domain message was first created.
	Created              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DomainMessage) Reset()         { *m = DomainMessage{} }
func (m *DomainMessage) String() string { return proto.CompactTextString(m) }
func (*DomainMessage) ProtoMessage()    {}
func (*DomainMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *DomainMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainMessage.Unmarshal(m, b)
}
func (m *DomainMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainMessage.Marshal(b, m, deterministic)
}
func (m *DomainMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainMessage.Merge(m, src)
}
func (m *DomainMessage) XXX_Size() int {
	return xxx_messageInfo_DomainMessage.Size(m)
}
func (m *DomainMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DomainMessage proto.InternalMessageInfo

func (m *DomainMessage) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *DomainMessage) GetMessageName() string {
	if m != nil {
		return m.MessageName
	}
	return ""
}

func (m *DomainMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DomainMessage) GetMetadata() map[string][]byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *DomainMessage) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *DomainMessage) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

type TestPayload struct {
	AString              string   `protobuf:"bytes,1,opt,name=a_string,json=aString,proto3" json:"a_string,omitempty"`
	AInt                 int64    `protobuf:"varint,2,opt,name=a_int,json=aInt,proto3" json:"a_int,omitempty"`
	ABool                bool     `protobuf:"varint,3,opt,name=a_bool,json=aBool,proto3" json:"a_bool,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestPayload) Reset()         { *m = TestPayload{} }
func (m *TestPayload) String() string { return proto.CompactTextString(m) }
func (*TestPayload) ProtoMessage()    {}
func (*TestPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *TestPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestPayload.Unmarshal(m, b)
}
func (m *TestPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestPayload.Marshal(b, m, deterministic)
}
func (m *TestPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestPayload.Merge(m, src)
}
func (m *TestPayload) XXX_Size() int {
	return xxx_messageInfo_TestPayload.Size(m)
}
func (m *TestPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_TestPayload.DiscardUnknown(m)
}

var xxx_messageInfo_TestPayload proto.InternalMessageInfo

func (m *TestPayload) GetAString() string {
	if m != nil {
		return m.AString
	}
	return ""
}

func (m *TestPayload) GetAInt() int64 {
	if m != nil {
		return m.AInt
	}
	return 0
}

func (m *TestPayload) GetABool() bool {
	if m != nil {
		return m.ABool
	}
	return false
}

func init() {
	proto.RegisterType((*DomainMessage)(nil), "com.github.go_cqrses.cqrses.messages.DomainMessage")
	proto.RegisterMapType((map[string][]byte)(nil), "com.github.go_cqrses.cqrses.messages.DomainMessage.MetadataEntry")
	proto.RegisterType((*TestPayload)(nil), "com.github.go_cqrses.cqrses.messages.TestPayload")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0x49, 0x93, 0xf4, 0xcf, 0x4d, 0x0b, 0x8f, 0x79, 0xef, 0x41, 0x2c, 0x88, 0xb1, 0xb8,
	0xc8, 0x6a, 0x0a, 0xd5, 0x85, 0xd8, 0x95, 0x45, 0x17, 0x5d, 0x54, 0x64, 0xec, 0x4a, 0x90, 0x70,
	0xd3, 0x8c, 0x31, 0x98, 0xc9, 0xd4, 0xcc, 0xb4, 0xd0, 0x8f, 0xe9, 0x37, 0x12, 0x27, 0x13, 0xa1,
	0x3b, 0x57, 0x73, 0xef, 0x9d, 0x73, 0xce, 0xdc, 0xf9, 0xc1, 0x48, 0x70, 0xa5, 0x30, 0xe7, 0x74,
	0x5b, 0x4b, 0x2d, 0xc9, 0xc5, 0x46, 0x0a, 0x9a, 0x17, 0xfa, 0x6d, 0x97, 0xd2, 0x5c, 0x26, 0x9b,
	0x8f, 0x5a, 0x71, 0x45, 0xed, 0x61, 0xa5, 0x6a, 0x7c, 0x96, 0x4b, 0x99, 0x97, 0x7c, 0x6a, 0x3c,
	0xe9, 0xee, 0x75, 0xaa, 0x0b, 0xc1, 0x95, 0x46, 0xb1, 0x6d, 0x62, 0x26, 0x9f, 0x1d, 0x18, 0xdd,
	0x49, 0x81, 0x45, 0xb5, 0x6a, 0x3c, 0xe4, 0x14, 0xc0, 0xda, 0x93, 0x22, 0x0b, 0x9d, 0xc8, 0x89,
	0x07, 0x6c, 0x60, 0x27, 0xcb, 0x8c, 0x9c, 0xc3, 0xb0, 0xbd, 0xae, 0x50, 0xf0, 0xb0, 0x63, 0x04,
	0x81, 0x9d, 0x3d, 0xa0, 0xe0, 0x84, 0x80, 0x97, 0xa1, 0xc6, 0xd0, 0x8d, 0x9c, 0x78, 0xc8, 0x4c,
	0x4d, 0x5e, 0xa0, 0x2f, 0xb8, 0x46, 0x33, 0xf7, 0x22, 0x37, 0x0e, 0x66, 0xb7, 0xf4, 0x37, 0x3f,
	0xa0, 0x47, 0xcb, 0xd1, 0x95, 0xcd, 0xb8, 0xaf, 0x74, 0x7d, 0x60, 0x3f, 0x91, 0x24, 0x84, 0xde,
	0x9e, 0xd7, 0xaa, 0x90, 0x55, 0xe8, 0x47, 0x4e, 0xec, 0xb1, 0xb6, 0x25, 0x57, 0xd0, 0xdb, 0xd4,
	0x1c, 0x35, 0xcf, 0xc2, 0x6e, 0xe4, 0xc4, 0xc1, 0x6c, 0x4c, 0x1b, 0x26, 0xb4, 0x65, 0x42, 0xd7,
	0x2d, 0x13, 0xd6, 0x4a, 0xc7, 0x73, 0x18, 0x1d, 0x3d, 0x45, 0xfe, 0x80, 0xfb, 0xce, 0x0f, 0x16,
	0xc7, 0x77, 0x49, 0xfe, 0x81, 0xbf, 0xc7, 0x72, 0xd7, 0x10, 0x18, 0xb2, 0xa6, 0xb9, 0xe9, 0x5c,
	0x3b, 0x93, 0x35, 0x04, 0x6b, 0xae, 0xf4, 0x23, 0x1e, 0x4a, 0x89, 0x19, 0x39, 0x81, 0x3e, 0x26,
	0x4a, 0xd7, 0x45, 0x95, 0x5b, 0x7f, 0x0f, 0x9f, 0x4c, 0x4b, 0xfe, 0x82, 0x8f, 0x49, 0x51, 0x69,
	0x93, 0xe1, 0x32, 0x0f, 0x97, 0x95, 0x26, 0xff, 0xa1, 0x8b, 0x49, 0x2a, 0x65, 0x69, 0x00, 0xf6,
	0x99, 0x8f, 0x0b, 0x29, 0xcb, 0x45, 0xf0, 0x3c, 0x98, 0xb7, 0x54, 0xd2, 0xae, 0x59, 0xfe, 0xf2,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x61, 0x4b, 0x0c, 0x62, 0x15, 0x02, 0x00, 0x00,
}
