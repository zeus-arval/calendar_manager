// Copyright 2021 Google LLC
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
// source: google/type/phone_number.proto

package phone_number

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

// An object representing a phone number, suitable as an API wire format.
//
// This representation:
//
//  - should not be used for locale-specific formatting of a phone number, such
//    as "+1 (650) 253-0000 ext. 123"
//
//  - is not designed for efficient storage
//  - may not be suitable for dialing - specialized libraries (see references)
//    should be used to parse the number for that purpose
//
// To do something meaningful with this number, such as format it for various
// use-cases, convert it to an `i18n.phonenumbers.PhoneNumber` object first.
//
// For instance, in Java this would be:
//
//    com.google.type.PhoneNumber wireProto =
//        com.google.type.PhoneNumber.newBuilder().build();
//    com.google.i18n.phonenumbers.Phonenumber.PhoneNumber phoneNumber =
//        PhoneNumberUtil.getInstance().parse(wireProto.getE164Number(), "ZZ");
//    if (!wireProto.getExtension().isEmpty()) {
//      phoneNumber.setExtension(wireProto.getExtension());
//    }
//
//  Reference(s):
//   - https://github.com/google/libphonenumber
type PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.  Either a regular number, or a short code.  New fields may be
	// added to the oneof below in the future, so clients should ignore phone
	// numbers for which none of the fields they coded against are set.
	//
	// Types that are assignable to Kind:
	//	*PhoneNumber_E164Number
	//	*PhoneNumber_ShortCode_
	Kind isPhoneNumber_Kind `protobuf_oneof:"kind"`
	// The phone number's extension. The extension is not standardized in ITU
	// recommendations, except for being defined as a series of numbers with a
	// maximum length of 40 digits. Other than digits, some other dialing
	// characters such as ',' (indicating a wait) or '#' may be stored here.
	//
	// Note that no regions currently use extensions with short codes, so this
	// field is normally only set in conjunction with an E.164 number. It is held
	// separately from the E.164 number to allow for short code extensions in the
	// future.
	Extension string `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (x *PhoneNumber) Reset() {
	*x = PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_type_phone_number_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneNumber) ProtoMessage() {}

func (x *PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_google_type_phone_number_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneNumber.ProtoReflect.Descriptor instead.
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return file_google_type_phone_number_proto_rawDescGZIP(), []int{0}
}

func (m *PhoneNumber) GetKind() isPhoneNumber_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *PhoneNumber) GetE164Number() string {
	if x, ok := x.GetKind().(*PhoneNumber_E164Number); ok {
		return x.E164Number
	}
	return ""
}

func (x *PhoneNumber) GetShortCode() *PhoneNumber_ShortCode {
	if x, ok := x.GetKind().(*PhoneNumber_ShortCode_); ok {
		return x.ShortCode
	}
	return nil
}

func (x *PhoneNumber) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

type isPhoneNumber_Kind interface {
	isPhoneNumber_Kind()
}

type PhoneNumber_E164Number struct {
	// The phone number, represented as a leading plus sign ('+'), followed by a
	// phone number that uses a relaxed ITU E.164 format consisting of the
	// country calling code (1 to 3 digits) and the subscriber number, with no
	// additional spaces or formatting, e.g.:
	//  - correct: "+15552220123"
	//  - incorrect: "+1 (555) 222-01234 x123".
	//
	// The ITU E.164 format limits the latter to 12 digits, but in practice not
	// all countries respect that, so we relax that restriction here.
	// National-only numbers are not allowed.
	//
	// References:
	//  - https://www.itu.int/rec/T-REC-E.164-201011-I
	//  - https://en.wikipedia.org/wiki/E.164.
	//  - https://en.wikipedia.org/wiki/List_of_country_calling_codes
	E164Number string `protobuf:"bytes,1,opt,name=e164_number,json=e164Number,proto3,oneof"`
}

type PhoneNumber_ShortCode_ struct {
	// A short code.
	//
	// Reference(s):
	//  - https://en.wikipedia.org/wiki/Short_code
	ShortCode *PhoneNumber_ShortCode `protobuf:"bytes,2,opt,name=short_code,json=shortCode,proto3,oneof"`
}

func (*PhoneNumber_E164Number) isPhoneNumber_Kind() {}

func (*PhoneNumber_ShortCode_) isPhoneNumber_Kind() {}

// An object representing a short code, which is a phone number that is
// typically much shorter than regular phone numbers and can be used to
// address messages in MMS and SMS systems, as well as for abbreviated dialing
// (e.g. "Text 611 to see how many minutes you have remaining on your plan.").
//
// Short codes are restricted to a region and are not internationally
// dialable, which means the same short code can exist in different regions,
// with different usage and pricing, even if those regions share the same
// country calling code (e.g. US and CA).
type PhoneNumber_ShortCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The BCP-47 region code of the location where calls to this
	// short code can be made, such as "US" and "BB".
	//
	// Reference(s):
	//  - http://www.unicode.org/reports/tr35/#unicode_region_subtag
	RegionCode string `protobuf:"bytes,1,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// Required. The short code digits, without a leading plus ('+') or country
	// calling code, e.g. "611".
	Number string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PhoneNumber_ShortCode) Reset() {
	*x = PhoneNumber_ShortCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_type_phone_number_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneNumber_ShortCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneNumber_ShortCode) ProtoMessage() {}

