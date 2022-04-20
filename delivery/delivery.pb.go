// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delivery/delivery.proto

package delivery

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

type DeliverRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliverRequest) Reset()         { *m = DeliverRequest{} }
func (m *DeliverRequest) String() string { return proto.CompactTextString(m) }
func (*DeliverRequest) ProtoMessage()    {}
func (*DeliverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15d787ed8ff9602d, []int{0}
}

func (m *DeliverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliverRequest.Unmarshal(m, b)
}
func (m *DeliverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliverRequest.Marshal(b, m, deterministic)
}
func (m *DeliverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliverRequest.Merge(m, src)
}
func (m *DeliverRequest) XXX_Size() int {
	return xxx_messageInfo_DeliverRequest.Size(m)
}
func (m *DeliverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeliverRequest proto.InternalMessageInfo

func (m *DeliverRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type DeliverResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliverResponse) Reset()         { *m = DeliverResponse{} }
func (m *DeliverResponse) String() string { return proto.CompactTextString(m) }
func (*DeliverResponse) ProtoMessage()    {}
func (*DeliverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15d787ed8ff9602d, []int{1}
}

func (m *DeliverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliverResponse.Unmarshal(m, b)
}
func (m *DeliverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliverResponse.Marshal(b, m, deterministic)
}
func (m *DeliverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliverResponse.Merge(m, src)
}
func (m *DeliverResponse) XXX_Size() int {
	return xxx_messageInfo_DeliverResponse.Size(m)
}
func (m *DeliverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeliverResponse proto.InternalMessageInfo

func (m *DeliverResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*DeliverRequest)(nil), "delivery.DeliverRequest")
	proto.RegisterType((*DeliverResponse)(nil), "delivery.DeliverResponse")
}

func init() { proto.RegisterFile("delivery/delivery.proto", fileDescriptor_15d787ed8ff9602d) }

var fileDescriptor_15d787ed8ff9602d = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x49, 0xcd, 0xc9,
	0x2c, 0x4b, 0x2d, 0xaa, 0xd4, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0x25, 0x6d, 0x2e, 0x3e, 0x17, 0x08, 0x3b, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48,
	0x92, 0x8b, 0x23, 0xbf, 0x28, 0x25, 0xb5, 0x28, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x88, 0x1d, 0xcc, 0xf7, 0x4c, 0x51, 0xd2, 0xe4, 0xe2, 0x87, 0x2b, 0x2e, 0x2e, 0xc8, 0xcf,
	0x2b, 0x4e, 0x15, 0x12, 0xe3, 0x62, 0x2b, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x06, 0xab, 0xe5, 0x08,
	0x82, 0xf2, 0x8c, 0xbc, 0xb9, 0xd8, 0xa1, 0x4a, 0x85, 0x1c, 0x10, 0x4c, 0x09, 0x3d, 0xb8, 0x43,
	0x50, 0x6d, 0x95, 0x92, 0xc4, 0x22, 0x03, 0xb1, 0x42, 0x89, 0xc1, 0xc9, 0x20, 0x4a, 0x2f, 0x3d,
	0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x38, 0x23, 0x33, 0x2f, 0x23, 0x31,
	0xbd, 0x34, 0x2f, 0x4f, 0xbf, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x17, 0xaa, 0x17, 0xee,
	0xcd, 0x24, 0x36, 0xb0, 0x3f, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xee, 0x19, 0xe1, 0x9f,
	0x02, 0x01, 0x00, 0x00,
}
