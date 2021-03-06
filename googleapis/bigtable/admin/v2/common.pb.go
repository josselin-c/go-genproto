// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/bigtable/admin/v2/common.proto
// DO NOT EDIT!

package google_bigtable_admin_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Storage media types for persisting Bigtable data.
type StorageType int32

const (
	// The user did not specify a storage type.
	StorageType_STORAGE_TYPE_UNSPECIFIED StorageType = 0
	// Flash (SSD) storage should be used.
	StorageType_SSD StorageType = 1
	// Magnetic drive (HDD) storage should be used.
	StorageType_HDD StorageType = 2
)

var StorageType_name = map[int32]string{
	0: "STORAGE_TYPE_UNSPECIFIED",
	1: "SSD",
	2: "HDD",
}
var StorageType_value = map[string]int32{
	"STORAGE_TYPE_UNSPECIFIED": 0,
	"SSD": 1,
	"HDD": 2,
}

func (x StorageType) String() string {
	return proto.EnumName(StorageType_name, int32(x))
}
func (StorageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func init() {
	proto.RegisterEnum("google.bigtable.admin.v2.StorageType", StorageType_name, StorageType_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/bigtable/admin/v2/common.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0xff, 0xfe, 0x82, 0xc2, 0x74, 0x33, 0xcc, 0xaa, 0x94, 0x3e, 0x81, 0x60, 0x2e, 0xd4,
	0xa5, 0xb8, 0xb0, 0xcd, 0xa8, 0xb3, 0xd1, 0xc1, 0x8c, 0x0b, 0x57, 0x25, 0x89, 0xe9, 0x35, 0x30,
	0xc9, 0x0d, 0x93, 0x74, 0xa0, 0x6f, 0x2f, 0xcd, 0x28, 0xae, 0x04, 0x77, 0x07, 0xce, 0xc9, 0xf7,
	0x91, 0x5b, 0x70, 0x24, 0xc2, 0xde, 0x30, 0xa4, 0x5e, 0x7a, 0x64, 0x34, 0x20, 0xa0, 0xf1, 0x61,
	0xa0, 0x44, 0x30, 0x55, 0x32, 0xd8, 0x08, 0xca, 0x62, 0x92, 0xaa, 0x37, 0x20, 0xdf, 0x9d, 0xf5,
	0x30, 0xae, 0x41, 0x93, 0x73, 0xe4, 0x59, 0x5e, 0x56, 0x8b, 0x2f, 0xca, 0xf7, 0x8c, 0xe5, 0x19,
	0x1b, 0xd7, 0xcb, 0xe6, 0x6f, 0x7c, 0x19, 0x2c, 0x44, 0x33, 0x8c, 0x56, 0x1b, 0x4d, 0x7e, 0x6f,
	0x11, 0xa4, 0xf7, 0x94, 0x64, 0xb2, 0xe4, 0xe3, 0x24, 0x59, 0xde, 0xa0, 0x4d, 0x1f, 0x07, 0xc5,
	0x34, 0x39, 0x98, 0x70, 0x90, 0x0b, 0x75, 0xd8, 0x43, 0x48, 0xc7, 0x60, 0x22, 0x24, 0xeb, 0x4c,
	0x4c, 0xd2, 0x85, 0x9f, 0x34, 0x3d, 0xbe, 0xbc, 0x2d, 0xe6, 0x22, 0xd1, 0x20, 0xd1, 0x74, 0xc7,
	0x60, 0xaa, 0x55, 0xb1, 0x10, 0xdd, 0xf3, 0xcb, 0xdd, 0x43, 0xbd, 0xeb, 0xde, 0xda, 0x7a, 0xf7,
	0xfa, 0x24, 0xda, 0x7a, 0xdb, 0xdc, 0x37, 0x35, 0x2f, 0xff, 0x55, 0x17, 0xc5, 0x99, 0x10, 0xbc,
	0x9c, 0x9d, 0xc2, 0x23, 0xe7, 0xe5, 0xff, 0xcd, 0x55, 0xb1, 0xd2, 0xe4, 0xd8, 0x6f, 0xdf, 0xdc,
	0xcc, 0xb7, 0xf9, 0x1c, 0xed, 0xc9, 0xd5, 0xce, 0xd4, 0x79, 0x96, 0x5e, 0x7f, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x7f, 0xfd, 0xc6, 0x6a, 0x5e, 0x01, 0x00, 0x00,
}
