// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ccm/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
type QueryCheckModuleContractRequest struct {
	ModuleName     string `protobuf:"bytes,1,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	ToContractAddr []byte `protobuf:"bytes,2,opt,name=toContractAddr,proto3" json:"toContractAddr,omitempty"`
	FromChainID    uint64 `protobuf:"varint,3,opt,name=fromChainID,proto3" json:"fromChainID,omitempty"`
}

func (m *QueryCheckModuleContractRequest) Reset()         { *m = QueryCheckModuleContractRequest{} }
func (m *QueryCheckModuleContractRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCheckModuleContractRequest) ProtoMessage()    {}
func (*QueryCheckModuleContractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd452b9385aefc5f, []int{0}
}
func (m *QueryCheckModuleContractRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCheckModuleContractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCheckModuleContractRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCheckModuleContractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCheckModuleContractRequest.Merge(m, src)
}
func (m *QueryCheckModuleContractRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCheckModuleContractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCheckModuleContractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCheckModuleContractRequest proto.InternalMessageInfo

func (m *QueryCheckModuleContractRequest) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *QueryCheckModuleContractRequest) GetToContractAddr() []byte {
	if m != nil {
		return m.ToContractAddr
	}
	return nil
}

func (m *QueryCheckModuleContractRequest) GetFromChainID() uint64 {
	if m != nil {
		return m.FromChainID
	}
	return 0
}

type QueryCheckModuleContractResponse struct {
	ModuleName string `protobuf:"bytes,1,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	Exist      bool   `protobuf:"varint,2,opt,name=exist,proto3" json:"exist,omitempty"`
}

func (m *QueryCheckModuleContractResponse) Reset()         { *m = QueryCheckModuleContractResponse{} }
func (m *QueryCheckModuleContractResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCheckModuleContractResponse) ProtoMessage()    {}
func (*QueryCheckModuleContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd452b9385aefc5f, []int{1}
}
func (m *QueryCheckModuleContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCheckModuleContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCheckModuleContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCheckModuleContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCheckModuleContractResponse.Merge(m, src)
}
func (m *QueryCheckModuleContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCheckModuleContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCheckModuleContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCheckModuleContractResponse proto.InternalMessageInfo

func (m *QueryCheckModuleContractResponse) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *QueryCheckModuleContractResponse) GetExist() bool {
	if m != nil {
		return m.Exist
	}
	return false
}

func init() {
	proto.RegisterType((*QueryCheckModuleContractRequest)(nil), "Switcheo.polynetworkcosmos.ccm.QueryCheckModuleContractRequest")
	proto.RegisterType((*QueryCheckModuleContractResponse)(nil), "Switcheo.polynetworkcosmos.ccm.QueryCheckModuleContractResponse")
}

func init() { proto.RegisterFile("ccm/query.proto", fileDescriptor_bd452b9385aefc5f) }

