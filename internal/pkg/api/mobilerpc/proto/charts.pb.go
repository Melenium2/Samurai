// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: charts.proto

package charts

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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	GsfId    int64  `protobuf:"varint,3,opt,name=gsfId,proto3" json:"gsfId,omitempty"`
	Token    string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Locale   string `protobuf:"bytes,5,opt,name=locale,proto3" json:"locale,omitempty"`
	Proxy    *Proxy `protobuf:"bytes,6,opt,name=proxy,proto3" json:"proxy,omitempty"`
	Device   string `protobuf:"bytes,7,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_charts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_charts_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Account) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Account) GetGsfId() int64 {
	if x != nil {
		return x.GsfId
	}
	return 0
}

func (x *Account) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Account) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Account) GetProxy() *Proxy {
	if x != nil {
		return x.Proxy
	}
	return nil
}

func (x *Account) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

type Proxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http  string `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Https string `protobuf:"bytes,2,opt,name=https,proto3" json:"https,omitempty"`
	No    string `protobuf:"bytes,3,opt,name=no,proto3" json:"no,omitempty"`
}

func (x *Proxy) Reset() {
	*x = Proxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proxy) ProtoMessage() {}

func (x *Proxy) ProtoReflect() protoreflect.Message {
	mi := &file_charts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proxy.ProtoReflect.Descriptor instead.
func (*Proxy) Descriptor() ([]byte, []int) {
	return file_charts_proto_rawDescGZIP(), []int{1}
}

func (x *Proxy) GetHttp() string {
	if x != nil {
		return x.Http
	}
	return ""
}

func (x *Proxy) GetHttps() string {
	if x != nil {
		return x.Https
	}
	return ""
}

func (x *Proxy) GetNo() string {
	if x != nil {
		return x.No
	}
	return ""
}

type ChartsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cat     string   `protobuf:"bytes,1,opt,name=cat,proto3" json:"cat,omitempty"`
	SubCat  string   `protobuf:"bytes,2,opt,name=subCat,proto3" json:"subCat,omitempty"`
	Account *Account `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *ChartsRequest) Reset() {
	*x = ChartsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartsRequest) ProtoMessage() {}

func (x *ChartsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_charts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartsRequest.ProtoReflect.Descriptor instead.
func (*ChartsRequest) Descriptor() ([]byte, []int) {
	return file_charts_proto_rawDescGZIP(), []int{2}
}

func (x *ChartsRequest) GetCat() string {
	if x != nil {
		return x.Cat
	}
	return ""
}

func (x *ChartsRequest) GetSubCat() string {
	if x != nil {
		return x.SubCat
	}
	return ""
}

func (x *ChartsRequest) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type ChartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids     []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Account *Account `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	ErrCode int32    `protobuf:"varint,3,opt,name=errCode,proto3" json:"errCode,omitempty"`
}

func (x *ChartResponse) Reset() {
	*x = ChartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartResponse) ProtoMessage() {}

func (x *ChartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_charts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartResponse.ProtoReflect.Descriptor instead.
func (*ChartResponse) Descriptor() ([]byte, []int) {
	return file_charts_proto_rawDescGZIP(), []int{3}
}

func (x *ChartResponse) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ChartResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *ChartResponse) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

var File_charts_proto protoreflect.FileDescriptor

var file_charts_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5,
	0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x73, 0x66, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x73, 0x66,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65,
	0x12, 0x1c, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x41, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x74, 0x74, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x74, 0x74, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x68, 0x74, 0x74, 0x70, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x6f, 0x22, 0x5d, 0x0a, 0x0d, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75,
	0x62, 0x43, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5f, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x22, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x36, 0x0a, 0x05, 0x47, 0x75, 0x61,
	0x72, 0x64, 0x12, 0x2d, 0x0a, 0x09, 0x54, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x12,
	0x0e, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_charts_proto_rawDescOnce sync.Once
	file_charts_proto_rawDescData = file_charts_proto_rawDesc
)

func file_charts_proto_rawDescGZIP() []byte {
	file_charts_proto_rawDescOnce.Do(func() {
		file_charts_proto_rawDescData = protoimpl.X.CompressGZIP(file_charts_proto_rawDescData)
	})
	return file_charts_proto_rawDescData
}

var file_charts_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_charts_proto_goTypes = []interface{}{
	(*Account)(nil),       // 0: Account
	(*Proxy)(nil),         // 1: Proxy
	(*ChartsRequest)(nil), // 2: ChartsRequest
	(*ChartResponse)(nil), // 3: ChartResponse
}
var file_charts_proto_depIdxs = []int32{
	1, // 0: Account.proxy:type_name -> Proxy
	0, // 1: ChartsRequest.account:type_name -> Account
	0, // 2: ChartResponse.account:type_name -> Account
	2, // 3: Guard.TopCharts:input_type -> ChartsRequest
	3, // 4: Guard.TopCharts:output_type -> ChartResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_charts_proto_init() }
func file_charts_proto_init() {
	if File_charts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_charts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_charts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proxy); i {
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
		file_charts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChartsRequest); i {
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
		file_charts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChartResponse); i {
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
			RawDescriptor: file_charts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_charts_proto_goTypes,
		DependencyIndexes: file_charts_proto_depIdxs,
		MessageInfos:      file_charts_proto_msgTypes,
	}.Build()
	File_charts_proto = out.File
	file_charts_proto_rawDesc = nil
	file_charts_proto_goTypes = nil
	file_charts_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GuardClient is the client API for Guard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GuardClient interface {
	TopCharts(ctx context.Context, in *ChartsRequest, opts ...grpc.CallOption) (*ChartResponse, error)
}

type guardClient struct {
	cc grpc.ClientConnInterface
}

func NewGuardClient(cc grpc.ClientConnInterface) GuardClient {
	return &guardClient{cc}
}

func (c *guardClient) TopCharts(ctx context.Context, in *ChartsRequest, opts ...grpc.CallOption) (*ChartResponse, error) {
	out := new(ChartResponse)
	err := c.cc.Invoke(ctx, "/Guard/TopCharts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuardServer is the server API for Guard service.
type GuardServer interface {
	TopCharts(context.Context, *ChartsRequest) (*ChartResponse, error)
}

// UnimplementedGuardServer can be embedded to have forward compatible implementations.
type UnimplementedGuardServer struct {
}

func (*UnimplementedGuardServer) TopCharts(context.Context, *ChartsRequest) (*ChartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopCharts not implemented")
}

func RegisterGuardServer(s *grpc.Server, srv GuardServer) {
	s.RegisterService(&_Guard_serviceDesc, srv)
}

func _Guard_TopCharts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChartsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuardServer).TopCharts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Guard/TopCharts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuardServer).TopCharts(ctx, req.(*ChartsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Guard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Guard",
	HandlerType: (*GuardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TopCharts",
			Handler:    _Guard_TopCharts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "charts.proto",
}
