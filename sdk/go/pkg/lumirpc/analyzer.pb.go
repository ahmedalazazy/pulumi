// Code generated by protoc-gen-go.
// source: analyzer.proto
// DO NOT EDIT!

/*
Package lumirpc is a generated protocol buffer package.

It is generated from these files:
	analyzer.proto
	engine.proto
	provider.proto

It has these top-level messages:
	AnalyzeRequest
	AnalyzeResponse
	AnalyzeFailure
	AnalyzeResourceRequest
	AnalyzeResourceResponse
	AnalyzeResourceFailure
	LogRequest
	CheckRequest
	CheckResponse
	CheckFailure
	NameRequest
	NameResponse
	CreateRequest
	CreateResponse
	GetRequest
	GetResponse
	InspectChangeRequest
	InspectChangeResponse
	UpdateRequest
	DeleteRequest
*/
package lumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

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

type AnalyzeRequest struct {
	Pkg string `protobuf:"bytes,1,opt,name=pkg" json:"pkg,omitempty"`
}

func (m *AnalyzeRequest) Reset()                    { *m = AnalyzeRequest{} }
func (m *AnalyzeRequest) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeRequest) ProtoMessage()               {}
func (*AnalyzeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AnalyzeRequest) GetPkg() string {
	if m != nil {
		return m.Pkg
	}
	return ""
}

type AnalyzeResponse struct {
	Failures []*AnalyzeFailure `protobuf:"bytes,1,rep,name=failures" json:"failures,omitempty"`
}

func (m *AnalyzeResponse) Reset()                    { *m = AnalyzeResponse{} }
func (m *AnalyzeResponse) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeResponse) ProtoMessage()               {}
func (*AnalyzeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AnalyzeResponse) GetFailures() []*AnalyzeFailure {
	if m != nil {
		return m.Failures
	}
	return nil
}

type AnalyzeFailure struct {
	Reason string `protobuf:"bytes,1,opt,name=reason" json:"reason,omitempty"`
}

func (m *AnalyzeFailure) Reset()                    { *m = AnalyzeFailure{} }
func (m *AnalyzeFailure) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeFailure) ProtoMessage()               {}
func (*AnalyzeFailure) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AnalyzeFailure) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type AnalyzeResourceRequest struct {
	Type       string                  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties *google_protobuf.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
	Unknowns   map[string]bool         `protobuf:"bytes,3,rep,name=unknowns" json:"unknowns,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *AnalyzeResourceRequest) Reset()                    { *m = AnalyzeResourceRequest{} }
func (m *AnalyzeResourceRequest) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeResourceRequest) ProtoMessage()               {}
func (*AnalyzeResourceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AnalyzeResourceRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AnalyzeResourceRequest) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *AnalyzeResourceRequest) GetUnknowns() map[string]bool {
	if m != nil {
		return m.Unknowns
	}
	return nil
}

type AnalyzeResourceResponse struct {
	Failures []*AnalyzeResourceFailure `protobuf:"bytes,1,rep,name=failures" json:"failures,omitempty"`
}

func (m *AnalyzeResourceResponse) Reset()                    { *m = AnalyzeResourceResponse{} }
func (m *AnalyzeResourceResponse) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeResourceResponse) ProtoMessage()               {}
func (*AnalyzeResourceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AnalyzeResourceResponse) GetFailures() []*AnalyzeResourceFailure {
	if m != nil {
		return m.Failures
	}
	return nil
}

type AnalyzeResourceFailure struct {
	Property string `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	Reason   string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *AnalyzeResourceFailure) Reset()                    { *m = AnalyzeResourceFailure{} }
