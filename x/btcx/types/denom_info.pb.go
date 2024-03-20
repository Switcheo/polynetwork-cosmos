// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Switcheo/carbon/btcx/denom_info.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

type DenomInfo struct {
	Creator          string                `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id               string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Denom            string                `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	AssetHash        string                `protobuf:"bytes,4,opt,name=asset_hash,json=assetHash,proto3" json:"asset_hash,omitempty"`
	TotalSupply      cosmossdk_io_math.Int `protobuf:"bytes,5,opt,name=total_supply,json=totalSupply,proto3,customtype=cosmossdk.io/math.Int" json:"total_supply"`
	RedeemScript     string                `protobuf:"bytes,6,opt,name=redeem_script,json=redeemScript,proto3" json:"redeem_script,omitempty"`
	RedeemScriptHash string                `protobuf:"bytes,7,opt,name=redeem_script_hash,json=redeemScriptHash,proto3" json:"redeem_script_hash,omitempty"`
}

func (m *DenomInfo) Reset()         { *m = DenomInfo{} }
func (m *DenomInfo) String() string { return proto.CompactTextString(m) }
func (*DenomInfo) ProtoMessage()    {}
func (*DenomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_db2efd66eb0c229c, []int{0}
}
func (m *DenomInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomInfo.Merge(m, src)
}
func (m *DenomInfo) XXX_Size() int {
	return m.Size()
}
func (m *DenomInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DenomInfo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DenomInfo)(nil), "Switcheo.carbon.btcx.DenomInfo")
}

func init() {
	proto.RegisterFile("Switcheo/carbon/btcx/denom_info.proto", fileDescriptor_db2efd66eb0c229c)
}

var fileDescriptor_db2efd66eb0c229c = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x93, 0xd8, 0x3f, 0x74, 0xad, 0x22, 0x4b, 0x85, 0x50, 0xe8, 0x56, 0x14, 0xd1, 0x83,
	0x64, 0x29, 0xde, 0x3c, 0x69, 0xf1, 0x60, 0xaf, 0xed, 0xcd, 0x4b, 0xd8, 0x6c, 0xb6, 0xcd, 0x62,
	0x93, 0x09, 0xd9, 0xad, 0xb6, 0x0f, 0x20, 0x78, 0xf4, 0x11, 0x7c, 0x9c, 0x1e, 0x7b, 0x14, 0x0f,
	0x45, 0xda, 0x17, 0x91, 0xec, 0x52, 0x50, 0x6f, 0x3b, 0xdf, 0x7c, 0xc3, 0xfc, 0xd8, 0x41, 0xe7,
	0xa3, 0x17, 0xa9, 0x79, 0x22, 0x80, 0x72, 0x56, 0x44, 0x90, 0xd1, 0x48, 0xf3, 0x39, 0x8d, 0x45,
	0x06, 0x69, 0x28, 0xb3, 0x31, 0x04, 0x79, 0x01, 0x1a, 0x70, 0x6b, 0xa7, 0x05, 0x56, 0x0b, 0x4a,
	0xad, 0xdd, 0x9a, 0xc0, 0x04, 0x8c, 0x40, 0xcb, 0x97, 0x75, 0xdb, 0x84, 0x83, 0x4a, 0x41, 0xd1,
	0x88, 0x29, 0x41, 0x9f, 0x7b, 0x91, 0xd0, 0xac, 0x47, 0x39, 0xc8, 0xcc, 0xf6, 0x4f, 0x5f, 0x3d,
	0xd4, 0xb8, 0x2f, 0x17, 0x0c, 0xb2, 0x31, 0x60, 0x1f, 0xd5, 0x79, 0x21, 0x98, 0x86, 0xc2, 0x77,
	0x4f, 0xdc, 0xcb, 0xc6, 0x70, 0x57, 0xe2, 0x43, 0xe4, 0xc9, 0xd8, 0xf7, 0x0c, 0xf4, 0x64, 0x8c,
	0x5b, 0xa8, 0x6a, 0x72, 0xf9, 0x7b, 0x06, 0xd9, 0x02, 0x77, 0x10, 0x62, 0x4a, 0x09, 0x1d, 0x26,
	0x4c, 0x25, 0x7e, 0xc5, 0xb4, 0x1a, 0x86, 0x3c, 0x30, 0x95, 0xe0, 0x5b, 0xd4, 0xd4, 0xa0, 0xd9,
	0x34, 0x54, 0xb3, 0x3c, 0x9f, 0x2e, 0xfc, 0x6a, 0x29, 0xf4, 0x3b, 0xcb, 0x75, 0xd7, 0xf9, 0x5a,
	0x77, 0x8f, 0x6d, 0x54, 0x15, 0x3f, 0x05, 0x12, 0x68, 0xca, 0x74, 0x12, 0x0c, 0x32, 0x3d, 0xdc,
	0x37, 0x23, 0x23, 0x33, 0x81, 0xcf, 0xd0, 0x41, 0x21, 0x62, 0x21, 0xd2, 0x50, 0xf1, 0x42, 0xe6,
	0xda, 0xaf, 0x99, 0x1d, 0x4d, 0x0b, 0x47, 0x86, 0xe1, 0x2b, 0x84, 0xff, 0x48, 0x36, 0x4d, 0xdd,
	0x98, 0x47, 0xbf, 0xcd, 0x32, 0xd4, 0x4d, 0xe5, 0xed, 0xa3, 0xeb, 0xf4, 0xef, 0x96, 0x1b, 0xe2,
	0xae, 0x36, 0xc4, 0xfd, 0xde, 0x10, 0xf7, 0x7d, 0x4b, 0x9c, 0xd5, 0x96, 0x38, 0x9f, 0x5b, 0xe2,
	0x3c, 0x5e, 0x4c, 0xa4, 0x4e, 0x66, 0x51, 0xc0, 0x21, 0xa5, 0xff, 0xef, 0x33, 0xb7, 0x17, 0xd2,
	0x8b, 0x5c, 0xa8, 0xa8, 0x66, 0x7e, 0xf4, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x76, 0xa8,
	0x2c, 0xc6, 0x01, 0x00, 0x00,
}

