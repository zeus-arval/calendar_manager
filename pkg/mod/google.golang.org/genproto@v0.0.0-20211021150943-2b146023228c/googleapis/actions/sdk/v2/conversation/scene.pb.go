// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/actions/sdk/v2/conversation/scene.proto

package conversation

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents the current status of slot filling.
type SlotFillingStatus int32

const (
	// Fallback value when the usage field is not populated.
	SlotFillingStatus_UNSPECIFIED SlotFillingStatus = 0
	// The slots have been initialized but slot filling has not started.
	SlotFillingStatus_INITIALIZED SlotFillingStatus = 1
	// The slot values are being collected.
	SlotFillingStatus_COLLECTING SlotFillingStatus = 2
	// All slot values are final and cannot be changed.
	SlotFillingStatus_FINAL SlotFillingStatus = 4
)

// Enum value maps for SlotFillingStatus.
var (
	SlotFillingStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "INITIALIZED",
		2: "COLLECTING",
		4: "FINAL",
	}
	SlotFillingStatus_value = map[string]int32{
		"UNSPECIFIED": 0,
		"INITIALIZED": 1,
		"COLLECTING":  2,
		"FINAL":       4,
	}
)

func (x SlotFillingStatus) Enum() *SlotFillingStatus {
	p := new(SlotFillingStatus)
	*p = x
	return p
}

func (x SlotFillingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SlotFillingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_actions_sdk_v2_conversation_scene_proto_enumTypes[0].Descriptor()
}

func (SlotFillingStatus) Type() protoreflect.EnumType {
	return &file_google_actions_sdk_v2_conversation_scene_proto_enumTypes[0]
}

func (x SlotFillingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SlotFillingStatus.Descriptor instead.
func (SlotFillingStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_scene_proto_rawDescGZIP(), []int{0}
}

// Represents the mode of a slot, that is, if it is required or not.
type Slot_SlotMode int32

const (
	// Fallback value when the usage field is not populated.
	Slot_MODE_UNSPECIFIED Slot_SlotMode = 0
	// Indicates that the slot is not required to complete slot filling.
	Slot_OPTIONAL Slot_SlotMode = 1
	// Indicates that the slot is required to complete slot filling.
	Slot_REQUIRED Slot_SlotMode = 2
)

// Enum value maps for Slot_SlotMode.
var (
	Slot_SlotMode_name = map[int32]string{
		0: "MODE_UNSPECIFIED",
		1: "OPTIONAL",
		2: "REQUIRED",
	}
	Slot_SlotMode_value = map[string]int32{
		"MODE_UNSPECIFIED": 0,
		"OPTIONAL":         1,
		"REQUIRED":         2,
	}
)

func (x Slot_SlotMode) Enum() *Slot_SlotMode {
	p := new(Slot_SlotMode)
	*p = x
	return p
}

func (x Slot_SlotMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Slot_SlotMode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_actions_sdk_v2_conversation_scene_proto_enumTypes[1].Descriptor()
}

func (Slot_SlotMode) Type() protoreflect.EnumType {
	return &file_google_actions_sdk_v2_conversation_scene_proto_enumTypes[1]
}

func (x Slot_SlotMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Slot_SlotMode.Descriptor instead.
func (Slot_SlotMode) EnumDescriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_scene_proto_rawDescGZIP(), []int{0, 0}
}

// Represents the status of a slot.
type Slot_SlotStatus int32

const (
	// Fallback value when the usage field is not populated.
	Slot_SLOT_UNSPECIFIED Slot_SlotStatus = 0
	// Indicates that the slot does not have any values. This status cannot be
	// modified through the response.
	Slot_EMPTY Slot_SlotStatus = 1
	// Indicates that the slot value is invalid. This status can be set
	// through the response.
	Slot_INVALID Slot_SlotStatus = 2
	// Indicates that the slot has a value. This status cannot be modified
	// through the response.
	Slot_FILLED Slot_SlotStatus = 3
)

// Enum value maps for Slot_SlotStatus.
var (
	Slot_SlotStatus_name = map[int32]string{
		0: "SLOT_UNSPECIFIED",
		1: "EMPTY",
		2: "INVALID",
		3: "FILLED",
	}
	Slot_SlotStatus_value = map[string]int32{
		"SLOT_UNSPECIFIED": 0,
		"EMPTY":            1,
		"INVALID":          2,
		"FILLED":           3,
	}
)

func (x Slot_SlotStatus) Enum() *Slot_SlotStatus {
	p := new(Slot_SlotStatus)
	*p = x
	return p
}

func (x Slot_SlotStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Slot_SlotStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_actions_sdk_v2_conversation_scene_proto_enumTypes[2].Descriptor()
}

func (Slot_SlotStatus) Type() protoreflect.EnumType {
	return &file_google_actions_sdk_v2_conversation_scene_proto_enumTypes[2]
}

func (x Slot_SlotStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Slot_SlotStatus.Descriptor instead.
func (Slot_SlotStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_scene_proto_rawDescGZIP(), []int{0, 1}
}

// Represents a slot.
type Slot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mode of the slot (required or optional). Can be set by developer.
	Mode Slot_SlotMode `protobuf:"varint,1,opt,name=mode,proto3,enum=google.actions.sdk.v2.conversation.Slot_SlotMode" json:"mode,omitempty"`
	// The status of the slot.
	Status Slot_SlotStatus `protobuf:"varint,2,opt,name=status,proto3,enum=google.actions.sdk.v2.conversation.Slot_SlotStatus" json:"status,omitempty"`
	// The value of the slot. Changing this value in the response, will
	// modify the value in slot filling.
	Value *structpb.Value `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// Indicates if the slot value was collected on the last turn.
	// This field is read-only.
	Updated bool `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	// Optional. This prompt is sent to the user when needed to fill a required
	// slot. This prompt overrides the existing prompt defined in the console.
	// This field is not included in the webhook request.
	Prompt *Prompt `protobuf:"bytes,5,opt,name=prompt,proto3" json:"prompt,omitempty"`
}

func (x *Slot) Reset() {
	*x = Slot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_conversation_scene_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slot) ProtoMessage() {}

func (x *Slot) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_conversation_scene_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slot.ProtoReflect.Descriptor instead.
func (*Slot) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_scene_proto_rawDescGZIP(), []int{0}
}

