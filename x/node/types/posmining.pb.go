// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: node/node/posmining.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Posmining struct {
	Owner           string     `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Coin            types.Coin `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin"`
	LastTransaction string     `protobuf:"bytes,3,opt,name=lastTransaction,proto3" json:"lastTransaction,omitempty"`
}

func (m *Posmining) Reset()         { *m = Posmining{} }
func (m *Posmining) String() string { return proto.CompactTextString(m) }
func (*Posmining) ProtoMessage()    {}
func (*Posmining) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a20eea0adfb6f72, []int{0}
}
func (m *Posmining) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Posmining) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Posmining.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Posmining) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Posmining.Merge(m, src)
}
func (m *Posmining) XXX_Size() int {
	return m.Size()
}
func (m *Posmining) XXX_DiscardUnknown() {
	xxx_messageInfo_Posmining.DiscardUnknown(m)
}

var xxx_messageInfo_Posmining proto.InternalMessageInfo

func (m *Posmining) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Posmining) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

func (m *Posmining) GetLastTransaction() string {
	if m != nil {
		return m.LastTransaction
	}
	return ""
}

func init() {
	proto.RegisterType((*Posmining)(nil), "node.node.Posmining")
}

func init() { proto.RegisterFile("node/node/posmining.proto", fileDescriptor_5a20eea0adfb6f72) }

var fileDescriptor_5a20eea0adfb6f72 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x28, 0x48, 0x31, 0x03, 0x52, 0xd4, 0x21, 0xed, 0x60, 0x2a, 0xa6, 0x2c, 0xb5,
	0x55, 0xfa, 0x06, 0x65, 0x60, 0x45, 0x15, 0x13, 0x9b, 0x6d, 0xac, 0x60, 0x89, 0xf8, 0x2c, 0xdb,
	0x05, 0xba, 0xf1, 0x08, 0x3c, 0x56, 0xc7, 0x8c, 0x4c, 0x08, 0x25, 0x2f, 0x82, 0x6c, 0xc3, 0xd2,
	0xe5, 0xb7, 0xef, 0xbf, 0xff, 0x3e, 0xe9, 0x0e, 0xcf, 0x0c, 0x3c, 0x29, 0x96, 0xc4, 0x82, 0xef,
	0xb4, 0xd1, 0xa6, 0xa5, 0xd6, 0x41, 0x80, 0xaa, 0x8c, 0x2e, 0x8d, 0x32, 0x9f, 0xb6, 0xd0, 0x42,
	0x72, 0x59, 0xfc, 0xe5, 0xc0, 0x9c, 0x48, 0xf0, 0x1d, 0x78, 0x26, 0xb8, 0x57, 0xec, 0x75, 0x25,
	0x54, 0xe0, 0x2b, 0x26, 0x41, 0x9b, 0xdc, 0xbf, 0xfe, 0x40, 0xb8, 0xbc, 0xff, 0x87, 0x56, 0x53,
	0x7c, 0x06, 0x6f, 0x46, 0xb9, 0x1a, 0x2d, 0x50, 0x53, 0x6e, 0x73, 0x51, 0xad, 0xf1, 0x24, 0x4e,
	0xd4, 0x27, 0x0b, 0xd4, 0x5c, 0xdc, 0xcc, 0x68, 0x46, 0xd2, 0x88, 0xa4, 0x7f, 0x48, 0x7a, 0x0b,
	0xda, 0x6c, 0x26, 0x87, 0xef, 0xab, 0x62, 0x9b, 0xc2, 0x55, 0x83, 0x2f, 0x5f, 0xb8, 0x0f, 0x0f,
	0x8e, 0x1b, 0xcf, 0x65, 0xd0, 0x60, 0xea, 0xd3, 0x04, 0x3d, 0xb6, 0x37, 0x77, 0x87, 0x81, 0xa0,
	0x7e, 0x20, 0xe8, 0x67, 0x20, 0xe8, 0x73, 0x24, 0x45, 0x3f, 0x92, 0xe2, 0x6b, 0x24, 0xc5, 0xe3,
	0xb2, 0xd5, 0xe1, 0x79, 0x27, 0xa8, 0x84, 0x8e, 0xc1, 0xce, 0x81, 0x00, 0x07, 0x7e, 0x29, 0xdd,
	0xde, 0x06, 0xc8, 0xf7, 0x78, 0xcf, 0x4f, 0xd8, 0x5b, 0xe5, 0xc5, 0x79, 0x5a, 0x69, 0xfd, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x1b, 0x31, 0x1d, 0x2a, 0x30, 0x01, 0x00, 0x00,
}

func (m *Posmining) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Posmining) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Posmining) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LastTransaction) > 0 {
		i -= len(m.LastTransaction)
		copy(dAtA[i:], m.LastTransaction)
		i = encodeVarintPosmining(dAtA, i, uint64(len(m.LastTransaction)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPosmining(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintPosmining(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPosmining(dAtA []byte, offset int, v uint64) int {
	offset -= sovPosmining(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Posmining) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovPosmining(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovPosmining(uint64(l))
	l = len(m.LastTransaction)
	if l > 0 {
		n += 1 + l + sovPosmining(uint64(l))
	}
	return n
}

func sovPosmining(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPosmining(x uint64) (n int) {
	return sovPosmining(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Posmining) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPosmining
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
			return fmt.Errorf("proto: Posmining: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Posmining: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosmining
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
				return ErrInvalidLengthPosmining
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPosmining
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosmining
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPosmining
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPosmining
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTransaction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosmining
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
				return ErrInvalidLengthPosmining
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPosmining
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastTransaction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPosmining(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPosmining
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
func skipPosmining(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPosmining
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
					return 0, ErrIntOverflowPosmining
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
					return 0, ErrIntOverflowPosmining
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
				return 0, ErrInvalidLengthPosmining
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPosmining
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPosmining
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPosmining        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPosmining          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPosmining = fmt.Errorf("proto: unexpected end of group")
)
