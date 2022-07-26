// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gRPC/proto/upvotes.proto

package backend

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

type NewVote struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceId            uint32   `protobuf:"varint,2,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	UserId               uint32   `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Vote                 string   `protobuf:"bytes,4,opt,name=vote,proto3" json:"vote,omitempty"`
	Comment              string   `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewVote) Reset()         { *m = NewVote{} }
func (m *NewVote) String() string { return proto.CompactTextString(m) }
func (*NewVote) ProtoMessage()    {}
func (*NewVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f57ddbc3e708c53, []int{0}
}

func (m *NewVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewVote.Unmarshal(m, b)
}
func (m *NewVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewVote.Marshal(b, m, deterministic)
}
func (m *NewVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewVote.Merge(m, src)
}
func (m *NewVote) XXX_Size() int {
	return xxx_messageInfo_NewVote.Size(m)
}
func (m *NewVote) XXX_DiscardUnknown() {
	xxx_messageInfo_NewVote.DiscardUnknown(m)
}

var xxx_messageInfo_NewVote proto.InternalMessageInfo

func (m *NewVote) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NewVote) GetServiceId() uint32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *NewVote) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *NewVote) GetVote() string {
	if m != nil {
		return m.Vote
	}
	return ""
}

func (m *NewVote) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type Vote struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Vote                 string   `protobuf:"bytes,4,opt,name=vote,proto3" json:"vote,omitempty"`
	Comment              string   `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f57ddbc3e708c53, []int{1}
}

func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Vote) GetVote() string {
	if m != nil {
		return m.Vote
	}
	return ""
}

func (m *Vote) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type FilterComment struct {
	Comment              string   `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterComment) Reset()         { *m = FilterComment{} }
func (m *FilterComment) String() string { return proto.CompactTextString(m) }
func (*FilterComment) ProtoMessage()    {}
func (*FilterComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f57ddbc3e708c53, []int{2}
}

func (m *FilterComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterComment.Unmarshal(m, b)
}
func (m *FilterComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterComment.Marshal(b, m, deterministic)
}
func (m *FilterComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterComment.Merge(m, src)
}
func (m *FilterComment) XXX_Size() int {
	return xxx_messageInfo_FilterComment.Size(m)
}
func (m *FilterComment) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterComment.DiscardUnknown(m)
}

var xxx_messageInfo_FilterComment proto.InternalMessageInfo

func (m *FilterComment) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type Votes struct {
	Votes                []*Vote  `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Votes) Reset()         { *m = Votes{} }
func (m *Votes) String() string { return proto.CompactTextString(m) }
func (*Votes) ProtoMessage()    {}
func (*Votes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f57ddbc3e708c53, []int{3}
}

func (m *Votes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Votes.Unmarshal(m, b)
}
func (m *Votes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Votes.Marshal(b, m, deterministic)
}
func (m *Votes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Votes.Merge(m, src)
}
func (m *Votes) XXX_Size() int {
	return xxx_messageInfo_Votes.Size(m)
}
func (m *Votes) XXX_DiscardUnknown() {
	xxx_messageInfo_Votes.DiscardUnknown(m)
}

var xxx_messageInfo_Votes proto.InternalMessageInfo

func (m *Votes) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

type VoteId struct {
	VoteId               uint32   `protobuf:"varint,1,opt,name=voteId,proto3" json:"voteId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteId) Reset()         { *m = VoteId{} }
func (m *VoteId) String() string { return proto.CompactTextString(m) }
func (*VoteId) ProtoMessage()    {}
func (*VoteId) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f57ddbc3e708c53, []int{4}
}

func (m *VoteId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteId.Unmarshal(m, b)
}
func (m *VoteId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteId.Marshal(b, m, deterministic)
}
func (m *VoteId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteId.Merge(m, src)
}
func (m *VoteId) XXX_Size() int {
	return xxx_messageInfo_VoteId.Size(m)
}
func (m *VoteId) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteId.DiscardUnknown(m)
}

var xxx_messageInfo_VoteId proto.InternalMessageInfo

func (m *VoteId) GetVoteId() uint32 {
	if m != nil {
		return m.VoteId
	}
	return 0
}

type UpdateVote struct {
	VoteId               uint32   `protobuf:"varint,1,opt,name=voteId,proto3" json:"voteId,omitempty"`
	Vote                 string   `protobuf:"bytes,4,opt,name=vote,proto3" json:"vote,omitempty"`
	Comment              string   `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateVote) Reset()         { *m = UpdateVote{} }
func (m *UpdateVote) String() string { return proto.CompactTextString(m) }
func (*UpdateVote) ProtoMessage()    {}
func (*UpdateVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f57ddbc3e708c53, []int{5}
}

func (m *UpdateVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateVote.Unmarshal(m, b)
}
func (m *UpdateVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateVote.Marshal(b, m, deterministic)
}
func (m *UpdateVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateVote.Merge(m, src)
}
func (m *UpdateVote) XXX_Size() int {
	return xxx_messageInfo_UpdateVote.Size(m)
}
func (m *UpdateVote) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateVote.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateVote proto.InternalMessageInfo

