// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vip_resolver.proto

package api

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

type GetVIPByNameRequest struct {
	Fqdn                 string   `protobuf:"bytes,1,opt,name=fqdn" json:"fqdn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVIPByNameRequest) Reset()         { *m = GetVIPByNameRequest{} }
func (m *GetVIPByNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetVIPByNameRequest) ProtoMessage()    {}
func (*GetVIPByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_vip_resolver_8c48ba218f39683a, []int{0}
}
func (m *GetVIPByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVIPByNameRequest.Unmarshal(m, b)
}
func (m *GetVIPByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVIPByNameRequest.Marshal(b, m, deterministic)
}
func (dst *GetVIPByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVIPByNameRequest.Merge(dst, src)
}
func (m *GetVIPByNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetVIPByNameRequest.Size(m)
}
func (m *GetVIPByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVIPByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVIPByNameRequest proto.InternalMessageInfo

func (m *GetVIPByNameRequest) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

type GetVIPByNameResponse struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVIPByNameResponse) Reset()         { *m = GetVIPByNameResponse{} }
func (m *GetVIPByNameResponse) String() string { return proto.CompactTextString(m) }
func (*GetVIPByNameResponse) ProtoMessage()    {}
func (*GetVIPByNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_vip_resolver_8c48ba218f39683a, []int{1}
}
func (m *GetVIPByNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVIPByNameResponse.Unmarshal(m, b)
}
func (m *GetVIPByNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVIPByNameResponse.Marshal(b, m, deterministic)
}
func (dst *GetVIPByNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVIPByNameResponse.Merge(dst, src)
}
func (m *GetVIPByNameResponse) XXX_Size() int {
	return xxx_messageInfo_GetVIPByNameResponse.Size(m)
}
func (m *GetVIPByNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVIPByNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVIPByNameResponse proto.InternalMessageInfo

func (m *GetVIPByNameResponse) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func init() {
	proto.RegisterType((*GetVIPByNameRequest)(nil), "api.GetVIPByNameRequest")
	proto.RegisterType((*GetVIPByNameResponse)(nil), "api.GetVIPByNameResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VIPResolverCopilotClient is the client API for VIPResolverCopilot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VIPResolverCopilotClient interface {
	GetVIPByName(ctx context.Context, in *GetVIPByNameRequest, opts ...grpc.CallOption) (*GetVIPByNameResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type vIPResolverCopilotClient struct {
	cc *grpc.ClientConn
}

func NewVIPResolverCopilotClient(cc *grpc.ClientConn) VIPResolverCopilotClient {
	return &vIPResolverCopilotClient{cc}
}

func (c *vIPResolverCopilotClient) GetVIPByName(ctx context.Context, in *GetVIPByNameRequest, opts ...grpc.CallOption) (*GetVIPByNameResponse, error) {
	out := new(GetVIPByNameResponse)
	err := c.cc.Invoke(ctx, "/api.VIPResolverCopilot/GetVIPByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vIPResolverCopilotClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/api.VIPResolverCopilot/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VIPResolverCopilot service

type VIPResolverCopilotServer interface {
	GetVIPByName(context.Context, *GetVIPByNameRequest) (*GetVIPByNameResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
}

func RegisterVIPResolverCopilotServer(s *grpc.Server, srv VIPResolverCopilotServer) {
	s.RegisterService(&_VIPResolverCopilot_serviceDesc, srv)
}

func _VIPResolverCopilot_GetVIPByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVIPByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VIPResolverCopilotServer).GetVIPByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.VIPResolverCopilot/GetVIPByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VIPResolverCopilotServer).GetVIPByName(ctx, req.(*GetVIPByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VIPResolverCopilot_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VIPResolverCopilotServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.VIPResolverCopilot/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VIPResolverCopilotServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VIPResolverCopilot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.VIPResolverCopilot",
	HandlerType: (*VIPResolverCopilotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVIPByName",
			Handler:    _VIPResolverCopilot_GetVIPByName_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _VIPResolverCopilot_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vip_resolver.proto",
}

func init() { proto.RegisterFile("vip_resolver.proto", fileDescriptor_vip_resolver_8c48ba218f39683a) }

var fileDescriptor_vip_resolver_8c48ba218f39683a = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xcb, 0x2c, 0x88,
	0x2f, 0x4a, 0x2d, 0xce, 0xcf, 0x29, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4e, 0x2c, 0xc8, 0x94, 0xe2, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x08, 0x29, 0x69, 0x72,
	0x09, 0xbb, 0xa7, 0x96, 0x84, 0x79, 0x06, 0x38, 0x55, 0xfa, 0x25, 0xe6, 0xa6, 0x06, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xa4, 0x15, 0xa6, 0xe4, 0x49, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x6a, 0x5c, 0x22, 0xa8, 0x4a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a,
	0x53, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x0b, 0xa0, 0x2a, 0x99, 0x32, 0x0b, 0x8c, 0x26, 0x30, 0x72,
	0x09, 0x85, 0x79, 0x06, 0x04, 0x41, 0xed, 0x76, 0xce, 0x2f, 0xc8, 0xcc, 0xc9, 0x2f, 0x11, 0x72,
	0xe5, 0xe2, 0x41, 0xd6, 0x2e, 0x24, 0xa1, 0x97, 0x58, 0x90, 0xa9, 0x87, 0xc5, 0x72, 0x29, 0x49,
	0x2c, 0x32, 0x10, 0xbb, 0x94, 0x18, 0x84, 0x8c, 0xb9, 0xd8, 0x3c, 0x52, 0x13, 0x73, 0x4a, 0x32,
	0x84, 0x84, 0xc0, 0xca, 0x20, 0x1c, 0x98, 0x56, 0x61, 0x14, 0x31, 0x98, 0xa6, 0x24, 0x36, 0xb0,
	0x67, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0x3f, 0x53, 0x98, 0x15, 0x01, 0x00, 0x00,
}
