// Code generated by protoc-gen-go. DO NOT EDIT.
// source: errors.proto

package errors

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Exception struct {
	Code                 *ErrorCode `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Exception            string     `protobuf:"bytes,2,opt,name=Exception,proto3" json:"Exception,omitempty"`
	BS                   []byte     `protobuf:"bytes,3,opt,name=BS,proto3" json:"BS,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Exception) Reset()         { *m = Exception{} }
func (m *Exception) String() string { return proto.CompactTextString(m) }
func (*Exception) ProtoMessage()    {}
func (*Exception) Descriptor() ([]byte, []int) {
	return fileDescriptor_errors_75c2645022d60ab3, []int{0}
}
func (m *Exception) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Exception.Unmarshal(m, b)
}
func (m *Exception) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Exception.Marshal(b, m, deterministic)
}
func (dst *Exception) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Exception.Merge(dst, src)
}
func (m *Exception) XXX_Size() int {
	return xxx_messageInfo_Exception.Size(m)
}
func (m *Exception) XXX_DiscardUnknown() {
	xxx_messageInfo_Exception.DiscardUnknown(m)
}

var xxx_messageInfo_Exception proto.InternalMessageInfo

func (m *Exception) GetCode() *ErrorCode {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *Exception) GetException() string {
	if m != nil {
		return m.Exception
	}
	return ""
}

func (m *Exception) GetBS() []byte {
	if m != nil {
		return m.BS
	}
	return nil
}

type ErrorCode struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorCode) Reset()         { *m = ErrorCode{} }
func (m *ErrorCode) String() string { return proto.CompactTextString(m) }
func (*ErrorCode) ProtoMessage()    {}
func (*ErrorCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_errors_75c2645022d60ab3, []int{1}
}
func (m *ErrorCode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorCode.Unmarshal(m, b)
}
func (m *ErrorCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorCode.Marshal(b, m, deterministic)
}
func (dst *ErrorCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorCode.Merge(dst, src)
}
func (m *ErrorCode) XXX_Size() int {
	return xxx_messageInfo_ErrorCode.Size(m)
}
func (m *ErrorCode) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorCode.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorCode proto.InternalMessageInfo

func (m *ErrorCode) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*Exception)(nil), "Exception")
	proto.RegisterType((*ErrorCode)(nil), "ErrorCode")
}

func init() { proto.RegisterFile("errors.proto", fileDescriptor_errors_75c2645022d60ab3) }

var fileDescriptor_errors_75c2645022d60ab3 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2d, 0x2a, 0xca,
	0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x8a, 0xe4, 0xe2, 0x74, 0xad, 0x48, 0x4e,
	0x2d, 0x28, 0xc9, 0xcc, 0xcf, 0x13, 0x92, 0xe3, 0x62, 0x71, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd2, 0x73, 0x05, 0xa9, 0x04, 0x89, 0x04, 0x81, 0xc5, 0x85, 0x64,
	0x90, 0x14, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x21, 0xe9, 0xe6, 0xe3, 0x62, 0x72, 0x0a,
	0x96, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x09, 0x62, 0x72, 0x0a, 0x56, 0x92, 0xe7, 0xe2, 0x84, 0x1b,
	0x20, 0x24, 0x84, 0x64, 0x34, 0x2f, 0xc4, 0xb8, 0x24, 0x36, 0xb0, 0x13, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x51, 0xa4, 0x54, 0xe6, 0x92, 0x00, 0x00, 0x00,
}