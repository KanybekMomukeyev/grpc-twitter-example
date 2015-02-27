// Code generated by protoc-gen-go.
// source: proto/twitter.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/twitter.proto

It has these top-level messages:
	Ack
	User
	Tweet
	Search
	Timeline
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type Ack struct {
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto1.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}

type User struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()    {}

type Tweet struct {
	Id   uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	User *User  `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
}

func (m *Tweet) Reset()         { *m = Tweet{} }
func (m *Tweet) String() string { return proto1.CompactTextString(m) }
func (*Tweet) ProtoMessage()    {}

func (m *Tweet) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Search struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Search) Reset()         { *m = Search{} }
func (m *Search) String() string { return proto1.CompactTextString(m) }
func (*Search) ProtoMessage()    {}

type Timeline struct {
	Tweet []*Tweet `protobuf:"bytes,1,rep,name=tweet" json:"tweet,omitempty"`
}

func (m *Timeline) Reset()         { *m = Timeline{} }
func (m *Timeline) String() string { return proto1.CompactTextString(m) }
func (*Timeline) ProtoMessage()    {}

func (m *Timeline) GetTweet() []*Tweet {
	if m != nil {
		return m.Tweet
	}
	return nil
}

func init() {
}