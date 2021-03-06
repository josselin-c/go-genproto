// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/type/date/date.proto
// DO NOT EDIT!

/*
Package google_type is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/type/date/date.proto

It has these top-level messages:
	Date
*/
package google_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents a whole calendar date, e.g. date of birth. The time of day and
// time zone are either specified elsewhere or are not significant. The date
// is relative to the Proleptic Gregorian Calendar. The day may be 0 to
// represent a year and month where the day is not significant, e.g. credit card
// expiration date. The year may be 0 to represent a month and day independent
// of year, e.g. anniversary date. Related types are [google.type.TimeOfDay][google.type.TimeOfDay]
// and `google.protobuf.Timestamp`.
type Date struct {
	// Year of date. Must be from 1 to 9999, or 0 if specifying a date without
	// a year.
	Year int32 `protobuf:"varint,1,opt,name=year" json:"year,omitempty"`
	// Month of year. Must be from 1 to 12.
	Month int32 `protobuf:"varint,2,opt,name=month" json:"month,omitempty"`
	// Day of month. Must be from 1 to 31 and valid for the year and month, or 0
	// if specifying a year/month where the day is not significant.
	Day int32 `protobuf:"varint,3,opt,name=day" json:"day,omitempty"`
}

func (m *Date) Reset()                    { *m = Date{} }
func (m *Date) String() string            { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()               {}
func (*Date) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Date)(nil), "google.type.Date")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/type/date/date.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4b, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0x4f,
	0xcd, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x87, 0x48, 0x25, 0x16, 0x64, 0x16, 0xeb, 0x97, 0x54,
	0x16, 0xa4, 0xea, 0xa7, 0x24, 0x96, 0x40, 0x08, 0x3d, 0xb0, 0xbc, 0x10, 0x37, 0x54, 0x2f, 0x48,
	0x52, 0xc9, 0x89, 0x8b, 0xc5, 0x25, 0xb1, 0x24, 0x55, 0x48, 0x88, 0x8b, 0xa5, 0x32, 0x35, 0xb1,
	0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x08, 0xcc, 0x16, 0x12, 0xe1, 0x62, 0xcd, 0xcd, 0xcf,
	0x2b, 0xc9, 0x90, 0x60, 0x02, 0x0b, 0x42, 0x38, 0x42, 0x02, 0x5c, 0xcc, 0x29, 0x89, 0x95, 0x12,
	0xcc, 0x60, 0x31, 0x10, 0xd3, 0x49, 0x85, 0x8b, 0x3f, 0x39, 0x3f, 0x57, 0x0f, 0xc9, 0x58, 0x27,
	0x4e, 0x90, 0xa1, 0x01, 0x20, 0xeb, 0x02, 0x18, 0x17, 0x31, 0x31, 0xbb, 0x87, 0x04, 0x24, 0xb1,
	0x81, 0x6d, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xd4, 0x98, 0x3f, 0xbb, 0x00, 0x00,
	0x00,
}
