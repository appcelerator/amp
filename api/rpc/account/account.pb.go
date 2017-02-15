// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/account/account.proto
// DO NOT EDIT!

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/account/account.proto

It has these top-level messages:
	SignUpRequest
	SignUpReply
	VerificationRequest
	LogInRequest
	LogInReply
	PasswordResetRequest
	PasswordResetReply
	PasswordSetRequest
	PasswordChangeRequest
	CreateOrganizationRequest
*/
package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

type SignUpRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *SignUpRequest) Reset()                    { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string            { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()               {}
func (*SignUpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignUpRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignUpRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type SignUpReply struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *SignUpReply) Reset()                    { *m = SignUpReply{} }
func (m *SignUpReply) String() string            { return proto.CompactTextString(m) }
func (*SignUpReply) ProtoMessage()               {}
func (*SignUpReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignUpReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SignUpReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerificationRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *VerificationRequest) Reset()                    { *m = VerificationRequest{} }
func (m *VerificationRequest) String() string            { return proto.CompactTextString(m) }
func (*VerificationRequest) ProtoMessage()               {}
func (*VerificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VerificationRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogInRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LogInRequest) Reset()                    { *m = LogInRequest{} }
func (m *LogInRequest) String() string            { return proto.CompactTextString(m) }
func (*LogInRequest) ProtoMessage()               {}
func (*LogInRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LogInRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LogInReply struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LogInReply) Reset()                    { *m = LogInReply{} }
func (m *LogInReply) String() string            { return proto.CompactTextString(m) }
func (*LogInReply) ProtoMessage()               {}
func (*LogInReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LogInReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type PasswordResetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PasswordResetRequest) Reset()                    { *m = PasswordResetRequest{} }
func (m *PasswordResetRequest) String() string            { return proto.CompactTextString(m) }
func (*PasswordResetRequest) ProtoMessage()               {}
func (*PasswordResetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PasswordResetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PasswordResetReply struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *PasswordResetReply) Reset()                    { *m = PasswordResetReply{} }
func (m *PasswordResetReply) String() string            { return proto.CompactTextString(m) }
func (*PasswordResetReply) ProtoMessage()               {}
func (*PasswordResetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PasswordResetReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type PasswordSetRequest struct {
	Token    string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *PasswordSetRequest) Reset()                    { *m = PasswordSetRequest{} }
func (m *PasswordSetRequest) String() string            { return proto.CompactTextString(m) }
func (*PasswordSetRequest) ProtoMessage()               {}
func (*PasswordSetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PasswordSetRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PasswordSetRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type PasswordChangeRequest struct {
	Name             string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ExistingPassword string `protobuf:"bytes,2,opt,name=existingPassword" json:"existingPassword,omitempty"`
	NewPassword      string `protobuf:"bytes,3,opt,name=newPassword" json:"newPassword,omitempty"`
}

func (m *PasswordChangeRequest) Reset()                    { *m = PasswordChangeRequest{} }
func (m *PasswordChangeRequest) String() string            { return proto.CompactTextString(m) }
func (*PasswordChangeRequest) ProtoMessage()               {}
func (*PasswordChangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PasswordChangeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PasswordChangeRequest) GetExistingPassword() string {
	if m != nil {
		return m.ExistingPassword
	}
	return ""
}

func (m *PasswordChangeRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type CreateOrganizationRequest struct {
	Owner string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *CreateOrganizationRequest) Reset()                    { *m = CreateOrganizationRequest{} }
func (m *CreateOrganizationRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateOrganizationRequest) ProtoMessage()               {}
func (*CreateOrganizationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CreateOrganizationRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CreateOrganizationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateOrganizationRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*SignUpRequest)(nil), "account.SignUpRequest")
	proto.RegisterType((*SignUpReply)(nil), "account.SignUpReply")
	proto.RegisterType((*VerificationRequest)(nil), "account.VerificationRequest")
	proto.RegisterType((*LogInRequest)(nil), "account.LogInRequest")
	proto.RegisterType((*LogInReply)(nil), "account.LogInReply")
	proto.RegisterType((*PasswordResetRequest)(nil), "account.PasswordResetRequest")
	proto.RegisterType((*PasswordResetReply)(nil), "account.PasswordResetReply")
	proto.RegisterType((*PasswordSetRequest)(nil), "account.PasswordSetRequest")
	proto.RegisterType((*PasswordChangeRequest)(nil), "account.PasswordChangeRequest")
	proto.RegisterType((*CreateOrganizationRequest)(nil), "account.CreateOrganizationRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Account service

type AccountClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpReply, error)
	Verify(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Login(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInReply, error)
	PasswordReset(ctx context.Context, in *PasswordResetRequest, opts ...grpc.CallOption) (*PasswordResetReply, error)
	PasswordSet(ctx context.Context, in *PasswordSetRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	PasswordChange(ctx context.Context, in *PasswordChangeRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpReply, error) {
	out := new(SignUpReply)
	err := grpc.Invoke(ctx, "/account.Account/SignUp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Verify(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/Verify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Login(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInReply, error) {
	out := new(LogInReply)
	err := grpc.Invoke(ctx, "/account.Account/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) PasswordReset(ctx context.Context, in *PasswordResetRequest, opts ...grpc.CallOption) (*PasswordResetReply, error) {
	out := new(PasswordResetReply)
	err := grpc.Invoke(ctx, "/account.Account/PasswordReset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) PasswordSet(ctx context.Context, in *PasswordSetRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/PasswordSet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) PasswordChange(ctx context.Context, in *PasswordChangeRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/PasswordChange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/CreateOrganization", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpReply, error)
	Verify(context.Context, *VerificationRequest) (*google_protobuf1.Empty, error)
	Login(context.Context, *LogInRequest) (*LogInReply, error)
	PasswordReset(context.Context, *PasswordResetRequest) (*PasswordResetReply, error)
	PasswordSet(context.Context, *PasswordSetRequest) (*google_protobuf1.Empty, error)
	PasswordChange(context.Context, *PasswordChangeRequest) (*google_protobuf1.Empty, error)
	CreateOrganization(context.Context, *CreateOrganizationRequest) (*google_protobuf1.Empty, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Verify(ctx, req.(*VerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Login(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_PasswordReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).PasswordReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/PasswordReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).PasswordReset(ctx, req.(*PasswordResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_PasswordSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).PasswordSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/PasswordSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).PasswordSet(ctx, req.(*PasswordSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_PasswordChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).PasswordChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/PasswordChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).PasswordChange(ctx, req.(*PasswordChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_CreateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).CreateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/CreateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).CreateOrganization(ctx, req.(*CreateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Account_SignUp_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Account_Verify_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Account_Login_Handler,
		},
		{
			MethodName: "PasswordReset",
			Handler:    _Account_PasswordReset_Handler,
		},
		{
			MethodName: "PasswordSet",
			Handler:    _Account_PasswordSet_Handler,
		},
		{
			MethodName: "PasswordChange",
			Handler:    _Account_PasswordChange_Handler,
		},
		{
			MethodName: "CreateOrganization",
			Handler:    _Account_CreateOrganization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/account/account.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/account/account.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x95, 0xc2, 0x36, 0x38, 0x65, 0x13, 0x78, 0x5d, 0x57, 0xb2, 0x0e, 0x15, 0x4b, 0xfc,
	0x51, 0x37, 0x25, 0x62, 0xbb, 0x62, 0x17, 0x48, 0x68, 0x02, 0x09, 0x69, 0x12, 0xa3, 0x53, 0xb9,
	0x41, 0x5c, 0xb8, 0x99, 0x9b, 0x59, 0xa4, 0xb6, 0x49, 0xdc, 0x76, 0x05, 0xed, 0x06, 0xf1, 0x06,
	0x3c, 0x1a, 0xaf, 0xc0, 0x1b, 0xf0, 0x02, 0xc8, 0xce, 0x9f, 0x26, 0x6d, 0xd2, 0x0b, 0xae, 0xda,
	0x63, 0x1f, 0xff, 0xbe, 0xcf, 0xf1, 0x77, 0xe0, 0xa5, 0xcf, 0xd4, 0xd5, 0x78, 0xe0, 0x78, 0x62,
	0xe4, 0x12, 0x29, 0x3d, 0x1a, 0xd0, 0x90, 0x28, 0x11, 0xba, 0x64, 0x24, 0x5d, 0x22, 0x99, 0x1b,
	0x4a, 0xcf, 0x25, 0x9e, 0x27, 0xc6, 0x5c, 0xa5, 0xbf, 0x8e, 0x0c, 0x85, 0x12, 0x68, 0x23, 0x29,
	0xed, 0xb6, 0x2f, 0x84, 0x1f, 0x50, 0xd3, 0x4e, 0x38, 0x17, 0x8a, 0x28, 0x26, 0x78, 0x14, 0xb7,
	0xd9, 0x7b, 0xc9, 0xae, 0xa9, 0x06, 0xe3, 0xa1, 0x4b, 0x47, 0x52, 0xcd, 0xe2, 0x4d, 0xdc, 0x87,
	0xcd, 0x0b, 0xe6, 0xf3, 0xbe, 0xec, 0xd1, 0xaf, 0x63, 0x1a, 0x29, 0x84, 0xe0, 0x36, 0x27, 0x23,
	0xda, 0xb2, 0x3a, 0xd6, 0xf3, 0xbb, 0x3d, 0xf3, 0x1f, 0xd9, 0x70, 0x47, 0x92, 0x28, 0x9a, 0x8a,
	0xf0, 0xb2, 0x55, 0x33, 0xeb, 0x59, 0x8d, 0x1a, 0xb0, 0x46, 0x47, 0x84, 0x05, 0xad, 0x5b, 0x66,
	0x23, 0x2e, 0xf0, 0x31, 0xd4, 0x53, 0xac, 0x0c, 0x66, 0x68, 0x0b, 0x6a, 0xec, 0x32, 0x41, 0xd6,
	0x98, 0x39, 0xa4, 0xc4, 0x17, 0xca, 0x13, 0x5a, 0x5c, 0xe0, 0x03, 0xd8, 0xfe, 0x48, 0x43, 0x36,
	0x64, 0x9e, 0xf1, 0x9f, 0x3a, 0xca, 0x9a, 0xad, 0x7c, 0xf3, 0x2b, 0xb8, 0x77, 0x26, 0xfc, 0x77,
	0xfc, 0x3f, 0x7d, 0x63, 0x0c, 0x90, 0x9c, 0xd7, 0x06, 0xcb, 0x35, 0xba, 0xd0, 0x38, 0x4f, 0xfa,
	0x7b, 0x34, 0xa2, 0x6a, 0x85, 0x16, 0xee, 0x02, 0x5a, 0xe8, 0xad, 0xe6, 0xbe, 0x9d, 0xf7, 0x5e,
	0xcc, 0xa9, 0xa5, 0xbd, 0x2b, 0xef, 0x30, 0x83, 0x9d, 0x94, 0x73, 0x7a, 0x45, 0xb8, 0x4f, 0x57,
	0x7d, 0x8c, 0x2e, 0xdc, 0xa7, 0xd7, 0x2c, 0x52, 0x8c, 0xfb, 0xe7, 0x45, 0xe0, 0xd2, 0x3a, 0xea,
	0x40, 0x9d, 0xd3, 0x69, 0xd6, 0x16, 0x3f, 0x6d, 0x7e, 0x09, 0x7f, 0x82, 0x87, 0xa7, 0x21, 0x25,
	0x8a, 0xbe, 0x0f, 0x7d, 0xc2, 0xd9, 0xb7, 0xc5, 0x17, 0x13, 0x53, 0x4e, 0xc3, 0xf4, 0x26, 0xa6,
	0xc8, 0x4c, 0xd5, 0x72, 0xa6, 0x4a, 0xd3, 0x73, 0xf4, 0x77, 0x0d, 0x36, 0x5e, 0xc7, 0xd9, 0x46,
	0x7d, 0x58, 0x8f, 0x93, 0x84, 0x9a, 0x4e, 0x1a, 0xff, 0x42, 0x62, 0xed, 0xc6, 0xd2, 0xba, 0x0c,
	0x66, 0x78, 0xff, 0xc7, 0xef, 0x3f, 0xbf, 0x6a, 0xbb, 0x18, 0xb9, 0x93, 0x17, 0xd9, 0xfc, 0x44,
	0xcc, 0xe7, 0x63, 0x79, 0x62, 0x75, 0xd1, 0x67, 0x58, 0x37, 0x59, 0x9b, 0xa1, 0x76, 0x76, 0xbc,
	0x24, 0x7c, 0x76, 0xd3, 0x89, 0xa7, 0xc7, 0x49, 0xa7, 0xc7, 0x79, 0xa3, 0xa7, 0xa7, 0x1c, 0x3f,
	0x31, 0x44, 0x8d, 0xff, 0x00, 0x6b, 0x67, 0xc2, 0x67, 0x1c, 0xed, 0x64, 0xf4, 0x7c, 0x5a, 0xed,
	0xed, 0xc5, 0x65, 0x6d, 0xb9, 0x6d, 0x98, 0x4d, 0xfc, 0x20, 0xcf, 0x0c, 0x34, 0x46, 0x23, 0xaf,
	0x61, 0xb3, 0x10, 0x30, 0xb4, 0x9f, 0x31, 0xca, 0x42, 0x6a, 0xef, 0x55, 0x6d, 0x6b, 0xa9, 0x03,
	0x23, 0xf5, 0x04, 0x77, 0xf2, 0x52, 0xdf, 0xf5, 0x93, 0xdc, 0xb8, 0x32, 0xdf, 0xae, 0x95, 0x87,
	0x50, 0xcf, 0xc5, 0x15, 0x2d, 0x83, 0xe7, 0x21, 0xae, 0xfc, 0x5e, 0xd8, 0x08, 0xb6, 0xf1, 0x6e,
	0x5e, 0x50, 0xce, 0xcf, 0x6b, 0x9d, 0x09, 0x6c, 0x15, 0xe3, 0x8c, 0x1e, 0x2d, 0x49, 0x15, 0x72,
	0x5e, 0xa9, 0x76, 0x68, 0xd4, 0x9e, 0xe2, 0xc7, 0x2b, 0xae, 0x17, 0x93, 0xb4, 0xee, 0x4f, 0x0b,
	0xd0, 0x72, 0x98, 0x11, 0xce, 0xc4, 0x2b, 0x93, 0x5e, 0x69, 0xe0, 0xc8, 0x18, 0x38, 0xc4, 0xcf,
	0x0a, 0x06, 0xcc, 0x1c, 0xdc, 0xb8, 0x22, 0x07, 0x72, 0x3d, 0xc3, 0x3e, 0xb1, 0xba, 0x83, 0x75,
	0xc3, 0x38, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x57, 0xd6, 0x19, 0x88, 0x12, 0x06, 0x00, 0x00,
}