func (m *DenomInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RedeemScriptHash) > 0 {
		i -= len(m.RedeemScriptHash)
		copy(dAtA[i:], m.RedeemScriptHash)
		i = encodeVarintDenomInfo(dAtA, i, uint64(len(m.RedeemScriptHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.RedeemScript) > 0 {
		i -= len(m.RedeemScript)
		copy(dAtA[i:], m.RedeemScript)
		i = encodeVarintDenomInfo(dAtA, i, uint64(len(m.RedeemScript)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.TotalSupply.Size()
		i -= size
		if _, err := m.TotalSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDenomInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.AssetHash) > 0 {
		i -= len(m.AssetHash)
		copy(dAtA[i:], m.AssetHash)
		i = encodeVarintDenomInfo(dAtA, i, uint64(len(m.AssetHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintDenomInfo(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDenomInfo(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDenomInfo(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDenomInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovDenomInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DenomInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDenomInfo(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDenomInfo(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovDenomInfo(uint64(l))
	}
	l = len(m.AssetHash)
	if l > 0 {
		n += 1 + l + sovDenomInfo(uint64(l))
	}
	l = m.TotalSupply.Size()
	n += 1 + l + sovDenomInfo(uint64(l))
	l = len(m.RedeemScript)
	if l > 0 {
		n += 1 + l + sovDenomInfo(uint64(l))
	}
	l = len(m.RedeemScriptHash)
	if l > 0 {
		n += 1 + l + sovDenomInfo(uint64(l))
	}
	return n
}

func sovDenomInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDenomInfo(x uint64) (n int) {
	return sovDenomInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DenomInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDenomInfo
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
			return fmt.Errorf("proto: DenomInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomInfo
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
				return ErrInvalidLengthDenomInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomInfo
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
				return ErrInvalidLengthDenomInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomInfo
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
				return ErrInvalidLengthDenomInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomInfo
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
				return ErrInvalidLengthDenomInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomInfo
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
				return ErrInvalidLengthDenomInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedeemScript", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomInfo
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
				return ErrInvalidLengthDenomInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedeemScript = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedeemScriptHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomInfo
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
				return ErrInvalidLengthDenomInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedeemScriptHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDenomInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDenomInfo
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
func skipDenomInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDenomInfo
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
					return 0, ErrIntOverflowDenomInfo
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
					return 0, ErrIntOverflowDenomInfo
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
				return 0, ErrInvalidLengthDenomInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDenomInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDenomInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDenomInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDenomInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDenomInfo = fmt.Errorf("proto: unexpected end of group")
)
