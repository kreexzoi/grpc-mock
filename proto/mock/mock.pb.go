// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: mock.proto

package grpc_mock_api

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddStubsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stubs []*Stub `protobuf:"bytes,1,rep,name=stubs,proto3" json:"stubs,omitempty"`
}

func (x *AddStubsRequest) Reset() {
	*x = AddStubsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStubsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStubsRequest) ProtoMessage() {}

func (x *AddStubsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStubsRequest.ProtoReflect.Descriptor instead.
func (*AddStubsRequest) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{0}
}

func (x *AddStubsRequest) GetStubs() []*Stub {
	if x != nil {
		return x.Stubs
	}
	return nil
}

type AddStubsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddStubsResponse) Reset() {
	*x = AddStubsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStubsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStubsResponse) ProtoMessage() {}

func (x *AddStubsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStubsResponse.ProtoReflect.Descriptor instead.
func (*AddStubsResponse) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{1}
}

type FindStubsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string           `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method  string           `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	In      *structpb.Struct `protobuf:"bytes,3,opt,name=in,proto3" json:"in,omitempty"`
}

func (x *FindStubsRequest) Reset() {
	*x = FindStubsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindStubsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStubsRequest) ProtoMessage() {}

func (x *FindStubsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStubsRequest.ProtoReflect.Descriptor instead.
func (*FindStubsRequest) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{2}
}

func (x *FindStubsRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *FindStubsRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *FindStubsRequest) GetIn() *structpb.Struct {
	if x != nil {
		return x.In
	}
	return nil
}

type FindStubsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stubs []*Stub `protobuf:"bytes,1,rep,name=stubs,proto3" json:"stubs,omitempty"`
}

func (x *FindStubsResponse) Reset() {
	*x = FindStubsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindStubsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStubsResponse) ProtoMessage() {}

func (x *FindStubsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStubsResponse.ProtoReflect.Descriptor instead.
func (*FindStubsResponse) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{3}
}

func (x *FindStubsResponse) GetStubs() []*Stub {
	if x != nil {
		return x.Stubs
	}
	return nil
}

type DeleteStubsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method  string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *DeleteStubsRequest) Reset() {
	*x = DeleteStubsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStubsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStubsRequest) ProtoMessage() {}

func (x *DeleteStubsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStubsRequest.ProtoReflect.Descriptor instead.
func (*DeleteStubsRequest) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteStubsRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *DeleteStubsRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type DeleteStubsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteStubsResponse) Reset() {
	*x = DeleteStubsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStubsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStubsResponse) ProtoMessage() {}

func (x *DeleteStubsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStubsResponse.ProtoReflect.Descriptor instead.
func (*DeleteStubsResponse) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{5}
}

type Stub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string  `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method  string  `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	In      *Input  `protobuf:"bytes,3,opt,name=in,proto3" json:"in,omitempty"`
	Out     *Output `protobuf:"bytes,4,opt,name=out,proto3" json:"out,omitempty"`
}

func (x *Stub) Reset() {
	*x = Stub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stub) ProtoMessage() {}

func (x *Stub) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stub.ProtoReflect.Descriptor instead.
func (*Stub) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{6}
}

func (x *Stub) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Stub) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Stub) GetIn() *Input {
	if x != nil {
		return x.In
	}
	return nil
}

