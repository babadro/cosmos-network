// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchange/dex/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3176e7bba247186e, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3176e7bba247186e, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetSellOrderBookRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetSellOrderBookRequest) Reset()         { *m = QueryGetSellOrderBookRequest{} }
func (m *QueryGetSellOrderBookRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetSellOrderBookRequest) ProtoMessage()    {}
func (*QueryGetSellOrderBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3176e7bba247186e, []int{2}
}
func (m *QueryGetSellOrderBookRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSellOrderBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSellOrderBookRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSellOrderBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSellOrderBookRequest.Merge(m, src)
}
func (m *QueryGetSellOrderBookRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSellOrderBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSellOrderBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSellOrderBookRequest proto.InternalMessageInfo

func (m *QueryGetSellOrderBookRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetSellOrderBookResponse struct {
	SellOrderBook SellOrderBook `protobuf:"bytes,1,opt,name=sellOrderBook,proto3" json:"sellOrderBook"`
}

func (m *QueryGetSellOrderBookResponse) Reset()         { *m = QueryGetSellOrderBookResponse{} }
func (m *QueryGetSellOrderBookResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetSellOrderBookResponse) ProtoMessage()    {}
func (*QueryGetSellOrderBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3176e7bba247186e, []int{3}
}
func (m *QueryGetSellOrderBookResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSellOrderBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSellOrderBookResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSellOrderBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSellOrderBookResponse.Merge(m, src)
}
func (m *QueryGetSellOrderBookResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSellOrderBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSellOrderBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSellOrderBookResponse proto.InternalMessageInfo

func (m *QueryGetSellOrderBookResponse) GetSellOrderBook() SellOrderBook {
	if m != nil {
		return m.SellOrderBook
	}
	return SellOrderBook{}
}

type QueryAllSellOrderBookRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllSellOrderBookRequest) Reset()         { *m = QueryAllSellOrderBookRequest{} }
func (m *QueryAllSellOrderBookRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllSellOrderBookRequest) ProtoMessage()    {}
func (*QueryAllSellOrderBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3176e7bba247186e, []int{4}
}
func (m *QueryAllSellOrderBookRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllSellOrderBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllSellOrderBookRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllSellOrderBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllSellOrderBookRequest.Merge(m, src)
}
func (m *QueryAllSellOrderBookRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllSellOrderBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllSellOrderBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllSellOrderBookRequest proto.InternalMessageInfo

func (m *QueryAllSellOrderBookRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllSellOrderBookResponse struct {
	SellOrderBook []SellOrderBook     `protobuf:"bytes,1,rep,name=sellOrderBook,proto3" json:"sellOrderBook"`
	Pagination    *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllSellOrderBookResponse) Reset()         { *m = QueryAllSellOrderBookResponse{} }
func (m *QueryAllSellOrderBookResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllSellOrderBookResponse) ProtoMessage()    {}
func (*QueryAllSellOrderBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3176e7bba247186e, []int{5}
}
func (m *QueryAllSellOrderBookResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllSellOrderBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllSellOrderBookResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllSellOrderBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllSellOrderBookResponse.Merge(m, src)
}
func (m *QueryAllSellOrderBookResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllSellOrderBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllSellOrderBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllSellOrderBookResponse proto.InternalMessageInfo

func (m *QueryAllSellOrderBookResponse) GetSellOrderBook() []SellOrderBook {
	if m != nil {
		return m.SellOrderBook
	}
	return nil
}

func (m *QueryAllSellOrderBookResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "interchange.dex.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "interchange.dex.QueryParamsResponse")
	proto.RegisterType((*QueryGetSellOrderBookRequest)(nil), "interchange.dex.QueryGetSellOrderBookRequest")
	proto.RegisterType((*QueryGetSellOrderBookResponse)(nil), "interchange.dex.QueryGetSellOrderBookResponse")
	proto.RegisterType((*QueryAllSellOrderBookRequest)(nil), "interchange.dex.QueryAllSellOrderBookRequest")
	proto.RegisterType((*QueryAllSellOrderBookResponse)(nil), "interchange.dex.QueryAllSellOrderBookResponse")
}

func init() { proto.RegisterFile("interchange/dex/query.proto", fileDescriptor_3176e7bba247186e) }

var fileDescriptor_3176e7bba247186e = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0x96, 0x46, 0x62, 0x51, 0x05, 0x5a, 0x22, 0x05, 0x42, 0x71, 0xab, 0xe5, 0x5f,
	0x54, 0x89, 0x5d, 0x52, 0xe0, 0x01, 0x9a, 0x03, 0x95, 0x10, 0x12, 0x25, 0xdc, 0xb8, 0x54, 0x9b,
	0x66, 0x30, 0x56, 0xb6, 0x1e, 0xd7, 0xbb, 0x45, 0xa9, 0x10, 0x17, 0x9e, 0x00, 0x89, 0x23, 0x17,
	0x1e, 0x82, 0x33, 0xe7, 0x1e, 0x2b, 0x71, 0xe1, 0x84, 0x50, 0xc2, 0x83, 0xa0, 0xec, 0x6e, 0x85,
	0x1d, 0xc7, 0x04, 0xf5, 0x96, 0x64, 0xe6, 0xfb, 0xe6, 0x37, 0x99, 0xcf, 0x26, 0x37, 0xe2, 0xc4,
	0x40, 0xb6, 0xff, 0x46, 0x26, 0x11, 0x88, 0x01, 0x8c, 0xc4, 0xe1, 0x11, 0x64, 0xc7, 0x3c, 0xcd,
	0xd0, 0x20, 0xbd, 0x9c, 0x2b, 0xf2, 0x01, 0x8c, 0x5a, 0x8d, 0x08, 0x23, 0xb4, 0x35, 0x31, 0xfd,
	0xe4, 0xda, 0x5a, 0x6b, 0x11, 0x62, 0xa4, 0x40, 0xc8, 0x34, 0x16, 0x32, 0x49, 0xd0, 0x48, 0x13,
	0x63, 0xa2, 0x7d, 0x75, 0x73, 0x1f, 0xf5, 0x01, 0x6a, 0xd1, 0x97, 0x1a, 0x9c, 0xbb, 0x78, 0xdb,
	0xe9, 0x83, 0x91, 0x1d, 0x91, 0xca, 0x28, 0x4e, 0x6c, 0xf3, 0x99, 0xd3, 0x2c, 0x4d, 0x2a, 0x33,
	0x79, 0x70, 0xe6, 0x74, 0x67, 0xb6, 0xaa, 0x41, 0xa9, 0x3d, 0xcc, 0x06, 0x90, 0xed, 0xf5, 0x11,
	0x87, 0xae, 0x8d, 0x35, 0x08, 0x7d, 0x31, 0x1d, 0xb3, 0x6b, 0xb5, 0x3d, 0x38, 0x3c, 0x02, 0x6d,
	0xd8, 0x33, 0x72, 0xb5, 0xf0, 0xab, 0x4e, 0x31, 0xd1, 0x40, 0x1f, 0x93, 0xba, 0x9b, 0x71, 0x2d,
	0xd8, 0x08, 0xda, 0x97, 0xb6, 0x9a, 0x7c, 0x66, 0x67, 0xee, 0x04, 0xdd, 0x0b, 0x27, 0x3f, 0xd7,
	0x6b, 0x3d, 0xdf, 0xcc, 0x1e, 0x91, 0x35, 0xeb, 0xb6, 0x03, 0xe6, 0x25, 0x28, 0xf5, 0x7c, 0xca,
	0xd0, 0x45, 0x1c, 0xfa, 0x69, 0xb4, 0x41, 0x56, 0xe2, 0x64, 0x00, 0x23, 0xeb, 0x7a, 0xb1, 0xe7,
	0xbe, 0xb0, 0x21, 0xb9, 0x59, 0xa1, 0xf2, 0x34, 0x4f, 0xc9, 0xaa, 0xce, 0x17, 0x3c, 0x54, 0x58,
	0x82, 0x2a, 0xc8, 0x3d, 0x5b, 0x51, 0xca, 0x5e, 0x7b, 0xc4, 0x6d, 0xa5, 0xe6, 0x22, 0x3e, 0x21,
	0xe4, 0xef, 0xff, 0xef, 0x07, 0xdd, 0xe5, 0xee, 0x58, 0x7c, 0x7a, 0x2c, 0xee, 0xa2, 0xe0, 0x8f,
	0xc5, 0x77, 0x65, 0x04, 0x5e, 0xdb, 0xcb, 0x29, 0xd9, 0xd7, 0xc0, 0x6f, 0x55, 0x1e, 0x54, 0xbd,
	0xd5, 0xf2, 0x39, 0xb7, 0xa2, 0x3b, 0x05, 0xea, 0x25, 0x4b, 0x7d, 0x6f, 0x21, 0xb5, 0x03, 0xc9,
	0x63, 0x6f, 0x7d, 0x5b, 0x26, 0x2b, 0x16, 0x9b, 0x1a, 0x52, 0x77, 0x37, 0xa6, 0xb7, 0x4a, 0x44,
	0xe5, 0x20, 0xb5, 0x6e, 0xff, 0xbb, 0xc9, 0x8d, 0x62, 0xeb, 0x1f, 0xbe, 0xff, 0xfe, 0xb4, 0x74,
	0x9d, 0x36, 0xc5, 0xfc, 0x48, 0xd3, 0x2f, 0x01, 0x59, 0x2d, 0xec, 0x4b, 0xef, 0xcf, 0x37, 0xae,
	0x88, 0x58, 0x8b, 0xff, 0x6f, 0xbb, 0x27, 0x7a, 0x60, 0x89, 0x36, 0x69, 0x5b, 0x2c, 0x78, 0x8c,
	0xc4, 0x3b, 0x9b, 0xd6, 0xf7, 0xf4, 0x73, 0x40, 0xae, 0x14, 0xbc, 0xb6, 0x95, 0xaa, 0xa2, 0xac,
	0x48, 0x59, 0x15, 0x65, 0x55, 0x56, 0x58, 0xdb, 0x52, 0x32, 0xba, 0xb1, 0x88, 0xb2, 0xdb, 0x39,
	0x19, 0x87, 0xc1, 0xe9, 0x38, 0x0c, 0x7e, 0x8d, 0xc3, 0xe0, 0xe3, 0x24, 0xac, 0x9d, 0x4e, 0xc2,
	0xda, 0x8f, 0x49, 0x58, 0x7b, 0xd5, 0xcc, 0x4b, 0x47, 0x56, 0x6c, 0x8e, 0x53, 0xd0, 0xfd, 0xba,
	0x7d, 0x41, 0x3c, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x33, 0x54, 0x07, 0xf5, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of SellOrderBook items.
	SellOrderBook(ctx context.Context, in *QueryGetSellOrderBookRequest, opts ...grpc.CallOption) (*QueryGetSellOrderBookResponse, error)
	SellOrderBookAll(ctx context.Context, in *QueryAllSellOrderBookRequest, opts ...grpc.CallOption) (*QueryAllSellOrderBookResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/interchange.dex.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SellOrderBook(ctx context.Context, in *QueryGetSellOrderBookRequest, opts ...grpc.CallOption) (*QueryGetSellOrderBookResponse, error) {
	out := new(QueryGetSellOrderBookResponse)
	err := c.cc.Invoke(ctx, "/interchange.dex.Query/SellOrderBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SellOrderBookAll(ctx context.Context, in *QueryAllSellOrderBookRequest, opts ...grpc.CallOption) (*QueryAllSellOrderBookResponse, error) {
	out := new(QueryAllSellOrderBookResponse)
	err := c.cc.Invoke(ctx, "/interchange.dex.Query/SellOrderBookAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of SellOrderBook items.
	SellOrderBook(context.Context, *QueryGetSellOrderBookRequest) (*QueryGetSellOrderBookResponse, error)
	SellOrderBookAll(context.Context, *QueryAllSellOrderBookRequest) (*QueryAllSellOrderBookResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) SellOrderBook(ctx context.Context, req *QueryGetSellOrderBookRequest) (*QueryGetSellOrderBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellOrderBook not implemented")
}
func (*UnimplementedQueryServer) SellOrderBookAll(ctx context.Context, req *QueryAllSellOrderBookRequest) (*QueryAllSellOrderBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellOrderBookAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interchange.dex.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SellOrderBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetSellOrderBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SellOrderBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interchange.dex.Query/SellOrderBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SellOrderBook(ctx, req.(*QueryGetSellOrderBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SellOrderBookAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllSellOrderBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SellOrderBookAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interchange.dex.Query/SellOrderBookAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SellOrderBookAll(ctx, req.(*QueryAllSellOrderBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interchange.dex.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "SellOrderBook",
			Handler:    _Query_SellOrderBook_Handler,
		},
		{
			MethodName: "SellOrderBookAll",
			Handler:    _Query_SellOrderBookAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interchange/dex/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetSellOrderBookRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSellOrderBookRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSellOrderBookRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetSellOrderBookResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSellOrderBookResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSellOrderBookResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SellOrderBook.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllSellOrderBookRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllSellOrderBookRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllSellOrderBookRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllSellOrderBookResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllSellOrderBookResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllSellOrderBookResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SellOrderBook) > 0 {
		for iNdEx := len(m.SellOrderBook) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SellOrderBook[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetSellOrderBookRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetSellOrderBookResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SellOrderBook.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllSellOrderBookRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllSellOrderBookResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SellOrderBook) > 0 {
		for _, e := range m.SellOrderBook {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetSellOrderBookRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetSellOrderBookRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSellOrderBookRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetSellOrderBookResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetSellOrderBookResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSellOrderBookResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellOrderBook", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SellOrderBook.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllSellOrderBookRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllSellOrderBookRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllSellOrderBookRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllSellOrderBookResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllSellOrderBookResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllSellOrderBookResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellOrderBook", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SellOrderBook = append(m.SellOrderBook, SellOrderBook{})
			if err := m.SellOrderBook[len(m.SellOrderBook)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)