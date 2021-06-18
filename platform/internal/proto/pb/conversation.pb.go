// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: conversation.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// ConversationUpdateType Enum
type ConversationUpdateType int32

const (
	ConversationUpdateType_CONVERSATION_NO_UPDATE_TYPE_SET ConversationUpdateType = 0
	ConversationUpdateType_CONVERSATION_CREATED            ConversationUpdateType = 1
	ConversationUpdateType_CONVERSATION_UPDATED            ConversationUpdateType = 2
	ConversationUpdateType_CONVERSATION_DELETED            ConversationUpdateType = 3
)

// Enum value maps for ConversationUpdateType.
var (
	ConversationUpdateType_name = map[int32]string{
		0: "CONVERSATION_NO_UPDATE_TYPE_SET",
		1: "CONVERSATION_CREATED",
		2: "CONVERSATION_UPDATED",
		3: "CONVERSATION_DELETED",
	}
	ConversationUpdateType_value = map[string]int32{
		"CONVERSATION_NO_UPDATE_TYPE_SET": 0,
		"CONVERSATION_CREATED":            1,
		"CONVERSATION_UPDATED":            2,
		"CONVERSATION_DELETED":            3,
	}
)

func (x ConversationUpdateType) Enum() *ConversationUpdateType {
	p := new(ConversationUpdateType)
	*p = x
	return p
}

func (x ConversationUpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversationUpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_conversation_proto_enumTypes[0].Descriptor()
}

func (ConversationUpdateType) Type() protoreflect.EnumType {
	return &file_conversation_proto_enumTypes[0]
}

func (x ConversationUpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversationUpdateType.Descriptor instead.
func (ConversationUpdateType) EnumDescriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{0}
}

// MessageUpdateType Enum
type MessageUpdateType int32

const (
	MessageUpdateType_MESSAGE_NO_UPDATE_TYPE_SET MessageUpdateType = 0
	MessageUpdateType_MESSAGE_CREATED            MessageUpdateType = 1
)

// Enum value maps for MessageUpdateType.
var (
	MessageUpdateType_name = map[int32]string{
		0: "MESSAGE_NO_UPDATE_TYPE_SET",
		1: "MESSAGE_CREATED",
	}
	MessageUpdateType_value = map[string]int32{
		"MESSAGE_NO_UPDATE_TYPE_SET": 0,
		"MESSAGE_CREATED":            1,
	}
)

func (x MessageUpdateType) Enum() *MessageUpdateType {
	p := new(MessageUpdateType)
	*p = x
	return p
}

func (x MessageUpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageUpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_conversation_proto_enumTypes[1].Descriptor()
}

func (MessageUpdateType) Type() protoreflect.EnumType {
	return &file_conversation_proto_enumTypes[1]
}

func (x MessageUpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageUpdateType.Descriptor instead.
func (MessageUpdateType) EnumDescriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{1}
}

// Message Protobuf Type
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ConversationId string                 `protobuf:"bytes,3,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	AuthorId       string                 `protobuf:"bytes,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Message        string                 `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Message) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

func (x *Message) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// ConversationHasParticipant Protobuf Type
type ConversationHasParticipant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversationId string `protobuf:"bytes,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	ParticipantId  string `protobuf:"bytes,2,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
}

func (x *ConversationHasParticipant) Reset() {
	*x = ConversationHasParticipant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationHasParticipant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationHasParticipant) ProtoMessage() {}

func (x *ConversationHasParticipant) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationHasParticipant.ProtoReflect.Descriptor instead.
func (*ConversationHasParticipant) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{1}
}

func (x *ConversationHasParticipant) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

func (x *ConversationHasParticipant) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

