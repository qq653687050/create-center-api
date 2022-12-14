// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: api/create/v1/words.proto

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

type CreateWordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateWordsRequest) Reset() {
	*x = CreateWordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_words_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWordsRequest) ProtoMessage() {}

func (x *CreateWordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_words_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWordsRequest.ProtoReflect.Descriptor instead.
func (*CreateWordsRequest) Descriptor() ([]byte, []int) {
	return file_api_create_v1_words_proto_rawDescGZIP(), []int{0}
}

type CreateWordsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateWordsReply) Reset() {
	*x = CreateWordsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_words_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWordsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWordsReply) ProtoMessage() {}

func (x *CreateWordsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_words_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWordsReply.ProtoReflect.Descriptor instead.
func (*CreateWordsReply) Descriptor() ([]byte, []int) {
	return file_api_create_v1_words_proto_rawDescGZIP(), []int{1}
}

type UpdateWordsPollDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Num int32 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *UpdateWordsPollDTO) Reset() {
	*x = UpdateWordsPollDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_words_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWordsPollDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWordsPollDTO) ProtoMessage() {}

func (x *UpdateWordsPollDTO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_words_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWordsPollDTO.ProtoReflect.Descriptor instead.
func (*UpdateWordsPollDTO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_words_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateWordsPollDTO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateWordsPollDTO) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type UpdateWordsPollVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateWordsPollVO) Reset() {
	*x = UpdateWordsPollVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_words_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWordsPollVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWordsPollVO) ProtoMessage() {}

