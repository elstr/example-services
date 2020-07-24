// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: stock.proto

package stock

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

// The item and quantity we will update
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item     int32 `protobuf:"varint,1,opt,name=item,proto3" json:"item,omitempty"`
	Quantity int32 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetItem() int32 {
	if x != nil {
		return x.Item
	}
	return 0
}

func (x *Request) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// The fake delivery date
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryDate string `protobuf:"bytes,1,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetDeliveryDate() string {
	if x != nil {
		return x.DeliveryDate
	}
	return ""
}

var File_stock_proto protoreflect.FileDescriptor

var file_stock_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x22, 0x39, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x2f, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65,
	0x32, 0x37, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x2e, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_stock_proto_rawDescOnce sync.Once
	file_stock_proto_rawDescData = file_stock_proto_rawDesc
)

func file_stock_proto_rawDescGZIP() []byte {
	file_stock_proto_rawDescOnce.Do(func() {
		file_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_stock_proto_rawDescData)
	})
	return file_stock_proto_rawDescData
}

var file_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stock_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: stock.Request
	(*Response)(nil), // 1: stock.Response
}
var file_stock_proto_depIdxs = []int32{
	0, // 0: stock.Stock.UpdateStock:input_type -> stock.Request
	1, // 1: stock.Stock.UpdateStock:output_type -> stock.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stock_proto_init() }
func file_stock_proto_init() {
	if File_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stock_proto_goTypes,
		DependencyIndexes: file_stock_proto_depIdxs,
		MessageInfos:      file_stock_proto_msgTypes,
	}.Build()
	File_stock_proto = out.File
	file_stock_proto_rawDesc = nil
	file_stock_proto_goTypes = nil
	file_stock_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StockClient is the client API for Stock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StockClient interface {
	// Updates stock quantity
	UpdateStock(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type stockClient struct {
	cc grpc.ClientConnInterface
}

func NewStockClient(cc grpc.ClientConnInterface) StockClient {
	return &stockClient{cc}
}

func (c *stockClient) UpdateStock(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/stock.Stock/UpdateStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServer is the server API for Stock service.
type StockServer interface {
	// Updates stock quantity
	UpdateStock(context.Context, *Request) (*Response, error)
}

// UnimplementedStockServer can be embedded to have forward compatible implementations.
type UnimplementedStockServer struct {
}

func (*UnimplementedStockServer) UpdateStock(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStock not implemented")
}

func RegisterStockServer(s *grpc.Server, srv StockServer) {
	s.RegisterService(&_Stock_serviceDesc, srv)
}

func _Stock_UpdateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).UpdateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.Stock/UpdateStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).UpdateStock(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stock.Stock",
	HandlerType: (*StockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateStock",
			Handler:    _Stock_UpdateStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stock.proto",
}