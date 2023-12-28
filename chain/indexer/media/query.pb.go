// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/indexer/media/query.proto

package media

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

type S3Operation int32

const (
	S3Operation_Put    S3Operation = 0
	S3Operation_Delete S3Operation = 1
)

var S3Operation_name = map[int32]string{
	0: "Put",
	1: "Delete",
}

var S3Operation_value = map[string]int32{
	"Put":    0,
	"Delete": 1,
}

func (x S3Operation) String() string {
	return proto.EnumName(S3Operation_name, int32(x))
}

func (S3Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_06796104fc06fe94, []int{0}
}

type PresignedURLRequest struct {
	Op  S3Operation `protobuf:"varint,1,opt,name=op,proto3,enum=flux.indexer.media.S3Operation" json:"op,omitempty"`
	Key string      `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *PresignedURLRequest) Reset()         { *m = PresignedURLRequest{} }
func (m *PresignedURLRequest) String() string { return proto.CompactTextString(m) }
func (*PresignedURLRequest) ProtoMessage()    {}
func (*PresignedURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06796104fc06fe94, []int{0}
}
func (m *PresignedURLRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PresignedURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PresignedURLRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PresignedURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PresignedURLRequest.Merge(m, src)
}
func (m *PresignedURLRequest) XXX_Size() int {
	return m.Size()
}
func (m *PresignedURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PresignedURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PresignedURLRequest proto.InternalMessageInfo

func (m *PresignedURLRequest) GetOp() S3Operation {
	if m != nil {
		return m.Op
	}
	return S3Operation_Put
}

func (m *PresignedURLRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type PresignedURLResponse struct {
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *PresignedURLResponse) Reset()         { *m = PresignedURLResponse{} }
func (m *PresignedURLResponse) String() string { return proto.CompactTextString(m) }
func (*PresignedURLResponse) ProtoMessage()    {}
func (*PresignedURLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06796104fc06fe94, []int{1}
}
func (m *PresignedURLResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PresignedURLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PresignedURLResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PresignedURLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PresignedURLResponse.Merge(m, src)
}
func (m *PresignedURLResponse) XXX_Size() int {
	return m.Size()
}
func (m *PresignedURLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PresignedURLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PresignedURLResponse proto.InternalMessageInfo

func (m *PresignedURLResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type GetMedataRequest struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Url   string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *GetMedataRequest) Reset()         { *m = GetMedataRequest{} }
func (m *GetMedataRequest) String() string { return proto.CompactTextString(m) }
func (*GetMedataRequest) ProtoMessage()    {}
func (*GetMedataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06796104fc06fe94, []int{2}
}
func (m *GetMedataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetMedataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetMedataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetMedataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMedataRequest.Merge(m, src)
}
func (m *GetMedataRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetMedataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMedataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMedataRequest proto.InternalMessageInfo

func (m *GetMedataRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GetMedataRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetMedataRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Metadata struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" bson:"owner"`
	Path  string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty" bson:"path"`
	Key   string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty" bson:"key"`
	Url   string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty" bson:"url"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_06796104fc06fe94, []int{3}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Metadata) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Metadata) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Metadata) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type SetMetadataResponse struct {
}

func (m *SetMetadataResponse) Reset()         { *m = SetMetadataResponse{} }
func (m *SetMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*SetMetadataResponse) ProtoMessage()    {}
func (*SetMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06796104fc06fe94, []int{4}
}
func (m *SetMetadataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetMetadataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMetadataResponse.Merge(m, src)
}
func (m *SetMetadataResponse) XXX_Size() int {
	return m.Size()
}
func (m *SetMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetMetadataResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("flux.indexer.media.S3Operation", S3Operation_name, S3Operation_value)
	proto.RegisterType((*PresignedURLRequest)(nil), "flux.indexer.media.PresignedURLRequest")
	proto.RegisterType((*PresignedURLResponse)(nil), "flux.indexer.media.PresignedURLResponse")
	proto.RegisterType((*GetMedataRequest)(nil), "flux.indexer.media.GetMedataRequest")
	proto.RegisterType((*Metadata)(nil), "flux.indexer.media.Metadata")
	proto.RegisterType((*SetMetadataResponse)(nil), "flux.indexer.media.SetMetadataResponse")
}

func init() { proto.RegisterFile("flux/indexer/media/query.proto", fileDescriptor_06796104fc06fe94) }

