// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: manifest.proto

package protos

import proto "mbfs/go-mbfs/gx/QmZHU2gx42NPTYXzw6pJkuX6xCE7bKECp6e8QcPdoLx8sx/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ManifestChange_Operation int32

const (
	ManifestChange_CREATE ManifestChange_Operation = 0
	ManifestChange_DELETE ManifestChange_Operation = 1
)

var ManifestChange_Operation_name = map[int32]string{
	0: "CREATE",
	1: "DELETE",
}
var ManifestChange_Operation_value = map[string]int32{
	"CREATE": 0,
	"DELETE": 1,
}

func (x ManifestChange_Operation) String() string {
	return proto.EnumName(ManifestChange_Operation_name, int32(x))
}
func (ManifestChange_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorManifest, []int{1, 0}
}

type ManifestChangeSet struct {
	// A set of changes that are applied atomically.
	Changes []*ManifestChange `protobuf:"bytes,1,rep,name=changes" json:"changes,omitempty"`
}

func (m *ManifestChangeSet) Reset()                    { *m = ManifestChangeSet{} }
func (m *ManifestChangeSet) String() string            { return proto.CompactTextString(m) }
func (*ManifestChangeSet) ProtoMessage()               {}
func (*ManifestChangeSet) Descriptor() ([]byte, []int) { return fileDescriptorManifest, []int{0} }

func (m *ManifestChangeSet) GetChanges() []*ManifestChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

type ManifestChange struct {
	Id    uint64                   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Op    ManifestChange_Operation `protobuf:"varint,2,opt,name=Op,proto3,enum=protos.ManifestChange_Operation" json:"Op,omitempty"`
	Level uint32                   `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (m *ManifestChange) Reset()                    { *m = ManifestChange{} }
func (m *ManifestChange) String() string            { return proto.CompactTextString(m) }
func (*ManifestChange) ProtoMessage()               {}
func (*ManifestChange) Descriptor() ([]byte, []int) { return fileDescriptorManifest, []int{1} }

func (m *ManifestChange) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ManifestChange) GetOp() ManifestChange_Operation {
	if m != nil {
		return m.Op
	}
	return ManifestChange_CREATE
}

func (m *ManifestChange) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*ManifestChangeSet)(nil), "protos.ManifestChangeSet")
	proto.RegisterType((*ManifestChange)(nil), "protos.ManifestChange")
	proto.RegisterEnum("protos.ManifestChange_Operation", ManifestChange_Operation_name, ManifestChange_Operation_value)
}
func (m *ManifestChangeSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManifestChangeSet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Changes) > 0 {
		for _, msg := range m.Changes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintManifest(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ManifestChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManifestChange) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.Id))
	}
	if m.Op != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.Op))
	}
	if m.Level != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.Level))
	}
	return i, nil
}

func encodeFixed64Manifest(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Manifest(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintManifest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ManifestChangeSet) Size() (n int) {
	var l int
	_ = l
	if len(m.Changes) > 0 {
		for _, e := range m.Changes {
			l = e.Size()
			n += 1 + l + sovManifest(uint64(l))
		}
	}
	return n
}

func (m *ManifestChange) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovManifest(uint64(m.Id))
	}
	if m.Op != 0 {
		n += 1 + sovManifest(uint64(m.Op))
	}
	if m.Level != 0 {
		n += 1 + sovManifest(uint64(m.Level))
	}
	return n
}

func sovManifest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozManifest(x uint64) (n int) {
	return sovManifest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ManifestChangeSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManifest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ManifestChangeSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManifestChangeSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthManifest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Changes = append(m.Changes, &ManifestChange{})
			if err := m.Changes[len(m.Changes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipManifest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManifest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ManifestChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManifest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ManifestChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManifestChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= (ManifestChange_Operation(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipManifest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManifest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipManifest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowManifest
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthManifest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowManifest
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipManifest(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthManifest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowManifest   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("manifest.proto", fileDescriptorManifest) }

var fileDescriptorManifest = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0xcc, 0xcb,
	0x4c, 0x4b, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a,
	0xae, 0x5c, 0x82, 0xbe, 0x50, 0x19, 0xe7, 0x8c, 0xc4, 0xbc, 0xf4, 0xd4, 0xe0, 0xd4, 0x12, 0x21,
	0x03, 0x2e, 0xf6, 0x64, 0x30, 0xa7, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x0c, 0xa2,
	0xab, 0x58, 0x0f, 0x55, 0x6d, 0x10, 0x4c, 0x99, 0x52, 0x2f, 0x23, 0x17, 0x1f, 0xaa, 0x9c, 0x10,
	0x1f, 0x17, 0x93, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x93, 0x67, 0x8a, 0x90,
	0x01, 0x17, 0x93, 0x7f, 0x81, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0x02, 0x76, 0xf3, 0xf4,
	0xfc, 0x0b, 0x52, 0x8b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x82, 0x98, 0xfc, 0x0b, 0x84, 0x44, 0xb8,
	0x58, 0x7d, 0x52, 0xcb, 0x52, 0x73, 0x24, 0x98, 0x15, 0x18, 0x35, 0x78, 0x83, 0x20, 0x1c, 0x25,
	0x65, 0x2e, 0x4e, 0xb8, 0x32, 0x21, 0x2e, 0x2e, 0x36, 0xe7, 0x20, 0x57, 0xc7, 0x10, 0x57, 0x01,
	0x06, 0x10, 0xdb, 0xc5, 0xd5, 0xc7, 0x35, 0xc4, 0x55, 0x80, 0xd1, 0x49, 0xe0, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x09, 0xe2,
	0x61, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0x6f, 0x23, 0xc9, 0x09, 0x01, 0x00, 0x00,
}
