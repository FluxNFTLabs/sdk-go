// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/svm/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

type AccountRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *AccountRequest) Reset()         { *m = AccountRequest{} }
func (m *AccountRequest) String() string { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()    {}
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99679498a27461b0, []int{0}
}
func (m *AccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRequest.Merge(m, src)
}
func (m *AccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *AccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRequest proto.InternalMessageInfo

func (m *AccountRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type AccountResponse struct {
	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *AccountResponse) Reset()         { *m = AccountResponse{} }
func (m *AccountResponse) String() string { return proto.CompactTextString(m) }
func (*AccountResponse) ProtoMessage()    {}
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99679498a27461b0, []int{1}
}
func (m *AccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountResponse.Merge(m, src)
}
func (m *AccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *AccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountResponse proto.InternalMessageInfo

func (m *AccountResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type ProgramAccountsRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *ProgramAccountsRequest) Reset()         { *m = ProgramAccountsRequest{} }
func (m *ProgramAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*ProgramAccountsRequest) ProtoMessage()    {}
func (*ProgramAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99679498a27461b0, []int{2}
}
func (m *ProgramAccountsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProgramAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProgramAccountsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProgramAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProgramAccountsRequest.Merge(m, src)
}
func (m *ProgramAccountsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProgramAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProgramAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProgramAccountsRequest proto.InternalMessageInfo

func (m *ProgramAccountsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ProgramAccountsResponse struct {
	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (m *ProgramAccountsResponse) Reset()         { *m = ProgramAccountsResponse{} }
func (m *ProgramAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*ProgramAccountsResponse) ProtoMessage()    {}
func (*ProgramAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99679498a27461b0, []int{3}
}
func (m *ProgramAccountsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProgramAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProgramAccountsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProgramAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProgramAccountsResponse.Merge(m, src)
}
func (m *ProgramAccountsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ProgramAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProgramAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProgramAccountsResponse proto.InternalMessageInfo

func (m *ProgramAccountsResponse) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountRequest)(nil), "flux.svm.v1beta1.AccountRequest")
	proto.RegisterType((*AccountResponse)(nil), "flux.svm.v1beta1.AccountResponse")
	proto.RegisterType((*ProgramAccountsRequest)(nil), "flux.svm.v1beta1.ProgramAccountsRequest")
	proto.RegisterType((*ProgramAccountsResponse)(nil), "flux.svm.v1beta1.ProgramAccountsResponse")
}

func init() { proto.RegisterFile("flux/svm/v1beta1/query.proto", fileDescriptor_99679498a27461b0) }

var fileDescriptor_99679498a27461b0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6b, 0xe2, 0x40,
	0x18, 0xc6, 0x8d, 0xcb, 0xae, 0x38, 0x0b, 0xeb, 0x32, 0x87, 0x5d, 0x37, 0x48, 0x70, 0xb3, 0x2c,
	0x58, 0x5b, 0x32, 0xf8, 0x07, 0x7a, 0x6e, 0x0f, 0x9e, 0x8a, 0xb4, 0xd2, 0x53, 0x2f, 0x65, 0x92,
	0x4c, 0x63, 0x20, 0xc9, 0xc4, 0xbc, 0x13, 0x51, 0xa4, 0x97, 0x7e, 0x82, 0x42, 0xaf, 0xbd, 0xf5,
	0xcb, 0xf4, 0x28, 0xf4, 0xd2, 0x63, 0xd1, 0x7e, 0x90, 0x92, 0x64, 0x54, 0x30, 0xfd, 0x77, 0xcc,
	0xfb, 0x3c, 0x4f, 0x9e, 0xdf, 0xfb, 0x32, 0xa8, 0x76, 0xe1, 0xc5, 0x13, 0x02, 0x63, 0x9f, 0x8c,
	0x5b, 0x26, 0x13, 0xb4, 0x45, 0x46, 0x31, 0x8b, 0xa6, 0x46, 0x18, 0x71, 0xc1, 0xf1, 0xcf, 0x44,
	0x35, 0x60, 0xec, 0x1b, 0x52, 0x55, 0x9b, 0x16, 0x07, 0x9f, 0x03, 0x31, 0x29, 0xb0, 0xcc, 0xba,
	0x0e, 0x86, 0xd4, 0x71, 0x03, 0x2a, 0x5c, 0x1e, 0x64, 0x69, 0xb5, 0xe6, 0x70, 0xee, 0x78, 0x8c,
	0xd0, 0xd0, 0x25, 0x34, 0x08, 0xb8, 0x48, 0x45, 0x90, 0xaa, 0x9a, 0x6b, 0x4e, 0x7a, 0x52, 0x4d,
	0x6f, 0xa2, 0x1f, 0x07, 0x96, 0xc5, 0xe3, 0x40, 0x0c, 0xd8, 0x28, 0x66, 0x20, 0x70, 0x15, 0x95,
	0xa8, 0x6d, 0x47, 0x0c, 0xa0, 0xaa, 0xd4, 0x95, 0x46, 0x79, 0xb0, 0xfa, 0xd4, 0x7b, 0xa8, 0xb2,
	0xf6, 0x42, 0xc8, 0x03, 0x60, 0xb8, 0x83, 0x4a, 0x34, 0x1b, 0xa5, 0xe6, 0xef, 0xed, 0x3f, 0xc6,
	0xf6, 0x22, 0xc6, 0x2a, 0xb3, 0x72, 0xea, 0x6d, 0xf4, 0xeb, 0x38, 0xe2, 0x4e, 0x44, 0x7d, 0x29,
	0xc1, 0xc7, 0xdd, 0xfb, 0xe8, 0x77, 0x2e, 0x23, 0x19, 0x6a, 0xa8, 0x2c, 0x5d, 0x2c, 0x89, 0x7d,
	0x69, 0x94, 0x07, 0x9b, 0x41, 0xfb, 0xae, 0x88, 0xbe, 0x9e, 0x24, 0xd7, 0xc3, 0x33, 0x54, 0x92,
	0x59, 0x5c, 0x7f, 0x9b, 0x32, 0x23, 0x51, 0xff, 0xbe, 0xe3, 0xc8, 0x7a, 0xf5, 0xdd, 0xab, 0x87,
	0xe7, 0x9b, 0xe2, 0x7f, 0xfc, 0x8f, 0xe4, 0xee, 0x2b, 0x37, 0x25, 0x33, 0xc9, 0x71, 0x89, 0x6f,
	0x15, 0x54, 0xd9, 0x5a, 0x00, 0x37, 0xf2, 0x1d, 0xaf, 0xdf, 0x45, 0xdd, 0xf9, 0x84, 0x53, 0x52,
	0x75, 0x53, 0x2a, 0x03, 0xef, 0xe5, 0xa9, 0xc2, 0x2c, 0x72, 0x2e, 0xe9, 0x60, 0x83, 0x77, 0xd8,
	0xbf, 0x5f, 0x68, 0xca, 0x7c, 0xa1, 0x29, 0x4f, 0x0b, 0x4d, 0xb9, 0x5e, 0x6a, 0x85, 0xf9, 0x52,
	0x2b, 0x3c, 0x2e, 0xb5, 0xc2, 0x59, 0xd7, 0x71, 0xc5, 0x30, 0x36, 0x0d, 0x8b, 0xfb, 0xa4, 0xe7,
	0xc5, 0x93, 0x7e, 0xef, 0xf4, 0x88, 0x9a, 0x90, 0xfe, 0xdd, 0x26, 0xd6, 0x90, 0xba, 0x01, 0xf1,
	0xb9, 0x1d, 0x7b, 0x0c, 0xd2, 0x32, 0x31, 0x0d, 0x19, 0x98, 0xdf, 0xd2, 0xd7, 0xd5, 0x79, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x84, 0xa2, 0x5e, 0xa2, 0xf5, 0x02, 0x00, 0x00,
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
	Account(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	ProgramAccounts(ctx context.Context, in *ProgramAccountsRequest, opts ...grpc.CallOption) (*ProgramAccountsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Account(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/flux.svm.v1beta1.Query/Account", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ProgramAccounts(ctx context.Context, in *ProgramAccountsRequest, opts ...grpc.CallOption) (*ProgramAccountsResponse, error) {
	out := new(ProgramAccountsResponse)
	err := c.cc.Invoke(ctx, "/flux.svm.v1beta1.Query/ProgramAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Account(context.Context, *AccountRequest) (*AccountResponse, error)
	ProgramAccounts(context.Context, *ProgramAccountsRequest) (*ProgramAccountsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Account(ctx context.Context, req *AccountRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Account not implemented")
}
func (*UnimplementedQueryServer) ProgramAccounts(ctx context.Context, req *ProgramAccountsRequest) (*ProgramAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProgramAccounts not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Account(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flux.svm.v1beta1.Query/Account",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Account(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ProgramAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProgramAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProgramAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flux.svm.v1beta1.Query/ProgramAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProgramAccounts(ctx, req.(*ProgramAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flux.svm.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Account",
			Handler:    _Query_Account_Handler,
		},
		{
			MethodName: "ProgramAccounts",
			Handler:    _Query_ProgramAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flux/svm/v1beta1/query.proto",
}

func (m *AccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Account != nil {
		{
			size, err := m.Account.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProgramAccountsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProgramAccountsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProgramAccountsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProgramAccountsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProgramAccountsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProgramAccountsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
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
func (m *AccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *AccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Account != nil {
		l = m.Account.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ProgramAccountsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ProgramAccountsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *AccountResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
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
			if m.Account == nil {
				m.Account = &Account{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ProgramAccountsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ProgramAccountsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProgramAccountsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *ProgramAccountsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ProgramAccountsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProgramAccountsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
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
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
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
