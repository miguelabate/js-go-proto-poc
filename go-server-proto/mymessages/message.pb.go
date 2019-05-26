// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package mymessages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MessageMike struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Lang                 string   `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Length               float32  `protobuf:"fixed32,3,opt,name=length,proto3" json:"length,omitempty"`
	Years                int32    `protobuf:"varint,4,opt,name=years,proto3" json:"years,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageMike) Reset()         { *m = MessageMike{} }
func (m *MessageMike) String() string { return proto.CompactTextString(m) }
func (*MessageMike) ProtoMessage()    {}
func (*MessageMike) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *MessageMike) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageMike.Unmarshal(m, b)
}
func (m *MessageMike) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageMike.Marshal(b, m, deterministic)
}
func (m *MessageMike) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageMike.Merge(m, src)
}
func (m *MessageMike) XXX_Size() int {
	return xxx_messageInfo_MessageMike.Size(m)
}
func (m *MessageMike) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageMike.DiscardUnknown(m)
}

var xxx_messageInfo_MessageMike proto.InternalMessageInfo

func (m *MessageMike) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *MessageMike) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *MessageMike) GetLength() float32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *MessageMike) GetYears() int32 {
	if m != nil {
		return m.Years
	}
	return 0
}

func init() {
	proto.RegisterType((*MessageMike)(nil), "MessageMike")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x4a, 0xe6, 0xe2, 0xf6, 0x85, 0x08,
	0xf8, 0x66, 0x66, 0xa7, 0x0a, 0x09, 0x71, 0xb1, 0x94, 0xa4, 0x56, 0x94, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x81, 0xd9, 0x20, 0xb1, 0x9c, 0xc4, 0xbc, 0x74, 0x09, 0x26, 0x88, 0x18, 0x88,
	0x2d, 0x24, 0xc6, 0xc5, 0x96, 0x93, 0x9a, 0x97, 0x5e, 0x92, 0x21, 0xc1, 0xac, 0xc0, 0xa8, 0xc1,
	0x14, 0x04, 0xe5, 0x09, 0x89, 0x70, 0xb1, 0x56, 0xa6, 0x26, 0x16, 0x15, 0x4b, 0xb0, 0x28, 0x30,
	0x6a, 0xb0, 0x06, 0x41, 0x38, 0x4e, 0x3c, 0x51, 0x5c, 0xb9, 0x95, 0x50, 0x7b, 0x8b, 0x93, 0xd8,
	0xc0, 0x36, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x93, 0xf3, 0x0a, 0xcf, 0x8a, 0x00, 0x00,
	0x00,
}