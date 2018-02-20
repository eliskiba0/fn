// Code generated by protoc-gen-go. DO NOT EDIT.
// source: poolmanager.proto

/*
Package poolmanager is a generated protocol buffer package.

It is generated from these files:
	poolmanager.proto

It has these top-level messages:
	Runner
	LBGroupId
	LBGroupMembership
	CapacitySnapshot
	CapacitySnapshotList
*/
package poolmanager

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Runner struct {
	Address   string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	ClientKey []byte `protobuf:"bytes,2,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
}

func (m *Runner) Reset()                    { *m = Runner{} }
func (m *Runner) String() string            { return proto.CompactTextString(m) }
func (*Runner) ProtoMessage()               {}
func (*Runner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Runner) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Runner) GetClientKey() []byte {
	if m != nil {
		return m.ClientKey
	}
	return nil
}

type LBGroupId struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *LBGroupId) Reset()                    { *m = LBGroupId{} }
func (m *LBGroupId) String() string            { return proto.CompactTextString(m) }
func (*LBGroupId) ProtoMessage()               {}
func (*LBGroupId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LBGroupId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type LBGroupMembership struct {
	GroupId *LBGroupId `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	Runners []*Runner  `protobuf:"bytes,2,rep,name=runners" json:"runners,omitempty"`
}

func (m *LBGroupMembership) Reset()                    { *m = LBGroupMembership{} }
func (m *LBGroupMembership) String() string            { return proto.CompactTextString(m) }
func (*LBGroupMembership) ProtoMessage()               {}
func (*LBGroupMembership) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LBGroupMembership) GetGroupId() *LBGroupId {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *LBGroupMembership) GetRunners() []*Runner {
	if m != nil {
		return m.Runners
	}
	return nil
}

type CapacitySnapshot struct {
	GroupId    *LBGroupId `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	MemMbTotal int32      `protobuf:"varint,2,opt,name=mem_mb_total,json=memMbTotal" json:"mem_mb_total,omitempty"`
}

func (m *CapacitySnapshot) Reset()                    { *m = CapacitySnapshot{} }
func (m *CapacitySnapshot) String() string            { return proto.CompactTextString(m) }
func (*CapacitySnapshot) ProtoMessage()               {}
func (*CapacitySnapshot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CapacitySnapshot) GetGroupId() *LBGroupId {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *CapacitySnapshot) GetMemMbTotal() int32 {
	if m != nil {
		return m.MemMbTotal
	}
	return 0
}

type CapacitySnapshotList struct {
	Ts        *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=ts" json:"ts,omitempty"`
	LbId      string                     `protobuf:"bytes,2,opt,name=lb_id,json=lbId" json:"lb_id,omitempty"`
	Snapshots []*CapacitySnapshot        `protobuf:"bytes,3,rep,name=snapshots" json:"snapshots,omitempty"`
}

func (m *CapacitySnapshotList) Reset()                    { *m = CapacitySnapshotList{} }
func (m *CapacitySnapshotList) String() string            { return proto.CompactTextString(m) }
func (*CapacitySnapshotList) ProtoMessage()               {}
func (*CapacitySnapshotList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CapacitySnapshotList) GetTs() *google_protobuf.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *CapacitySnapshotList) GetLbId() string {
	if m != nil {
		return m.LbId
	}
	return ""
}

func (m *CapacitySnapshotList) GetSnapshots() []*CapacitySnapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

