// Code generated by protoc-gen-go. DO NOT EDIT.
// source: other.proto

package im

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type S1 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  string   `protobuf:"bytes,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S1) Reset()         { *m = S1{} }
func (m *S1) String() string { return proto.CompactTextString(m) }
func (*S1) ProtoMessage()    {}
func (*S1) Descriptor() ([]byte, []int) {
	return fileDescriptor_b402626a50d68a3b, []int{0}
}

func (m *S1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S1.Unmarshal(m, b)
}
func (m *S1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S1.Marshal(b, m, deterministic)
}
func (m *S1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S1.Merge(m, src)
}
func (m *S1) XXX_Size() int {
	return xxx_messageInfo_S1.Size(m)
}
func (m *S1) XXX_DiscardUnknown() {
	xxx_messageInfo_S1.DiscardUnknown(m)
}

var xxx_messageInfo_S1 proto.InternalMessageInfo

func (m *S1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *S1) GetAge() string {
	if m != nil {
		return m.Age
	}
	return ""
}

func init() {
	proto.RegisterType((*S1)(nil), "im.S1")
}

func init() { proto.RegisterFile("other.proto", fileDescriptor_b402626a50d68a3b) }

var fileDescriptor_b402626a50d68a3b = []byte{
	// 81 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2f, 0xc9, 0x48,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xca, 0xcc, 0x55, 0xd2, 0xe2, 0x62, 0x0a,
	0x36, 0x14, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x13, 0xd3, 0x53, 0x25, 0x98, 0xc0, 0x42, 0x20, 0x66, 0x12,
	0x1b, 0x58, 0x9b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x31, 0xcd, 0x7f, 0x45, 0x00, 0x00,
	0x00,
}
