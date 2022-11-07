// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: api/create/v1/framer.proto

package v1

import (
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

type CreateFramerDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address        string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	InsgramAccount string `protobuf:"bytes,2,opt,name=insgramAccount,proto3" json:"insgramAccount,omitempty"`
}

func (x *CreateFramerDTO) Reset() {
	*x = CreateFramerDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_framer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFramerDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFramerDTO) ProtoMessage() {}

func (x *CreateFramerDTO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_framer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFramerDTO.ProtoReflect.Descriptor instead.
func (*CreateFramerDTO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_framer_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFramerDTO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateFramerDTO) GetInsgramAccount() string {
	if x != nil {
		return x.InsgramAccount
	}
	return ""
}

type CreateFramerVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateFramerVO) Reset() {
	*x = CreateFramerVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_framer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFramerVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFramerVO) ProtoMessage() {}

func (x *CreateFramerVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_framer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFramerVO.ProtoReflect.Descriptor instead.
func (*CreateFramerVO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_framer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFramerVO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateFramerDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address        string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	InsgramAccount *string `protobuf:"bytes,2,opt,name=insgramAccount,proto3,oneof" json:"insgramAccount,omitempty"`
	CreationNumber *int32  `protobuf:"varint,3,opt,name=creationNumber,proto3,oneof" json:"creationNumber,omitempty"`
	MaxPoll        *int32  `protobuf:"varint,4,opt,name=maxPoll,proto3,oneof" json:"maxPoll,omitempty"`
	Credit         *int32  `protobuf:"varint,5,opt,name=credit,proto3,oneof" json:"credit,omitempty"`
}

func (x *UpdateFramerDTO) Reset() {
	*x = UpdateFramerDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_framer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFramerDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFramerDTO) ProtoMessage() {}

func (x *UpdateFramerDTO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_framer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFramerDTO.ProtoReflect.Descriptor instead.
func (*UpdateFramerDTO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_framer_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateFramerDTO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateFramerDTO) GetInsgramAccount() string {
	if x != nil && x.InsgramAccount != nil {
		return *x.InsgramAccount
	}
	return ""
}

func (x *UpdateFramerDTO) GetCreationNumber() int32 {
	if x != nil && x.CreationNumber != nil {
		return *x.CreationNumber
	}
	return 0
}

func (x *UpdateFramerDTO) GetMaxPoll() int32 {
	if x != nil && x.MaxPoll != nil {
		return *x.MaxPoll
	}
	return 0
}

func (x *UpdateFramerDTO) GetCredit() int32 {
	if x != nil && x.Credit != nil {
		return *x.Credit
	}
	return 0
}

type UpdateFramerVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *UpdateFramerVO) Reset() {
	*x = UpdateFramerVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_framer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFramerVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFramerVO) ProtoMessage() {}

func (x *UpdateFramerVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_framer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFramerVO.ProtoReflect.Descriptor instead.
func (*UpdateFramerVO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_framer_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFramerVO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetFramerDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetFramerDTO) Reset() {
	*x = GetFramerDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_framer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFramerDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFramerDTO) ProtoMessage() {}

func (x *GetFramerDTO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_framer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFramerDTO.ProtoReflect.Descriptor instead.
func (*GetFramerDTO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_framer_proto_rawDescGZIP(), []int{4}
}

func (x *GetFramerDTO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetFramerVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address        string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CreationNumber int32  `protobuf:"varint,2,opt,name=creationNumber,proto3" json:"creationNumber,omitempty"`
	MaxPoll        int32  `protobuf:"varint,3,opt,name=maxPoll,proto3" json:"maxPoll,omitempty"`
	Credit         int32  `protobuf:"varint,4,opt,name=credit,proto3" json:"credit,omitempty"`
}

func (x *GetFramerVO) Reset() {
	*x = GetFramerVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_framer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFramerVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFramerVO) ProtoMessage() {}

func (x *GetFramerVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_framer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFramerVO.ProtoReflect.Descriptor instead.
func (*GetFramerVO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_framer_proto_rawDescGZIP(), []int{5}
}

func (x *GetFramerVO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetFramerVO) GetCreationNumber() int32 {
	if x != nil {
		return x.CreationNumber
	}
	return 0
}

func (x *GetFramerVO) GetMaxPoll() int32 {
	if x != nil {
		return x.MaxPoll
	}
	return 0
}

func (x *GetFramerVO) GetCredit() int32 {
	if x != nil {
		return x.Credit
	}
	return 0
}

type ListFramerDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PageNum  int32  `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListFramerDTO) Reset() {
	*x = ListFramerDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_framer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFramerDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFramerDTO) ProtoMessage() {}

func (x *ListFramerDTO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_framer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFramerDTO.ProtoReflect.Descriptor instead.
func (*ListFramerDTO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_framer_proto_rawDescGZIP(), []int{6}
}

func (x *ListFramerDTO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ListFramerDTO) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListFramerDTO) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListFramerVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total      int64         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageNum    int32         `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize   int32         `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	FramerInfo []*FramerInfo `protobuf:"bytes,4,rep,name=framerInfo,proto3" json:"framerInfo,omitempty"`
}

func (x *ListFramerVO) Reset() {
	*x = ListFramerVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_framer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFramerVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFramerVO) ProtoMessage() {}

func (x *ListFramerVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_framer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFramerVO.ProtoReflect.Descriptor instead.
func (*ListFramerVO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_framer_proto_rawDescGZIP(), []int{7}
}

func (x *ListFramerVO) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListFramerVO) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListFramerVO) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListFramerVO) GetFramerInfo() []*FramerInfo {
	if x != nil {
		return x.FramerInfo
	}
	return nil
}

type FramerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address        string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	InsgramAccount string `protobuf:"bytes,3,opt,name=insgramAccount,proto3" json:"insgramAccount,omitempty"`
	CreationNumber int32  `protobuf:"varint,4,opt,name=creationNumber,proto3" json:"creationNumber,omitempty"`
	MaxPoll        int32  `protobuf:"varint,5,opt,name=maxPoll,proto3" json:"maxPoll,omitempty"`
	Credit         int32  `protobuf:"varint,6,opt,name=credit,proto3" json:"credit,omitempty"`
}

func (x *FramerInfo) Reset() {
	*x = FramerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_framer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FramerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FramerInfo) ProtoMessage() {}

func (x *FramerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_framer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FramerInfo.ProtoReflect.Descriptor instead.
func (*FramerInfo) Descriptor() ([]byte, []int) {
	return file_api_create_v1_framer_proto_rawDescGZIP(), []int{8}
}

func (x *FramerInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FramerInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *FramerInfo) GetInsgramAccount() string {
	if x != nil {
		return x.InsgramAccount
	}
	return ""
}

func (x *FramerInfo) GetCreationNumber() int32 {
	if x != nil {
		return x.CreationNumber
	}
	return 0
}

func (x *FramerInfo) GetMaxPoll() int32 {
	if x != nil {
		return x.MaxPoll
	}
	return 0
}

func (x *FramerInfo) GetCredit() int32 {
	if x != nil {
		return x.Credit
	}
	return 0
}

var File_api_create_v1_framer_proto protoreflect.FileDescriptor

var file_api_create_v1_framer_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x53, 0x0a, 0x0f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x67,
	0x72, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x69, 0x6e, 0x73, 0x67, 0x72, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x20, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72,
	0x56, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xfe, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x2b, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x67, 0x72, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x67,
	0x72, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a,
	0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x61,
	0x78, 0x50, 0x6f, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x07, 0x6d,
	0x61, 0x78, 0x50, 0x6f, 0x6c, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x06, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x69, 0x6e, 0x73, 0x67, 0x72,
	0x61, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x6d, 0x61, 0x78, 0x50, 0x6f, 0x6c, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x72, 0x56, 0x4f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x28, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x4f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x61, 0x78, 0x50, 0x6f, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61,
	0x78, 0x50, 0x6f, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x22, 0x5f, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x95,
	0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x4f, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x0a, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x67, 0x72, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x67, 0x72, 0x61, 0x6d,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x50, 0x6f, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6d, 0x61, 0x78, 0x50, 0x6f, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x32, 0xb5, 0x02, 0x0a, 0x06, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x1a, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x4f, 0x12, 0x4d, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x1a, 0x1d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x4f, 0x12, 0x44, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x72, 0x44, 0x54, 0x4f, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x4f,
	0x12, 0x47, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x1a, 0x1b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x4f, 0x42, 0x31, 0x0a, 0x0d, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_create_v1_framer_proto_rawDescOnce sync.Once
	file_api_create_v1_framer_proto_rawDescData = file_api_create_v1_framer_proto_rawDesc
)

func file_api_create_v1_framer_proto_rawDescGZIP() []byte {
	file_api_create_v1_framer_proto_rawDescOnce.Do(func() {
		file_api_create_v1_framer_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_create_v1_framer_proto_rawDescData)
	})
	return file_api_create_v1_framer_proto_rawDescData
}

var file_api_create_v1_framer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_create_v1_framer_proto_goTypes = []interface{}{
	(*CreateFramerDTO)(nil), // 0: api.create.v1.CreateFramerDTO
	(*CreateFramerVO)(nil),  // 1: api.create.v1.CreateFramerVO
	(*UpdateFramerDTO)(nil), // 2: api.create.v1.UpdateFramerDTO
	(*UpdateFramerVO)(nil),  // 3: api.create.v1.UpdateFramerVO
	(*GetFramerDTO)(nil),    // 4: api.create.v1.GetFramerDTO
	(*GetFramerVO)(nil),     // 5: api.create.v1.GetFramerVO
	(*ListFramerDTO)(nil),   // 6: api.create.v1.ListFramerDTO
	(*ListFramerVO)(nil),    // 7: api.create.v1.ListFramerVO
	(*FramerInfo)(nil),      // 8: api.create.v1.FramerInfo
}
var file_api_create_v1_framer_proto_depIdxs = []int32{
	8, // 0: api.create.v1.ListFramerVO.framerInfo:type_name -> api.create.v1.FramerInfo
	0, // 1: api.create.v1.Framer.CreateFramer:input_type -> api.create.v1.CreateFramerDTO
	2, // 2: api.create.v1.Framer.UpdateFramer:input_type -> api.create.v1.UpdateFramerDTO
	4, // 3: api.create.v1.Framer.GetFramer:input_type -> api.create.v1.GetFramerDTO
	6, // 4: api.create.v1.Framer.ListFramer:input_type -> api.create.v1.ListFramerDTO
	1, // 5: api.create.v1.Framer.CreateFramer:output_type -> api.create.v1.CreateFramerVO
	3, // 6: api.create.v1.Framer.UpdateFramer:output_type -> api.create.v1.UpdateFramerVO
	5, // 7: api.create.v1.Framer.GetFramer:output_type -> api.create.v1.GetFramerVO
	7, // 8: api.create.v1.Framer.ListFramer:output_type -> api.create.v1.ListFramerVO
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_create_v1_framer_proto_init() }
func file_api_create_v1_framer_proto_init() {
	if File_api_create_v1_framer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_create_v1_framer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFramerDTO); i {
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
		file_api_create_v1_framer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFramerVO); i {
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
		file_api_create_v1_framer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFramerDTO); i {
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
		file_api_create_v1_framer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFramerVO); i {
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
		file_api_create_v1_framer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFramerDTO); i {
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
		file_api_create_v1_framer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFramerVO); i {
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
		file_api_create_v1_framer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFramerDTO); i {
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
		file_api_create_v1_framer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFramerVO); i {
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
		file_api_create_v1_framer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FramerInfo); i {
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
	file_api_create_v1_framer_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_create_v1_framer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_create_v1_framer_proto_goTypes,
		DependencyIndexes: file_api_create_v1_framer_proto_depIdxs,
		MessageInfos:      file_api_create_v1_framer_proto_msgTypes,
	}.Build()
	File_api_create_v1_framer_proto = out.File
	file_api_create_v1_framer_proto_rawDesc = nil
	file_api_create_v1_framer_proto_goTypes = nil
	file_api_create_v1_framer_proto_depIdxs = nil
}
