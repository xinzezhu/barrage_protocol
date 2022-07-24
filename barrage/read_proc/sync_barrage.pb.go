// Code generated by protoc-gen-go.
// source: sync_barrage.proto
// DO NOT EDIT!

/*
Package read_proc is a generated protocol buffer package.

It is generated from these files:
	sync_barrage.proto

It has these top-level messages:
	Request
	Response
	Header
	BarragePerSecond
	Body
*/
package read_proc

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

type Request struct {
	Header *Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Body   *Body   `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetBody() *Body {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	Header *Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Body   *Body   `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() *Body {
	if m != nil {
		return m.Body
	}
	return nil
}

type Header struct {
	RoomId uint32 `protobuf:"varint,1,opt,name=room_id,json=roomId" json:"room_id,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Header) GetRoomId() uint32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

type BarragePerSecond struct {
	Content []string `protobuf:"bytes,1,rep,name=content" json:"content,omitempty"`
}

func (m *BarragePerSecond) Reset()                    { *m = BarragePerSecond{} }
func (m *BarragePerSecond) String() string            { return proto.CompactTextString(m) }
func (*BarragePerSecond) ProtoMessage()               {}
func (*BarragePerSecond) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BarragePerSecond) GetContent() []string {
	if m != nil {
		return m.Content
	}
	return nil
}

type Body struct {
	Barrage []*BarragePerSecond `protobuf:"bytes,1,rep,name=barrage" json:"barrage,omitempty"`
}

func (m *Body) Reset()                    { *m = Body{} }
func (m *Body) String() string            { return proto.CompactTextString(m) }
func (*Body) ProtoMessage()               {}
func (*Body) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Body) GetBarrage() []*BarragePerSecond {
	if m != nil {
		return m.Barrage
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "read_proc.Request")
	proto.RegisterType((*Response)(nil), "read_proc.Response")
	proto.RegisterType((*Header)(nil), "read_proc.Header")
	proto.RegisterType((*BarragePerSecond)(nil), "read_proc.BarragePerSecond")
	proto.RegisterType((*Body)(nil), "read_proc.Body")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ReadProc service

type ReadProcClient interface {
	// 定义SayHello方法
	SyncBarrage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type readProcClient struct {
	cc *grpc.ClientConn
}

func NewReadProcClient(cc *grpc.ClientConn) ReadProcClient {
	return &readProcClient{cc}
}

func (c *readProcClient) SyncBarrage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/read_proc.ReadProc/SyncBarrage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReadProc service

type ReadProcServer interface {
	// 定义SayHello方法
	SyncBarrage(context.Context, *Request) (*Response, error)
}

func RegisterReadProcServer(s *grpc.Server, srv ReadProcServer) {
	s.RegisterService(&_ReadProc_serviceDesc, srv)
}

func _ReadProc_SyncBarrage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadProcServer).SyncBarrage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/read_proc.ReadProc/SyncBarrage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadProcServer).SyncBarrage(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReadProc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "read_proc.ReadProc",
	HandlerType: (*ReadProcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncBarrage",
			Handler:    _ReadProc_SyncBarrage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sync_barrage.proto",
}

func init() { proto.RegisterFile("sync_barrage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x91, 0x31, 0x4f, 0x84, 0x40,
	0x10, 0x46, 0x45, 0x2f, 0xe0, 0x0d, 0x31, 0xea, 0x58, 0x48, 0xb4, 0x39, 0xb1, 0x39, 0x13, 0x43,
	0x81, 0x31, 0xb1, 0xb1, 0x21, 0x16, 0xda, 0x5d, 0xf6, 0x2a, 0xaf, 0x21, 0xcb, 0xee, 0x44, 0x2d,
	0xdc, 0xc1, 0x01, 0x0b, 0xfe, 0xbd, 0x61, 0xe1, 0x14, 0xad, 0x2d, 0x67, 0xe7, 0xed, 0xcb, 0xcc,
	0x37, 0x80, 0x4d, 0xe7, 0x4c, 0x59, 0x69, 0x11, 0xfd, 0x42, 0x59, 0x2d, 0xdc, 0x32, 0xce, 0x85,
	0xb4, 0x2d, 0x6b, 0x61, 0x93, 0x3e, 0x43, 0xa4, 0xe8, 0xe3, 0x93, 0x9a, 0x16, 0xaf, 0x20, 0x7c,
	0x25, 0x6d, 0x49, 0x92, 0x60, 0x11, 0x2c, 0xe3, 0xfc, 0x38, 0xfb, 0xc6, 0xb2, 0x47, 0xdf, 0x50,
	0x23, 0x80, 0x97, 0x30, 0xab, 0xd8, 0x76, 0xc9, 0xae, 0x07, 0x0f, 0x27, 0x60, 0xc1, 0xb6, 0x53,
	0xbe, 0x99, 0x6e, 0x60, 0x5f, 0x51, 0x53, 0xb3, 0x6b, 0xe8, 0xdf, 0xdd, 0x17, 0x10, 0x0e, 0xdf,
	0xf0, 0x14, 0x22, 0x61, 0x7e, 0x2f, 0xdf, 0xac, 0x57, 0x1f, 0xa8, 0xb0, 0x2f, 0x9f, 0x6c, 0x7a,
	0x0d, 0x47, 0xc5, 0xb0, 0xf5, 0x8a, 0x64, 0x4d, 0x86, 0x9d, 0xc5, 0x04, 0x22, 0xc3, 0xae, 0x25,
	0xd7, 0x26, 0xc1, 0x62, 0x6f, 0x39, 0x57, 0xdb, 0x32, 0xbd, 0x87, 0x59, 0xaf, 0xc7, 0x5b, 0x88,
	0xc6, 0xac, 0x3c, 0x11, 0xe7, 0xe7, 0xd3, 0x01, 0xfe, 0xf8, 0xd4, 0x96, 0xcd, 0x1f, 0xfa, 0x5d,
	0xb5, 0x5d, 0x09, 0x1b, 0xbc, 0x83, 0x78, 0xdd, 0x39, 0x33, 0xc2, 0x88, 0x13, 0xc1, 0x18, 0xf5,
	0xd9, 0xc9, 0xaf, 0xb7, 0x21, 0xa3, 0x74, 0xa7, 0x88, 0x37, 0x3f, 0x97, 0xa9, 0x42, 0x7f, 0xab,
	0x9b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xa8, 0x89, 0xbe, 0xc1, 0x01, 0x00, 0x00,
}
