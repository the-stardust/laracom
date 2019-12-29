// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package laracom_user_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	StripeId             string   `protobuf:"bytes,6,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	CardBrand            string   `protobuf:"bytes,7,opt,name=card_brand,json=cardBrand,proto3" json:"card_brand,omitempty"`
	CardLastFour         string   `protobuf:"bytes,8,opt,name=card_last_four,json=cardLastFour,proto3" json:"card_last_four,omitempty"`
	TrialEndsAt          string   `protobuf:"bytes,9,opt,name=trial_ends_at,json=trialEndsAt,proto3" json:"trial_ends_at,omitempty"`
	RememberToken        string   `protobuf:"bytes,11,opt,name=remember_token,json=rememberToken,proto3" json:"remember_token,omitempty"`
	CreatedAt            string   `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *User) GetStripeId() string {
	if m != nil {
		return m.StripeId
	}
	return ""
}

func (m *User) GetCardBrand() string {
	if m != nil {
		return m.CardBrand
	}
	return ""
}

func (m *User) GetCardLastFour() string {
	if m != nil {
		return m.CardLastFour
	}
	return ""
}

func (m *User) GetTrialEndsAt() string {
	if m != nil {
		return m.TrialEndsAt
	}
	return ""
}

func (m *User) GetRememberToken() string {
	if m != nil {
		return m.RememberToken
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Token)(nil), "laracom.user.service.Token")
	proto.RegisterType((*User)(nil), "laracom.user.service.User")
	proto.RegisterType((*Request)(nil), "laracom.user.service.Request")
	proto.RegisterType((*Response)(nil), "laracom.user.service.Response")
	proto.RegisterType((*Error)(nil), "laracom.user.service.Error")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x6d, 0x9c, 0xd8, 0x8d, 0xc7, 0x4d, 0x0e, 0xa3, 0x82, 0xac, 0x54, 0x45, 0x91, 0x05, 0x52,
	0x4e, 0x01, 0xa5, 0x67, 0x0e, 0xa6, 0x6a, 0xab, 0x0a, 0x4e, 0xe6, 0xe3, 0x6a, 0x6d, 0xb2, 0x83,
	0x6a, 0xe1, 0x78, 0xcd, 0xee, 0xba, 0xfc, 0x1e, 0x7e, 0x24, 0x57, 0x84, 0x76, 0xd6, 0x41, 0x1c,
	0xd2, 0x80, 0xe8, 0x25, 0xd9, 0x79, 0xef, 0xed, 0xcc, 0xd3, 0xbe, 0x49, 0xe0, 0x49, 0xab, 0x95,
	0x55, 0x2f, 0x3b, 0x43, 0x9a, 0x3f, 0x96, 0x5c, 0xe3, 0x69, 0x2d, 0xb4, 0xd8, 0xa8, 0xed, 0x92,
	0x31, 0x43, 0xfa, 0xbe, 0xda, 0x50, 0x76, 0x07, 0xe1, 0x07, 0xf5, 0x85, 0x1a, 0x3c, 0x85, 0xd0,
	0xba, 0x43, 0x3a, 0x98, 0x0f, 0x16, 0x71, 0xe1, 0x0b, 0x87, 0xde, 0x8b, 0xba, 0x92, 0x69, 0x30,
	0x1f, 0x2c, 0xc6, 0x85, 0x2f, 0xf0, 0x02, 0x22, 0xd2, 0x5a, 0x69, 0x93, 0x0e, 0xe7, 0xc3, 0x45,
	0xb2, 0x3a, 0x5b, 0xee, 0xeb, 0xbd, 0xbc, 0x72, 0x9a, 0xa2, 0x97, 0x66, 0x3f, 0x03, 0x18, 0x7d,
	0x34, 0xa4, 0x71, 0x0a, 0x41, 0x25, 0xfb, 0x31, 0x41, 0x25, 0x11, 0x61, 0xd4, 0x88, 0x2d, 0xf1,
	0x88, 0xb8, 0xe0, 0xb3, 0x9b, 0x4b, 0x5b, 0x51, 0xd5, 0xe9, 0xd0, 0xbb, 0xe1, 0x02, 0x67, 0x30,
	0x6e, 0x85, 0x31, 0xdf, 0x94, 0x96, 0xe9, 0x88, 0x89, 0xdf, 0x35, 0x3e, 0x85, 0xc8, 0x58, 0x61,
	0x3b, 0x93, 0x86, 0xcc, 0xf4, 0x15, 0x9e, 0x41, 0x6c, 0xac, 0xae, 0x5a, 0x2a, 0x2b, 0x99, 0x46,
	0xfe, 0x92, 0x07, 0x6e, 0x25, 0x9e, 0x03, 0x6c, 0x84, 0x96, 0xe5, 0x5a, 0x8b, 0x46, 0xa6, 0xc7,
	0xcc, 0xc6, 0x0e, 0x79, 0xe3, 0x00, 0x7c, 0x0e, 0x53, 0xa6, 0x6b, 0x61, 0x6c, 0xf9, 0x59, 0x75,
	0x3a, 0x1d, 0xb3, 0xe4, 0xc4, 0xa1, 0xef, 0x84, 0xb1, 0xd7, 0xaa, 0xd3, 0x98, 0xc1, 0xc4, 0xea,
	0x4a, 0xd4, 0x25, 0x35, 0xd2, 0x94, 0xc2, 0xa6, 0x31, 0x8b, 0x12, 0x06, 0xaf, 0x1a, 0x69, 0x72,
	0xeb, 0x06, 0x49, 0xaa, 0xc9, 0x92, 0x74, 0x02, 0xf0, 0x83, 0x7a, 0x24, 0xb7, 0xf8, 0x02, 0xa6,
	0x9a, 0xb6, 0xb4, 0x5d, 0x93, 0x2e, 0x7d, 0x0a, 0x09, 0x4b, 0x26, 0x3b, 0xd4, 0x67, 0xe4, 0xec,
	0x6a, 0x12, 0x7d, 0x97, 0x93, 0xde, 0xae, 0x47, 0xfc, 0x90, 0xae, 0x95, 0x3b, 0x7a, 0xe2, 0xe9,
	0x1e, 0xc9, 0x6d, 0x16, 0xc3, 0x71, 0x41, 0x5f, 0x3b, 0x32, 0x36, 0xfb, 0x3e, 0x80, 0x71, 0x41,
	0xa6, 0x55, 0x8d, 0x21, 0x5c, 0xc2, 0xc8, 0xc5, 0xc6, 0x89, 0x24, 0xab, 0xd9, 0xfe, 0x2c, 0x5d,
	0x72, 0x05, 0xeb, 0xf0, 0x15, 0x84, 0xee, 0xdb, 0xa4, 0x01, 0x87, 0x7f, 0xe8, 0x82, 0x17, 0xfe,
	0xdf, 0xbe, 0xbc, 0x86, 0x90, 0x01, 0xb7, 0x1f, 0x1b, 0x25, 0x89, 0xfd, 0x85, 0x05, 0x9f, 0x71,
	0x0e, 0x89, 0x24, 0xb3, 0xd1, 0x55, 0x6b, 0x2b, 0xd5, 0xf4, 0xab, 0xf3, 0x27, 0xb4, 0xfa, 0x11,
	0x40, 0xe2, 0x3c, 0xbc, 0xf7, 0xcd, 0xf1, 0x1a, 0xa2, 0x4b, 0x7e, 0x29, 0x3c, 0x60, 0x78, 0xf6,
	0x6c, 0x3f, 0xb7, 0x7b, 0xab, 0xec, 0x08, 0x2f, 0x61, 0x78, 0x43, 0xf6, 0x91, 0x4d, 0x6e, 0x21,
	0xba, 0x21, 0x9b, 0xd7, 0x35, 0x9e, 0x3f, 0xa4, 0xe5, 0xa0, 0xfe, 0xa1, 0x55, 0x0e, 0xa3, 0xbc,
	0xb3, 0x77, 0x07, 0x0d, 0x3d, 0xf0, 0xde, 0xbc, 0x54, 0xd9, 0x11, 0xbe, 0x85, 0xc9, 0x27, 0xf7,
	0xbb, 0x16, 0x96, 0xfc, 0x9e, 0x1d, 0xd2, 0xff, 0xa5, 0xd9, 0x3a, 0xe2, 0x7f, 0x9b, 0x8b, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xd4, 0xc7, 0x06, 0x86, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "laracom.user.service"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}
