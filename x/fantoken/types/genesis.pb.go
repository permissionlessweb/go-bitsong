// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fantoken/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the fantoken module's genesis state
type GenesisState struct {
	Params      Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Tokens      []FanToken   `protobuf:"bytes,2,rep,name=tokens,proto3" json:"tokens"`
	BurnedCoins []types.Coin `protobuf:"bytes,3,rep,name=burned_coins,json=burnedCoins,proto3" json:"burned_coins"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_43788810d38dc1c4, []int{0}
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

func (m *GenesisState) GetTokens() []FanToken {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *GenesisState) GetBurnedCoins() []types.Coin {
	if m != nil {
		return m.BurnedCoins
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "bitsong.fantoken.GenesisState")
}

func init() { proto.RegisterFile("fantoken/genesis.proto", fileDescriptor_43788810d38dc1c4) }

var fileDescriptor_43788810d38dc1c4 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x80, 0x63, 0x8a, 0x3a, 0x24, 0x1d, 0x50, 0x84, 0x20, 0x64, 0x30, 0x15, 0x53, 0x27, 0x5b,
	0x6d, 0x25, 0xc4, 0x1c, 0x24, 0x98, 0x90, 0x10, 0x30, 0xb1, 0x20, 0x3b, 0x38, 0xc1, 0x82, 0xf8,
	0xa2, 0x9c, 0x8b, 0xe0, 0x2d, 0x78, 0x28, 0x86, 0x8e, 0x1d, 0x99, 0x10, 0x4a, 0x5e, 0x04, 0xe5,
	0xc7, 0x1d, 0x60, 0xbb, 0xe4, 0xbb, 0x4f, 0xfe, 0x74, 0xfe, 0x41, 0x26, 0x8c, 0x85, 0x67, 0x65,
	0x78, 0xae, 0x8c, 0x42, 0x8d, 0xac, 0xac, 0xc0, 0x42, 0xb8, 0x27, 0xb5, 0x45, 0x30, 0x39, 0x73,
	0x3c, 0xde, 0xcf, 0x21, 0x87, 0x0e, 0xf2, 0x76, 0xea, 0xf7, 0xe2, 0xc3, 0xad, 0xef, 0x86, 0x01,
	0xd0, 0x14, 0xb0, 0x00, 0xe4, 0x52, 0xa0, 0xe2, 0xaf, 0x73, 0xa9, 0xac, 0x98, 0xf3, 0x14, 0xf4,
	0xc0, 0x4f, 0x3e, 0x89, 0x3f, 0xb9, 0xec, 0x9f, 0xbc, 0xb5, 0xc2, 0xaa, 0xf0, 0xd4, 0x1f, 0x97,
	0xa2, 0x12, 0x05, 0x46, 0x64, 0x4a, 0x66, 0xc1, 0x22, 0x62, 0x7f, 0x13, 0xd8, 0x75, 0xc7, 0x93,
	0xdd, 0xf5, 0xf7, 0xb1, 0x77, 0x33, 0x6c, 0x87, 0x67, 0xfe, 0xb8, 0xa3, 0x18, 0xed, 0x4c, 0x47,
	0xb3, 0x60, 0x11, 0xff, 0xf7, 0x2e, 0x84, 0xb9, 0x6b, 0x07, 0x67, 0xf6, 0xfb, 0x61, 0xe2, 0x4f,
	0xe4, 0xaa, 0x32, 0xea, 0xf1, 0xa1, 0xed, 0xc2, 0x68, 0xd4, 0xf9, 0x47, 0xac, 0x2f, 0x67, 0x6d,
	0x39, 0x1b, 0xca, 0xd9, 0x39, 0x68, 0xa7, 0x07, 0xbd, 0xd4, 0xfe, 0xc1, 0xe4, 0x6a, 0x5d, 0x53,
	0xb2, 0xa9, 0x29, 0xf9, 0xa9, 0x29, 0xf9, 0x68, 0xa8, 0xb7, 0x69, 0xa8, 0xf7, 0xd5, 0x50, 0xef,
	0x7e, 0x99, 0x6b, 0xfb, 0xb4, 0x92, 0x2c, 0x85, 0x82, 0x0f, 0x45, 0x90, 0x65, 0x3a, 0xd5, 0xe2,
	0xc5, 0x7d, 0xf3, 0xb7, 0xed, 0xd5, 0xb8, 0x7d, 0x2f, 0x15, 0xca, 0x71, 0x77, 0x9c, 0xe5, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0x42, 0xca, 0xbe, 0x97, 0x01, 0x00, 0x00,
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
	if len(m.BurnedCoins) > 0 {
		for iNdEx := len(m.BurnedCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BurnedCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BurnedCoins) > 0 {
		for _, e := range m.BurnedCoins {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
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
			m.Tokens = append(m.Tokens, FanToken{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnedCoins", wireType)
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
			m.BurnedCoins = append(m.BurnedCoins, types.Coin{})
			if err := m.BurnedCoins[len(m.BurnedCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
