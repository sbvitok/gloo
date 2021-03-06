// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/artifact.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//@solo-kit:resource.short_name=art
//@solo-kit:resource.plural_name=artifacts
//
//Gloo Artifacts are used by Gloo to store small bits of binary or file data.
//
//Certain options such as the gRPC option read and write artifacts to one of Gloo's configured
//storage layer.
//
//Artifacts can be backed by files on disk, Kubernetes ConfigMaps, and Consul Key/Value pairs.
//
//Supported artifact backends can be selected in Gloo's boostrap options.
type Artifact struct {
	// Raw data data being stored
	Data map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_c52f0e475c0b5a35, []int{0}
}
func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Artifact) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*Artifact)(nil), "gloo.solo.io.Artifact")
	proto.RegisterMapType((map[string]string)(nil), "gloo.solo.io.Artifact.DataEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/artifact.proto", fileDescriptor_c52f0e475c0b5a35)
}

var fileDescriptor_c52f0e475c0b5a35 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x50, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0xfd, 0xdc, 0xa4, 0x1f, 0x8d, 0xc3, 0x80, 0xac, 0x0a, 0x45, 0x19, 0x42, 0xc4, 0x94, 0x05,
	0x1b, 0x0a, 0x82, 0xaa, 0x4c, 0x54, 0x30, 0xb2, 0x64, 0x64, 0x73, 0x43, 0x08, 0x26, 0x29, 0x37,
	0x72, 0x6e, 0x2b, 0x75, 0xcd, 0xd3, 0xf0, 0x28, 0xec, 0xec, 0x0c, 0xbc, 0x41, 0xde, 0x00, 0xe5,
	0x57, 0x2c, 0x48, 0x6c, 0xf7, 0xdc, 0xf3, 0x27, 0x1d, 0x7a, 0x9d, 0x28, 0x7c, 0xde, 0xac, 0x78,
	0x04, 0x6b, 0x51, 0x40, 0x06, 0x27, 0x0a, 0x44, 0x92, 0x01, 0x88, 0x5c, 0xc3, 0x4b, 0x1c, 0x61,
	0xd1, 0x22, 0x99, 0x2b, 0xb1, 0x3d, 0x13, 0x52, 0xa3, 0x7a, 0x92, 0x11, 0xf2, 0x5c, 0x03, 0x02,
	0xdb, 0xaf, 0x39, 0x5e, 0xdb, 0xb8, 0x02, 0x77, 0x9a, 0x40, 0x02, 0x0d, 0x21, 0xea, 0xab, 0xd5,
	0xb8, 0x5e, 0x93, 0x9a, 0x2a, 0xec, 0x33, 0xd6, 0x31, 0xca, 0x47, 0x89, 0xf2, 0x37, 0xbe, 0xc7,
	0x2d, 0x7f, 0xfc, 0x41, 0xe8, 0xe4, 0xa6, 0xab, 0x65, 0x17, 0xd4, 0xac, 0xad, 0x0e, 0xf1, 0x8d,
	0xc0, 0x9e, 0xf9, 0xfc, 0x67, 0x3f, 0xef, 0x55, 0xfc, 0x56, 0xa2, 0xbc, 0x7b, 0x45, 0xbd, 0x0b,
	0x1b, 0x35, 0x9b, 0xd3, 0x49, 0x5f, 0xea, 0xec, 0xf9, 0x24, 0xb0, 0x67, 0x87, 0x3c, 0x02, 0x1d,
	0x0f, 0xce, 0xfb, 0x8e, 0x5d, 0x9a, 0xef, 0x9f, 0x47, 0xff, 0xc2, 0x41, 0xed, 0x5e, 0x51, 0x6b,
	0x08, 0x63, 0x07, 0xd4, 0x48, 0xe3, 0x9d, 0x43, 0x7c, 0x12, 0x58, 0x61, 0x7d, 0xb2, 0x29, 0x1d,
	0x6f, 0x65, 0xb6, 0x89, 0x9d, 0x51, 0xf3, 0x6b, 0xc1, 0x62, 0x34, 0x27, 0x0b, 0xaf, 0xac, 0xcc,
	0x31, 0x35, 0xa4, 0xc6, 0xb2, 0x32, 0x6d, 0x66, 0xf5, 0xc3, 0x15, 0x65, 0x65, 0x8e, 0x02, 0xb2,
	0xbc, 0x7c, 0xfb, 0xf2, 0xc8, 0xc3, 0xe9, 0xdf, 0xc6, 0xcf, 0xd3, 0xa4, 0x1b, 0x67, 0xf5, 0xbf,
	0x19, 0xe5, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x79, 0x85, 0x7c, 0x1a, 0xb7, 0x01, 0x00, 0x00,
}

func (this *Artifact) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Artifact)
	if !ok {
		that2, ok := that.(Artifact)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Data) != len(that1.Data) {
		return false
	}
	for i := range this.Data {
		if this.Data[i] != that1.Data[i] {
			return false
		}
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
