// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nomad/stream/proto/subscribe.proto

package pbstream

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type Topic int32

const (
	Topic_All        Topic = 0
	Topic_Deployment Topic = 1
	Topic_Evaluation Topic = 2
	Topic_Allocation Topic = 3
	Topic_Job        Topic = 4
	Topic_Node       Topic = 5
)

var Topic_name = map[int32]string{
	0: "All",
	1: "Deployment",
	2: "Evaluation",
	3: "Allocation",
	4: "Job",
	5: "Node",
}

var Topic_value = map[string]int32{
	"All":        0,
	"Deployment": 1,
	"Evaluation": 2,
	"Allocation": 3,
	"Job":        4,
	"Node":       5,
}

func (x Topic) String() string {
	return proto.EnumName(Topic_name, int32(x))
}

func (Topic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1fbaa8c5f6fb9626, []int{0}
}

type TopicFilter struct {
	Topic                Topic    `protobuf:"varint,1,opt,name=Topic,proto3,enum=hashicorp.nomad.stream.proto.Topic" json:"Topic,omitempty"`
	FilterKeys           []string `protobuf:"bytes,2,rep,name=FilterKeys,proto3" json:"FilterKeys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicFilter) Reset()         { *m = TopicFilter{} }
func (m *TopicFilter) String() string { return proto.CompactTextString(m) }
func (*TopicFilter) ProtoMessage()    {}
func (*TopicFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbaa8c5f6fb9626, []int{0}
}

func (m *TopicFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicFilter.Unmarshal(m, b)
}
func (m *TopicFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicFilter.Marshal(b, m, deterministic)
}
func (m *TopicFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicFilter.Merge(m, src)
}
func (m *TopicFilter) XXX_Size() int {
	return xxx_messageInfo_TopicFilter.Size(m)
}
func (m *TopicFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TopicFilter proto.InternalMessageInfo

func (m *TopicFilter) GetTopic() Topic {
	if m != nil {
		return m.Topic
	}
	return Topic_All
}

func (m *TopicFilter) GetFilterKeys() []string {
	if m != nil {
		return m.FilterKeys
	}
	return nil
}

type SubscribeRequest struct {
	Token                string         `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Index                uint64         `protobuf:"varint,2,opt,name=Index,proto3" json:"Index,omitempty"`
	Namespace            string         `protobuf:"bytes,3,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	Topics               []*TopicFilter `protobuf:"bytes,4,rep,name=Topics,proto3" json:"Topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbaa8c5f6fb9626, []int{1}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SubscribeRequest) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *SubscribeRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SubscribeRequest) GetTopics() []*TopicFilter {
	if m != nil {
		return m.Topics
	}
	return nil
}

type EventBatch struct {
	Index                uint64   `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Event                []*Event `protobuf:"bytes,2,rep,name=Event,proto3" json:"Event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventBatch) Reset()         { *m = EventBatch{} }
func (m *EventBatch) String() string { return proto.CompactTextString(m) }
func (*EventBatch) ProtoMessage()    {}
func (*EventBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbaa8c5f6fb9626, []int{2}
}

func (m *EventBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventBatch.Unmarshal(m, b)
}
func (m *EventBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventBatch.Marshal(b, m, deterministic)
}
func (m *EventBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBatch.Merge(m, src)
}
func (m *EventBatch) XXX_Size() int {
	return xxx_messageInfo_EventBatch.Size(m)
}
func (m *EventBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBatch.DiscardUnknown(m)
}

var xxx_messageInfo_EventBatch proto.InternalMessageInfo

func (m *EventBatch) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *EventBatch) GetEvent() []*Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type Event struct {
	Topic                Topic           `protobuf:"varint,1,opt,name=Topic,proto3,enum=hashicorp.nomad.stream.proto.Topic" json:"Topic,omitempty"`
	Type                 string          `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Key                  string          `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	Namespace            string          `protobuf:"bytes,4,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	FilterKeys           []string        `protobuf:"bytes,5,rep,name=FilterKeys,proto3" json:"FilterKeys,omitempty"`
	Index                uint64          `protobuf:"varint,6,opt,name=Index,proto3" json:"Index,omitempty"`
	Payload              *_struct.Struct `protobuf:"bytes,7,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbaa8c5f6fb9626, []int{3}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetTopic() Topic {
	if m != nil {
		return m.Topic
	}
	return Topic_All
}

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Event) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Event) GetFilterKeys() []string {
	if m != nil {
		return m.FilterKeys
	}
	return nil
}

