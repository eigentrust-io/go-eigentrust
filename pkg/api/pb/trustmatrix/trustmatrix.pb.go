// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: trustmatrix.proto

package trustmatrixpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              *string  `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	TimestampQwords []uint64 `protobuf:"varint,2,rep,packed,name=timestamp_qwords,json=timestampQwords,proto3" json:"timestamp_qwords,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustmatrix_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_trustmatrix_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_trustmatrix_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Header) GetTimestampQwords() []uint64 {
	if x != nil {
		return x.TimestampQwords
	}
	return nil
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Truster string  `protobuf:"bytes,1,opt,name=truster,proto3" json:"truster,omitempty"`
	Trustee string  `protobuf:"bytes,2,opt,name=trustee,proto3" json:"trustee,omitempty"`
	Value   float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustmatrix_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_trustmatrix_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_trustmatrix_proto_rawDescGZIP(), []int{1}
}

func (x *Entry) GetTruster() string {
	if x != nil {
		return x.Truster
	}
	return ""
}

func (x *Entry) GetTrustee() string {
	if x != nil {
		return x.Trustee
	}
	return ""
}

func (x *Entry) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustmatrix_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trustmatrix_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_trustmatrix_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustmatrix_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trustmatrix_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_trustmatrix_proto_rawDescGZIP(), []int{3}
}

func (x *CreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustmatrix_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trustmatrix_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_trustmatrix_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Part:
	//
	//	*GetResponse_Header
	//	*GetResponse_Entry
	Part isGetResponse_Part `protobuf_oneof:"part"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustmatrix_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trustmatrix_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_trustmatrix_proto_rawDescGZIP(), []int{5}
}

func (m *GetResponse) GetPart() isGetResponse_Part {
	if m != nil {
		return m.Part
	}
	return nil
}

func (x *GetResponse) GetHeader() *Header {
	if x, ok := x.GetPart().(*GetResponse_Header); ok {
		return x.Header
	}
	return nil
}

func (x *GetResponse) GetEntry() *Entry {
	if x, ok := x.GetPart().(*GetResponse_Entry); ok {
		return x.Entry
	}
	return nil
}

type isGetResponse_Part interface {
	isGetResponse_Part()
}

type GetResponse_Header struct {
	Header *Header `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type GetResponse_Entry struct {
	Entry *Entry `protobuf:"bytes,2,opt,name=entry,proto3,oneof"`
}

func (*GetResponse_Header) isGetResponse_Part() {}

func (*GetResponse_Entry) isGetResponse_Part() {}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Entries []*Entry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustmatrix_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trustmatrix_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_trustmatrix_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRequest) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *UpdateRequest) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustmatrix_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trustmatrix_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_trustmatrix_proto_rawDescGZIP(), []int{7}
}

type FlushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FlushRequest) Reset() {
	*x = FlushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustmatrix_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushRequest) ProtoMessage() {}

