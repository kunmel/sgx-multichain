// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multichain.proto

package multichain_proto // import "sgx-multichain/multichain_proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NewFabRequest struct {
	CCName               string   `protobuf:"bytes,1,opt,name=CCName,proto3" json:"CCName,omitempty"`
	FcnName              string   `protobuf:"bytes,2,opt,name=FcnName,proto3" json:"FcnName,omitempty"`
	CCArg                []string `protobuf:"bytes,3,rep,name=CCArg,proto3" json:"CCArg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewFabRequest) Reset()         { *m = NewFabRequest{} }
func (m *NewFabRequest) String() string { return proto.CompactTextString(m) }
func (*NewFabRequest) ProtoMessage()    {}
func (*NewFabRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_multichain_e73f529096addbdf, []int{0}
}
func (m *NewFabRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewFabRequest.Unmarshal(m, b)
}
func (m *NewFabRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewFabRequest.Marshal(b, m, deterministic)
}
func (dst *NewFabRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewFabRequest.Merge(dst, src)
}
func (m *NewFabRequest) XXX_Size() int {
	return xxx_messageInfo_NewFabRequest.Size(m)
}
func (m *NewFabRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewFabRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewFabRequest proto.InternalMessageInfo

func (m *NewFabRequest) GetCCName() string {
	if m != nil {
		return m.CCName
	}
	return ""
}

func (m *NewFabRequest) GetFcnName() string {
	if m != nil {
		return m.FcnName
	}
	return ""
}

func (m *NewFabRequest) GetCCArg() []string {
	if m != nil {
		return m.CCArg
	}
	return nil
}

type NewEthRequest struct {
	KeystorePath         string   `protobuf:"bytes,1,opt,name=KeystorePath,proto3" json:"KeystorePath,omitempty"`
	ChainID              string   `protobuf:"bytes,2,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Pass                 string   `protobuf:"bytes,3,opt,name=Pass,proto3" json:"Pass,omitempty"`
	ContractAddress      string   `protobuf:"bytes,4,opt,name=ContractAddress,proto3" json:"ContractAddress,omitempty"`
	FcnName              string   `protobuf:"bytes,5,opt,name=FcnName,proto3" json:"FcnName,omitempty"`
	ContractArg          string   `protobuf:"bytes,6,opt,name=ContractArg,proto3" json:"ContractArg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewEthRequest) Reset()         { *m = NewEthRequest{} }
func (m *NewEthRequest) String() string { return proto.CompactTextString(m) }
func (*NewEthRequest) ProtoMessage()    {}
func (*NewEthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_multichain_e73f529096addbdf, []int{1}
}
func (m *NewEthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewEthRequest.Unmarshal(m, b)
}
func (m *NewEthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewEthRequest.Marshal(b, m, deterministic)
}
func (dst *NewEthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewEthRequest.Merge(dst, src)
}
func (m *NewEthRequest) XXX_Size() int {
	return xxx_messageInfo_NewEthRequest.Size(m)
}
func (m *NewEthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewEthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewEthRequest proto.InternalMessageInfo

func (m *NewEthRequest) GetKeystorePath() string {
	if m != nil {
		return m.KeystorePath
	}
	return ""
}

func (m *NewEthRequest) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *NewEthRequest) GetPass() string {
	if m != nil {
		return m.Pass
	}
	return ""
}

func (m *NewEthRequest) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *NewEthRequest) GetFcnName() string {
	if m != nil {
		return m.FcnName
	}
	return ""
}

func (m *NewEthRequest) GetContractArg() string {
	if m != nil {
		return m.ContractArg
	}
	return ""
}

type CommitResponse struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	TxID                 string   `protobuf:"bytes,2,opt,name=txID,proto3" json:"txID,omitempty"`
	BlockID              string   `protobuf:"bytes,3,opt,name=blockID,proto3" json:"blockID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitResponse) Reset()         { *m = CommitResponse{} }
