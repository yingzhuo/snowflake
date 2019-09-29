// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snowflake.proto

package protobuf

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

type IdList struct {
	Ids                  []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdList) Reset()         { *m = IdList{} }
func (m *IdList) String() string { return proto.CompactTextString(m) }
func (*IdList) ProtoMessage()    {}
func (*IdList) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa8bfc3cc8f3970, []int{0}
}

func (m *IdList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdList.Unmarshal(m, b)
}
func (m *IdList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdList.Marshal(b, m, deterministic)
}
func (m *IdList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdList.Merge(m, src)
}
func (m *IdList) XXX_Size() int {
	return xxx_messageInfo_IdList.Size(m)
}
func (m *IdList) XXX_DiscardUnknown() {
	xxx_messageInfo_IdList.DiscardUnknown(m)
}

var xxx_messageInfo_IdList proto.InternalMessageInfo

func (m *IdList) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*IdList)(nil), "snowflake.IdList")
}

func init() { proto.RegisterFile("snowflake.proto", fileDescriptor_aaa8bfc3cc8f3970) }

var fileDescriptor_aaa8bfc3cc8f3970 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xce, 0xcb, 0x2f,
	0x4f, 0xcb, 0x49, 0xcc, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x49, 0x71, 0xb1, 0x79, 0xa6, 0xf8, 0x64, 0x16, 0x97, 0x08, 0x09, 0x70, 0x31, 0x67, 0xa6, 0x14,
	0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x30, 0x07, 0x81, 0x98, 0x4e, 0xc1, 0x5c, 0xca, 0xc9, 0xf9, 0xb9,
	0x7a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0x95, 0x99, 0x79, 0xe9, 0x55, 0x19, 0xa5, 0xf9,
	0x7a, 0x68, 0xa6, 0x39, 0xf1, 0x05, 0xc3, 0x04, 0x02, 0x40, 0xfc, 0x28, 0x19, 0xa8, 0x86, 0xe4,
	0xfc, 0x5c, 0x7d, 0x98, 0x26, 0x7d, 0xb0, 0xd2, 0xa4, 0xd2, 0xb4, 0x24, 0x36, 0x30, 0xcb, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x82, 0xf1, 0xb5, 0x0e, 0x95, 0x00, 0x00, 0x00,
}
