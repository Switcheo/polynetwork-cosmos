// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lockproxy/event.proto

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

type LockEvent struct {
	Denom         string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	FromLockProxy string `protobuf:"bytes,2,opt,name=from_lock_proxy,json=fromLockProxy,proto3" json:"from_lock_proxy,omitempty" yaml:"from_lock_proxy"`
	FromAssetId   string `protobuf:"bytes,3,opt,name=from_asset_id,json=fromAssetId,proto3" json:"from_asset_id,omitempty" yaml:"from_asset_id"`
	FromAddress   string `protobuf:"bytes,4,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	ToChainId     string `protobuf:"bytes,5,opt,name=to_chain_id,json=toChainId,proto3" json:"to_chain_id,omitempty" yaml:"to_chain_id"`
	ToLockProxy   string `protobuf:"bytes,6,opt,name=to_lock_proxy,json=toLockProxy,proto3" json:"to_lock_proxy,omitempty" yaml:"to_lock_proxy"`
	ToAssetId     string `protobuf:"bytes,7,opt,name=to_asset_id,json=toAssetId,proto3" json:"to_asset_id,omitempty" yaml:"to_asset_id"`
	ToAddress     string `protobuf:"bytes,8,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	Amount        string `protobuf:"bytes,9,opt,name=amount,proto3" json:"amount,omitempty"`
	FeeAmount     string `protobuf:"bytes,10,opt,name=fee_amount,json=feeAmount,proto3" json:"fee_amount,omitempty" yaml:"fee_amount"`
	FeeAddress    string `protobuf:"bytes,11,opt,name=fee_address,json=feeAddress,proto3" json:"fee_address,omitempty" yaml:"fee_address"`
	Nonce         string `protobuf:"bytes,12,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *LockEvent) Reset()         { *m = LockEvent{} }
func (m *LockEvent) String() string { return proto.CompactTextString(m) }
func (*LockEvent) ProtoMessage()    {}
func (*LockEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_56c1279822ec82cd, []int{0}
}
func (m *LockEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LockEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LockEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LockEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockEvent.Merge(m, src)
}
func (m *LockEvent) XXX_Size() int {
	return m.Size()
}
func (m *LockEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LockEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LockEvent proto.InternalMessageInfo

func (m *LockEvent) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *LockEvent) GetFromLockProxy() string {
	if m != nil {
		return m.FromLockProxy
	}
	return ""
}

func (m *LockEvent) GetFromAssetId() string {
	if m != nil {
		return m.FromAssetId
	}
	return ""
}

func (m *LockEvent) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *LockEvent) GetToChainId() string {
	if m != nil {
		return m.ToChainId
	}
	return ""
}

func (m *LockEvent) GetToLockProxy() string {
	if m != nil {
		return m.ToLockProxy
	}
	return ""
}

func (m *LockEvent) GetToAssetId() string {
	if m != nil {
		return m.ToAssetId
	}
	return ""
}

func (m *LockEvent) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *LockEvent) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *LockEvent) GetFeeAmount() string {
	if m != nil {
		return m.FeeAmount
	}
	return ""
}

func (m *LockEvent) GetFeeAddress() string {
	if m != nil {
		return m.FeeAddress
	}
	return ""
}

func (m *LockEvent) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

type UnlockEvent struct {
	ToAssetId   string `protobuf:"bytes,1,opt,name=to_asset_id,json=toAssetId,proto3" json:"to_asset_id,omitempty" yaml:"to_asset_id"`
	ToAddress   string `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	Amount      string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	FromAddress string `protobuf:"bytes,4,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	FromAssetId string `protobuf:"bytes,5,opt,name=from_asset_id,json=fromAssetId,proto3" json:"from_asset_id,omitempty" yaml:"from_asset_id"`
	FeeAmount   string `protobuf:"bytes,6,opt,name=fee_amount,json=feeAmount,proto3" json:"fee_amount,omitempty" yaml:"fee_amount"`
	FeeAddress  string `protobuf:"bytes,7,opt,name=fee_address,json=feeAddress,proto3" json:"fee_address,omitempty" yaml:"fee_address"`
	Nonce       string `protobuf:"bytes,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *UnlockEvent) Reset()         { *m = UnlockEvent{} }
func (m *UnlockEvent) String() string { return proto.CompactTextString(m) }
func (*UnlockEvent) ProtoMessage()    {}
func (*UnlockEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_56c1279822ec82cd, []int{1}
}
func (m *UnlockEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnlockEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnlockEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnlockEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnlockEvent.Merge(m, src)
}
func (m *UnlockEvent) XXX_Size() int {
	return m.Size()
}
func (m *UnlockEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_UnlockEvent.DiscardUnknown(m)
}

var xxx_messageInfo_UnlockEvent proto.InternalMessageInfo

func (m *UnlockEvent) GetToAssetId() string {
	if m != nil {
		return m.ToAssetId
	}
	return ""
}

func (m *UnlockEvent) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *UnlockEvent) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *UnlockEvent) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *UnlockEvent) GetFromAssetId() string {
	if m != nil {
		return m.FromAssetId
	}
	return ""
}

