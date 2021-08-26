// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protocol/calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
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

type Addition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstInteger  int32 `protobuf:"varint,1,opt,name=first_integer,json=firstInteger,proto3" json:"first_integer,omitempty"`
	SecondInteger int32 `protobuf:"varint,2,opt,name=second_integer,json=secondInteger,proto3" json:"second_integer,omitempty"`
}

func (x *Addition) Reset() {
	*x = Addition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Addition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Addition) ProtoMessage() {}

func (x *Addition) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Addition.ProtoReflect.Descriptor instead.
func (*Addition) Descriptor() ([]byte, []int) {
	return file_protocol_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *Addition) GetFirstInteger() int32 {
	if x != nil {
		return x.FirstInteger
	}
	return 0
}

func (x *Addition) GetSecondInteger() int32 {
	if x != nil {
		return x.SecondInteger
	}
	return 0
}

type AdditionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addition *Addition `protobuf:"bytes,1,opt,name=addition,proto3" json:"addition,omitempty"`
}

func (x *AdditionRequest) Reset() {
	*x = AdditionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionRequest) ProtoMessage() {}

func (x *AdditionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionRequest.ProtoReflect.Descriptor instead.
func (*AdditionRequest) Descriptor() ([]byte, []int) {
	return file_protocol_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *AdditionRequest) GetAddition() *Addition {
	if x != nil {
		return x.Addition
	}
	return nil
}

type AdditionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AdditionResponse) Reset() {
	*x = AdditionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionResponse) ProtoMessage() {}

