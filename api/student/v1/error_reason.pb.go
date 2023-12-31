// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: student/v1/error_reason.proto

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

type CreateErrorReasonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateErrorReasonRequest) Reset() {
	*x = CreateErrorReasonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_error_reason_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateErrorReasonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateErrorReasonRequest) ProtoMessage() {}

func (x *CreateErrorReasonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_error_reason_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateErrorReasonRequest.ProtoReflect.Descriptor instead.
func (*CreateErrorReasonRequest) Descriptor() ([]byte, []int) {
	return file_student_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

type CreateErrorReasonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateErrorReasonReply) Reset() {
	*x = CreateErrorReasonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_error_reason_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateErrorReasonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateErrorReasonReply) ProtoMessage() {}

func (x *CreateErrorReasonReply) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_error_reason_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateErrorReasonReply.ProtoReflect.Descriptor instead.
func (*CreateErrorReasonReply) Descriptor() ([]byte, []int) {
	return file_student_v1_error_reason_proto_rawDescGZIP(), []int{1}
}

type UpdateErrorReasonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateErrorReasonRequest) Reset() {
	*x = UpdateErrorReasonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_error_reason_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateErrorReasonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateErrorReasonRequest) ProtoMessage() {}

func (x *UpdateErrorReasonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_error_reason_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateErrorReasonRequest.ProtoReflect.Descriptor instead.
func (*UpdateErrorReasonRequest) Descriptor() ([]byte, []int) {
	return file_student_v1_error_reason_proto_rawDescGZIP(), []int{2}
}

type UpdateErrorReasonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateErrorReasonReply) Reset() {
	*x = UpdateErrorReasonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_error_reason_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateErrorReasonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateErrorReasonReply) ProtoMessage() {}

func (x *UpdateErrorReasonReply) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_error_reason_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateErrorReasonReply.ProtoReflect.Descriptor instead.
func (*UpdateErrorReasonReply) Descriptor() ([]byte, []int) {
	return file_student_v1_error_reason_proto_rawDescGZIP(), []int{3}
}

type DeleteErrorReasonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteErrorReasonRequest) Reset() {
	*x = DeleteErrorReasonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_error_reason_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteErrorReasonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteErrorReasonRequest) ProtoMessage() {}

func (x *DeleteErrorReasonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_error_reason_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteErrorReasonRequest.ProtoReflect.Descriptor instead.
func (*DeleteErrorReasonRequest) Descriptor() ([]byte, []int) {
	return file_student_v1_error_reason_proto_rawDescGZIP(), []int{4}
}

type DeleteErrorReasonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteErrorReasonReply) Reset() {
	*x = DeleteErrorReasonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_error_reason_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteErrorReasonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteErrorReasonReply) ProtoMessage() {}

func (x *DeleteErrorReasonReply) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_error_reason_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteErrorReasonReply.ProtoReflect.Descriptor instead.
func (*DeleteErrorReasonReply) Descriptor() ([]byte, []int) {
	return file_student_v1_error_reason_proto_rawDescGZIP(), []int{5}
}

type GetErrorReasonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetErrorReasonRequest) Reset() {
	*x = GetErrorReasonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_error_reason_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetErrorReasonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetErrorReasonRequest) ProtoMessage() {}

func (x *GetErrorReasonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_error_reason_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetErrorReasonRequest.ProtoReflect.Descriptor instead.
func (*GetErrorReasonRequest) Descriptor() ([]byte, []int) {
	return file_student_v1_error_reason_proto_rawDescGZIP(), []int{6}
}

type GetErrorReasonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetErrorReasonReply) Reset() {
	*x = GetErrorReasonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_error_reason_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetErrorReasonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetErrorReasonReply) ProtoMessage() {}

func (x *GetErrorReasonReply) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_error_reason_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetErrorReasonReply.ProtoReflect.Descriptor instead.
func (*GetErrorReasonReply) Descriptor() ([]byte, []int) {
	return file_student_v1_error_reason_proto_rawDescGZIP(), []int{7}
}

type ListErrorReasonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListErrorReasonRequest) Reset() {
	*x = ListErrorReasonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_error_reason_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListErrorReasonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListErrorReasonRequest) ProtoMessage() {}

func (x *ListErrorReasonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_error_reason_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListErrorReasonRequest.ProtoReflect.Descriptor instead.
func (*ListErrorReasonRequest) Descriptor() ([]byte, []int) {
	return file_student_v1_error_reason_proto_rawDescGZIP(), []int{8}
}

type ListErrorReasonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListErrorReasonReply) Reset() {
	*x = ListErrorReasonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_error_reason_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListErrorReasonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListErrorReasonReply) ProtoMessage() {}