func (x *UpdateWordsPollVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_words_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWordsPollVO.ProtoReflect.Descriptor instead.
func (*UpdateWordsPollVO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_words_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateWordsPollVO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteWordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteWordsRequest) Reset() {
	*x = DeleteWordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_words_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWordsRequest) ProtoMessage() {}

func (x *DeleteWordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_words_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWordsRequest.ProtoReflect.Descriptor instead.
func (*DeleteWordsRequest) Descriptor() ([]byte, []int) {
	return file_api_create_v1_words_proto_rawDescGZIP(), []int{4}
}

type DeleteWordsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteWordsReply) Reset() {
	*x = DeleteWordsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_words_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWordsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWordsReply) ProtoMessage() {}

func (x *DeleteWordsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_words_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWordsReply.ProtoReflect.Descriptor instead.
func (*DeleteWordsReply) Descriptor() ([]byte, []int) {
	return file_api_create_v1_words_proto_rawDescGZIP(), []int{5}
}

type GetWordsDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWordsDTO) Reset() {
	*x = GetWordsDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_words_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWordsDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWordsDTO) ProtoMessage() {}

func (x *GetWordsDTO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_words_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWordsDTO.ProtoReflect.Descriptor instead.
func (*GetWordsDTO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_words_proto_rawDescGZIP(), []int{6}
}

func (x *GetWordsDTO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetWordsVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Theme              string `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
	IsPublish          int32  `protobuf:"varint,3,opt,name=isPublish,proto3" json:"isPublish,omitempty"`
	CreationActivityId int64  `protobuf:"varint,4,opt,name=creationActivityId,proto3" json:"creationActivityId,omitempty"`
	AdvancedParameter  string `protobuf:"bytes,5,opt,name=advancedParameter,proto3" json:"advancedParameter,omitempty"`
	CreationPicture    string `protobuf:"bytes,6,opt,name=creationPicture,proto3" json:"creationPicture,omitempty"`
	Prompt             string `protobuf:"bytes,7,opt,name=prompt,proto3" json:"prompt,omitempty"`
	PromptPrivate      int32  `protobuf:"varint,8,opt,name=promptPrivate,proto3" json:"promptPrivate,omitempty"`
	Address            string `protobuf:"bytes,9,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetWordsVO) Reset() {
	*x = GetWordsVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_words_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWordsVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWordsVO) ProtoMessage() {}

func (x *GetWordsVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_words_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWordsVO.ProtoReflect.Descriptor instead.
func (*GetWordsVO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_words_proto_rawDescGZIP(), []int{7}
}

func (x *GetWordsVO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetWordsVO) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

func (x *GetWordsVO) GetIsPublish() int32 {
	if x != nil {
		return x.IsPublish
	}
	return 0
}

func (x *GetWordsVO) GetCreationActivityId() int64 {
	if x != nil {
		return x.CreationActivityId
	}
	return 0
}

func (x *GetWordsVO) GetAdvancedParameter() string {
	if x != nil {
		return x.AdvancedParameter
	}
	return ""
}

func (x *GetWordsVO) GetCreationPicture() string {
	if x != nil {
		return x.CreationPicture
	}
	return ""
}

func (x *GetWordsVO) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *GetWordsVO) GetPromptPrivate() int32 {
	if x != nil {
		return x.PromptPrivate
	}
	return 0
}

func (x *GetWordsVO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ListWordsDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address            string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	IsPublish          *int32 `protobuf:"varint,3,opt,name=isPublish,proto3,oneof" json:"isPublish,omitempty"`
	PollRank           string `protobuf:"bytes,4,opt,name=pollRank,proto3" json:"pollRank,omitempty"`
	PageNum            int32  `protobuf:"varint,5,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize           int32  `protobuf:"varint,6,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	CreationActivityId *int64 `protobuf:"varint,7,opt,name=creationActivityId,proto3,oneof" json:"creationActivityId,omitempty"`
	IdRank             *bool  `protobuf:"varint,8,opt,name=idRank,proto3,oneof" json:"idRank,omitempty"`
	PublishTimeRank    *bool  `protobuf:"varint,9,opt,name=publishTimeRank,proto3,oneof" json:"publishTimeRank,omitempty"`
}

func (x *ListWordsDTO) Reset() {
	*x = ListWordsDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_words_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWordsDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWordsDTO) ProtoMessage() {}

func (x *ListWordsDTO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_words_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWordsDTO.ProtoReflect.Descriptor instead.
func (*ListWordsDTO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_words_proto_rawDescGZIP(), []int{8}
}

func (x *ListWordsDTO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListWordsDTO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ListWordsDTO) GetIsPublish() int32 {
	if x != nil && x.IsPublish != nil {
		return *x.IsPublish
	}
	return 0
}

func (x *ListWordsDTO) GetPollRank() string {
	if x != nil {
		return x.PollRank
	}
	return ""
}

func (x *ListWordsDTO) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListWordsDTO) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListWordsDTO) GetCreationActivityId() int64 {
	if x != nil && x.CreationActivityId != nil {
		return *x.CreationActivityId
	}
	return 0
}

func (x *ListWordsDTO) GetIdRank() bool {
	if x != nil && x.IdRank != nil {
		return *x.IdRank
	}
	return false
}

func (x *ListWordsDTO) GetPublishTimeRank() bool {
	if x != nil && x.PublishTimeRank != nil {
		return *x.PublishTimeRank
	}
	return false
}

type ListWordsVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     int64        `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageNum   int32        `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize  int32        `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	WordsInfo []*WordsInfo `protobuf:"bytes,4,rep,name=wordsInfo,proto3" json:"wordsInfo,omitempty"`
}

func (x *ListWordsVO) Reset() {
	*x = ListWordsVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_words_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWordsVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWordsVO) ProtoMessage() {}

func (x *ListWordsVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_words_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWordsVO.ProtoReflect.Descriptor instead.
func (*ListWordsVO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_words_proto_rawDescGZIP(), []int{9}
}

func (x *ListWordsVO) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListWordsVO) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListWordsVO) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListWordsVO) GetWordsInfo() []*WordsInfo {
	if x != nil {
		return x.WordsInfo
	}
	return nil
}

type WordsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Theme             string `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
	IsPublish         int32  `protobuf:"varint,3,opt,name=isPublish,proto3" json:"isPublish,omitempty"`
	Poll              int32  `protobuf:"varint,4,opt,name=poll,proto3" json:"poll,omitempty"`
	PublishTime       string `protobuf:"bytes,5,opt,name=publishTime,proto3" json:"publishTime,omitempty"`
	Address           string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	CreationPicture   string `protobuf:"bytes,7,opt,name=creationPicture,proto3" json:"creationPicture,omitempty"`
	Prompt            string `protobuf:"bytes,8,opt,name=prompt,proto3" json:"prompt,omitempty"`
	PromptPrivate     int32  `protobuf:"varint,9,opt,name=promptPrivate,proto3" json:"promptPrivate,omitempty"`
	AdvancedParameter string `protobuf:"bytes,10,opt,name=advancedParameter,proto3" json:"advancedParameter,omitempty"`
}

func (x *WordsInfo) Reset() {
	*x = WordsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_words_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WordsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WordsInfo) ProtoMessage() {}

