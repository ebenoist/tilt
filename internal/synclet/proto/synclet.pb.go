// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/windmilleng/tilt/internal/synclet/synclet.proto

package proto // import "github.com/windmilleng/tilt/internal/synclet/proto"

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogLevel int32

const (
	LogLevel_INFO    LogLevel = 0
	LogLevel_VERBOSE LogLevel = 1
	LogLevel_DEBUG   LogLevel = 2
)

var LogLevel_name = map[int32]string{
	0: "INFO",
	1: "VERBOSE",
	2: "DEBUG",
}
var LogLevel_value = map[string]int32{
	"INFO":    0,
	"VERBOSE": 1,
	"DEBUG":   2,
}

func (x LogLevel) String() string {
	return proto.EnumName(LogLevel_name, int32(x))
}
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_synclet_4ecdb7569c74f134, []int{0}
}

type Cmd struct {
	Argv                 []string `protobuf:"bytes,1,rep,name=argv,proto3" json:"argv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cmd) Reset()         { *m = Cmd{} }
func (m *Cmd) String() string { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()    {}
func (*Cmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_synclet_4ecdb7569c74f134, []int{0}
}
func (m *Cmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cmd.Unmarshal(m, b)
}
func (m *Cmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cmd.Marshal(b, m, deterministic)
}
func (dst *Cmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cmd.Merge(dst, src)
}
func (m *Cmd) XXX_Size() int {
	return xxx_messageInfo_Cmd.Size(m)
}
func (m *Cmd) XXX_DiscardUnknown() {
	xxx_messageInfo_Cmd.DiscardUnknown(m)
}

var xxx_messageInfo_Cmd proto.InternalMessageInfo

func (m *Cmd) GetArgv() []string {
	if m != nil {
		return m.Argv
	}
	return nil
}

type UpdateContainerRequest struct {
	ContainerId          string    `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	TarArchive           []byte    `protobuf:"bytes,2,opt,name=tar_archive,json=tarArchive,proto3" json:"tar_archive,omitempty"`
	FilesToDelete        []string  `protobuf:"bytes,3,rep,name=files_to_delete,json=filesToDelete,proto3" json:"files_to_delete,omitempty"`
	Commands             []*Cmd    `protobuf:"bytes,4,rep,name=commands,proto3" json:"commands,omitempty"`
	LogStyle             *LogStyle `protobuf:"bytes,5,opt,name=log_style,json=logStyle,proto3" json:"log_style,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateContainerRequest) Reset()         { *m = UpdateContainerRequest{} }
func (m *UpdateContainerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateContainerRequest) ProtoMessage()    {}
func (*UpdateContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_synclet_4ecdb7569c74f134, []int{1}
}
func (m *UpdateContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateContainerRequest.Unmarshal(m, b)
}
func (m *UpdateContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateContainerRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateContainerRequest.Merge(dst, src)
}
func (m *UpdateContainerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateContainerRequest.Size(m)
}
func (m *UpdateContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateContainerRequest proto.InternalMessageInfo

func (m *UpdateContainerRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *UpdateContainerRequest) GetTarArchive() []byte {
	if m != nil {
		return m.TarArchive
	}
	return nil
}

func (m *UpdateContainerRequest) GetFilesToDelete() []string {
	if m != nil {
		return m.FilesToDelete
	}
	return nil
}

func (m *UpdateContainerRequest) GetCommands() []*Cmd {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *UpdateContainerRequest) GetLogStyle() *LogStyle {
	if m != nil {
		return m.LogStyle
	}
	return nil
}

type LogMessage struct {
	Level                LogLevel `protobuf:"varint,1,opt,name=level,proto3,enum=synclet.LogLevel" json:"level,omitempty"`
	Message              []byte   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogMessage) Reset()         { *m = LogMessage{} }
func (m *LogMessage) String() string { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()    {}
func (*LogMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_synclet_4ecdb7569c74f134, []int{2}
}
func (m *LogMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogMessage.Unmarshal(m, b)
}
func (m *LogMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogMessage.Marshal(b, m, deterministic)
}
func (dst *LogMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogMessage.Merge(dst, src)
}
func (m *LogMessage) XXX_Size() int {
	return xxx_messageInfo_LogMessage.Size(m)
}
func (m *LogMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LogMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LogMessage proto.InternalMessageInfo

func (m *LogMessage) GetLevel() LogLevel {
	if m != nil {
		return m.Level
	}
	return LogLevel_INFO
}

func (m *LogMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type UpdateContainerReply struct {
	LogMessage           *LogMessage `protobuf:"bytes,1,opt,name=log_message,json=logMessage,proto3" json:"log_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateContainerReply) Reset()         { *m = UpdateContainerReply{} }
func (m *UpdateContainerReply) String() string { return proto.CompactTextString(m) }
func (*UpdateContainerReply) ProtoMessage()    {}
func (*UpdateContainerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_synclet_4ecdb7569c74f134, []int{3}
}
func (m *UpdateContainerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateContainerReply.Unmarshal(m, b)
}
func (m *UpdateContainerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateContainerReply.Marshal(b, m, deterministic)
}
func (dst *UpdateContainerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateContainerReply.Merge(dst, src)
}
func (m *UpdateContainerReply) XXX_Size() int {
	return xxx_messageInfo_UpdateContainerReply.Size(m)
}
func (m *UpdateContainerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateContainerReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateContainerReply proto.InternalMessageInfo

func (m *UpdateContainerReply) GetLogMessage() *LogMessage {
	if m != nil {
		return m.LogMessage
	}
	return nil
}

type LogStyle struct {
	ColorsEnabled        bool     `protobuf:"varint,1,opt,name=colors_enabled,json=colorsEnabled,proto3" json:"colors_enabled,omitempty"`
	Level                LogLevel `protobuf:"varint,2,opt,name=level,proto3,enum=synclet.LogLevel" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogStyle) Reset()         { *m = LogStyle{} }
func (m *LogStyle) String() string { return proto.CompactTextString(m) }
func (*LogStyle) ProtoMessage()    {}
func (*LogStyle) Descriptor() ([]byte, []int) {
	return fileDescriptor_synclet_4ecdb7569c74f134, []int{4}
}
func (m *LogStyle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogStyle.Unmarshal(m, b)
}
func (m *LogStyle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogStyle.Marshal(b, m, deterministic)
}
func (dst *LogStyle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogStyle.Merge(dst, src)
}
func (m *LogStyle) XXX_Size() int {
	return xxx_messageInfo_LogStyle.Size(m)
}
func (m *LogStyle) XXX_DiscardUnknown() {
	xxx_messageInfo_LogStyle.DiscardUnknown(m)
}

var xxx_messageInfo_LogStyle proto.InternalMessageInfo

func (m *LogStyle) GetColorsEnabled() bool {
	if m != nil {
		return m.ColorsEnabled
	}
	return false
}

func (m *LogStyle) GetLevel() LogLevel {
	if m != nil {
		return m.Level
	}
	return LogLevel_INFO
}

func init() {
	proto.RegisterType((*Cmd)(nil), "synclet.Cmd")
	proto.RegisterType((*UpdateContainerRequest)(nil), "synclet.UpdateContainerRequest")
	proto.RegisterType((*LogMessage)(nil), "synclet.LogMessage")
	proto.RegisterType((*UpdateContainerReply)(nil), "synclet.UpdateContainerReply")
	proto.RegisterType((*LogStyle)(nil), "synclet.LogStyle")
	proto.RegisterEnum("synclet.LogLevel", LogLevel_name, LogLevel_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SyncletClient is the client API for Synclet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncletClient interface {
	// updates the specified container and then restarts it
	// (much functionality packed into one rpc to minimize latency)
	UpdateContainer(ctx context.Context, in *UpdateContainerRequest, opts ...grpc.CallOption) (Synclet_UpdateContainerClient, error)
}

type syncletClient struct {
	cc *grpc.ClientConn
}

func NewSyncletClient(cc *grpc.ClientConn) SyncletClient {
	return &syncletClient{cc}
}

func (c *syncletClient) UpdateContainer(ctx context.Context, in *UpdateContainerRequest, opts ...grpc.CallOption) (Synclet_UpdateContainerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Synclet_serviceDesc.Streams[0], "/synclet.Synclet/UpdateContainer", opts...)
	if err != nil {
		return nil, err
	}
	x := &syncletUpdateContainerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Synclet_UpdateContainerClient interface {
	Recv() (*UpdateContainerReply, error)
	grpc.ClientStream
}

type syncletUpdateContainerClient struct {
	grpc.ClientStream
}

func (x *syncletUpdateContainerClient) Recv() (*UpdateContainerReply, error) {
	m := new(UpdateContainerReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SyncletServer is the server API for Synclet service.
type SyncletServer interface {
	// updates the specified container and then restarts it
	// (much functionality packed into one rpc to minimize latency)
	UpdateContainer(*UpdateContainerRequest, Synclet_UpdateContainerServer) error
}

func RegisterSyncletServer(s *grpc.Server, srv SyncletServer) {
	s.RegisterService(&_Synclet_serviceDesc, srv)
}

func _Synclet_UpdateContainer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpdateContainerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyncletServer).UpdateContainer(m, &syncletUpdateContainerServer{stream})
}

type Synclet_UpdateContainerServer interface {
	Send(*UpdateContainerReply) error
	grpc.ServerStream
}

type syncletUpdateContainerServer struct {
	grpc.ServerStream
}

func (x *syncletUpdateContainerServer) Send(m *UpdateContainerReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Synclet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "synclet.Synclet",
	HandlerType: (*SyncletServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateContainer",
			Handler:       _Synclet_UpdateContainer_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/windmilleng/tilt/internal/synclet/synclet.proto",
}

func init() {
	proto.RegisterFile("github.com/windmilleng/tilt/internal/synclet/synclet.proto", fileDescriptor_synclet_4ecdb7569c74f134)
}

var fileDescriptor_synclet_4ecdb7569c74f134 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0xe7, 0xfe, 0xa0, 0xe9, 0xa5, 0xdb, 0x8a, 0x41, 0x28, 0x20, 0xa1, 0x85, 0x48, 0x40,
	0x84, 0x50, 0x8b, 0x0a, 0x0f, 0x68, 0x3c, 0xd1, 0xae, 0xa0, 0x49, 0x85, 0x4a, 0x2e, 0xe5, 0x61,
	0x2f, 0xc1, 0x4d, 0x8e, 0xcc, 0x92, 0x63, 0x97, 0xc4, 0x2b, 0xea, 0x3f, 0xcb, 0xdf, 0x82, 0xea,
	0x34, 0xe1, 0x37, 0xd2, 0x5e, 0xec, 0xf3, 0xe7, 0x7c, 0xbe, 0xef, 0xd7, 0x3a, 0x38, 0x4d, 0x85,
	0xb9, 0xbc, 0x5a, 0x0d, 0x62, 0x9d, 0x0d, 0xbf, 0x0a, 0x95, 0x64, 0x42, 0x4a, 0x54, 0xe9, 0xd0,
	0x08, 0x69, 0x86, 0x42, 0x19, 0xcc, 0x15, 0x97, 0xc3, 0x62, 0xab, 0x62, 0x89, 0xa6, 0xda, 0x07,
	0xeb, 0x5c, 0x1b, 0x4d, 0x3b, 0xfb, 0x63, 0x70, 0x17, 0x9a, 0x93, 0x2c, 0xa1, 0x14, 0x5a, 0x3c,
	0x4f, 0x37, 0x1e, 0xf1, 0x9b, 0x61, 0x97, 0xd9, 0x38, 0xf8, 0x46, 0xe0, 0xce, 0x72, 0x9d, 0x70,
	0x83, 0x13, 0xad, 0x0c, 0x17, 0x0a, 0x73, 0x86, 0x5f, 0xae, 0xb0, 0x30, 0xf4, 0x01, 0xf4, 0xe2,
	0x8a, 0x45, 0x22, 0xf1, 0x88, 0x4f, 0xc2, 0x2e, 0x73, 0x6b, 0x76, 0x9e, 0xd0, 0x13, 0x70, 0x0d,
	0xcf, 0x23, 0x9e, 0xc7, 0x97, 0x62, 0x83, 0x5e, 0xc3, 0x27, 0x61, 0x8f, 0x81, 0xe1, 0xf9, 0xeb,
	0x92, 0xd0, 0x47, 0x70, 0xfc, 0x59, 0x48, 0x2c, 0x22, 0xa3, 0xa3, 0x04, 0x25, 0x1a, 0xf4, 0x9a,
	0xb6, 0xfb, 0xa1, 0xc5, 0x1f, 0xf4, 0x99, 0x85, 0x34, 0x04, 0x27, 0xd6, 0x59, 0xc6, 0x55, 0x52,
	0x78, 0x2d, 0xbf, 0x19, 0xba, 0xa3, 0xde, 0xa0, 0x32, 0x33, 0xc9, 0x12, 0x56, 0x67, 0xe9, 0x00,
	0xba, 0x52, 0xa7, 0x51, 0x61, 0xb6, 0x12, 0xbd, 0xb6, 0x4f, 0x42, 0x77, 0x74, 0xb3, 0xbe, 0x3a,
	0xd3, 0xe9, 0x62, 0x97, 0x60, 0x8e, 0xdc, 0x47, 0xc1, 0x1c, 0x60, 0xa6, 0xd3, 0x77, 0x58, 0x14,
	0x3c, 0x45, 0xfa, 0x18, 0xda, 0x12, 0x37, 0x28, 0xad, 0x99, 0xa3, 0x5f, 0x2b, 0x67, 0xbb, 0x04,
	0x2b, 0xf3, 0xd4, 0x83, 0x4e, 0x56, 0xd6, 0xec, 0x5d, 0x55, 0xc7, 0x60, 0x06, 0xb7, 0xff, 0xf8,
	0xb0, 0xb5, 0xdc, 0xd2, 0x17, 0xe0, 0xee, 0x84, 0x55, 0x55, 0xc4, 0x4a, 0xbb, 0xf5, 0x73, 0x83,
	0xbd, 0x08, 0x06, 0xb2, 0x8e, 0x83, 0x0b, 0x70, 0x2a, 0xd1, 0xf4, 0x21, 0x1c, 0xc5, 0x5a, 0xea,
	0xbc, 0x88, 0x50, 0xf1, 0x95, 0xc4, 0xf2, 0xcb, 0x1d, 0x76, 0x58, 0xd2, 0x69, 0x09, 0x7f, 0x78,
	0x68, 0xfc, 0xdf, 0xc3, 0x93, 0xa7, 0xf6, 0x6d, 0x8b, 0xa8, 0x03, 0xad, 0xf3, 0xf7, 0x6f, 0xe6,
	0xfd, 0x03, 0xea, 0x42, 0xe7, 0xe3, 0x94, 0x8d, 0xe7, 0x8b, 0x69, 0x9f, 0xd0, 0x2e, 0xb4, 0xcf,
	0xa6, 0xe3, 0xe5, 0xdb, 0x7e, 0x63, 0xf4, 0x09, 0x3a, 0x8b, 0xf2, 0x21, 0xba, 0x84, 0xe3, 0xdf,
	0x2c, 0xd2, 0x93, 0xba, 0xcb, 0xdf, 0xa7, 0xe5, 0xde, 0xfd, 0x7f, 0x5f, 0x58, 0xcb, 0x6d, 0x70,
	0xf0, 0x8c, 0x8c, 0x4f, 0x2f, 0x5e, 0x5e, 0x6b, 0x9a, 0xed, 0x14, 0xbf, 0xb2, 0xeb, 0xea, 0x86,
	0xdd, 0x9e, 0x7f, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x15, 0xd9, 0xf3, 0xab, 0x10, 0x03, 0x00, 0x00,
}
