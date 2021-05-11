// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/threads.proto

package threads

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

// Api Endpoints for Threads service

func NewThreadsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Threads service

type ThreadsService interface {
	// Create a thread
	CreateThread(ctx context.Context, in *CreateThreadRequest, opts ...client.CallOption) (*CreateThreadResponse, error)
	// Read a thread using its ID, can filter using group ID if provided
	ReadThread(ctx context.Context, in *ReadThreadRequest, opts ...client.CallOption) (*ReadThreadResponse, error)
	// Update a threads topic
	UpdateThread(ctx context.Context, in *UpdateThreadRequest, opts ...client.CallOption) (*UpdateThreadResponse, error)
	// Delete a thread and all the messages within
	DeleteThread(ctx context.Context, in *DeleteThreadRequest, opts ...client.CallOption) (*DeleteThreadResponse, error)
	// List all the threads for a group
	ListThreads(ctx context.Context, in *ListThreadsRequest, opts ...client.CallOption) (*ListThreadsResponse, error)
	// Create a message within a thread
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...client.CallOption) (*CreateMessageResponse, error)
	// List the messages within a thread in reverse chronological order, using sent_before to
	// offset as older messages need to be loaded
	ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...client.CallOption) (*ListMessagesResponse, error)
	// RecentMessages returns the most recent messages in a group of threads. By default the
	// most messages retrieved per thread is 25, however this can be overriden using the
	// limit_per_thread option
	RecentMessages(ctx context.Context, in *RecentMessagesRequest, opts ...client.CallOption) (*RecentMessagesResponse, error)
}

type threadsService struct {
	c    client.Client
	name string
}

func NewThreadsService(name string, c client.Client) ThreadsService {
	return &threadsService{
		c:    c,
		name: name,
	}
}

