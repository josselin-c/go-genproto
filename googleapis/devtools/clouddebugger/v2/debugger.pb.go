// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/debugger.proto
// DO NOT EDIT!

package google_devtools_clouddebugger_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf3 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request to set a breakpoint
type SetBreakpointRequest struct {
	// ID of the debuggee where the breakpoint is to be set.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId" json:"debuggee_id,omitempty"`
	// Breakpoint specification to set.
	// The field 'location' of the breakpoint must be set.
	Breakpoint *Breakpoint `protobuf:"bytes,2,opt,name=breakpoint" json:"breakpoint,omitempty"`
	// The client version making the call.
	// Following: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion string `protobuf:"bytes,4,opt,name=client_version,json=clientVersion" json:"client_version,omitempty"`
}

func (m *SetBreakpointRequest) Reset()                    { *m = SetBreakpointRequest{} }
func (m *SetBreakpointRequest) String() string            { return proto.CompactTextString(m) }
func (*SetBreakpointRequest) ProtoMessage()               {}
func (*SetBreakpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *SetBreakpointRequest) GetBreakpoint() *Breakpoint {
	if m != nil {
		return m.Breakpoint
	}
	return nil
}

// Response for setting a breakpoint.
type SetBreakpointResponse struct {
	// Breakpoint resource.
	// The field `id` is guaranteed to be set (in addition to the echoed fileds).
	Breakpoint *Breakpoint `protobuf:"bytes,1,opt,name=breakpoint" json:"breakpoint,omitempty"`
}

func (m *SetBreakpointResponse) Reset()                    { *m = SetBreakpointResponse{} }
func (m *SetBreakpointResponse) String() string            { return proto.CompactTextString(m) }
func (*SetBreakpointResponse) ProtoMessage()               {}
func (*SetBreakpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *SetBreakpointResponse) GetBreakpoint() *Breakpoint {
	if m != nil {
		return m.Breakpoint
	}
	return nil
}

// Request to get breakpoint information.
type GetBreakpointRequest struct {
	// ID of the debuggee whose breakpoint to get.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId" json:"debuggee_id,omitempty"`
	// ID of the breakpoint to get.
	BreakpointId string `protobuf:"bytes,2,opt,name=breakpoint_id,json=breakpointId" json:"breakpoint_id,omitempty"`
	// The client version making the call.
	// Following: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion string `protobuf:"bytes,4,opt,name=client_version,json=clientVersion" json:"client_version,omitempty"`
}