func (x *Stub) GetOut() *Output {
	if x != nil {
		return x.Out
	}
	return nil
}

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Rule:
	//	*Input_Equals
	//	*Input_Contains
	//	*Input_Matches
	Rule isInput_Rule `protobuf_oneof:"rule"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{7}
}

func (m *Input) GetRule() isInput_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (x *Input) GetEquals() *structpb.Struct {
	if x, ok := x.GetRule().(*Input_Equals); ok {
		return x.Equals
	}
	return nil
}

func (x *Input) GetContains() *structpb.Struct {
	if x, ok := x.GetRule().(*Input_Contains); ok {
		return x.Contains
	}
	return nil
}

func (x *Input) GetMatches() *structpb.Struct {
	if x, ok := x.GetRule().(*Input_Matches); ok {
		return x.Matches
	}
	return nil
}

type isInput_Rule interface {
	isInput_Rule()
}

type Input_Equals struct {
	Equals *structpb.Struct `protobuf:"bytes,1,opt,name=equals,proto3,oneof"`
}

type Input_Contains struct {
	Contains *structpb.Struct `protobuf:"bytes,2,opt,name=contains,proto3,oneof"`
}

type Input_Matches struct {
	Matches *structpb.Struct `protobuf:"bytes,3,opt,name=matches,proto3,oneof"`
}

func (*Input_Equals) isInput_Rule() {}

func (*Input_Contains) isInput_Rule() {}

func (*Input_Matches) isInput_Rule() {}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *structpb.Struct `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code  int32            `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Error string           `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mock_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_mock_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_mock_proto_rawDescGZIP(), []int{8}
}

func (x *Output) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Output) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Output) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_mock_proto protoreflect.FileDescriptor

var file_mock_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x72, 0x65, 0x65, 0x78, 0x7a, 0x6f,
	0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x63, 0x6b, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x53, 0x74,
	0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x74,
	0x75, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x72, 0x65, 0x65, 0x78, 0x7a, 0x6f, 0x69, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x75, 0x62, 0x52, 0x05,
	0x73, 0x74, 0x75, 0x62, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x62,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d, 0x0a, 0x10, 0x46, 0x69, 0x6e,
	0x64, 0x53, 0x74, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x27, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x02, 0x69, 0x6e, 0x22, 0x4e, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64,
	0x53, 0x74, 0x75, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x05, 0x73, 0x74, 0x75, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x72, 0x65, 0x65, 0x78, 0x7a,
	0x6f, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x75,
	0x62, 0x52, 0x05, 0x73, 0x74, 0x75, 0x62, 0x73, 0x22, 0x46, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x74, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x62, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x04, 0x53, 0x74, 0x75, 0x62,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x72, 0x65, 0x65,
	0x78, 0x7a, 0x6f, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x52, 0x02, 0x69, 0x6e, 0x12, 0x37, 0x0a, 0x03, 0x6f, 0x75, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6b, 0x72, 0x65, 0x65, 0x78, 0x7a, 0x6f, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x03, 0x6f, 0x75,
	0x74, 0x22, 0xae, 0x01, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x65,
	0x71, 0x75, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x06, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x12, 0x35,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48,
	0x00, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x72, 0x75,
	0x6c, 0x65, 0x22, 0x5f, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x32, 0x98, 0x03, 0x0a, 0x04, 0x4d, 0x6f, 0x63, 0x6b, 0x12, 0x81, 0x01, 0x0a,
	0x08, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x62, 0x73, 0x12, 0x2e, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x72, 0x65, 0x65, 0x78, 0x7a, 0x6f, 0x69, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75,
	0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x72, 0x65, 0x65, 0x78, 0x7a, 0x6f, 0x69, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75,
	0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x75, 0x62, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x81, 0x01, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x75, 0x62, 0x73, 0x12, 0x2f,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x72, 0x65, 0x65,
	0x78, 0x7a, 0x6f, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x53, 0x74, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x72, 0x65,
	0x65, 0x78, 0x7a, 0x6f, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x63, 0x6b, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x75, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x75, 0x62, 0x73, 0x12, 0x87, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x74, 0x75, 0x62, 0x73, 0x12, 0x31, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6b, 0x72, 0x65, 0x65, 0x78, 0x7a, 0x6f, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x62, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x72, 0x65, 0x65, 0x78, 0x7a, 0x6f, 0x69, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x75, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0b, 0x2a, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x75, 0x62, 0x73, 0x42, 0x16,
	0x5a, 0x14, 0x2e, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f,
	0x63, 0x6b, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mock_proto_rawDescOnce sync.Once
	file_mock_proto_rawDescData = file_mock_proto_rawDesc
)

