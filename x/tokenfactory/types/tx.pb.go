// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenfactory/tokenfactory/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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

type MsgCreateDenom struct {
	Owner              string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Denom              string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Description        string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Ticker             string `protobuf:"bytes,4,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Precision          int32  `protobuf:"varint,5,opt,name=precision,proto3" json:"precision,omitempty"`
	Url                string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	MaxSupply          int32  `protobuf:"varint,7,opt,name=maxSupply,proto3" json:"maxSupply,omitempty"`
	Supply             int32  `protobuf:"varint,8,opt,name=supply,proto3" json:"supply,omitempty"`
	CanChangeMaxSupply bool   `protobuf:"varint,9,opt,name=canChangeMaxSupply,proto3" json:"canChangeMaxSupply,omitempty"`
}

func (m *MsgCreateDenom) Reset()         { *m = MsgCreateDenom{} }
func (m *MsgCreateDenom) String() string { return proto.CompactTextString(m) }
func (*MsgCreateDenom) ProtoMessage()    {}
func (*MsgCreateDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_1be956d2c30629a7, []int{0}
}
func (m *MsgCreateDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateDenom.Merge(m, src)
}
func (m *MsgCreateDenom) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateDenom.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateDenom proto.InternalMessageInfo

func (m *MsgCreateDenom) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgCreateDenom) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *MsgCreateDenom) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgCreateDenom) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *MsgCreateDenom) GetPrecision() int32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *MsgCreateDenom) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *MsgCreateDenom) GetMaxSupply() int32 {
	if m != nil {
		return m.MaxSupply
	}
	return 0
}

func (m *MsgCreateDenom) GetSupply() int32 {
	if m != nil {
		return m.Supply
	}
	return 0
}

func (m *MsgCreateDenom) GetCanChangeMaxSupply() bool {
	if m != nil {
		return m.CanChangeMaxSupply
	}
	return false
}

type MsgCreateDenomResponse struct {
}

func (m *MsgCreateDenomResponse) Reset()         { *m = MsgCreateDenomResponse{} }
func (m *MsgCreateDenomResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateDenomResponse) ProtoMessage()    {}
func (*MsgCreateDenomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1be956d2c30629a7, []int{1}
}
func (m *MsgCreateDenomResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateDenomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateDenomResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateDenomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateDenomResponse.Merge(m, src)
}
func (m *MsgCreateDenomResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateDenomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateDenomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateDenomResponse proto.InternalMessageInfo

type MsgUpdateDenom struct {
	Owner              string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Denom              string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Description        string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Ticker             string `protobuf:"bytes,4,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Precision          int32  `protobuf:"varint,5,opt,name=precision,proto3" json:"precision,omitempty"`
	Url                string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	MaxSupply          int32  `protobuf:"varint,7,opt,name=maxSupply,proto3" json:"maxSupply,omitempty"`
	Supply             int32  `protobuf:"varint,8,opt,name=supply,proto3" json:"supply,omitempty"`
	CanChangeMaxSupply bool   `protobuf:"varint,9,opt,name=canChangeMaxSupply,proto3" json:"canChangeMaxSupply,omitempty"`
}

func (m *MsgUpdateDenom) Reset()         { *m = MsgUpdateDenom{} }
func (m *MsgUpdateDenom) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateDenom) ProtoMessage()    {}
func (*MsgUpdateDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_1be956d2c30629a7, []int{2}
}
func (m *MsgUpdateDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateDenom.Merge(m, src)
}
func (m *MsgUpdateDenom) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateDenom.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateDenom proto.InternalMessageInfo

func (m *MsgUpdateDenom) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgUpdateDenom) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *MsgUpdateDenom) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgUpdateDenom) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *MsgUpdateDenom) GetPrecision() int32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *MsgUpdateDenom) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *MsgUpdateDenom) GetMaxSupply() int32 {
	if m != nil {
		return m.MaxSupply
	}
	return 0
}

func (m *MsgUpdateDenom) GetSupply() int32 {
	if m != nil {
		return m.Supply
	}
	return 0
}

func (m *MsgUpdateDenom) GetCanChangeMaxSupply() bool {
	if m != nil {
		return m.CanChangeMaxSupply
	}
	return false
}

type MsgUpdateDenomResponse struct {
}

func (m *MsgUpdateDenomResponse) Reset()         { *m = MsgUpdateDenomResponse{} }
func (m *MsgUpdateDenomResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateDenomResponse) ProtoMessage()    {}
func (*MsgUpdateDenomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1be956d2c30629a7, []int{3}
}
func (m *MsgUpdateDenomResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateDenomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateDenomResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateDenomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateDenomResponse.Merge(m, src)
}
func (m *MsgUpdateDenomResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateDenomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateDenomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateDenomResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateDenom)(nil), "tokenfactory.tokenfactory.MsgCreateDenom")
	proto.RegisterType((*MsgCreateDenomResponse)(nil), "tokenfactory.tokenfactory.MsgCreateDenomResponse")
	proto.RegisterType((*MsgUpdateDenom)(nil), "tokenfactory.tokenfactory.MsgUpdateDenom")
	proto.RegisterType((*MsgUpdateDenomResponse)(nil), "tokenfactory.tokenfactory.MsgUpdateDenomResponse")
}

