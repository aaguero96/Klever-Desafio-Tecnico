// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/service.proto

package pb_service

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

type NewService struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Site                 string   `protobuf:"bytes,2,opt,name=site,proto3" json:"site,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewService) Reset()         { *m = NewService{} }
func (m *NewService) String() string { return proto.CompactTextString(m) }
func (*NewService) ProtoMessage()    {}
func (*NewService) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{0}
}

func (m *NewService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewService.Unmarshal(m, b)
}
func (m *NewService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewService.Marshal(b, m, deterministic)
}
func (m *NewService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewService.Merge(m, src)
}
func (m *NewService) XXX_Size() int {
	return xxx_messageInfo_NewService.Size(m)
}
func (m *NewService) XXX_DiscardUnknown() {
	xxx_messageInfo_NewService.DiscardUnknown(m)
}

var xxx_messageInfo_NewService proto.InternalMessageInfo

func (m *NewService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewService) GetSite() string {
	if m != nil {
		return m.Site
	}
	return ""
}

type Service struct {
	ServiceId            string   `protobuf:"bytes,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Site                 string   `protobuf:"bytes,3,opt,name=site,proto3" json:"site,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{1}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetSite() string {
	if m != nil {
		return m.Site
	}
	return ""
}

func init() {
	proto.RegisterType((*NewService)(nil), "NewService")
	proto.RegisterType((*Service)(nil), "Service")
}

func init() { proto.RegisterFile("proto/service.proto", fileDescriptor_c33392ef2c1961ba) }

var fileDescriptor_c33392ef2c1961ba = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0xf3, 0x94, 0x4c, 0xb8, 0xb8,
	0xfc, 0x52, 0xcb, 0x83, 0x21, 0x62, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0x36, 0x48, 0xac, 0x38, 0xb3, 0x24, 0x55, 0x82, 0x09, 0x22,
	0x06, 0x62, 0x2b, 0xf9, 0x73, 0xb1, 0xc3, 0xb4, 0xc8, 0x70, 0x71, 0x42, 0x4d, 0xf4, 0x4c, 0x81,
	0xea, 0x43, 0x08, 0xc0, 0x0d, 0x64, 0xc2, 0x62, 0x20, 0x33, 0xc2, 0x40, 0x23, 0x63, 0x2e, 0x3e,
	0xa8, 0x81, 0x30, 0x73, 0x15, 0xb9, 0xd8, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x85, 0xb8, 0xf5,
	0x10, 0x2e, 0x94, 0xe2, 0xd0, 0x83, 0xb2, 0x94, 0x18, 0x9c, 0x24, 0xa2, 0xc4, 0xf4, 0x0b, 0x92,
	0x60, 0x1e, 0xb2, 0x2e, 0x48, 0x8a, 0x87, 0x32, 0x93, 0xd8, 0xc0, 0x9e, 0x33, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xcc, 0x73, 0x4d, 0x3f, 0xf3, 0x00, 0x00, 0x00,
}