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
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x51, 0x6f, 0xd3, 0x3c,
	0x14, 0x86, 0x95, 0xaf, 0xdf, 0xba, 0xd6, 0x6d, 0xa7, 0xc9, 0x48, 0x60, 0x81, 0x80, 0x32, 0x84,
	0xd4, 0x1b, 0x92, 0x0b, 0x7e, 0x01, 0x95, 0x40, 0x2b, 0x02, 0x31, 0x65, 0x02, 0x2e, 0x23, 0x2f,
	0x3d, 0x75, 0x2c, 0x5c, 0xdb, 0xd8, 0x27, 0x43, 0xdb, 0x8f, 0xe5, 0xb7, 0x20, 0x1f, 0x27, 0x12,
	0x17, 0x5c, 0x70, 0xe9, 0xe7, 0x7d, 0xce, 0x6b, 0xc7, 0x31, 0x7b, 0xa7, 0x9c, 0x53, 0x06, 0x4a,
	0xe5, 0x8c, 0xb4, 0xaa, 0x74, 0x41, 0x55, 0x0a, 0xac, 0x0f, 0x0e, 0x5d, 0x95, 0x23, 0xe9, 0x75,
	0xac, 0x8c, 0x53, 0x4a, 0x5b, 0x55, 0xe1, 0x9d, 0x87, 0xaa, 0x43, 0xf4, 0x4d, 0x80, 0x1f, 0x3d,
	0x44, 0x2c, 0x49, 0xe5, 0x0f, 0x86, 0x9a, 0xc1, 0x2b, 0x93, 0xf7, 0x78, 0xf7, 0x6f, 0xdd, 0xd2,
	0xeb, 0x2a, 0x42, 0xb8, 0xd5, 0x2d, 0xb4, 0xce, 0x1e, 0xb4, 0xaa, 0xa4, 0xb5, 0x0e, 0x25, 0x6a,
	0x67, 0x63, 0xee, 0xbf, 0xf8, 0x35, 0x61, 0x8b, 0x4b, 0x44, 0x5f, 0xe7, 0x5d, 0xf9, 0x2b, 0x76,
	0x36, 0x1c, 0xa0, 0x39, 0x02, 0x76, 0x6e, 0x2f, 0x8a, 0x75, 0xb1, 0x99, 0xd7, 0xab, 0x81, 0x7e,
	0x22, 0xc8, 0x9f, 0xb3, 0xc5, 0xa8, 0xf5, 0xc1, 0x88, 0xff, 0xc8, 0x61, 0x03, 0xfa, 0x12, 0x0c,
	0x7f, 0xc1, 0x96, 0xa3, 0x10, 0xf5, 0x3d, 0x88, 0xc9, 0xba, 0xd8, 0x4c, 0xea, 0x71, 0xe8, 0x5a,
	0xdf, 0x03, 0x7f, 0xc8, 0xa6, 0x11, 0x25, 0xf6, 0x51, 0xfc, 0xbf, 0x2e, 0x36, 0x27, 0xf5, 0xb0,
	0xe2, 0x2f, 0xd9, 0x2a, 0x40, 0xf4, 0xce, 0x46, 0xc8, 0xb3, 0x27, 0x34, 0xbb, 0x1c, 0x21, 0x0d,
	0x3f, 0x65, 0xac, 0x8f, 0x10, 0x1a, 0xa9, 0xc0, 0xa2, 0x98, 0xd2, 0xfe, 0xf3, 0x44, 0xde, 0x26,
	0xc0, 0x9f, 0xb0, 0x79, 0x80, 0xa3, 0x43, 0x68, 0xb4, 0x17, 0xa7, 0x94, 0xce, 0x32, 0xd8, 0xf9,
	0x14, 0xa6, 0x6b, 0x81, 0x90, 0xc2, 0x55, 0x0e, 0x33, 0xd8, 0x79, 0x2e, 0xd8, 0x69, 0x80, 0x03,
	0x04, 0x08, 0x62, 0x46, 0xd1, 0xb8, 0x4c, 0x9f, 0xd4, 0xca, 0xb6, 0x83, 0xc6, 0x38, 0xf7, 0xbd,
	0xf7, 0x62, 0xb1, 0x2e, 0x36, 0xb3, 0x7a, 0x41, 0xec, 0x23, 0xa1, 0xd4, 0x9c, 0x95, 0x4e, 0xa3,
	0x98, 0x53, 0x3e, 0x23, 0x70, 0xa9, 0x91, 0x7f, 0x60, 0x17, 0x39, 0xbc, 0x95, 0x46, 0xef, 0x25,
	0xc2, 0xbe, 0xf9, 0xa9, 0xb1, 0x6b, 0x5c, 0xd0, 0x4a, 0xdb, 0x26, 0x9f, 0x40, 0x30, 0x9a, 0x7a,
	0x46, 0xe6, 0xd7, 0x51, 0xfc, 0xa6, 0xb1, 0xfb, 0x4c, 0xda, 0x35, 0x59, 0x7c, 0xc3, 0xce, 0x73,
	0xd7, 0x41, 0x1b, 0xd3, 0xdc, 0xdc, 0x21, 0x44, 0xb1, 0xa4, 0x6b, 0x3a, 0x23, 0xfe, 0x5e, 0x1b,
	0xb3, 0x4d, 0x74, 0xfb, 0x9a, 0x3d, 0x6a, 0xdd, 0xb1, 0xfc, 0xcb, 0x33, 0xda, 0x9e, 0xff, 0xf1,
	0xe3, 0xaf, 0xd2, 0x6b, 0xb8, 0x2a, 0x6e, 0xa6, 0xf4, 0x2c, 0xde, 0xfc, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xdf, 0xd7, 0xa9, 0xed, 0xbf, 0x02, 0x00, 0x00,
}
