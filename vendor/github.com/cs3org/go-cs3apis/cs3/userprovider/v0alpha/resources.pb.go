// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/userprovider/v0alpha/resources.proto

package userproviderv0alphapb

import (
	fmt "fmt"
	types "github.com/cs3org/go-cs3apis/cs3/types"
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

type User struct {
	Id                   *types.UserId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Issuer               string        `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Subject              string        `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Username             string        `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Mail                 string        `protobuf:"bytes,5,opt,name=mail,proto3" json:"mail,omitempty"`
	DisplayName          string        `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Groups               []string      `protobuf:"bytes,7,rep,name=groups,proto3" json:"groups,omitempty"`
	Opaque               *types.Opaque `protobuf:"bytes,8,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ff151dbcd044679, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() *types.UserId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *User) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *User) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *User) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "cs3.userproviderv0alpha.User")
}

func init() {
	proto.RegisterFile("cs3/userprovider/v0alpha/resources.proto", fileDescriptor_7ff151dbcd044679)
}

var fileDescriptor_7ff151dbcd044679 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4e, 0x02, 0x31,
	0x14, 0x86, 0xd3, 0x01, 0x07, 0x2c, 0xc6, 0x68, 0x13, 0xa0, 0xc1, 0x0d, 0xb8, 0xc2, 0xcd, 0x40,
	0x9c, 0x13, 0x08, 0xb2, 0x60, 0x23, 0x93, 0x12, 0x88, 0x31, 0x26, 0xa6, 0xcc, 0x34, 0x5a, 0x03,
	0xb6, 0xf6, 0x51, 0x12, 0x0e, 0xe1, 0x25, 0x5c, 0x7a, 0x14, 0x6f, 0xe3, 0x0d, 0x4c, 0x3b, 0x83,
	0x31, 0x11, 0x37, 0x4d, 0xff, 0xf7, 0x7f, 0xef, 0xe5, 0xf5, 0x2f, 0xee, 0xa6, 0x10, 0xf7, 0x2c,
	0x08, 0xa3, 0x8d, 0xda, 0xc8, 0x4c, 0x98, 0xde, 0xa6, 0xcf, 0x97, 0xfa, 0x89, 0xf7, 0x8c, 0x00,
	0x65, 0x4d, 0x2a, 0x20, 0xd2, 0x46, 0xad, 0x15, 0x69, 0xa6, 0x10, 0x47, 0xbf, 0xc9, 0x02, 0x6c,
	0xd5, 0xdd, 0x88, 0xf5, 0x56, 0x0b, 0xc8, 0xcf, 0x9c, 0x3f, 0xff, 0x42, 0xb8, 0x3c, 0x03, 0x61,
	0x48, 0x07, 0x07, 0x32, 0xa3, 0xa8, 0x8d, 0xba, 0xb5, 0xcb, 0xd3, 0xc8, 0x4d, 0xc9, 0x31, 0x67,
	0x8e, 0x33, 0x16, 0xc8, 0x8c, 0x34, 0x70, 0x28, 0x01, 0xac, 0x30, 0x34, 0x68, 0xa3, 0xee, 0x21,
	0x2b, 0x14, 0xa1, 0xb8, 0x02, 0x76, 0xf1, 0x2c, 0xd2, 0x35, 0x2d, 0x79, 0x63, 0x27, 0x49, 0x0b,
	0x57, 0xdd, 0x2e, 0x2f, 0x7c, 0x25, 0x68, 0xd9, 0x5b, 0x3f, 0x9a, 0x10, 0x5c, 0x5e, 0x71, 0xb9,
	0xa4, 0x07, 0xbe, 0xee, 0xef, 0xa4, 0x83, 0x8f, 0x32, 0x09, 0x7a, 0xc9, 0xb7, 0x0f, 0xbe, 0x27,
	0xf4, 0x5e, 0xad, 0xa8, 0xdd, 0xb8, 0xb6, 0x06, 0x0e, 0x1f, 0x8d, 0xb2, 0x1a, 0x68, 0xa5, 0x5d,
	0x72, 0x4b, 0xe4, 0x8a, 0x5c, 0xe0, 0x50, 0x69, 0xfe, 0x6a, 0x05, 0xad, 0xfe, 0x79, 0xc3, 0xc4,
	0x1b, 0xac, 0x00, 0x06, 0x6f, 0x08, 0x9f, 0xa5, 0x6a, 0x15, 0xfd, 0x13, 0xd5, 0xe0, 0x98, 0xed,
	0x42, 0x4d, 0x5c, 0x46, 0x09, 0xba, 0xab, 0xef, 0xc1, 0xf4, 0xe2, 0x3d, 0x38, 0x19, 0x0e, 0x26,
	0xb7, 0xb3, 0xe9, 0x88, 0x25, 0x6c, 0x32, 0x1f, 0x5f, 0x8f, 0xd8, 0x47, 0xd0, 0x1c, 0x4e, 0x63,
	0x1f, 0x5a, 0x52, 0xe0, 0xf3, 0xfe, 0x95, 0xc3, 0x3f, 0xbd, 0x73, 0xbf, 0xc7, 0x59, 0x84, 0xfe,
	0x2b, 0xe2, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0xfd, 0x87, 0x2b, 0xe6, 0x01, 0x00, 0x00,
}