func file_mock_proto_rawDescGZIP() []byte {
	file_mock_proto_rawDescOnce.Do(func() {
		file_mock_proto_rawDescData = protoimpl.X.CompressGZIP(file_mock_proto_rawDescData)
	})
	return file_mock_proto_rawDescData
}

var file_mock_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mock_proto_goTypes = []interface{}{
	(*AddStubsRequest)(nil),     // 0: github.com.kreexzoi.grpc_mock.AddStubsRequest
	(*AddStubsResponse)(nil),    // 1: github.com.kreexzoi.grpc_mock.AddStubsResponse
	(*FindStubsRequest)(nil),    // 2: github.com.kreexzoi.grpc_mock.FindStubsRequest
	(*FindStubsResponse)(nil),   // 3: github.com.kreexzoi.grpc_mock.FindStubsResponse
	(*DeleteStubsRequest)(nil),  // 4: github.com.kreexzoi.grpc_mock.DeleteStubsRequest
	(*DeleteStubsResponse)(nil), // 5: github.com.kreexzoi.grpc_mock.DeleteStubsResponse
	(*Stub)(nil),                // 6: github.com.kreexzoi.grpc_mock.Stub
	(*Input)(nil),               // 7: github.com.kreexzoi.grpc_mock.Input
	(*Output)(nil),              // 8: github.com.kreexzoi.grpc_mock.Output
	(*structpb.Struct)(nil),     // 9: google.protobuf.Struct
}
var file_mock_proto_depIdxs = []int32{
	6,  // 0: github.com.kreexzoi.grpc_mock.AddStubsRequest.stubs:type_name -> github.com.kreexzoi.grpc_mock.Stub
	9,  // 1: github.com.kreexzoi.grpc_mock.FindStubsRequest.in:type_name -> google.protobuf.Struct
	6,  // 2: github.com.kreexzoi.grpc_mock.FindStubsResponse.stubs:type_name -> github.com.kreexzoi.grpc_mock.Stub
	7,  // 3: github.com.kreexzoi.grpc_mock.Stub.in:type_name -> github.com.kreexzoi.grpc_mock.Input
	8,  // 4: github.com.kreexzoi.grpc_mock.Stub.out:type_name -> github.com.kreexzoi.grpc_mock.Output
	9,  // 5: github.com.kreexzoi.grpc_mock.Input.equals:type_name -> google.protobuf.Struct
	9,  // 6: github.com.kreexzoi.grpc_mock.Input.contains:type_name -> google.protobuf.Struct
	9,  // 7: github.com.kreexzoi.grpc_mock.Input.matches:type_name -> google.protobuf.Struct
	9,  // 8: github.com.kreexzoi.grpc_mock.Output.data:type_name -> google.protobuf.Struct
	0,  // 9: github.com.kreexzoi.grpc_mock.Mock.AddStubs:input_type -> github.com.kreexzoi.grpc_mock.AddStubsRequest
	2,  // 10: github.com.kreexzoi.grpc_mock.Mock.FindStubs:input_type -> github.com.kreexzoi.grpc_mock.FindStubsRequest
	4,  // 11: github.com.kreexzoi.grpc_mock.Mock.DeleteStubs:input_type -> github.com.kreexzoi.grpc_mock.DeleteStubsRequest
	1,  // 12: github.com.kreexzoi.grpc_mock.Mock.AddStubs:output_type -> github.com.kreexzoi.grpc_mock.AddStubsResponse
	3,  // 13: github.com.kreexzoi.grpc_mock.Mock.FindStubs:output_type -> github.com.kreexzoi.grpc_mock.FindStubsResponse
	5,  // 14: github.com.kreexzoi.grpc_mock.Mock.DeleteStubs:output_type -> github.com.kreexzoi.grpc_mock.DeleteStubsResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_mock_proto_init() }
func file_mock_proto_init() {
	if File_mock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStubsRequest); i {
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
		file_mock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStubsResponse); i {
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
		file_mock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindStubsRequest); i {
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
		file_mock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindStubsResponse); i {
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
		file_mock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStubsRequest); i {
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
		file_mock_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStubsResponse); i {
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
		file_mock_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stub); i {
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
		file_mock_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_mock_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
	file_mock_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*Input_Equals)(nil),
		(*Input_Contains)(nil),
		(*Input_Matches)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mock_proto_goTypes,
		DependencyIndexes: file_mock_proto_depIdxs,
		MessageInfos:      file_mock_proto_msgTypes,
	}.Build()
	File_mock_proto = out.File
	file_mock_proto_rawDesc = nil
	file_mock_proto_goTypes = nil
	file_mock_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MockClient is the client API for Mock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MockClient interface {
	AddStubs(ctx context.Context, in *AddStubsRequest, opts ...grpc.CallOption) (*AddStubsResponse, error)
	FindStubs(ctx context.Context, in *FindStubsRequest, opts ...grpc.CallOption) (*FindStubsResponse, error)
	DeleteStubs(ctx context.Context, in *DeleteStubsRequest, opts ...grpc.CallOption) (*DeleteStubsResponse, error)
}