func (m *CommitResponse) String() string { return proto.CompactTextString(m) }
func (*CommitResponse) ProtoMessage()    {}
func (*CommitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_multichain_e73f529096addbdf, []int{2}
}
func (m *CommitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitResponse.Unmarshal(m, b)
}
func (m *CommitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitResponse.Marshal(b, m, deterministic)
}
func (dst *CommitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitResponse.Merge(dst, src)
}
func (m *CommitResponse) XXX_Size() int {
	return xxx_messageInfo_CommitResponse.Size(m)
}
func (m *CommitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommitResponse proto.InternalMessageInfo

func (m *CommitResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *CommitResponse) GetTxID() string {
	if m != nil {
		return m.TxID
	}
	return ""
}

func (m *CommitResponse) GetBlockID() string {
	if m != nil {
		return m.BlockID
	}
	return ""
}

func init() {
	proto.RegisterType((*NewFabRequest)(nil), "multichain_proto.NewFabRequest")
	proto.RegisterType((*NewEthRequest)(nil), "multichain_proto.NewEthRequest")
	proto.RegisterType((*CommitResponse)(nil), "multichain_proto.CommitResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MultichainServiceClient is the client API for MultichainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MultichainServiceClient interface {
	NewFabTx(ctx context.Context, in *NewFabRequest, opts ...grpc.CallOption) (MultichainService_NewFabTxClient, error)
	NewEthTx(ctx context.Context, in *NewEthRequest, opts ...grpc.CallOption) (MultichainService_NewEthTxClient, error)
}

type multichainServiceClient struct {
	cc *grpc.ClientConn
}

func NewMultichainServiceClient(cc *grpc.ClientConn) MultichainServiceClient {
	return &multichainServiceClient{cc}
}

func (c *multichainServiceClient) NewFabTx(ctx context.Context, in *NewFabRequest, opts ...grpc.CallOption) (MultichainService_NewFabTxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MultichainService_serviceDesc.Streams[0], "/multichain_proto.multichainService/NewFabTx", opts...)
	if err != nil {
		return nil, err
	}
	x := &multichainServiceNewFabTxClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MultichainService_NewFabTxClient interface {
	Recv() (*CommitResponse, error)
	grpc.ClientStream
}

type multichainServiceNewFabTxClient struct {
	grpc.ClientStream
}

func (x *multichainServiceNewFabTxClient) Recv() (*CommitResponse, error) {
	m := new(CommitResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *multichainServiceClient) NewEthTx(ctx context.Context, in *NewEthRequest, opts ...grpc.CallOption) (MultichainService_NewEthTxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MultichainService_serviceDesc.Streams[1], "/multichain_proto.multichainService/NewEthTx", opts...)
	if err != nil {
		return nil, err
	}
	x := &multichainServiceNewEthTxClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MultichainService_NewEthTxClient interface {
	Recv() (*CommitResponse, error)
	grpc.ClientStream
}

type multichainServiceNewEthTxClient struct {
	grpc.ClientStream
}

func (x *multichainServiceNewEthTxClient) Recv() (*CommitResponse, error) {
	m := new(CommitResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MultichainServiceServer is the server API for MultichainService service.
type MultichainServiceServer interface {
	NewFabTx(*NewFabRequest, MultichainService_NewFabTxServer) error
	NewEthTx(*NewEthRequest, MultichainService_NewEthTxServer) error
}

func RegisterMultichainServiceServer(s *grpc.Server, srv MultichainServiceServer) {
	s.RegisterService(&_MultichainService_serviceDesc, srv)
}

func _MultichainService_NewFabTx_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NewFabRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MultichainServiceServer).NewFabTx(m, &multichainServiceNewFabTxServer{stream})
}

type MultichainService_NewFabTxServer interface {
	Send(*CommitResponse) error
	grpc.ServerStream
}

type multichainServiceNewFabTxServer struct {
	grpc.ServerStream
}

func (x *multichainServiceNewFabTxServer) Send(m *CommitResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MultichainService_NewEthTx_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NewEthRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MultichainServiceServer).NewEthTx(m, &multichainServiceNewEthTxServer{stream})
}

type MultichainService_NewEthTxServer interface {
	Send(*CommitResponse) error
	grpc.ServerStream
}

type multichainServiceNewEthTxServer struct {
	grpc.ServerStream
}

func (x *multichainServiceNewEthTxServer) Send(m *CommitResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MultichainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "multichain_proto.multichainService",
	HandlerType: (*MultichainServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NewFabTx",
			Handler:       _MultichainService_NewFabTx_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NewEthTx",
			Handler:       _MultichainService_NewEthTx_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "multichain.proto",
}

func init() { proto.RegisterFile("multichain.proto", fileDescriptor_multichain_e73f529096addbdf) }

var fileDescriptor_multichain_e73f529096addbdf = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xdd, 0x4e, 0xc2, 0x30,
	0x18, 0x75, 0xf2, 0xa3, 0x7c, 0xfe, 0x37, 0xc6, 0x34, 0x5e, 0xc8, 0xdc, 0x15, 0x37, 0xa2, 0xd1,
	0x27, 0xc0, 0x02, 0x09, 0x31, 0x21, 0x38, 0x4c, 0x4c, 0xbc, 0xd0, 0x94, 0xd1, 0xb0, 0x45, 0xb6,
	0x62, 0xfb, 0xa1, 0xf8, 0x68, 0x3e, 0x81, 0xaf, 0x65, 0x56, 0x36, 0xd8, 0x66, 0x4c, 0xbc, 0x5a,
	0xcf, 0x39, 0xdf, 0x4e, 0xcf, 0xd7, 0x03, 0x87, 0xe1, 0x7c, 0x8a, 0x81, 0xe7, 0xf3, 0x20, 0x6a,
	0xce, 0x94, 0x44, 0x49, 0x32, 0xcc, 0x8b, 0x61, 0x9c, 0x47, 0xd8, 0xeb, 0x8b, 0x8f, 0x2e, 0x1f,
	0xb9, 0xe2, 0x6d, 0x2e, 0x34, 0x92, 0x13, 0xa8, 0x32, 0xd6, 0xe7, 0xa1, 0xa0, 0x96, 0x6d, 0x35,
	0x6a, 0x6e, 0x82, 0x08, 0x85, 0xad, 0xae, 0x17, 0x19, 0x61, 0xd3, 0x08, 0x29, 0x24, 0xc7, 0x50,
	0x61, 0xac, 0xa5, 0x26, 0xb4, 0x64, 0x97, 0x1a, 0x35, 0x77, 0x09, 0x9c, 0x6f, 0xcb, 0x38, 0x77,
	0xd0, 0x4f, 0x9d, 0x1d, 0xd8, 0xbd, 0x13, 0x9f, 0x1a, 0xa5, 0x12, 0x03, 0x8e, 0x7e, 0xe2, 0x9f,
	0xe3, 0xe2, 0x5b, 0x58, 0x9c, 0xae, 0xd7, 0x4e, 0x6f, 0x49, 0x20, 0x21, 0x50, 0x1e, 0x70, 0xad,
	0x69, 0xc9, 0xd0, 0xe6, 0x4c, 0x1a, 0x70, 0xc0, 0x64, 0x84, 0x8a, 0x7b, 0xd8, 0x1a, 0x8f, 0x95,
	0xd0, 0x9a, 0x96, 0x8d, 0x5c, 0xa4, 0xb3, 0xe9, 0x2b, 0xf9, 0xf4, 0x36, 0xec, 0xac, 0x86, 0xd5,
	0x84, 0x56, 0x8d, 0x9a, 0xa5, 0x9c, 0x67, 0xd8, 0x67, 0x32, 0x0c, 0x03, 0x74, 0x85, 0x9e, 0xc9,
	0x48, 0x0b, 0x72, 0x06, 0x30, 0x44, 0x8e, 0x73, 0xcd, 0xe4, 0x78, 0xf9, 0x4e, 0x15, 0x37, 0xc3,
	0xc4, 0x59, 0x71, 0xb1, 0x5a, 0xc1, 0x9c, 0xe3, 0x04, 0xa3, 0xa9, 0xf4, 0x5e, 0x7b, 0xed, 0x64,
	0x85, 0x14, 0x5e, 0x7f, 0x59, 0x70, 0xb4, 0xee, 0x65, 0x28, 0xd4, 0x7b, 0xe0, 0x09, 0x72, 0x0f,
	0xdb, 0xcb, 0x62, 0x1e, 0x16, 0xa4, 0xde, 0x2c, 0xf6, 0xd6, 0xcc, 0x95, 0x76, 0x6a, 0xff, 0x1e,
	0xc8, 0x47, 0x76, 0x36, 0xae, 0xac, 0xc4, 0xb2, 0x83, 0xfe, 0x9f, 0x96, 0xeb, 0xb6, 0xfe, 0x67,
	0x79, 0x7b, 0xfe, 0x54, 0xd7, 0x93, 0xc5, 0xc5, 0x7a, 0xf4, 0xb2, 0xf8, 0xd7, 0xa8, 0x6a, 0x3e,
	0x37, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xe1, 0x1e, 0x78, 0x8e, 0x02, 0x00, 0x00,
}
