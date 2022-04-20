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
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Price                float64  `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
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

func (m *DeliverRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeliverRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DeliverRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *DeliverRequest) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
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
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4e, 0xc4, 0x20,
	0x10, 0x86, 0x45, 0xd7, 0xdd, 0x75, 0x0e, 0x9a, 0x10, 0xa3, 0xe8, 0x69, 0xd3, 0xd3, 0x7a, 0x90,
	0x1a, 0x7d, 0x01, 0x63, 0xbc, 0x79, 0xe3, 0xe8, 0x8d, 0x6d, 0x27, 0x05, 0x63, 0xa1, 0x32, 0xd0,
	0xc4, 0xb7, 0x37, 0x81, 0xb6, 0xc6, 0xc4, 0xdb, 0xf7, 0xfd, 0x40, 0x7e, 0x66, 0xe0, 0xba, 0xc5,
	0x4f, 0x3b, 0x62, 0xf8, 0xae, 0x67, 0x90, 0x43, 0xf0, 0xd1, 0xf3, 0xed, 0xec, 0xd5, 0x07, 0x9c,
	0xbf, 0x16, 0x56, 0xf8, 0x95, 0x90, 0x22, 0xe7, 0xb0, 0x72, 0xba, 0x47, 0xc1, 0x76, 0x6c, 0x7f,
	0xa6, 0x32, 0x73, 0x01, 0x1b, 0xdd, 0xb6, 0x01, 0x89, 0xc4, 0x71, 0x8e, 0x67, 0xe5, 0x97, 0x70,
	0x3a, 0x18, 0xef, 0x50, 0x9c, 0xe4, 0xbc, 0x48, 0x4e, 0x83, 0x6d, 0x50, 0xac, 0x76, 0x6c, 0xcf,
	0x54, 0x91, 0xea, 0x0e, 0x2e, 0x96, 0x2e, 0x1a, 0xbc, 0x23, 0xe4, 0x57, 0xb0, 0xa6, 0xa8, 0x63,
	0xa2, 0x5c, 0xb7, 0x55, 0x93, 0x3d, 0xbe, 0xc1, 0x66, 0xba, 0xca, 0x9f, 0x7f, 0x51, 0xc8, 0x65,
	0x8e, 0xbf, 0x9f, 0xbe, 0xbd, 0xf9, 0xe7, 0xa4, 0x54, 0x54, 0x47, 0x2f, 0x0f, 0xef, 0xb2, 0xb3,
	0xd1, 0xa4, 0x83, 0x6c, 0x7c, 0x5f, 0x93, 0xb1, 0xce, 0xe8, 0x2e, 0x39, 0x57, 0x13, 0x86, 0xd1,
	0x36, 0x78, 0x3f, 0xbd, 0x5d, 0xb6, 0x74, 0x58, 0xe7, 0x35, 0x3d, 0xfd, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0xec, 0x49, 0xb9, 0x41, 0x01, 0x00, 0x00,
}