type mockClient struct {
	cc grpc.ClientConnInterface
}

func NewMockClient(cc grpc.ClientConnInterface) MockClient {
	return &mockClient{cc}
}

func (c *mockClient) AddStubs(ctx context.Context, in *AddStubsRequest, opts ...grpc.CallOption) (*AddStubsResponse, error) {
	out := new(AddStubsResponse)
	err := c.cc.Invoke(ctx, "/github.com.kreexzoi.grpc_mock.Mock/AddStubs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockClient) FindStubs(ctx context.Context, in *FindStubsRequest, opts ...grpc.CallOption) (*FindStubsResponse, error) {
	out := new(FindStubsResponse)
	err := c.cc.Invoke(ctx, "/github.com.kreexzoi.grpc_mock.Mock/FindStubs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockClient) DeleteStubs(ctx context.Context, in *DeleteStubsRequest, opts ...grpc.CallOption) (*DeleteStubsResponse, error) {
	out := new(DeleteStubsResponse)
	err := c.cc.Invoke(ctx, "/github.com.kreexzoi.grpc_mock.Mock/DeleteStubs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MockServer is the server API for Mock service.
type MockServer interface {
	AddStubs(context.Context, *AddStubsRequest) (*AddStubsResponse, error)
	FindStubs(context.Context, *FindStubsRequest) (*FindStubsResponse, error)
	DeleteStubs(context.Context, *DeleteStubsRequest) (*DeleteStubsResponse, error)
}

// UnimplementedMockServer can be embedded to have forward compatible implementations.
type UnimplementedMockServer struct {
}

func (*UnimplementedMockServer) AddStubs(context.Context, *AddStubsRequest) (*AddStubsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStubs not implemented")
}
func (*UnimplementedMockServer) FindStubs(context.Context, *FindStubsRequest) (*FindStubsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStubs not implemented")
}
func (*UnimplementedMockServer) DeleteStubs(context.Context, *DeleteStubsRequest) (*DeleteStubsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStubs not implemented")
}

func RegisterMockServer(s *grpc.Server, srv MockServer) {
	s.RegisterService(&_Mock_serviceDesc, srv)
}

func _Mock_AddStubs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStubsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockServer).AddStubs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.kreexzoi.grpc_mock.Mock/AddStubs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockServer).AddStubs(ctx, req.(*AddStubsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mock_FindStubs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindStubsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockServer).FindStubs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.kreexzoi.grpc_mock.Mock/FindStubs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockServer).FindStubs(ctx, req.(*FindStubsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mock_DeleteStubs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStubsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockServer).DeleteStubs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.kreexzoi.grpc_mock.Mock/DeleteStubs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockServer).DeleteStubs(ctx, req.(*DeleteStubsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.kreexzoi.grpc_mock.Mock",
	HandlerType: (*MockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddStubs",
			Handler:    _Mock_AddStubs_Handler,
		},
		{
			MethodName: "FindStubs",
			Handler:    _Mock_FindStubs_Handler,
		},
		{
			MethodName: "DeleteStubs",
			Handler:    _Mock_DeleteStubs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mock.proto",
}