func (x *AdditionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionResponse.ProtoReflect.Descriptor instead.
func (*AdditionResponse) Descriptor() ([]byte, []int) {
	return file_protocol_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *AdditionResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type AddManyTimesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addition *Addition `protobuf:"bytes,1,opt,name=addition,proto3" json:"addition,omitempty"`
}

func (x *AddManyTimesRequest) Reset() {
	*x = AddManyTimesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddManyTimesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddManyTimesRequest) ProtoMessage() {}

func (x *AddManyTimesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddManyTimesRequest.ProtoReflect.Descriptor instead.
func (*AddManyTimesRequest) Descriptor() ([]byte, []int) {
	return file_protocol_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *AddManyTimesRequest) GetAddition() *Addition {
	if x != nil {
		return x.Addition
	}
	return nil
}

type AddManyTimesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddManyTimesResponse) Reset() {
	*x = AddManyTimesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddManyTimesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddManyTimesResponse) ProtoMessage() {}

func (x *AddManyTimesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddManyTimesResponse.ProtoReflect.Descriptor instead.
func (*AddManyTimesResponse) Descriptor() ([]byte, []int) {
	return file_protocol_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *AddManyTimesResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_protocol_calculator_calculatorpb_calculator_proto protoreflect.FileDescriptor

var file_protocol_calculator_calculatorpb_calculator_proto_rawDesc = []byte{
	0x0a, 0x31, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22,
	0x56, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x22, 0x43, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x10,
	0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x47, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x4d,
	0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x2e, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0xac, 0x01, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1b, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x41, 0x64, 0x64,
	0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x22, 0x5a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_calculator_calculatorpb_calculator_proto_rawDescOnce sync.Once
	file_protocol_calculator_calculatorpb_calculator_proto_rawDescData = file_protocol_calculator_calculatorpb_calculator_proto_rawDesc
)

func file_protocol_calculator_calculatorpb_calculator_proto_rawDescGZIP() []byte {
	file_protocol_calculator_calculatorpb_calculator_proto_rawDescOnce.Do(func() {
		file_protocol_calculator_calculatorpb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_calculator_calculatorpb_calculator_proto_rawDescData)
	})
	return file_protocol_calculator_calculatorpb_calculator_proto_rawDescData
}

var file_protocol_calculator_calculatorpb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protocol_calculator_calculatorpb_calculator_proto_goTypes = []interface{}{
	(*Addition)(nil),             // 0: calculator.Addition
	(*AdditionRequest)(nil),      // 1: calculator.AdditionRequest
	(*AdditionResponse)(nil),     // 2: calculator.AdditionResponse
	(*AddManyTimesRequest)(nil),  // 3: calculator.AddManyTimesRequest
	(*AddManyTimesResponse)(nil), // 4: calculator.AddManyTimesResponse
}
var file_protocol_calculator_calculatorpb_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.AdditionRequest.addition:type_name -> calculator.Addition
	0, // 1: calculator.AddManyTimesRequest.addition:type_name -> calculator.Addition
	1, // 2: calculator.AdditionService.Add:input_type -> calculator.AdditionRequest
	3, // 3: calculator.AdditionService.AddManyTimes:input_type -> calculator.AddManyTimesRequest
	2, // 4: calculator.AdditionService.Add:output_type -> calculator.AdditionResponse
	4, // 5: calculator.AdditionService.AddManyTimes:output_type -> calculator.AddManyTimesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protocol_calculator_calculatorpb_calculator_proto_init() }
func file_protocol_calculator_calculatorpb_calculator_proto_init() {
	if File_protocol_calculator_calculatorpb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Addition); i {
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
		file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionRequest); i {
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
		file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionResponse); i {
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
		file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddManyTimesRequest); i {
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
		file_protocol_calculator_calculatorpb_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddManyTimesResponse); i {
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
			RawDescriptor: file_protocol_calculator_calculatorpb_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_calculator_calculatorpb_calculator_proto_goTypes,
		DependencyIndexes: file_protocol_calculator_calculatorpb_calculator_proto_depIdxs,
		MessageInfos:      file_protocol_calculator_calculatorpb_calculator_proto_msgTypes,
	}.Build()
	File_protocol_calculator_calculatorpb_calculator_proto = out.File
	file_protocol_calculator_calculatorpb_calculator_proto_rawDesc = nil
	file_protocol_calculator_calculatorpb_calculator_proto_goTypes = nil
	file_protocol_calculator_calculatorpb_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdditionServiceClient is the client API for AdditionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdditionServiceClient interface {
	// Unary
	Add(ctx context.Context, in *AdditionRequest, opts ...grpc.CallOption) (*AdditionResponse, error)
	// Server Streaming
	AddManyTimes(ctx context.Context, in *AddManyTimesRequest, opts ...grpc.CallOption) (AdditionService_AddManyTimesClient, error)
}

type additionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdditionServiceClient(cc grpc.ClientConnInterface) AdditionServiceClient {
	return &additionServiceClient{cc}
}

func (c *additionServiceClient) Add(ctx context.Context, in *AdditionRequest, opts ...grpc.CallOption) (*AdditionResponse, error) {
	out := new(AdditionResponse)
	err := c.cc.Invoke(ctx, "/calculator.AdditionService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *additionServiceClient) AddManyTimes(ctx context.Context, in *AddManyTimesRequest, opts ...grpc.CallOption) (AdditionService_AddManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AdditionService_serviceDesc.Streams[0], "/calculator.AdditionService/AddManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &additionServiceAddManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AdditionService_AddManyTimesClient interface {
	Recv() (*AddManyTimesResponse, error)
	grpc.ClientStream
}

type additionServiceAddManyTimesClient struct {
	grpc.ClientStream
}

func (x *additionServiceAddManyTimesClient) Recv() (*AddManyTimesResponse, error) {
	m := new(AddManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AdditionServiceServer is the server API for AdditionService service.
type AdditionServiceServer interface {
	// Unary
	Add(context.Context, *AdditionRequest) (*AdditionResponse, error)
	// Server Streaming
	AddManyTimes(*AddManyTimesRequest, AdditionService_AddManyTimesServer) error
}

// UnimplementedAdditionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdditionServiceServer struct {
}

func (*UnimplementedAdditionServiceServer) Add(context.Context, *AdditionRequest) (*AdditionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedAdditionServiceServer) AddManyTimes(*AddManyTimesRequest, AdditionService_AddManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method AddManyTimes not implemented")
}

func RegisterAdditionServiceServer(s *grpc.Server, srv AdditionServiceServer) {
	s.RegisterService(&_AdditionService_serviceDesc, srv)
}

func _AdditionService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdditionServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.AdditionService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdditionServiceServer).Add(ctx, req.(*AdditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdditionService_AddManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AddManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdditionServiceServer).AddManyTimes(m, &additionServiceAddManyTimesServer{stream})
}

type AdditionService_AddManyTimesServer interface {
	Send(*AddManyTimesResponse) error
	grpc.ServerStream
}

type additionServiceAddManyTimesServer struct {
	grpc.ServerStream
}

func (x *additionServiceAddManyTimesServer) Send(m *AddManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _AdditionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.AdditionService",
	HandlerType: (*AdditionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AdditionService_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddManyTimes",
			Handler:       _AdditionService_AddManyTimes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protocol/calculator/calculatorpb/calculator.proto",
}