func (m *UpdateVote) GetVoteId() uint32 {
	if m != nil {
		return m.VoteId
	}
	return 0
}

func (m *UpdateVote) GetVote() string {
	if m != nil {
		return m.Vote
	}
	return ""
}

func (m *UpdateVote) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type ServiceId struct {
	ServiceId            uint32   `protobuf:"varint,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceId) Reset()         { *m = ServiceId{} }
func (m *ServiceId) String() string { return proto.CompactTextString(m) }
func (*ServiceId) ProtoMessage()    {}
func (*ServiceId) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f57ddbc3e708c53, []int{6}
}

func (m *ServiceId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceId.Unmarshal(m, b)
}
func (m *ServiceId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceId.Marshal(b, m, deterministic)
}
func (m *ServiceId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceId.Merge(m, src)
}
func (m *ServiceId) XXX_Size() int {
	return xxx_messageInfo_ServiceId.Size(m)
}
func (m *ServiceId) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceId.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceId proto.InternalMessageInfo

func (m *ServiceId) GetServiceId() uint32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

type Report struct {
	ServiceId            uint32   `protobuf:"varint,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	TotalUpvotes         uint64   `protobuf:"varint,2,opt,name=totalUpvotes,proto3" json:"totalUpvotes,omitempty"`
	TotalDownvotes       uint64   `protobuf:"varint,3,opt,name=totalDownvotes,proto3" json:"totalDownvotes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f57ddbc3e708c53, []int{7}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetServiceId() uint32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *Report) GetTotalUpvotes() uint64 {
	if m != nil {
		return m.TotalUpvotes
	}
	return 0
}

func (m *Report) GetTotalDownvotes() uint64 {
	if m != nil {
		return m.TotalDownvotes
	}
	return 0
}

type Comment struct {
	ServiceId            uint32   `protobuf:"varint,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f57ddbc3e708c53, []int{8}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetServiceId() uint32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *Comment) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type Comments struct {
	Comments             []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Comments) Reset()         { *m = Comments{} }
func (m *Comments) String() string { return proto.CompactTextString(m) }
func (*Comments) ProtoMessage()    {}
func (*Comments) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f57ddbc3e708c53, []int{9}
}

func (m *Comments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comments.Unmarshal(m, b)
}
func (m *Comments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comments.Marshal(b, m, deterministic)
}
func (m *Comments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comments.Merge(m, src)
}
func (m *Comments) XXX_Size() int {
	return xxx_messageInfo_Comments.Size(m)
}
func (m *Comments) XXX_DiscardUnknown() {
	xxx_messageInfo_Comments.DiscardUnknown(m)
}

var xxx_messageInfo_Comments proto.InternalMessageInfo

func (m *Comments) GetComments() []*Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

func init() {
	proto.RegisterType((*NewVote)(nil), "proto.NewVote")
	proto.RegisterType((*Vote)(nil), "proto.Vote")
	proto.RegisterType((*FilterComment)(nil), "proto.FilterComment")
	proto.RegisterType((*Votes)(nil), "proto.Votes")
	proto.RegisterType((*VoteId)(nil), "proto.voteId")
	proto.RegisterType((*UpdateVote)(nil), "proto.UpdateVote")
	proto.RegisterType((*ServiceId)(nil), "proto.serviceId")
	proto.RegisterType((*Report)(nil), "proto.Report")
	proto.RegisterType((*Comment)(nil), "proto.Comment")
	proto.RegisterType((*Comments)(nil), "proto.Comments")
}

func init() { proto.RegisterFile("gRPC/proto/upvotes.proto", fileDescriptor_4f57ddbc3e708c53) }

