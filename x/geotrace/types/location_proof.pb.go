// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: geotrace/geotrace/location_proof.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LocationProof struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Lat       string `protobuf:"bytes,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon       string `protobuf:"bytes,3,opt,name=lon,proto3" json:"lon,omitempty"`
	Timestamp string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *LocationProof) Reset()         { *m = LocationProof{} }
func (m *LocationProof) String() string { return proto.CompactTextString(m) }
func (*LocationProof) ProtoMessage()    {}
func (*LocationProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a981f916f0673c, []int{0}
}
func (m *LocationProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocationProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocationProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LocationProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationProof.Merge(m, src)
}
func (m *LocationProof) XXX_Size() int {
	return m.Size()
}
func (m *LocationProof) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationProof.DiscardUnknown(m)
}

var xxx_messageInfo_LocationProof proto.InternalMessageInfo

func (m *LocationProof) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *LocationProof) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *LocationProof) GetLon() string {
	if m != nil {
		return m.Lon
	}
	return ""
}

func (m *LocationProof) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func init() {
	proto.RegisterType((*LocationProof)(nil), "geotrace.geotrace.LocationProof")
}

func init() {
	proto.RegisterFile("geotrace/geotrace/location_proof.proto", fileDescriptor_f4a981f916f0673c)
}

var fileDescriptor_f4a981f916f0673c = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0x4f, 0xcd, 0x2f,
	0x29, 0x4a, 0x4c, 0x4e, 0xd5, 0x87, 0x33, 0x72, 0xf2, 0x93, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0xe2,
	0x0b, 0x8a, 0xf2, 0xf3, 0xd3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x61, 0xd2, 0x7a,
	0x30, 0x86, 0x52, 0x36, 0x17, 0xaf, 0x0f, 0x54, 0x69, 0x00, 0x48, 0xa5, 0x90, 0x04, 0x17, 0x7b,
	0x72, 0x51, 0x6a, 0x62, 0x49, 0x7e, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b,
	0x24, 0xc0, 0xc5, 0x9c, 0x93, 0x58, 0x22, 0xc1, 0x04, 0x16, 0x05, 0x31, 0xc1, 0x22, 0xf9, 0x79,
	0x12, 0xcc, 0x50, 0x91, 0xfc, 0x3c, 0x21, 0x19, 0x2e, 0xce, 0x92, 0xcc, 0xdc, 0xd4, 0xe2, 0x92,
	0xc4, 0xdc, 0x02, 0x09, 0x16, 0xb0, 0x38, 0x42, 0xc0, 0xc9, 0xf8, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x24, 0xe1, 0x0e, 0xaf, 0x40, 0xf8, 0xa1, 0xa4, 0xb2, 0x20, 0xb5,
	0x38, 0x89, 0x0d, 0xec, 0x76, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x60, 0xf7, 0x1b,
	0xe5, 0x00, 0x00, 0x00,
}

func (m *LocationProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocationProof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LocationProof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Timestamp) > 0 {
		i -= len(m.Timestamp)
		copy(dAtA[i:], m.Timestamp)
		i = encodeVarintLocationProof(dAtA, i, uint64(len(m.Timestamp)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Lon) > 0 {
		i -= len(m.Lon)
		copy(dAtA[i:], m.Lon)
		i = encodeVarintLocationProof(dAtA, i, uint64(len(m.Lon)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Lat) > 0 {
		i -= len(m.Lat)
		copy(dAtA[i:], m.Lat)
		i = encodeVarintLocationProof(dAtA, i, uint64(len(m.Lat)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintLocationProof(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLocationProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocationProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LocationProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovLocationProof(uint64(l))
	}
	l = len(m.Lat)
	if l > 0 {
		n += 1 + l + sovLocationProof(uint64(l))
	}
	l = len(m.Lon)
	if l > 0 {
		n += 1 + l + sovLocationProof(uint64(l))
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovLocationProof(uint64(l))
	}
	return n
}

func sovLocationProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocationProof(x uint64) (n int) {
	return sovLocationProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LocationProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocationProof
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LocationProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocationProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocationProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLocationProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocationProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lat", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocationProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLocationProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocationProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lat = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocationProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLocationProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocationProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocationProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLocationProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocationProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocationProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLocationProof
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
func skipLocationProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocationProof
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
					return 0, ErrIntOverflowLocationProof
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLocationProof
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
			if length < 0 {
				return 0, ErrInvalidLengthLocationProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLocationProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLocationProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLocationProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocationProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLocationProof = fmt.Errorf("proto: unexpected end of group")
)
