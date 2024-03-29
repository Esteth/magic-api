// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GetWaitTimesRequest struct {
	ParkId               string   `protobuf:"bytes,2,opt,name=park_id,json=parkId,proto3" json:"park_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWaitTimesRequest) Reset()         { *m = GetWaitTimesRequest{} }
func (m *GetWaitTimesRequest) String() string { return proto.CompactTextString(m) }
func (*GetWaitTimesRequest) ProtoMessage()    {}
func (*GetWaitTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *GetWaitTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWaitTimesRequest.Unmarshal(m, b)
}
func (m *GetWaitTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWaitTimesRequest.Marshal(b, m, deterministic)
}
func (m *GetWaitTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWaitTimesRequest.Merge(m, src)
}
func (m *GetWaitTimesRequest) XXX_Size() int {
	return xxx_messageInfo_GetWaitTimesRequest.Size(m)
}
func (m *GetWaitTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWaitTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWaitTimesRequest proto.InternalMessageInfo

func (m *GetWaitTimesRequest) GetParkId() string {
	if m != nil {
		return m.ParkId
	}
	return ""
}

type GetWaitHistoryRequest struct {
	AttractionId         string   `protobuf:"bytes,1,opt,name=attraction_id,json=attractionId,proto3" json:"attraction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWaitHistoryRequest) Reset()         { *m = GetWaitHistoryRequest{} }