func (x *ListErrorReasonReply) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_error_reason_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListErrorReasonReply.ProtoReflect.Descriptor instead.
func (*ListErrorReasonReply) Descriptor() ([]byte, []int) {
	return file_student_v1_error_reason_proto_rawDescGZIP(), []int{9}
}

var File_student_v1_error_reason_proto protoreflect.FileDescriptor

var file_student_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22,
	0x1a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x0a, 0x18, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x18, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0x81, 0x04, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x65, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x65, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x65, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5f, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x31, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1d, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_student_v1_error_reason_proto_rawDescOnce sync.Once
	file_student_v1_error_reason_proto_rawDescData = file_student_v1_error_reason_proto_rawDesc
)

func file_student_v1_error_reason_proto_rawDescGZIP() []byte {
	file_student_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_student_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_student_v1_error_reason_proto_rawDescData)
	})
	return file_student_v1_error_reason_proto_rawDescData
}

var file_student_v1_error_reason_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_student_v1_error_reason_proto_goTypes = []interface{}{
	(*CreateErrorReasonRequest)(nil), // 0: api.student.v1.CreateErrorReasonRequest
	(*CreateErrorReasonReply)(nil),   // 1: api.student.v1.CreateErrorReasonReply
	(*UpdateErrorReasonRequest)(nil), // 2: api.student.v1.UpdateErrorReasonRequest
	(*UpdateErrorReasonReply)(nil),   // 3: api.student.v1.UpdateErrorReasonReply
	(*DeleteErrorReasonRequest)(nil), // 4: api.student.v1.DeleteErrorReasonRequest
	(*DeleteErrorReasonReply)(nil),   // 5: api.student.v1.DeleteErrorReasonReply
	(*GetErrorReasonRequest)(nil),    // 6: api.student.v1.GetErrorReasonRequest
	(*GetErrorReasonReply)(nil),      // 7: api.student.v1.GetErrorReasonReply
	(*ListErrorReasonRequest)(nil),   // 8: api.student.v1.ListErrorReasonRequest
	(*ListErrorReasonReply)(nil),     // 9: api.student.v1.ListErrorReasonReply
}
var file_student_v1_error_reason_proto_depIdxs = []int32{
	0, // 0: api.student.v1.ErrorReason.CreateErrorReason:input_type -> api.student.v1.CreateErrorReasonRequest
	2, // 1: api.student.v1.ErrorReason.UpdateErrorReason:input_type -> api.student.v1.UpdateErrorReasonRequest
	4, // 2: api.student.v1.ErrorReason.DeleteErrorReason:input_type -> api.student.v1.DeleteErrorReasonRequest
	6, // 3: api.student.v1.ErrorReason.GetErrorReason:input_type -> api.student.v1.GetErrorReasonRequest
	8, // 4: api.student.v1.ErrorReason.ListErrorReason:input_type -> api.student.v1.ListErrorReasonRequest
	1, // 5: api.student.v1.ErrorReason.CreateErrorReason:output_type -> api.student.v1.CreateErrorReasonReply
	3, // 6: api.student.v1.ErrorReason.UpdateErrorReason:output_type -> api.student.v1.UpdateErrorReasonReply
	5, // 7: api.student.v1.ErrorReason.DeleteErrorReason:output_type -> api.student.v1.DeleteErrorReasonReply
	7, // 8: api.student.v1.ErrorReason.GetErrorReason:output_type -> api.student.v1.GetErrorReasonReply
	9, // 9: api.student.v1.ErrorReason.ListErrorReason:output_type -> api.student.v1.ListErrorReasonReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_student_v1_error_reason_proto_init() }
func file_student_v1_error_reason_proto_init() {
	if File_student_v1_error_reason_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_student_v1_error_reason_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateErrorReasonRequest); i {
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
		file_student_v1_error_reason_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateErrorReasonReply); i {
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
		file_student_v1_error_reason_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateErrorReasonRequest); i {
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
		file_student_v1_error_reason_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateErrorReasonReply); i {
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
		file_student_v1_error_reason_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteErrorReasonRequest); i {
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
		file_student_v1_error_reason_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteErrorReasonReply); i {
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
		file_student_v1_error_reason_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetErrorReasonRequest); i {
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
		file_student_v1_error_reason_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetErrorReasonReply); i {
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
		file_student_v1_error_reason_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListErrorReasonRequest); i {
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
		file_student_v1_error_reason_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListErrorReasonReply); i {
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
			RawDescriptor: file_student_v1_error_reason_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_student_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_student_v1_error_reason_proto_depIdxs,
		MessageInfos:      file_student_v1_error_reason_proto_msgTypes,
	}.Build()
	File_student_v1_error_reason_proto = out.File
	file_student_v1_error_reason_proto_rawDesc = nil
	file_student_v1_error_reason_proto_goTypes = nil
	file_student_v1_error_reason_proto_depIdxs = nil
}
