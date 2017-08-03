// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opensds.proto

/*
Package opensds is a generated protocol buffer package.

It is generated from these files:
	opensds.proto

It has these top-level messages:
	DockRequest
	DockResponse
*/
package opensds

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// The DockRequest message containing all properties of
// a request to dock service.
type DockRequest struct {
	VolumeId              string `protobuf:"bytes,1,opt,name=volumeId" json:"volumeId,omitempty"`
	VolumeName            string `protobuf:"bytes,2,opt,name=volumeName" json:"volumeName,omitempty"`
	VolumeDescription     string `protobuf:"bytes,3,opt,name=volumeDescription" json:"volumeDescription,omitempty"`
	VolumeSize            int64  `protobuf:"varint,4,opt,name=volumeSize" json:"volumeSize,omitempty"`
	AttachmentId          string `protobuf:"bytes,5,opt,name=attachmentId" json:"attachmentId,omitempty"`
	AttachmentName        string `protobuf:"bytes,6,opt,name=attachmentName" json:"attachmentName,omitempty"`
	AttachmentDescription string `protobuf:"bytes,7,opt,name=attachmentDescription" json:"attachmentDescription,omitempty"`
	DoLocalAttach         bool   `protobuf:"varint,8,opt,name=doLocalAttach" json:"doLocalAttach,omitempty"`
	MultiPath             bool   `protobuf:"varint,9,opt,name=multiPath" json:"multiPath,omitempty"`
	HostInfo              string `protobuf:"bytes,10,opt,name=hostInfo" json:"hostInfo,omitempty"`
	Mountpoint            string `protobuf:"bytes,11,opt,name=mountpoint" json:"mountpoint,omitempty"`
	SnapshotId            string `protobuf:"bytes,12,opt,name=snapshotId" json:"snapshotId,omitempty"`
	SnapshotName          string `protobuf:"bytes,13,opt,name=snapshotName" json:"snapshotName,omitempty"`
	SnapshotDescription   string `protobuf:"bytes,14,opt,name=snapshotDescription" json:"snapshotDescription,omitempty"`
	DockInfo              string `protobuf:"bytes,15,opt,name=dockInfo" json:"dockInfo,omitempty"`
	PoolId                string `protobuf:"bytes,16,opt,name=poolId" json:"poolId,omitempty"`
	ProfileId             string `protobuf:"bytes,17,opt,name=profileId" json:"profileId,omitempty"`
}

func (m *DockRequest) Reset()                    { *m = DockRequest{} }
func (m *DockRequest) String() string            { return proto.CompactTextString(m) }
func (*DockRequest) ProtoMessage()               {}
func (*DockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DockRequest) GetVolumeId() string {
	if m != nil {
		return m.VolumeId
	}
	return ""
}

func (m *DockRequest) GetVolumeName() string {
	if m != nil {
		return m.VolumeName
	}
	return ""
}

func (m *DockRequest) GetVolumeDescription() string {
	if m != nil {
		return m.VolumeDescription
	}
	return ""
}

func (m *DockRequest) GetVolumeSize() int64 {
	if m != nil {
		return m.VolumeSize
	}
	return 0
}

func (m *DockRequest) GetAttachmentId() string {
	if m != nil {
		return m.AttachmentId
	}
	return ""
}

func (m *DockRequest) GetAttachmentName() string {
	if m != nil {
		return m.AttachmentName
	}
	return ""
}

func (m *DockRequest) GetAttachmentDescription() string {
	if m != nil {
		return m.AttachmentDescription
	}
	return ""
}

func (m *DockRequest) GetDoLocalAttach() bool {
	if m != nil {
		return m.DoLocalAttach
	}
	return false
}

func (m *DockRequest) GetMultiPath() bool {
	if m != nil {
		return m.MultiPath
	}
	return false
}

func (m *DockRequest) GetHostInfo() string {
	if m != nil {
		return m.HostInfo
	}
	return ""
}

func (m *DockRequest) GetMountpoint() string {
	if m != nil {
		return m.Mountpoint
	}
	return ""
}

func (m *DockRequest) GetSnapshotId() string {
	if m != nil {
		return m.SnapshotId
	}
	return ""
}

func (m *DockRequest) GetSnapshotName() string {
	if m != nil {
		return m.SnapshotName
	}
	return ""
}

func (m *DockRequest) GetSnapshotDescription() string {
	if m != nil {
		return m.SnapshotDescription
	}
	return ""
}