func (m *AnalyzeResourceFailure) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeResourceFailure) ProtoMessage()               {}
func (*AnalyzeResourceFailure) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AnalyzeResourceFailure) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *AnalyzeResourceFailure) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*AnalyzeRequest)(nil), "lumirpc.AnalyzeRequest")
	proto.RegisterType((*AnalyzeResponse)(nil), "lumirpc.AnalyzeResponse")
	proto.RegisterType((*AnalyzeFailure)(nil), "lumirpc.AnalyzeFailure")
	proto.RegisterType((*AnalyzeResourceRequest)(nil), "lumirpc.AnalyzeResourceRequest")
	proto.RegisterType((*AnalyzeResourceResponse)(nil), "lumirpc.AnalyzeResourceResponse")
	proto.RegisterType((*AnalyzeResourceFailure)(nil), "lumirpc.AnalyzeResourceFailure")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Analyzer service

type AnalyzerClient interface {
	// Analyze analyzes an entire project/stack/snapshot, and returns any errors that it finds.
	Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error)
	// AnalyzeResource analyzes a single resource object, and returns any errors that it finds.
	AnalyzeResource(ctx context.Context, in *AnalyzeResourceRequest, opts ...grpc.CallOption) (*AnalyzeResourceResponse, error)
}

type analyzerClient struct {
	cc *grpc.ClientConn
}

func NewAnalyzerClient(cc *grpc.ClientConn) AnalyzerClient {
	return &analyzerClient{cc}
}

