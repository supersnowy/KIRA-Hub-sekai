// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/ubi/ubi.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type UBIRecord struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// distribution-start & distribution-end - defines the exact time period (Unix
	// timestamps) between which tokens will be distributed to the pool, allowing
	// for a precise funds spending.
	DistributionStart uint64 `protobuf:"varint,2,opt,name=distribution_start,json=distributionStart,proto3" json:"distribution_start,omitempty"`
	DistributionEnd   uint64 `protobuf:"varint,3,opt,name=distribution_end,json=distributionEnd,proto3" json:"distribution_end,omitempty"`
	// distribution-last- timestamp of the last distribution
	DistributionLast uint64 `protobuf:"varint,4,opt,name=distribution_last,json=distributionLast,proto3" json:"distribution_last,omitempty"`
	// amount - the amount of KEX tokens to be minted and distributed every period
	// number of seconds into pool
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount" yaml:"amount"`
	// period - time in seconds that must elapse sincedistribution-last
	// for the funds to be distributed automatically into pool
	Period uint64 `protobuf:"varint,6,opt,name=period,proto3" json:"period,omitempty"`
	// pool - spending pool name where exact amount of KEX should be deposited
	Pool string `protobuf:"bytes,7,opt,name=pool,proto3" json:"pool,omitempty"`
}

func (m *UBIRecord) Reset()         { *m = UBIRecord{} }
func (m *UBIRecord) String() string { return proto.CompactTextString(m) }
func (*UBIRecord) ProtoMessage()    {}
func (*UBIRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_27b7349fe391b1c8, []int{0}
}
func (m *UBIRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UBIRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UBIRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UBIRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UBIRecord.Merge(m, src)
}
func (m *UBIRecord) XXX_Size() int {
	return m.Size()
}
func (m *UBIRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_UBIRecord.DiscardUnknown(m)
}

var xxx_messageInfo_UBIRecord proto.InternalMessageInfo

func (m *UBIRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UBIRecord) GetDistributionStart() uint64 {
	if m != nil {
		return m.DistributionStart
	}
	return 0
}

func (m *UBIRecord) GetDistributionEnd() uint64 {
	if m != nil {
		return m.DistributionEnd
	}
	return 0
}

func (m *UBIRecord) GetDistributionLast() uint64 {
	if m != nil {
		return m.DistributionLast
	}
	return 0
}

func (m *UBIRecord) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *UBIRecord) GetPool() string {
	if m != nil {
		return m.Pool
	}
	return ""
}

func init() {
	proto.RegisterType((*UBIRecord)(nil), "kira.ubi.UBIRecord")
}

func init() { proto.RegisterFile("kira/ubi/ubi.proto", fileDescriptor_27b7349fe391b1c8) }

