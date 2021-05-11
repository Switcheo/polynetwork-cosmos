// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: headersync/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// this line is used by starport scaffolding # proto/tx/message
type MsgSyncGenesis struct {
	Syncer        string `protobuf:"bytes,1,opt,name=syncer,proto3" json:"syncer,omitempty"`
	GenesisHeader string `protobuf:"bytes,2,opt,name=genesisHeader,proto3" json:"genesisHeader,omitempty"`
}

func (m *MsgSyncGenesis) Reset()         { *m = MsgSyncGenesis{} }
func (m *MsgSyncGenesis) String() string { return proto.CompactTextString(m) }
func (*MsgSyncGenesis) ProtoMessage()    {}
func (*MsgSyncGenesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_032c830984a3391a, []int{0}
}
func (m *MsgSyncGenesis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSyncGenesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSyncGenesis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSyncGenesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSyncGenesis.Merge(m, src)
}
func (m *MsgSyncGenesis) XXX_Size() int {
	return m.Size()
}
func (m *MsgSyncGenesis) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSyncGenesis.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSyncGenesis proto.InternalMessageInfo

func (m *MsgSyncGenesis) GetSyncer() string {
	if m != nil {
		return m.Syncer
	}
	return ""
}

func (m *MsgSyncGenesis) GetGenesisHeader() string {
	if m != nil {
		return m.GenesisHeader
	}
	return ""
}

type MsgSyncGenesisResponse struct {
}

