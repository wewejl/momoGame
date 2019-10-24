// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

//包名

package mmomsg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//定义具体协议
//sync：同步,async:异步
//msgid :1
type Syncid struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Syncid) Reset()         { *m = Syncid{} }
func (m *Syncid) String() string { return proto.CompactTextString(m) }
func (*Syncid) ProtoMessage()    {}
func (*Syncid) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *Syncid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Syncid.Unmarshal(m, b)
}
func (m *Syncid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Syncid.Marshal(b, m, deterministic)
}
func (m *Syncid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Syncid.Merge(m, src)
}
func (m *Syncid) XXX_Size() int {
	return xxx_messageInfo_Syncid.Size(m)
}
func (m *Syncid) XXX_DiscardUnknown() {
	xxx_messageInfo_Syncid.DiscardUnknown(m)
}

var xxx_messageInfo_Syncid proto.InternalMessageInfo

func (m *Syncid) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

//定义一个Position协议，用于存放位置
type Position struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	V                    float32  `protobuf:"fixed32,4,opt,name=v,proto3" json:"v,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Position) GetV() float32 {
	if m != nil {
		return m.V
	}
	return 0
}

//msgid : 200
//定 义 一 个 BroadCast 广 播 协 议 , 根 据 类 型 , 发 送 消 息
type BroadCast struct {
	Pid int32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Tp  int32 `protobuf:"varint,2,opt,name=Tp,proto3" json:"Tp,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*BroadCast_Content
	//	*BroadCast_P
	//	*BroadCast_ActionData
	Data                 isBroadCast_Data `protobuf_oneof:"Data"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BroadCast) Reset()         { *m = BroadCast{} }
func (m *BroadCast) String() string { return proto.CompactTextString(m) }
func (*BroadCast) ProtoMessage()    {}
func (*BroadCast) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *BroadCast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadCast.Unmarshal(m, b)
}
func (m *BroadCast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadCast.Marshal(b, m, deterministic)
}
func (m *BroadCast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadCast.Merge(m, src)
}
func (m *BroadCast) XXX_Size() int {
	return xxx_messageInfo_BroadCast.Size(m)
}
func (m *BroadCast) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadCast.DiscardUnknown(m)
}

var xxx_messageInfo_BroadCast proto.InternalMessageInfo

func (m *BroadCast) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *BroadCast) GetTp() int32 {
	if m != nil {
		return m.Tp
	}
	return 0
}

type isBroadCast_Data interface {
	isBroadCast_Data()
}

type BroadCast_Content struct {
	Content string `protobuf:"bytes,3,opt,name=content,proto3,oneof"`
}

type BroadCast_P struct {
	P *Position `protobuf:"bytes,4,opt,name=p,proto3,oneof"`
}

type BroadCast_ActionData struct {
	ActionData int32 `protobuf:"varint,5,opt,name=ActionData,proto3,oneof"`
}

func (*BroadCast_Content) isBroadCast_Data() {}

func (*BroadCast_P) isBroadCast_Data() {}

func (*BroadCast_ActionData) isBroadCast_Data() {}

func (m *BroadCast) GetData() isBroadCast_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BroadCast) GetContent() string {
	if x, ok := m.GetData().(*BroadCast_Content); ok {
		return x.Content
	}
	return ""
}

func (m *BroadCast) GetP() *Position {
	if x, ok := m.GetData().(*BroadCast_P); ok {
		return x.P
	}
	return nil
}

func (m *BroadCast) GetActionData() int32 {
	if x, ok := m.GetData().(*BroadCast_ActionData); ok {
		return x.ActionData
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BroadCast) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BroadCast_Content)(nil),
		(*BroadCast_P)(nil),
		(*BroadCast_ActionData)(nil),
	}
}

//另一个talk
type Talk struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Talk) Reset()         { *m = Talk{} }
func (m *Talk) String() string { return proto.CompactTextString(m) }
func (*Talk) ProtoMessage()    {}
func (*Talk) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *Talk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Talk.Unmarshal(m, b)
}
func (m *Talk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Talk.Marshal(b, m, deterministic)
}
func (m *Talk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Talk.Merge(m, src)
}
func (m *Talk) XXX_Size() int {
	return xxx_messageInfo_Talk.Size(m)
}
func (m *Talk) XXX_DiscardUnknown() {
	xxx_messageInfo_Talk.DiscardUnknown(m)
}

var xxx_messageInfo_Talk proto.InternalMessageInfo

func (m *Talk) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

//用于同步
type SyncPlayers struct {
	Ps                   []*SyncPlayers_Player `protobuf:"bytes,1,rep,name=ps,proto3" json:"ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SyncPlayers) Reset()         { *m = SyncPlayers{} }
func (m *SyncPlayers) String() string { return proto.CompactTextString(m) }
func (*SyncPlayers) ProtoMessage()    {}
func (*SyncPlayers) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *SyncPlayers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPlayers.Unmarshal(m, b)
}
func (m *SyncPlayers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPlayers.Marshal(b, m, deterministic)
}
func (m *SyncPlayers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPlayers.Merge(m, src)
}
func (m *SyncPlayers) XXX_Size() int {
	return xxx_messageInfo_SyncPlayers.Size(m)
}
func (m *SyncPlayers) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPlayers.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPlayers proto.InternalMessageInfo

