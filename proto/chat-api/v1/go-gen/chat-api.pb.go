// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0--rc3
// source: chat-api.proto

package proto

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

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender   string `protobuf:"bytes,1,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Receiver string `protobuf:"bytes,2,opt,name=Receiver,proto3" json:"Receiver,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_chat_api_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *SendMessageRequest) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *SendMessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Done bool `protobuf:"varint,1,opt,name=Done,proto3" json:"Done,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_chat_api_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageResponse) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type ReceiveMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ReceiveMessagesRequest) Reset() {
	*x = ReceiveMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveMessagesRequest) ProtoMessage() {}

func (x *ReceiveMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveMessagesRequest.ProtoReflect.Descriptor instead.
func (*ReceiveMessagesRequest) Descriptor() ([]byte, []int) {
	return file_chat_api_proto_rawDescGZIP(), []int{2}
}

func (x *ReceiveMessagesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReceiveMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender    string `protobuf:"bytes,1,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *ReceiveMessagesResponse) Reset() {
	*x = ReceiveMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveMessagesResponse) ProtoMessage() {}

func (x *ReceiveMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveMessagesResponse.ProtoReflect.Descriptor instead.
func (*ReceiveMessagesResponse) Descriptor() ([]byte, []int) {
	return file_chat_api_proto_rawDescGZIP(), []int{3}
}

func (x *ReceiveMessagesResponse) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ReceiveMessagesResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ReceiveMessagesResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type RegisterDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	DeviceToken string `protobuf:"bytes,2,opt,name=DeviceToken,proto3" json:"DeviceToken,omitempty"`
	Platform    string `protobuf:"bytes,3,opt,name=Platform,proto3" json:"Platform,omitempty"`
}

func (x *RegisterDeviceRequest) Reset() {
	*x = RegisterDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDeviceRequest) ProtoMessage() {}

func (x *RegisterDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDeviceRequest.ProtoReflect.Descriptor instead.
func (*RegisterDeviceRequest) Descriptor() ([]byte, []int) {
	return file_chat_api_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterDeviceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterDeviceRequest) GetDeviceToken() string {
	if x != nil {
		return x.DeviceToken
	}
	return ""
}

func (x *RegisterDeviceRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

type RegisterDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Done bool `protobuf:"varint,1,opt,name=Done,proto3" json:"Done,omitempty"`
}

func (x *RegisterDeviceResponse) Reset() {
	*x = RegisterDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDeviceResponse) ProtoMessage() {}

func (x *RegisterDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDeviceResponse.ProtoReflect.Descriptor instead.
func (*RegisterDeviceResponse) Descriptor() ([]byte, []int) {
	return file_chat_api_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterDeviceResponse) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type GetChatHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DialogueOwner string `protobuf:"bytes,1,opt,name=DialogueOwner,proto3" json:"DialogueOwner,omitempty"`
	Partner       string `protobuf:"bytes,2,opt,name=Partner,proto3" json:"Partner,omitempty"`
}

func (x *GetChatHistoryRequest) Reset() {
	*x = GetChatHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatHistoryRequest) ProtoMessage() {}

func (x *GetChatHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetChatHistoryRequest) Descriptor() ([]byte, []int) {
	return file_chat_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetChatHistoryRequest) GetDialogueOwner() string {
	if x != nil {
		return x.DialogueOwner
	}
	return ""
}

func (x *GetChatHistoryRequest) GetPartner() string {
	if x != nil {
		return x.Partner
	}
	return ""
}

type GetChatHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*ChatMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *GetChatHistoryResponse) Reset() {
	*x = GetChatHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatHistoryResponse) ProtoMessage() {}

func (x *GetChatHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetChatHistoryResponse) Descriptor() ([]byte, []int) {
	return file_chat_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetChatHistoryResponse) GetMessages() []*ChatMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender    string `protobuf:"bytes,1,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chat_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_chat_api_proto_rawDescGZIP(), []int{8}
}

func (x *ChatMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ChatMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatMessage) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_chat_api_proto protoreflect.FileDescriptor

var file_chat_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x76, 0x31, 0x22, 0x62,
	0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x29, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x22, 0x28, 0x0a,
	0x16, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x17, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x65, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x2c, 0x0a, 0x16, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x22, 0x57, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75,
	0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x22, 0x4f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x22, 0x5d, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x32, 0xfd, 0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x52, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x20, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x5b, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x73, 0x70, 0x33, 0x63, 0x74, 0x61, 0x68, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_api_proto_rawDescOnce sync.Once
	file_chat_api_proto_rawDescData = file_chat_api_proto_rawDesc
)

