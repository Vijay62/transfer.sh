// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/mutate_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible mutate errors.
type MutateErrorEnum_MutateError int32

const (
	// Enum unspecified.
	MutateErrorEnum_UNSPECIFIED MutateErrorEnum_MutateError = 0
	// The received error code is not known in this version.
	MutateErrorEnum_UNKNOWN MutateErrorEnum_MutateError = 1
	// Requested resource was not found.
	MutateErrorEnum_RESOURCE_NOT_FOUND MutateErrorEnum_MutateError = 3
	// Cannot mutate the same resource twice in one request.
	MutateErrorEnum_ID_EXISTS_IN_MULTIPLE_MUTATES MutateErrorEnum_MutateError = 7
	// The field's contents don't match another field that represents the same
	// data.
	MutateErrorEnum_INCONSISTENT_FIELD_VALUES MutateErrorEnum_MutateError = 8
	// Mutates are not allowed for the requested resource.
	MutateErrorEnum_MUTATE_NOT_ALLOWED MutateErrorEnum_MutateError = 9
	// The resource isn't in Google Ads. It belongs to another ads system.
	MutateErrorEnum_RESOURCE_NOT_IN_GOOGLE_ADS MutateErrorEnum_MutateError = 10
)

var MutateErrorEnum_MutateError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "RESOURCE_NOT_FOUND",
	7:  "ID_EXISTS_IN_MULTIPLE_MUTATES",
	8:  "INCONSISTENT_FIELD_VALUES",
	9:  "MUTATE_NOT_ALLOWED",
	10: "RESOURCE_NOT_IN_GOOGLE_ADS",
}
var MutateErrorEnum_MutateError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"RESOURCE_NOT_FOUND":            3,
	"ID_EXISTS_IN_MULTIPLE_MUTATES": 7,
	"INCONSISTENT_FIELD_VALUES":     8,
	"MUTATE_NOT_ALLOWED":            9,
	"RESOURCE_NOT_IN_GOOGLE_ADS":    10,
}

func (x MutateErrorEnum_MutateError) String() string {
	return proto.EnumName(MutateErrorEnum_MutateError_name, int32(x))
}
func (MutateErrorEnum_MutateError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mutate_error_91b1822685171f8a, []int{0, 0}
}

// Container for enum describing possible mutate errors.
type MutateErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateErrorEnum) Reset()         { *m = MutateErrorEnum{} }
func (m *MutateErrorEnum) String() string { return proto.CompactTextString(m) }
func (*MutateErrorEnum) ProtoMessage()    {}
func (*MutateErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_mutate_error_91b1822685171f8a, []int{0}
}
func (m *MutateErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateErrorEnum.Unmarshal(m, b)
}
func (m *MutateErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateErrorEnum.Marshal(b, m, deterministic)
}
func (dst *MutateErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateErrorEnum.Merge(dst, src)
}
func (m *MutateErrorEnum) XXX_Size() int {
	return xxx_messageInfo_MutateErrorEnum.Size(m)
}
func (m *MutateErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MutateErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MutateErrorEnum)(nil), "google.ads.googleads.v0.errors.MutateErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.MutateErrorEnum_MutateError", MutateErrorEnum_MutateError_name, MutateErrorEnum_MutateError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/mutate_error.proto", fileDescriptor_mutate_error_91b1822685171f8a)
}