var fileDescriptor_bd452b9385aefc5f = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0x1b, 0x41,
	0x18, 0xc6, 0x33, 0xdb, 0xa6, 0xb4, 0xd3, 0xd2, 0xc2, 0xb6, 0x87, 0x10, 0xca, 0x74, 0xc9, 0xa1,
	0x84, 0x42, 0x77, 0x48, 0xfa, 0x01, 0xfa, 0x27, 0xbd, 0x08, 0x2a, 0xb8, 0x5e, 0xc4, 0x4b, 0x98,
	0x9d, 0x8c, 0xbb, 0x43, 0xb2, 0xf3, 0x6e, 0x76, 0x66, 0x4d, 0x42, 0xc8, 0x45, 0x8f, 0x22, 0x08,
	0x7e, 0x29, 0x8f, 0x01, 0x2f, 0x1e, 0x25, 0xf1, 0x83, 0x48, 0x76, 0x13, 0x8c, 0x1a, 0x0d, 0x78,
	0x9c, 0x67, 0x9e, 0xf7, 0xc7, 0xf3, 0x3e, 0xbc, 0xf8, 0x13, 0xe7, 0x11, 0xed, 0xa6, 0x22, 0x19,
	0xb8, 0x71, 0x02, 0x06, 0x6c, 0xb2, 0xdb, 0x93, 0x86, 0x87, 0x02, 0xdc, 0x18, 0x3a, 0x03, 0x25,
	0x4c, 0x0f, 0x92, 0x36, 0x07, 0x1d, 0x81, 0x76, 0x39, 0x8f, 0xca, 0x5f, 0x03, 0x80, 0xa0, 0x23,
	0x28, 0x8b, 0x25, 0x65, 0x4a, 0x81, 0x61, 0x46, 0x82, 0xd2, 0xf9, 0x74, 0xf9, 0x47, 0xee, 0xa4,
	0x3e, 0xd3, 0x22, 0xc7, 0xd2, 0xc3, 0x9a, 0x2f, 0x0c, 0xab, 0xd1, 0x98, 0x05, 0x52, 0x65, 0xe6,
	0xdc, 0x5b, 0x39, 0x41, 0xf8, 0xdb, 0xce, 0xcc, 0xd2, 0x08, 0x05, 0x6f, 0x6f, 0x41, 0x2b, 0xed,
	0x88, 0x06, 0x28, 0x93, 0x30, 0x6e, 0x3c, 0xd1, 0x4d, 0x85, 0x36, 0x36, 0xc1, 0x38, 0xca, 0x3e,
	0xb6, 0x59, 0x24, 0x4a, 0xc8, 0x41, 0xd5, 0x77, 0xde, 0x92, 0x62, 0x7f, 0xc7, 0x1f, 0x0d, 0x2c,
	0x86, 0xfe, 0xb6, 0x5a, 0x49, 0xc9, 0x72, 0x50, 0xf5, 0x83, 0xf7, 0x40, 0xb5, 0x1d, 0xfc, 0xfe,
	0x20, 0x81, 0xa8, 0x11, 0x32, 0xa9, 0x36, 0xfe, 0x97, 0x5e, 0x39, 0xa8, 0xfa, 0xda, 0x5b, 0x96,
	0x2a, 0x7b, 0xd8, 0x79, 0x3a, 0x8c, 0x8e, 0x41, 0x69, 0xb1, 0x36, 0xcd, 0x17, 0x5c, 0x14, 0x7d,
	0xa9, 0x4d, 0x16, 0xe2, 0xad, 0x97, 0x3f, 0xea, 0xa7, 0x16, 0x2e, 0x66, 0x68, 0xfb, 0xd8, 0xc2,
	0x9f, 0x57, 0xf0, 0xed, 0xdf, 0xee, 0xf3, 0xa5, 0xbb, 0x6b, 0x6a, 0x2a, 0xff, 0x79, 0x39, 0x20,
	0x5f, 0xad, 0x22, 0x8f, 0x2e, 0x6f, 0xce, 0x2d, 0x6e, 0x33, 0xba, 0x20, 0xd1, 0x47, 0x24, 0x3a,
	0xbb, 0x15, 0x3e, 0x83, 0x34, 0xf3, 0xa5, 0x9b, 0x7c, 0x8e, 0xa1, 0xc3, 0xbb, 0x16, 0x46, 0x74,
	0x78, 0xbf, 0xfc, 0x11, 0x1d, 0x2e, 0x15, 0x3d, 0xfa, 0xb7, 0x79, 0x31, 0x21, 0x68, 0x3c, 0x21,
	0xe8, 0x7a, 0x42, 0xd0, 0xd9, 0x94, 0x14, 0xc6, 0x53, 0x52, 0xb8, 0x9a, 0x92, 0xc2, 0x7e, 0x3d,
	0x90, 0x26, 0x4c, 0x7d, 0x97, 0x43, 0xb4, 0x32, 0xc6, 0xcf, 0x79, 0x8e, 0x7e, 0x96, 0xc4, 0x0c,
	0x62, 0xa1, 0xfd, 0x37, 0xd9, 0x31, 0xfd, 0xba, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xa8, 0x98,
	0x4b, 0xc9, 0x02, 0x00, 0x00,
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
	CheckModuleContract(ctx context.Context, in *QueryCheckModuleContractRequest, opts ...grpc.CallOption) (*QueryCheckModuleContractResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CheckModuleContract(ctx context.Context, in *QueryCheckModuleContractRequest, opts ...grpc.CallOption) (*QueryCheckModuleContractResponse, error) {
	out := new(QueryCheckModuleContractResponse)
	err := c.cc.Invoke(ctx, "/Switcheo.polynetworkcosmos.ccm.Query/CheckModuleContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// this line is used by starport scaffolding # 2
	CheckModuleContract(context.Context, *QueryCheckModuleContractRequest) (*QueryCheckModuleContractResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CheckModuleContract(ctx context.Context, req *QueryCheckModuleContractRequest) (*QueryCheckModuleContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckModuleContract not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CheckModuleContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCheckModuleContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CheckModuleContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Switcheo.polynetworkcosmos.ccm.Query/CheckModuleContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CheckModuleContract(ctx, req.(*QueryCheckModuleContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Switcheo.polynetworkcosmos.ccm.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckModuleContract",
			Handler:    _Query_CheckModuleContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ccm/query.proto",
}

func (m *QueryCheckModuleContractRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCheckModuleContractRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCheckModuleContractRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FromChainID != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.FromChainID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ToContractAddr) > 0 {
		i -= len(m.ToContractAddr)
		copy(dAtA[i:], m.ToContractAddr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ToContractAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCheckModuleContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCheckModuleContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCheckModuleContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Exist {
		i--
		if m.Exist {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *QueryCheckModuleContractRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ToContractAddr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.FromChainID != 0 {
		n += 1 + sovQuery(uint64(m.FromChainID))
	}
	return n
}

func (m *QueryCheckModuleContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Exist {
		n += 2
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCheckModuleContractRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCheckModuleContractRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCheckModuleContractRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
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
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToContractAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToContractAddr = append(m.ToContractAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.ToContractAddr == nil {
				m.ToContractAddr = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromChainID", wireType)
			}
			m.FromChainID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FromChainID |= uint64(b&0x7F) << shift
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
func (m *QueryCheckModuleContractResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCheckModuleContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCheckModuleContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
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
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exist", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Exist = bool(v != 0)
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