func file_chat_api_proto_rawDescGZIP() []byte {
	file_chat_api_proto_rawDescOnce.Do(func() {
		file_chat_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_api_proto_rawDescData)
	})
	return file_chat_api_proto_rawDescData
}

var file_chat_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_chat_api_proto_goTypes = []interface{}{
	(*SendMessageRequest)(nil),      // 0: chat_serv.v1.SendMessageRequest
	(*SendMessageResponse)(nil),     // 1: chat_serv.v1.SendMessageResponse
	(*ReceiveMessagesRequest)(nil),  // 2: chat_serv.v1.ReceiveMessagesRequest
	(*ReceiveMessagesResponse)(nil), // 3: chat_serv.v1.ReceiveMessagesResponse
	(*RegisterDeviceRequest)(nil),   // 4: chat_serv.v1.RegisterDeviceRequest
	(*RegisterDeviceResponse)(nil),  // 5: chat_serv.v1.RegisterDeviceResponse
	(*GetChatHistoryRequest)(nil),   // 6: chat_serv.v1.GetChatHistoryRequest
	(*GetChatHistoryResponse)(nil),  // 7: chat_serv.v1.GetChatHistoryResponse
	(*ChatMessage)(nil),             // 8: chat_serv.v1.ChatMessage
}
var file_chat_api_proto_depIdxs = []int32{
	8, // 0: chat_serv.v1.GetChatHistoryResponse.messages:type_name -> chat_serv.v1.ChatMessage
	0, // 1: chat_serv.v1.ChatService.SendMessage:input_type -> chat_serv.v1.SendMessageRequest
	2, // 2: chat_serv.v1.ChatService.ReceiveMessages:input_type -> chat_serv.v1.ReceiveMessagesRequest
	4, // 3: chat_serv.v1.ChatService.RegisterDevice:input_type -> chat_serv.v1.RegisterDeviceRequest
	6, // 4: chat_serv.v1.ChatService.GetChatHistory:input_type -> chat_serv.v1.GetChatHistoryRequest
	1, // 5: chat_serv.v1.ChatService.SendMessage:output_type -> chat_serv.v1.SendMessageResponse
	3, // 6: chat_serv.v1.ChatService.ReceiveMessages:output_type -> chat_serv.v1.ReceiveMessagesResponse
	5, // 7: chat_serv.v1.ChatService.RegisterDevice:output_type -> chat_serv.v1.RegisterDeviceResponse
	7, // 8: chat_serv.v1.ChatService.GetChatHistory:output_type -> chat_serv.v1.GetChatHistoryResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_api_proto_init() }
func file_chat_api_proto_init() {
	if File_chat_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_chat_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_chat_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveMessagesRequest); i {
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
		file_chat_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveMessagesResponse); i {
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
		file_chat_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDeviceRequest); i {
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
		file_chat_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDeviceResponse); i {
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
		file_chat_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatHistoryRequest); i {
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
		file_chat_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatHistoryResponse); i {
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
		file_chat_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
			RawDescriptor: file_chat_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_api_proto_goTypes,
		DependencyIndexes: file_chat_api_proto_depIdxs,
		MessageInfos:      file_chat_api_proto_msgTypes,
	}.Build()
	File_chat_api_proto = out.File
	file_chat_api_proto_rawDesc = nil
	file_chat_api_proto_goTypes = nil
	file_chat_api_proto_depIdxs = nil
}