func init() {
	proto.RegisterFile("tokenfactory/tokenfactory/tx.proto", fileDescriptor_1be956d2c30629a7)
}

var fileDescriptor_1be956d2c30629a7 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x53, 0xc1, 0x4e, 0x32, 0x31,
	0x18, 0xa4, 0xf0, 0xc3, 0x0f, 0x25, 0x31, 0xa6, 0x31, 0xa4, 0x12, 0xd3, 0xac, 0x24, 0x26, 0x78,
	0x59, 0xa3, 0x1e, 0xbd, 0x89, 0x57, 0x2e, 0x6b, 0xbc, 0x78, 0x5b, 0x97, 0x4f, 0xdc, 0x00, 0x6d,
	0xd3, 0x96, 0x08, 0xef, 0xe0, 0xc1, 0xc7, 0xf2, 0xc8, 0xd1, 0x9b, 0x06, 0x5e, 0xc4, 0xb4, 0x0b,
	0xd8, 0x35, 0x10, 0x7c, 0x00, 0x6f, 0x3b, 0x33, 0xdf, 0xec, 0x64, 0x67, 0xbf, 0x0f, 0xb7, 0x8c,
	0x18, 0x00, 0x7f, 0x8c, 0x13, 0x23, 0xd4, 0xf4, 0x2c, 0x0f, 0x26, 0xa1, 0x54, 0xc2, 0x08, 0x72,
	0xe8, 0xd3, 0xa1, 0x0f, 0x9a, 0x27, 0xdb, 0xed, 0x3d, 0xe0, 0x62, 0x94, 0xbd, 0xa1, 0xf5, 0x52,
	0xc4, 0x7b, 0x5d, 0xdd, 0xef, 0x28, 0x88, 0x0d, 0xdc, 0x58, 0x81, 0x1c, 0xe0, 0xb2, 0x78, 0xe6,
	0xa0, 0x28, 0x0a, 0x50, 0xbb, 0x16, 0x65, 0xc0, 0xb2, 0xce, 0x47, 0x8b, 0x19, 0xeb, 0x00, 0x09,
	0x70, 0xbd, 0x07, 0x3a, 0x51, 0xa9, 0x34, 0xa9, 0xe0, 0xb4, 0xe4, 0x34, 0x9f, 0x22, 0x0d, 0x5c,
	0x31, 0x69, 0x32, 0x00, 0x45, 0xff, 0x39, 0x71, 0x89, 0xc8, 0x11, 0xae, 0x49, 0x05, 0x49, 0xaa,
	0xad, 0xaf, 0x1c, 0xa0, 0x76, 0x39, 0xfa, 0x26, 0xc8, 0x3e, 0x2e, 0x8d, 0xd5, 0x90, 0x56, 0x9c,
	0xc5, 0x3e, 0xda, 0xf9, 0x51, 0x3c, 0xb9, 0x1d, 0x4b, 0x39, 0x9c, 0xd2, 0xff, 0xd9, 0xfc, 0x9a,
	0xb0, 0x29, 0x3a, 0x93, 0xaa, 0x4e, 0x5a, 0x22, 0x12, 0x62, 0x92, 0xc4, 0xbc, 0xf3, 0x14, 0xf3,
	0x3e, 0x74, 0xd7, 0xf6, 0x5a, 0x80, 0xda, 0xd5, 0x68, 0x83, 0xd2, 0xa2, 0xb8, 0x91, 0x6f, 0x23,
	0x02, 0x2d, 0x05, 0xd7, 0xb0, 0x2a, 0xea, 0x4e, 0xf6, 0xfe, 0x8a, 0x5a, 0x17, 0xe5, 0xb5, 0xb1,
	0x2a, 0xea, 0xe2, 0x03, 0xe1, 0x52, 0x57, 0xf7, 0xc9, 0x00, 0xd7, 0xfd, 0xad, 0x3a, 0x0d, 0xb7,
	0xee, 0x6a, 0x98, 0xaf, 0xbc, 0x79, 0xfe, 0xeb, 0xd1, 0x55, 0xa8, 0x0d, 0xf3, 0xff, 0xcc, 0x8e,
	0x30, 0x6f, 0x74, 0x57, 0xd8, 0x86, 0x2f, 0xbc, 0xbe, 0x7a, 0x9b, 0x33, 0x34, 0x9b, 0x33, 0xf4,
	0x39, 0x67, 0xe8, 0x75, 0xc1, 0x0a, 0xb3, 0x05, 0x2b, 0xbc, 0x2f, 0x58, 0xe1, 0xfe, 0x38, 0x77,
	0x67, 0x93, 0x1f, 0x57, 0x3b, 0x95, 0xa0, 0x1f, 0x2a, 0xee, 0xee, 0x2e, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x5d, 0xc6, 0x6b, 0xa0, 0xdf, 0x03, 0x00, 0x00,
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
	CreateDenom(ctx context.Context, in *MsgCreateDenom, opts ...grpc.CallOption) (*MsgCreateDenomResponse, error)
	UpdateDenom(ctx context.Context, in *MsgUpdateDenom, opts ...grpc.CallOption) (*MsgUpdateDenomResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateDenom(ctx context.Context, in *MsgCreateDenom, opts ...grpc.CallOption) (*MsgCreateDenomResponse, error) {
	out := new(MsgCreateDenomResponse)
	err := c.cc.Invoke(ctx, "/tokenfactory.tokenfactory.Msg/CreateDenom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateDenom(ctx context.Context, in *MsgUpdateDenom, opts ...grpc.CallOption) (*MsgUpdateDenomResponse, error) {
	out := new(MsgUpdateDenomResponse)
	err := c.cc.Invoke(ctx, "/tokenfactory.tokenfactory.Msg/UpdateDenom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateDenom(context.Context, *MsgCreateDenom) (*MsgCreateDenomResponse, error)
	UpdateDenom(context.Context, *MsgUpdateDenom) (*MsgUpdateDenomResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateDenom(ctx context.Context, req *MsgCreateDenom) (*MsgCreateDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDenom not implemented")
}
func (*UnimplementedMsgServer) UpdateDenom(ctx context.Context, req *MsgUpdateDenom) (*MsgUpdateDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDenom not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateDenom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenfactory.tokenfactory.Msg/CreateDenom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateDenom(ctx, req.(*MsgCreateDenom))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateDenom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenfactory.tokenfactory.Msg/UpdateDenom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateDenom(ctx, req.(*MsgUpdateDenom))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tokenfactory.tokenfactory.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDenom",
			Handler:    _Msg_CreateDenom_Handler,
		},
		{
			MethodName: "UpdateDenom",
			Handler:    _Msg_UpdateDenom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tokenfactory/tokenfactory/tx.proto",
}

func (m *MsgCreateDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CanChangeMaxSupply {
		i--
		if m.CanChangeMaxSupply {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.Supply != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Supply))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxSupply != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MaxSupply))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x32
	}
	if m.Precision != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Precision))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Ticker) > 0 {
		i -= len(m.Ticker)
		copy(dAtA[i:], m.Ticker)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Ticker)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateDenomResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateDenomResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateDenomResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CanChangeMaxSupply {
		i--
		if m.CanChangeMaxSupply {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.Supply != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Supply))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxSupply != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MaxSupply))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x32
	}
	if m.Precision != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Precision))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Ticker) > 0 {
		i -= len(m.Ticker)
		copy(dAtA[i:], m.Ticker)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Ticker)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateDenomResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateDenomResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateDenomResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Ticker)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Precision != 0 {
		n += 1 + sovTx(uint64(m.Precision))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.MaxSupply != 0 {
		n += 1 + sovTx(uint64(m.MaxSupply))
	}
	if m.Supply != 0 {
		n += 1 + sovTx(uint64(m.Supply))
	}
	if m.CanChangeMaxSupply {
		n += 2
	}
	return n
}

func (m *MsgCreateDenomResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Ticker)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Precision != 0 {
		n += 1 + sovTx(uint64(m.Precision))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.MaxSupply != 0 {
		n += 1 + sovTx(uint64(m.MaxSupply))
	}
	if m.Supply != 0 {
		n += 1 + sovTx(uint64(m.Supply))
	}
	if m.CanChangeMaxSupply {
		n += 2
	}
	return n
}

func (m *MsgUpdateDenomResponse) Size() (n int) {
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
func (m *MsgCreateDenom) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
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
			m.Ticker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
			}
			m.Precision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Precision |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
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
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			m.MaxSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSupply |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			m.Supply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Supply |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanChangeMaxSupply", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.CanChangeMaxSupply = bool(v != 0)
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
func (m *MsgCreateDenomResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateDenomResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateDenomResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateDenom) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
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
			m.Ticker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
			}
			m.Precision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Precision |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
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
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			m.MaxSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSupply |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			m.Supply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Supply |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanChangeMaxSupply", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.CanChangeMaxSupply = bool(v != 0)
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
func (m *MsgUpdateDenomResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateDenomResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateDenomResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