var fileDescriptor_4f57ddbc3e708c53 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xb5, 0x64, 0x59, 0x76, 0x26, 0xb6, 0xdb, 0x0c, 0xa1, 0x2c, 0xa6, 0x07, 0x75, 0x0f, 0xad,
	0x1a, 0x8a, 0x0d, 0x29, 0xe4, 0x92, 0x53, 0x2d, 0x53, 0xf0, 0xa5, 0x94, 0x0d, 0xe9, 0xa1, 0x37,
	0xc5, 0x3b, 0x14, 0x51, 0xdb, 0x2b, 0xa4, 0x4d, 0x42, 0x0e, 0xfd, 0xd4, 0xfe, 0x4b, 0xd1, 0xee,
	0x5a, 0x91, 0x9c, 0xc6, 0x90, 0x93, 0x76, 0xde, 0xbc, 0x99, 0xd1, 0xbc, 0x37, 0xc0, 0x7e, 0x89,
	0xef, 0xc9, 0x2c, 0x2f, 0x94, 0x56, 0xb3, 0xdb, 0xfc, 0x4e, 0x69, 0x2a, 0xa7, 0x26, 0xc2, 0x9e,
	0xf9, 0xf0, 0x3f, 0xd0, 0xff, 0x46, 0xf7, 0x3f, 0x94, 0x26, 0x1c, 0x83, 0x9f, 0x49, 0xe6, 0x45,
	0x5e, 0x3c, 0x12, 0x7e, 0x26, 0xf1, 0x2d, 0x1c, 0x95, 0x54, 0xdc, 0x65, 0x2b, 0x5a, 0x4a, 0xe6,
	0x1b, 0xf8, 0x11, 0xc0, 0x37, 0x10, 0xde, 0x96, 0x54, 0x2c, 0x25, 0xeb, 0x9a, 0x94, 0x8b, 0x10,
	0x21, 0xa8, 0xc6, 0xb0, 0x20, 0xf2, 0xe2, 0x23, 0x61, 0xde, 0xc8, 0xa0, 0xbf, 0x52, 0x9b, 0x0d,
	0x6d, 0x35, 0xeb, 0x19, 0x78, 0x17, 0xf2, 0x05, 0x04, 0xff, 0x9d, 0xfd, 0xb2, 0x2e, 0x1f, 0x61,
	0xf4, 0x35, 0x5b, 0x6b, 0x2a, 0x12, 0x0b, 0x34, 0xa9, 0x5e, 0x9b, 0x7a, 0x06, 0xbd, 0x6a, 0x60,
	0x89, 0xef, 0xa0, 0x67, 0xe4, 0x60, 0x5e, 0xd4, 0x8d, 0x8f, 0xcf, 0x8f, 0xad, 0x2c, 0xd3, 0x2a,
	0x29, 0x6c, 0x86, 0x47, 0x10, 0x56, 0x0f, 0xbb, 0xac, 0x7d, 0xb9, 0x5f, 0x74, 0x11, 0x17, 0x00,
	0xd7, 0xb9, 0x4c, 0x35, 0x99, 0x25, 0x9e, 0x61, 0xbd, 0x78, 0x99, 0x86, 0xca, 0x2d, 0x0f, 0xbc,
	0x3d, 0x0f, 0x78, 0x01, 0xa1, 0xa0, 0x5c, 0x15, 0xfa, 0x30, 0x0f, 0x39, 0x0c, 0xb5, 0xd2, 0xe9,
	0xfa, 0xda, 0x5e, 0x80, 0x31, 0x33, 0x10, 0x2d, 0x0c, 0xdf, 0xc3, 0xd8, 0xc4, 0x0b, 0x75, 0xbf,
	0xb5, 0xac, 0xae, 0x61, 0xed, 0xa1, 0xfc, 0x0b, 0xf4, 0x77, 0x2a, 0x1f, 0x1e, 0xda, 0xd8, 0xd0,
	0x6f, 0x6f, 0x78, 0x01, 0x03, 0xd7, 0xa2, 0xc4, 0x33, 0x18, 0x38, 0x78, 0xe7, 0xc4, 0xd8, 0x39,
	0xe1, 0x28, 0xa2, 0xce, 0x9f, 0xff, 0xf5, 0x61, 0x64, 0x7f, 0xf7, 0xca, 0x4e, 0xc1, 0x0f, 0x10,
	0x26, 0x05, 0xa5, 0xd5, 0x01, 0xb9, 0x2a, 0x77, 0xcc, 0x93, 0xa6, 0x9f, 0xbc, 0x83, 0x9f, 0x20,
	0x10, 0x94, 0x4a, 0x3c, 0x75, 0x70, 0xeb, 0x5c, 0x26, 0xc3, 0x06, 0xb9, 0xe4, 0x1d, 0x8c, 0x61,
	0x50, 0xb1, 0xe7, 0x0f, 0x4b, 0x89, 0x23, 0x97, 0xb3, 0x5e, 0x3e, 0xed, 0x1b, 0xda, 0x03, 0xc0,
	0x13, 0x97, 0x78, 0xbc, 0x87, 0x49, 0xbb, 0xd4, 0xf4, 0x0d, 0x17, 0xb4, 0x26, 0x4d, 0xfb, 0x5d,
	0x9f, 0x30, 0x2f, 0xe0, 0xc4, 0x3a, 0x3b, 0x7f, 0xb8, 0xaa, 0x15, 0x7d, 0xed, 0x58, 0xb5, 0xc6,
	0x75, 0x9d, 0xe5, 0xf2, 0x0e, 0x5e, 0xc2, 0x69, 0xa2, 0x36, 0x64, 0xf4, 0x3a, 0x5c, 0xfa, 0xaa,
	0x2d, 0x73, 0xc9, 0x3b, 0xf3, 0xe1, 0x4f, 0x98, 0xce, 0x2e, 0x6f, 0xd2, 0xd5, 0x6f, 0xda, 0xca,
	0x9b, 0xd0, 0xe4, 0x3f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x76, 0xcf, 0x57, 0x24, 0x43, 0x04,
	0x00, 0x00,
}
