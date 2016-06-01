// Code generated by protoc-gen-go.
// source: op.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	op.proto

It has these top-level messages:
	OpRequest
	OpReply
*/
package pb

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

type OpRequest struct {
	Op     string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Bucket string `protobuf:"bytes,4,opt,name=bucket" json:"bucket,omitempty"`
}

func (m *OpRequest) Reset() {
	*m = OpRequest{}
}
func (m *OpRequest) String() string {
	return proto.CompactTextString(m)
}
func (*OpRequest) ProtoMessage() {}
func (*OpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

type OpReply struct {
	Status  int32  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Value   []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	ErrCode int32  `protobuf:"varint,4,opt,name=errCode" json:"errCode,omitempty"`
}

func (m *OpReply) Reset() {
	*m = OpReply{}
}
func (m *OpReply) String() string {
	return proto.CompactTextString(m)
}
func (*OpReply) ProtoMessage() {}
func (*OpReply) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1}
}

func init() {
	proto.RegisterType((*OpRequest)(nil), "pb.OpRequest")
	proto.RegisterType((*OpReply)(nil), "pb.OpReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for RiotGossip service

type RiotGossipClient interface {
	OpRPC(ctx context.Context, in *OpRequest, opts ...grpc.CallOption) (*OpReply, error)
}

type riotGossipClient struct {
	cc *grpc.ClientConn
}

func NewRiotGossipClient(cc *grpc.ClientConn) RiotGossipClient {
	return &riotGossipClient{cc}
}

func (c *riotGossipClient) OpRPC(ctx context.Context, in *OpRequest, opts ...grpc.CallOption) (*OpReply, error) {
	out := new(OpReply)
	err := grpc.Invoke(ctx, "/pb.RiotGossip/OpRPC", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RiotGossip service

type RiotGossipServer interface {
	OpRPC(context.Context, *OpRequest) (*OpReply, error)
}

func RegisterRiotGossipServer(s *grpc.Server, srv RiotGossipServer) {
	s.RegisterService(&_RiotGossip_serviceDesc, srv)
}

func _RiotGossip_OpRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotGossipServer).OpRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RiotGossip/OpRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotGossipServer).OpRPC(ctx, req.(*OpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RiotGossip_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RiotGossip",
	HandlerType: (*RiotGossipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpRPC",
			Handler:    _RiotGossip_OpRPC_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x2f, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x8a, 0xe6, 0xe2, 0xf4, 0x2f, 0x08, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x2f, 0x90, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0xb2, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x98, 0xc0, 0x02, 0x20,
	0xa6, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0x33, 0x50, 0x8c, 0x27, 0x08,
	0xc2, 0x11, 0x12, 0xe3, 0x62, 0x4b, 0x2a, 0x4d, 0xce, 0x4e, 0x2d, 0x91, 0x60, 0x01, 0x2b, 0x85,
	0xf2, 0x94, 0x92, 0xb9, 0xd8, 0x41, 0x86, 0x17, 0xe4, 0x54, 0x82, 0x94, 0x14, 0x97, 0x24, 0x96,
	0x94, 0x16, 0x83, 0x8d, 0x67, 0x0d, 0x82, 0xf2, 0x40, 0x56, 0xe4, 0x16, 0xa7, 0xc3, 0xac, 0x00,
	0x32, 0x71, 0x58, 0x21, 0xc1, 0xc5, 0x9e, 0x5a, 0x54, 0xe4, 0x9c, 0x9f, 0x92, 0x0a, 0xb6, 0x83,
	0x35, 0x08, 0xc6, 0x35, 0x32, 0xe6, 0xe2, 0x0a, 0xca, 0xcc, 0x2f, 0x71, 0xcf, 0x2f, 0x2e, 0xce,
	0x2c, 0x10, 0x52, 0xe5, 0x62, 0x05, 0x5a, 0x19, 0xe0, 0x2c, 0xc4, 0xab, 0x57, 0x90, 0xa4, 0x07,
	0xf7, 0x9a, 0x14, 0x37, 0x8c, 0x0b, 0x74, 0x8c, 0x12, 0x43, 0x12, 0x1b, 0x38, 0x04, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x32, 0x36, 0xca, 0x53, 0x0d, 0x01, 0x00, 0x00,
}
