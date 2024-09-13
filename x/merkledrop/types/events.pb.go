// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitsong/merkledrop/v1beta1/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type EventCreate struct {
	Owner        string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	MerkledropId uint64 `protobuf:"varint,2,opt,name=merkledrop_id,json=merkledropId,proto3" json:"merkledrop_id,omitempty"`
}

func (m *EventCreate) Reset()         { *m = EventCreate{} }
func (m *EventCreate) String() string { return proto.CompactTextString(m) }
func (*EventCreate) ProtoMessage()    {}
func (*EventCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_3042ab6a9db80a59, []int{0}
}
func (m *EventCreate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreate.Merge(m, src)
}
func (m *EventCreate) XXX_Size() int {
	return m.Size()
}
func (m *EventCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreate.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreate proto.InternalMessageInfo

type EventClaim struct {
	MerkledropId uint64                                  `protobuf:"varint,1,opt,name=merkledrop_id,json=merkledropId,proto3" json:"merkledrop_id,omitempty"`
	Index        uint64                                  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Coin         github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=coin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"coin"`
}

func (m *EventClaim) Reset()         { *m = EventClaim{} }
func (m *EventClaim) String() string { return proto.CompactTextString(m) }
func (*EventClaim) ProtoMessage()    {}
func (*EventClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_3042ab6a9db80a59, []int{1}
}
func (m *EventClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventClaim.Merge(m, src)
}
func (m *EventClaim) XXX_Size() int {
	return m.Size()
}
func (m *EventClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_EventClaim.DiscardUnknown(m)
}

var xxx_messageInfo_EventClaim proto.InternalMessageInfo

type EventWithdraw struct {
	MerkledropId uint64                                  `protobuf:"varint,1,opt,name=merkledrop_id,json=merkledropId,proto3" json:"merkledrop_id,omitempty"`
	Coin         github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=coin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"coin"`
}

func (m *EventWithdraw) Reset()         { *m = EventWithdraw{} }
func (m *EventWithdraw) String() string { return proto.CompactTextString(m) }
func (*EventWithdraw) ProtoMessage()    {}
func (*EventWithdraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_3042ab6a9db80a59, []int{2}
}
func (m *EventWithdraw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventWithdraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventWithdraw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventWithdraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventWithdraw.Merge(m, src)
}
func (m *EventWithdraw) XXX_Size() int {
	return m.Size()
}
func (m *EventWithdraw) XXX_DiscardUnknown() {
	xxx_messageInfo_EventWithdraw.DiscardUnknown(m)
}

var xxx_messageInfo_EventWithdraw proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventCreate)(nil), "bitsong.merkledrop.v1beta1.EventCreate")
	proto.RegisterType((*EventClaim)(nil), "bitsong.merkledrop.v1beta1.EventClaim")
	proto.RegisterType((*EventWithdraw)(nil), "bitsong.merkledrop.v1beta1.EventWithdraw")
}

func init() {
	proto.RegisterFile("bitsong/merkledrop/v1beta1/events.proto", fileDescriptor_3042ab6a9db80a59)
}