func (m *MsgSyncGenesisResponse) Reset()         { *m = MsgSyncGenesisResponse{} }
func (m *MsgSyncGenesisResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSyncGenesisResponse) ProtoMessage()    {}
func (*MsgSyncGenesisResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_032c830984a3391a, []int{1}
}
func (m *MsgSyncGenesisResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSyncGenesisResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSyncGenesisResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSyncGenesisResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSyncGenesisResponse.Merge(m, src)
}
func (m *MsgSyncGenesisResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSyncGenesisResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSyncGenesisResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSyncGenesisResponse proto.InternalMessageInfo

type MsgSyncHeaders struct {
	Syncer  string   `protobuf:"bytes,1,opt,name=syncer,proto3" json:"syncer,omitempty"`
	Headers []string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (m *MsgSyncHeaders) Reset()         { *m = MsgSyncHeaders{} }
func (m *MsgSyncHeaders) String() string { return proto.CompactTextString(m) }
func (*MsgSyncHeaders) ProtoMessage()    {}
func (*MsgSyncHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_032c830984a3391a, []int{2}
}
func (m *MsgSyncHeaders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSyncHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSyncHeaders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSyncHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSyncHeaders.Merge(m, src)
}
func (m *MsgSyncHeaders) XXX_Size() int {
	return m.Size()
}
func (m *MsgSyncHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSyncHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSyncHeaders proto.InternalMessageInfo

func (m *MsgSyncHeaders) GetSyncer() string {
	if m != nil {
		return m.Syncer
	}
	return ""
}

func (m *MsgSyncHeaders) GetHeaders() []string {
	if m != nil {
		return m.Headers
	}
	return nil
}

type MsgSyncHeadersResponse struct {
}

func (m *MsgSyncHeadersResponse) Reset()         { *m = MsgSyncHeadersResponse{} }
func (m *MsgSyncHeadersResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSyncHeadersResponse) ProtoMessage()    {}
func (*MsgSyncHeadersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_032c830984a3391a, []int{3}
}
func (m *MsgSyncHeadersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSyncHeadersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSyncHeadersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSyncHeadersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSyncHeadersResponse.Merge(m, src)
}
func (m *MsgSyncHeadersResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSyncHeadersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSyncHeadersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSyncHeadersResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSyncGenesis)(nil), "Switcheo.polynetworkcosmos.headersync.MsgSyncGenesis")
	proto.RegisterType((*MsgSyncGenesisResponse)(nil), "Switcheo.polynetworkcosmos.headersync.MsgSyncGenesisResponse")
	proto.RegisterType((*MsgSyncHeaders)(nil), "Switcheo.polynetworkcosmos.headersync.MsgSyncHeaders")
	proto.RegisterType((*MsgSyncHeadersResponse)(nil), "Switcheo.polynetworkcosmos.headersync.MsgSyncHeadersResponse")
}

func init() { proto.RegisterFile("headersync/tx.proto", fileDescriptor_032c830984a3391a) }

var fileDescriptor_032c830984a3391a = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x48, 0x4d, 0x4c,
	0x49, 0x2d, 0x2a, 0xae, 0xcc, 0x4b, 0xd6, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x0d, 0x2e, 0xcf, 0x2c, 0x49, 0xce, 0x48, 0xcd, 0xd7, 0x2b, 0xc8, 0xcf, 0xa9, 0xcc, 0x4b,
	0x2d, 0x29, 0xcf, 0x2f, 0xca, 0x4e, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0x43, 0xa8, 0x57, 0xf2,
	0xe3, 0xe2, 0xf3, 0x2d, 0x4e, 0x0f, 0xae, 0xcc, 0x4b, 0x76, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c,
	0x16, 0x12, 0xe3, 0x62, 0x03, 0xc9, 0xa4, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41,
	0x79, 0x42, 0x2a, 0x5c, 0xbc, 0xe9, 0x10, 0x25, 0x1e, 0x60, 0xed, 0x12, 0x4c, 0x60, 0x69, 0x54,
	0x41, 0x25, 0x09, 0x2e, 0x31, 0x54, 0xf3, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95,
	0x9c, 0xe0, 0x36, 0x41, 0x94, 0xe2, 0xb6, 0x49, 0x82, 0x8b, 0x1d, 0xea, 0x42, 0x09, 0x26, 0x05,
	0x66, 0x0d, 0xce, 0x20, 0x18, 0x17, 0xc9, 0x74, 0xa8, 0x19, 0x30, 0xd3, 0x8d, 0x26, 0x32, 0x71,
	0x31, 0xfb, 0x16, 0xa7, 0x0b, 0x35, 0x33, 0x72, 0x71, 0x23, 0xfb, 0xc6, 0x54, 0x8f, 0xa8, 0x70,
	0xd0, 0x43, 0x75, 0xb4, 0x94, 0x2d, 0x59, 0xda, 0x60, 0xae, 0x81, 0xbb, 0x02, 0xe6, 0x53, 0x12,
	0x5d, 0x01, 0xd5, 0x46, 0xaa, 0x2b, 0xd0, 0xc2, 0xc4, 0x29, 0xf8, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x2c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73,
	0xf5, 0x61, 0x56, 0xe8, 0x23, 0x59, 0xa1, 0x0b, 0xb1, 0x43, 0xbf, 0x42, 0x1f, 0x39, 0x69, 0x55,
	0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x93, 0x97, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xae,
	0xa7, 0x5d, 0x75, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	SyncGenesis(ctx context.Context, in *MsgSyncGenesis, opts ...grpc.CallOption) (*MsgSyncGenesisResponse, error)
	SyncHeaders(ctx context.Context, in *MsgSyncHeaders, opts ...grpc.CallOption) (*MsgSyncHeadersResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SyncGenesis(ctx context.Context, in *MsgSyncGenesis, opts ...grpc.CallOption) (*MsgSyncGenesisResponse, error) {
	out := new(MsgSyncGenesisResponse)
	err := c.cc.Invoke(ctx, "/Switcheo.polynetworkcosmos.headersync.Msg/SyncGenesis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SyncHeaders(ctx context.Context, in *MsgSyncHeaders, opts ...grpc.CallOption) (*MsgSyncHeadersResponse, error) {
	out := new(MsgSyncHeadersResponse)
	err := c.cc.Invoke(ctx, "/Switcheo.polynetworkcosmos.headersync.Msg/SyncHeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	SyncGenesis(context.Context, *MsgSyncGenesis) (*MsgSyncGenesisResponse, error)
	SyncHeaders(context.Context, *MsgSyncHeaders) (*MsgSyncHeadersResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SyncGenesis(ctx context.Context, req *MsgSyncGenesis) (*MsgSyncGenesisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncGenesis not implemented")
}
func (*UnimplementedMsgServer) SyncHeaders(ctx context.Context, req *MsgSyncHeaders) (*MsgSyncHeadersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncHeaders not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SyncGenesis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSyncGenesis)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SyncGenesis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Switcheo.polynetworkcosmos.headersync.Msg/SyncGenesis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SyncGenesis(ctx, req.(*MsgSyncGenesis))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SyncHeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSyncHeaders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SyncHeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Switcheo.polynetworkcosmos.headersync.Msg/SyncHeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SyncHeaders(ctx, req.(*MsgSyncHeaders))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Switcheo.polynetworkcosmos.headersync.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncGenesis",
			Handler:    _Msg_SyncGenesis_Handler,
		},
		{
			MethodName: "SyncHeaders",
			Handler:    _Msg_SyncHeaders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "headersync/tx.proto",
}

func (m *MsgSyncGenesis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSyncGenesis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSyncGenesis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GenesisHeader) > 0 {
		i -= len(m.GenesisHeader)
		copy(dAtA[i:], m.GenesisHeader)
		i = encodeVarintTx(dAtA, i, uint64(len(m.GenesisHeader)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Syncer) > 0 {
		i -= len(m.Syncer)
		copy(dAtA[i:], m.Syncer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Syncer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSyncGenesisResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSyncGenesisResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSyncGenesisResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSyncHeaders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSyncHeaders) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSyncHeaders) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Headers[iNdEx])
			copy(dAtA[i:], m.Headers[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Headers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Syncer) > 0 {
		i -= len(m.Syncer)
		copy(dAtA[i:], m.Syncer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Syncer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSyncHeadersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSyncHeadersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSyncHeadersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSyncGenesis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Syncer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.GenesisHeader)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSyncGenesisResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSyncHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Syncer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Headers) > 0 {
		for _, s := range m.Headers {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgSyncHeadersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSyncGenesis) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSyncGenesis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSyncGenesis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Syncer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Syncer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisHeader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisHeader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSyncGenesisResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSyncGenesisResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSyncGenesisResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSyncHeaders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSyncHeaders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSyncHeaders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Syncer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Syncer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSyncHeadersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSyncHeadersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSyncHeadersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)