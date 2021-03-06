// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package golang1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/amsokol/protoc-gen-gotag/tagger"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Data struct {
	BoolValue            bool                  `protobuf:"varint,1,opt,name=boolValue,proto3" json:"boolValue,omitempty"`
	Int64Value           int64                 `protobuf:"varint,2,opt,name=int64Value,proto3" json:"int64Value,omitempty"`
	DoubleValue          float64               `protobuf:"fixed64,3,opt,name=doubleValue,proto3" json:"doubleValue,omitempty"`
	StringValue          string                `protobuf:"bytes,4,opt,name=stringValue,proto3" json:"stringValue,omitempty"`
	TimestampValue       *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=timestampValue,proto3" json:"timestampValue,omitempty" bson:"date"`
	BoolWrappedValue     *wrappers.BoolValue   `protobuf:"bytes,6,opt,name=boolWrappedValue,proto3" json:"boolWrappedValue,omitempty"`
	Int64WrappedValue    *wrappers.Int64Value  `protobuf:"bytes,7,opt,name=int64WrappedValue,proto3" json:"int64WrappedValue,omitempty"`
	DoubleWrappedValue   *wrappers.DoubleValue `protobuf:"bytes,8,opt,name=doubleWrappedValue,proto3" json:"doubleWrappedValue,omitempty"`
	StringWrappedValue   *wrappers.StringValue `protobuf:"bytes,9,opt,name=stringWrappedValue,proto3" json:"stringWrappedValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" bson:"-"`
	XXX_unrecognized     []byte                `json:"-" bson:"-"`
	XXX_sizecache        int32                 `json:"-" bson:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_fb2e72a0fe485112, []int{0}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (dst *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(dst, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *Data) GetInt64Value() int64 {
	if m != nil {
		return m.Int64Value
	}
	return 0
}

func (m *Data) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *Data) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *Data) GetTimestampValue() *timestamp.Timestamp {
	if m != nil {
		return m.TimestampValue
	}
	return nil
}

func (m *Data) GetBoolWrappedValue() *wrappers.BoolValue {
	if m != nil {
		return m.BoolWrappedValue
	}
	return nil
}

func (m *Data) GetInt64WrappedValue() *wrappers.Int64Value {
	if m != nil {
		return m.Int64WrappedValue
	}
	return nil
}

func (m *Data) GetDoubleWrappedValue() *wrappers.DoubleValue {
	if m != nil {
		return m.DoubleWrappedValue
	}
	return nil
}

func (m *Data) GetStringWrappedValue() *wrappers.StringValue {
	if m != nil {
		return m.StringWrappedValue
	}
	return nil
}

func init() {
	proto.RegisterType((*Data)(nil), "main.Data")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_data_fb2e72a0fe485112) }

var fileDescriptor_data_fb2e72a0fe485112 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xc9, 0xbb, 0xbe, 0x73, 0x4b, 0x41, 0x66, 0x4e, 0x52, 0xc7, 0x0c, 0xc5, 0x43, 0x2f,
	0x6b, 0x41, 0xc5, 0x83, 0xc7, 0x31, 0x84, 0x81, 0xa7, 0x38, 0xf4, 0xfc, 0xaf, 0x8d, 0xa1, 0xd0,
	0x36, 0xa5, 0x4d, 0xf1, 0x0b, 0x78, 0xf5, 0x0b, 0x78, 0xf0, 0xb3, 0x4a, 0x93, 0x75, 0x6d, 0xd7,
	0xed, 0x14, 0x78, 0xf2, 0xfb, 0xff, 0xe0, 0x79, 0x30, 0x8e, 0x40, 0x81, 0x9f, 0x17, 0x52, 0x49,
	0x62, 0xa5, 0x10, 0x67, 0xce, 0xb5, 0x90, 0x52, 0x24, 0x3c, 0xd0, 0x59, 0x58, 0x7d, 0x04, 0x2a,
	0x4e, 0x79, 0xa9, 0x20, 0xcd, 0x0d, 0xe6, 0x2c, 0x0e, 0x81, 0xcf, 0x02, 0xf2, 0x9c, 0x17, 0xe5,
	0xee, 0xff, 0x46, 0x3f, 0xef, 0x4b, 0xc1, 0xb3, 0xa5, 0x90, 0x0a, 0x44, 0xa0, 0x40, 0x08, 0x5e,
	0xec, 0x1e, 0x43, 0xb9, 0xdf, 0x16, 0xb6, 0xd6, 0xa0, 0x80, 0xcc, 0xf1, 0x34, 0x94, 0x32, 0x79,
	0x85, 0xa4, 0xe2, 0x97, 0x88, 0x22, 0x6f, 0xc2, 0xda, 0x80, 0x2c, 0x30, 0x8e, 0x33, 0xf5, 0x70,
	0x6f, 0xbe, 0xff, 0x51, 0xe4, 0x8d, 0x58, 0x27, 0x21, 0x14, 0xdb, 0x91, 0xac, 0xc2, 0x84, 0x1b,
	0x60, 0x44, 0x91, 0x87, 0x58, 0x37, 0xaa, 0x89, 0x52, 0x15, 0x71, 0x26, 0x0c, 0x61, 0x51, 0xe4,
	0x4d, 0x59, 0x37, 0x22, 0x5b, 0x7c, 0xbe, 0xef, 0x68, 0xa0, 0xff, 0x14, 0x79, 0xf6, 0xad, 0xe3,
	0x9b, 0xa6, 0x7e, 0xd3, 0xd4, 0xdf, 0x36, 0xd8, 0x6a, 0xf6, 0xf3, 0xf5, 0x3b, 0xb2, 0xc3, 0x52,
	0x66, 0x8f, 0x6e, 0x04, 0x8a, 0xbb, 0xec, 0xc0, 0x41, 0x9e, 0xf0, 0xac, 0xae, 0xf1, 0xa6, 0xc7,
	0x89, 0x8c, 0x77, 0x7c, 0xc2, 0xbb, 0x6a, 0xfa, 0xb2, 0xc1, 0x0d, 0xd9, 0xe0, 0x0b, 0xdd, 0xb7,
	0x27, 0x3a, 0xd3, 0xa2, 0xab, 0x81, 0x68, 0xb3, 0x5f, 0x86, 0x0d, 0xaf, 0xc8, 0x33, 0x26, 0x66,
	0x99, 0x9e, 0x6b, 0xa2, 0x5d, 0xf3, 0x81, 0x6b, 0xdd, 0x8e, 0xc8, 0x8e, 0xdc, 0xd5, 0x36, 0xb3,
	0x62, 0xcf, 0x36, 0x3d, 0x61, 0x7b, 0x69, 0x07, 0x67, 0x47, 0xee, 0xc2, 0xb1, 0x26, 0xef, 0xfe,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xd8, 0xc5, 0x37, 0x91, 0x02, 0x00, 0x00,
}
