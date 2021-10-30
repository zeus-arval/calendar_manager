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
// source: google/cloud/bigquery/migration/v2alpha/migration_metrics.proto

package migration

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	distribution "google.golang.org/genproto/googleapis/api/distribution"
	metric "google.golang.org/genproto/googleapis/api/metric"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The metrics object for a SubTask.
type TimeSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the metric.
	//
	// If the metric is not known by the service yet, it will be auto-created.
	Metric string `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	// Required. The value type of the time series.
	ValueType metric.MetricDescriptor_ValueType `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=google.api.MetricDescriptor_ValueType" json:"value_type,omitempty"`
	// Optional. The metric kind of the time series.
	//
	// If present, it must be the same as the metric kind of the associated
	// metric. If the associated metric's descriptor must be auto-created, then
	// this field specifies the metric kind of the new descriptor and must be
	// either `GAUGE` (the default) or `CUMULATIVE`.
	MetricKind metric.MetricDescriptor_MetricKind `protobuf:"varint,3,opt,name=metric_kind,json=metricKind,proto3,enum=google.api.MetricDescriptor_MetricKind" json:"metric_kind,omitempty"`
	// Required. The data points of this time series. When listing time series, points are
	// returned in reverse time order.
	//
	// When creating a time series, this field must contain exactly one point and
	// the point's type must be the same as the value type of the associated
	// metric. If the associated metric's descriptor must be auto-created, then
	// the value type of the descriptor is determined by the point's type, which
	// must be `BOOL`, `INT64`, `DOUBLE`, or `DISTRIBUTION`.
	Points []*Point `protobuf:"bytes,4,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *TimeSeries) Reset() {
	*x = TimeSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeries) ProtoMessage() {}

func (x *TimeSeries) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeries.ProtoReflect.Descriptor instead.
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *TimeSeries) GetMetric() string {
	if x != nil {
		return x.Metric
	}
	return ""
}

func (x *TimeSeries) GetValueType() metric.MetricDescriptor_ValueType {
	if x != nil {
		return x.ValueType
	}
	return metric.MetricDescriptor_VALUE_TYPE_UNSPECIFIED
}

func (x *TimeSeries) GetMetricKind() metric.MetricDescriptor_MetricKind {
	if x != nil {
		return x.MetricKind
	}
	return metric.MetricDescriptor_METRIC_KIND_UNSPECIFIED
}

func (x *TimeSeries) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

// A single data point in a time series.
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time interval to which the data point applies.  For `GAUGE` metrics,
	// the start time does not need to be supplied, but if it is supplied, it must
	// equal the end time.  For `DELTA` metrics, the start and end time should
	// specify a non-zero interval, with subsequent points specifying contiguous
	// and non-overlapping intervals.  For `CUMULATIVE` metrics, the start and end
	// time should specify a non-zero interval, with subsequent points specifying
	// the same start time and increasing end times, until an event resets the
	// cumulative value to zero and sets a new start time for the following
	// points.
	Interval *TimeInterval `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	// The value of the data point.
	Value *TypedValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *Point) GetInterval() *TimeInterval {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *Point) GetValue() *TypedValue {
	if x != nil {
		return x.Value
	}
	return nil
}

// A time interval extending just after a start time through an end time.
// If the start time is the same as the end time, then the interval
// represents a single point in time.
type TimeInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The beginning of the time interval.  The default value
	// for the start time is the end time. The start time must not be
	// later than the end time.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Required. The end of the time interval.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *TimeInterval) Reset() {
	*x = TimeInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeInterval) ProtoMessage() {}

func (x *TimeInterval) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeInterval.ProtoReflect.Descriptor instead.
func (*TimeInterval) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDescGZIP(), []int{2}
}

func (x *TimeInterval) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TimeInterval) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

// A single strongly-typed value.
type TypedValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The typed value field.
	//
	// Types that are assignable to Value:
	//	*TypedValue_BoolValue
	//	*TypedValue_Int64Value
	//	*TypedValue_DoubleValue
	//	*TypedValue_StringValue
	//	*TypedValue_DistributionValue
	Value isTypedValue_Value `protobuf_oneof:"value"`
}

func (x *TypedValue) Reset() {
	*x = TypedValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypedValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypedValue) ProtoMessage() {}

func (x *TypedValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypedValue.ProtoReflect.Descriptor instead.
func (*TypedValue) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDescGZIP(), []int{3}
}