func (m *SyncPlayers) GetPs() []*SyncPlayers_Player {
	if m != nil {
		return m.Ps
	}
	return nil
}

type SyncPlayers_Player struct {
	Pid                  int32     `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	P                    *Position `protobuf:"bytes,2,opt,name=p,proto3" json:"p,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SyncPlayers_Player) Reset()         { *m = SyncPlayers_Player{} }
func (m *SyncPlayers_Player) String() string { return proto.CompactTextString(m) }
func (*SyncPlayers_Player) ProtoMessage()    {}
func (*SyncPlayers_Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4, 0}
}

func (m *SyncPlayers_Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPlayers_Player.Unmarshal(m, b)
}
func (m *SyncPlayers_Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPlayers_Player.Marshal(b, m, deterministic)
}
func (m *SyncPlayers_Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPlayers_Player.Merge(m, src)
}
func (m *SyncPlayers_Player) XXX_Size() int {
	return xxx_messageInfo_SyncPlayers_Player.Size(m)
}
func (m *SyncPlayers_Player) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPlayers_Player.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPlayers_Player proto.InternalMessageInfo

func (m *SyncPlayers_Player) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *SyncPlayers_Player) GetP() *Position {
	if m != nil {
		return m.P
	}
	return nil
}

func init() {
	proto.RegisterType((*Syncid)(nil), "mmomsg.syncid")
	proto.RegisterType((*Position)(nil), "mmomsg.Position")
	proto.RegisterType((*BroadCast)(nil), "mmomsg.BroadCast")
	proto.RegisterType((*Talk)(nil), "mmomsg.talk")
	proto.RegisterType((*SyncPlayers)(nil), "mmomsg.SyncPlayers")
	proto.RegisterType((*SyncPlayers_Player)(nil), "mmomsg.SyncPlayers.Player")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x6b, 0xe7, 0x02, 0x39, 0x41, 0xa8, 0xf2, 0x64, 0x65, 0x40, 0x51, 0xa6, 0x8a, 0x21,
	0x43, 0xd9, 0xd8, 0x08, 0x0c, 0x19, 0x2b, 0xd3, 0x17, 0x30, 0x49, 0x55, 0x45, 0x34, 0xb6, 0x15,
	0x9b, 0xaa, 0xee, 0x7b, 0xf0, 0xbe, 0xc8, 0xb6, 0x02, 0x95, 0xe8, 0x74, 0xfc, 0x1d, 0x5b, 0xff,
	0x45, 0x86, 0x6c, 0xd4, 0xfb, 0x5a, 0x4d, 0xd2, 0x48, 0x92, 0x8e, 0xa3, 0x1c, 0xf5, 0xbe, 0x2a,
	0x20, 0xd5, 0x56, 0x74, 0x43, 0x4f, 0x96, 0x10, 0xa9, 0xa1, 0xa7, 0xa8, 0x44, 0xab, 0x84, 0xb9,
	0x63, 0xd5, 0xc0, 0xed, 0x46, 0xea, 0xc1, 0x0c, 0x52, 0x90, 0x3b, 0x40, 0x27, 0x7f, 0x87, 0x19,
	0x3a, 0x39, 0xb2, 0x14, 0x07, 0xb2, 0x8e, 0xce, 0x34, 0x0a, 0x74, 0x76, 0x74, 0xa4, 0x71, 0xa0,
	0x63, 0xf5, 0x8d, 0x20, 0x6b, 0x26, 0xc9, 0xfb, 0x57, 0xae, 0xcd, 0x7f, 0x0f, 0x72, 0x0f, 0x78,
	0xab, 0xbc, 0x54, 0xc2, 0xf0, 0x56, 0x91, 0x02, 0x6e, 0x3a, 0x29, 0xcc, 0x4e, 0x18, 0xaf, 0x98,
	0xb5, 0x0b, 0x36, 0x2f, 0x48, 0x09, 0x48, 0x79, 0xe5, 0x7c, 0xbd, 0xac, 0x43, 0xfe, 0x7a, 0x0e,
	0xd8, 0x2e, 0x18, 0x52, 0xa4, 0x04, 0x78, 0xe9, 0x1c, 0xbe, 0x71, 0xc3, 0x69, 0xe2, 0x54, 0xdb,
	0x05, 0xbb, 0xd8, 0x35, 0x29, 0xc4, 0x6e, 0x56, 0x25, 0xc4, 0x86, 0x1f, 0x3e, 0x09, 0xfd, 0xf3,
	0x73, 0xa9, 0xb2, 0x5f, 0xb7, 0xea, 0x0b, 0xf2, 0x77, 0x2b, 0xba, 0xcd, 0x81, 0xdb, 0xdd, 0xa4,
	0xc9, 0x23, 0x60, 0xa5, 0x29, 0x2a, 0xa3, 0x55, 0xbe, 0x2e, 0x66, 0xf7, 0x8b, 0x07, 0x75, 0x98,
	0x0c, 0x2b, 0x5d, 0x3c, 0x43, 0x1a, 0xe8, 0x4a, 0xe1, 0x07, 0x57, 0x02, 0x5f, 0x2f, 0xc1, 0x90,
	0xfa, 0x48, 0xfd, 0xff, 0x3c, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xce, 0x37, 0x5c, 0x79, 0xac,
	0x01, 0x00, 0x00,
}
