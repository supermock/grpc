// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: chat/chat.proto

package chat

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	recipe "github.com/supermock/grpc/recipe"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_chat_chat_proto protoreflect.FileDescriptor

var file_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73,
	0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb9, 0x04, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7f, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x0f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x73, 0x61, 0x79, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x3a,
	0x01, 0x2a, 0x92, 0x41, 0x32, 0x0a, 0x03, 0x72, 0x70, 0x63, 0x12, 0x15, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x3a, 0x20, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x72, 0x70,
	0x63, 0x1a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x53,
	0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x3e, 0x0a, 0x18, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x12, 0x78, 0x0a, 0x18, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x73,
	0x61, 0x79, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x72, 0x6f,
	0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x92, 0x41, 0x05, 0x0a, 0x03, 0x72, 0x70, 0x63, 0x30,
	0x01, 0x12, 0x43, 0x0a, 0x1b, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x12, 0x0f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x81, 0x01, 0x0a, 0x1a, 0x57, 0x69, 0x74, 0x68, 0x6f,
	0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x61, 0x6e, 0x64, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x92, 0x41, 0x05, 0x0a, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x26, 0x92, 0x41, 0x23, 0x12,
	0x21, 0x43, 0x68, 0x61, 0x74, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x61, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x92, 0x41, 0x3e, 0x12, 0x13, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x74, 0x20,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x2a, 0x03, 0x01, 0x02,
	0x04, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_chat_chat_proto_goTypes = []interface{}{
	(*recipe.Message)(nil), // 0: recipe.Message
	(*empty.Empty)(nil),    // 1: google.protobuf.Empty
}
var file_chat_chat_proto_depIdxs = []int32{
	0, // 0: chat.ChatService.SayHello:input_type -> recipe.Message
	0, // 1: chat.ChatService.SayHelloStreamFromClient:input_type -> recipe.Message
	1, // 2: chat.ChatService.SayHelloStreamFromServer:input_type -> google.protobuf.Empty
	0, // 3: chat.ChatService.SayHelloStreamBidirectional:input_type -> recipe.Message
	1, // 4: chat.ChatService.WithoutParametersAndReturn:input_type -> google.protobuf.Empty
	0, // 5: chat.ChatService.SayHello:output_type -> recipe.Message
	0, // 6: chat.ChatService.SayHelloStreamFromClient:output_type -> recipe.Message
	0, // 7: chat.ChatService.SayHelloStreamFromServer:output_type -> recipe.Message
	0, // 8: chat.ChatService.SayHelloStreamBidirectional:output_type -> recipe.Message
	1, // 9: chat.ChatService.WithoutParametersAndReturn:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_chat_proto_init() }
func file_chat_chat_proto_init() {
	if File_chat_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_chat_proto_goTypes,
		DependencyIndexes: file_chat_chat_proto_depIdxs,
	}.Build()
	File_chat_chat_proto = out.File
	file_chat_chat_proto_rawDesc = nil
	file_chat_chat_proto_goTypes = nil
	file_chat_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	SayHello(ctx context.Context, in *recipe.Message, opts ...grpc.CallOption) (*recipe.Message, error)
	SayHelloStreamFromClient(ctx context.Context, opts ...grpc.CallOption) (ChatService_SayHelloStreamFromClientClient, error)
	SayHelloStreamFromServer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ChatService_SayHelloStreamFromServerClient, error)
	SayHelloStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (ChatService_SayHelloStreamBidirectionalClient, error)
	WithoutParametersAndReturn(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SayHello(ctx context.Context, in *recipe.Message, opts ...grpc.CallOption) (*recipe.Message, error) {
	out := new(recipe.Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SayHelloStreamFromClient(ctx context.Context, opts ...grpc.CallOption) (ChatService_SayHelloStreamFromClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/chat.ChatService/SayHelloStreamFromClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSayHelloStreamFromClientClient{stream}
	return x, nil
}

type ChatService_SayHelloStreamFromClientClient interface {
	Send(*recipe.Message) error
	CloseAndRecv() (*recipe.Message, error)
	grpc.ClientStream
}

type chatServiceSayHelloStreamFromClientClient struct {
	grpc.ClientStream
}

func (x *chatServiceSayHelloStreamFromClientClient) Send(m *recipe.Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSayHelloStreamFromClientClient) CloseAndRecv() (*recipe.Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(recipe.Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) SayHelloStreamFromServer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ChatService_SayHelloStreamFromServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[1], "/chat.ChatService/SayHelloStreamFromServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSayHelloStreamFromServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_SayHelloStreamFromServerClient interface {
	Recv() (*recipe.Message, error)
	grpc.ClientStream
}

type chatServiceSayHelloStreamFromServerClient struct {
	grpc.ClientStream
}

func (x *chatServiceSayHelloStreamFromServerClient) Recv() (*recipe.Message, error) {
	m := new(recipe.Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) SayHelloStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (ChatService_SayHelloStreamBidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[2], "/chat.ChatService/SayHelloStreamBidirectional", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSayHelloStreamBidirectionalClient{stream}
	return x, nil
}

type ChatService_SayHelloStreamBidirectionalClient interface {
	Send(*recipe.Message) error
	Recv() (*recipe.Message, error)
	grpc.ClientStream
}

type chatServiceSayHelloStreamBidirectionalClient struct {
	grpc.ClientStream
}

func (x *chatServiceSayHelloStreamBidirectionalClient) Send(m *recipe.Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSayHelloStreamBidirectionalClient) Recv() (*recipe.Message, error) {
	m := new(recipe.Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) WithoutParametersAndReturn(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chat.ChatService/WithoutParametersAndReturn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	SayHello(context.Context, *recipe.Message) (*recipe.Message, error)
	SayHelloStreamFromClient(ChatService_SayHelloStreamFromClientServer) error
	SayHelloStreamFromServer(*empty.Empty, ChatService_SayHelloStreamFromServerServer) error
	SayHelloStreamBidirectional(ChatService_SayHelloStreamBidirectionalServer) error
	WithoutParametersAndReturn(context.Context, *empty.Empty) (*empty.Empty, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) SayHello(context.Context, *recipe.Message) (*recipe.Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedChatServiceServer) SayHelloStreamFromClient(ChatService_SayHelloStreamFromClientServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStreamFromClient not implemented")
}
func (*UnimplementedChatServiceServer) SayHelloStreamFromServer(*empty.Empty, ChatService_SayHelloStreamFromServerServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStreamFromServer not implemented")
}
func (*UnimplementedChatServiceServer) SayHelloStreamBidirectional(ChatService_SayHelloStreamBidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStreamBidirectional not implemented")
}
func (*UnimplementedChatServiceServer) WithoutParametersAndReturn(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithoutParametersAndReturn not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(recipe.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SayHello(ctx, req.(*recipe.Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SayHelloStreamFromClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SayHelloStreamFromClient(&chatServiceSayHelloStreamFromClientServer{stream})
}

type ChatService_SayHelloStreamFromClientServer interface {
	SendAndClose(*recipe.Message) error
	Recv() (*recipe.Message, error)
	grpc.ServerStream
}

type chatServiceSayHelloStreamFromClientServer struct {
	grpc.ServerStream
}

func (x *chatServiceSayHelloStreamFromClientServer) SendAndClose(m *recipe.Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSayHelloStreamFromClientServer) Recv() (*recipe.Message, error) {
	m := new(recipe.Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_SayHelloStreamFromServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).SayHelloStreamFromServer(m, &chatServiceSayHelloStreamFromServerServer{stream})
}

type ChatService_SayHelloStreamFromServerServer interface {
	Send(*recipe.Message) error
	grpc.ServerStream
}

type chatServiceSayHelloStreamFromServerServer struct {
	grpc.ServerStream
}

func (x *chatServiceSayHelloStreamFromServerServer) Send(m *recipe.Message) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_SayHelloStreamBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SayHelloStreamBidirectional(&chatServiceSayHelloStreamBidirectionalServer{stream})
}

type ChatService_SayHelloStreamBidirectionalServer interface {
	Send(*recipe.Message) error
	Recv() (*recipe.Message, error)
	grpc.ServerStream
}

type chatServiceSayHelloStreamBidirectionalServer struct {
	grpc.ServerStream
}

func (x *chatServiceSayHelloStreamBidirectionalServer) Send(m *recipe.Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSayHelloStreamBidirectionalServer) Recv() (*recipe.Message, error) {
	m := new(recipe.Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_WithoutParametersAndReturn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).WithoutParametersAndReturn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/WithoutParametersAndReturn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).WithoutParametersAndReturn(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ChatService_SayHello_Handler,
		},
		{
			MethodName: "WithoutParametersAndReturn",
			Handler:    _ChatService_WithoutParametersAndReturn_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloStreamFromClient",
			Handler:       _ChatService_SayHelloStreamFromClient_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloStreamFromServer",
			Handler:       _ChatService_SayHelloStreamFromServer_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloStreamBidirectional",
			Handler:       _ChatService_SayHelloStreamBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat/chat.proto",
}
