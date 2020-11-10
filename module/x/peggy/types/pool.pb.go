// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: module/peggy/v1beta1/pool.proto

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

// OutgoingTx is a withdrawal on the bridged contract
type OutgoingTx struct {
	Sender    string     `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	DestAddr  []byte     `protobuf:"bytes,2,opt,name=dest_addr,json=destAddr,proto3" json:"dest_addr,omitempty"`
	Amount    types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
	BridgeFee types.Coin `protobuf:"bytes,4,opt,name=bridge_fee,json=bridgeFee,proto3" json:"bridge_fee"`
}

func (m *OutgoingTx) Reset()         { *m = OutgoingTx{} }
func (m *OutgoingTx) String() string { return proto.CompactTextString(m) }
func (*OutgoingTx) ProtoMessage()    {}
func (*OutgoingTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_42e6de9b297dce86, []int{0}
}
func (m *OutgoingTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutgoingTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTx.Merge(m, src)
}
func (m *OutgoingTx) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingTx) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTx.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTx proto.InternalMessageInfo

func (m *OutgoingTx) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *OutgoingTx) GetDestAddr() []byte {
	if m != nil {
		return m.DestAddr
	}
	return nil
}

func (m *OutgoingTx) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *OutgoingTx) GetBridgeFee() types.Coin {
	if m != nil {
		return m.BridgeFee
	}
	return types.Coin{}
}

// BridgedDenominator track and identify the ported ERC20 tokens into Peggy.
// An ERC20 token on the Ethereum side can be uniquely identified by the ERC20 contract address and the human readable symbol
// used for it in the contract
// In Peggy this is represented as "vouchers" that get minted and burned when interacting with the Ethereum side. These "vouchers"
// are identified by a prefixed string representation. See VoucherDenom type.
type BridgedDenominator struct {
	TokenContractAddress []byte `protobuf:"bytes,1,opt,name=token_contract_address,json=tokenContractAddress,proto3" json:"token_contract_address,omitempty"`
	Symbol               string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	CosmosVoucherDenom   string `protobuf:"bytes,3,opt,name=cosmos_voucher_denom,json=cosmosVoucherDenom,proto3" json:"cosmos_voucher_denom,omitempty"`
}

func (m *BridgedDenominator) Reset()         { *m = BridgedDenominator{} }
func (m *BridgedDenominator) String() string { return proto.CompactTextString(m) }
func (*BridgedDenominator) ProtoMessage()    {}
func (*BridgedDenominator) Descriptor() ([]byte, []int) {
	return fileDescriptor_42e6de9b297dce86, []int{1}
}
func (m *BridgedDenominator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BridgedDenominator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BridgedDenominator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BridgedDenominator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgedDenominator.Merge(m, src)
}
func (m *BridgedDenominator) XXX_Size() int {
	return m.Size()
}
func (m *BridgedDenominator) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgedDenominator.DiscardUnknown(m)
}

var xxx_messageInfo_BridgedDenominator proto.InternalMessageInfo

func (m *BridgedDenominator) GetTokenContractAddress() []byte {
	if m != nil {
		return m.TokenContractAddress
	}
	return nil
}

func (m *BridgedDenominator) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *BridgedDenominator) GetCosmosVoucherDenom() string {
	if m != nil {
		return m.CosmosVoucherDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*OutgoingTx)(nil), "module.peggy.v1beta1.OutgoingTx")
	proto.RegisterType((*BridgedDenominator)(nil), "module.peggy.v1beta1.BridgedDenominator")
}

func init() { proto.RegisterFile("module/peggy/v1beta1/pool.proto", fileDescriptor_42e6de9b297dce86) }

var fileDescriptor_42e6de9b297dce86 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x31, 0xaf, 0xd3, 0x30,
	0x18, 0x8c, 0xe1, 0xa9, 0x22, 0xe6, 0x4d, 0x56, 0xf4, 0x14, 0x1e, 0x52, 0x5a, 0x75, 0xea, 0x42,
	0xd2, 0x02, 0x12, 0x1b, 0x52, 0x5b, 0xc4, 0xc0, 0x82, 0x14, 0x21, 0x06, 0x96, 0xc8, 0x89, 0x3f,
	0xdc, 0x88, 0xc4, 0x5f, 0x64, 0x3b, 0x55, 0xfb, 0x2f, 0x18, 0xf8, 0x35, 0xfc, 0x82, 0x8e, 0x1d,
	0x99, 0x10, 0x6a, 0xff, 0x08, 0x8a, 0x1d, 0x98, 0xdf, 0xe6, 0xfb, 0xee, 0x4e, 0xbe, 0xd3, 0xd1,
	0x69, 0x8b, 0xa2, 0x6f, 0x20, 0xeb, 0x40, 0xca, 0x63, 0xb6, 0x5f, 0x95, 0x60, 0xf9, 0x2a, 0xeb,
	0x10, 0x9b, 0xb4, 0xd3, 0x68, 0x91, 0x45, 0x5e, 0x90, 0x3a, 0x41, 0x3a, 0x0a, 0xee, 0x93, 0x0a,
	0x4d, 0x8b, 0x26, 0x2b, 0xb9, 0x81, 0xff, 0xae, 0x0a, 0x6b, 0xe5, 0x5d, 0xf7, 0x91, 0x44, 0x89,
	0xee, 0x99, 0x0d, 0x2f, 0x7f, 0x9d, 0xff, 0x24, 0x94, 0x7e, 0xec, 0xad, 0xc4, 0x5a, 0xc9, 0x4f,
	0x07, 0x76, 0x47, 0x27, 0x06, 0x94, 0x00, 0x1d, 0x93, 0x19, 0x59, 0x84, 0xf9, 0x88, 0xd8, 0x73,
	0x1a, 0x0a, 0x30, 0xb6, 0xe0, 0x42, 0xe8, 0xf8, 0xd1, 0x8c, 0x2c, 0x6e, 0xf3, 0x27, 0xc3, 0x61,
	0x2d, 0x84, 0x66, 0x6f, 0xe8, 0x84, 0xb7, 0xd8, 0x2b, 0x1b, 0x3f, 0x9e, 0x91, 0xc5, 0xd3, 0x97,
	0xcf, 0x52, 0x1f, 0x25, 0x1d, 0xa2, 0xfc, 0xcb, 0x97, 0x6e, 0xb1, 0x56, 0x9b, 0x9b, 0xd3, 0xef,
	0x69, 0x90, 0x8f, 0x72, 0xf6, 0x96, 0xd2, 0x52, 0xd7, 0x42, 0x42, 0xf1, 0x15, 0x20, 0xbe, 0x79,
	0x98, 0x39, 0xf4, 0x96, 0xf7, 0x00, 0xf3, 0x1f, 0x84, 0xb2, 0x8d, 0x43, 0xe2, 0x1d, 0x28, 0x6c,
	0x6b, 0xc5, 0x2d, 0x6a, 0xf6, 0x9a, 0xde, 0x59, 0xfc, 0x06, 0xaa, 0xa8, 0x50, 0x59, 0xcd, 0x2b,
	0x1f, 0x1b, 0x8c, 0x71, 0xa5, 0x6e, 0xf3, 0xc8, 0xb1, 0xdb, 0x91, 0x5c, 0x7b, 0xce, 0x55, 0x3f,
	0xb6, 0x25, 0x36, 0xae, 0xdf, 0x50, 0xdd, 0x21, 0xb6, 0xa4, 0x91, 0x4f, 0x54, 0xec, 0xb1, 0xaf,
	0x76, 0xa0, 0x0b, 0x31, 0xfc, 0xe5, 0xba, 0x86, 0x39, 0xf3, 0xdc, 0x67, 0x4f, 0xb9, 0x14, 0x9b,
	0x0f, 0xa7, 0x4b, 0x42, 0xce, 0x97, 0x84, 0xfc, 0xb9, 0x24, 0xe4, 0xfb, 0x35, 0x09, 0xce, 0xd7,
	0x24, 0xf8, 0x75, 0x4d, 0x82, 0x2f, 0x4b, 0x59, 0xdb, 0x5d, 0x5f, 0xa6, 0x15, 0xb6, 0x19, 0x6f,
	0xec, 0x0e, 0xf8, 0x0b, 0x05, 0x76, 0x5c, 0x7a, 0x9c, 0xfd, 0x30, 0x42, 0x7b, 0xec, 0xc0, 0x94,
	0x13, 0x37, 0xd3, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x4a, 0xad, 0x66, 0x15, 0x02,
	0x00, 0x00,
}

func (m *OutgoingTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutgoingTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BridgeFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.DestAddr) > 0 {
		i -= len(m.DestAddr)
		copy(dAtA[i:], m.DestAddr)
		i = encodeVarintPool(dAtA, i, uint64(len(m.DestAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BridgedDenominator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BridgedDenominator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BridgedDenominator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CosmosVoucherDenom) > 0 {
		i -= len(m.CosmosVoucherDenom)
		copy(dAtA[i:], m.CosmosVoucherDenom)
		i = encodeVarintPool(dAtA, i, uint64(len(m.CosmosVoucherDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TokenContractAddress) > 0 {
		i -= len(m.TokenContractAddress)
		copy(dAtA[i:], m.TokenContractAddress)
		i = encodeVarintPool(dAtA, i, uint64(len(m.TokenContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutgoingTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.DestAddr)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.BridgeFee.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func (m *BridgedDenominator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenContractAddress)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.CosmosVoucherDenom)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutgoingTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: OutgoingTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAddr = append(m.DestAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.DestAddr == nil {
				m.DestAddr = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BridgeFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPool
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
func (m *BridgedDenominator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: BridgedDenominator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BridgedDenominator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenContractAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenContractAddress = append(m.TokenContractAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.TokenContractAddress == nil {
				m.TokenContractAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosVoucherDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CosmosVoucherDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)