func (m *GetBreakpointRequest) Reset()                    { *m = GetBreakpointRequest{} }
func (m *GetBreakpointRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBreakpointRequest) ProtoMessage()               {}
func (*GetBreakpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

// Response for getting breakpoint information.
type GetBreakpointResponse struct {
	// Complete breakpoint state.
	// The fields `id` and `location` are guaranteed to be set.
	Breakpoint *Breakpoint `protobuf:"bytes,1,opt,name=breakpoint" json:"breakpoint,omitempty"`
}

func (m *GetBreakpointResponse) Reset()                    { *m = GetBreakpointResponse{} }
func (m *GetBreakpointResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBreakpointResponse) ProtoMessage()               {}
func (*GetBreakpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *GetBreakpointResponse) GetBreakpoint() *Breakpoint {
	if m != nil {
		return m.Breakpoint
	}
	return nil
}

// Request to delete a breakpoint.
type DeleteBreakpointRequest struct {
	// ID of the debuggee whose breakpoint to delete.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId" json:"debuggee_id,omitempty"`
	// ID of the breakpoint to delete.
	BreakpointId string `protobuf:"bytes,2,opt,name=breakpoint_id,json=breakpointId" json:"breakpoint_id,omitempty"`
	// The client version making the call.
	// Following: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion string `protobuf:"bytes,3,opt,name=client_version,json=clientVersion" json:"client_version,omitempty"`
}

func (m *DeleteBreakpointRequest) Reset()                    { *m = DeleteBreakpointRequest{} }
func (m *DeleteBreakpointRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteBreakpointRequest) ProtoMessage()               {}
func (*DeleteBreakpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

// Request to list breakpoints.
type ListBreakpointsRequest struct {
	// ID of the debuggee whose breakpoints to list.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId" json:"debuggee_id,omitempty"`
	// When set to `true`, the response includes the list of breakpoints set by
	// any user. Otherwise, it includes only breakpoints set by the caller.
	IncludeAllUsers bool `protobuf:"varint,2,opt,name=include_all_users,json=includeAllUsers" json:"include_all_users,omitempty"`
	// When set to `true`, the response includes active and inactive
	// breakpoints. Otherwise, it includes only active breakpoints.
	IncludeInactive bool `protobuf:"varint,3,opt,name=include_inactive,json=includeInactive" json:"include_inactive,omitempty"`
	// When set, the response includes only breakpoints with the specified action.
	Action *ListBreakpointsRequest_BreakpointActionValue `protobuf:"bytes,4,opt,name=action" json:"action,omitempty"`
	// When set to `true`, the response breakpoints are stripped of the
	// results fields: `stack_frames`, `evaluated_expressions` and
	// `variable_table`.
	StripResults bool `protobuf:"varint,5,opt,name=strip_results,json=stripResults" json:"strip_results,omitempty"`
	// A wait token that, if specified, blocks the call until the breakpoints
	// list has changed, or a server selected timeout has expired.  The value
	// should be set from the last response. The error code
	// `google.rpc.Code.ABORTED` (RPC) is returned on wait timeout, which
	// should be called again with the same `wait_token`.
	WaitToken string `protobuf:"bytes,6,opt,name=wait_token,json=waitToken" json:"wait_token,omitempty"`
	// The client version making the call.
	// Following: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion string `protobuf:"bytes,8,opt,name=client_version,json=clientVersion" json:"client_version,omitempty"`
}

func (m *ListBreakpointsRequest) Reset()                    { *m = ListBreakpointsRequest{} }
func (m *ListBreakpointsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListBreakpointsRequest) ProtoMessage()               {}
func (*ListBreakpointsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *ListBreakpointsRequest) GetAction() *ListBreakpointsRequest_BreakpointActionValue {
	if m != nil {
		return m.Action
	}
	return nil
}

// Wrapper message for `Breakpoint.Action`. Defines a filter on the action
// field of breakpoints.
type ListBreakpointsRequest_BreakpointActionValue struct {
	// Only breakpoints with the specified action will pass the filter.
	Value Breakpoint_Action `protobuf:"varint,1,opt,name=value,enum=google.devtools.clouddebugger.v2.Breakpoint_Action" json:"value,omitempty"`
}

func (m *ListBreakpointsRequest_BreakpointActionValue) Reset() {
	*m = ListBreakpointsRequest_BreakpointActionValue{}
}
func (m *ListBreakpointsRequest_BreakpointActionValue) String() string {
	return proto.CompactTextString(m)
}
func (*ListBreakpointsRequest_BreakpointActionValue) ProtoMessage() {}
func (*ListBreakpointsRequest_BreakpointActionValue) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{5, 0}
}

// Response for listing breakpoints.
type ListBreakpointsResponse struct {
	// List of all breakpoints with complete state.
	// The fields `id` and `location` are guaranteed to be set on each breakpoint.
	Breakpoints []*Breakpoint `protobuf:"bytes,1,rep,name=breakpoints" json:"breakpoints,omitempty"`
	// A wait token that can be used in the next call to `list` (REST) or
	// `ListBreakpoints` (RPC) to block until the list of breakpoints has changes.
	NextWaitToken string `protobuf:"bytes,2,opt,name=next_wait_token,json=nextWaitToken" json:"next_wait_token,omitempty"`
}

func (m *ListBreakpointsResponse) Reset()                    { *m = ListBreakpointsResponse{} }
func (m *ListBreakpointsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListBreakpointsResponse) ProtoMessage()               {}
func (*ListBreakpointsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *ListBreakpointsResponse) GetBreakpoints() []*Breakpoint {
	if m != nil {
		return m.Breakpoints
	}
	return nil
}

// Request to list debuggees.
type ListDebuggeesRequest struct {
	// Project number of a Google Cloud project whose debuggees to list.
	Project string `protobuf:"bytes,2,opt,name=project" json:"project,omitempty"`
	// When set to `true`, the result includes all debuggees. Otherwise, the
	// result includes only debuggees that are active.
	IncludeInactive bool `protobuf:"varint,3,opt,name=include_inactive,json=includeInactive" json:"include_inactive,omitempty"`
	// The client version making the call.
	// Following: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion string `protobuf:"bytes,4,opt,name=client_version,json=clientVersion" json:"client_version,omitempty"`
}

func (m *ListDebuggeesRequest) Reset()                    { *m = ListDebuggeesRequest{} }
func (m *ListDebuggeesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDebuggeesRequest) ProtoMessage()               {}
func (*ListDebuggeesRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

// Response for listing debuggees.
type ListDebuggeesResponse struct {
	// List of debuggees accessible to the calling user.
	// Note that the `description` field is the only human readable field
	// that should be displayed to the user.
	// The fields `debuggee.id` and  `description` fields are guaranteed to be
	// set on each debuggee.
	Debuggees []*Debuggee `protobuf:"bytes,1,rep,name=debuggees" json:"debuggees,omitempty"`
}

func (m *ListDebuggeesResponse) Reset()                    { *m = ListDebuggeesResponse{} }
func (m *ListDebuggeesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDebuggeesResponse) ProtoMessage()               {}
func (*ListDebuggeesResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *ListDebuggeesResponse) GetDebuggees() []*Debuggee {
	if m != nil {
		return m.Debuggees
	}
	return nil
}

func init() {
	proto.RegisterType((*SetBreakpointRequest)(nil), "google.devtools.clouddebugger.v2.SetBreakpointRequest")
	proto.RegisterType((*SetBreakpointResponse)(nil), "google.devtools.clouddebugger.v2.SetBreakpointResponse")
	proto.RegisterType((*GetBreakpointRequest)(nil), "google.devtools.clouddebugger.v2.GetBreakpointRequest")
	proto.RegisterType((*GetBreakpointResponse)(nil), "google.devtools.clouddebugger.v2.GetBreakpointResponse")
	proto.RegisterType((*DeleteBreakpointRequest)(nil), "google.devtools.clouddebugger.v2.DeleteBreakpointRequest")
	proto.RegisterType((*ListBreakpointsRequest)(nil), "google.devtools.clouddebugger.v2.ListBreakpointsRequest")
	proto.RegisterType((*ListBreakpointsRequest_BreakpointActionValue)(nil), "google.devtools.clouddebugger.v2.ListBreakpointsRequest.BreakpointActionValue")
	proto.RegisterType((*ListBreakpointsResponse)(nil), "google.devtools.clouddebugger.v2.ListBreakpointsResponse")
	proto.RegisterType((*ListDebuggeesRequest)(nil), "google.devtools.clouddebugger.v2.ListDebuggeesRequest")
	proto.RegisterType((*ListDebuggeesResponse)(nil), "google.devtools.clouddebugger.v2.ListDebuggeesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Debugger2 service

type Debugger2Client interface {
	// Sets the breakpoint to the debuggee.
	SetBreakpoint(ctx context.Context, in *SetBreakpointRequest, opts ...grpc.CallOption) (*SetBreakpointResponse, error)
	// Gets breakpoint information.
	GetBreakpoint(ctx context.Context, in *GetBreakpointRequest, opts ...grpc.CallOption) (*GetBreakpointResponse, error)
	// Deletes the breakpoint from the debuggee.
	DeleteBreakpoint(ctx context.Context, in *DeleteBreakpointRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error)
	// Lists all breakpoints for the debuggee.
	ListBreakpoints(ctx context.Context, in *ListBreakpointsRequest, opts ...grpc.CallOption) (*ListBreakpointsResponse, error)
	// Lists all the debuggees that the user can set breakpoints to.
	ListDebuggees(ctx context.Context, in *ListDebuggeesRequest, opts ...grpc.CallOption) (*ListDebuggeesResponse, error)
}

type debugger2Client struct {
	cc *grpc.ClientConn
}

func NewDebugger2Client(cc *grpc.ClientConn) Debugger2Client {
	return &debugger2Client{cc}
}

func (c *debugger2Client) SetBreakpoint(ctx context.Context, in *SetBreakpointRequest, opts ...grpc.CallOption) (*SetBreakpointResponse, error) {
	out := new(SetBreakpointResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/SetBreakpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugger2Client) GetBreakpoint(ctx context.Context, in *GetBreakpointRequest, opts ...grpc.CallOption) (*GetBreakpointResponse, error) {
	out := new(GetBreakpointResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/GetBreakpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugger2Client) DeleteBreakpoint(ctx context.Context, in *DeleteBreakpointRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error) {
	out := new(google_protobuf3.Empty)
	err := grpc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/DeleteBreakpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugger2Client) ListBreakpoints(ctx context.Context, in *ListBreakpointsRequest, opts ...grpc.CallOption) (*ListBreakpointsResponse, error) {
	out := new(ListBreakpointsResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/ListBreakpoints", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugger2Client) ListDebuggees(ctx context.Context, in *ListDebuggeesRequest, opts ...grpc.CallOption) (*ListDebuggeesResponse, error) {
	out := new(ListDebuggeesResponse)
	err := grpc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/ListDebuggees", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Debugger2 service

type Debugger2Server interface {
	// Sets the breakpoint to the debuggee.
	SetBreakpoint(context.Context, *SetBreakpointRequest) (*SetBreakpointResponse, error)
	// Gets breakpoint information.
	GetBreakpoint(context.Context, *GetBreakpointRequest) (*GetBreakpointResponse, error)
	// Deletes the breakpoint from the debuggee.
	DeleteBreakpoint(context.Context, *DeleteBreakpointRequest) (*google_protobuf3.Empty, error)
	// Lists all breakpoints for the debuggee.
	ListBreakpoints(context.Context, *ListBreakpointsRequest) (*ListBreakpointsResponse, error)
	// Lists all the debuggees that the user can set breakpoints to.
	ListDebuggees(context.Context, *ListDebuggeesRequest) (*ListDebuggeesResponse, error)
}

func RegisterDebugger2Server(s *grpc.Server, srv Debugger2Server) {
	s.RegisterService(&_Debugger2_serviceDesc, srv)
}

func _Debugger2_SetBreakpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBreakpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).SetBreakpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/SetBreakpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).SetBreakpoint(ctx, req.(*SetBreakpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debugger2_GetBreakpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBreakpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).GetBreakpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/GetBreakpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).GetBreakpoint(ctx, req.(*GetBreakpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debugger2_DeleteBreakpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBreakpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).DeleteBreakpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/DeleteBreakpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).DeleteBreakpoint(ctx, req.(*DeleteBreakpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debugger2_ListBreakpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBreakpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).ListBreakpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/ListBreakpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).ListBreakpoints(ctx, req.(*ListBreakpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debugger2_ListDebuggees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDebuggeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).ListDebuggees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/ListDebuggees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).ListDebuggees(ctx, req.(*ListDebuggeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Debugger2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouddebugger.v2.Debugger2",
	HandlerType: (*Debugger2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetBreakpoint",
			Handler:    _Debugger2_SetBreakpoint_Handler,
		},
		{
			MethodName: "GetBreakpoint",
			Handler:    _Debugger2_GetBreakpoint_Handler,
		},
		{
			MethodName: "DeleteBreakpoint",
			Handler:    _Debugger2_DeleteBreakpoint_Handler,
		},
		{
			MethodName: "ListBreakpoints",
			Handler:    _Debugger2_ListBreakpoints_Handler,
		},
		{
			MethodName: "ListDebuggees",
			Handler:    _Debugger2_ListDebuggees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/devtools/clouddebugger/v2/debugger.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 781 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x55, 0xdd, 0x6a, 0x13, 0x4f,
	0x1c, 0x65, 0xda, 0x7f, 0x3f, 0xf2, 0x4b, 0xd3, 0xf6, 0x3f, 0xf4, 0x63, 0x09, 0xa8, 0x61, 0xfd,
	0xa0, 0x16, 0xd9, 0x91, 0xad, 0xd8, 0xd6, 0x2b, 0x1b, 0x2a, 0x69, 0xa4, 0x94, 0x12, 0xb5, 0x5e,
	0x86, 0xcd, 0x66, 0xba, 0x8e, 0xdd, 0xee, 0xac, 0x3b, 0xb3, 0xd1, 0x52, 0x7a, 0x53, 0xc1, 0x7b,
	0xf1, 0x05, 0x7c, 0x00, 0xc1, 0x27, 0x10, 0xbc, 0x13, 0xbc, 0xf5, 0x15, 0x7c, 0x10, 0xd9, 0xaf,
	0x66, 0x13, 0x17, 0x93, 0x4d, 0xa1, 0x37, 0x61, 0x73, 0x66, 0x7e, 0x67, 0xce, 0x39, 0xf3, 0x9b,
	0x19, 0xd8, 0xb3, 0x38, 0xb7, 0x6c, 0xaa, 0x59, 0xdc, 0x36, 0x1c, 0x4b, 0xe3, 0x9e, 0x45, 0x2c,
	0xea, 0xb8, 0x1e, 0x97, 0x9c, 0x44, 0x43, 0x86, 0xcb, 0x04, 0x69, 0xd3, 0x8e, 0xe4, 0xdc, 0x16,
	0xc4, 0xb4, 0xb9, 0xdf, 0x6e, 0xd3, 0x96, 0x6f, 0x59, 0xd4, 0x23, 0x1d, 0x9d, 0x24, 0xdf, 0x5a,
	0x58, 0x83, 0x2b, 0x31, 0x5f, 0x52, 0xa0, 0xf5, 0x14, 0x68, 0x1d, 0xbd, 0x5c, 0x1f, 0x6e, 0x45,
	0xc3, 0x65, 0x44, 0x50, 0xaf, 0xc3, 0x4c, 0x6a, 0x72, 0xe7, 0x90, 0x59, 0xc4, 0x70, 0x1c, 0x2e,
	0x0d, 0xc9, 0xb8, 0x23, 0xa2, 0xc5, 0xca, 0x4f, 0x2f, 0x2d, 0xde, 0x90, 0x46, 0xcc, 0xb5, 0x66,
	0x31, 0xf9, 0xca, 0x6f, 0x69, 0x26, 0x3f, 0x26, 0x11, 0x1f, 0x09, 0x07, 0x5a, 0xfe, 0x21, 0x71,
	0xe5, 0x89, 0x4b, 0x05, 0xa1, 0xc7, 0xae, 0x3c, 0x89, 0x7e, 0xa3, 0x22, 0xf5, 0x0b, 0x82, 0x85,
	0x67, 0x54, 0x56, 0x3d, 0x6a, 0x1c, 0xb9, 0x9c, 0x39, 0xb2, 0x41, 0xdf, 0xf8, 0x54, 0x48, 0x7c,
	0x03, 0x8a, 0xf1, 0x3a, 0xb4, 0xc9, 0xda, 0x0a, 0xaa, 0xa0, 0x95, 0x42, 0x03, 0x12, 0xa8, 0xde,
	0xc6, 0xbb, 0x00, 0xad, 0x8b, 0x2a, 0x65, 0xac, 0x82, 0x56, 0x8a, 0xfa, 0x3d, 0x6d, 0x50, 0x78,
	0x5a, 0x6a, 0xa5, 0x54, 0x3d, 0xbe, 0x0d, 0xb3, 0xa6, 0xcd, 0xa8, 0x23, 0x9b, 0x1d, 0xea, 0x09,
	0xc6, 0x1d, 0xe5, 0xbf, 0x70, 0xc5, 0x52, 0x84, 0x1e, 0x44, 0xa0, 0x4a, 0x61, 0xb1, 0x4f, 0xad,
	0x70, 0xb9, 0x23, 0x68, 0x9f, 0x1a, 0x74, 0x39, 0x35, 0xea, 0x7b, 0x04, 0x0b, 0xb5, 0x91, 0x52,
	0xb9, 0x09, 0xa5, 0x2e, 0x4f, 0x30, 0x65, 0x2c, 0x9c, 0x32, 0xd3, 0x05, 0xeb, 0xed, 0x1c, 0x66,
	0x6b, 0x57, 0x60, 0xf6, 0x03, 0x82, 0xe5, 0x6d, 0x6a, 0x53, 0x49, 0xaf, 0xce, 0xef, 0x78, 0x96,
	0xdf, 0x1f, 0xe3, 0xb0, 0xb4, 0xcb, 0x44, 0xca, 0xb1, 0x18, 0x5a, 0xc7, 0x2a, 0xfc, 0xcf, 0x1c,
	0xd3, 0xf6, 0xdb, 0xb4, 0x69, 0xd8, 0x76, 0xd3, 0x17, 0xd4, 0x13, 0xa1, 0x96, 0xe9, 0xc6, 0x5c,
	0x3c, 0xb0, 0x65, 0xdb, 0x2f, 0x02, 0x18, 0xdf, 0x85, 0xf9, 0x64, 0x2e, 0x73, 0x0c, 0x53, 0xb2,
	0x0e, 0x0d, 0x05, 0x75, 0xa7, 0xd6, 0x63, 0x18, 0x1f, 0xc2, 0x64, 0xf0, 0x15, 0xef, 0x50, 0x51,
	0xdf, 0x1b, 0x9c, 0x72, 0xb6, 0x83, 0x54, 0xf8, 0x5b, 0x21, 0xe1, 0x81, 0x61, 0xfb, 0xb4, 0x11,
	0xb3, 0x07, 0x31, 0x0a, 0xe9, 0x31, 0xb7, 0xe9, 0x51, 0xe1, 0xdb, 0x52, 0x28, 0x13, 0xa1, 0x9e,
	0x99, 0x10, 0x6c, 0x44, 0x18, 0xbe, 0x06, 0xf0, 0xd6, 0x60, 0xb2, 0x29, 0xf9, 0x11, 0x75, 0x94,
	0xc9, 0x30, 0x83, 0x42, 0x80, 0x3c, 0x0f, 0x80, 0x8c, 0x94, 0xa7, 0x33, 0x52, 0x2e, 0xb7, 0x60,
	0x31, 0x53, 0x0b, 0xae, 0xc3, 0x44, 0x27, 0xf8, 0x08, 0xd3, 0x9d, 0xd5, 0xd7, 0xf2, 0x34, 0x94,
	0x16, 0x11, 0x35, 0x22, 0x06, 0xf5, 0x23, 0x82, 0xe5, 0xbf, 0x72, 0x88, 0x9b, 0x77, 0x0f, 0x8a,
	0xdd, 0xe6, 0x10, 0x0a, 0xaa, 0x8c, 0xe7, 0xee, 0xde, 0x34, 0x01, 0xbe, 0x03, 0x73, 0x0e, 0x7d,
	0x27, 0x9b, 0xa9, 0x68, 0xa2, 0x1e, 0x2c, 0x05, 0xf0, 0xcb, 0x24, 0x1e, 0xf5, 0x1c, 0xc1, 0x42,
	0xa0, 0x69, 0x3b, 0x6e, 0x9a, 0x8b, 0xde, 0x52, 0x60, 0xca, 0xf5, 0xf8, 0x6b, 0x6a, 0xca, 0xb8,
	0x30, 0xf9, 0x9b, 0xa7, 0x51, 0x86, 0x3c, 0xd2, 0x06, 0x2c, 0xf6, 0x69, 0x88, 0x53, 0xd9, 0x81,
	0x42, 0xd2, 0xcd, 0x49, 0x26, 0xab, 0x83, 0x33, 0x49, 0x78, 0x1a, 0xdd, 0x62, 0xfd, 0xdb, 0x14,
	0x14, 0x62, 0xdc, 0xd3, 0xf1, 0x4f, 0x04, 0xa5, 0x9e, 0x1b, 0x13, 0x3f, 0x1c, 0x4c, 0x9b, 0xf5,
	0x20, 0x94, 0xd7, 0x73, 0xd7, 0x45, 0xd6, 0xd4, 0x9d, 0xf3, 0x5f, 0xbf, 0x3f, 0x8d, 0x55, 0xd5,
	0x07, 0xe9, 0xc7, 0x96, 0x5c, 0x08, 0x26, 0xa7, 0xa9, 0x93, 0x7d, 0x46, 0x52, 0x5b, 0x4b, 0x04,
	0x95, 0x8f, 0xd2, 0x8f, 0x44, 0x60, 0xa6, 0x96, 0xd7, 0x4c, 0x6d, 0x44, 0x33, 0xb5, 0x7f, 0x99,
	0xc1, 0x8f, 0x73, 0x9b, 0x39, 0xed, 0xb9, 0x27, 0xcf, 0xf0, 0x57, 0x04, 0xf3, 0xfd, 0xd7, 0x2e,
	0xde, 0x1c, 0x66, 0xcf, 0x33, 0xaf, 0xea, 0xf2, 0x52, 0x52, 0x9a, 0x3c, 0xfa, 0xda, 0x93, 0xe0,
	0x9d, 0x4f, 0x14, 0xaf, 0x5e, 0x5e, 0xf1, 0x77, 0x04, 0x73, 0x7d, 0xa7, 0x1a, 0x6f, 0x8c, 0x7a,
	0x21, 0x96, 0x37, 0x47, 0xa8, 0x8c, 0x37, 0x61, 0x23, 0xb4, 0xa4, 0xe3, 0xfb, 0x79, 0x2d, 0xe1,
	0xcf, 0x08, 0x4a, 0x3d, 0x07, 0x70, 0x98, 0x0e, 0xca, 0xba, 0x35, 0x86, 0xe9, 0xa0, 0xcc, 0x93,
	0xae, 0x5e, 0x0f, 0xc5, 0x2b, 0x78, 0x29, 0x5b, 0x7c, 0x75, 0x1d, 0x6e, 0x99, 0xfc, 0x78, 0x20,
	0x7b, 0xb5, 0x94, 0x1c, 0xf2, 0xfd, 0x60, 0xbf, 0xf7, 0x51, 0x6b, 0x32, 0xdc, 0xf8, 0xb5, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x05, 0x32, 0xe5, 0x11, 0x0b, 0x00, 0x00,
}
