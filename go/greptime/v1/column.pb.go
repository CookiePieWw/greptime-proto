// Copyright 2023 Greptime Team
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: greptime/v1/column.proto

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

type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColumnName   string       `protobuf:"bytes,1,opt,name=column_name,json=columnName,proto3" json:"column_name,omitempty"`
	SemanticType SemanticType `protobuf:"varint,2,opt,name=semantic_type,json=semanticType,proto3,enum=greptime.v1.SemanticType" json:"semantic_type,omitempty"`
	// The array of non-null values in this column.
	//
	// For example: suppose there is a column "foo" that contains some int32
	// values (1, 2, 3, 4, 5, null, 7, 8, 9, null);
	//
	//	column:
	//	  column_name: foo
	//	  semantic_type: Tag
	//	  values: 1, 2, 3, 4, 5, 7, 8, 9
	//	  null_masks: 00100000 00000010
	Values *Column_Values `protobuf:"bytes,3,opt,name=values,proto3" json:"values,omitempty"`
	// Mask maps the positions of null values.
	// If a bit in null_mask is 1, it indicates that the column value at that
	// position is null.
	NullMask []byte `protobuf:"bytes,4,opt,name=null_mask,json=nullMask,proto3" json:"null_mask,omitempty"`
	// Helpful in creating vector from column.
	Datatype ColumnDataType `protobuf:"varint,5,opt,name=datatype,proto3,enum=greptime.v1.ColumnDataType" json:"datatype,omitempty"`
	// Extension for ColumnDataType.
	DatatypeExtension *ColumnDataTypeExtension `protobuf:"bytes,6,opt,name=datatype_extension,json=datatypeExtension,proto3" json:"datatype_extension,omitempty"`
}

func (x *Column) Reset() {
	*x = Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_column_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_column_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_greptime_v1_column_proto_rawDescGZIP(), []int{0}
}

func (x *Column) GetColumnName() string {
	if x != nil {
		return x.ColumnName
	}
	return ""
}

func (x *Column) GetSemanticType() SemanticType {
	if x != nil {
		return x.SemanticType
	}
	return SemanticType_TAG
}

func (x *Column) GetValues() *Column_Values {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Column) GetNullMask() []byte {
	if x != nil {
		return x.NullMask
	}
	return nil
}

func (x *Column) GetDatatype() ColumnDataType {
	if x != nil {
		return x.Datatype
	}
	return ColumnDataType_BOOLEAN
}

func (x *Column) GetDatatypeExtension() *ColumnDataTypeExtension {
	if x != nil {
		return x.DatatypeExtension
	}
	return nil
}