var fileDescriptor_06796104fc06fe94 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xe3, 0x52, 0xe8, 0xa6, 0x2a, 0xd6, 0x36, 0x48, 0x95, 0x15, 0x39, 0xd6, 0x52, 0x95,
	0x88, 0x83, 0x17, 0xb5, 0x37, 0x6e, 0x44, 0x55, 0x2b, 0x50, 0x0a, 0x91, 0x03, 0x12, 0xe2, 0xb6,
	0xa9, 0x07, 0xc7, 0xc2, 0xf1, 0xba, 0xf6, 0x5a, 0x4d, 0x39, 0x72, 0x40, 0xe2, 0x86, 0xc4, 0x91,
	0x1f, 0xe2, 0x58, 0x89, 0x0b, 0xa7, 0x08, 0x25, 0x7c, 0x41, 0xbe, 0x00, 0xed, 0x6e, 0x2d, 0x92,
	0xd6, 0x2a, 0xb7, 0xc9, 0xcc, 0xdb, 0xf7, 0xde, 0xbc, 0x89, 0x91, 0xf3, 0x3e, 0x2e, 0x26, 0x34,
	0x4a, 0x02, 0x98, 0x40, 0x46, 0xc7, 0x10, 0x44, 0x8c, 0x9e, 0x15, 0x90, 0x5d, 0x78, 0x69, 0xc6,
	0x05, 0xc7, 0x58, 0xce, 0xbd, 0xab, 0xb9, 0xa7, 0xe6, 0x76, 0x2b, 0xe4, 0x3c, 0x8c, 0x81, 0xb2,
	0x34, 0xa2, 0x2c, 0x49, 0xb8, 0x60, 0x22, 0xe2, 0x49, 0xae, 0x5f, 0xd8, 0xcd, 0x90, 0x87, 0x5c,
	0x95, 0x54, 0x56, 0xba, 0x4b, 0xde, 0xa2, 0xed, 0x7e, 0x06, 0x79, 0x14, 0x26, 0x10, 0xbc, 0xf1,
	0x7b, 0x3e, 0x9c, 0x15, 0x90, 0x0b, 0x4c, 0x51, 0x9d, 0xa7, 0x3b, 0x86, 0x6b, 0x74, 0xb6, 0xf6,
	0xdb, 0xde, 0x4d, 0x2d, 0x6f, 0x70, 0xf0, 0x2a, 0x85, 0x4c, 0x09, 0xf8, 0x75, 0x9e, 0x62, 0x0b,
	0x99, 0x1f, 0xe0, 0x62, 0xa7, 0xee, 0x1a, 0x9d, 0x0d, 0x5f, 0x96, 0xa4, 0x83, 0x9a, 0xab, 0xcc,
	0x79, 0xca, 0x93, 0x1c, 0x24, 0xb2, 0xc8, 0x62, 0xc5, 0xbd, 0xe1, 0xcb, 0x92, 0xf4, 0x90, 0x75,
	0x0c, 0xe2, 0x04, 0x02, 0x26, 0x58, 0x69, 0xa0, 0x89, 0xee, 0xf0, 0xf3, 0x04, 0xb2, 0x2b, 0x9c,
	0xfe, 0x71, 0x53, 0xa5, 0x64, 0x33, 0xff, 0xb1, 0x7d, 0x37, 0xd0, 0xbd, 0x13, 0x10, 0x4c, 0xb2,
	0xe1, 0xbd, 0x15, 0x9a, 0xae, 0xb5, 0x98, 0xb6, 0x37, 0x87, 0x39, 0x4f, 0x9e, 0x12, 0xd5, 0x26,
	0x25, 0xf1, 0x43, 0xb4, 0x96, 0x32, 0x31, 0xd2, 0xcc, 0xdd, 0xfb, 0x8b, 0x69, 0xbb, 0xa1, 0x61,
	0xb2, 0x4b, 0x7c, 0x35, 0xc4, 0xae, 0x56, 0x57, 0x5a, 0xdd, 0xad, 0xc5, 0xb4, 0x8d, 0x34, 0x46,
	0x6e, 0xab, 0xdd, 0xb8, 0xda, 0xcd, 0xda, 0x75, 0x84, 0xf4, 0xa5, 0xdd, 0x3d, 0x40, 0xdb, 0x03,
	0xb9, 0xab, 0xf6, 0x57, 0x86, 0xf2, 0x98, 0xa0, 0xc6, 0x52, 0xa2, 0xf8, 0x2e, 0x32, 0xfb, 0x85,
	0xb0, 0x6a, 0x18, 0xa1, 0xf5, 0x43, 0x88, 0x41, 0x80, 0x65, 0xec, 0x7f, 0x36, 0x91, 0xf9, 0xac,
	0xff, 0x1c, 0x7f, 0x31, 0xd0, 0xe6, 0x72, 0xb2, 0xf8, 0x51, 0xd5, 0x81, 0x2a, 0xae, 0x6a, 0x77,
	0xfe, 0x0f, 0xd4, 0x7e, 0xc8, 0xde, 0xa7, 0x9f, 0x7f, 0xbe, 0xd5, 0x5d, 0xe2, 0xd0, 0x8a, 0xff,
	0x61, 0xaa, 0x5f, 0x14, 0x59, 0x8c, 0xcf, 0x51, 0xe3, 0x58, 0xaf, 0x73, 0x28, 0xe3, 0xde, 0xad,
	0x12, 0xb8, 0x7e, 0x5b, 0xbb, 0x55, 0x85, 0x2a, 0x23, 0x21, 0xbb, 0x4a, 0xda, 0xc1, 0xad, 0x2a,
	0xe9, 0x71, 0x79, 0xd8, 0x8f, 0xa8, 0x31, 0x58, 0x12, 0xbe, 0x95, 0xd2, 0xae, 0x0c, 0xa8, 0xe2,
	0x0c, 0xa5, 0x36, 0xb9, 0x55, 0xbb, 0xfb, 0xe2, 0xc7, 0xcc, 0x31, 0x2e, 0x67, 0x8e, 0xf1, 0x7b,
	0xe6, 0x18, 0x5f, 0xe7, 0x4e, 0xed, 0x72, 0xee, 0xd4, 0x7e, 0xcd, 0x9d, 0xda, 0xbb, 0x27, 0x61,
	0x24, 0x46, 0xc5, 0xd0, 0x3b, 0xe5, 0x63, 0x7a, 0x14, 0x17, 0x93, 0x97, 0x47, 0xaf, 0x7b, 0x6c,
	0x98, 0x2b, 0xb6, 0x80, 0x9e, 0x8e, 0x58, 0x94, 0xac, 0x92, 0x0e, 0xd7, 0xd5, 0x67, 0x78, 0xf0,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x15, 0xb8, 0xc4, 0xf6, 0xf0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	PresignedURL(ctx context.Context, in *PresignedURLRequest, opts ...grpc.CallOption) (*PresignedURLResponse, error)
	GetMetaData(ctx context.Context, in *GetMedataRequest, opts ...grpc.CallOption) (*Metadata, error)
	SetMetaData(ctx context.Context, in *Metadata, opts ...grpc.CallOption) (*SetMetadataResponse, error)
}