func (m *Event) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Event) GetPayload() *_struct.Struct {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("hashicorp.nomad.stream.proto.Topic", Topic_name, Topic_value)
	proto.RegisterType((*TopicFilter)(nil), "hashicorp.nomad.stream.proto.TopicFilter")
	proto.RegisterType((*SubscribeRequest)(nil), "hashicorp.nomad.stream.proto.SubscribeRequest")
	proto.RegisterType((*EventBatch)(nil), "hashicorp.nomad.stream.proto.EventBatch")
	proto.RegisterType((*Event)(nil), "hashicorp.nomad.stream.proto.Event")
}

func init() {
	proto.RegisterFile("nomad/stream/proto/subscribe.proto", fileDescriptor_1fbaa8c5f6fb9626)
}

var fileDescriptor_1fbaa8c5f6fb9626 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x6b, 0xdb, 0x30,
	0x18, 0xad, 0x63, 0x27, 0xa9, 0x3f, 0x43, 0x31, 0x62, 0x30, 0x53, 0xc2, 0x30, 0xde, 0xc5, 0xdb,
	0x41, 0xd9, 0xb2, 0x53, 0x8f, 0x29, 0xeb, 0x60, 0x2b, 0x94, 0xa1, 0xf4, 0x34, 0xd8, 0x41, 0x76,
	0xbe, 0x35, 0x66, 0x8a, 0xe5, 0xd9, 0x72, 0xa9, 0x7f, 0xcd, 0x7e, 0xe2, 0xfe, 0xc2, 0xf0, 0xa7,
	0x64, 0x49, 0x73, 0xe8, 0x02, 0xbd, 0xe9, 0x49, 0xef, 0x49, 0xef, 0x7b, 0x7a, 0x90, 0x94, 0x7a,
	0x2d, 0x97, 0xd3, 0xc6, 0xd4, 0x28, 0xd7, 0xd3, 0xaa, 0xd6, 0x46, 0x4f, 0x9b, 0x36, 0x6b, 0xf2,
	0xba, 0xc8, 0x90, 0x13, 0x66, 0x93, 0x95, 0x6c, 0x56, 0x45, 0xae, 0xeb, 0x8a, 0x13, 0x9b, 0x5b,
	0xb6, 0x3d, 0x3d, 0x9f, 0xdc, 0x69, 0x7d, 0xa7, 0xd0, 0x6a, 0xb3, 0xf6, 0x47, 0x7f, 0x57, 0x9b,
	0x1b, 0x7b, 0x9a, 0xac, 0x20, 0xb8, 0xd5, 0x55, 0x91, 0x7f, 0x2a, 0x94, 0xc1, 0x9a, 0x5d, 0xc0,
	0x90, 0x60, 0xe4, 0xc4, 0x4e, 0x7a, 0x36, 0x7b, 0xcd, 0x9f, 0xba, 0x9a, 0x13, 0x55, 0x58, 0x05,
	0x7b, 0x05, 0x60, 0x2f, 0xb9, 0xc6, 0xae, 0x89, 0x06, 0xb1, 0x9b, 0xfa, 0x62, 0x6f, 0x27, 0xf9,
	0xed, 0x40, 0xb8, 0xd8, 0x3a, 0x17, 0xf8, 0xab, 0xc5, 0xc6, 0xb0, 0x17, 0xfd, 0x7b, 0x3f, 0xb1,
	0xa4, 0xf7, 0x7c, 0x61, 0x41, 0xbf, 0xfb, 0xb9, 0x5c, 0xe2, 0x43, 0x34, 0x88, 0x9d, 0xd4, 0x13,
	0x16, 0xb0, 0x09, 0xf8, 0x37, 0x72, 0x8d, 0x4d, 0x25, 0x73, 0x8c, 0x5c, 0xe2, 0xef, 0x36, 0xd8,
	0x1c, 0x46, 0xe4, 0xa3, 0x89, 0xbc, 0xd8, 0x4d, 0x83, 0xd9, 0x9b, 0x23, 0xac, 0x5b, 0x77, 0x62,
	0x23, 0x4c, 0xbe, 0x03, 0x5c, 0xdd, 0x63, 0x69, 0x2e, 0xa5, 0xc9, 0x57, 0x3b, 0x13, 0xce, 0xbe,
	0x89, 0x0b, 0x18, 0x12, 0x87, 0x06, 0x0c, 0xfe, 0x17, 0x10, 0x51, 0x85, 0x55, 0x24, 0x7f, 0x9c,
	0x8d, 0xf6, 0x39, 0x29, 0x33, 0xf0, 0x6e, 0xbb, 0x0a, 0x29, 0x19, 0x5f, 0xd0, 0x9a, 0x85, 0xe0,
	0x5e, 0x63, 0xb7, 0x89, 0xa4, 0x5f, 0x3e, 0x8e, 0xca, 0x3b, 0x8c, 0xea, 0xf1, 0x4f, 0x0d, 0x0f,
	0x7f, 0x6a, 0x37, 0xf9, 0x68, 0x7f, 0xf2, 0xf7, 0x30, 0xfe, 0x2a, 0x3b, 0xa5, 0xe5, 0x32, 0x1a,
	0xc7, 0x4e, 0x1a, 0xcc, 0x5e, 0x72, 0xdb, 0x2c, 0xbe, 0x6d, 0x16, 0x5f, 0x50, 0xb3, 0xc4, 0x96,
	0xf7, 0x76, 0xb1, 0x99, 0x93, 0x8d, 0xc1, 0x9d, 0x2b, 0x15, 0x9e, 0xb0, 0x33, 0x80, 0x8f, 0x58,
	0x29, 0xdd, 0xad, 0xb1, 0x34, 0xa1, 0xd3, 0xe3, 0xab, 0x7b, 0xa9, 0x5a, 0x69, 0x0a, 0x5d, 0x86,
	0x83, 0x1e, 0xcf, 0x95, 0xd2, 0xb9, 0xc5, 0x6e, 0x2f, 0xfc, 0xa2, 0xb3, 0xd0, 0x63, 0xa7, 0xe0,
	0xdd, 0xe8, 0x25, 0x86, 0xc3, 0xd9, 0x03, 0x04, 0x94, 0xe2, 0x82, 0x32, 0x62, 0x05, 0xf8, 0xff,
	0x5a, 0xc5, 0xf8, 0xd3, 0x49, 0x1e, 0xd6, 0xef, 0x3c, 0x3d, 0xe2, 0xfb, 0xa8, 0x0d, 0xc9, 0xc9,
	0x3b, 0xe7, 0x12, 0xbe, 0x9d, 0x56, 0x99, 0x25, 0x64, 0x23, 0x62, 0x7c, 0xf8, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0xe9, 0xee, 0x6e, 0xfe, 0xa0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventStreamClient is the client API for EventStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventStreamClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (EventStream_SubscribeClient, error)
}

type eventStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewEventStreamClient(cc grpc.ClientConnInterface) EventStreamClient {
	return &eventStreamClient{cc}
}

func (c *eventStreamClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (EventStream_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventStream_serviceDesc.Streams[0], "/hashicorp.nomad.stream.proto.EventStream/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventStreamSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventStream_SubscribeClient interface {
	Recv() (*EventBatch, error)
	grpc.ClientStream
}

type eventStreamSubscribeClient struct {
	grpc.ClientStream
}

func (x *eventStreamSubscribeClient) Recv() (*EventBatch, error) {
	m := new(EventBatch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventStreamServer is the server API for EventStream service.
type EventStreamServer interface {
	Subscribe(*SubscribeRequest, EventStream_SubscribeServer) error
}

// UnimplementedEventStreamServer can be embedded to have forward compatible implementations.
type UnimplementedEventStreamServer struct {
}

func (*UnimplementedEventStreamServer) Subscribe(req *SubscribeRequest, srv EventStream_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterEventStreamServer(s *grpc.Server, srv EventStreamServer) {
	s.RegisterService(&_EventStream_serviceDesc, srv)
}

func _EventStream_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventStreamServer).Subscribe(m, &eventStreamSubscribeServer{stream})
}

type EventStream_SubscribeServer interface {
	Send(*EventBatch) error
	grpc.ServerStream
}

type eventStreamSubscribeServer struct {
	grpc.ServerStream
}

func (x *eventStreamSubscribeServer) Send(m *EventBatch) error {
	return x.ServerStream.SendMsg(m)
}

var _EventStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad.stream.proto.EventStream",
	HandlerType: (*EventStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _EventStream_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "nomad/stream/proto/subscribe.proto",
}
