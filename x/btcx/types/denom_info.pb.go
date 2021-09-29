// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: btcx/denom_info.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type DenomInfo struct {
	Creator          string                                 `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id               string                                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Denom            string                                 `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	AssetHash        string                                 `protobuf:"bytes,4,opt,name=asset_hash,json=assetHash,proto3" json:"asset_hash,omitempty" yaml:"asset_hash"`
	TotalSupply      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=total_supply,json=totalSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_supply" yaml:"total_supply"`
	RedeemScript     string                                 `protobuf:"bytes,6,opt,name=redeem_script,json=redeemScript,proto3" json:"redeem_script,omitempty" yaml:"redeem_script"`
	RedeemScriptHash string                                 `protobuf:"bytes,7,opt,name=redeem_script_hash,json=redeemScriptHash,proto3" json:"redeem_script_hash,omitempty" yaml:"redeem_script_hash"`
}

func (m *DenomInfo) Reset()         { *m = DenomInfo{} }
func (m *DenomInfo) String() string { return proto.CompactTextString(m) }
func (*DenomInfo) ProtoMessage()    {}
func (*DenomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c32e8847bb2457b, []int{0}
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
	proto.RegisterType((*DenomInfo)(nil), "Switcheo.polynetworkcosmos.btcx.DenomInfo")
}

func init() { proto.RegisterFile("btcx/denom_info.proto", fileDescriptor_1c32e8847bb2457b) }

var fileDescriptor_1c32e8847bb2457b = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x4d, 0xe2, 0x7e, 0x30, 0xed, 0x2a, 0xda, 0xce, 0x42, 0xbb, 0x60, 0x5a, 0x72, 0x10, 0x2f,
	0x9b, 0x66, 0x59, 0x4f, 0x0b, 0x5e, 0x16, 0x05, 0x17, 0xf1, 0x92, 0xb9, 0x79, 0x09, 0x9d, 0xa4,
	0x77, 0x12, 0x26, 0x49, 0x85, 0x74, 0x8f, 0x33, 0xf9, 0x07, 0x1e, 0xfd, 0x09, 0xfe, 0x9c, 0x39,
	0xce, 0x51, 0x3c, 0x04, 0x99, 0xf9, 0x07, 0xb9, 0x0b, 0x92, 0xee, 0x19, 0xcc, 0xc0, 0x9e, 0x92,
	0x57, 0xaf, 0xde, 0xab, 0xbc, 0xaa, 0xa0, 0xf3, 0x48, 0xc5, 0x4b, 0x96, 0x88, 0x12, 0x8a, 0x30,
	0x2b, 0xef, 0xc1, 0xaf, 0x6a, 0x50, 0x80, 0xe9, 0x64, 0x91, 0xa9, 0x38, 0x15, 0xe0, 0x57, 0x90,
	0x37, 0xa5, 0x50, 0x0b, 0xa8, 0x67, 0x31, 0xc8, 0x02, 0xa4, 0xdf, 0x2b, 0x2e, 0xc6, 0x53, 0x98,
	0x82, 0xee, 0x65, 0xfd, 0x9b, 0x91, 0x5d, 0xb8, 0xa6, 0x85, 0x45, 0x5c, 0x0a, 0xf6, 0xed, 0x2a,
	0x12, 0x8a, 0x5f, 0xb1, 0x18, 0xb2, 0xd2, 0xf0, 0xde, 0x5f, 0x07, 0x8d, 0x3e, 0xf4, 0xb3, 0xee,
	0xca, 0x7b, 0xc0, 0x04, 0x9d, 0xc6, 0xb5, 0xe0, 0x0a, 0x6a, 0x62, 0xbf, 0xb6, 0xdf, 0x8e, 0x82,
	0x3d, 0xc4, 0x4f, 0x91, 0x93, 0x25, 0xc4, 0xd1, 0x45, 0x27, 0x4b, 0xf0, 0x18, 0x1d, 0xeb, 0x4f,
	0x24, 0x8f, 0x74, 0xc9, 0x00, 0xfc, 0x0e, 0x21, 0x2e, 0xa5, 0x50, 0x61, 0xca, 0x65, 0x4a, 0x8e,
	0x7a, 0xea, 0xf6, 0xbc, 0x6b, 0xe9, 0xf3, 0x86, 0x17, 0xf9, 0x8d, 0xf7, 0x9f, 0xf3, 0x82, 0x91,
	0x06, 0x9f, 0xb8, 0x4c, 0x71, 0x8a, 0xce, 0x14, 0x28, 0x9e, 0x87, 0x72, 0x5e, 0x55, 0x79, 0x43,
	0x8e, 0xb5, 0xee, 0xe3, 0xaa, 0xa5, 0xd6, 0xef, 0x96, 0xbe, 0x99, 0x66, 0x2a, 0x9d, 0x47, 0x7e,
	0x0c, 0x05, 0xdb, 0x85, 0x31, 0x8f, 0x4b, 0x99, 0xcc, 0x98, 0x6a, 0x2a, 0x21, 0xfd, 0xbb, 0x52,
	0x75, 0x2d, 0x7d, 0x61, 0xa6, 0x0c, 0xbd, 0xbc, 0xe0, 0xb1, 0x86, 0x13, 0x8d, 0xf0, 0x7b, 0xf4,
	0xa4, 0x16, 0x89, 0x10, 0x45, 0x28, 0xe3, 0x3a, 0xab, 0x14, 0x39, 0xd1, 0xa3, 0x48, 0xd7, 0xd2,
	0xb1, 0x11, 0x1f, 0xd0, 0x5e, 0x70, 0x66, 0xf0, 0x44, 0x43, 0xfc, 0x19, 0xe1, 0x03, 0xde, 0xc4,
	0x3c, 0xd5, 0x1e, 0xaf, 0xba, 0x96, 0xbe, 0x7c, 0xc0, 0x63, 0x17, 0xf7, 0xd9, 0xd0, 0xa8, 0x4f,
	0x7d, 0x73, 0xf4, 0xfd, 0x27, 0xb5, 0x6e, 0xbf, 0xac, 0x36, 0xae, 0xbd, 0xde, 0xb8, 0xf6, 0x9f,
	0x8d, 0x6b, 0xff, 0xd8, 0xba, 0xd6, 0x7a, 0xeb, 0x5a, 0xbf, 0xb6, 0xae, 0xf5, 0xf5, 0x7a, 0x90,
	0x7b, 0x7f, 0x7b, 0x36, 0xb8, 0xfd, 0xe5, 0x6e, 0x19, 0x4b, 0xa6, 0x7f, 0x18, 0xbd, 0x88, 0xe8,
	0x44, 0x5f, 0xf5, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0xcf, 0x55, 0x5e, 0x45, 0x02,
	0x00, 0x00,
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