func (x *WordsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_words_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WordsInfo.ProtoReflect.Descriptor instead.
func (*WordsInfo) Descriptor() ([]byte, []int) {
	return file_api_create_v1_words_proto_rawDescGZIP(), []int{10}
}

func (x *WordsInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WordsInfo) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

func (x *WordsInfo) GetIsPublish() int32 {
	if x != nil {
		return x.IsPublish
	}
	return 0
}

func (x *WordsInfo) GetPoll() int32 {
	if x != nil {
		return x.Poll
	}
	return 0
}

func (x *WordsInfo) GetPublishTime() string {
	if x != nil {
		return x.PublishTime
	}
	return ""
}

func (x *WordsInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *WordsInfo) GetCreationPicture() string {
	if x != nil {
		return x.CreationPicture
	}
	return ""
}

func (x *WordsInfo) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *WordsInfo) GetPromptPrivate() int32 {
	if x != nil {
		return x.PromptPrivate
	}
	return 0
}

func (x *WordsInfo) GetAdvancedParameter() string {
	if x != nil {
		return x.AdvancedParameter
	}
	return ""
}

type ListWordsByIdsDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WordsIds []string `protobuf:"bytes,1,rep,name=wordsIds,proto3" json:"wordsIds,omitempty"`
}

func (x *ListWordsByIdsDTO) Reset() {
	*x = ListWordsByIdsDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_create_v1_words_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWordsByIdsDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWordsByIdsDTO) ProtoMessage() {}

func (x *ListWordsByIdsDTO) ProtoReflect() protoreflect.Message {
	mi := &file_api_create_v1_words_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWordsByIdsDTO.ProtoReflect.Descriptor instead.
func (*ListWordsByIdsDTO) Descriptor() ([]byte, []int) {
	return file_api_create_v1_words_proto_rawDescGZIP(), []int{11}
}

func (x *ListWordsByIdsDTO) GetWordsIds() []string {
	if x != nil {
		return x.WordsIds
	}
	return nil
}

var File_api_create_v1_words_proto protoreflect.FileDescriptor

var file_api_create_v1_words_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x12, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x36, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f,
	0x72, 0x64, 0x73, 0x50, 0x6f, 0x6c, 0x6c, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x23, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x50, 0x6f, 0x6c, 0x6c, 0x56,
	0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb0, 0x02, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x56, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x65,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x2e, 0x0a,
	0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x11, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xf2, 0x02,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x44, 0x54, 0x4f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x09, 0x69, 0x73, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x69,
	0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x6c, 0x6c, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6f, 0x6c, 0x6c, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x33, 0x0a,
	0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x12, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x69, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x02, 0x52, 0x06, 0x69, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x88, 0x01, 0x01, 0x12,
	0x2d, 0x0a, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61,
	0x6e, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x0f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x88, 0x01, 0x01, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x42, 0x15, 0x0a, 0x13,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x69, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x42, 0x12,
	0x0a, 0x10, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61,
	0x6e, 0x6b, 0x22, 0x91, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73,
	0x56, 0x4f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x36,
	0x0a, 0x09, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xb5, 0x02, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x64, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69,
	0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6c, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x6c, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x12, 0x2c, 0x0a, 0x11, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x22, 0x2f,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73,
	0x44, 0x54, 0x4f, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x49, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x49, 0x64, 0x73, 0x32,
	0xde, 0x03, 0x0a, 0x05, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x51, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x56, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x50, 0x6f, 0x6c, 0x6c, 0x12,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x50, 0x6f, 0x6c, 0x6c, 0x44,
	0x54, 0x4f, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x50, 0x6f,
	0x6c, 0x6c, 0x56, 0x4f, 0x12, 0x51, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x57, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x44, 0x54, 0x4f, 0x1a,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x56, 0x4f, 0x12, 0x44, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64,
	0x73, 0x44, 0x54, 0x4f, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x56, 0x4f,
	0x12, 0x4e, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x79, 0x49,
	0x64, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x79, 0x49, 0x64,
	0x73, 0x44, 0x54, 0x4f, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x56, 0x4f,
	0x42, 0x31, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x1e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_create_v1_words_proto_rawDescOnce sync.Once
	file_api_create_v1_words_proto_rawDescData = file_api_create_v1_words_proto_rawDesc
)

