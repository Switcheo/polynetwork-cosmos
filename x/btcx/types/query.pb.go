// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmospoly/btcx/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// this line is used by starport scaffolding # 3
type QueryGetDenomInfoRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryGetDenomInfoRequest) Reset()         { *m = QueryGetDenomInfoRequest{} }
func (m *QueryGetDenomInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetDenomInfoRequest) ProtoMessage()    {}
func (*QueryGetDenomInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8421dffc0f307fe9, []int{0}
}
func (m *QueryGetDenomInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDenomInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDenomInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDenomInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDenomInfoRequest.Merge(m, src)
}
func (m *QueryGetDenomInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDenomInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDenomInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDenomInfoRequest proto.InternalMessageInfo

type QueryGetDenomInfoResponse struct {
	DenomInfo DenomInfo `protobuf:"bytes,1,opt,name=denom_info,json=denomInfo,proto3" json:"denom_info"`
}

func (m *QueryGetDenomInfoResponse) Reset()         { *m = QueryGetDenomInfoResponse{} }
func (m *QueryGetDenomInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetDenomInfoResponse) ProtoMessage()    {}
func (*QueryGetDenomInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8421dffc0f307fe9, []int{1}
}
func (m *QueryGetDenomInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDenomInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDenomInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDenomInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDenomInfoResponse.Merge(m, src)
}
func (m *QueryGetDenomInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDenomInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDenomInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDenomInfoResponse proto.InternalMessageInfo

func (m *QueryGetDenomInfoResponse) GetDenomInfo() DenomInfo {
	if m != nil {
		return m.DenomInfo
	}
	return DenomInfo{}
}

type QueryGetDenomCrossChainInfoRequest struct {
	Denom   string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	ChainId uint64 `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *QueryGetDenomCrossChainInfoRequest) Reset()         { *m = QueryGetDenomCrossChainInfoRequest{} }
func (m *QueryGetDenomCrossChainInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetDenomCrossChainInfoRequest) ProtoMessage()    {}
func (*QueryGetDenomCrossChainInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8421dffc0f307fe9, []int{2}
}
func (m *QueryGetDenomCrossChainInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDenomCrossChainInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDenomCrossChainInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDenomCrossChainInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDenomCrossChainInfoRequest.Merge(m, src)
}
func (m *QueryGetDenomCrossChainInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDenomCrossChainInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDenomCrossChainInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDenomCrossChainInfoRequest proto.InternalMessageInfo

type QueryGetDenomCrossChainInfoResponse struct {
	DenomInfo   DenomInfo `protobuf:"bytes,1,opt,name=denom_info,json=denomInfo,proto3" json:"denom_info"`
	ToChainId   uint64    `protobuf:"varint,2,opt,name=to_chain_id,json=toChainId,proto3" json:"to_chain_id,omitempty"`
	ToAssetHash string    `protobuf:"bytes,3,opt,name=to_asset_hash,json=toAssetHash,proto3" json:"to_asset_hash,omitempty"`
}

func (m *QueryGetDenomCrossChainInfoResponse) Reset()         { *m = QueryGetDenomCrossChainInfoResponse{} }
func (m *QueryGetDenomCrossChainInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetDenomCrossChainInfoResponse) ProtoMessage()    {}
func (*QueryGetDenomCrossChainInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8421dffc0f307fe9, []int{3}
}
func (m *QueryGetDenomCrossChainInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDenomCrossChainInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDenomCrossChainInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDenomCrossChainInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDenomCrossChainInfoResponse.Merge(m, src)
}
func (m *QueryGetDenomCrossChainInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDenomCrossChainInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDenomCrossChainInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDenomCrossChainInfoResponse proto.InternalMessageInfo

func (m *QueryGetDenomCrossChainInfoResponse) GetDenomInfo() DenomInfo {
	if m != nil {
		return m.DenomInfo
	}
	return DenomInfo{}
}

func (m *QueryGetDenomCrossChainInfoResponse) GetToChainId() uint64 {
	if m != nil {
		return m.ToChainId
	}
	return 0
}

func (m *QueryGetDenomCrossChainInfoResponse) GetToAssetHash() string {
	if m != nil {
		return m.ToAssetHash
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryGetDenomInfoRequest)(nil), "cosmospoly.btcx.QueryGetDenomInfoRequest")
	proto.RegisterType((*QueryGetDenomInfoResponse)(nil), "cosmospoly.btcx.QueryGetDenomInfoResponse")
	proto.RegisterType((*QueryGetDenomCrossChainInfoRequest)(nil), "cosmospoly.btcx.QueryGetDenomCrossChainInfoRequest")
	proto.RegisterType((*QueryGetDenomCrossChainInfoResponse)(nil), "cosmospoly.btcx.QueryGetDenomCrossChainInfoResponse")
}

func init() { proto.RegisterFile("cosmospoly/btcx/query.proto", fileDescriptor_8421dffc0f307fe9) }

var fileDescriptor_8421dffc0f307fe9 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0x2e, 0xc8, 0xcf, 0xa9, 0xd4, 0x4f, 0x2a, 0x49, 0xae, 0xd0, 0x2f, 0x2c, 0x4d, 0x2d,
	0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x47, 0x48, 0xea, 0x81, 0x24, 0xa5, 0x44,
	0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x72, 0xfa, 0x20, 0x16, 0x44, 0x99, 0x94, 0x4c, 0x7a, 0x7e, 0x7e,
	0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66,
	0x7e, 0x5e, 0x31, 0x54, 0x56, 0x01, 0xdd, 0x86, 0x94, 0xd4, 0xbc, 0xfc, 0xdc, 0xf8, 0xcc, 0xbc,
	0x34, 0xa8, 0x7e, 0x25, 0x33, 0x2e, 0x89, 0x40, 0x90, 0xad, 0xee, 0xa9, 0x25, 0x2e, 0x20, 0x39,
	0xcf, 0xbc, 0xb4, 0xfc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x11, 0x2e, 0x56, 0xb0,
	0x7a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0xc7, 0x8a, 0xa5, 0x63, 0x81, 0x3c, 0x83,
	0x52, 0x0c, 0x97, 0x24, 0x16, 0x7d, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xf6, 0x5c, 0x5c,
	0x08, 0x8b, 0xc0, 0xba, 0xb9, 0x8d, 0xa4, 0xf4, 0xd0, 0x3c, 0xa4, 0x07, 0xd7, 0xe7, 0xc4, 0x72,
	0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x67, 0x0a, 0x4c, 0x40, 0x29, 0x9a, 0x4b, 0x09, 0xc5, 0x74, 0xe7,
	0xa2, 0xfc, 0xe2, 0x62, 0xe7, 0x8c, 0xc4, 0xcc, 0x3c, 0x82, 0xee, 0x13, 0x92, 0xe4, 0xe2, 0x48,
	0x06, 0xa9, 0x8c, 0xcf, 0x4c, 0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x09, 0x62, 0x07, 0xf3, 0x3d,
	0x53, 0xa0, 0x4e, 0x5f, 0xc5, 0xc8, 0xa5, 0x8c, 0xd7, 0x74, 0x2a, 0xf9, 0x42, 0x48, 0x8e, 0x8b,
	0xbb, 0x24, 0x3f, 0x1e, 0xcd, 0x31, 0x9c, 0x25, 0xf9, 0x10, 0xab, 0x52, 0x84, 0x94, 0xb8, 0x78,
	0x4b, 0xf2, 0xe3, 0x13, 0x8b, 0x8b, 0x53, 0x4b, 0xe2, 0x33, 0x12, 0x8b, 0x33, 0x24, 0x98, 0xc1,
	0xfe, 0xe0, 0x2e, 0xc9, 0x77, 0x04, 0x89, 0x79, 0x24, 0x16, 0x67, 0x18, 0x7d, 0x62, 0xe2, 0x62,
	0x05, 0x3b, 0x56, 0x68, 0x1e, 0x23, 0x17, 0x27, 0xdc, 0x32, 0x21, 0x4d, 0x0c, 0x87, 0xe0, 0x8a,
	0x46, 0x29, 0x2d, 0x62, 0x94, 0x42, 0xfc, 0xac, 0x64, 0xd5, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x13,
	0x21, 0x23, 0xfd, 0xe0, 0xf2, 0xcc, 0x92, 0xe4, 0x8c, 0xd4, 0x7c, 0x7d, 0x90, 0xb6, 0xbc, 0xd4,
	0x92, 0xf2, 0xfc, 0xa2, 0x6c, 0x88, 0x39, 0x90, 0x94, 0x04, 0xd7, 0xab, 0x5f, 0x0d, 0xf6, 0x71,
	0xad, 0xd0, 0x71, 0x46, 0x2e, 0x61, 0x2c, 0xe1, 0x29, 0x64, 0x8c, 0xdf, 0x7e, 0xac, 0x71, 0x2b,
	0x65, 0x42, 0x9a, 0x26, 0xa8, 0xf3, 0x5d, 0xc1, 0xce, 0xb7, 0x17, 0xb2, 0x25, 0xdd, 0xf9, 0xfa,
	0xd5, 0xb0, 0x88, 0xaa, 0x75, 0xf2, 0x3d, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x28, 0xe3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xac, 0x56, 0xe8, 0x42,
	0xed, 0xa8, 0x80, 0xd8, 0x52, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xce, 0x6a, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x61, 0xe8, 0xb4, 0xf0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// this line is used by starport scaffolding # 2
	DenomInfo(ctx context.Context, in *QueryGetDenomInfoRequest, opts ...grpc.CallOption) (*QueryGetDenomInfoResponse, error)
	DenomCrossChainInfo(ctx context.Context, in *QueryGetDenomCrossChainInfoRequest, opts ...grpc.CallOption) (*QueryGetDenomCrossChainInfoResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DenomInfo(ctx context.Context, in *QueryGetDenomInfoRequest, opts ...grpc.CallOption) (*QueryGetDenomInfoResponse, error) {
	out := new(QueryGetDenomInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmospoly.btcx.Query/DenomInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomCrossChainInfo(ctx context.Context, in *QueryGetDenomCrossChainInfoRequest, opts ...grpc.CallOption) (*QueryGetDenomCrossChainInfoResponse, error) {
	out := new(QueryGetDenomCrossChainInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmospoly.btcx.Query/DenomCrossChainInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// this line is used by starport scaffolding # 2
	DenomInfo(context.Context, *QueryGetDenomInfoRequest) (*QueryGetDenomInfoResponse, error)
	DenomCrossChainInfo(context.Context, *QueryGetDenomCrossChainInfoRequest) (*QueryGetDenomCrossChainInfoResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DenomInfo(ctx context.Context, req *QueryGetDenomInfoRequest) (*QueryGetDenomInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomInfo not implemented")
}
func (*UnimplementedQueryServer) DenomCrossChainInfo(ctx context.Context, req *QueryGetDenomCrossChainInfoRequest) (*QueryGetDenomCrossChainInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomCrossChainInfo not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DenomInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetDenomInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmospoly.btcx.Query/DenomInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomInfo(ctx, req.(*QueryGetDenomInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomCrossChainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetDenomCrossChainInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomCrossChainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmospoly.btcx.Query/DenomCrossChainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomCrossChainInfo(ctx, req.(*QueryGetDenomCrossChainInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmospoly.btcx.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DenomInfo",
			Handler:    _Query_DenomInfo_Handler,
		},
		{
			MethodName: "DenomCrossChainInfo",
			Handler:    _Query_DenomCrossChainInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmospoly/btcx/query.proto",
}

func (m *QueryGetDenomInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDenomInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDenomInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetDenomInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDenomInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDenomInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DenomInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetDenomCrossChainInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDenomCrossChainInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDenomCrossChainInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ChainId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetDenomCrossChainInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDenomCrossChainInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDenomCrossChainInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ToAssetHash) > 0 {
		i -= len(m.ToAssetHash)
		copy(dAtA[i:], m.ToAssetHash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ToAssetHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ToChainId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ToChainId))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.DenomInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetDenomInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetDenomInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DenomInfo.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetDenomCrossChainInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.ChainId != 0 {
		n += 1 + sovQuery(uint64(m.ChainId))
	}
	return n
}

func (m *QueryGetDenomCrossChainInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DenomInfo.Size()
	n += 1 + l + sovQuery(uint64(l))
	if m.ToChainId != 0 {
		n += 1 + sovQuery(uint64(m.ToChainId))
	}
	l = len(m.ToAssetHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetDenomInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetDenomInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDenomInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetDenomInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetDenomInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDenomInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DenomInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetDenomCrossChainInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetDenomCrossChainInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDenomCrossChainInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetDenomCrossChainInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetDenomCrossChainInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDenomCrossChainInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DenomInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToChainId", wireType)
			}
			m.ToChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ToChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAssetHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAssetHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
