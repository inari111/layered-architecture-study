// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Profile              *Profile             `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_1f49687ee0ebebeb, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Profile struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  uint32   `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_1f49687ee0ebebeb, []int{1}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (dst *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(dst, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Profile) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UserServiceRegisterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceRegisterRequest) Reset()         { *m = UserServiceRegisterRequest{} }
func (m *UserServiceRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*UserServiceRegisterRequest) ProtoMessage()    {}
func (*UserServiceRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_1f49687ee0ebebeb, []int{2}
}
func (m *UserServiceRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceRegisterRequest.Unmarshal(m, b)
}
func (m *UserServiceRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceRegisterRequest.Marshal(b, m, deterministic)
}
func (dst *UserServiceRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceRegisterRequest.Merge(dst, src)
}
func (m *UserServiceRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_UserServiceRegisterRequest.Size(m)
}
func (m *UserServiceRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceRegisterRequest proto.InternalMessageInfo

type UserServiceRegisterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceRegisterResponse) Reset()         { *m = UserServiceRegisterResponse{} }
func (m *UserServiceRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*UserServiceRegisterResponse) ProtoMessage()    {}
func (*UserServiceRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_1f49687ee0ebebeb, []int{3}
}
func (m *UserServiceRegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceRegisterResponse.Unmarshal(m, b)
}
func (m *UserServiceRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceRegisterResponse.Marshal(b, m, deterministic)
}
func (dst *UserServiceRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceRegisterResponse.Merge(dst, src)
}
func (m *UserServiceRegisterResponse) XXX_Size() int {
	return xxx_messageInfo_UserServiceRegisterResponse.Size(m)
}
func (m *UserServiceRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceRegisterResponse proto.InternalMessageInfo

type UserServiceUpdateProfileRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  uint32   `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceUpdateProfileRequest) Reset()         { *m = UserServiceUpdateProfileRequest{} }
func (m *UserServiceUpdateProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UserServiceUpdateProfileRequest) ProtoMessage()    {}
func (*UserServiceUpdateProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_1f49687ee0ebebeb, []int{4}
}
func (m *UserServiceUpdateProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceUpdateProfileRequest.Unmarshal(m, b)
}
func (m *UserServiceUpdateProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceUpdateProfileRequest.Marshal(b, m, deterministic)
}
func (dst *UserServiceUpdateProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceUpdateProfileRequest.Merge(dst, src)
}
func (m *UserServiceUpdateProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UserServiceUpdateProfileRequest.Size(m)
}
func (m *UserServiceUpdateProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceUpdateProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceUpdateProfileRequest proto.InternalMessageInfo

func (m *UserServiceUpdateProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserServiceUpdateProfileRequest) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserServiceUpdateProfileRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UserServiceUpdateProfileResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceUpdateProfileResponse) Reset()         { *m = UserServiceUpdateProfileResponse{} }
func (m *UserServiceUpdateProfileResponse) String() string { return proto.CompactTextString(m) }
func (*UserServiceUpdateProfileResponse) ProtoMessage()    {}
func (*UserServiceUpdateProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_1f49687ee0ebebeb, []int{5}
}
func (m *UserServiceUpdateProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceUpdateProfileResponse.Unmarshal(m, b)
}
func (m *UserServiceUpdateProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceUpdateProfileResponse.Marshal(b, m, deterministic)
}
func (dst *UserServiceUpdateProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceUpdateProfileResponse.Merge(dst, src)
}
func (m *UserServiceUpdateProfileResponse) XXX_Size() int {
	return xxx_messageInfo_UserServiceUpdateProfileResponse.Size(m)
}
func (m *UserServiceUpdateProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceUpdateProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceUpdateProfileResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "article.User")
	proto.RegisterType((*Profile)(nil), "article.Profile")
	proto.RegisterType((*UserServiceRegisterRequest)(nil), "article.UserServiceRegisterRequest")
	proto.RegisterType((*UserServiceRegisterResponse)(nil), "article.UserServiceRegisterResponse")
	proto.RegisterType((*UserServiceUpdateProfileRequest)(nil), "article.UserServiceUpdateProfileRequest")
	proto.RegisterType((*UserServiceUpdateProfileResponse)(nil), "article.UserServiceUpdateProfileResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_1f49687ee0ebebeb) }

var fileDescriptor_user_1f49687ee0ebebeb = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0x34, 0xa2, 0xf4, 0xaa, 0xa2, 0xea, 0xa6, 0x2a, 0x80, 0x1a, 0x05, 0x86, 0xc2,
	0x90, 0x4a, 0x65, 0x62, 0x2c, 0x4f, 0x00, 0x86, 0x2e, 0x2c, 0x28, 0x4d, 0xae, 0x91, 0xa5, 0xb6,
	0x36, 0xb6, 0xc3, 0xb3, 0xf1, 0x0c, 0x3c, 0x15, 0xaa, 0x1d, 0xa3, 0x22, 0xa5, 0xc0, 0xc0, 0x66,
	0x9d, 0xbf, 0xbb, 0xff, 0xcb, 0xc5, 0x00, 0xb5, 0x26, 0x95, 0x49, 0x25, 0x8c, 0xc0, 0x6e, 0xae,
	0x0c, 0x2f, 0xd6, 0x14, 0x8f, 0x2b, 0x21, 0xaa, 0x35, 0x4d, 0x6d, 0x79, 0x59, 0xaf, 0xa6, 0x86,
	0x6f, 0x48, 0x9b, 0x7c, 0x23, 0x1d, 0x99, 0xbe, 0x07, 0x10, 0x2d, 0x34, 0x29, 0x3c, 0x81, 0x90,
	0x97, 0xa3, 0x20, 0x09, 0x26, 0x3d, 0x16, 0xf2, 0x12, 0xaf, 0xa1, 0x2b, 0x95, 0x58, 0xf1, 0x35,
	0x8d, 0xc2, 0x24, 0x98, 0xf4, 0x67, 0xc3, 0xac, 0x19, 0x9a, 0xdd, 0xbb, 0x3a, 0xf3, 0x00, 0xde,
	0x02, 0x14, 0x8a, 0x72, 0x43, 0xe5, 0x4b, 0x6e, 0x46, 0x1d, 0x8b, 0xc7, 0x99, 0x8b, 0xce, 0x7c,
	0x74, 0xf6, 0xe4, 0xa3, 0x59, 0xaf, 0xa1, 0xe7, 0x66, 0xd7, 0x5a, 0xcb, 0xd2, 0xb7, 0x46, 0xbf,
	0xb7, 0x36, 0xf4, 0xdc, 0xa4, 0x0f, 0xd0, 0x6d, 0x4c, 0x10, 0x21, 0xda, 0xe6, 0x1b, 0x6a, 0xf4,
	0xed, 0x19, 0x87, 0xd0, 0xc9, 0x2b, 0x27, 0x3f, 0x60, 0xbb, 0x23, 0x26, 0xd0, 0x2f, 0x49, 0x17,
	0x8a, 0x4b, 0xc3, 0xc5, 0xd6, 0x7a, 0xf6, 0xd8, 0x7e, 0x29, 0x3d, 0x83, 0x78, 0xb7, 0x8c, 0x47,
	0x52, 0x6f, 0xbc, 0x20, 0x46, 0x15, 0xd7, 0x86, 0x14, 0xa3, 0xd7, 0x9a, 0xb4, 0x49, 0xcf, 0xe1,
	0xb4, 0xf5, 0x56, 0x4b, 0xb1, 0xd5, 0x94, 0x72, 0x18, 0xef, 0x5d, 0x2f, 0xac, 0xa7, 0x5f, 0x95,
	0x9b, 0xf0, 0x6f, 0x9e, 0x29, 0x24, 0x87, 0xa3, 0x9c, 0xce, 0xec, 0x23, 0x80, 0xfe, 0x1e, 0x84,
	0x0b, 0x38, 0xf6, 0xca, 0x78, 0xf1, 0xf5, 0x2f, 0x0f, 0x7f, 0x6e, 0x7c, 0xf9, 0x33, 0xe4, 0x62,
	0xb0, 0x84, 0xc1, 0xb7, 0x7c, 0x9c, 0xb4, 0xb5, 0xb5, 0x6d, 0x23, 0xbe, 0xfa, 0x03, 0xe9, 0x52,
	0xee, 0xa2, 0xe7, 0x50, 0x2e, 0x97, 0x47, 0xf6, 0x41, 0xdc, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x82, 0x98, 0xf3, 0x18, 0xeb, 0x02, 0x00, 0x00,
}