func (m *UnlockEvent) GetFeeAmount() string {
	if m != nil {
		return m.FeeAmount
	}
	return ""
}

func (m *UnlockEvent) GetFeeAddress() string {
	if m != nil {
		return m.FeeAddress
	}
	return ""
}

func (m *UnlockEvent) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func init() {
	proto.RegisterType((*LockEvent)(nil), "Switcheo.carbon.lockproxy.LockEvent")
	proto.RegisterType((*UnlockEvent)(nil), "Switcheo.carbon.lockproxy.UnlockEvent")
}

func init() { proto.RegisterFile("lockproxy/event.proto", fileDescriptor_56c1279822ec82cd) }

var fileDescriptor_56c1279822ec82cd = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x9b, 0x26, 0xad, 0xd7, 0x45, 0x08, 0x93, 0x06, 0xd3, 0x83, 0x8d, 0xf6, 0xc4, 0x85,
	0xf8, 0x00, 0x02, 0x54, 0x71, 0x69, 0x10, 0x87, 0x4a, 0x1c, 0x90, 0x11, 0x17, 0x2e, 0x96, 0x63,
	0x6f, 0x3e, 0x94, 0x78, 0x27, 0xb2, 0xb7, 0xb4, 0xf9, 0x17, 0x88, 0x5f, 0xc5, 0xb1, 0x47, 0x4e,
	0x56, 0x95, 0xfc, 0x03, 0xff, 0x02, 0x34, 0xbb, 0x76, 0xec, 0x18, 0x2e, 0x69, 0x6f, 0x7e, 0x7e,
	0xfb, 0xde, 0xce, 0x9b, 0x19, 0x2d, 0x39, 0x5d, 0x40, 0x38, 0x5f, 0x26, 0x70, 0xb3, 0x72, 0xd9,
	0x0f, 0xc6, 0xc5, 0x60, 0x99, 0x80, 0x00, 0xf3, 0xf9, 0xd7, 0xeb, 0x99, 0x08, 0xa7, 0x0c, 0x06,
	0x61, 0x90, 0x8c, 0x80, 0x0f, 0xb6, 0xc7, 0xce, 0x7a, 0x13, 0x98, 0x80, 0x3c, 0xe5, 0xe2, 0x97,
	0x12, 0xd0, 0xbb, 0x43, 0xa2, 0x7f, 0x86, 0x70, 0xfe, 0x09, 0x4d, 0xcc, 0x1e, 0xe9, 0x44, 0x8c,
	0x43, 0x6c, 0x69, 0x2f, 0xb4, 0x97, 0xba, 0xa7, 0x80, 0x39, 0x24, 0x8f, 0xc7, 0x09, 0xc4, 0x3e,
	0x7a, 0xf9, 0xd2, 0xcc, 0x3a, 0x40, 0x7e, 0x78, 0x96, 0x67, 0x4e, 0x7f, 0x15, 0xc4, 0x8b, 0x73,
	0xda, 0x38, 0x40, 0xbd, 0x47, 0xf8, 0x07, 0x9d, 0xbf, 0x20, 0x36, 0x3f, 0x10, 0xf9, 0xc3, 0x0f,
	0xd2, 0x94, 0x09, 0x7f, 0x16, 0x59, 0x6d, 0xe9, 0x60, 0xe5, 0x99, 0xd3, 0xab, 0x39, 0x94, 0x34,
	0xf5, 0x0c, 0xc4, 0x17, 0x08, 0x2f, 0x23, 0xf3, 0x9c, 0x9c, 0x28, 0x3a, 0x8a, 0x12, 0x96, 0xa6,
	0xd6, 0xa1, 0x14, 0x3f, 0xcb, 0x33, 0xe7, 0x69, 0x5d, 0xac, 0xd8, 0x52, 0xab, 0x90, 0xf9, 0x96,
	0x18, 0x02, 0xfc, 0x70, 0x1a, 0xcc, 0x38, 0xde, 0xdb, 0x91, 0xd2, 0x7e, 0x9e, 0x39, 0xa6, 0x92,
	0xd6, 0x48, 0xea, 0xe9, 0x02, 0x3e, 0x22, 0xb8, 0x8c, 0xb0, 0x62, 0x01, 0xf5, 0xcc, 0xdd, 0x66,
	0xc5, 0x3b, 0x34, 0xf5, 0x0c, 0x01, 0x55, 0x5e, 0x75, 0xeb, 0x36, 0xed, 0xd1, 0x7f, 0x6e, 0xad,
	0xb2, 0xea, 0x02, 0xca, 0xa4, 0x6f, 0x08, 0x41, 0xaa, 0xc8, 0x79, 0x2c, 0x65, 0xa7, 0x79, 0xe6,
	0x3c, 0xa9, 0x64, 0x65, 0x4a, 0x54, 0x15, 0x19, 0xfb, 0xa4, 0x1b, 0xc4, 0x70, 0xc5, 0x85, 0xa5,
	0xcb, 0xc1, 0x15, 0x08, 0xdd, 0xc6, 0x8c, 0xf9, 0x05, 0x47, 0x9a, 0x6e, 0x15, 0x47, 0x3d, 0x7d,
	0xcc, 0xd8, 0x85, 0x52, 0xbd, 0x23, 0x86, 0x64, 0x8a, 0x22, 0x8c, 0x66, 0xed, 0x35, 0x92, 0x7a,
	0x78, 0x41, 0x59, 0x46, 0x8f, 0x74, 0x38, 0xf0, 0x90, 0x59, 0x27, 0x6a, 0x7d, 0x24, 0xa0, 0xbf,
	0xda, 0xc4, 0xf8, 0xc6, 0x17, 0xdb, 0x25, 0x6b, 0xb4, 0x46, 0xbb, 0x5f, 0x6b, 0x0e, 0xf6, 0x6e,
	0x4d, 0x7b, 0xa7, 0x35, 0x0f, 0x59, 0xa9, 0x7f, 0x96, 0xb9, 0xb3, 0xcf, 0x32, 0xef, 0x0e, 0xa5,
	0x7b, 0xbf, 0xa1, 0x1c, 0xed, 0x3f, 0x94, 0xe3, 0xda, 0x50, 0x86, 0xde, 0xef, 0xb5, 0xad, 0xdd,
	0xae, 0x6d, 0xed, 0x6e, 0x6d, 0x6b, 0x3f, 0x37, 0x76, 0xeb, 0x76, 0x63, 0xb7, 0xfe, 0x6c, 0xec,
	0xd6, 0xf7, 0xf7, 0x93, 0x99, 0x98, 0x5e, 0x8d, 0x06, 0x21, 0xc4, 0x6e, 0xf9, 0x9a, 0xb8, 0x4b,
	0x58, 0xac, 0x38, 0x13, 0xd7, 0x90, 0xcc, 0x5f, 0x85, 0x90, 0xc6, 0x90, 0xba, 0x37, 0x6e, 0xf5,
	0x04, 0x89, 0xd5, 0x92, 0xa5, 0xa3, 0xae, 0x7c, 0x52, 0x5e, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x3c, 0xf5, 0x40, 0xd9, 0x9c, 0x04, 0x00, 0x00,
}