var fileDescriptor_27b7349fe391b1c8 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x6d, 0xe7, 0xac, 0x2e, 0x20, 0x6a, 0x10, 0x29, 0x3b, 0xb4, 0x63, 0xa0, 0x4c, 0x64, 0xcd,
	0xc1, 0x9b, 0x97, 0xc1, 0xc4, 0xc3, 0xd0, 0x53, 0x45, 0x04, 0x2f, 0x23, 0x5d, 0x63, 0x0d, 0x6b,
	0xfa, 0x95, 0x24, 0x05, 0xf7, 0x2f, 0xfc, 0x1f, 0xfe, 0x91, 0x1d, 0x77, 0x14, 0x0f, 0x43, 0xb6,
	0x7f, 0xe0, 0x2f, 0x90, 0xa4, 0x3b, 0x6c, 0x87, 0x90, 0x97, 0xf7, 0xbd, 0xbc, 0x3c, 0x5e, 0x10,
	0x9e, 0x72, 0x49, 0x49, 0x95, 0x70, 0xb3, 0xa2, 0x52, 0x82, 0x06, 0x7c, 0x68, 0xb8, 0xa8, 0x4a,
	0x78, 0x3b, 0xcc, 0x00, 0xb2, 0x9c, 0x11, 0xcb, 0x27, 0xd5, 0x1b, 0xd1, 0x5c, 0x30, 0xa5, 0xa9,
	0x28, 0x6b, 0x69, 0xfb, 0x2c, 0x83, 0x0c, 0x2c, 0x24, 0x06, 0xd5, 0x6c, 0xf7, 0xab, 0x81, 0x5a,
	0xcf, 0xc3, 0x51, 0xcc, 0x26, 0x20, 0x53, 0x8c, 0x51, 0xb3, 0xa0, 0x82, 0xf9, 0x6e, 0xc7, 0xed,
	0xb5, 0x62, 0x8b, 0x71, 0x1f, 0xe1, 0x94, 0x2b, 0x2d, 0x79, 0x52, 0x69, 0x0e, 0xc5, 0x58, 0x69,
	0x2a, 0xb5, 0xdf, 0xe8, 0xb8, 0xbd, 0x66, 0x7c, 0xba, 0x3d, 0x79, 0x32, 0x03, 0x7c, 0x85, 0x4e,
	0x76, 0xe4, 0xac, 0x48, 0xfd, 0x3d, 0x2b, 0x3e, 0xde, 0xe6, 0xef, 0x8b, 0x14, 0x5f, 0xa3, 0x9d,
	0xfb, 0xe3, 0x9c, 0x2a, 0xed, 0x37, 0xad, 0x76, 0xc7, 0xe3, 0x91, 0x2a, 0x8d, 0x5f, 0x90, 0x47,
	0x05, 0x54, 0x85, 0xf6, 0xf7, 0x4d, 0xb8, 0xe1, 0x60, 0xbe, 0x0c, 0x9d, 0x9f, 0x65, 0x78, 0x99,
	0x71, 0xfd, 0x5e, 0x25, 0xd1, 0x04, 0x04, 0x99, 0x80, 0x12, 0xa0, 0x36, 0x5b, 0x5f, 0xa5, 0x53,
	0xa2, 0x67, 0x25, 0x53, 0xd1, 0xa8, 0xd0, 0x7f, 0xcb, 0xf0, 0x68, 0x46, 0x45, 0x7e, 0xdb, 0xad,
	0x5d, 0xba, 0xf1, 0xc6, 0x0e, 0x9f, 0x23, 0xaf, 0x64, 0x92, 0x43, 0xea, 0x7b, 0xf6, 0xe9, 0xcd,
	0xc9, 0x74, 0x51, 0x02, 0xe4, 0xfe, 0x41, 0xdd, 0x85, 0xc1, 0xc3, 0xc1, 0x7c, 0x15, 0xb8, 0x8b,
	0x55, 0xe0, 0xfe, 0xae, 0x02, 0xf7, 0x73, 0x1d, 0x38, 0x8b, 0x75, 0xe0, 0x7c, 0xaf, 0x03, 0xe7,
	0xf5, 0x62, 0x2b, 0xc6, 0x03, 0x97, 0xf4, 0x0e, 0x24, 0x23, 0x8a, 0x4d, 0x29, 0x27, 0x1f, 0xf6,
	0xcf, 0x6c, 0x92, 0xc4, 0xb3, 0xad, 0xdf, 0xfc, 0x07, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xa7, 0x04,
	0x3a, 0xcc, 0x01, 0x00, 0x00,
}

func (m *UBIRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UBIRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UBIRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pool) > 0 {
		i -= len(m.Pool)
		copy(dAtA[i:], m.Pool)
		i = encodeVarintUbi(dAtA, i, uint64(len(m.Pool)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Period != 0 {
		i = encodeVarintUbi(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintUbi(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.DistributionLast != 0 {
		i = encodeVarintUbi(dAtA, i, uint64(m.DistributionLast))
		i--
		dAtA[i] = 0x20
	}
	if m.DistributionEnd != 0 {
		i = encodeVarintUbi(dAtA, i, uint64(m.DistributionEnd))
		i--
		dAtA[i] = 0x18
	}
	if m.DistributionStart != 0 {
		i = encodeVarintUbi(dAtA, i, uint64(m.DistributionStart))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUbi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUbi(dAtA []byte, offset int, v uint64) int {
	offset -= sovUbi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UBIRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUbi(uint64(l))
	}
	if m.DistributionStart != 0 {
		n += 1 + sovUbi(uint64(m.DistributionStart))
	}
	if m.DistributionEnd != 0 {
		n += 1 + sovUbi(uint64(m.DistributionEnd))
	}
	if m.DistributionLast != 0 {
		n += 1 + sovUbi(uint64(m.DistributionLast))
	}
	l = m.Amount.Size()
	n += 1 + l + sovUbi(uint64(l))
	if m.Period != 0 {
		n += 1 + sovUbi(uint64(m.Period))
	}
	l = len(m.Pool)
	if l > 0 {
		n += 1 + l + sovUbi(uint64(l))
	}
	return n
}

func sovUbi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUbi(x uint64) (n int) {
	return sovUbi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UBIRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUbi
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
			return fmt.Errorf("proto: UBIRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UBIRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbi
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
				return ErrInvalidLengthUbi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUbi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionStart", wireType)
			}
			m.DistributionStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistributionStart |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionEnd", wireType)
			}
			m.DistributionEnd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistributionEnd |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionLast", wireType)
			}
			m.DistributionLast = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistributionLast |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbi
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
				return ErrInvalidLengthUbi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUbi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUbi
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
				return ErrInvalidLengthUbi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUbi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pool = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUbi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUbi
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
func skipUbi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUbi
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
					return 0, ErrIntOverflowUbi
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
					return 0, ErrIntOverflowUbi
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
				return 0, ErrInvalidLengthUbi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUbi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUbi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUbi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUbi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUbi = fmt.Errorf("proto: unexpected end of group")
)