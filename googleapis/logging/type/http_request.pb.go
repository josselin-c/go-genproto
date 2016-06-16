// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/logging/type/http_request.proto
// DO NOT EDIT!

/*
Package google_logging_type is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/logging/type/http_request.proto
	google.golang.org/genproto/googleapis/logging/type/log_severity.proto

It has these top-level messages:
	HttpRequest
*/
package google_logging_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A common proto for logging HTTP requests.
//
type HttpRequest struct {
	// The request method. Examples: `"GET"`, `"HEAD"`, `"PUT"`, `"POST"`.
	RequestMethod string `protobuf:"bytes,1,opt,name=request_method,json=requestMethod" json:"request_method,omitempty"`
	// The scheme (http, https), the host name, the path and the query
	// portion of the URL that was requested.
	// Example: `"http://example.com/some/info?color=red"`.
	RequestUrl string `protobuf:"bytes,2,opt,name=request_url,json=requestUrl" json:"request_url,omitempty"`
	// The size of the HTTP request message in bytes, including the request
	// headers and the request body.
	RequestSize int64 `protobuf:"varint,3,opt,name=request_size,json=requestSize" json:"request_size,omitempty"`
	// The response code indicating the status of response.
	// Examples: 200, 404.
	Status int32 `protobuf:"varint,4,opt,name=status" json:"status,omitempty"`
	// The size of the HTTP response message sent back to the client, in bytes,
	// including the response headers and the response body.
	ResponseSize int64 `protobuf:"varint,5,opt,name=response_size,json=responseSize" json:"response_size,omitempty"`
	// The user agent sent by the client. Example:
	// `"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; Q312461; .NET CLR 1.0.3705)"`.
	UserAgent string `protobuf:"bytes,6,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	// The IP address (IPv4 or IPv6) of the client that issued the HTTP
	// request. Examples: `"192.168.1.1"`, `"FE80::0202:B3FF:FE1E:8329"`.
	RemoteIp string `protobuf:"bytes,7,opt,name=remote_ip,json=remoteIp" json:"remote_ip,omitempty"`
	// The IP address (IPv4 or IPv6) of the origin server that the request was
	// sent to.
	ServerIp string `protobuf:"bytes,13,opt,name=server_ip,json=serverIp" json:"server_ip,omitempty"`
	// The referer URL of the request, as defined in
	// [HTTP/1.1 Header Field Definitions](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).
	Referer string `protobuf:"bytes,8,opt,name=referer" json:"referer,omitempty"`
	// Whether or not a cache lookup was attempted.
	CacheLookup bool `protobuf:"varint,11,opt,name=cache_lookup,json=cacheLookup" json:"cache_lookup,omitempty"`
	// Whether or not an entity was served from cache
	// (with or without validation).
	CacheHit bool `protobuf:"varint,9,opt,name=cache_hit,json=cacheHit" json:"cache_hit,omitempty"`
	// Whether or not the response was validated with the origin server before
	// being served from cache. This field is only meaningful if `cache_hit` is
	// True.
	CacheValidatedWithOriginServer bool `protobuf:"varint,10,opt,name=cache_validated_with_origin_server,json=cacheValidatedWithOriginServer" json:"cache_validated_with_origin_server,omitempty"`
	// The number of HTTP response bytes inserted into cache. Set only when a
	// cache fill was attempted.
	CacheFillBytes int64 `protobuf:"varint,12,opt,name=cache_fill_bytes,json=cacheFillBytes" json:"cache_fill_bytes,omitempty"`
}

func (m *HttpRequest) Reset()                    { *m = HttpRequest{} }
func (m *HttpRequest) String() string            { return proto.CompactTextString(m) }
func (*HttpRequest) ProtoMessage()               {}
func (*HttpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*HttpRequest)(nil), "google.logging.type.HttpRequest")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/logging/type/http_request.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0x86, 0x35, 0x0c, 0x9d, 0xce, 0x38, 0xd3, 0xaa, 0x32, 0x12, 0x58, 0x20, 0x6e, 0x45, 0x48,
	0xdd, 0x90, 0x2c, 0x78, 0x02, 0x46, 0x02, 0xb5, 0x08, 0x44, 0x95, 0x0a, 0x58, 0x5a, 0x69, 0xe6,
	0x8c, 0x63, 0xe1, 0x89, 0x8d, 0xed, 0x14, 0x95, 0x87, 0xe5, 0x59, 0x38, 0x3e, 0x4e, 0x24, 0x16,
	0x2c, 0xba, 0x89, 0x94, 0xef, 0xff, 0xce, 0xc5, 0x89, 0xd9, 0x7b, 0x65, 0xad, 0x32, 0x50, 0x2a,
	0x6b, 0x9a, 0x5e, 0x95, 0xd6, 0xab, 0x4a, 0x41, 0xef, 0xbc, 0x8d, 0xb6, 0xca, 0x51, 0xe3, 0x74,
	0xa8, 0x8c, 0x55, 0x4a, 0xf7, 0xaa, 0x8a, 0xb7, 0x0e, 0xaa, 0x2e, 0x46, 0x27, 0x3d, 0xfc, 0x1c,
	0x20, 0xc4, 0x92, 0x54, 0xfe, 0x60, 0x6c, 0x33, 0x7a, 0x65, 0xf2, 0x1e, 0x5f, 0xdc, 0xad, 0x37,
	0x3e, 0xaa, 0x00, 0xfe, 0x46, 0xb7, 0xd0, 0xda, 0x7e, 0xa7, 0x55, 0xd5, 0xf4, 0xbd, 0x8d, 0x4d,
	0xd4, 0xb6, 0x0f, 0xb9, 0xff, 0xe9, 0x9f, 0x39, 0x2b, 0xce, 0x71, 0x6c, 0x9d, 0xa7, 0xf2, 0xd7,
	0xec, 0x78, 0x5c, 0x40, 0xee, 0x21, 0x76, 0x76, 0x2b, 0x66, 0x2f, 0x66, 0x67, 0xab, 0xfa, 0x68,
	0xa4, 0x9f, 0x09, 0xf2, 0xe7, 0xac, 0x98, 0xb4, 0xc1, 0x1b, 0x71, 0x8f, 0x1c, 0x36, 0xa2, 0xaf,
	0xde, 0xf0, 0x97, 0x6c, 0x3d, 0x09, 0x41, 0xff, 0x06, 0x31, 0x47, 0x63, 0x5e, 0x4f, 0x45, 0x57,
	0x88, 0xf8, 0x43, 0xb6, 0x08, 0xb8, 0xcc, 0x10, 0xc4, 0x7d, 0x0c, 0x0f, 0xea, 0xf1, 0x8d, 0xbf,
	0x62, 0x38, 0x2c, 0x38, 0xdc, 0x11, 0x72, 0xed, 0x01, 0xd5, 0xae, 0x27, 0x48, 0xc5, 0x4f, 0x19,
	0x1b, 0xf0, 0x6c, 0xb2, 0xc1, 0x83, 0x47, 0xb1, 0xa0, 0xf9, 0xab, 0x44, 0xde, 0x25, 0xc0, 0x9f,
	0xb0, 0x95, 0x87, 0xbd, 0x8d, 0x20, 0xb5, 0x13, 0x87, 0x94, 0x2e, 0x33, 0xb8, 0x70, 0x29, 0x4c,
	0x9f, 0x05, 0xab, 0x31, 0x3c, 0xca, 0x61, 0x06, 0x18, 0x0a, 0x76, 0xe8, 0x61, 0x07, 0x1e, 0xbc,
	0x58, 0x52, 0x34, 0xbd, 0xa6, 0x23, 0xb5, 0x4d, 0xdb, 0x81, 0x34, 0xd6, 0xfe, 0x18, 0x9c, 0x28,
	0x30, 0x5e, 0xd6, 0x05, 0xb1, 0x4f, 0x84, 0x52, 0xe7, 0xac, 0x74, 0x3a, 0x8a, 0x15, 0xe5, 0x4b,
	0x02, 0xe7, 0x3a, 0xf2, 0x8f, 0xec, 0x34, 0x87, 0x37, 0x8d, 0xd1, 0xdb, 0x26, 0xc2, 0x56, 0xfe,
	0xd2, 0xb1, 0x93, 0xd6, 0x6b, 0xfc, 0xb5, 0x32, 0x6f, 0x20, 0x18, 0x55, 0x3d, 0x23, 0xf3, 0xdb,
	0x24, 0x7e, 0x47, 0xef, 0x0b, 0x69, 0x57, 0x64, 0xf1, 0x33, 0x76, 0x92, 0x7b, 0xed, 0xb4, 0x31,
	0xf2, 0xfa, 0x36, 0x42, 0x10, 0x6b, 0xfa, 0x4c, 0xc7, 0xc4, 0x3f, 0x20, 0xde, 0x24, 0xba, 0x79,
	0xc3, 0x1e, 0xb5, 0x76, 0x5f, 0xfe, 0xe7, 0x1a, 0x6d, 0x4e, 0xfe, 0xf9, 0xf1, 0x97, 0xe9, 0x36,
	0x5c, 0xce, 0xae, 0x17, 0x74, 0x2d, 0xde, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xd7, 0xa9,
	0xed, 0xbf, 0x02, 0x00, 0x00,
}