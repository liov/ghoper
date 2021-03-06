// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package protobuf

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

// Timestamp from public import google/protobuf/timestamp.proto
type Timestamp = timestamp.Timestamp

// The request message containing the user's name.
type User struct {
	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// google.protobuf.Timestamp activatedAt = 3;
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Sex       string `protobuf:"bytes,6,opt,name=sex,proto3" json:"sex,omitempty"`
	Score     uint64 `protobuf:"varint,7,opt,name=Score,proto3" json:"Score,omitempty"`
	Signature string `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
	AvatarURL string `protobuf:"bytes,9,opt,name=avatarURL,proto3" json:"avatarURL,omitempty"`
	Role      uint32 `protobuf:"varint,10,opt,name=Role,proto3" json:"Role,omitempty"`
	// google.protobuf.Timestamp createdAt = 12;
	Status               int32    `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `gorm:"-" json:"-"`
	XXX_unrecognized     []byte   `gorm:"-" json:"-"`
	XXX_sizecache        int32    `gorm:"-" json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8d2a4c100c22267f, []int{0}
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

func (m *User) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func (m *User) GetScore() uint64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *User) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *User) GetAvatarURL() string {
	if m != nil {
		return m.AvatarURL
	}
	return ""
}

func (m *User) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *User) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type GetReq struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}
func (*GetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8d2a4c100c22267f, []int{1}
}
func (m *GetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReq.Unmarshal(m, b)
}
func (m *GetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReq.Marshal(b, m, deterministic)
}
func (dst *GetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReq.Merge(dst, src)
}
func (m *GetReq) XXX_Size() int {
	return xxx_messageInfo_GetReq.Size(m)
}
func (m *GetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetReq proto.InternalMessageInfo

func (m *GetReq) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type SignupReq struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Sex                  string   `protobuf:"bytes,7,opt,name=sex,proto3" json:"sex,omitempty"`
	Score                uint64   `protobuf:"varint,8,opt,name=Score,proto3" json:"Score,omitempty"`
	Signature            string   `protobuf:"bytes,9,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignupReq) Reset()         { *m = SignupReq{} }
func (m *SignupReq) String() string { return proto.CompactTextString(m) }
func (*SignupReq) ProtoMessage()    {}
func (*SignupReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8d2a4c100c22267f, []int{2}
}
func (m *SignupReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupReq.Unmarshal(m, b)
}
func (m *SignupReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupReq.Marshal(b, m, deterministic)
}
func (dst *SignupReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupReq.Merge(dst, src)
}
func (m *SignupReq) XXX_Size() int {
	return xxx_messageInfo_SignupReq.Size(m)
}
func (m *SignupReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignupReq proto.InternalMessageInfo

func (m *SignupReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignupReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignupReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignupReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *SignupReq) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func (m *SignupReq) GetScore() uint64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *SignupReq) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type LoginReq struct {
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	LuosimaoRes          string   `protobuf:"bytes,3,opt,name=luosimaoRes,proto3" json:"luosimaoRes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8d2a4c100c22267f, []int{3}
}
func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (dst *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(dst, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginReq) GetLuosimaoRes() string {
	if m != nil {
		return m.LuosimaoRes
	}
	return ""
}

type LoginRep struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRep) Reset()         { *m = LoginRep{} }
func (m *LoginRep) String() string { return proto.CompactTextString(m) }
func (*LoginRep) ProtoMessage()    {}
func (*LoginRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8d2a4c100c22267f, []int{4}
}
func (m *LoginRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRep.Unmarshal(m, b)
}
func (m *LoginRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRep.Marshal(b, m, deterministic)
}
func (dst *LoginRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRep.Merge(dst, src)
}
func (m *LoginRep) XXX_Size() int {
	return xxx_messageInfo_LoginRep.Size(m)
}
func (m *LoginRep) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRep.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRep proto.InternalMessageInfo

func (m *LoginRep) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginRep) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *LoginRep) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type LogoutReq struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutReq) Reset()         { *m = LogoutReq{} }
func (m *LogoutReq) String() string { return proto.CompactTextString(m) }
func (*LogoutReq) ProtoMessage()    {}
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8d2a4c100c22267f, []int{5}
}
func (m *LogoutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutReq.Unmarshal(m, b)
}
func (m *LogoutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutReq.Marshal(b, m, deterministic)
}
func (dst *LogoutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutReq.Merge(dst, src)
}
func (m *LogoutReq) XXX_Size() int {
	return xxx_messageInfo_LogoutReq.Size(m)
}
func (m *LogoutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutReq.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutReq proto.InternalMessageInfo

func (m *LogoutReq) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type LogoutRep struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRep) Reset()         { *m = LogoutRep{} }
func (m *LogoutRep) String() string { return proto.CompactTextString(m) }
func (*LogoutRep) ProtoMessage()    {}
func (*LogoutRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8d2a4c100c22267f, []int{6}
}
func (m *LogoutRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRep.Unmarshal(m, b)
}
func (m *LogoutRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRep.Marshal(b, m, deterministic)
}
func (dst *LogoutRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRep.Merge(dst, src)
}
func (m *LogoutRep) XXX_Size() int {
	return xxx_messageInfo_LogoutRep.Size(m)
}
func (m *LogoutRep) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRep.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRep proto.InternalMessageInfo

func (m *LogoutRep) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "protobuf.User")
	proto.RegisterType((*GetReq)(nil), "protobuf.GetReq")
	proto.RegisterType((*SignupReq)(nil), "protobuf.SignupReq")
	proto.RegisterType((*LoginReq)(nil), "protobuf.LoginReq")
	proto.RegisterType((*LoginRep)(nil), "protobuf.LoginRep")
	proto.RegisterType((*LogoutReq)(nil), "protobuf.LogoutReq")
	proto.RegisterType((*LogoutRep)(nil), "protobuf.LogoutRep")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_8d2a4c100c22267f) }

var fileDescriptor_user_8d2a4c100c22267f = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0xde, 0xa4, 0x49, 0x36, 0x99, 0x15, 0x55, 0x65, 0x2a, 0x64, 0x05, 0x10, 0x91, 0x4f, 0x7b,
	0x61, 0x2b, 0x5a, 0x1e, 0xa1, 0x52, 0x55, 0x69, 0x0f, 0xc8, 0xab, 0x72, 0x44, 0x72, 0xcb, 0x10,
	0x22, 0x36, 0x71, 0x88, 0xed, 0xc2, 0x99, 0xf7, 0xe1, 0x85, 0x78, 0x1a, 0x64, 0x7b, 0x93, 0x2c,
	0x4b, 0x56, 0x3d, 0x65, 0xbe, 0x6f, 0x7e, 0x3c, 0xf3, 0xcd, 0x04, 0xc0, 0x28, 0xec, 0x56, 0x6d,
	0x27, 0xb5, 0x24, 0xa9, 0xfb, 0xdc, 0x9b, 0x2f, 0xf9, 0x9b, 0x52, 0xca, 0x72, 0x8b, 0x17, 0x3d,
	0x71, 0xa1, 0xab, 0x1a, 0x95, 0x16, 0x75, 0xeb, 0x43, 0xd9, 0xaf, 0x10, 0xa2, 0x3b, 0x85, 0x1d,
	0x39, 0x85, 0xf0, 0xf6, 0x9a, 0x06, 0x45, 0xb0, 0x8c, 0x78, 0x78, 0x7b, 0x4d, 0x08, 0x44, 0x8d,
	0xa8, 0x91, 0x86, 0x45, 0xb0, 0xcc, 0xb8, 0xb3, 0x49, 0x0e, 0x69, 0x2b, 0x94, 0xfa, 0x21, 0xbb,
	0xcf, 0xf4, 0xc4, 0xf1, 0x03, 0x26, 0xe7, 0x10, 0x63, 0x2d, 0xaa, 0x2d, 0x8d, 0x9c, 0xc3, 0x03,
	0xcb, 0xb6, 0x5f, 0x65, 0x83, 0x34, 0xf6, 0xac, 0x03, 0xe4, 0x0c, 0x4e, 0x14, 0xfe, 0xa4, 0x89,
	0xe3, 0xac, 0x69, 0xe3, 0x36, 0x0f, 0xb2, 0x43, 0x3a, 0x77, 0x0d, 0x78, 0x40, 0x5e, 0x41, 0xa6,
	0xaa, 0xb2, 0x11, 0xda, 0x74, 0x48, 0x53, 0x17, 0x3d, 0x12, 0xd6, 0x2b, 0x1e, 0x85, 0x16, 0xdd,
	0x1d, 0x5f, 0xd3, 0xcc, 0x7b, 0x07, 0xc2, 0xf6, 0xcf, 0xe5, 0x16, 0x29, 0x14, 0xc1, 0xf2, 0x19,
	0x77, 0x36, 0x79, 0x01, 0x89, 0xd2, 0x42, 0x1b, 0x45, 0x17, 0x45, 0xb0, 0x8c, 0xf9, 0x0e, 0x31,
	0x0a, 0xc9, 0x0d, 0x6a, 0x8e, 0xdf, 0x0f, 0x55, 0x60, 0xbf, 0x03, 0xc8, 0x36, 0x55, 0xd9, 0x98,
	0xd6, 0x7a, 0x9f, 0xd2, 0x24, 0x3a, 0xa6, 0x49, 0x3c, 0xa9, 0x49, 0x32, 0xa1, 0xc9, 0x7c, 0x42,
	0x93, 0xf4, 0xa8, 0x26, 0xd9, 0x81, 0x26, 0xec, 0x13, 0xa4, 0x6b, 0x59, 0x56, 0x8d, 0xed, 0xf6,
	0x1c, 0xe2, 0xaa, 0x69, 0x8d, 0x76, 0xe3, 0x64, 0xdc, 0x83, 0x7f, 0xfa, 0x0d, 0x0f, 0xfa, 0x2d,
	0x60, 0xb1, 0x35, 0x52, 0x55, 0xb5, 0x90, 0x1c, 0xd5, 0x6e, 0xc5, 0xfb, 0x14, 0xfb, 0x38, 0xd4,
	0x6f, 0x6d, 0x7d, 0x2d, 0xbf, 0x61, 0xd3, 0xd7, 0x77, 0x80, 0x30, 0x88, 0xec, 0x25, 0xba, 0xda,
	0x8b, 0xcb, 0xd3, 0x55, 0x7f, 0x79, 0x2b, 0x7b, 0x65, 0xdc, 0xf9, 0xec, 0xac, 0xb5, 0x2a, 0x77,
	0xf5, 0xad, 0xc9, 0x5e, 0x42, 0xb6, 0x96, 0xa5, 0x34, 0x93, 0x4b, 0x78, 0x3d, 0x3a, 0xdb, 0x3e,
	0x37, 0x18, 0x72, 0x2f, 0xff, 0x04, 0xb0, 0xb0, 0xc5, 0x37, 0xd8, 0x3d, 0x56, 0x0f, 0x48, 0xae,
	0x20, 0xf1, 0x2b, 0x23, 0xcf, 0xc7, 0xd7, 0x87, 0x25, 0xe6, 0x64, 0x24, 0xfb, 0x51, 0xd8, 0x8c,
	0xbc, 0x83, 0xd8, 0x21, 0xf2, 0xbf, 0xfb, 0x58, 0xca, 0x7b, 0x48, 0x7c, 0x5b, 0xfb, 0xef, 0x0c,
	0x53, 0xe4, 0x13, 0xa4, 0xcd, 0x7a, 0x0b, 0xf3, 0x1b, 0xd4, 0xee, 0x97, 0x3b, 0x1b, 0x23, 0xfc,
	0xf9, 0xe5, 0x07, 0x72, 0xb1, 0xd9, 0x87, 0xd9, 0x7d, 0xe2, 0xa8, 0xab, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xe5, 0xca, 0x2a, 0x36, 0xe1, 0x03, 0x00, 0x00,
}