func (c *threadsService) CreateThread(ctx context.Context, in *CreateThreadRequest, opts ...client.CallOption) (*CreateThreadResponse, error) {
	req := c.c.NewRequest(c.name, "Threads.CreateThread", in)
	out := new(CreateThreadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadsService) ReadThread(ctx context.Context, in *ReadThreadRequest, opts ...client.CallOption) (*ReadThreadResponse, error) {
	req := c.c.NewRequest(c.name, "Threads.ReadThread", in)
	out := new(ReadThreadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadsService) UpdateThread(ctx context.Context, in *UpdateThreadRequest, opts ...client.CallOption) (*UpdateThreadResponse, error) {
	req := c.c.NewRequest(c.name, "Threads.UpdateThread", in)
	out := new(UpdateThreadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadsService) DeleteThread(ctx context.Context, in *DeleteThreadRequest, opts ...client.CallOption) (*DeleteThreadResponse, error) {
	req := c.c.NewRequest(c.name, "Threads.DeleteThread", in)
	out := new(DeleteThreadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadsService) ListThreads(ctx context.Context, in *ListThreadsRequest, opts ...client.CallOption) (*ListThreadsResponse, error) {
	req := c.c.NewRequest(c.name, "Threads.ListThreads", in)
	out := new(ListThreadsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadsService) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...client.CallOption) (*CreateMessageResponse, error) {
	req := c.c.NewRequest(c.name, "Threads.CreateMessage", in)
	out := new(CreateMessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadsService) ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...client.CallOption) (*ListMessagesResponse, error) {
	req := c.c.NewRequest(c.name, "Threads.ListMessages", in)
	out := new(ListMessagesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadsService) RecentMessages(ctx context.Context, in *RecentMessagesRequest, opts ...client.CallOption) (*RecentMessagesResponse, error) {
	req := c.c.NewRequest(c.name, "Threads.RecentMessages", in)
	out := new(RecentMessagesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Threads service

type ThreadsHandler interface {
	// Create a thread
	CreateThread(context.Context, *CreateThreadRequest, *CreateThreadResponse) error
	// Read a thread using its ID, can filter using group ID if provided
	ReadThread(context.Context, *ReadThreadRequest, *ReadThreadResponse) error
	// Update a threads topic
	UpdateThread(context.Context, *UpdateThreadRequest, *UpdateThreadResponse) error
	// Delete a thread and all the messages within
	DeleteThread(context.Context, *DeleteThreadRequest, *DeleteThreadResponse) error
	// List all the threads for a group
	ListThreads(context.Context, *ListThreadsRequest, *ListThreadsResponse) error
	// Create a message within a thread
	CreateMessage(context.Context, *CreateMessageRequest, *CreateMessageResponse) error
	// List the messages within a thread in reverse chronological order, using sent_before to
	// offset as older messages need to be loaded
	ListMessages(context.Context, *ListMessagesRequest, *ListMessagesResponse) error
	// RecentMessages returns the most recent messages in a group of threads. By default the
	// most messages retrieved per thread is 25, however this can be overriden using the
	// limit_per_thread option
	RecentMessages(context.Context, *RecentMessagesRequest, *RecentMessagesResponse) error
}

func RegisterThreadsHandler(s server.Server, hdlr ThreadsHandler, opts ...server.HandlerOption) error {
	type threads interface {
		CreateThread(ctx context.Context, in *CreateThreadRequest, out *CreateThreadResponse) error
		ReadThread(ctx context.Context, in *ReadThreadRequest, out *ReadThreadResponse) error
		UpdateThread(ctx context.Context, in *UpdateThreadRequest, out *UpdateThreadResponse) error
		DeleteThread(ctx context.Context, in *DeleteThreadRequest, out *DeleteThreadResponse) error
		ListThreads(ctx context.Context, in *ListThreadsRequest, out *ListThreadsResponse) error
		CreateMessage(ctx context.Context, in *CreateMessageRequest, out *CreateMessageResponse) error
		ListMessages(ctx context.Context, in *ListMessagesRequest, out *ListMessagesResponse) error
		RecentMessages(ctx context.Context, in *RecentMessagesRequest, out *RecentMessagesResponse) error
	}
	type Threads struct {
		threads
	}
	h := &threadsHandler{hdlr}
	return s.Handle(s.NewHandler(&Threads{h}, opts...))
}

type threadsHandler struct {
	ThreadsHandler
}

func (h *threadsHandler) CreateThread(ctx context.Context, in *CreateThreadRequest, out *CreateThreadResponse) error {
	return h.ThreadsHandler.CreateThread(ctx, in, out)
}

func (h *threadsHandler) ReadThread(ctx context.Context, in *ReadThreadRequest, out *ReadThreadResponse) error {
	return h.ThreadsHandler.ReadThread(ctx, in, out)
}

func (h *threadsHandler) UpdateThread(ctx context.Context, in *UpdateThreadRequest, out *UpdateThreadResponse) error {
	return h.ThreadsHandler.UpdateThread(ctx, in, out)
}

func (h *threadsHandler) DeleteThread(ctx context.Context, in *DeleteThreadRequest, out *DeleteThreadResponse) error {
	return h.ThreadsHandler.DeleteThread(ctx, in, out)
}

func (h *threadsHandler) ListThreads(ctx context.Context, in *ListThreadsRequest, out *ListThreadsResponse) error {
	return h.ThreadsHandler.ListThreads(ctx, in, out)
}

func (h *threadsHandler) CreateMessage(ctx context.Context, in *CreateMessageRequest, out *CreateMessageResponse) error {
	return h.ThreadsHandler.CreateMessage(ctx, in, out)
}

func (h *threadsHandler) ListMessages(ctx context.Context, in *ListMessagesRequest, out *ListMessagesResponse) error {
	return h.ThreadsHandler.ListMessages(ctx, in, out)
}

func (h *threadsHandler) RecentMessages(ctx context.Context, in *RecentMessagesRequest, out *RecentMessagesResponse) error {
	return h.ThreadsHandler.RecentMessages(ctx, in, out)
}
