// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/currency.proto

package currency

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Currency service

func NewCurrencyEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Currency service

type CurrencyService interface {
	Codes(ctx context.Context, in *CodesRequest, opts ...client.CallOption) (*CodesResponse, error)
	Rates(ctx context.Context, in *RatesRequest, opts ...client.CallOption) (*RatesResponse, error)
	Convert(ctx context.Context, in *ConvertRequest, opts ...client.CallOption) (*ConvertResponse, error)
}

type currencyService struct {
	c    client.Client
	name string
}

func NewCurrencyService(name string, c client.Client) CurrencyService {
	return &currencyService{
		c:    c,
		name: name,
	}
}

func (c *currencyService) Codes(ctx context.Context, in *CodesRequest, opts ...client.CallOption) (*CodesResponse, error) {
	req := c.c.NewRequest(c.name, "Currency.Codes", in)
	out := new(CodesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyService) Rates(ctx context.Context, in *RatesRequest, opts ...client.CallOption) (*RatesResponse, error) {
	req := c.c.NewRequest(c.name, "Currency.Rates", in)
	out := new(RatesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyService) Convert(ctx context.Context, in *ConvertRequest, opts ...client.CallOption) (*ConvertResponse, error) {
	req := c.c.NewRequest(c.name, "Currency.Convert", in)
	out := new(ConvertResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Currency service

type CurrencyHandler interface {
	Codes(context.Context, *CodesRequest, *CodesResponse) error
	Rates(context.Context, *RatesRequest, *RatesResponse) error
	Convert(context.Context, *ConvertRequest, *ConvertResponse) error
}

func RegisterCurrencyHandler(s server.Server, hdlr CurrencyHandler, opts ...server.HandlerOption) error {
	type currency interface {
		Codes(ctx context.Context, in *CodesRequest, out *CodesResponse) error
		Rates(ctx context.Context, in *RatesRequest, out *RatesResponse) error
		Convert(ctx context.Context, in *ConvertRequest, out *ConvertResponse) error
	}
	type Currency struct {
		currency
	}
	h := &currencyHandler{hdlr}
	return s.Handle(s.NewHandler(&Currency{h}, opts...))
}

type currencyHandler struct {
	CurrencyHandler
}

func (h *currencyHandler) Codes(ctx context.Context, in *CodesRequest, out *CodesResponse) error {
	return h.CurrencyHandler.Codes(ctx, in, out)
}

func (h *currencyHandler) Rates(ctx context.Context, in *RatesRequest, out *RatesResponse) error {
	return h.CurrencyHandler.Rates(ctx, in, out)
}

func (h *currencyHandler) Convert(ctx context.Context, in *ConvertRequest, out *ConvertResponse) error {
	return h.CurrencyHandler.Convert(ctx, in, out)
}