func init() {
	proto.RegisterType((*Runner)(nil), "Runner")
	proto.RegisterType((*LBGroupId)(nil), "LBGroupId")
	proto.RegisterType((*LBGroupMembership)(nil), "LBGroupMembership")
	proto.RegisterType((*CapacitySnapshot)(nil), "CapacitySnapshot")
	proto.RegisterType((*CapacitySnapshotList)(nil), "CapacitySnapshotList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodePoolScaler service

type NodePoolScalerClient interface {
	AdvertiseCapacity(ctx context.Context, in *CapacitySnapshotList, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type nodePoolScalerClient struct {
	cc *grpc.ClientConn
}

func NewNodePoolScalerClient(cc *grpc.ClientConn) NodePoolScalerClient {
	return &nodePoolScalerClient{cc}
}

func (c *nodePoolScalerClient) AdvertiseCapacity(ctx context.Context, in *CapacitySnapshotList, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/NodePoolScaler/AdvertiseCapacity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodePoolScaler service

type NodePoolScalerServer interface {
	AdvertiseCapacity(context.Context, *CapacitySnapshotList) (*google_protobuf1.Empty, error)
}

func RegisterNodePoolScalerServer(s *grpc.Server, srv NodePoolScalerServer) {
	s.RegisterService(&_NodePoolScaler_serviceDesc, srv)
}

func _NodePoolScaler_AdvertiseCapacity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapacitySnapshotList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodePoolScalerServer).AdvertiseCapacity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodePoolScaler/AdvertiseCapacity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodePoolScalerServer).AdvertiseCapacity(ctx, req.(*CapacitySnapshotList))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodePoolScaler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NodePoolScaler",
	HandlerType: (*NodePoolScalerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdvertiseCapacity",
			Handler:    _NodePoolScaler_AdvertiseCapacity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "poolmanager.proto",
}

// Client API for RunnerManager service

type RunnerManagerClient interface {
	GetLBGroup(ctx context.Context, in *LBGroupId, opts ...grpc.CallOption) (*LBGroupMembership, error)
}

type runnerManagerClient struct {
	cc *grpc.ClientConn
}

func NewRunnerManagerClient(cc *grpc.ClientConn) RunnerManagerClient {
	return &runnerManagerClient{cc}
}

func (c *runnerManagerClient) GetLBGroup(ctx context.Context, in *LBGroupId, opts ...grpc.CallOption) (*LBGroupMembership, error) {
	out := new(LBGroupMembership)
	err := grpc.Invoke(ctx, "/RunnerManager/GetLBGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RunnerManager service

type RunnerManagerServer interface {
	GetLBGroup(context.Context, *LBGroupId) (*LBGroupMembership, error)
}

func RegisterRunnerManagerServer(s *grpc.Server, srv RunnerManagerServer) {
	s.RegisterService(&_RunnerManager_serviceDesc, srv)
}

func _RunnerManager_GetLBGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LBGroupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerManagerServer).GetLBGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RunnerManager/GetLBGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerManagerServer).GetLBGroup(ctx, req.(*LBGroupId))
	}
	return interceptor(ctx, in, info, handler)
}

var _RunnerManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RunnerManager",
	HandlerType: (*RunnerManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLBGroup",
			Handler:    _RunnerManager_GetLBGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "poolmanager.proto",
}

func init() { proto.RegisterFile("poolmanager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x5f, 0x6b, 0x14, 0x31,
	0x14, 0xc5, 0xbb, 0x53, 0xdb, 0x75, 0x6f, 0x6b, 0x71, 0xe3, 0x1f, 0x86, 0x2d, 0xe2, 0x38, 0x20,
	0x2c, 0x3e, 0x64, 0x61, 0xfc, 0x00, 0x52, 0xff, 0x50, 0x8a, 0x5d, 0x91, 0xb4, 0xf8, 0x22, 0x32,
	0x64, 0x36, 0xd7, 0x69, 0x30, 0x99, 0x84, 0x24, 0x2b, 0xcc, 0x37, 0xf0, 0x63, 0xcb, 0xcc, 0x64,
	0x54, 0xb6, 0x3e, 0xf4, 0x31, 0x27, 0x67, 0x7e, 0x73, 0xee, 0xc9, 0x85, 0xb9, 0x35, 0x46, 0x69,
	0xde, 0xf0, 0x1a, 0x1d, 0xb5, 0xce, 0x04, 0xb3, 0x78, 0x5e, 0x1b, 0x53, 0x2b, 0x5c, 0xf5, 0xa7,
	0x6a, 0xfb, 0x7d, 0x15, 0xa4, 0x46, 0x1f, 0xb8, 0xb6, 0xd1, 0x70, 0xba, 0x6b, 0x40, 0x6d, 0x43,
	0x3b, 0x5c, 0xe6, 0x67, 0x70, 0xc8, 0xb6, 0x4d, 0x83, 0x8e, 0xa4, 0x30, 0xe5, 0x42, 0x38, 0xf4,
	0x3e, 0x9d, 0x64, 0x93, 0xe5, 0x8c, 0x8d, 0x47, 0xf2, 0x0c, 0x60, 0xa3, 0x24, 0x36, 0xa1, 0xfc,
	0x81, 0x6d, 0x9a, 0x64, 0x93, 0xe5, 0x31, 0x9b, 0x0d, 0xca, 0x47, 0x6c, 0xf3, 0x53, 0x98, 0x5d,
	0xbe, 0x3d, 0x77, 0x66, 0x6b, 0x2f, 0x04, 0x39, 0x81, 0x44, 0x8a, 0x08, 0x48, 0xa4, 0xc8, 0xbf,
	0xc1, 0x3c, 0x5e, 0xae, 0x51, 0x57, 0xe8, 0xfc, 0x8d, 0xb4, 0xe4, 0x25, 0xdc, 0xaf, 0x3b, 0xa9,
	0x8c, 0xd6, 0xa3, 0x02, 0xe8, 0x1f, 0x04, 0x9b, 0xd6, 0x91, 0xf5, 0x02, 0xa6, 0xae, 0xcf, 0xe6,
	0xd3, 0x24, 0xdb, 0x5f, 0x1e, 0x15, 0x53, 0x3a, 0x64, 0x65, 0xa3, 0x9e, 0x7f, 0x85, 0x87, 0xef,
	0xb8, 0xe5, 0x1b, 0x19, 0xda, 0xab, 0x86, 0x5b, 0x7f, 0x63, 0xc2, 0x5d, 0xe9, 0x19, 0x1c, 0x6b,
	0xd4, 0xa5, 0xae, 0xca, 0x60, 0x02, 0x57, 0xfd, 0x5c, 0x07, 0x0c, 0x34, 0xea, 0x75, 0x75, 0xdd,
	0x29, 0xf9, 0xaf, 0x09, 0x3c, 0xde, 0xa5, 0x5f, 0x4a, 0x1f, 0xc8, 0x2b, 0x48, 0x82, 0x8f, 0xec,
	0x05, 0x1d, 0xea, 0xa5, 0x63, 0xbd, 0xf4, 0x7a, 0xec, 0x9f, 0x25, 0xc1, 0x93, 0x47, 0x70, 0xa0,
	0xaa, 0x2e, 0x4a, 0xd2, 0x77, 0x72, 0x4f, 0x55, 0x17, 0x82, 0xac, 0x60, 0xe6, 0x23, 0xd0, 0xa7,
	0xfb, 0xfd, 0x6c, 0x73, 0xba, 0xfb, 0x2b, 0xf6, 0xd7, 0x53, 0x7c, 0x81, 0x93, 0x4f, 0x46, 0xe0,
	0x67, 0x63, 0xd4, 0xd5, 0x86, 0x2b, 0x74, 0xe4, 0x3d, 0xcc, 0xcf, 0xc4, 0x4f, 0x74, 0x41, 0x7a,
	0x1c, 0xbf, 0x24, 0x4f, 0xe8, 0xff, 0xf2, 0x2e, 0x9e, 0xde, 0xca, 0xf8, 0xa1, 0x5b, 0x81, 0x7c,
	0xaf, 0x78, 0x03, 0x0f, 0x86, 0x4a, 0xd7, 0xc3, 0x4e, 0x11, 0x0a, 0x70, 0x8e, 0x21, 0xd6, 0x45,
	0xfe, 0x29, 0x6e, 0x41, 0xe8, 0xad, 0x87, 0xcc, 0xf7, 0xaa, 0xc3, 0x1e, 0xf9, 0xfa, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1f, 0x74, 0x50, 0x3b, 0x99, 0x02, 0x00, 0x00,
}
