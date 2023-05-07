// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/superfluid-mosmo/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type GenesisState struct {
	Params                              Params                                `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	SuperfluidMosmoIntermediaryAccounts []*SuperfluidmOsmoIntermediaryAccount `protobuf:"bytes,2,rep,name=superfluid_mosmo_intermediary_accounts,json=superfluidMosmoIntermediaryAccounts,proto3" json:"superfluid_mosmo_intermediary_accounts,omitempty"`
	SuperfluidMosmoDelegationRecords    []*SuperfluidmOsmoDelegationRecord    `protobuf:"bytes,3,rep,name=superfluid_mosmo_delegation_records,json=superfluidMosmoDelegationRecords,proto3" json:"superfluid_mosmo_delegation_records,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59a2e040e8c174a, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetSuperfluidMosmoIntermediaryAccounts() []*SuperfluidmOsmoIntermediaryAccount {
	if m != nil {
		return m.SuperfluidMosmoIntermediaryAccounts
	}
	return nil
}

func (m *GenesisState) GetSuperfluidMosmoDelegationRecords() []*SuperfluidmOsmoDelegationRecord {
	if m != nil {
		return m.SuperfluidMosmoDelegationRecords
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "osmosis.superfluidmosmo.GenesisState")
}

func init() {
	proto.RegisterFile("osmosis/superfluid-mosmo/genesis.proto", fileDescriptor_b59a2e040e8c174a)
}

var fileDescriptor_b59a2e040e8c174a = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0x2f, 0xce, 0xcd,
	0x2f, 0xce, 0x2c, 0xd6, 0x2f, 0x2e, 0x2d, 0x48, 0x2d, 0x4a, 0xcb, 0x29, 0xcd, 0x4c, 0xd1, 0xcd,
	0x05, 0x89, 0xe9, 0xa7, 0xa7, 0xe6, 0xa5, 0x16, 0x67, 0x16, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4,
	0x0b, 0x89, 0x43, 0xd5, 0xe9, 0x21, 0xd4, 0x81, 0x95, 0x49, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83,
	0xd5, 0xe8, 0x83, 0x58, 0x10, 0xe5, 0x52, 0xaa, 0x38, 0x8d, 0x2d, 0x48, 0x2c, 0x4a, 0xcc, 0x85,
	0x9a, 0x2a, 0xa5, 0x8f, 0x53, 0x19, 0x42, 0x20, 0x1e, 0x2c, 0x00, 0xd1, 0xa0, 0xf4, 0x8a, 0x89,
	0x8b, 0xc7, 0x1d, 0xe2, 0xb0, 0xe0, 0x92, 0xc4, 0x92, 0x54, 0x21, 0x5b, 0x2e, 0x36, 0x88, 0x89,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xf2, 0x7a, 0x38, 0x1c, 0xaa, 0x17, 0x00, 0x56, 0xe6,
	0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x54, 0x93, 0xd0, 0x04, 0x46, 0x2e, 0x35, 0x74, 0xab,
	0xe2, 0x33, 0xf3, 0x4a, 0x52, 0x8b, 0x72, 0x53, 0x53, 0x32, 0x13, 0x8b, 0x2a, 0xe3, 0x13, 0x93,
	0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0x8a, 0x25, 0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0xac, 0x71, 0x9a,
	0x1f, 0x8c, 0xe0, 0xfb, 0x17, 0xe7, 0xe6, 0x7b, 0x22, 0x19, 0xe2, 0x08, 0x31, 0x23, 0x48, 0x19,
	0xa1, 0xc7, 0x37, 0x1f, 0xbb, 0x9a, 0x62, 0xa1, 0x76, 0x46, 0x2e, 0x65, 0x0c, 0x27, 0xa5, 0xa4,
	0xe6, 0xa4, 0xa6, 0x27, 0x96, 0x64, 0xe6, 0xe7, 0xc5, 0x17, 0xa5, 0x26, 0xe7, 0x17, 0xa5, 0x14,
	0x4b, 0x30, 0x83, 0xdd, 0x63, 0x41, 0xac, 0x7b, 0x5c, 0xe0, 0x26, 0x04, 0x81, 0x0d, 0x08, 0x52,
	0x40, 0x73, 0x0c, 0xba, 0x82, 0x62, 0xa7, 0xb0, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0x88, 0xb2, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0x85, 0x45, 0xa1,
	0x6e, 0x4e, 0x62, 0x52, 0x31, 0x3c, 0x3e, 0xcb, 0x0c, 0x4d, 0xf5, 0x2b, 0x30, 0x63, 0xb5, 0xa4,
	0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0x97, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97,
	0x14, 0x2d, 0x7a, 0x7c, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SuperfluidMosmoDelegationRecords) > 0 {
		for iNdEx := len(m.SuperfluidMosmoDelegationRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SuperfluidMosmoDelegationRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SuperfluidMosmoIntermediaryAccounts) > 0 {
		for iNdEx := len(m.SuperfluidMosmoIntermediaryAccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SuperfluidMosmoIntermediaryAccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.SuperfluidMosmoIntermediaryAccounts) > 0 {
		for _, e := range m.SuperfluidMosmoIntermediaryAccounts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SuperfluidMosmoDelegationRecords) > 0 {
		for _, e := range m.SuperfluidMosmoDelegationRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuperfluidMosmoIntermediaryAccounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SuperfluidMosmoIntermediaryAccounts = append(m.SuperfluidMosmoIntermediaryAccounts, &SuperfluidmOsmoIntermediaryAccount{})
			if err := m.SuperfluidMosmoIntermediaryAccounts[len(m.SuperfluidMosmoIntermediaryAccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuperfluidMosmoDelegationRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SuperfluidMosmoDelegationRecords = append(m.SuperfluidMosmoDelegationRecords, &SuperfluidmOsmoDelegationRecord{})
			if err := m.SuperfluidMosmoDelegationRecords[len(m.SuperfluidMosmoDelegationRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