func (x *Slot) GetMode() Slot_SlotMode {
	if x != nil {
		return x.Mode
	}
	return Slot_MODE_UNSPECIFIED
}

func (x *Slot) GetStatus() Slot_SlotStatus {
	if x != nil {
		return x.Status
	}
	return Slot_SLOT_UNSPECIFIED
}

func (x *Slot) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Slot) GetUpdated() bool {
	if x != nil {
		return x.Updated
	}
	return false
}

func (x *Slot) GetPrompt() *Prompt {
	if x != nil {
		return x.Prompt
	}
	return nil
}

var File_google_actions_sdk_v2_conversation_scene_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_conversation_scene_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x03, 0x0a, 0x04, 0x53,
	0x6c, 0x6f, 0x74, 0x12, 0x45, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x2e, 0x53, 0x6c, 0x6f, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x6c, 0x6f, 0x74, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x42, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x06, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x22, 0x3c, 0x0a, 0x08, 0x53, 0x6c, 0x6f, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41,
	0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x02, 0x22, 0x46, 0x0a, 0x0a, 0x53, 0x6c, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x10, 0x53, 0x4c, 0x4f, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x50, 0x0a, 0x11, 0x53, 0x6c, 0x6f,
	0x74, 0x46, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x10, 0x04, 0x42, 0x86, 0x01, 0x0a, 0x26,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_conversation_scene_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_conversation_scene_proto_rawDescData = file_google_actions_sdk_v2_conversation_scene_proto_rawDesc
)

func file_google_actions_sdk_v2_conversation_scene_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_conversation_scene_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_conversation_scene_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_conversation_scene_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_conversation_scene_proto_rawDescData
}

var file_google_actions_sdk_v2_conversation_scene_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_actions_sdk_v2_conversation_scene_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_actions_sdk_v2_conversation_scene_proto_goTypes = []interface{}{
	(SlotFillingStatus)(0), // 0: google.actions.sdk.v2.conversation.SlotFillingStatus
	(Slot_SlotMode)(0),     // 1: google.actions.sdk.v2.conversation.Slot.SlotMode
	(Slot_SlotStatus)(0),   // 2: google.actions.sdk.v2.conversation.Slot.SlotStatus
	(*Slot)(nil),           // 3: google.actions.sdk.v2.conversation.Slot
	(*structpb.Value)(nil), // 4: google.protobuf.Value
	(*Prompt)(nil),         // 5: google.actions.sdk.v2.conversation.Prompt
}
var file_google_actions_sdk_v2_conversation_scene_proto_depIdxs = []int32{
	1, // 0: google.actions.sdk.v2.conversation.Slot.mode:type_name -> google.actions.sdk.v2.conversation.Slot.SlotMode
	2, // 1: google.actions.sdk.v2.conversation.Slot.status:type_name -> google.actions.sdk.v2.conversation.Slot.SlotStatus
	4, // 2: google.actions.sdk.v2.conversation.Slot.value:type_name -> google.protobuf.Value
	5, // 3: google.actions.sdk.v2.conversation.Slot.prompt:type_name -> google.actions.sdk.v2.conversation.Prompt
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_conversation_scene_proto_init() }
func file_google_actions_sdk_v2_conversation_scene_proto_init() {
	if File_google_actions_sdk_v2_conversation_scene_proto != nil {
		return
	}
	file_google_actions_sdk_v2_conversation_prompt_prompt_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_conversation_scene_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slot); i {
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
			RawDescriptor: file_google_actions_sdk_v2_conversation_scene_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_conversation_scene_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_conversation_scene_proto_depIdxs,
		EnumInfos:         file_google_actions_sdk_v2_conversation_scene_proto_enumTypes,
		MessageInfos:      file_google_actions_sdk_v2_conversation_scene_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_conversation_scene_proto = out.File
	file_google_actions_sdk_v2_conversation_scene_proto_rawDesc = nil
	file_google_actions_sdk_v2_conversation_scene_proto_goTypes = nil
	file_google_actions_sdk_v2_conversation_scene_proto_depIdxs = nil
}
