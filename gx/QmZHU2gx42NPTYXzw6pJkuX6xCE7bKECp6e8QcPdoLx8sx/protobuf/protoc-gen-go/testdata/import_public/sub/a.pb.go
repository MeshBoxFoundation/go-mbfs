// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/sub/a.proto

package sub // import "github.com/golang/protobuf/protoc-gen-go/testdata/import_public/sub"

import proto "mbfs/go-mbfs/gx/QmZHU2gx42NPTYXzw6pJkuX6xCE7bKECp6e8QcPdoLx8sx/protobuf/proto"
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

type E int32

const (
	E_ZERO E = 0
)

var E_name = map[int32]string{
	0: "ZERO",
}
var E_value = map[string]int32{
	"ZERO": 0,
}

func (x E) String() string {
	return proto.EnumName(E_name, int32(x))
}
func (E) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a_91ca0264a534463a, []int{0}
}

type M struct {
	// Field using a type in the same Go package, but a different source file.
	M2                   *M2      `protobuf:"bytes,1,opt,name=m2,proto3" json:"m2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M) Reset()         { *m = M{} }
func (m *M) String() string { return proto.CompactTextString(m) }
func (*M) ProtoMessage()    {}
func (*M) Descriptor() ([]byte, []int) {
	return fileDescriptor_a_91ca0264a534463a, []int{0}
}
func (m *M) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M.Unmarshal(m, b)
}
func (m *M) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M.Marshal(b, m, deterministic)
}
func (dst *M) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M.Merge(dst, src)
}
func (m *M) XXX_Size() int {
	return xxx_messageInfo_M.Size(m)
}
func (m *M) XXX_DiscardUnknown() {
	xxx_messageInfo_M.DiscardUnknown(m)
}

var xxx_messageInfo_M proto.InternalMessageInfo

func (m *M) GetM2() *M2 {
	if m != nil {
		return m.M2
	}
	return nil
}

func init() {
	proto.RegisterType((*M)(nil), "goproto.test.import_public.sub.M")
	proto.RegisterEnum("goproto.test.import_public.sub.E", E_name, E_value)
}

func init() { proto.RegisterFile("import_public/sub/a.proto", fileDescriptor_a_91ca0264a534463a) }

var fileDescriptor_a_91ca0264a534463a = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x89, 0x2f, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xd6, 0x2f, 0x2e, 0x4d, 0xd2, 0x4f, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4b, 0xcf, 0x07, 0x33, 0xf4, 0x4a, 0x52, 0x8b, 0x4b,
	0xf4, 0x50, 0xd4, 0xe9, 0x15, 0x97, 0x26, 0x49, 0x61, 0xd1, 0x9a, 0x04, 0xd1, 0xaa, 0x64, 0xce,
	0xc5, 0xe8, 0x2b, 0x64, 0xc4, 0xc5, 0x94, 0x6b, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4,
	0xa4, 0x87, 0xdf, 0x30, 0x3d, 0x5f, 0xa3, 0x20, 0xa6, 0x5c, 0x23, 0x2d, 0x5e, 0x2e, 0x46, 0x57,
	0x21, 0x0e, 0x2e, 0x96, 0x28, 0xd7, 0x20, 0x7f, 0x01, 0x06, 0x27, 0xd7, 0x28, 0xe7, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xf4, 0xfc, 0x9c, 0xc4, 0xbc, 0x74, 0x7d,
	0xb0, 0x39, 0x49, 0xa5, 0x69, 0x10, 0x46, 0xb2, 0x6e, 0x7a, 0x6a, 0x9e, 0x6e, 0x7a, 0xbe, 0x3e,
	0xc8, 0xe0, 0x94, 0xc4, 0x92, 0x44, 0x7d, 0x0c, 0x67, 0x25, 0xb1, 0x81, 0x55, 0x1a, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x81, 0xcc, 0x07, 0x7d, 0xed, 0x00, 0x00, 0x00,
}
