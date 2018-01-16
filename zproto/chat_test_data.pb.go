// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat_test_data.proto

package zproto

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

type ChatMessage struct {
	SenderSessionId string `protobuf:"bytes,1,opt,name=sender_session_id,json=senderSessionId" json:"sender_session_id,omitempty"`
	MessageData     string `protobuf:"bytes,3,opt,name=message_data,json=messageData" json:"message_data,omitempty"`
}

func (m *ChatMessage) Reset()                    { *m = ChatMessage{} }
func (m *ChatMessage) String() string            { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()               {}
func (*ChatMessage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ChatMessage) GetSenderSessionId() string {
	if m != nil {
		return m.SenderSessionId
	}
	return ""
}

func (m *ChatMessage) GetMessageData() string {
	if m != nil {
		return m.MessageData
	}
	return ""
}

type ChatSession struct {
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
}

func (m *ChatSession) Reset()                    { *m = ChatSession{} }
func (m *ChatSession) String() string            { return proto.CompactTextString(m) }
func (*ChatSession) ProtoMessage()               {}
func (*ChatSession) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ChatSession) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type VoidRsp2 struct {
}

func (m *VoidRsp2) Reset()                    { *m = VoidRsp2{} }
func (m *VoidRsp2) String() string            { return proto.CompactTextString(m) }
func (*VoidRsp2) ProtoMessage()               {}
func (*VoidRsp2) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func init() {
	proto.RegisterType((*ChatMessage)(nil), "zproto.ChatMessage")
	proto.RegisterType((*ChatSession)(nil), "zproto.ChatSession")
	proto.RegisterType((*VoidRsp2)(nil), "zproto.VoidRsp2")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChatTest service

type ChatTestClient interface {
	Connect(ctx context.Context, in *ChatSession, opts ...grpc.CallOption) (ChatTest_ConnectClient, error)
	SendChat(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*VoidRsp2, error)
}

type chatTestClient struct {
	cc *grpc.ClientConn
}

func NewChatTestClient(cc *grpc.ClientConn) ChatTestClient {
	return &chatTestClient{cc}
}

func (c *chatTestClient) Connect(ctx context.Context, in *ChatSession, opts ...grpc.CallOption) (ChatTest_ConnectClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChatTest_serviceDesc.Streams[0], c.cc, "/zproto.ChatTest/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatTestConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatTest_ConnectClient interface {
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatTestConnectClient struct {
	grpc.ClientStream
}

func (x *chatTestConnectClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatTestClient) SendChat(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*VoidRsp2, error) {
	out := new(VoidRsp2)
	err := grpc.Invoke(ctx, "/zproto.ChatTest/SendChat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChatTest service

type ChatTestServer interface {
	Connect(*ChatSession, ChatTest_ConnectServer) error
	SendChat(context.Context, *ChatMessage) (*VoidRsp2, error)
}

func RegisterChatTestServer(s *grpc.Server, srv ChatTestServer) {
	s.RegisterService(&_ChatTest_serviceDesc, srv)
}

func _ChatTest_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChatSession)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatTestServer).Connect(m, &chatTestConnectServer{stream})
}

type ChatTest_ConnectServer interface {
	Send(*ChatMessage) error
	grpc.ServerStream
}

type chatTestConnectServer struct {
	grpc.ServerStream
}

func (x *chatTestConnectServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatTest_SendChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatTestServer).SendChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zproto.ChatTest/SendChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatTestServer).SendChat(ctx, req.(*ChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatTest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zproto.ChatTest",
	HandlerType: (*ChatTestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendChat",
			Handler:    _ChatTest_SendChat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _ChatTest_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat_test_data.proto",
}

func init() { proto.RegisterFile("chat_test_data.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x42, 0x4d, 0xa7, 0xe2, 0xc7, 0xd4, 0x43, 0x29, 0x08, 0x9a, 0x93, 0x8a, 0x04,
	0xad, 0xf8, 0x07, 0x5a, 0x0f, 0x7a, 0x10, 0x4a, 0x2b, 0x1e, 0x44, 0x08, 0xd3, 0xec, 0x60, 0x03,
	0x66, 0xb7, 0x64, 0xc6, 0x8b, 0xbf, 0x5e, 0x92, 0xdd, 0x05, 0xc5, 0x9e, 0x06, 0x9e, 0x77, 0xde,
	0x67, 0x67, 0xe1, 0xa4, 0x5c, 0x93, 0x16, 0xca, 0xa2, 0x85, 0x21, 0xa5, 0x7c, 0xd3, 0x38, 0x75,
	0xd8, 0xfb, 0xee, 0x66, 0xf6, 0x0e, 0x83, 0xd9, 0x9a, 0xf4, 0x99, 0x45, 0xe8, 0x83, 0xf1, 0x0a,
	0x8e, 0x85, 0xad, 0xe1, 0xa6, 0x10, 0x16, 0xa9, 0x9c, 0x2d, 0x2a, 0x33, 0x4a, 0xce, 0x92, 0x8b,
	0xfe, 0xe2, 0xd0, 0x07, 0x4b, 0xcf, 0x9f, 0x0c, 0x9e, 0xc3, 0x7e, 0xed, 0x6b, 0x9d, 0x78, 0xb4,
	0xdb, 0xad, 0x0d, 0x02, 0x7b, 0x20, 0xa5, 0xec, 0xda, 0xdb, 0x43, 0x07, 0x4f, 0x01, 0xfe, 0x69,
	0xfb, 0x12, 0x85, 0x19, 0x40, 0xfa, 0xea, 0x2a, 0xb3, 0x90, 0xcd, 0x64, 0xa2, 0x90, 0xb6, 0xcd,
	0x17, 0x16, 0xc5, 0x7b, 0xd8, 0x9b, 0x39, 0x6b, 0xb9, 0x54, 0x1c, 0xe6, 0xfe, 0xee, 0xfc, 0x97,
	0x76, 0xfc, 0x07, 0x86, 0x9f, 0xdc, 0x24, 0x78, 0x0b, 0xe9, 0x92, 0xad, 0x69, 0x21, 0x6e, 0x5b,
	0x19, 0x1f, 0x45, 0x18, 0x5f, 0x9d, 0x5e, 0xc2, 0xb0, 0x74, 0x75, 0x6e, 0x79, 0xf5, 0xf5, 0x49,
	0x55, 0x1d, 0xf2, 0xe9, 0xc1, 0xdb, 0xbc, 0x9d, 0xf1, 0xa0, 0xc7, 0x9d, 0x79, 0xb2, 0xea, 0x75,
	0xd1, 0xdd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0xe7, 0x5b, 0x83, 0x5f, 0x01, 0x00, 0x00,
}