// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: api/create/v1/apply_record.proto

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

type CreateApplyRecordDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationActivityId int64  `protobuf:"varint,1,opt,name=creationActivityId,proto3" json:"creationActivityId"`
	Address            string `protobuf:"bytes,2,opt,name=address,proto3" json:"address"`
}

func (x *CreateApplyRecordDTO) Reset() {
	*x = CreateApplyRecordDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_apply_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApplyRecordDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApplyRecordDTO) ProtoMessage() {}

func (x *CreateApplyRecordDTO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_apply_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApplyRecordDTO.ProtoReflect.Descriptor instead.
func (*CreateApplyRecordDTO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_apply_record_proto_rawDescGZIP(), []int{0}
}

func (x *CreateApplyRecordDTO) GetCreationActivityId() int64 {
	if x != nil {
		return x.CreationActivityId
	}
	return 0
}

func (x *CreateApplyRecordDTO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type CreateApplyRecordVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *CreateApplyRecordVO) Reset() {
	*x = CreateApplyRecordVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_apply_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApplyRecordVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApplyRecordVO) ProtoMessage() {}

func (x *CreateApplyRecordVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_apply_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApplyRecordVO.ProtoReflect.Descriptor instead.
func (*CreateApplyRecordVO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_apply_record_proto_rawDescGZIP(), []int{1}
}

func (x *CreateApplyRecordVO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateApplyRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateApplyRecordRequest) Reset() {
	*x = UpdateApplyRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_apply_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateApplyRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateApplyRecordRequest) ProtoMessage() {}

func (x *UpdateApplyRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_apply_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateApplyRecordRequest.ProtoReflect.Descriptor instead.
func (*UpdateApplyRecordRequest) Descriptor() ([]byte, []int) {
	return file_api_create_v1_apply_record_proto_rawDescGZIP(), []int{2}
}

type UpdateApplyRecordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateApplyRecordReply) Reset() {
	*x = UpdateApplyRecordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_apply_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateApplyRecordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateApplyRecordReply) ProtoMessage() {}

func (x *UpdateApplyRecordReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_apply_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateApplyRecordReply.ProtoReflect.Descriptor instead.
func (*UpdateApplyRecordReply) Descriptor() ([]byte, []int) {
	return file_api_create_v1_apply_record_proto_rawDescGZIP(), []int{3}
}

type DeleteApplyRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteApplyRecordRequest) Reset() {
	*x = DeleteApplyRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_apply_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApplyRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApplyRecordRequest) ProtoMessage() {}

func (x *DeleteApplyRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_apply_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApplyRecordRequest.ProtoReflect.Descriptor instead.
func (*DeleteApplyRecordRequest) Descriptor() ([]byte, []int) {
	return file_api_create_v1_apply_record_proto_rawDescGZIP(), []int{4}
}

type DeleteApplyRecordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteApplyRecordReply) Reset() {
	*x = DeleteApplyRecordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_apply_record_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApplyRecordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApplyRecordReply) ProtoMessage() {}

func (x *DeleteApplyRecordReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_apply_record_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApplyRecordReply.ProtoReflect.Descriptor instead.
func (*DeleteApplyRecordReply) Descriptor() ([]byte, []int) {
	return file_api_create_v1_apply_record_proto_rawDescGZIP(), []int{5}
}

type GetApplyRecordDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationActivityId int64  `protobuf:"varint,1,opt,name=creationActivityId,proto3" json:"creationActivityId"`
	Address            string `protobuf:"bytes,2,opt,name=address,proto3" json:"address"`
}

func (x *GetApplyRecordDTO) Reset() {
	*x = GetApplyRecordDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_apply_record_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplyRecordDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplyRecordDTO) ProtoMessage() {}

func (x *GetApplyRecordDTO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_apply_record_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplyRecordDTO.ProtoReflect.Descriptor instead.
func (*GetApplyRecordDTO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_apply_record_proto_rawDescGZIP(), []int{6}
}

func (x *GetApplyRecordDTO) GetCreationActivityId() int64 {
	if x != nil {
		return x.CreationActivityId
	}
	return 0
}

func (x *GetApplyRecordDTO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetApplyRecordVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CreationActivityId int64  `protobuf:"varint,2,opt,name=creationActivityId,proto3" json:"creationActivityId"`
	Address            string `protobuf:"bytes,3,opt,name=address,proto3" json:"address"`
}

func (x *GetApplyRecordVO) Reset() {
	*x = GetApplyRecordVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_apply_record_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplyRecordVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplyRecordVO) ProtoMessage() {}

func (x *GetApplyRecordVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_apply_record_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplyRecordVO.ProtoReflect.Descriptor instead.
func (*GetApplyRecordVO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_apply_record_proto_rawDescGZIP(), []int{7}
}

func (x *GetApplyRecordVO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetApplyRecordVO) GetCreationActivityId() int64 {
	if x != nil {
		return x.CreationActivityId
	}
	return 0
}

func (x *GetApplyRecordVO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ListApplyRecordDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationActivityId []int64 `protobuf:"varint,1,rep,packed,name=creationActivityId,proto3" json:"creationActivityId"`
	Address            string  `protobuf:"bytes,2,opt,name=address,proto3" json:"address"`
}

func (x *ListApplyRecordDTO) Reset() {
	*x = ListApplyRecordDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_apply_record_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplyRecordDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplyRecordDTO) ProtoMessage() {}

func (x *ListApplyRecordDTO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_apply_record_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplyRecordDTO.ProtoReflect.Descriptor instead.
func (*ListApplyRecordDTO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_apply_record_proto_rawDescGZIP(), []int{8}
}

func (x *ListApplyRecordDTO) GetCreationActivityId() []int64 {
	if x != nil {
		return x.CreationActivityId
	}
	return nil
}

func (x *ListApplyRecordDTO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ListApplyRecordVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyRecordInfo []*ApplyRecordInfo `protobuf:"bytes,1,rep,name=applyRecordInfo,proto3" json:"applyRecordInfo"`
}

func (x *ListApplyRecordVO) Reset() {
	*x = ListApplyRecordVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_apply_record_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplyRecordVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplyRecordVO) ProtoMessage() {}

func (x *ListApplyRecordVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_apply_record_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplyRecordVO.ProtoReflect.Descriptor instead.
func (*ListApplyRecordVO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_apply_record_proto_rawDescGZIP(), []int{9}
}

func (x *ListApplyRecordVO) GetApplyRecordInfo() []*ApplyRecordInfo {
	if x != nil {
		return x.ApplyRecordInfo
	}
	return nil
}

type ApplyRecordInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CreationActivityId int64  `protobuf:"varint,2,opt,name=creationActivityId,proto3" json:"creationActivityId"`
	Address            string `protobuf:"bytes,3,opt,name=address,proto3" json:"address"`
}

func (x *ApplyRecordInfo) Reset() {
	*x = ApplyRecordInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_apply_record_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyRecordInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyRecordInfo) ProtoMessage() {}

func (x *ApplyRecordInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_apply_record_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyRecordInfo.ProtoReflect.Descriptor instead.
func (*ApplyRecordInfo) Descriptor() ([]byte, []int) {
	return file_api_create_v1_apply_record_proto_rawDescGZIP(), []int{10}
}

func (x *ApplyRecordInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ApplyRecordInfo) GetCreationActivityId() int64 {
	if x != nil {
		return x.CreationActivityId
	}
	return 0
}

func (x *ApplyRecordInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_api_create_v1_apply_record_proto protoreflect.FileDescriptor

var file_api_create_v1_apply_record_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x22, 0x60, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x54, 0x4f, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x56, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x5d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x54, 0x4f, 0x12, 0x2e, 0x0a, 0x12, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x56, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x5e, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x54, 0x4f, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x5d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x56, 0x4f, 0x12, 0x48, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x6b, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32,
	0xe2, 0x03, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x5c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x54, 0x4f, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x56, 0x4f, 0x12, 0x63, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x63, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x54, 0x4f, 0x1a, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x56, 0x4f, 0x12, 0x56, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44,
	0x54, 0x4f, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x56, 0x4f, 0x42, 0x37, 0x0a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x23, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_create_v1_apply_record_proto_rawDescOnce sync.Once
	file_api_create_v1_apply_record_proto_rawDescData = file_api_create_v1_apply_record_proto_rawDesc
)

func file_api_create_v1_apply_record_proto_rawDescGZIP() []byte {
	file_api_create_v1_apply_record_proto_rawDescOnce.Do(func() {
		file_api_create_v1_apply_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_create_v1_apply_record_proto_rawDescData)
	})
	return file_api_create_v1_apply_record_proto_rawDescData
}

var file_api_create_v1_apply_record_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_create_v1_apply_record_proto_goTypes = []interface{}{
	(*CreateApplyRecordDTO)(nil),     // 0: api.create.v1.CreateApplyRecordDTO
	(*CreateApplyRecordVO)(nil),      // 1: api.create.v1.CreateApplyRecordVO
	(*UpdateApplyRecordRequest)(nil), // 2: api.create.v1.UpdateApplyRecordRequest
	(*UpdateApplyRecordReply)(nil),   // 3: api.create.v1.UpdateApplyRecordReply
	(*DeleteApplyRecordRequest)(nil), // 4: api.create.v1.DeleteApplyRecordRequest
	(*DeleteApplyRecordReply)(nil),   // 5: api.create.v1.DeleteApplyRecordReply
	(*GetApplyRecordDTO)(nil),        // 6: api.create.v1.GetApplyRecordDTO
	(*GetApplyRecordVO)(nil),         // 7: api.create.v1.GetApplyRecordVO
	(*ListApplyRecordDTO)(nil),       // 8: api.create.v1.ListApplyRecordDTO
	(*ListApplyRecordVO)(nil),        // 9: api.create.v1.ListApplyRecordVO
	(*ApplyRecordInfo)(nil),          // 10: api.create.v1.ApplyRecordInfo
}
var file_api_create_v1_apply_record_proto_depIdxs = []int32{
	10, // 0: api.create.v1.ListApplyRecordVO.applyRecordInfo:type_name -> api.create.v1.ApplyRecordInfo
	0,  // 1: api.create.v1.ApplyRecord.CreateApplyRecord:input_type -> api.create.v1.CreateApplyRecordDTO
	2,  // 2: api.create.v1.ApplyRecord.UpdateApplyRecord:input_type -> api.create.v1.UpdateApplyRecordRequest
	4,  // 3: api.create.v1.ApplyRecord.DeleteApplyRecord:input_type -> api.create.v1.DeleteApplyRecordRequest
	6,  // 4: api.create.v1.ApplyRecord.GetApplyRecord:input_type -> api.create.v1.GetApplyRecordDTO
	8,  // 5: api.create.v1.ApplyRecord.ListApplyRecord:input_type -> api.create.v1.ListApplyRecordDTO
	1,  // 6: api.create.v1.ApplyRecord.CreateApplyRecord:output_type -> api.create.v1.CreateApplyRecordVO
	3,  // 7: api.create.v1.ApplyRecord.UpdateApplyRecord:output_type -> api.create.v1.UpdateApplyRecordReply
	5,  // 8: api.create.v1.ApplyRecord.DeleteApplyRecord:output_type -> api.create.v1.DeleteApplyRecordReply
	7,  // 9: api.create.v1.ApplyRecord.GetApplyRecord:output_type -> api.create.v1.GetApplyRecordVO
	9,  // 10: api.create.v1.ApplyRecord.ListApplyRecord:output_type -> api.create.v1.ListApplyRecordVO
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_create_v1_apply_record_proto_init() }
func file_api_create_v1_apply_record_proto_init() {
	if File_api_create_v1_apply_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_create_v1_apply_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApplyRecordDTO); i {
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
		file_api_create_v1_apply_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApplyRecordVO); i {
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
		file_api_create_v1_apply_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateApplyRecordRequest); i {
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
		file_api_create_v1_apply_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateApplyRecordReply); i {
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
		file_api_create_v1_apply_record_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApplyRecordRequest); i {
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
		file_api_create_v1_apply_record_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApplyRecordReply); i {
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
		file_api_create_v1_apply_record_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplyRecordDTO); i {
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
		file_api_create_v1_apply_record_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplyRecordVO); i {
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
		file_api_create_v1_apply_record_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplyRecordDTO); i {
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
		file_api_create_v1_apply_record_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplyRecordVO); i {
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
		file_api_create_v1_apply_record_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyRecordInfo); i {
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
			RawDescriptor: file_api_create_v1_apply_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_create_v1_apply_record_proto_goTypes,
		DependencyIndexes: file_api_create_v1_apply_record_proto_depIdxs,
		MessageInfos:      file_api_create_v1_apply_record_proto_msgTypes,
	}.Build()
	File_api_create_v1_apply_record_proto = out.File
	file_api_create_v1_apply_record_proto_rawDesc = nil
	file_api_create_v1_apply_record_proto_goTypes = nil
	file_api_create_v1_apply_record_proto_depIdxs = nil
}
