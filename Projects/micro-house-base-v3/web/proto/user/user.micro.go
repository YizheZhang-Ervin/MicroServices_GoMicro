// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	SendSms(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Register(ctx context.Context, in *RegReq, opts ...client.CallOption) (*Response, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) SendSms(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.SendSms", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Register(ctx context.Context, in *RegReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	SendSms(context.Context, *Request, *Response) error
	Register(context.Context, *RegReq, *Response) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		SendSms(ctx context.Context, in *Request, out *Response) error
		Register(ctx context.Context, in *RegReq, out *Response) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) SendSms(ctx context.Context, in *Request, out *Response) error {
	return h.UserHandler.SendSms(ctx, in, out)
}

func (h *userHandler) Register(ctx context.Context, in *RegReq, out *Response) error {
	return h.UserHandler.Register(ctx, in, out)
}