func file_api_create_v1_words_proto_rawDescGZIP() []byte {
	file_api_create_v1_words_proto_rawDescOnce.Do(func() {
		file_api_create_v1_words_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_create_v1_words_proto_rawDescData)
	})
	return file_api_create_v1_words_proto_rawDescData
}

var file_api_create_v1_words_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_create_v1_words_proto_goTypes = []interface{}{
	(*CreateWordsRequest)(nil), // 0: api.create.v1.CreateWordsRequest
	(*CreateWordsReply)(nil),   // 1: api.create.v1.CreateWordsReply
	(*UpdateWordsPollDTO)(nil), // 2: api.create.v1.UpdateWordsPollDTO
	(*UpdateWordsPollVO)(nil),  // 3: api.create.v1.UpdateWordsPollVO
	(*DeleteWordsRequest)(nil), // 4: api.create.v1.DeleteWordsRequest
	(*DeleteWordsReply)(nil),   // 5: api.create.v1.DeleteWordsReply
	(*GetWordsDTO)(nil),        // 6: api.create.v1.GetWordsDTO
	(*GetWordsVO)(nil),         // 7: api.create.v1.GetWordsVO
	(*ListWordsDTO)(nil),       // 8: api.create.v1.ListWordsDTO
	(*ListWordsVO)(nil),        // 9: api.create.v1.ListWordsVO
	(*WordsInfo)(nil),          // 10: api.create.v1.WordsInfo
	(*ListWordsByIdsDTO)(nil),  // 11: api.create.v1.ListWordsByIdsDTO
}
var file_api_create_v1_words_proto_depIdxs = []int32{
	10, // 0: api.create.v1.ListWordsVO.wordsInfo:type_name -> api.create.v1.WordsInfo
	0,  // 1: api.create.v1.Words.CreateWords:input_type -> api.create.v1.CreateWordsRequest
	2,  // 2: api.create.v1.Words.UpdateWordsPoll:input_type -> api.create.v1.UpdateWordsPollDTO
	4,  // 3: api.create.v1.Words.DeleteWords:input_type -> api.create.v1.DeleteWordsRequest
	6,  // 4: api.create.v1.Words.GetWords:input_type -> api.create.v1.GetWordsDTO
	8,  // 5: api.create.v1.Words.ListWords:input_type -> api.create.v1.ListWordsDTO
	11, // 6: api.create.v1.Words.ListWordsByIds:input_type -> api.create.v1.ListWordsByIdsDTO
	1,  // 7: api.create.v1.Words.CreateWords:output_type -> api.create.v1.CreateWordsReply
	3,  // 8: api.create.v1.Words.UpdateWordsPoll:output_type -> api.create.v1.UpdateWordsPollVO
	5,  // 9: api.create.v1.Words.DeleteWords:output_type -> api.create.v1.DeleteWordsReply
	7,  // 10: api.create.v1.Words.GetWords:output_type -> api.create.v1.GetWordsVO
	9,  // 11: api.create.v1.Words.ListWords:output_type -> api.create.v1.ListWordsVO
	9,  // 12: api.create.v1.Words.ListWordsByIds:output_type -> api.create.v1.ListWordsVO
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_create_v1_words_proto_init() }
func file_api_create_v1_words_proto_init() {
	if File_api_create_v1_words_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_create_v1_words_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWordsRequest); i {
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
		file_api_create_v1_words_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWordsReply); i {
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
		file_api_create_v1_words_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWordsPollDTO); i {
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
		file_api_create_v1_words_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWordsPollVO); i {
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
		file_api_create_v1_words_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWordsRequest); i {
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
		file_api_create_v1_words_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWordsReply); i {
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
		file_api_create_v1_words_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWordsDTO); i {
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
		file_api_create_v1_words_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWordsVO); i {
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
		file_api_create_v1_words_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWordsDTO); i {
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
		file_api_create_v1_words_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWordsVO); i {
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
		file_api_create_v1_words_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WordsInfo); i {
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
		file_api_create_v1_words_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWordsByIdsDTO); i {
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
	file_api_create_v1_words_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_create_v1_words_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_create_v1_words_proto_goTypes,
		DependencyIndexes: file_api_create_v1_words_proto_depIdxs,
		MessageInfos:      file_api_create_v1_words_proto_msgTypes,
	}.Build()
	File_api_create_v1_words_proto = out.File
	file_api_create_v1_words_proto_rawDesc = nil
	file_api_create_v1_words_proto_goTypes = nil
	file_api_create_v1_words_proto_depIdxs = nil
}
