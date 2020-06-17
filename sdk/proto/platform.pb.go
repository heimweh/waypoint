// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Deploy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deploy) Reset()         { *m = Deploy{} }
func (m *Deploy) String() string { return proto.CompactTextString(m) }
func (*Deploy) ProtoMessage()    {}
func (*Deploy) Descriptor() ([]byte, []int) {
	return fileDescriptor_918f3d50bfb447e4, []int{0}
}

func (m *Deploy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deploy.Unmarshal(m, b)
}
func (m *Deploy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deploy.Marshal(b, m, deterministic)
}
func (m *Deploy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deploy.Merge(m, src)
}
func (m *Deploy) XXX_Size() int {
	return xxx_messageInfo_Deploy.Size(m)
}
func (m *Deploy) XXX_DiscardUnknown() {
	xxx_messageInfo_Deploy.DiscardUnknown(m)
}

var xxx_messageInfo_Deploy proto.InternalMessageInfo

type Deploy_Resp struct {
	// result is the resulting opaque data type
	Result               *any.Any `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deploy_Resp) Reset()         { *m = Deploy_Resp{} }
func (m *Deploy_Resp) String() string { return proto.CompactTextString(m) }
func (*Deploy_Resp) ProtoMessage()    {}
func (*Deploy_Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_918f3d50bfb447e4, []int{0, 0}
}

func (m *Deploy_Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deploy_Resp.Unmarshal(m, b)
}
func (m *Deploy_Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deploy_Resp.Marshal(b, m, deterministic)
}
func (m *Deploy_Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deploy_Resp.Merge(m, src)
}
func (m *Deploy_Resp) XXX_Size() int {
	return xxx_messageInfo_Deploy_Resp.Size(m)
}
func (m *Deploy_Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Deploy_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Deploy_Resp proto.InternalMessageInfo

func (m *Deploy_Resp) GetResult() *any.Any {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Deploy)(nil), "proto.Deploy")
	proto.RegisterType((*Deploy_Resp)(nil), "proto.Deploy.Resp")
}

func init() {
	proto.RegisterFile("platform.proto", fileDescriptor_918f3d50bfb447e4)
}

var fileDescriptor_918f3d50bfb447e4 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4b, 0x4b, 0x2b, 0x31,
	0x14, 0xc7, 0x29, 0xf4, 0x79, 0x3a, 0xf7, 0x41, 0xe8, 0xbd, 0xd4, 0x11, 0x44, 0x5c, 0x29, 0x4a,
	0x8a, 0xb5, 0xe8, 0xc2, 0x52, 0xac, 0xad, 0x42, 0xc1, 0x85, 0xb4, 0xe0, 0xc2, 0xdd, 0x58, 0x4f,
	0xc7, 0x42, 0x66, 0x12, 0xf3, 0x58, 0xcc, 0xe7, 0xf2, 0x0b, 0x4a, 0x93, 0xb4, 0x68, 0xcb, 0x40,
	0xeb, 0x2a, 0x24, 0xe7, 0xf7, 0x3b, 0xf9, 0x4f, 0xce, 0xc0, 0x6f, 0xc1, 0x22, 0x3d, 0xe3, 0x32,
	0xa1, 0x42, 0x72, 0xcd, 0x49, 0xc9, 0x2e, 0xe1, 0x5e, 0xcc, 0x79, 0xcc, 0xb0, 0x65, 0x77, 0x2f,
	0x66, 0xd6, 0x8a, 0xd2, 0xcc, 0x11, 0xe1, 0xfe, 0x7a, 0x09, 0x13, 0xa1, 0x97, 0xc5, 0x40, 0x30,
	0x13, 0xcf, 0x53, 0xb7, 0x3b, 0xea, 0x41, 0x79, 0x88, 0x82, 0xf1, 0x2c, 0xec, 0x40, 0x71, 0x8c,
	0x4a, 0x90, 0x33, 0x28, 0x4b, 0x54, 0x86, 0xe9, 0x66, 0xe1, 0xb0, 0x70, 0x5c, 0x6f, 0x37, 0xa8,
	0xeb, 0x46, 0x97, 0xdd, 0x68, 0x3f, 0xcd, 0xc6, 0x9e, 0x69, 0x7f, 0x94, 0xa0, 0xfa, 0xe8, 0xf3,
	0x91, 0x0e, 0x14, 0xfb, 0x46, 0xbf, 0x91, 0x86, 0x63, 0xe9, 0xbd, 0x49, 0xa7, 0x13, 0x81, 0x53,
	0xda, 0x97, 0xb1, 0x0a, 0xff, 0x6f, 0x34, 0xba, 0x5b, 0xc4, 0x22, 0x27, 0x50, 0x5d, 0x58, 0x0b,
	0x90, 0x04, 0xde, 0xb4, 0x95, 0xf0, 0xcf, 0x5a, 0x1f, 0xd2, 0x85, 0xe0, 0x29, 0x62, 0xf3, 0xd7,
	0x48, 0xe3, 0x0f, 0x2e, 0x3a, 0x87, 0xbf, 0x5f, 0xed, 0x6d, 0x2e, 0xbc, 0x81, 0x60, 0xc0, 0xd3,
	0xd9, 0x3c, 0x9e, 0x68, 0x69, 0xa6, 0x9a, 0xe4, 0xb4, 0x0e, 0x9b, 0x5e, 0x74, 0x30, 0x75, 0xb4,
	0x7d, 0xce, 0x01, 0xd4, 0xdc, 0xa1, 0x91, 0x48, 0x0e, 0xbe, 0x63, 0xab, 0xc2, 0x18, 0xdf, 0x0d,
	0x2a, 0x9d, 0x9b, 0xfc, 0x14, 0xc0, 0x4d, 0x69, 0x9b, 0xcc, 0xed, 0xe5, 0x48, 0x73, 0x9e, 0x87,
	0xf8, 0x53, 0x07, 0x51, 0x9b, 0xb2, 0x0b, 0xf5, 0x91, 0x1a, 0xa2, 0xd2, 0x92, 0x67, 0x28, 0x73,
	0x3f, 0xf3, 0x9f, 0x57, 0x47, 0x89, 0x60, 0x98, 0x60, 0xaa, 0x95, 0xb5, 0x2f, 0xa1, 0xee, 0x5d,
	0x1b, 0x20, 0xcf, 0xde, 0x48, 0x7a, 0x05, 0x15, 0xef, 0xed, 0x38, 0xc9, 0x1e, 0xfc, 0x1a, 0xa9,
	0x07, 0x1e, 0xaf, 0xfe, 0xbc, 0xdd, 0x02, 0xdf, 0xd6, 0x9e, 0x2b, 0xf4, 0xda, 0xa1, 0x65, 0xbb,
	0x5c, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0x52, 0x61, 0x66, 0x29, 0x66, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlatformClient is the client API for Platform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlatformClient interface {
	Auth(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error)
	AuthSpec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	ValidateAuth(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error)
	ValidateAuthSpec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	ConfigStruct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Config_StructResp, error)
	Configure(ctx context.Context, in *Config_ConfigureRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeploySpec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	Deploy(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*Deploy_Resp, error)
	// component.Destroyer optional implementation
	IsDestroyer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementsResp, error)
	DestroySpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	Destroy(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error)
	// IsLogPlatform returns true if this platform also implements LogPlatform.
	IsLogPlatform(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementsResp, error)
}

type platformClient struct {
	cc grpc.ClientConnInterface
}

func NewPlatformClient(cc grpc.ClientConnInterface) PlatformClient {
	return &platformClient{cc}
}

func (c *platformClient) Auth(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Platform/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) AuthSpec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/proto.Platform/AuthSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) ValidateAuth(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Platform/ValidateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) ValidateAuthSpec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/proto.Platform/ValidateAuthSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) ConfigStruct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Config_StructResp, error) {
	out := new(Config_StructResp)
	err := c.cc.Invoke(ctx, "/proto.Platform/ConfigStruct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) Configure(ctx context.Context, in *Config_ConfigureRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Platform/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) DeploySpec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/proto.Platform/DeploySpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) Deploy(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*Deploy_Resp, error) {
	out := new(Deploy_Resp)
	err := c.cc.Invoke(ctx, "/proto.Platform/Deploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) IsDestroyer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementsResp, error) {
	out := new(ImplementsResp)
	err := c.cc.Invoke(ctx, "/proto.Platform/IsDestroyer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) DestroySpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/proto.Platform/DestroySpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) Destroy(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Platform/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) IsLogPlatform(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementsResp, error) {
	out := new(ImplementsResp)
	err := c.cc.Invoke(ctx, "/proto.Platform/IsLogPlatform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformServer is the server API for Platform service.
type PlatformServer interface {
	Auth(context.Context, *FuncSpec_Args) (*empty.Empty, error)
	AuthSpec(context.Context, *Empty) (*FuncSpec, error)
	ValidateAuth(context.Context, *FuncSpec_Args) (*empty.Empty, error)
	ValidateAuthSpec(context.Context, *Empty) (*FuncSpec, error)
	ConfigStruct(context.Context, *empty.Empty) (*Config_StructResp, error)
	Configure(context.Context, *Config_ConfigureRequest) (*empty.Empty, error)
	DeploySpec(context.Context, *Empty) (*FuncSpec, error)
	Deploy(context.Context, *FuncSpec_Args) (*Deploy_Resp, error)
	// component.Destroyer optional implementation
	IsDestroyer(context.Context, *empty.Empty) (*ImplementsResp, error)
	DestroySpec(context.Context, *empty.Empty) (*FuncSpec, error)
	Destroy(context.Context, *FuncSpec_Args) (*empty.Empty, error)
	// IsLogPlatform returns true if this platform also implements LogPlatform.
	IsLogPlatform(context.Context, *empty.Empty) (*ImplementsResp, error)
}

// UnimplementedPlatformServer can be embedded to have forward compatible implementations.
type UnimplementedPlatformServer struct {
}

func (*UnimplementedPlatformServer) Auth(ctx context.Context, req *FuncSpec_Args) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedPlatformServer) AuthSpec(ctx context.Context, req *Empty) (*FuncSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthSpec not implemented")
}
func (*UnimplementedPlatformServer) ValidateAuth(ctx context.Context, req *FuncSpec_Args) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAuth not implemented")
}
func (*UnimplementedPlatformServer) ValidateAuthSpec(ctx context.Context, req *Empty) (*FuncSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAuthSpec not implemented")
}
func (*UnimplementedPlatformServer) ConfigStruct(ctx context.Context, req *empty.Empty) (*Config_StructResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigStruct not implemented")
}
func (*UnimplementedPlatformServer) Configure(ctx context.Context, req *Config_ConfigureRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedPlatformServer) DeploySpec(ctx context.Context, req *Empty) (*FuncSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeploySpec not implemented")
}
func (*UnimplementedPlatformServer) Deploy(ctx context.Context, req *FuncSpec_Args) (*Deploy_Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}
func (*UnimplementedPlatformServer) IsDestroyer(ctx context.Context, req *empty.Empty) (*ImplementsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsDestroyer not implemented")
}
func (*UnimplementedPlatformServer) DestroySpec(ctx context.Context, req *empty.Empty) (*FuncSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroySpec not implemented")
}
func (*UnimplementedPlatformServer) Destroy(ctx context.Context, req *FuncSpec_Args) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (*UnimplementedPlatformServer) IsLogPlatform(ctx context.Context, req *empty.Empty) (*ImplementsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLogPlatform not implemented")
}

func RegisterPlatformServer(s *grpc.Server, srv PlatformServer) {
	s.RegisterService(&_Platform_serviceDesc, srv)
}

func _Platform_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).Auth(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_AuthSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).AuthSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/AuthSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).AuthSpec(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_ValidateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).ValidateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/ValidateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).ValidateAuth(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_ValidateAuthSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).ValidateAuthSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/ValidateAuthSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).ValidateAuthSpec(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_ConfigStruct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).ConfigStruct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/ConfigStruct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).ConfigStruct(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config_ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).Configure(ctx, req.(*Config_ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_DeploySpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).DeploySpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/DeploySpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).DeploySpec(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).Deploy(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_IsDestroyer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).IsDestroyer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/IsDestroyer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).IsDestroyer(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_DestroySpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).DestroySpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/DestroySpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).DestroySpec(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).Destroy(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_IsLogPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).IsLogPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/IsLogPlatform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).IsLogPlatform(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Platform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Platform",
	HandlerType: (*PlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Platform_Auth_Handler,
		},
		{
			MethodName: "AuthSpec",
			Handler:    _Platform_AuthSpec_Handler,
		},
		{
			MethodName: "ValidateAuth",
			Handler:    _Platform_ValidateAuth_Handler,
		},
		{
			MethodName: "ValidateAuthSpec",
			Handler:    _Platform_ValidateAuthSpec_Handler,
		},
		{
			MethodName: "ConfigStruct",
			Handler:    _Platform_ConfigStruct_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Platform_Configure_Handler,
		},
		{
			MethodName: "DeploySpec",
			Handler:    _Platform_DeploySpec_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _Platform_Deploy_Handler,
		},
		{
			MethodName: "IsDestroyer",
			Handler:    _Platform_IsDestroyer_Handler,
		},
		{
			MethodName: "DestroySpec",
			Handler:    _Platform_DestroySpec_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _Platform_Destroy_Handler,
		},
		{
			MethodName: "IsLogPlatform",
			Handler:    _Platform_IsLogPlatform_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform.proto",
}
