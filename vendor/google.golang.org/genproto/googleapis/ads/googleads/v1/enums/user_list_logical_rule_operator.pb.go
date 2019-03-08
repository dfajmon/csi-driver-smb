// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/user_list_logical_rule_operator.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible user list logical rule operators.
type UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator int32

const (
	// Not specified.
	UserListLogicalRuleOperatorEnum_UNSPECIFIED UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator = 0
	// Used for return value only. Represents value unknown in this version.
	UserListLogicalRuleOperatorEnum_UNKNOWN UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator = 1
	// And - all of the operands.
	UserListLogicalRuleOperatorEnum_ALL UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator = 2
	// Or - at least one of the operands.
	UserListLogicalRuleOperatorEnum_ANY UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator = 3
	// Not - none of the operands.
	UserListLogicalRuleOperatorEnum_NONE UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator = 4
)

var UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ALL",
	3: "ANY",
	4: "NONE",
}
var UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ALL":         2,
	"ANY":         3,
	"NONE":        4,
}

func (x UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator) String() string {
	return proto.EnumName(UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator_name, int32(x))
}
func (UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_list_logical_rule_operator_e05d8545d5195be3, []int{0, 0}
}

// The logical operator of the rule.
type UserListLogicalRuleOperatorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListLogicalRuleOperatorEnum) Reset()         { *m = UserListLogicalRuleOperatorEnum{} }
func (m *UserListLogicalRuleOperatorEnum) String() string { return proto.CompactTextString(m) }
func (*UserListLogicalRuleOperatorEnum) ProtoMessage()    {}
func (*UserListLogicalRuleOperatorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_list_logical_rule_operator_e05d8545d5195be3, []int{0}
}
func (m *UserListLogicalRuleOperatorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListLogicalRuleOperatorEnum.Unmarshal(m, b)
}
func (m *UserListLogicalRuleOperatorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListLogicalRuleOperatorEnum.Marshal(b, m, deterministic)
}
func (dst *UserListLogicalRuleOperatorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListLogicalRuleOperatorEnum.Merge(dst, src)
}
func (m *UserListLogicalRuleOperatorEnum) XXX_Size() int {
	return xxx_messageInfo_UserListLogicalRuleOperatorEnum.Size(m)
}
func (m *UserListLogicalRuleOperatorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListLogicalRuleOperatorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListLogicalRuleOperatorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserListLogicalRuleOperatorEnum)(nil), "google.ads.googleads.v1.enums.UserListLogicalRuleOperatorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator", UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator_name, UserListLogicalRuleOperatorEnum_UserListLogicalRuleOperator_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/user_list_logical_rule_operator.proto", fileDescriptor_user_list_logical_rule_operator_e05d8545d5195be3)
}

var fileDescriptor_user_list_logical_rule_operator_e05d8545d5195be3 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xc1, 0x6e, 0xe2, 0x30,
	0x14, 0x5c, 0x02, 0x5a, 0x56, 0xe6, 0xb0, 0x51, 0x8e, 0xbb, 0x45, 0x2d, 0x7c, 0x80, 0xa3, 0xa8,
	0x37, 0xf7, 0x14, 0x68, 0x8a, 0x50, 0x23, 0x83, 0x5a, 0x01, 0x6a, 0x15, 0x29, 0x72, 0x89, 0x65,
	0x45, 0x32, 0x76, 0xe4, 0x97, 0x70, 0xe8, 0xe7, 0xf4, 0xd8, 0x4f, 0xe9, 0xa7, 0xf4, 0xde, 0x7b,
	0x15, 0x1b, 0xb8, 0x95, 0x8b, 0x35, 0xf2, 0xcc, 0x9b, 0x79, 0x6f, 0xd0, 0x54, 0x68, 0x2d, 0x24,
	0x0f, 0x59, 0x01, 0xa1, 0x83, 0x2d, 0xda, 0x47, 0x21, 0x57, 0xcd, 0x0e, 0xc2, 0x06, 0xb8, 0xc9,
	0x65, 0x09, 0x75, 0x2e, 0xb5, 0x28, 0xb7, 0x4c, 0xe6, 0xa6, 0x91, 0x3c, 0xd7, 0x15, 0x37, 0xac,
	0xd6, 0x06, 0x57, 0x46, 0xd7, 0x3a, 0x18, 0xba, 0x49, 0xcc, 0x0a, 0xc0, 0x27, 0x13, 0xbc, 0x8f,
	0xb0, 0x35, 0xf9, 0x77, 0x71, 0xcc, 0xa8, 0xca, 0x90, 0x29, 0xa5, 0x6b, 0x56, 0x97, 0x5a, 0x81,
	0x1b, 0x1e, 0xbf, 0xa2, 0xcb, 0x15, 0x70, 0x93, 0x96, 0x50, 0xa7, 0x2e, 0xe3, 0xa1, 0x91, 0x7c,
	0x71, 0x48, 0x48, 0x54, 0xb3, 0x1b, 0x6f, 0xd0, 0xff, 0x33, 0x92, 0xe0, 0x2f, 0x1a, 0xac, 0xe8,
	0xe3, 0x32, 0x99, 0xce, 0xef, 0xe6, 0xc9, 0xad, 0xff, 0x2b, 0x18, 0xa0, 0xfe, 0x8a, 0xde, 0xd3,
	0xc5, 0x86, 0xfa, 0x9d, 0xa0, 0x8f, 0xba, 0x71, 0x9a, 0xfa, 0x9e, 0x05, 0xf4, 0xc9, 0xef, 0x06,
	0x7f, 0x50, 0x8f, 0x2e, 0x68, 0xe2, 0xf7, 0x26, 0x5f, 0x1d, 0x34, 0xda, 0xea, 0x1d, 0x3e, 0xbb,
	0xff, 0xe4, 0xea, 0x4c, 0xf8, 0xb2, 0xbd, 0x61, 0xd9, 0x79, 0x9e, 0x1c, 0x2c, 0x84, 0x96, 0x4c,
	0x09, 0xac, 0x8d, 0x08, 0x05, 0x57, 0xf6, 0xc2, 0x63, 0xaf, 0x55, 0x09, 0x3f, 0xd4, 0x7c, 0x63,
	0xdf, 0x37, 0xaf, 0x3b, 0x8b, 0xe3, 0x77, 0x6f, 0x38, 0x73, 0x56, 0x71, 0x01, 0xd8, 0xc1, 0x16,
	0xad, 0x23, 0xdc, 0x76, 0x01, 0x1f, 0x47, 0x3e, 0x8b, 0x0b, 0xc8, 0x4e, 0x7c, 0xb6, 0x8e, 0x32,
	0xcb, 0x7f, 0x7a, 0x23, 0xf7, 0x49, 0x48, 0x5c, 0x00, 0x21, 0x27, 0x05, 0x21, 0xeb, 0x88, 0x10,
	0xab, 0x79, 0xf9, 0x6d, 0x17, 0xbb, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x91, 0xa1, 0x11, 0x29,
	0xfe, 0x01, 0x00, 0x00,
}