var fileDescriptor_mutate_error_91b1822685171f8a = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0xea, 0xd3, 0x30,
	0x14, 0xc6, 0xed, 0xfe, 0xe0, 0x34, 0xbb, 0x58, 0xc8, 0x85, 0xa0, 0xb0, 0x81, 0x7d, 0x80, 0xb4,
	0xe2, 0x5d, 0xbc, 0xca, 0xd6, 0xac, 0x04, 0xbb, 0xa4, 0x2c, 0x6d, 0x27, 0x52, 0x08, 0xd5, 0x96,
	0x22, 0x6c, 0xcb, 0x68, 0xb6, 0x3d, 0x90, 0x97, 0xbe, 0x83, 0x2f, 0xe0, 0x9d, 0xaf, 0xa1, 0x2f,
	0x21, 0x6d, 0xdc, 0x98, 0x17, 0xfe, 0xaf, 0xf2, 0xe5, 0xf0, 0xfb, 0xce, 0xc9, 0xf9, 0x02, 0xde,
	0xb4, 0xc6, 0xb4, 0xbb, 0x26, 0xa8, 0x6a, 0x1b, 0x38, 0xd9, 0xab, 0x4b, 0x18, 0x34, 0x5d, 0x67,
	0x3a, 0x1b, 0xec, 0xcf, 0xa7, 0xea, 0xd4, 0xe8, 0xe1, 0x86, 0x8f, 0x9d, 0x39, 0x19, 0x34, 0x77,
	0x1c, 0xae, 0x6a, 0x8b, 0x6f, 0x16, 0x7c, 0x09, 0xb1, 0xb3, 0xf8, 0x3f, 0x3d, 0x30, 0x5d, 0x0f,
	0x36, 0xd6, 0x17, 0xd8, 0xe1, 0xbc, 0xf7, 0xbf, 0x7b, 0x60, 0x72, 0x57, 0x43, 0x53, 0x30, 0xc9,
	0x85, 0x4a, 0xd9, 0x92, 0xaf, 0x38, 0x8b, 0xe0, 0x13, 0x34, 0x01, 0xe3, 0x5c, 0xbc, 0x17, 0x72,
	0x2b, 0xa0, 0x87, 0x5e, 0x00, 0xb4, 0x61, 0x4a, 0xe6, 0x9b, 0x25, 0xd3, 0x42, 0x66, 0x7a, 0x25,
	0x73, 0x11, 0xc1, 0x07, 0xf4, 0x1a, 0xcc, 0x78, 0xa4, 0xd9, 0x07, 0xae, 0x32, 0xa5, 0xb9, 0xd0,
	0xeb, 0x3c, 0xc9, 0x78, 0x9a, 0x30, 0xbd, 0xce, 0x33, 0x9a, 0x31, 0x05, 0xc7, 0x68, 0x06, 0x5e,
	0x72, 0xb1, 0x94, 0x42, 0x71, 0x95, 0x31, 0x91, 0xe9, 0x15, 0x67, 0x49, 0xa4, 0x0b, 0x9a, 0xe4,
	0x4c, 0xc1, 0x67, 0x7d, 0x67, 0xc7, 0x0e, 0x7d, 0x69, 0x92, 0xc8, 0x2d, 0x8b, 0xe0, 0x73, 0x34,
	0x07, 0xaf, 0xfe, 0x99, 0xc8, 0x85, 0x8e, 0xa5, 0x8c, 0x13, 0xa6, 0x69, 0xa4, 0x20, 0x58, 0xfc,
	0xf6, 0x80, 0xff, 0xd9, 0xec, 0xf1, 0xe3, 0xab, 0x2f, 0xe0, 0xdd, 0x8e, 0x69, 0x1f, 0x56, 0xea,
	0x7d, 0x8c, 0xfe, 0x7a, 0x5a, 0xb3, 0xab, 0x0e, 0x2d, 0x36, 0x5d, 0x1b, 0xb4, 0xcd, 0x61, 0x88,
	0xf2, 0x9a, 0xf8, 0xf1, 0x8b, 0xfd, 0xdf, 0x07, 0xbc, 0x73, 0xc7, 0xd7, 0xd1, 0x43, 0x4c, 0xe9,
	0xb7, 0xd1, 0x3c, 0x76, 0xcd, 0x68, 0x6d, 0xb1, 0x93, 0xbd, 0x2a, 0x42, 0x3c, 0x8c, 0xb4, 0x3f,
	0xae, 0x40, 0x49, 0x6b, 0x5b, 0xde, 0x80, 0xb2, 0x08, 0x4b, 0x07, 0xfc, 0x1a, 0xf9, 0xae, 0x4a,
	0x08, 0xad, 0x2d, 0x21, 0x37, 0x84, 0x90, 0x22, 0x24, 0xc4, 0x41, 0x9f, 0x9e, 0x0e, 0xaf, 0x7b,
	0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x58, 0x98, 0xf4, 0xa2, 0x1d, 0x02, 0x00, 0x00,
}
