// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prey.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FindPreyRequest struct {
	PreyName             string   `protobuf:"bytes,1,opt,name=prey_name,json=preyName,proto3" json:"prey_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPreyRequest) Reset()         { *m = FindPreyRequest{} }
func (m *FindPreyRequest) String() string { return proto.CompactTextString(m) }
func (*FindPreyRequest) ProtoMessage()    {}
func (*FindPreyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde12c6e752d6325, []int{0}
}

func (m *FindPreyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPreyRequest.Unmarshal(m, b)
}
func (m *FindPreyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPreyRequest.Marshal(b, m, deterministic)
}
func (m *FindPreyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPreyRequest.Merge(m, src)
}
func (m *FindPreyRequest) XXX_Size() int {
	return xxx_messageInfo_FindPreyRequest.Size(m)
}
func (m *FindPreyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPreyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindPreyRequest proto.InternalMessageInfo

func (m *FindPreyRequest) GetPreyName() string {
	if m != nil {
		return m.PreyName
	}
	return ""
}

type FindPreyResponse struct {
	PreyName             string   `protobuf:"bytes,1,opt,name=prey_name,json=preyName,proto3" json:"prey_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPreyResponse) Reset()         { *m = FindPreyResponse{} }
func (m *FindPreyResponse) String() string { return proto.CompactTextString(m) }
func (*FindPreyResponse) ProtoMessage()    {}
func (*FindPreyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde12c6e752d6325, []int{1}
}

func (m *FindPreyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPreyResponse.Unmarshal(m, b)
}
func (m *FindPreyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPreyResponse.Marshal(b, m, deterministic)
}
func (m *FindPreyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPreyResponse.Merge(m, src)
}
func (m *FindPreyResponse) XXX_Size() int {
	return xxx_messageInfo_FindPreyResponse.Size(m)
}
func (m *FindPreyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPreyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindPreyResponse proto.InternalMessageInfo

func (m *FindPreyResponse) GetPreyName() string {
	if m != nil {
		return m.PreyName
	}
	return ""
}

func init() {
	proto.RegisterType((*FindPreyRequest)(nil), "proto.FindPreyRequest")
	proto.RegisterType((*FindPreyResponse)(nil), "proto.FindPreyResponse")
}

func init() { proto.RegisterFile("prey.proto", fileDescriptor_fde12c6e752d6325) }

var fileDescriptor_fde12c6e752d6325 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x28, 0x4a, 0xad,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x7a, 0x5c, 0xfc, 0x6e, 0x99,
	0x79, 0x29, 0x01, 0x45, 0xa9, 0x95, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xd2, 0x5c,
	0x9c, 0x20, 0x75, 0xf1, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x1c,
	0x20, 0x01, 0xbf, 0xc4, 0xdc, 0x54, 0x25, 0x7d, 0x2e, 0x01, 0x84, 0xfa, 0xe2, 0x82, 0xfc, 0xbc,
	0xe2, 0x54, 0xbc, 0x1a, 0x8c, 0x7c, 0xb8, 0xb8, 0x7d, 0x32, 0xf3, 0xf3, 0x82, 0x53, 0x8b, 0xca,
	0x32, 0x93, 0x53, 0x85, 0x6c, 0xb9, 0x38, 0x60, 0xfa, 0x85, 0xc4, 0x20, 0x4e, 0xd1, 0x43, 0x73,
	0x80, 0x94, 0x38, 0x86, 0x38, 0xc4, 0x22, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x8c, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x04, 0x9a, 0xa2, 0xab, 0xca, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LionServiceClient is the client API for LionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LionServiceClient interface {
	FindPrey(ctx context.Context, in *FindPreyRequest, opts ...grpc.CallOption) (*FindPreyResponse, error)
}

type lionServiceClient struct {
	cc *grpc.ClientConn
}

func NewLionServiceClient(cc *grpc.ClientConn) LionServiceClient {
	return &lionServiceClient{cc}
}

func (c *lionServiceClient) FindPrey(ctx context.Context, in *FindPreyRequest, opts ...grpc.CallOption) (*FindPreyResponse, error) {
	out := new(FindPreyResponse)
	err := c.cc.Invoke(ctx, "/proto.LionService/FindPrey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LionServiceServer is the server API for LionService service.
type LionServiceServer interface {
	FindPrey(context.Context, *FindPreyRequest) (*FindPreyResponse, error)
}

// UnimplementedLionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLionServiceServer struct {
}

func (*UnimplementedLionServiceServer) FindPrey(ctx context.Context, req *FindPreyRequest) (*FindPreyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPrey not implemented")
}

func RegisterLionServiceServer(s *grpc.Server, srv LionServiceServer) {
	s.RegisterService(&_LionService_serviceDesc, srv)
}

func _LionService_FindPrey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPreyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LionServiceServer).FindPrey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LionService/FindPrey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LionServiceServer).FindPrey(ctx, req.(*FindPreyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LionService",
	HandlerType: (*LionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindPrey",
			Handler:    _LionService_FindPrey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prey.proto",
}