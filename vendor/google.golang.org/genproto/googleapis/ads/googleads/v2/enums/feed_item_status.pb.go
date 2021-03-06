// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/feed_item_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Possible statuses of a feed item.
type FeedItemStatusEnum_FeedItemStatus int32

const (
	// Not specified.
	FeedItemStatusEnum_UNSPECIFIED FeedItemStatusEnum_FeedItemStatus = 0
	// Used for return value only. Represents value unknown in this version.
	FeedItemStatusEnum_UNKNOWN FeedItemStatusEnum_FeedItemStatus = 1
	// Feed item is enabled.
	FeedItemStatusEnum_ENABLED FeedItemStatusEnum_FeedItemStatus = 2
	// Feed item has been removed.
	FeedItemStatusEnum_REMOVED FeedItemStatusEnum_FeedItemStatus = 3
)

var FeedItemStatusEnum_FeedItemStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var FeedItemStatusEnum_FeedItemStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x FeedItemStatusEnum_FeedItemStatus) String() string {
	return proto.EnumName(FeedItemStatusEnum_FeedItemStatus_name, int32(x))
}

func (FeedItemStatusEnum_FeedItemStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33adf1cbb79e30ca, []int{0, 0}
}

// Container for enum describing possible statuses of a feed item.
type FeedItemStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItemStatusEnum) Reset()         { *m = FeedItemStatusEnum{} }
func (m *FeedItemStatusEnum) String() string { return proto.CompactTextString(m) }
func (*FeedItemStatusEnum) ProtoMessage()    {}
func (*FeedItemStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_33adf1cbb79e30ca, []int{0}
}

func (m *FeedItemStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemStatusEnum.Unmarshal(m, b)
}
func (m *FeedItemStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemStatusEnum.Marshal(b, m, deterministic)
}
func (m *FeedItemStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemStatusEnum.Merge(m, src)
}
func (m *FeedItemStatusEnum) XXX_Size() int {
	return xxx_messageInfo_FeedItemStatusEnum.Size(m)
}
func (m *FeedItemStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.FeedItemStatusEnum_FeedItemStatus", FeedItemStatusEnum_FeedItemStatus_name, FeedItemStatusEnum_FeedItemStatus_value)
	proto.RegisterType((*FeedItemStatusEnum)(nil), "google.ads.googleads.v2.enums.FeedItemStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/feed_item_status.proto", fileDescriptor_33adf1cbb79e30ca)
}

var fileDescriptor_33adf1cbb79e30ca = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdf, 0x4a, 0xfb, 0x30,
	0x14, 0xfe, 0xad, 0x83, 0x9f, 0x90, 0x81, 0x96, 0x7a, 0x27, 0xee, 0x62, 0x7b, 0x80, 0x04, 0xaa,
	0x57, 0xf1, 0x2a, 0x75, 0xd9, 0x1c, 0x6a, 0x37, 0x1c, 0xab, 0x20, 0xc5, 0x11, 0x4d, 0x0c, 0x85,
	0x35, 0x19, 0x4b, 0xba, 0x07, 0xf2, 0xd2, 0x47, 0xf1, 0x49, 0xc4, 0xa7, 0x90, 0xa4, 0x6b, 0x61,
	0x17, 0x7a, 0x53, 0xbe, 0x73, 0xbe, 0x3f, 0xfd, 0x72, 0xc0, 0xa5, 0xd4, 0x5a, 0xae, 0x05, 0x62,
	0xdc, 0xa0, 0x1a, 0x3a, 0xb4, 0x8b, 0x91, 0x50, 0x55, 0x69, 0xd0, 0x9b, 0x10, 0x7c, 0x55, 0x58,
	0x51, 0xae, 0x8c, 0x65, 0xb6, 0x32, 0x70, 0xb3, 0xd5, 0x56, 0x47, 0xfd, 0x5a, 0x0a, 0x19, 0x37,
	0xb0, 0x75, 0xc1, 0x5d, 0x0c, 0xbd, 0xeb, 0xec, 0xbc, 0x09, 0xdd, 0x14, 0x88, 0x29, 0xa5, 0x2d,
	0xb3, 0x85, 0x56, 0x7b, 0xf3, 0xf0, 0x19, 0x44, 0x63, 0x21, 0xf8, 0xd4, 0x8a, 0x72, 0xe1, 0x43,
	0xa9, 0xaa, 0xca, 0xe1, 0x0d, 0x38, 0x3e, 0xdc, 0x46, 0x27, 0xa0, 0xb7, 0x4c, 0x17, 0x73, 0x7a,
	0x3d, 0x1d, 0x4f, 0xe9, 0x28, 0xfc, 0x17, 0xf5, 0xc0, 0xd1, 0x32, 0xbd, 0x4d, 0x67, 0x8f, 0x69,
	0xd8, 0x71, 0x03, 0x4d, 0x49, 0x72, 0x47, 0x47, 0x61, 0xe0, 0x86, 0x07, 0x7a, 0x3f, 0xcb, 0xe8,
	0x28, 0xec, 0x26, 0x5f, 0x1d, 0x30, 0x78, 0xd5, 0x25, 0xfc, 0xb3, 0x63, 0x72, 0x7a, 0xf8, 0xb7,
	0xb9, 0xab, 0x36, 0xef, 0x3c, 0x25, 0x7b, 0x97, 0xd4, 0x6b, 0xa6, 0x24, 0xd4, 0x5b, 0x89, 0xa4,
	0x50, 0xbe, 0x78, 0x73, 0x9f, 0x4d, 0x61, 0x7e, 0x39, 0xd7, 0x95, 0xff, 0xbe, 0x07, 0xdd, 0x09,
	0x21, 0x1f, 0x41, 0x7f, 0x52, 0x47, 0x11, 0x6e, 0x60, 0x0d, 0x1d, 0xca, 0x62, 0xe8, 0xde, 0x6b,
	0x3e, 0x1b, 0x3e, 0x27, 0xdc, 0xe4, 0x2d, 0x9f, 0x67, 0x71, 0xee, 0xf9, 0xef, 0x60, 0x50, 0x2f,
	0x31, 0x26, 0xdc, 0x60, 0xdc, 0x2a, 0x30, 0xce, 0x62, 0x8c, 0xbd, 0xe6, 0xe5, 0xbf, 0x2f, 0x76,
	0xf1, 0x13, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x58, 0x9e, 0x52, 0xc6, 0x01, 0x00, 0x00,
}