func (x *FlushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trustmatrix_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushRequest.ProtoReflect.Descriptor instead.
func (*FlushRequest) Descriptor() ([]byte, []int) {
	return file_trustmatrix_proto_rawDescGZIP(), []int{8}
}

func (x *FlushRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FlushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlushResponse) Reset() {
	*x = FlushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustmatrix_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushResponse) ProtoMessage() {}

func (x *FlushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trustmatrix_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushResponse.ProtoReflect.Descriptor instead.
func (*FlushResponse) Descriptor() ([]byte, []int) {
	return file_trustmatrix_proto_rawDescGZIP(), []int{9}
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustmatrix_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trustmatrix_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_trustmatrix_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustmatrix_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trustmatrix_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_trustmatrix_proto_rawDescGZIP(), []int{11}
}

var File_trustmatrix_proto protoreflect.FileDescriptor

var file_trustmatrix_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x22, 0x4f, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x29, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x71, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x51, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69,
	0x64, 0x22, 0x51, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x70, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61, 0x74, 0x72,
	0x69, 0x78, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0x06, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x22, 0x6a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61,
	0x74, 0x72, 0x69, 0x78, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd8, 0x02, 0x0a, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x17, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a,
	0x05, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x12, 0x19, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61,
	0x74, 0x72, 0x69, 0x78, 0x2e, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e,
	0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x6b, 0x33, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x67,
	0x6f, 0x2d, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x3b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trustmatrix_proto_rawDescOnce sync.Once
	file_trustmatrix_proto_rawDescData = file_trustmatrix_proto_rawDesc
)

func file_trustmatrix_proto_rawDescGZIP() []byte {
	file_trustmatrix_proto_rawDescOnce.Do(func() {
		file_trustmatrix_proto_rawDescData = protoimpl.X.CompressGZIP(file_trustmatrix_proto_rawDescData)
	})
	return file_trustmatrix_proto_rawDescData
}

var file_trustmatrix_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_trustmatrix_proto_goTypes = []interface{}{
	(*Header)(nil),         // 0: trustmatrix.Header
	(*Entry)(nil),          // 1: trustmatrix.Entry
	(*CreateRequest)(nil),  // 2: trustmatrix.CreateRequest
	(*CreateResponse)(nil), // 3: trustmatrix.CreateResponse
	(*GetRequest)(nil),     // 4: trustmatrix.GetRequest
	(*GetResponse)(nil),    // 5: trustmatrix.GetResponse
	(*UpdateRequest)(nil),  // 6: trustmatrix.UpdateRequest
	(*UpdateResponse)(nil), // 7: trustmatrix.UpdateResponse
	(*FlushRequest)(nil),   // 8: trustmatrix.FlushRequest
	(*FlushResponse)(nil),  // 9: trustmatrix.FlushResponse
	(*DeleteRequest)(nil),  // 10: trustmatrix.DeleteRequest
	(*DeleteResponse)(nil), // 11: trustmatrix.DeleteResponse
}
var file_trustmatrix_proto_depIdxs = []int32{
	0,  // 0: trustmatrix.GetResponse.header:type_name -> trustmatrix.Header
	1,  // 1: trustmatrix.GetResponse.entry:type_name -> trustmatrix.Entry
	0,  // 2: trustmatrix.UpdateRequest.header:type_name -> trustmatrix.Header
	1,  // 3: trustmatrix.UpdateRequest.entries:type_name -> trustmatrix.Entry
	2,  // 4: trustmatrix.Service.Create:input_type -> trustmatrix.CreateRequest
	4,  // 5: trustmatrix.Service.Get:input_type -> trustmatrix.GetRequest
	6,  // 6: trustmatrix.Service.Update:input_type -> trustmatrix.UpdateRequest
	8,  // 7: trustmatrix.Service.Flush:input_type -> trustmatrix.FlushRequest
	10, // 8: trustmatrix.Service.Delete:input_type -> trustmatrix.DeleteRequest
	3,  // 9: trustmatrix.Service.Create:output_type -> trustmatrix.CreateResponse
	5,  // 10: trustmatrix.Service.Get:output_type -> trustmatrix.GetResponse
	7,  // 11: trustmatrix.Service.Update:output_type -> trustmatrix.UpdateResponse
	9,  // 12: trustmatrix.Service.Flush:output_type -> trustmatrix.FlushResponse
	11, // 13: trustmatrix.Service.Delete:output_type -> trustmatrix.DeleteResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_trustmatrix_proto_init() }
func file_trustmatrix_proto_init() {
	if File_trustmatrix_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trustmatrix_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_trustmatrix_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_trustmatrix_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_trustmatrix_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_trustmatrix_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_trustmatrix_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_trustmatrix_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_trustmatrix_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_trustmatrix_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushRequest); i {
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
		file_trustmatrix_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushResponse); i {
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
		file_trustmatrix_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_trustmatrix_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
	file_trustmatrix_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_trustmatrix_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*GetResponse_Header)(nil),
		(*GetResponse_Entry)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trustmatrix_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trustmatrix_proto_goTypes,
		DependencyIndexes: file_trustmatrix_proto_depIdxs,
		MessageInfos:      file_trustmatrix_proto_msgTypes,
	}.Build()
	File_trustmatrix_proto = out.File
	file_trustmatrix_proto_rawDesc = nil
	file_trustmatrix_proto_goTypes = nil
	file_trustmatrix_proto_depIdxs = nil
}
