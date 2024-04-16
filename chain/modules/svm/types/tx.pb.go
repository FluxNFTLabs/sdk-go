// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/svm/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgTransaction struct {
	// signers are cosmos addresses that signs this message
	CosmosSigners []string       `protobuf:"bytes,1,rep,name=cosmos_signers,json=cosmosSigners,proto3" json:"cosmos_signers,omitempty"`
	Accounts      []string       `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	Instructions  []*Instruction `protobuf:"bytes,3,rep,name=instructions,proto3" json:"instructions,omitempty"`
	ComputeBudget uint64         `protobuf:"varint,4,opt,name=compute_budget,json=computeBudget,proto3" json:"compute_budget,omitempty"`
}

func (m *MsgTransaction) Reset()         { *m = MsgTransaction{} }
func (m *MsgTransaction) String() string { return proto.CompactTextString(m) }
func (*MsgTransaction) ProtoMessage()    {}
func (*MsgTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_6010395245cd7264, []int{0}
}
func (m *MsgTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransaction.Merge(m, src)
}
func (m *MsgTransaction) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransaction proto.InternalMessageInfo

type MsgTransactionResponse struct {
	UnitConsumed uint64 `protobuf:"varint,1,opt,name=unit_consumed,json=unitConsumed,proto3" json:"unit_consumed,omitempty"`
}

func (m *MsgTransactionResponse) Reset()         { *m = MsgTransactionResponse{} }
func (m *MsgTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTransactionResponse) ProtoMessage()    {}
func (*MsgTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6010395245cd7264, []int{1}
}
func (m *MsgTransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransactionResponse.Merge(m, src)
}
func (m *MsgTransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransactionResponse proto.InternalMessageInfo

func (m *MsgTransactionResponse) GetUnitConsumed() uint64 {
	if m != nil {
		return m.UnitConsumed
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgTransaction)(nil), "flux.svm.v1beta1.MsgTransaction")
	proto.RegisterType((*MsgTransactionResponse)(nil), "flux.svm.v1beta1.MsgTransactionResponse")
}

func init() { proto.RegisterFile("flux/svm/v1beta1/tx.proto", fileDescriptor_6010395245cd7264) }

var fileDescriptor_6010395245cd7264 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0x63, 0x52, 0x50, 0x31, 0x6d, 0x05, 0x16, 0x82, 0xeb, 0x49, 0x1c, 0x51, 0x11, 0x52,
	0x94, 0xe1, 0xac, 0x14, 0x26, 0x24, 0x06, 0x8a, 0x54, 0x09, 0x89, 0x76, 0x08, 0x15, 0x03, 0x4b,
	0xe4, 0xf3, 0x19, 0xd7, 0x52, 0x6c, 0x47, 0x79, 0x76, 0x14, 0x36, 0xc4, 0x84, 0x98, 0xf8, 0x08,
	0x7c, 0x84, 0x7c, 0x0c, 0xc6, 0x8e, 0x8c, 0x28, 0x19, 0x32, 0xf1, 0x1d, 0x90, 0x7d, 0x6e, 0x51,
	0x9a, 0x81, 0xe5, 0xee, 0xde, 0xef, 0xff, 0xf7, 0x7b, 0xff, 0x77, 0xc6, 0xfb, 0x1f, 0x47, 0x7e,
	0x46, 0x61, 0xaa, 0xe9, 0xb4, 0x5f, 0x09, 0xc7, 0xfa, 0xd4, 0xcd, 0xca, 0xf1, 0xc4, 0x3a, 0x4b,
	0xee, 0x06, 0xa9, 0x84, 0xa9, 0x2e, 0x93, 0x94, 0xef, 0x73, 0x0b, 0xda, 0xc2, 0x30, 0xea, 0xb4,
	0x29, 0x1a, 0x73, 0xfe, 0xb0, 0xa9, 0xa8, 0x06, 0x49, 0xa7, 0xfd, 0xf0, 0x4a, 0xc2, 0x7d, 0x69,
	0xa5, 0x6d, 0x0e, 0x84, 0xaf, 0x44, 0xef, 0x31, 0xad, 0x8c, 0xa5, 0xf1, 0x99, 0x50, 0x91, 0x3a,
	0x54, 0x0c, 0xc4, 0x55, 0x18, 0x6e, 0x95, 0x49, 0x7a, 0xbe, 0x91, 0x34, 0x44, 0x8b, 0xda, 0xc1,
	0x1f, 0x84, 0xf7, 0x4e, 0x40, 0x9e, 0x4d, 0x98, 0x01, 0xc6, 0x9d, 0xb2, 0x86, 0x3c, 0xc5, 0x7b,
	0x29, 0x2d, 0x28, 0x69, 0xc4, 0x04, 0x32, 0xd4, 0x69, 0x77, 0x6f, 0x0f, 0x76, 0x1b, 0xfa, 0xae,
	0x81, 0x24, 0xc7, 0xdb, 0x8c, 0x73, 0xeb, 0x8d, 0x83, 0xec, 0x46, 0x34, 0x5c, 0xd5, 0xe4, 0x15,
	0xde, 0x51, 0x06, 0xdc, 0xc4, 0xc7, 0x8e, 0x90, 0xb5, 0x3b, 0xed, 0xee, 0x9d, 0xc3, 0x47, 0xe5,
	0xf5, 0xff, 0x52, 0xbe, 0xf9, 0xe7, 0x1a, 0xac, 0x1d, 0x69, 0x52, 0xe8, 0xb1, 0x77, 0x62, 0x58,
	0xf9, 0x5a, 0x0a, 0x97, 0x6d, 0x75, 0x50, 0x77, 0x2b, 0xa4, 0x88, 0xf4, 0x28, 0xc2, 0x17, 0xe5,
	0xd7, 0x1f, 0x8f, 0x5b, 0x5f, 0x56, 0xf3, 0xde, 0xb5, 0xcc, 0xdf, 0x56, 0xf3, 0x1e, 0x09, 0x2b,
	0xaf, 0x2f, 0x77, 0xf0, 0x12, 0x3f, 0x58, 0x27, 0x03, 0x01, 0x63, 0x6b, 0x40, 0x90, 0x27, 0x78,
	0xd7, 0x1b, 0xe5, 0x86, 0xdc, 0x1a, 0xf0, 0x5a, 0xd4, 0x19, 0x8a, 0xf3, 0x76, 0x02, 0x7c, 0x9d,
	0xd8, 0x61, 0x8d, 0xdb, 0x27, 0x20, 0xc9, 0x7b, 0xbc, 0x7d, 0xd9, 0x82, 0x74, 0x36, 0xb7, 0x5a,
	0x9f, 0x90, 0x77, 0xff, 0xe7, 0xb8, 0xcc, 0x90, 0xdf, 0xfc, 0xbc, 0x9a, 0xf7, 0xd0, 0xd1, 0xe9,
	0xcf, 0x45, 0x81, 0x2e, 0x16, 0x05, 0xfa, 0xbd, 0x28, 0xd0, 0xf7, 0x65, 0xd1, 0xba, 0x58, 0x16,
	0xad, 0x5f, 0xcb, 0xa2, 0xf5, 0xe1, 0xb9, 0x54, 0xee, 0xdc, 0x57, 0x25, 0xb7, 0x9a, 0x1e, 0x8f,
	0xfc, 0xec, 0xf4, 0xf8, 0xec, 0x2d, 0xab, 0x80, 0x86, 0x01, 0x35, 0xe5, 0xe7, 0x4c, 0x19, 0xaa,
	0x6d, 0xed, 0x47, 0x02, 0xe2, 0x85, 0xbb, 0x4f, 0x63, 0x01, 0xd5, 0xad, 0x78, 0xd7, 0xcf, 0xfe,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x49, 0x81, 0x09, 0xaa, 0xb3, 0x02, 0x00, 0x00,
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
	Transact(ctx context.Context, in *MsgTransaction, opts ...grpc.CallOption) (*MsgTransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Transact(ctx context.Context, in *MsgTransaction, opts ...grpc.CallOption) (*MsgTransactionResponse, error) {
	out := new(MsgTransactionResponse)
	err := c.cc.Invoke(ctx, "/flux.svm.v1beta1.Msg/Transact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Transact(context.Context, *MsgTransaction) (*MsgTransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Transact(ctx context.Context, req *MsgTransaction) (*MsgTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transact not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Transact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Transact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flux.svm.v1beta1.Msg/Transact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Transact(ctx, req.(*MsgTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flux.svm.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transact",
			Handler:    _Msg_Transact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flux/svm/v1beta1/tx.proto",
}

func (m *MsgTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ComputeBudget != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ComputeBudget))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Instructions) > 0 {
		for iNdEx := len(m.Instructions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Instructions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Accounts[iNdEx])
			copy(dAtA[i:], m.Accounts[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Accounts[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.CosmosSigners) > 0 {
		for iNdEx := len(m.CosmosSigners) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CosmosSigners[iNdEx])
			copy(dAtA[i:], m.CosmosSigners[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.CosmosSigners[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnitConsumed != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.UnitConsumed))
		i--
		dAtA[i] = 0x8
	}
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
func (m *MsgTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CosmosSigners) > 0 {
		for _, s := range m.CosmosSigners {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Accounts) > 0 {
		for _, s := range m.Accounts {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Instructions) > 0 {
		for _, e := range m.Instructions {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.ComputeBudget != 0 {
		n += 1 + sovTx(uint64(m.ComputeBudget))
	}
	return n
}

func (m *MsgTransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UnitConsumed != 0 {
		n += 1 + sovTx(uint64(m.UnitConsumed))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgTransaction) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosSigners", wireType)
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
			m.CosmosSigners = append(m.CosmosSigners, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
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
			m.Accounts = append(m.Accounts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instructions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Instructions = append(m.Instructions, &Instruction{})
			if err := m.Instructions[len(m.Instructions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputeBudget", wireType)
			}
			m.ComputeBudget = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComputeBudget |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgTransactionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgTransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnitConsumed", wireType)
			}
			m.UnitConsumed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnitConsumed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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