type Column_Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I8Values                   []int32                 `protobuf:"varint,1,rep,packed,name=i8_values,json=i8Values,proto3" json:"i8_values,omitempty"`
	I16Values                  []int32                 `protobuf:"varint,2,rep,packed,name=i16_values,json=i16Values,proto3" json:"i16_values,omitempty"`
	I32Values                  []int32                 `protobuf:"varint,3,rep,packed,name=i32_values,json=i32Values,proto3" json:"i32_values,omitempty"`
	I64Values                  []int64                 `protobuf:"varint,4,rep,packed,name=i64_values,json=i64Values,proto3" json:"i64_values,omitempty"`
	U8Values                   []uint32                `protobuf:"varint,5,rep,packed,name=u8_values,json=u8Values,proto3" json:"u8_values,omitempty"`
	U16Values                  []uint32                `protobuf:"varint,6,rep,packed,name=u16_values,json=u16Values,proto3" json:"u16_values,omitempty"`
	U32Values                  []uint32                `protobuf:"varint,7,rep,packed,name=u32_values,json=u32Values,proto3" json:"u32_values,omitempty"`
	U64Values                  []uint64                `protobuf:"varint,8,rep,packed,name=u64_values,json=u64Values,proto3" json:"u64_values,omitempty"`
	F32Values                  []float32               `protobuf:"fixed32,9,rep,packed,name=f32_values,json=f32Values,proto3" json:"f32_values,omitempty"`
	F64Values                  []float64               `protobuf:"fixed64,10,rep,packed,name=f64_values,json=f64Values,proto3" json:"f64_values,omitempty"`
	BoolValues                 []bool                  `protobuf:"varint,11,rep,packed,name=bool_values,json=boolValues,proto3" json:"bool_values,omitempty"`
	BinaryValues               [][]byte                `protobuf:"bytes,12,rep,name=binary_values,json=binaryValues,proto3" json:"binary_values,omitempty"`
	StringValues               []string                `protobuf:"bytes,13,rep,name=string_values,json=stringValues,proto3" json:"string_values,omitempty"`
	DateValues                 []int32                 `protobuf:"varint,14,rep,packed,name=date_values,json=dateValues,proto3" json:"date_values,omitempty"`
	DatetimeValues             []int64                 `protobuf:"varint,15,rep,packed,name=datetime_values,json=datetimeValues,proto3" json:"datetime_values,omitempty"`
	TimestampSecondValues      []int64                 `protobuf:"varint,16,rep,packed,name=timestamp_second_values,json=timestampSecondValues,proto3" json:"timestamp_second_values,omitempty"`
	TimestampMillisecondValues []int64                 `protobuf:"varint,17,rep,packed,name=timestamp_millisecond_values,json=timestampMillisecondValues,proto3" json:"timestamp_millisecond_values,omitempty"`
	TimestampMicrosecondValues []int64                 `protobuf:"varint,18,rep,packed,name=timestamp_microsecond_values,json=timestampMicrosecondValues,proto3" json:"timestamp_microsecond_values,omitempty"`
	TimestampNanosecondValues  []int64                 `protobuf:"varint,19,rep,packed,name=timestamp_nanosecond_values,json=timestampNanosecondValues,proto3" json:"timestamp_nanosecond_values,omitempty"`
	TimeSecondValues           []int64                 `protobuf:"varint,20,rep,packed,name=time_second_values,json=timeSecondValues,proto3" json:"time_second_values,omitempty"`
	TimeMillisecondValues      []int64                 `protobuf:"varint,21,rep,packed,name=time_millisecond_values,json=timeMillisecondValues,proto3" json:"time_millisecond_values,omitempty"`
	TimeMicrosecondValues      []int64                 `protobuf:"varint,22,rep,packed,name=time_microsecond_values,json=timeMicrosecondValues,proto3" json:"time_microsecond_values,omitempty"`
	TimeNanosecondValues       []int64                 `protobuf:"varint,23,rep,packed,name=time_nanosecond_values,json=timeNanosecondValues,proto3" json:"time_nanosecond_values,omitempty"`
	IntervalYearMonthValues    []int32                 `protobuf:"varint,24,rep,packed,name=interval_year_month_values,json=intervalYearMonthValues,proto3" json:"interval_year_month_values,omitempty"`
	IntervalDayTimeValues      []int64                 `protobuf:"varint,25,rep,packed,name=interval_day_time_values,json=intervalDayTimeValues,proto3" json:"interval_day_time_values,omitempty"`
	IntervalMonthDayNanoValues []*IntervalMonthDayNano `protobuf:"bytes,26,rep,name=interval_month_day_nano_values,json=intervalMonthDayNanoValues,proto3" json:"interval_month_day_nano_values,omitempty"`
	Decimal128Values           []*Decimal128           `protobuf:"bytes,31,rep,name=decimal128_values,json=decimal128Values,proto3" json:"decimal128_values,omitempty"`
}

func (x *Column_Values) Reset() {
	*x = Column_Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_column_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column_Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column_Values) ProtoMessage() {}

func (x *Column_Values) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_column_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column_Values.ProtoReflect.Descriptor instead.
func (*Column_Values) Descriptor() ([]byte, []int) {
	return file_greptime_v1_column_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Column_Values) GetI8Values() []int32 {
	if x != nil {
		return x.I8Values
	}
	return nil
}

func (x *Column_Values) GetI16Values() []int32 {
	if x != nil {
		return x.I16Values
	}
	return nil
}