func (m *DockRequest) GetDockInfo() string {
	if m != nil {
		return m.DockInfo
	}
	return ""
}

func (m *DockRequest) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

func (m *DockRequest) GetProfileId() string {
	if m != nil {
		return m.ProfileId
	}
	return ""
}

// The DockResponse message containing all properties of
// resource response from dock service.
type DockResponse struct {
	Status  string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DockResponse) Reset()                    { *m = DockResponse{} }
func (m *DockResponse) String() string            { return proto.CompactTextString(m) }
func (*DockResponse) ProtoMessage()               {}
func (*DockResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DockResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DockResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DockResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*DockRequest)(nil), "opensds.DockRequest")
	proto.RegisterType((*DockResponse)(nil), "opensds.DockResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Dock service

type DockClient interface {
	// Create a volume
	CreateVolume(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error)
	// Get a volume
	GetVolume(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error)
	// Delete a volume
	DeleteVolume(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error)
	// Create a volume attachment
	CreateVolumeAttachment(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error)
	// Update a volume attachment
	UpdateVolumeAttachment(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error)
	// Delete a volume attachment
	DeleteVolumeAttachment(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error)
	// Create a volume snapshot
	CreateVolumeSnapshot(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error)
	// Get a volume snapshot
	GetVolumeSnapshot(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error)
	// Delete a volume snapshot
	DeleteVolumeSnapshot(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error)
}

type dockClient struct {
	cc *grpc.ClientConn
}

func NewDockClient(cc *grpc.ClientConn) DockClient {
	return &dockClient{cc}
}

func (c *dockClient) CreateVolume(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error) {
	out := new(DockResponse)
	err := grpc.Invoke(ctx, "/opensds.Dock/CreateVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockClient) GetVolume(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error) {
	out := new(DockResponse)
	err := grpc.Invoke(ctx, "/opensds.Dock/GetVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockClient) DeleteVolume(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error) {
	out := new(DockResponse)
	err := grpc.Invoke(ctx, "/opensds.Dock/DeleteVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockClient) CreateVolumeAttachment(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error) {
	out := new(DockResponse)
	err := grpc.Invoke(ctx, "/opensds.Dock/CreateVolumeAttachment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockClient) UpdateVolumeAttachment(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error) {
	out := new(DockResponse)
	err := grpc.Invoke(ctx, "/opensds.Dock/UpdateVolumeAttachment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockClient) DeleteVolumeAttachment(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error) {
	out := new(DockResponse)
	err := grpc.Invoke(ctx, "/opensds.Dock/DeleteVolumeAttachment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockClient) CreateVolumeSnapshot(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error) {
	out := new(DockResponse)
	err := grpc.Invoke(ctx, "/opensds.Dock/CreateVolumeSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockClient) GetVolumeSnapshot(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error) {
	out := new(DockResponse)
	err := grpc.Invoke(ctx, "/opensds.Dock/GetVolumeSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockClient) DeleteVolumeSnapshot(ctx context.Context, in *DockRequest, opts ...grpc.CallOption) (*DockResponse, error) {
	out := new(DockResponse)
	err := grpc.Invoke(ctx, "/opensds.Dock/DeleteVolumeSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dock service

type DockServer interface {
	// Create a volume
	CreateVolume(context.Context, *DockRequest) (*DockResponse, error)
	// Get a volume
	GetVolume(context.Context, *DockRequest) (*DockResponse, error)
	// Delete a volume
	DeleteVolume(context.Context, *DockRequest) (*DockResponse, error)
	// Create a volume attachment
	CreateVolumeAttachment(context.Context, *DockRequest) (*DockResponse, error)
	// Update a volume attachment
	UpdateVolumeAttachment(context.Context, *DockRequest) (*DockResponse, error)
	// Delete a volume attachment
	DeleteVolumeAttachment(context.Context, *DockRequest) (*DockResponse, error)
	// Create a volume snapshot
	CreateVolumeSnapshot(context.Context, *DockRequest) (*DockResponse, error)
	// Get a volume snapshot
	GetVolumeSnapshot(context.Context, *DockRequest) (*DockResponse, error)
	// Delete a volume snapshot
	DeleteVolumeSnapshot(context.Context, *DockRequest) (*DockResponse, error)
}

func RegisterDockServer(s *grpc.Server, srv DockServer) {
	s.RegisterService(&_Dock_serviceDesc, srv)
}

func _Dock_CreateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).CreateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensds.Dock/CreateVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).CreateVolume(ctx, req.(*DockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dock_GetVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).GetVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensds.Dock/GetVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).GetVolume(ctx, req.(*DockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dock_DeleteVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).DeleteVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensds.Dock/DeleteVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).DeleteVolume(ctx, req.(*DockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dock_CreateVolumeAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).CreateVolumeAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensds.Dock/CreateVolumeAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).CreateVolumeAttachment(ctx, req.(*DockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dock_UpdateVolumeAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).UpdateVolumeAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensds.Dock/UpdateVolumeAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).UpdateVolumeAttachment(ctx, req.(*DockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dock_DeleteVolumeAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).DeleteVolumeAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensds.Dock/DeleteVolumeAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).DeleteVolumeAttachment(ctx, req.(*DockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dock_CreateVolumeSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).CreateVolumeSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensds.Dock/CreateVolumeSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).CreateVolumeSnapshot(ctx, req.(*DockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dock_GetVolumeSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).GetVolumeSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensds.Dock/GetVolumeSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).GetVolumeSnapshot(ctx, req.(*DockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dock_DeleteVolumeSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).DeleteVolumeSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensds.Dock/DeleteVolumeSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).DeleteVolumeSnapshot(ctx, req.(*DockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opensds.Dock",
	HandlerType: (*DockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVolume",
			Handler:    _Dock_CreateVolume_Handler,
		},
		{
			MethodName: "GetVolume",
			Handler:    _Dock_GetVolume_Handler,
		},
		{
			MethodName: "DeleteVolume",
			Handler:    _Dock_DeleteVolume_Handler,
		},
		{
			MethodName: "CreateVolumeAttachment",
			Handler:    _Dock_CreateVolumeAttachment_Handler,
		},
		{
			MethodName: "UpdateVolumeAttachment",
			Handler:    _Dock_UpdateVolumeAttachment_Handler,
		},
		{
			MethodName: "DeleteVolumeAttachment",
			Handler:    _Dock_DeleteVolumeAttachment_Handler,
		},
		{
			MethodName: "CreateVolumeSnapshot",
			Handler:    _Dock_CreateVolumeSnapshot_Handler,
		},
		{
			MethodName: "GetVolumeSnapshot",
			Handler:    _Dock_GetVolumeSnapshot_Handler,
		},
		{
			MethodName: "DeleteVolumeSnapshot",
			Handler:    _Dock_DeleteVolumeSnapshot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opensds.proto",
}

func init() { proto.RegisterFile("opensds.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x25, 0x34, 0xcd, 0xc7, 0xd4, 0x29, 0x64, 0x49, 0xab, 0x55, 0x85, 0x50, 0x14, 0x21, 0x94,
	0x03, 0xaa, 0x10, 0x70, 0x42, 0xea, 0xa1, 0x90, 0xaa, 0x8a, 0x84, 0x10, 0x4a, 0x45, 0xef, 0x8b,
	0x3d, 0xc5, 0x56, 0xed, 0x9d, 0xc5, 0xbb, 0xe6, 0xc0, 0xbf, 0xe5, 0x07, 0xf0, 0x1f, 0xd0, 0xee,
	0xda, 0xf1, 0xb6, 0xe4, 0x82, 0xc5, 0x2d, 0xef, 0xbd, 0xc9, 0xf3, 0xcc, 0xf3, 0x93, 0x61, 0x42,
	0x0a, 0xa5, 0x4e, 0xf4, 0xa9, 0x2a, 0xc9, 0x10, 0x1b, 0xd6, 0x70, 0xf1, 0xab, 0x0f, 0x07, 0x2b,
	0x8a, 0x6f, 0x37, 0xf8, 0xbd, 0x42, 0x6d, 0xd8, 0x09, 0x8c, 0x7e, 0x50, 0x5e, 0x15, 0xb8, 0x4e,
	0x78, 0x6f, 0xde, 0x5b, 0x8e, 0x37, 0x5b, 0xcc, 0x9e, 0x01, 0xf8, 0xdf, 0x9f, 0x44, 0x81, 0xfc,
	0xa1, 0x53, 0x03, 0x86, 0xbd, 0x84, 0xa9, 0x47, 0x2b, 0xd4, 0x71, 0x99, 0x29, 0x93, 0x91, 0xe4,
	0x7b, 0x6e, 0xec, 0x6f, 0xa1, 0x75, 0xbb, 0xca, 0x7e, 0x22, 0xef, 0xcf, 0x7b, 0xcb, 0xbd, 0x4d,
	0xc0, 0xb0, 0x05, 0x44, 0xc2, 0x18, 0x11, 0xa7, 0x05, 0x4a, 0xb3, 0x4e, 0xf8, 0xbe, 0x33, 0xba,
	0xc3, 0xb1, 0x17, 0x70, 0xd8, 0x62, 0xb7, 0xd5, 0xc0, 0x4d, 0xdd, 0x63, 0xd9, 0x5b, 0x38, 0x6a,
	0x99, 0x70, 0xbb, 0xa1, 0x1b, 0xdf, 0x2d, 0xb2, 0xe7, 0x30, 0x49, 0xe8, 0x23, 0xc5, 0x22, 0x3f,
	0x77, 0x3a, 0x1f, 0xcd, 0x7b, 0xcb, 0xd1, 0xe6, 0x2e, 0xc9, 0x9e, 0xc2, 0xb8, 0xa8, 0x72, 0x93,
	0x7d, 0x16, 0x26, 0xe5, 0x63, 0x37, 0xd1, 0x12, 0x36, 0xcf, 0x94, 0xb4, 0x59, 0xcb, 0x1b, 0xe2,
	0xe0, 0xf3, 0x6c, 0xb0, 0x4d, 0xa0, 0xa0, 0x4a, 0x1a, 0x45, 0x99, 0x34, 0xfc, 0xc0, 0xe7, 0xd9,
	0x32, 0x56, 0xd7, 0x52, 0x28, 0x9d, 0x92, 0xbd, 0x3f, 0xf2, 0x7a, 0xcb, 0xd8, 0x84, 0x1a, 0xe4,
	0x6e, 0x9f, 0xf8, 0x84, 0x42, 0x8e, 0xbd, 0x82, 0x27, 0x0d, 0x0e, 0xef, 0x3e, 0x74, 0xa3, 0xbb,
	0x24, 0xbb, 0x71, 0x42, 0xf1, 0xad, 0xdb, 0xf8, 0x91, 0xdf, 0xb8, 0xc1, 0xec, 0x18, 0x06, 0x8a,
	0x28, 0x5f, 0x27, 0xfc, 0xb1, 0x53, 0x6a, 0x64, 0x33, 0x50, 0x25, 0xdd, 0x64, 0xb9, 0xad, 0xcd,
	0xd4, 0x49, 0x2d, 0xb1, 0xb8, 0x86, 0xc8, 0x57, 0x4c, 0x2b, 0x92, 0x1a, 0xad, 0x8b, 0x36, 0xc2,
	0x54, 0xba, 0x6e, 0x58, 0x8d, 0x18, 0x87, 0x61, 0x81, 0x5a, 0x8b, 0x6f, 0x4d, 0xb9, 0x1a, 0xc8,
	0x66, 0xb0, 0x8f, 0x65, 0x49, 0x65, 0xdd, 0x26, 0x0f, 0x5e, 0xff, 0xee, 0x43, 0xdf, 0x1a, 0xb3,
	0x33, 0x88, 0x3e, 0x94, 0x28, 0x0c, 0x5e, 0xbb, 0xfa, 0xb0, 0xd9, 0x69, 0xd3, 0xf6, 0xa0, 0xda,
	0x27, 0x47, 0xf7, 0x58, 0xbf, 0xcd, 0xe2, 0x01, 0x7b, 0x07, 0xe3, 0x4b, 0x34, 0xdd, 0xfe, 0x7b,
	0x06, 0xd1, 0x0a, 0x73, 0xec, 0xfa, 0xe8, 0x4b, 0x38, 0x0e, 0x37, 0x3f, 0xdf, 0xf6, 0xb0, 0x83,
	0xd1, 0x17, 0x95, 0xfc, 0x1f, 0xa3, 0xf0, 0xa0, 0xee, 0x46, 0x17, 0x30, 0x0b, 0x4f, 0xbb, 0xaa,
	0xab, 0xf6, 0xaf, 0x36, 0xef, 0x61, 0xba, 0x7d, 0x39, 0x5d, 0x3d, 0x2e, 0x60, 0x16, 0xde, 0xd4,
	0xd1, 0xe6, 0xeb, 0xc0, 0x7d, 0x3b, 0xdf, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x85, 0x1c, 0xca,
	0xbf, 0x4c, 0x05, 0x00, 0x00,
}