func (m *GetWaitHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetWaitHistoryRequest) ProtoMessage()    {}
func (*GetWaitHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *GetWaitHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWaitHistoryRequest.Unmarshal(m, b)
}
func (m *GetWaitHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWaitHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetWaitHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWaitHistoryRequest.Merge(m, src)
}
func (m *GetWaitHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetWaitHistoryRequest.Size(m)
}
func (m *GetWaitHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWaitHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWaitHistoryRequest proto.InternalMessageInfo

func (m *GetWaitHistoryRequest) GetAttractionId() string {
	if m != nil {
		return m.AttractionId
	}
	return ""
}

type GetWaitTimesResponse struct {
	WaitTime             []*WaitTime `protobuf:"bytes,1,rep,name=wait_time,json=waitTime,proto3" json:"wait_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetWaitTimesResponse) Reset()         { *m = GetWaitTimesResponse{} }
func (m *GetWaitTimesResponse) String() string { return proto.CompactTextString(m) }
func (*GetWaitTimesResponse) ProtoMessage()    {}
func (*GetWaitTimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *GetWaitTimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWaitTimesResponse.Unmarshal(m, b)
}
func (m *GetWaitTimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWaitTimesResponse.Marshal(b, m, deterministic)
}
func (m *GetWaitTimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWaitTimesResponse.Merge(m, src)
}
func (m *GetWaitTimesResponse) XXX_Size() int {
	return xxx_messageInfo_GetWaitTimesResponse.Size(m)
}
func (m *GetWaitTimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWaitTimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetWaitTimesResponse proto.InternalMessageInfo

func (m *GetWaitTimesResponse) GetWaitTime() []*WaitTime {
	if m != nil {
		return m.WaitTime
	}
	return nil
}

type WaitTime struct {
	ParkId               string               `protobuf:"bytes,1,opt,name=park_id,json=parkId,proto3" json:"park_id,omitempty"`
	AttractionId         string               `protobuf:"bytes,2,opt,name=attraction_id,json=attractionId,proto3" json:"attraction_id,omitempty"`
	WaitTime             int32                `protobuf:"varint,3,opt,name=wait_time,json=waitTime,proto3" json:"wait_time,omitempty"`
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WaitTime) Reset()         { *m = WaitTime{} }
func (m *WaitTime) String() string { return proto.CompactTextString(m) }
func (*WaitTime) ProtoMessage()    {}
func (*WaitTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *WaitTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitTime.Unmarshal(m, b)
}
func (m *WaitTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitTime.Marshal(b, m, deterministic)
}
func (m *WaitTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitTime.Merge(m, src)
}
func (m *WaitTime) XXX_Size() int {
	return xxx_messageInfo_WaitTime.Size(m)
}
func (m *WaitTime) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitTime.DiscardUnknown(m)
}

var xxx_messageInfo_WaitTime proto.InternalMessageInfo

func (m *WaitTime) GetParkId() string {
	if m != nil {
		return m.ParkId
	}
	return ""
}

func (m *WaitTime) GetAttractionId() string {
	if m != nil {
		return m.AttractionId
	}
	return ""
}

func (m *WaitTime) GetWaitTime() int32 {
	if m != nil {
		return m.WaitTime
	}
	return 0
}

func (m *WaitTime) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type GetAttractionsRequest struct {
	ParkId               string   `protobuf:"bytes,1,opt,name=park_id,json=parkId,proto3" json:"park_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAttractionsRequest) Reset()         { *m = GetAttractionsRequest{} }
func (m *GetAttractionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAttractionsRequest) ProtoMessage()    {}
func (*GetAttractionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *GetAttractionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttractionsRequest.Unmarshal(m, b)
}
func (m *GetAttractionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttractionsRequest.Marshal(b, m, deterministic)
}
func (m *GetAttractionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttractionsRequest.Merge(m, src)
}
func (m *GetAttractionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAttractionsRequest.Size(m)
}
func (m *GetAttractionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttractionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttractionsRequest proto.InternalMessageInfo

func (m *GetAttractionsRequest) GetParkId() string {
	if m != nil {
		return m.ParkId
	}
	return ""
}

type GetAttractionsResponse struct {
	Attraction           []*Attraction `protobuf:"bytes,1,rep,name=attraction,proto3" json:"attraction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetAttractionsResponse) Reset()         { *m = GetAttractionsResponse{} }
func (m *GetAttractionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAttractionsResponse) ProtoMessage()    {}
func (*GetAttractionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *GetAttractionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttractionsResponse.Unmarshal(m, b)
}
func (m *GetAttractionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttractionsResponse.Marshal(b, m, deterministic)
}
func (m *GetAttractionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttractionsResponse.Merge(m, src)
}
func (m *GetAttractionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAttractionsResponse.Size(m)
}
func (m *GetAttractionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttractionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttractionsResponse proto.InternalMessageInfo

func (m *GetAttractionsResponse) GetAttraction() []*Attraction {
	if m != nil {
		return m.Attraction
	}
	return nil
}

type Attraction struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attraction) Reset()         { *m = Attraction{} }
func (m *Attraction) String() string { return proto.CompactTextString(m) }
func (*Attraction) ProtoMessage()    {}
func (*Attraction) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *Attraction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attraction.Unmarshal(m, b)
}
func (m *Attraction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attraction.Marshal(b, m, deterministic)
}
func (m *Attraction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attraction.Merge(m, src)
}
func (m *Attraction) XXX_Size() int {
	return xxx_messageInfo_Attraction.Size(m)
}
func (m *Attraction) XXX_DiscardUnknown() {
	xxx_messageInfo_Attraction.DiscardUnknown(m)
}

var xxx_messageInfo_Attraction proto.InternalMessageInfo

func (m *Attraction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Attraction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetWaitTimesRequest)(nil), "proto.GetWaitTimesRequest")
	proto.RegisterType((*GetWaitHistoryRequest)(nil), "proto.GetWaitHistoryRequest")
	proto.RegisterType((*GetWaitTimesResponse)(nil), "proto.GetWaitTimesResponse")
	proto.RegisterType((*WaitTime)(nil), "proto.WaitTime")
	proto.RegisterType((*GetAttractionsRequest)(nil), "proto.GetAttractionsRequest")
	proto.RegisterType((*GetAttractionsResponse)(nil), "proto.GetAttractionsResponse")
	proto.RegisterType((*Attraction)(nil), "proto.Attraction")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0x33, 0xe5, 0xcf, 0x85, 0x03, 0x97, 0x9b, 0x3b, 0xf7, 0xaa, 0xa4, 0x68, 0x6c, 0xc6,
	0x4d, 0x17, 0xa6, 0x20, 0x6e, 0x75, 0x61, 0x62, 0x82, 0x84, 0xb0, 0x69, 0x30, 0x2e, 0xc9, 0x40,
	0x47, 0x32, 0x11, 0x98, 0xda, 0x0e, 0x21, 0xbe, 0x8f, 0x8f, 0xe7, 0x43, 0x98, 0x4e, 0xa7, 0x74,
	0x2a, 0xe8, 0xaa, 0x33, 0x73, 0xbe, 0x73, 0xbe, 0x2f, 0xbf, 0x1e, 0xa8, 0xd3, 0x90, 0x7b, 0x61,
	0x24, 0xa4, 0xc0, 0x15, 0xf5, 0xb1, 0xcf, 0x17, 0x42, 0x2c, 0x96, 0xac, 0xab, 0x6e, 0xb3, 0xcd,
	0x73, 0x57, 0xf2, 0x15, 0x8b, 0x25, 0x5d, 0x85, 0xa9, 0x8e, 0x78, 0xf0, 0x6f, 0xc0, 0xe4, 0x13,
	0xe5, 0x72, 0x92, 0x54, 0x7c, 0xf6, 0xba, 0x61, 0xb1, 0xc4, 0x27, 0xf0, 0x2b, 0xa4, 0xd1, 0xcb,
	0x94, 0x07, 0x6d, 0xcb, 0x41, 0x6e, 0xdd, 0xaf, 0x26, 0xd7, 0x61, 0x40, 0x6e, 0xe0, 0x48, 0xeb,
	0x1f, 0x78, 0x2c, 0x45, 0xf4, 0x96, 0x75, 0x5c, 0xc0, 0x6f, 0x2a, 0x65, 0x44, 0xe7, 0x92, 0x8b,
	0x75, 0xd2, 0x87, 0x54, 0x5f, 0x33, 0x7f, 0x1c, 0x06, 0xe4, 0x1e, 0xfe, 0x17, 0xdd, 0xe2, 0x50,
	0xac, 0x63, 0x86, 0x2f, 0xa1, 0xbe, 0xa5, 0x5c, 0x4e, 0x93, 0x74, 0x6d, 0xe4, 0x94, 0xdc, 0x46,
	0xff, 0x4f, 0x1a, 0xd0, 0xcb, 0xc4, 0x7e, 0x6d, 0xab, 0x4f, 0xe4, 0x1d, 0x41, 0x2d, 0x7b, 0x36,
	0x93, 0x22, 0x33, 0xe9, 0x7e, 0x20, 0x6b, 0x3f, 0x10, 0xee, 0x98, 0xc6, 0x25, 0x07, 0xb9, 0x95,
	0xdc, 0x07, 0xdf, 0x42, 0x73, 0x49, 0x63, 0x39, 0xdd, 0x84, 0x01, 0x95, 0x2c, 0x68, 0x97, 0x1d,
	0xe4, 0x36, 0xfa, 0xb6, 0x97, 0x32, 0xf5, 0x32, 0xa6, 0xde, 0x24, 0x63, 0xea, 0x37, 0x12, 0xfd,
	0x63, 0x2a, 0x27, 0x3d, 0x85, 0xea, 0x6e, 0x67, 0x77, 0x08, 0x6e, 0x21, 0x32, 0x19, 0xc1, 0xf1,
	0xd7, 0x0e, 0x0d, 0xe8, 0x0a, 0x20, 0xcf, 0xad, 0x09, 0xfd, 0xd5, 0x84, 0x72, 0xbd, 0x6f, 0x88,
	0x48, 0x0f, 0x20, 0xaf, 0xe0, 0x16, 0x58, 0x3b, 0x3b, 0x8b, 0x07, 0x18, 0x43, 0x79, 0x4d, 0x57,
	0x4c, 0x43, 0x51, 0xe7, 0xfe, 0x07, 0x82, 0xca, 0x98, 0x2e, 0xf8, 0x1c, 0x0f, 0xa0, 0x69, 0xfe,
	0x27, 0x6c, 0x6b, 0xab, 0x03, 0xab, 0x62, 0x77, 0x0e, 0xd6, 0x74, 0xee, 0x11, 0xb4, 0x8a, 0xeb,
	0x82, 0x4f, 0x8b, 0xf2, 0xe2, 0x16, 0xfd, 0x3c, 0x6c, 0xac, 0x86, 0x19, 0x78, 0xcc, 0x61, 0xfb,
	0x9c, 0xed, 0xb3, 0x6f, 0xaa, 0xe9, 0xb8, 0x59, 0x55, 0x55, 0xaf, 0x3f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xe8, 0xfe, 0xb9, 0xf5, 0x36, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MagicClient is the client API for Magic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MagicClient interface {
	GetWaitTimes(ctx context.Context, in *GetWaitTimesRequest, opts ...grpc.CallOption) (*GetWaitTimesResponse, error)
	GetWaitHistory(ctx context.Context, in *GetWaitHistoryRequest, opts ...grpc.CallOption) (*GetWaitTimesResponse, error)
	GetAttractions(ctx context.Context, in *GetAttractionsRequest, opts ...grpc.CallOption) (*GetAttractionsResponse, error)
}

type magicClient struct {
	cc *grpc.ClientConn
}

func NewMagicClient(cc *grpc.ClientConn) MagicClient {
	return &magicClient{cc}
}

func (c *magicClient) GetWaitTimes(ctx context.Context, in *GetWaitTimesRequest, opts ...grpc.CallOption) (*GetWaitTimesResponse, error) {
	out := new(GetWaitTimesResponse)
	err := c.cc.Invoke(ctx, "/proto.Magic/GetWaitTimes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *magicClient) GetWaitHistory(ctx context.Context, in *GetWaitHistoryRequest, opts ...grpc.CallOption) (*GetWaitTimesResponse, error) {
	out := new(GetWaitTimesResponse)
	err := c.cc.Invoke(ctx, "/proto.Magic/GetWaitHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *magicClient) GetAttractions(ctx context.Context, in *GetAttractionsRequest, opts ...grpc.CallOption) (*GetAttractionsResponse, error) {
	out := new(GetAttractionsResponse)
	err := c.cc.Invoke(ctx, "/proto.Magic/GetAttractions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MagicServer is the server API for Magic service.
type MagicServer interface {
	GetWaitTimes(context.Context, *GetWaitTimesRequest) (*GetWaitTimesResponse, error)
	GetWaitHistory(context.Context, *GetWaitHistoryRequest) (*GetWaitTimesResponse, error)
	GetAttractions(context.Context, *GetAttractionsRequest) (*GetAttractionsResponse, error)
}

// UnimplementedMagicServer can be embedded to have forward compatible implementations.
type UnimplementedMagicServer struct {
}

func (*UnimplementedMagicServer) GetWaitTimes(ctx context.Context, req *GetWaitTimesRequest) (*GetWaitTimesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWaitTimes not implemented")
}
func (*UnimplementedMagicServer) GetWaitHistory(ctx context.Context, req *GetWaitHistoryRequest) (*GetWaitTimesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWaitHistory not implemented")
}
func (*UnimplementedMagicServer) GetAttractions(ctx context.Context, req *GetAttractionsRequest) (*GetAttractionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttractions not implemented")
}

func RegisterMagicServer(s *grpc.Server, srv MagicServer) {
	s.RegisterService(&_Magic_serviceDesc, srv)
}

func _Magic_GetWaitTimes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWaitTimesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagicServer).GetWaitTimes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Magic/GetWaitTimes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagicServer).GetWaitTimes(ctx, req.(*GetWaitTimesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Magic_GetWaitHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWaitHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagicServer).GetWaitHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Magic/GetWaitHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagicServer).GetWaitHistory(ctx, req.(*GetWaitHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Magic_GetAttractions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttractionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagicServer).GetAttractions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Magic/GetAttractions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagicServer).GetAttractions(ctx, req.(*GetAttractionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Magic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Magic",
	HandlerType: (*MagicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWaitTimes",
			Handler:    _Magic_GetWaitTimes_Handler,
		},
		{
			MethodName: "GetWaitHistory",
			Handler:    _Magic_GetWaitHistory_Handler,
		},
		{
			MethodName: "GetAttractions",
			Handler:    _Magic_GetAttractions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