var fileDescriptor_3042ab6a9db80a59 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x3d, 0x4f, 0x3a, 0x31,
	0x18, 0xbf, 0xf2, 0xe7, 0x6f, 0x62, 0x91, 0xe5, 0xc2, 0x80, 0x0c, 0x85, 0xe0, 0x00, 0x0b, 0x6d,
	0xd0, 0xc5, 0x19, 0x62, 0xa2, 0x2b, 0x83, 0x26, 0x0e, 0x9a, 0x7b, 0x29, 0x47, 0x03, 0x77, 0x0f,
	0x69, 0x2b, 0xe0, 0xb7, 0x70, 0xf0, 0x3b, 0xf8, 0x55, 0x18, 0x19, 0x8d, 0x03, 0x51, 0xf8, 0x22,
	0xe6, 0xda, 0xc3, 0x23, 0x61, 0x71, 0x71, 0xba, 0x7b, 0xda, 0xdf, 0x6b, 0xfa, 0xe0, 0x96, 0x2f,
	0xb4, 0x82, 0x24, 0x62, 0x31, 0x97, 0xe3, 0x09, 0x0f, 0x25, 0x4c, 0xd9, 0xac, 0xeb, 0x73, 0xed,
	0x75, 0x19, 0x9f, 0xf1, 0x44, 0x2b, 0x3a, 0x95, 0xa0, 0xc1, 0xad, 0x65, 0x40, 0x9a, 0x03, 0x69,
	0x06, 0xac, 0x55, 0x22, 0x88, 0xc0, 0xc0, 0x58, 0xfa, 0x67, 0x19, 0x35, 0x12, 0x80, 0x8a, 0x41,
	0x31, 0xdf, 0x53, 0xfc, 0x47, 0x33, 0x00, 0x91, 0xd8, 0xfb, 0xe6, 0x35, 0x2e, 0x5d, 0xa5, 0x0e,
	0x7d, 0xc9, 0x3d, 0xcd, 0xdd, 0x0a, 0xfe, 0x0f, 0xf3, 0x84, 0xcb, 0x2a, 0x6a, 0xa0, 0xf6, 0xf1,
	0xc0, 0x0e, 0xee, 0x19, 0x2e, 0xe7, 0x86, 0x8f, 0x22, 0xac, 0x16, 0x1a, 0xa8, 0x5d, 0x1c, 0x9c,
	0xe4, 0x87, 0x37, 0x61, 0xf3, 0x0d, 0x61, 0x6c, 0xa5, 0x26, 0x9e, 0x88, 0x0f, 0x39, 0xe8, 0x90,
	0x93, 0xda, 0x89, 0x24, 0xe4, 0x8b, 0x4c, 0xd0, 0x0e, 0xee, 0x03, 0x2e, 0xa6, 0x09, 0xab, 0xff,
	0x1a, 0xa8, 0x5d, 0x3a, 0x3f, 0xa5, 0xb6, 0x02, 0x4d, 0x2b, 0xec, 0xda, 0xd2, 0x3e, 0x88, 0xa4,
	0xc7, 0x96, 0xeb, 0xba, 0xf3, 0xb1, 0xae, 0xb7, 0x22, 0xa1, 0x47, 0x4f, 0x3e, 0x0d, 0x20, 0x66,
	0x59, 0x5f, 0xfb, 0xe9, 0xa8, 0x70, 0xcc, 0xf4, 0xf3, 0x94, 0x2b, 0x43, 0x18, 0x18, 0xdd, 0xe6,
	0x2b, 0xc2, 0x65, 0x93, 0xf4, 0x4e, 0xe8, 0x51, 0x28, 0xbd, 0xf9, 0xef, 0xc2, 0xee, 0x62, 0x15,
	0xfe, 0x26, 0x56, 0xef, 0x76, 0xf9, 0x45, 0x9c, 0xe5, 0x86, 0xa0, 0xd5, 0x86, 0xa0, 0xcf, 0x0d,
	0x41, 0x2f, 0x5b, 0xe2, 0xac, 0xb6, 0xc4, 0x79, 0xdf, 0x12, 0xe7, 0xfe, 0x72, 0x4f, 0x2c, 0xdb,
	0x02, 0x18, 0x0e, 0x45, 0x20, 0xbc, 0x09, 0x8b, 0xa0, 0xb3, 0xdb, 0xa0, 0xc5, 0xfe, 0x0e, 0x19,
	0x0b, 0xff, 0xc8, 0xbc, 0xf4, 0xc5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0xcc, 0x06, 0x92,
	0x66, 0x02, 0x00, 0x00,
}

func (m *EventCreate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MerkledropId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.MerkledropId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Coin.Size()
		i -= size
		if _, err := m.Coin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Index != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if m.MerkledropId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.MerkledropId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventWithdraw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventWithdraw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventWithdraw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Coin.Size()
		i -= size
		if _, err := m.Coin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.MerkledropId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.MerkledropId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.MerkledropId != 0 {
		n += 1 + sovEvents(uint64(m.MerkledropId))
	}
	return n
}

func (m *EventClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MerkledropId != 0 {
		n += 1 + sovEvents(uint64(m.MerkledropId))
	}
	if m.Index != 0 {
		n += 1 + sovEvents(uint64(m.Index))
	}
	l = m.Coin.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *EventWithdraw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MerkledropId != 0 {
		n += 1 + sovEvents(uint64(m.MerkledropId))
	}
	l = m.Coin.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkledropId", wireType)
			}
			m.MerkledropId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MerkledropId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkledropId", wireType)
			}
			m.MerkledropId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MerkledropId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventWithdraw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventWithdraw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventWithdraw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkledropId", wireType)
			}
			m.MerkledropId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MerkledropId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
