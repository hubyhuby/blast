// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/management/management.proto

package management

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	raft "github.com/mosuka/blast/protobuf/raft"
	grpc "google.golang.org/grpc"
	math "math"
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

type ManagementCommand_Type int32

const (
	ManagementCommand_UNKNOWN_COMMAND       ManagementCommand_Type = 0
	ManagementCommand_SET_METADATA          ManagementCommand_Type = 1
	ManagementCommand_DELETE_METADATA       ManagementCommand_Type = 2
	ManagementCommand_PUT_KEY_VALUE_PAIR    ManagementCommand_Type = 3
	ManagementCommand_DELETE_KEY_VALUE_PAIR ManagementCommand_Type = 4
)

var ManagementCommand_Type_name = map[int32]string{
	0: "UNKNOWN_COMMAND",
	1: "SET_METADATA",
	2: "DELETE_METADATA",
	3: "PUT_KEY_VALUE_PAIR",
	4: "DELETE_KEY_VALUE_PAIR",
}

var ManagementCommand_Type_value = map[string]int32{
	"UNKNOWN_COMMAND":       0,
	"SET_METADATA":          1,
	"DELETE_METADATA":       2,
	"PUT_KEY_VALUE_PAIR":    3,
	"DELETE_KEY_VALUE_PAIR": 4,
}

func (x ManagementCommand_Type) String() string {
	return proto.EnumName(ManagementCommand_Type_name, int32(x))
}

func (ManagementCommand_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e030ad796566078, []int{1, 0}
}

type KeyValuePair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *any.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValuePair) Reset()         { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()    {}
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e030ad796566078, []int{0}
}

