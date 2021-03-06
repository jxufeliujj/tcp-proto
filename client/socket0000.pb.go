// Code generated by protoc-gen-go.
// source: socket0000.proto
// DO NOT EDIT!

/*
Package client is a generated protocol buffer package.

It is generated from these files:
	socket0000.proto

It has these top-level messages:
	SCS100000
	SSC100000
	TestData
	SCS100001
	SSC100001
	SCS100002
*/
package client

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// 请求登陆
type SCS100000 struct {
	UserId           *string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCS100000) Reset()         { *m = SCS100000{} }
func (m *SCS100000) String() string { return proto.CompactTextString(m) }
func (*SCS100000) ProtoMessage()    {}

func (m *SCS100000) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

type SSC100000 struct {
	Type             *uint32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SSC100000) Reset()         { *m = SSC100000{} }
func (m *SSC100000) String() string { return proto.CompactTextString(m) }
func (*SSC100000) ProtoMessage()    {}

func (m *SSC100000) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type TestData struct {
	Time             *int64  `protobuf:"zigzag64,1,opt,name=time" json:"time,omitempty"`
	UserId           *string `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
	State            *bool   `protobuf:"varint,3,opt,name=state" json:"state,omitempty"`
	Data             *int32  `protobuf:"zigzag32,4,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TestData) Reset()         { *m = TestData{} }
func (m *TestData) String() string { return proto.CompactTextString(m) }
func (*TestData) ProtoMessage()    {}

func (m *TestData) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *TestData) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *TestData) GetState() bool {
	if m != nil && m.State != nil {
		return *m.State
	}
	return false
}

func (m *TestData) GetData() int32 {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return 0
}

// 测试数据
type SCS100001 struct {
	Num              *int64    `protobuf:"zigzag64,1,opt,name=num" json:"num,omitempty"`
	Msg              *TestData `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SCS100001) Reset()         { *m = SCS100001{} }
func (m *SCS100001) String() string { return proto.CompactTextString(m) }
func (*SCS100001) ProtoMessage()    {}

func (m *SCS100001) GetNum() int64 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *SCS100001) GetMsg() *TestData {
	if m != nil {
		return m.Msg
	}
	return nil
}

// 测试数据返回
type SSC100001 struct {
	Num              *int64 `protobuf:"zigzag64,1,opt,name=num" json:"num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SSC100001) Reset()         { *m = SSC100001{} }
func (m *SSC100001) String() string { return proto.CompactTextString(m) }
func (*SSC100001) ProtoMessage()    {}

func (m *SSC100001) GetNum() int64 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

// 心跳包
type SCS100002 struct {
	Beat             *string `protobuf:"bytes,1,opt,name=beat,def=beatHeart" json:"beat,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCS100002) Reset()         { *m = SCS100002{} }
func (m *SCS100002) String() string { return proto.CompactTextString(m) }
func (*SCS100002) ProtoMessage()    {}

const Default_SCS100002_Beat string = "beatHeart"

func (m *SCS100002) GetBeat() string {
	if m != nil && m.Beat != nil {
		return *m.Beat
	}
	return Default_SCS100002_Beat
}

func init() {
}