func (x *Column_Values) GetI32Values() []int32 {
	if x != nil {
		return x.I32Values
	}
	return nil
}

func (x *Column_Values) GetI64Values() []int64 {
	if x != nil {
		return x.I64Values
	}
	return nil
}

func (x *Column_Values) GetU8Values() []uint32 {
	if x != nil {
		return x.U8Values
	}
	return nil
}

func (x *Column_Values) GetU16Values() []uint32 {
	if x != nil {
		return x.U16Values
	}
	return nil
}

func (x *Column_Values) GetU32Values() []uint32 {
	if x != nil {
		return x.U32Values
	}
	return nil
}

func (x *Column_Values) GetU64Values() []uint64 {
	if x != nil {
		return x.U64Values
	}
	return nil
}

func (x *Column_Values) GetF32Values() []float32 {
	if x != nil {
		return x.F32Values
	}
	return nil
}

func (x *Column_Values) GetF64Values() []float64 {
	if x != nil {
		return x.F64Values
	}
	return nil
}

func (x *Column_Values) GetBoolValues() []bool {
	if x != nil {
		return x.BoolValues
	}
	return nil
}

func (x *Column_Values) GetBinaryValues() [][]byte {
	if x != nil {
		return x.BinaryValues
	}
	return nil
}

func (x *Column_Values) GetStringValues() []string {
	if x != nil {
		return x.StringValues
	}
	return nil
}

func (x *Column_Values) GetDateValues() []int32 {
	if x != nil {
		return x.DateValues
	}
	return nil
}

func (x *Column_Values) GetDatetimeValues() []int64 {
	if x != nil {
		return x.DatetimeValues
	}
	return nil
}

func (x *Column_Values) GetTimestampSecondValues() []int64 {
	if x != nil {
		return x.TimestampSecondValues
	}
	return nil
}

func (x *Column_Values) GetTimestampMillisecondValues() []int64 {
	if x != nil {
		return x.TimestampMillisecondValues
	}
	return nil
}

func (x *Column_Values) GetTimestampMicrosecondValues() []int64 {
	if x != nil {
		return x.TimestampMicrosecondValues
	}
	return nil
}

func (x *Column_Values) GetTimestampNanosecondValues() []int64 {
	if x != nil {
		return x.TimestampNanosecondValues
	}
	return nil
}

func (x *Column_Values) GetTimeSecondValues() []int64 {
	if x != nil {
		return x.TimeSecondValues
	}
	return nil
}

func (x *Column_Values) GetTimeMillisecondValues() []int64 {
	if x != nil {
		return x.TimeMillisecondValues
	}
	return nil
}

func (x *Column_Values) GetTimeMicrosecondValues() []int64 {
	if x != nil {
		return x.TimeMicrosecondValues
	}
	return nil
}

func (x *Column_Values) GetTimeNanosecondValues() []int64 {
	if x != nil {
		return x.TimeNanosecondValues
	}
	return nil
}

func (x *Column_Values) GetIntervalYearMonthValues() []int32 {
	if x != nil {
		return x.IntervalYearMonthValues
	}
	return nil
}

func (x *Column_Values) GetIntervalDayTimeValues() []int64 {
	if x != nil {
		return x.IntervalDayTimeValues
	}
	return nil
}

func (x *Column_Values) GetIntervalMonthDayNanoValues() []*IntervalMonthDayNano {
	if x != nil {
		return x.IntervalMonthDayNanoValues
	}
	return nil
}

func (x *Column_Values) GetDecimal128Values() []*Decimal128 {
	if x != nil {
		return x.Decimal128Values
	}
	return nil
}

var File_greptime_v1_column_proto protoreflect.FileDescriptor

