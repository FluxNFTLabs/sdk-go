// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/strategy/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

type QueryFoo struct {
}

func (m *QueryFoo) Reset()         { *m = QueryFoo{} }
func (m *QueryFoo) String() string { return proto.CompactTextString(m) }
func (*QueryFoo) ProtoMessage()    {}
func (*QueryFoo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aa31d246011d73a, []int{0}
}
func (m *QueryFoo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFoo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFoo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFoo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFoo.Merge(m, src)
}
func (m *QueryFoo) XXX_Size() int {
	return m.Size()
}
func (m *QueryFoo) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFoo.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFoo proto.InternalMessageInfo

type QueryFooResponse struct {
}

func (m *QueryFooResponse) Reset()         { *m = QueryFooResponse{} }
func (m *QueryFooResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFooResponse) ProtoMessage()    {}
func (*QueryFooResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aa31d246011d73a, []int{1}
}
func (m *QueryFooResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFooResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFooResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFooResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFooResponse.Merge(m, src)
}
func (m *QueryFooResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFooResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFooResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFooResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryFoo)(nil), "flux.strategy.v1beta1.QueryFoo")
	proto.RegisterType((*QueryFooResponse)(nil), "flux.strategy.v1beta1.QueryFooResponse")
}

func init() { proto.RegisterFile("flux/strategy/v1beta1/query.proto", fileDescriptor_0aa31d246011d73a) }

var fileDescriptor_0aa31d246011d73a = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xbd, 0x4a, 0xf4, 0x40,
	0x14, 0x86, 0x77, 0xbf, 0x0f, 0x45, 0xa6, 0x92, 0x61, 0x6d, 0x82, 0x8c, 0x18, 0x04, 0xc1, 0x62,
	0x0e, 0xab, 0x95, 0xad, 0x45, 0x2a, 0x11, 0xfc, 0xa9, 0xec, 0x26, 0xd9, 0xd9, 0xd9, 0xc1, 0x64,
	0x4e, 0xcc, 0x4c, 0x64, 0xd3, 0x7a, 0x05, 0x82, 0x37, 0x65, 0xb9, 0x60, 0x63, 0x29, 0x89, 0x17,
	0x22, 0x19, 0x93, 0xa8, 0xb0, 0x60, 0x77, 0x98, 0xf7, 0x99, 0xe7, 0xbc, 0x87, 0xec, 0xcf, 0xd3,
	0x72, 0x09, 0xd6, 0x15, 0xc2, 0x49, 0x55, 0xc1, 0xc3, 0x34, 0x96, 0x4e, 0x4c, 0xe1, 0xbe, 0x94,
	0x45, 0xc5, 0xf3, 0x02, 0x1d, 0xd2, 0x9d, 0x16, 0xe1, 0x3d, 0xc2, 0x3b, 0x24, 0x98, 0x28, 0x54,
	0xe8, 0x09, 0x68, 0xa7, 0x2f, 0x38, 0x60, 0x09, 0xda, 0x0c, 0x2d, 0xc4, 0xc2, 0xca, 0xc1, 0x96,
	0xa0, 0x36, 0x5d, 0x7e, 0xf4, 0x33, 0xf7, 0x5b, 0x06, 0x2a, 0x17, 0x4a, 0x1b, 0xe1, 0x34, 0xf6,
	0xec, 0xae, 0x42, 0x54, 0xa9, 0x04, 0x91, 0x6b, 0x10, 0xc6, 0xa0, 0xf3, 0xa1, 0xed, 0xd2, 0x83,
	0xf5, 0xcd, 0x87, 0x9e, 0x9e, 0x0a, 0x09, 0xd9, 0xba, 0x6c, 0xb7, 0x44, 0x88, 0x21, 0x25, 0xdb,
	0xfd, 0x7c, 0x25, 0x6d, 0x8e, 0xc6, 0xca, 0x63, 0x47, 0x36, 0xfc, 0x1b, 0xbd, 0x23, 0xff, 0x23,
	0x44, 0xba, 0xc7, 0xd7, 0x5e, 0xcb, 0xfb, 0x8f, 0xc1, 0xe1, 0x1f, 0x40, 0x6f, 0x0e, 0x83, 0xc7,
	0xd7, 0x8f, 0xe7, 0x7f, 0x13, 0x4a, 0xe1, 0x77, 0xd1, 0x39, 0xe2, 0xd9, 0xf5, 0x4b, 0xcd, 0xc6,
	0xab, 0x9a, 0x8d, 0xdf, 0x6b, 0x36, 0x7e, 0x6a, 0xd8, 0x68, 0xd5, 0xb0, 0xd1, 0x5b, 0xc3, 0x46,
	0xb7, 0xa7, 0x4a, 0xbb, 0x45, 0x19, 0xf3, 0x04, 0x33, 0x88, 0xd2, 0x72, 0x79, 0x11, 0xdd, 0x9c,
	0x8b, 0xd8, 0x7a, 0xc7, 0x0c, 0x92, 0x85, 0xd0, 0x06, 0x32, 0x9c, 0x95, 0xa9, 0xb4, 0xdf, 0x4a,
	0x57, 0xe5, 0xd2, 0xc6, 0x9b, 0xfe, 0xe2, 0x93, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xb9,
	0xe9, 0x00, 0xd3, 0x01, 0x00, 0x00,
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
	Foo(ctx context.Context, in *QueryFoo, opts ...grpc.CallOption) (*QueryFooResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Foo(ctx context.Context, in *QueryFoo, opts ...grpc.CallOption) (*QueryFooResponse, error) {
	out := new(QueryFooResponse)
	err := c.cc.Invoke(ctx, "/flux.strategy.v1beta1.Query/Foo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Foo(context.Context, *QueryFoo) (*QueryFooResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Foo(ctx context.Context, req *QueryFoo) (*QueryFooResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Foo not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFoo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flux.strategy.v1beta1.Query/Foo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Foo(ctx, req.(*QueryFoo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flux.strategy.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Foo",
			Handler:    _Query_Foo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flux/strategy/v1beta1/query.proto",
}

func (m *QueryFoo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFoo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFoo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryFooResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFooResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFooResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *QueryFoo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryFooResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryFoo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFoo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFoo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryFooResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFooResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFooResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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