func (x *PhoneNumber_ShortCode) ProtoReflect() protoreflect.Message {
	mi := &file_google_type_phone_number_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneNumber_ShortCode.ProtoReflect.Descriptor instead.
func (*PhoneNumber_ShortCode) Descriptor() ([]byte, []int) {
	return file_google_type_phone_number_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PhoneNumber_ShortCode) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *PhoneNumber_ShortCode) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

var File_google_type_phone_number_proto protoreflect.FileDescriptor

var file_google_type_phone_number_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x22, 0xe1, 0x01,
	0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0b, 0x65, 0x31, 0x36, 0x34, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x31, 0x36, 0x34, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x43, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x1a, 0x44, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x42, 0x74, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x42, 0x10, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x3b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0xf8, 0x01,
	0x01, 0xa2, 0x02, 0x03, 0x47, 0x54, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_type_phone_number_proto_rawDescOnce sync.Once
	file_google_type_phone_number_proto_rawDescData = file_google_type_phone_number_proto_rawDesc
)

func file_google_type_phone_number_proto_rawDescGZIP() []byte {
	file_google_type_phone_number_proto_rawDescOnce.Do(func() {
		file_google_type_phone_number_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_type_phone_number_proto_rawDescData)
	})
	return file_google_type_phone_number_proto_rawDescData
}

var file_google_type_phone_number_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_type_phone_number_proto_goTypes = []interface{}{
	(*PhoneNumber)(nil),           // 0: google.type.PhoneNumber
	(*PhoneNumber_ShortCode)(nil), // 1: google.type.PhoneNumber.ShortCode
}
var file_google_type_phone_number_proto_depIdxs = []int32{
	1, // 0: google.type.PhoneNumber.short_code:type_name -> google.type.PhoneNumber.ShortCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_type_phone_number_proto_init() }
func file_google_type_phone_number_proto_init() {
	if File_google_type_phone_number_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_type_phone_number_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneNumber); i {
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
		file_google_type_phone_number_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneNumber_ShortCode); i {
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
	file_google_type_phone_number_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PhoneNumber_E164Number)(nil),
		(*PhoneNumber_ShortCode_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_type_phone_number_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_type_phone_number_proto_goTypes,
		DependencyIndexes: file_google_type_phone_number_proto_depIdxs,
		MessageInfos:      file_google_type_phone_number_proto_msgTypes,
	}.Build()
	File_google_type_phone_number_proto = out.File
	file_google_type_phone_number_proto_rawDesc = nil
	file_google_type_phone_number_proto_goTypes = nil
	file_google_type_phone_number_proto_depIdxs = nil
}