var file_greptime_v1_column_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x65, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xad, 0x0c, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a,
	0x0d, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x37,
	0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x12, 0x53, 0x0a, 0x12, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x79, 0x70, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0xe2, 0x09, 0x0a,
	0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x38, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x69, 0x38, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x31, 0x36, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x69, 0x31, 0x36, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x69, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x69, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x38, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x75, 0x38, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x31, 0x36, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x09, 0x75, 0x31, 0x36, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x09, 0x75, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x09, 0x75, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x02, 0x52,
	0x09, 0x66, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x36,
	0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x01, 0x52, 0x09,
	0x66, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x6f,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x08, 0x52, 0x0a,
	0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x0c, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0e,
	0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x36,
	0x0a, 0x17, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x15, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x1c, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x03, 0x52, 0x1a, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x1c, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x03, 0x52, 0x1a,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x1b, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x19, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4e, 0x61, 0x6e, 0x6f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x03, 0x52, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x03, 0x52, 0x15, 0x74, 0x69, 0x6d, 0x65, 0x4d,
	0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x36, 0x0a, 0x17, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x15, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x03, 0x52, 0x14, 0x74, 0x69, 0x6d, 0x65, 0x4e, 0x61,
	0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x3b,
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x5f,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x18, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x59, 0x65, 0x61, 0x72,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x19, 0x20, 0x03, 0x28, 0x03, 0x52, 0x15, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x44, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x65, 0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x44, 0x61, 0x79, 0x4e, 0x61, 0x6e, 0x6f, 0x52,
	0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x44, 0x61,
	0x79, 0x4e, 0x61, 0x6e, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x11, 0x64,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x31, 0x32, 0x38, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x1f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x31, 0x32, 0x38, 0x52,
	0x10, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x31, 0x32, 0x38, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x42, 0x50, 0x0a, 0x0e, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x07, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_column_proto_rawDescOnce sync.Once
	file_greptime_v1_column_proto_rawDescData = file_greptime_v1_column_proto_rawDesc
)

func file_greptime_v1_column_proto_rawDescGZIP() []byte {
	file_greptime_v1_column_proto_rawDescOnce.Do(func() {
		file_greptime_v1_column_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_column_proto_rawDescData)
	})
	return file_greptime_v1_column_proto_rawDescData
}

var file_greptime_v1_column_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_greptime_v1_column_proto_goTypes = []interface{}{
	(*Column)(nil),                  // 0: greptime.v1.Column
	(*Column_Values)(nil),           // 1: greptime.v1.Column.Values
	(SemanticType)(0),               // 2: greptime.v1.SemanticType
	(ColumnDataType)(0),             // 3: greptime.v1.ColumnDataType
	(*ColumnDataTypeExtension)(nil), // 4: greptime.v1.ColumnDataTypeExtension
	(*IntervalMonthDayNano)(nil),    // 5: greptime.v1.IntervalMonthDayNano
	(*Decimal128)(nil),              // 6: greptime.v1.Decimal128
}
var file_greptime_v1_column_proto_depIdxs = []int32{
	2, // 0: greptime.v1.Column.semantic_type:type_name -> greptime.v1.SemanticType
	1, // 1: greptime.v1.Column.values:type_name -> greptime.v1.Column.Values
	3, // 2: greptime.v1.Column.datatype:type_name -> greptime.v1.ColumnDataType
	4, // 3: greptime.v1.Column.datatype_extension:type_name -> greptime.v1.ColumnDataTypeExtension
	5, // 4: greptime.v1.Column.Values.interval_month_day_nano_values:type_name -> greptime.v1.IntervalMonthDayNano
	6, // 5: greptime.v1.Column.Values.decimal128_values:type_name -> greptime.v1.Decimal128
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_greptime_v1_column_proto_init() }
func file_greptime_v1_column_proto_init() {
	if File_greptime_v1_column_proto != nil {
		return
	}
	file_greptime_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_column_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column); i {
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
		file_greptime_v1_column_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column_Values); i {
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
			RawDescriptor: file_greptime_v1_column_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_greptime_v1_column_proto_goTypes,
		DependencyIndexes: file_greptime_v1_column_proto_depIdxs,
		MessageInfos:      file_greptime_v1_column_proto_msgTypes,
	}.Build()
	File_greptime_v1_column_proto = out.File
	file_greptime_v1_column_proto_rawDesc = nil
	file_greptime_v1_column_proto_goTypes = nil
	file_greptime_v1_column_proto_depIdxs = nil
}