func (m *KeyValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValuePair.Unmarshal(m, b)
}
func (m *KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValuePair.Marshal(b, m, deterministic)
}
func (m *KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValuePair.Merge(m, src)
}
func (m *KeyValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyValuePair.Size(m)
}
func (m *KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValuePair proto.InternalMessageInfo

func (m *KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValuePair) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type ManagementCommand struct {
	Type                 ManagementCommand_Type `protobuf:"varint,1,opt,name=type,proto3,enum=management.ManagementCommand_Type" json:"type,omitempty"`
	Data                 *any.Any               `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ManagementCommand) Reset()         { *m = ManagementCommand{} }
func (m *ManagementCommand) String() string { return proto.CompactTextString(m) }
func (*ManagementCommand) ProtoMessage()    {}
func (*ManagementCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e030ad796566078, []int{1}
}

func (m *ManagementCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagementCommand.Unmarshal(m, b)
}
func (m *ManagementCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagementCommand.Marshal(b, m, deterministic)
}
func (m *ManagementCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagementCommand.Merge(m, src)
}
func (m *ManagementCommand) XXX_Size() int {
	return xxx_messageInfo_ManagementCommand.Size(m)
}
func (m *ManagementCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagementCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ManagementCommand proto.InternalMessageInfo

func (m *ManagementCommand) GetType() ManagementCommand_Type {
	if m != nil {
		return m.Type
	}
	return ManagementCommand_UNKNOWN_COMMAND
}

func (m *ManagementCommand) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("management.ManagementCommand_Type", ManagementCommand_Type_name, ManagementCommand_Type_value)
	proto.RegisterType((*KeyValuePair)(nil), "management.KeyValuePair")
	proto.RegisterType((*ManagementCommand)(nil), "management.ManagementCommand")
}

func init() {
	proto.RegisterFile("protobuf/management/management.proto", fileDescriptor_5e030ad796566078)
}

var fileDescriptor_5e030ad796566078 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x51, 0x6f, 0x93, 0x50,
	0x14, 0xc7, 0x61, 0xb0, 0xa9, 0xc7, 0xa9, 0x78, 0xd4, 0xa5, 0xc3, 0x97, 0x85, 0xf8, 0xd0, 0xa8,
	0x83, 0xa4, 0x46, 0x13, 0xe3, 0x5e, 0xb0, 0xbd, 0x69, 0xb4, 0x2d, 0x6b, 0x5a, 0x3a, 0xa3, 0x2f,
	0xe4, 0xd6, 0x9e, 0x75, 0xcd, 0x80, 0x4b, 0xda, 0xcb, 0x12, 0x3e, 0xa0, 0xdf, 0xc6, 0x0f, 0x61,
	0x80, 0xad, 0x25, 0x9b, 0x34, 0xc6, 0x17, 0x72, 0xf3, 0xff, 0xff, 0xfe, 0xe7, 0x9c, 0xcb, 0x01,
	0x78, 0x95, 0x2c, 0x85, 0x14, 0xd3, 0xf4, 0xdc, 0x89, 0x78, 0xcc, 0xe7, 0x14, 0x51, 0x2c, 0x2b,
	0x47, 0xbb, 0xb0, 0x11, 0x36, 0x8a, 0x79, 0x38, 0x17, 0x62, 0x1e, 0x92, 0xb3, 0x0e, 0xf2, 0x38,
	0x2b, 0x31, 0xf3, 0xe5, 0x6d, 0x8b, 0xa2, 0x44, 0xde, 0x98, 0x8d, 0xb5, 0xba, 0xe4, 0xe7, 0xb2,
	0x78, 0x94, 0x8e, 0xd5, 0x87, 0xfd, 0x1e, 0x65, 0x67, 0x3c, 0x4c, 0x69, 0xc8, 0x17, 0x4b, 0x34,
	0x40, 0xbb, 0xa4, 0xac, 0xa1, 0x1e, 0xa9, 0xcd, 0x07, 0xa3, 0xfc, 0x88, 0xaf, 0x61, 0xf7, 0x2a,
	0xb7, 0x1b, 0x3b, 0x47, 0x6a, 0xf3, 0x61, 0xeb, 0xb9, 0x5d, 0x36, 0xb2, 0x6f, 0x4a, 0xda, 0x6e,
	0x9c, 0x8d, 0x4a, 0xc4, 0xfa, 0xad, 0xc2, 0xd3, 0xc1, 0x7a, 0xdc, 0xb6, 0x88, 0x22, 0x1e, 0xcf,
	0xf0, 0x03, 0xe8, 0x32, 0x4b, 0xa8, 0x28, 0xfa, 0xb8, 0x65, 0xd9, 0x95, 0x2b, 0xde, 0x81, 0x6d,
	0x3f, 0x4b, 0x68, 0x54, 0xf0, 0xd8, 0x04, 0x7d, 0xc6, 0x25, 0xdf, 0xda, 0xb8, 0x20, 0xac, 0x14,
	0xf4, 0x3c, 0x87, 0xcf, 0xe0, 0xc9, 0xc4, 0xeb, 0x79, 0xa7, 0xdf, 0xbc, 0xa0, 0x7d, 0x3a, 0x18,
	0xb8, 0x5e, 0xc7, 0x50, 0xd0, 0x80, 0xfd, 0x31, 0xf3, 0x83, 0x01, 0xf3, 0xdd, 0x8e, 0xeb, 0xbb,
	0x86, 0x9a, 0x63, 0x1d, 0xd6, 0x67, 0x3e, 0xdb, 0x88, 0x3b, 0x78, 0x00, 0x38, 0x9c, 0xf8, 0x41,
	0x8f, 0x7d, 0x0f, 0xce, 0xdc, 0xfe, 0x84, 0x05, 0x43, 0xf7, 0xcb, 0xc8, 0xd0, 0xf0, 0x10, 0x5e,
	0x5c, 0xc3, 0xb7, 0x2c, 0xbd, 0xf5, 0x4b, 0x03, 0xd8, 0xdc, 0x00, 0xdf, 0x82, 0xfe, 0x55, 0x2c,
	0x62, 0x04, 0xbb, 0x78, 0xc1, 0x9e, 0x98, 0x91, 0x79, 0x70, 0x67, 0x6a, 0x96, 0xef, 0xc5, 0x52,
	0xf0, 0x18, 0x76, 0xfb, 0xc4, 0xaf, 0xe8, 0x1f, 0x71, 0x07, 0xee, 0x75, 0x49, 0xe6, 0x10, 0xd6,
	0x40, 0x66, 0xa5, 0x90, 0xa5, 0xe0, 0x7b, 0x80, 0x2e, 0xc9, 0x76, 0x98, 0xae, 0x24, 0x2d, 0x6b,
	0x33, 0x8f, 0xca, 0xcc, 0x35, 0x66, 0x29, 0x78, 0x02, 0xf7, 0xc7, 0x31, 0x4f, 0x56, 0x17, 0x42,
	0xd6, 0x86, 0xea, 0xa7, 0xfc, 0x04, 0x5a, 0x97, 0x24, 0x36, 0xaa, 0x3b, 0xae, 0x7e, 0x5f, 0x66,
	0xad, 0x63, 0x29, 0xf8, 0x11, 0xb4, 0xf1, 0xd6, 0x70, 0x7d, 0xdf, 0x13, 0xd8, 0xeb, 0x50, 0x48,
	0x92, 0xfe, 0x27, 0xfd, 0xf9, 0xf8, 0xc7, 0x9b, 0xf9, 0x42, 0x5e, 0xa4, 0x53, 0xfb, 0xa7, 0x88,
	0x9c, 0x48, 0xac, 0xd2, 0x4b, 0xee, 0x4c, 0x43, 0xbe, 0x92, 0xce, 0x5f, 0x7e, 0xd1, 0xe9, 0x5e,
	0x21, 0xbe, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x64, 0x5d, 0xf3, 0xc0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ManagementClient is the client API for Management service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManagementClient interface {
	Join(ctx context.Context, in *raft.Node, opts ...grpc.CallOption) (*empty.Empty, error)
	Leave(ctx context.Context, in *raft.Node, opts ...grpc.CallOption) (*empty.Empty, error)
	GetNode(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.Node, error)
	GetCluster(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.Cluster, error)
	Snapshot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*KeyValuePair, error)
	Set(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*empty.Empty, error)
}

type managementClient struct {
	cc *grpc.ClientConn
}

func NewManagementClient(cc *grpc.ClientConn) ManagementClient {
	return &managementClient{cc}
}

func (c *managementClient) Join(ctx context.Context, in *raft.Node, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/management.Management/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) Leave(ctx context.Context, in *raft.Node, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/management.Management/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) GetNode(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.Node, error) {
	out := new(raft.Node)
	err := c.cc.Invoke(ctx, "/management.Management/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) GetCluster(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.Cluster, error) {
	out := new(raft.Cluster)
	err := c.cc.Invoke(ctx, "/management.Management/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) Snapshot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/management.Management/Snapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) Get(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*KeyValuePair, error) {
	out := new(KeyValuePair)
	err := c.cc.Invoke(ctx, "/management.Management/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) Set(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/management.Management/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) Delete(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/management.Management/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagementServer is the server API for Management service.
type ManagementServer interface {
	Join(context.Context, *raft.Node) (*empty.Empty, error)
	Leave(context.Context, *raft.Node) (*empty.Empty, error)
	GetNode(context.Context, *empty.Empty) (*raft.Node, error)
	GetCluster(context.Context, *empty.Empty) (*raft.Cluster, error)
	Snapshot(context.Context, *empty.Empty) (*empty.Empty, error)
	Get(context.Context, *KeyValuePair) (*KeyValuePair, error)
	Set(context.Context, *KeyValuePair) (*empty.Empty, error)
	Delete(context.Context, *KeyValuePair) (*empty.Empty, error)
}

func RegisterManagementServer(s *grpc.Server, srv ManagementServer) {
	s.RegisterService(&_Management_serviceDesc, srv)
}

func _Management_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).Join(ctx, req.(*raft.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).Leave(ctx, req.(*raft.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).GetNode(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).GetCluster(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).Snapshot(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValuePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).Get(ctx, req.(*KeyValuePair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValuePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).Set(ctx, req.(*KeyValuePair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValuePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).Delete(ctx, req.(*KeyValuePair))
	}
	return interceptor(ctx, in, info, handler)
}

var _Management_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.Management",
	HandlerType: (*ManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Management_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Management_Leave_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _Management_GetNode_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _Management_GetCluster_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _Management_Snapshot_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Management_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Management_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Management_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/management/management.proto",
}