func (m *TypedValue) GetValue() isTypedValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *TypedValue) GetBoolValue() bool {
	if x, ok := x.GetValue().(*TypedValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *TypedValue) GetInt64Value() int64 {
	if x, ok := x.GetValue().(*TypedValue_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (x *TypedValue) GetDoubleValue() float64 {
	if x, ok := x.GetValue().(*TypedValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *TypedValue) GetStringValue() string {
	if x, ok := x.GetValue().(*TypedValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *TypedValue) GetDistributionValue() *distribution.Distribution {
	if x, ok := x.GetValue().(*TypedValue_DistributionValue); ok {
		return x.DistributionValue
	}
	return nil
}

type isTypedValue_Value interface {
	isTypedValue_Value()
}

type TypedValue_BoolValue struct {
	// A Boolean value: `true` or `false`.
	BoolValue bool `protobuf:"varint,1,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type TypedValue_Int64Value struct {
	// A 64-bit integer. Its range is approximately +/-9.2x10^18.
	Int64Value int64 `protobuf:"varint,2,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type TypedValue_DoubleValue struct {
	// A 64-bit double-precision floating-point number. Its magnitude
	// is approximately +/-10^(+/-300) and it has 16 significant digits of
	// precision.
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type TypedValue_StringValue struct {
	// A variable-length string value.
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type TypedValue_DistributionValue struct {
	// A distribution value.
	DistributionValue *distribution.Distribution `protobuf:"bytes,5,opt,name=distribution_value,json=distributionValue,proto3,oneof"`
}

func (*TypedValue_BoolValue) isTypedValue_Value() {}

func (*TypedValue_Int64Value) isTypedValue_Value() {}

func (*TypedValue_DoubleValue) isTypedValue_Value() {}

func (*TypedValue_StringValue) isTypedValue_Value() {}

func (*TypedValue_DistributionValue) isTypedValue_Value() {}

var File_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto protoreflect.FileDescriptor

var file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1d, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x02, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x12, 0x4a, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4d, 0x0a, 0x0b,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4b, 0x69, 0x6e, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x4b, 0x0a, 0x06, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x05, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x51, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x49, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xee, 0x01,
	0x0a, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a,
	0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a,
	0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x49, 0x0a, 0x12, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x11, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xec,
	0x01, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x15,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xaa, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x32, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0xca, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x4d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDescOnce sync.Once
	file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDescData = file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDesc
)

func file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDescGZIP() []byte {
	file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDescOnce.Do(func() {
		file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDescData)
	})
	return file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDescData
}

var file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_goTypes = []interface{}{
	(*TimeSeries)(nil),                      // 0: google.cloud.bigquery.migration.v2alpha.TimeSeries
	(*Point)(nil),                           // 1: google.cloud.bigquery.migration.v2alpha.Point
	(*TimeInterval)(nil),                    // 2: google.cloud.bigquery.migration.v2alpha.TimeInterval
	(*TypedValue)(nil),                      // 3: google.cloud.bigquery.migration.v2alpha.TypedValue
	(metric.MetricDescriptor_ValueType)(0),  // 4: google.api.MetricDescriptor.ValueType
	(metric.MetricDescriptor_MetricKind)(0), // 5: google.api.MetricDescriptor.MetricKind
	(*timestamppb.Timestamp)(nil),           // 6: google.protobuf.Timestamp
	(*distribution.Distribution)(nil),       // 7: google.api.Distribution
}
var file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_depIdxs = []int32{
	4, // 0: google.cloud.bigquery.migration.v2alpha.TimeSeries.value_type:type_name -> google.api.MetricDescriptor.ValueType
	5, // 1: google.cloud.bigquery.migration.v2alpha.TimeSeries.metric_kind:type_name -> google.api.MetricDescriptor.MetricKind
	1, // 2: google.cloud.bigquery.migration.v2alpha.TimeSeries.points:type_name -> google.cloud.bigquery.migration.v2alpha.Point
	2, // 3: google.cloud.bigquery.migration.v2alpha.Point.interval:type_name -> google.cloud.bigquery.migration.v2alpha.TimeInterval
	3, // 4: google.cloud.bigquery.migration.v2alpha.Point.value:type_name -> google.cloud.bigquery.migration.v2alpha.TypedValue
	6, // 5: google.cloud.bigquery.migration.v2alpha.TimeInterval.start_time:type_name -> google.protobuf.Timestamp
	6, // 6: google.cloud.bigquery.migration.v2alpha.TimeInterval.end_time:type_name -> google.protobuf.Timestamp
	7, // 7: google.cloud.bigquery.migration.v2alpha.TypedValue.distribution_value:type_name -> google.api.Distribution
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_init() }
func file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_init() {
	if File_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeries); i {
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
		file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeInterval); i {
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
		file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypedValue); i {
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
	file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*TypedValue_BoolValue)(nil),
		(*TypedValue_Int64Value)(nil),
		(*TypedValue_DoubleValue)(nil),
		(*TypedValue_StringValue)(nil),
		(*TypedValue_DistributionValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_goTypes,
		DependencyIndexes: file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_depIdxs,
		MessageInfos:      file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_msgTypes,
	}.Build()
	File_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto = out.File
	file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_rawDesc = nil
	file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_goTypes = nil
	file_google_cloud_bigquery_migration_v2alpha_migration_metrics_proto_depIdxs = nil
}