// Conversation Protobuf Type
type Conversation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label         string                        `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	IsPublic      bool                          `protobuf:"varint,3,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	LastMessageOn *timestamppb.Timestamp        `protobuf:"bytes,4,opt,name=last_message_on,json=lastMessageOn,proto3" json:"last_message_on,omitempty"`
	Participants  []*ConversationHasParticipant `protobuf:"bytes,5,rep,name=participants,proto3" json:"participants,omitempty"`
	Messages      []*Message                    `protobuf:"bytes,6,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Conversation) Reset() {
	*x = Conversation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conversation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conversation) ProtoMessage() {}

func (x *Conversation) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conversation.ProtoReflect.Descriptor instead.
func (*Conversation) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{2}
}

func (x *Conversation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Conversation) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Conversation) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *Conversation) GetLastMessageOn() *timestamppb.Timestamp {
	if x != nil {
		return x.LastMessageOn
	}
	return nil
}

func (x *Conversation) GetParticipants() []*ConversationHasParticipant {
	if x != nil {
		return x.Participants
	}
	return nil
}

func (x *Conversation) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

// ConversationList Protobuf Type
type ConversationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conversations []*Conversation `protobuf:"bytes,1,rep,name=conversations,proto3" json:"conversations,omitempty"`
}

func (x *ConversationList) Reset() {
	*x = ConversationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationList) ProtoMessage() {}

func (x *ConversationList) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationList.ProtoReflect.Descriptor instead.
func (*ConversationList) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{3}
}

func (x *ConversationList) GetConversations() []*Conversation {
	if x != nil {
		return x.Conversations
	}
	return nil
}

// ConversationUpdate Protobuf Type
type ConversationUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateType   ConversationUpdateType `protobuf:"varint,1,opt,name=update_type,json=updateType,proto3,enum=pb.ConversationUpdateType" json:"update_type,omitempty"`
	Conversation *Conversation          `protobuf:"bytes,2,opt,name=conversation,proto3" json:"conversation,omitempty"`
}

func (x *ConversationUpdate) Reset() {
	*x = ConversationUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationUpdate) ProtoMessage() {}

func (x *ConversationUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationUpdate.ProtoReflect.Descriptor instead.
func (*ConversationUpdate) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{4}
}

func (x *ConversationUpdate) GetUpdateType() ConversationUpdateType {
	if x != nil {
		return x.UpdateType
	}
	return ConversationUpdateType_CONVERSATION_NO_UPDATE_TYPE_SET
}

func (x *ConversationUpdate) GetConversation() *Conversation {
	if x != nil {
		return x.Conversation
	}
	return nil
}

// MessageUpdate Protobuf Type
type MessageUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateType MessageUpdateType `protobuf:"varint,1,opt,name=update_type,json=updateType,proto3,enum=pb.MessageUpdateType" json:"update_type,omitempty"`
	Message    *Message          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageUpdate) Reset() {
	*x = MessageUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageUpdate) ProtoMessage() {}

func (x *MessageUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageUpdate.ProtoReflect.Descriptor instead.
func (*MessageUpdate) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{5}
}

func (x *MessageUpdate) GetUpdateType() MessageUpdateType {
	if x != nil {
		return x.UpdateType
	}
	return MessageUpdateType_MESSAGE_NO_UPDATE_TYPE_SET
}

func (x *MessageUpdate) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

// ChatStreamUpdate Protobuf Type
type ChatStreamUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UpdateType:
	//	*ChatStreamUpdate_ConversationInfo
	//	*ChatStreamUpdate_MessageInfo
	UpdateType isChatStreamUpdate_UpdateType `protobuf_oneof:"update_type"`
}

func (x *ChatStreamUpdate) Reset() {
	*x = ChatStreamUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatStreamUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatStreamUpdate) ProtoMessage() {}

func (x *ChatStreamUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatStreamUpdate.ProtoReflect.Descriptor instead.
func (*ChatStreamUpdate) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{6}
}

func (m *ChatStreamUpdate) GetUpdateType() isChatStreamUpdate_UpdateType {
	if m != nil {
		return m.UpdateType
	}
	return nil
}

func (x *ChatStreamUpdate) GetConversationInfo() *ConversationUpdate {
	if x, ok := x.GetUpdateType().(*ChatStreamUpdate_ConversationInfo); ok {
		return x.ConversationInfo
	}
	return nil
}

func (x *ChatStreamUpdate) GetMessageInfo() *MessageUpdate {
	if x, ok := x.GetUpdateType().(*ChatStreamUpdate_MessageInfo); ok {
		return x.MessageInfo
	}
	return nil
}

type isChatStreamUpdate_UpdateType interface {
	isChatStreamUpdate_UpdateType()
}

type ChatStreamUpdate_ConversationInfo struct {
	ConversationInfo *ConversationUpdate `protobuf:"bytes,1,opt,name=conversation_info,json=conversationInfo,proto3,oneof"`
}

type ChatStreamUpdate_MessageInfo struct {
	MessageInfo *MessageUpdate `protobuf:"bytes,2,opt,name=message_info,json=messageInfo,proto3,oneof"`
}

func (*ChatStreamUpdate_ConversationInfo) isChatStreamUpdate_UpdateType() {}

func (*ChatStreamUpdate_MessageInfo) isChatStreamUpdate_UpdateType() {}

var File_conversation_proto protoreflect.FileDescriptor

var file_conversation_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6c, 0x0a,
	0x1a, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x82, 0x02, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12,
	0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x6e, 0x12, 0x42, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x4a, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x87, 0x01, 0x0a,
	0x12, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x34, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6e, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x25, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x45, 0x0a, 0x11, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00,
	0x52, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0d, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x8b, 0x01, 0x0a, 0x16, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a,
	0x14, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x48, 0x0a, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a,
	0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f,
	0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x32, 0xa1, 0x03, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conversation_proto_rawDescOnce sync.Once
	file_conversation_proto_rawDescData = file_conversation_proto_rawDesc
)

func file_conversation_proto_rawDescGZIP() []byte {
	file_conversation_proto_rawDescOnce.Do(func() {
		file_conversation_proto_rawDescData = protoimpl.X.CompressGZIP(file_conversation_proto_rawDescData)
	})
	return file_conversation_proto_rawDescData
}

var file_conversation_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_conversation_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_conversation_proto_goTypes = []interface{}{
	(ConversationUpdateType)(0),        // 0: pb.ConversationUpdateType
	(MessageUpdateType)(0),             // 1: pb.MessageUpdateType
	(*Message)(nil),                    // 2: pb.Message
	(*ConversationHasParticipant)(nil), // 3: pb.ConversationHasParticipant
	(*Conversation)(nil),               // 4: pb.Conversation
	(*ConversationList)(nil),           // 5: pb.ConversationList
	(*ConversationUpdate)(nil),         // 6: pb.ConversationUpdate
	(*MessageUpdate)(nil),              // 7: pb.MessageUpdate
	(*ChatStreamUpdate)(nil),           // 8: pb.ChatStreamUpdate
	(*timestamppb.Timestamp)(nil),      // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),              // 10: google.protobuf.Empty
	(*wrapperspb.StringValue)(nil),     // 11: google.protobuf.StringValue
}
var file_conversation_proto_depIdxs = []int32{
	9,  // 0: pb.Message.created_at:type_name -> google.protobuf.Timestamp
	9,  // 1: pb.Conversation.last_message_on:type_name -> google.protobuf.Timestamp
	3,  // 2: pb.Conversation.participants:type_name -> pb.ConversationHasParticipant
	2,  // 3: pb.Conversation.messages:type_name -> pb.Message
	4,  // 4: pb.ConversationList.conversations:type_name -> pb.Conversation
	0,  // 5: pb.ConversationUpdate.update_type:type_name -> pb.ConversationUpdateType
	4,  // 6: pb.ConversationUpdate.conversation:type_name -> pb.Conversation
	1,  // 7: pb.MessageUpdate.update_type:type_name -> pb.MessageUpdateType
	2,  // 8: pb.MessageUpdate.message:type_name -> pb.Message
	6,  // 9: pb.ChatStreamUpdate.conversation_info:type_name -> pb.ConversationUpdate
	7,  // 10: pb.ChatStreamUpdate.message_info:type_name -> pb.MessageUpdate
	4,  // 11: pb.ConversationService.CreateConversation:input_type -> pb.Conversation
	3,  // 12: pb.ConversationService.AddUserToConversation:input_type -> pb.ConversationHasParticipant
	10, // 13: pb.ConversationService.GetPublicConversations:input_type -> google.protobuf.Empty
	11, // 14: pb.ConversationService.GetPrivateConversationsForUser:input_type -> google.protobuf.StringValue
	2,  // 15: pb.ConversationService.CreateMessage:input_type -> pb.Message
	10, // 16: pb.ConversationService.ChatStream:input_type -> google.protobuf.Empty
	4,  // 17: pb.ConversationService.CreateConversation:output_type -> pb.Conversation
	4,  // 18: pb.ConversationService.AddUserToConversation:output_type -> pb.Conversation
	5,  // 19: pb.ConversationService.GetPublicConversations:output_type -> pb.ConversationList
	5,  // 20: pb.ConversationService.GetPrivateConversationsForUser:output_type -> pb.ConversationList
	2,  // 21: pb.ConversationService.CreateMessage:output_type -> pb.Message
	8,  // 22: pb.ConversationService.ChatStream:output_type -> pb.ChatStreamUpdate
	17, // [17:23] is the sub-list for method output_type
	11, // [11:17] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_conversation_proto_init() }
func file_conversation_proto_init() {
	if File_conversation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conversation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_conversation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationHasParticipant); i {
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
		file_conversation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conversation); i {
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
		file_conversation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationList); i {
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
		file_conversation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationUpdate); i {
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
		file_conversation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageUpdate); i {
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
		file_conversation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatStreamUpdate); i {
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
	file_conversation_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ChatStreamUpdate_ConversationInfo)(nil),
		(*ChatStreamUpdate_MessageInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conversation_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_conversation_proto_goTypes,
		DependencyIndexes: file_conversation_proto_depIdxs,
		EnumInfos:         file_conversation_proto_enumTypes,
		MessageInfos:      file_conversation_proto_msgTypes,
	}.Build()
	File_conversation_proto = out.File
	file_conversation_proto_rawDesc = nil
	file_conversation_proto_goTypes = nil
	file_conversation_proto_depIdxs = nil
}