func (c *analyzerClient) Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error) {
	out := new(AnalyzeResponse)
	err := grpc.Invoke(ctx, "/lumirpc.Analyzer/Analyze", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) AnalyzeResource(ctx context.Context, in *AnalyzeResourceRequest, opts ...grpc.CallOption) (*AnalyzeResourceResponse, error) {
	out := new(AnalyzeResourceResponse)
	err := grpc.Invoke(ctx, "/lumirpc.Analyzer/AnalyzeResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Analyzer service

type AnalyzerServer interface {
	// Analyze analyzes an entire project/stack/snapshot, and returns any errors that it finds.
	Analyze(context.Context, *AnalyzeRequest) (*AnalyzeResponse, error)
	// AnalyzeResource analyzes a single resource object, and returns any errors that it finds.
	AnalyzeResource(context.Context, *AnalyzeResourceRequest) (*AnalyzeResourceResponse, error)
}

func RegisterAnalyzerServer(s *grpc.Server, srv AnalyzerServer) {
	s.RegisterService(&_Analyzer_serviceDesc, srv)
}

func _Analyzer_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.Analyzer/Analyze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).Analyze(ctx, req.(*AnalyzeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_AnalyzeResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).AnalyzeResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.Analyzer/AnalyzeResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).AnalyzeResource(ctx, req.(*AnalyzeResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Analyzer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lumirpc.Analyzer",
	HandlerType: (*AnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyze",
			Handler:    _Analyzer_Analyze_Handler,
		},
		{
			MethodName: "AnalyzeResource",
			Handler:    _Analyzer_AnalyzeResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analyzer.proto",
}

func init() { proto.RegisterFile("analyzer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xbe, 0xd3, 0xde, 0xdb, 0xe6, 0x9e, 0xcb, 0xad, 0x32, 0x48, 0x1b, 0x82, 0xd0, 0x30, 0xab,
	0x6e, 0x4c, 0xa1, 0x5d, 0x28, 0x16, 0x04, 0x17, 0x16, 0x04, 0x57, 0x11, 0xbb, 0x4f, 0xcb, 0xb4,
	0x94, 0xc4, 0x99, 0x71, 0x7e, 0x94, 0xf8, 0x38, 0xbe, 0x9f, 0xef, 0x20, 0x9d, 0x4c, 0x62, 0x6c,
	0x2c, 0xee, 0xce, 0xc9, 0xf9, 0xce, 0x7c, 0x3f, 0x27, 0xd0, 0x4b, 0x58, 0x92, 0xe5, 0xaf, 0x54,
	0x46, 0x42, 0x72, 0xcd, 0x71, 0x37, 0x33, 0x8f, 0x5b, 0x29, 0x56, 0xc1, 0xe9, 0x86, 0xf3, 0x4d,
	0x46, 0xc7, 0xf6, 0xf3, 0xd2, 0xac, 0xc7, 0x4a, 0x4b, 0xb3, 0xd2, 0x05, 0x8c, 0x10, 0xe8, 0x5d,
	0x17, 0x8b, 0x31, 0x7d, 0x32, 0x54, 0x69, 0x7c, 0x0c, 0x6d, 0x91, 0x6e, 0x7c, 0x14, 0xa2, 0xd1,
	0xdf, 0x78, 0x57, 0x92, 0x39, 0x1c, 0x55, 0x18, 0x25, 0x38, 0x53, 0x14, 0x4f, 0xc1, 0x5b, 0x27,
	0xdb, 0xcc, 0x48, 0xaa, 0x7c, 0x14, 0xb6, 0x47, 0xff, 0x26, 0x83, 0xc8, 0x11, 0x46, 0x0e, 0x3b,
	0x2f, 0xe6, 0x71, 0x05, 0x24, 0xa3, 0x8a, 0xcb, 0xcd, 0x70, 0x1f, 0x3a, 0x92, 0x26, 0x8a, 0x33,
	0x47, 0xe7, 0x3a, 0xf2, 0x8e, 0xa0, 0xff, 0x49, 0xc9, 0x8d, 0x5c, 0x55, 0xf2, 0x30, 0xfc, 0xd6,
	0xb9, 0xa0, 0x6e, 0xc1, 0xd6, 0xf8, 0x1c, 0x40, 0x48, 0x2e, 0xa8, 0xd4, 0x5b, 0xaa, 0xfc, 0x56,
	0x88, 0xac, 0x9e, 0xc2, 0x77, 0x54, 0xfa, 0x8e, 0xee, 0xad, 0xef, 0xb8, 0x06, 0xc5, 0xb7, 0xe0,
	0x19, 0x96, 0x32, 0xfe, 0xc2, 0x94, 0xdf, 0xb6, 0x36, 0xce, 0xf6, 0x6d, 0xec, 0xf1, 0x47, 0x0f,
	0x0e, 0x7f, 0xc3, 0xb4, 0xcc, 0xe3, 0x6a, 0x3d, 0x98, 0xc1, 0xff, 0x2f, 0xa3, 0x5d, 0x8e, 0x29,
	0xcd, 0xcb, 0x1c, 0x53, 0x9a, 0xe3, 0x13, 0xf8, 0xf3, 0x9c, 0x64, 0x86, 0x5a, 0x85, 0x5e, 0x5c,
	0x34, 0x97, 0xad, 0x0b, 0x44, 0x16, 0x30, 0x68, 0xd0, 0xb9, 0xa4, 0x67, 0x8d, 0xa4, 0x87, 0x87,
	0x24, 0x36, 0x13, 0xbf, 0x6b, 0xc4, 0x58, 0x26, 0x1f, 0x80, 0xe7, 0x72, 0x28, 0x25, 0x56, 0x7d,
	0xed, 0x2a, 0xad, 0xfa, 0x55, 0x26, 0x6f, 0x08, 0x3c, 0xf7, 0x9c, 0xc4, 0x57, 0xd0, 0x75, 0x35,
	0x1e, 0x34, 0x05, 0xd9, 0xac, 0x02, 0xff, 0x1b, 0xa5, 0xd6, 0x15, 0xf9, 0x85, 0x17, 0xf5, 0x9f,
	0xca, 0x4a, 0xc3, 0xc3, 0x1f, 0xb2, 0x0f, 0xc2, 0xc3, 0x80, 0xf2, 0xdd, 0x65, 0xc7, 0xde, 0x7b,
	0xfa, 0x11, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x75, 0x5c, 0x82, 0x10, 0x03, 0x00, 0x00,
}