func (m *LockEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LockEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LockEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.FeeAddress) > 0 {
		i -= len(m.FeeAddress)
		copy(dAtA[i:], m.FeeAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.FeeAddress)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.FeeAmount) > 0 {
		i -= len(m.FeeAmount)
		copy(dAtA[i:], m.FeeAmount)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.FeeAmount)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ToAssetId) > 0 {
		i -= len(m.ToAssetId)
		copy(dAtA[i:], m.ToAssetId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ToAssetId)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ToLockProxy) > 0 {
		i -= len(m.ToLockProxy)
		copy(dAtA[i:], m.ToLockProxy)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ToLockProxy)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ToChainId) > 0 {
		i -= len(m.ToChainId)
		copy(dAtA[i:], m.ToChainId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ToChainId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FromAssetId) > 0 {
		i -= len(m.FromAssetId)
		copy(dAtA[i:], m.FromAssetId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.FromAssetId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FromLockProxy) > 0 {
		i -= len(m.FromLockProxy)
		copy(dAtA[i:], m.FromLockProxy)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.FromLockProxy)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UnlockEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnlockEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnlockEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.FeeAddress) > 0 {
		i -= len(m.FeeAddress)
		copy(dAtA[i:], m.FeeAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.FeeAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.FeeAmount) > 0 {
		i -= len(m.FeeAmount)
		copy(dAtA[i:], m.FeeAmount)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.FeeAmount)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.FromAssetId) > 0 {
		i -= len(m.FromAssetId)
		copy(dAtA[i:], m.FromAssetId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.FromAssetId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ToAssetId) > 0 {
		i -= len(m.ToAssetId)
		copy(dAtA[i:], m.ToAssetId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ToAssetId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LockEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.FromLockProxy)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.FromAssetId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.ToChainId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.ToLockProxy)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.ToAssetId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.FeeAmount)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.FeeAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *UnlockEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ToAssetId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.FromAssetId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.FeeAmount)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.FeeAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LockEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: LockEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LockEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromLockProxy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromLockProxy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToLockProxy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToLockProxy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *UnlockEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: UnlockEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnlockEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