type aPIClient struct {
	cc grpc1.ClientConn
}

func NewAPIClient(cc grpc1.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) PresignedURL(ctx context.Context, in *PresignedURLRequest, opts ...grpc.CallOption) (*PresignedURLResponse, error) {
	out := new(PresignedURLResponse)
	err := c.cc.Invoke(ctx, "/flux.indexer.media.API/PresignedURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetMetaData(ctx context.Context, in *GetMedataRequest, opts ...grpc.CallOption) (*Metadata, error) {
	out := new(Metadata)
	err := c.cc.Invoke(ctx, "/flux.indexer.media.API/GetMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) SetMetaData(ctx context.Context, in *Metadata, opts ...grpc.CallOption) (*SetMetadataResponse, error) {
	out := new(SetMetadataResponse)
	err := c.cc.Invoke(ctx, "/flux.indexer.media.API/SetMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	PresignedURL(context.Context, *PresignedURLRequest) (*PresignedURLResponse, error)
	GetMetaData(context.Context, *GetMedataRequest) (*Metadata, error)
	SetMetaData(context.Context, *Metadata) (*SetMetadataResponse, error)
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) PresignedURL(ctx context.Context, req *PresignedURLRequest) (*PresignedURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PresignedURL not implemented")
}
func (*UnimplementedAPIServer) GetMetaData(ctx context.Context, req *GetMedataRequest) (*Metadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetaData not implemented")
}
func (*UnimplementedAPIServer) SetMetaData(ctx context.Context, req *Metadata) (*SetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMetaData not implemented")
}

func RegisterAPIServer(s grpc1.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_PresignedURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PresignedURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).PresignedURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flux.indexer.media.API/PresignedURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).PresignedURL(ctx, req.(*PresignedURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flux.indexer.media.API/GetMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetMetaData(ctx, req.(*GetMedataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_SetMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Metadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).SetMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flux.indexer.media.API/SetMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).SetMetaData(ctx, req.(*Metadata))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flux.indexer.media.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PresignedURL",
			Handler:    _API_PresignedURL_Handler,
		},
		{
			MethodName: "GetMetaData",
			Handler:    _API_GetMetaData_Handler,
		},
		{
			MethodName: "SetMetaData",
			Handler:    _API_SetMetaData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flux/indexer/media/query.proto",
}

func (m *PresignedURLRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PresignedURLRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PresignedURLRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if m.Op != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Op))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PresignedURLResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PresignedURLResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PresignedURLResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetMedataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMedataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetMedataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SetMetadataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetMetadataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetMetadataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *PresignedURLRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Op != 0 {
		n += 1 + sovQuery(uint64(m.Op))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *PresignedURLResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *GetMedataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *SetMetadataResponse) Size() (n int) {
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
func (m *PresignedURLRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PresignedURLRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PresignedURLRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= S3Operation(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
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
			m.Key = string(dAtA[iNdEx:postIndex])
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
func (m *PresignedURLResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PresignedURLResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PresignedURLResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
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
			m.Url = string(dAtA[iNdEx:postIndex])
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
func (m *GetMedataRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetMedataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMedataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
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
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
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
			m.Url = string(dAtA[iNdEx:postIndex])
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
func (m *Metadata) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
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
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
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
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
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
			m.Url = string(dAtA[iNdEx:postIndex])
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
func (m *SetMetadataResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SetMetadataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetMetadataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
