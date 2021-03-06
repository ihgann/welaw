// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v1/upstream.proto

package grpc_welaw_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Upstream struct {
	Ident               string            `protobuf:"bytes,1,opt,name=ident" json:"ident,omitempty"`
	UpstreamName        string            `protobuf:"bytes,2,opt,name=upstream_name,json=upstreamName" json:"upstream_name,omitempty"`
	UpstreamDescription string            `protobuf:"bytes,3,opt,name=upstream_description,json=upstreamDescription" json:"upstream_description,omitempty"`
	Name                string            `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Description         string            `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Url                 string            `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	Email               string            `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	DefaultUser         string            `protobuf:"bytes,8,opt,name=default_user,json=defaultUser" json:"default_user,omitempty"`
	GeoCoords           string            `protobuf:"bytes,9,opt,name=geo_coords,json=geoCoords" json:"geo_coords,omitempty"`
	Tags                []*LawTag         `protobuf:"bytes,10,rep,name=tags" json:"tags,omitempty"`
	Laws                int32             `protobuf:"varint,11,opt,name=laws" json:"laws,omitempty"`
	Users               int32             `protobuf:"varint,12,opt,name=users" json:"users,omitempty"`
	Metadata            *UpstreamMetadata `protobuf:"bytes,13,opt,name=metadata" json:"metadata,omitempty"`
	Uid                 string            `protobuf:"bytes,14,opt,name=uid" json:"uid,omitempty"`
}

func (m *Upstream) Reset()                    { *m = Upstream{} }
func (m *Upstream) String() string            { return proto.CompactTextString(m) }
func (*Upstream) ProtoMessage()               {}
func (*Upstream) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Upstream) GetIdent() string {
	if m != nil {
		return m.Ident
	}
	return ""
}

func (m *Upstream) GetUpstreamName() string {
	if m != nil {
		return m.UpstreamName
	}
	return ""
}

func (m *Upstream) GetUpstreamDescription() string {
	if m != nil {
		return m.UpstreamDescription
	}
	return ""
}

func (m *Upstream) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Upstream) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Upstream) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Upstream) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Upstream) GetDefaultUser() string {
	if m != nil {
		return m.DefaultUser
	}
	return ""
}

func (m *Upstream) GetGeoCoords() string {
	if m != nil {
		return m.GeoCoords
	}
	return ""
}

func (m *Upstream) GetTags() []*LawTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Upstream) GetLaws() int32 {
	if m != nil {
		return m.Laws
	}
	return 0
}

func (m *Upstream) GetUsers() int32 {
	if m != nil {
		return m.Users
	}
	return 0
}

func (m *Upstream) GetMetadata() *UpstreamMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Upstream) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type UpstreamMetadata struct {
	Email     string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Wikipedia string `protobuf:"bytes,2,opt,name=wikipedia" json:"wikipedia,omitempty"`
	Twitter   string `protobuf:"bytes,3,opt,name=twitter" json:"twitter,omitempty"`
	ApiUrl    string `protobuf:"bytes,4,opt,name=api_url,json=apiUrl" json:"api_url,omitempty"`
	Facebook  string `protobuf:"bytes,5,opt,name=facebook" json:"facebook,omitempty"`
}

func (m *UpstreamMetadata) Reset()                    { *m = UpstreamMetadata{} }
func (m *UpstreamMetadata) String() string            { return proto.CompactTextString(m) }
func (*UpstreamMetadata) ProtoMessage()               {}
func (*UpstreamMetadata) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *UpstreamMetadata) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpstreamMetadata) GetWikipedia() string {
	if m != nil {
		return m.Wikipedia
	}
	return ""
}

func (m *UpstreamMetadata) GetTwitter() string {
	if m != nil {
		return m.Twitter
	}
	return ""
}

func (m *UpstreamMetadata) GetApiUrl() string {
	if m != nil {
		return m.ApiUrl
	}
	return ""
}

func (m *UpstreamMetadata) GetFacebook() string {
	if m != nil {
		return m.Facebook
	}
	return ""
}

type CreateUpstreamRequest struct {
	Upstream *Upstream `protobuf:"bytes,1,opt,name=upstream" json:"upstream,omitempty"`
}

func (m *CreateUpstreamRequest) Reset()                    { *m = CreateUpstreamRequest{} }
func (m *CreateUpstreamRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUpstreamRequest) ProtoMessage()               {}
func (*CreateUpstreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *CreateUpstreamRequest) GetUpstream() *Upstream {
	if m != nil {
		return m.Upstream
	}
	return nil
}

type CreateUpstreamReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *CreateUpstreamReply) Reset()                    { *m = CreateUpstreamReply{} }
func (m *CreateUpstreamReply) String() string            { return proto.CompactTextString(m) }
func (*CreateUpstreamReply) ProtoMessage()               {}
func (*CreateUpstreamReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *CreateUpstreamReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type GetUpstreamRequest struct {
	Ident string `protobuf:"bytes,1,opt,name=ident" json:"ident,omitempty"`
}

func (m *GetUpstreamRequest) Reset()                    { *m = GetUpstreamRequest{} }
func (m *GetUpstreamRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUpstreamRequest) ProtoMessage()               {}
func (*GetUpstreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *GetUpstreamRequest) GetIdent() string {
	if m != nil {
		return m.Ident
	}
	return ""
}

type GetUpstreamReply struct {
	Upstream *Upstream `protobuf:"bytes,1,opt,name=upstream" json:"upstream,omitempty"`
	Laws     []*LawSet `protobuf:"bytes,2,rep,name=laws" json:"laws,omitempty"`
	Users    []*User   `protobuf:"bytes,3,rep,name=users" json:"users,omitempty"`
	Err      string    `protobuf:"bytes,4,opt,name=err" json:"err,omitempty"`
}

func (m *GetUpstreamReply) Reset()                    { *m = GetUpstreamReply{} }
func (m *GetUpstreamReply) String() string            { return proto.CompactTextString(m) }
func (*GetUpstreamReply) ProtoMessage()               {}
func (*GetUpstreamReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *GetUpstreamReply) GetUpstream() *Upstream {
	if m != nil {
		return m.Upstream
	}
	return nil
}

func (m *GetUpstreamReply) GetLaws() []*LawSet {
	if m != nil {
		return m.Laws
	}
	return nil
}

func (m *GetUpstreamReply) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *GetUpstreamReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ListUpstreamsRequest struct {
}

func (m *ListUpstreamsRequest) Reset()                    { *m = ListUpstreamsRequest{} }
func (m *ListUpstreamsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListUpstreamsRequest) ProtoMessage()               {}
func (*ListUpstreamsRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

type ListUpstreamsReply struct {
	Upstreams []*Upstream `protobuf:"bytes,1,rep,name=upstreams" json:"upstreams,omitempty"`
	Err       string      `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *ListUpstreamsReply) Reset()                    { *m = ListUpstreamsReply{} }
func (m *ListUpstreamsReply) String() string            { return proto.CompactTextString(m) }
func (*ListUpstreamsReply) ProtoMessage()               {}
func (*ListUpstreamsReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *ListUpstreamsReply) GetUpstreams() []*Upstream {
	if m != nil {
		return m.Upstreams
	}
	return nil
}

func (m *ListUpstreamsReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type UpdateUpstreamRequest struct {
	Upstream *Upstream `protobuf:"bytes,1,opt,name=upstream" json:"upstream,omitempty"`
}

func (m *UpdateUpstreamRequest) Reset()                    { *m = UpdateUpstreamRequest{} }
func (m *UpdateUpstreamRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUpstreamRequest) ProtoMessage()               {}
func (*UpdateUpstreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *UpdateUpstreamRequest) GetUpstream() *Upstream {
	if m != nil {
		return m.Upstream
	}
	return nil
}

type UpdateUpstreamReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *UpdateUpstreamReply) Reset()                    { *m = UpdateUpstreamReply{} }
func (m *UpdateUpstreamReply) String() string            { return proto.CompactTextString(m) }
func (*UpdateUpstreamReply) ProtoMessage()               {}
func (*UpdateUpstreamReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *UpdateUpstreamReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Upstream)(nil), "grpc.welaw.v1.Upstream")
	proto.RegisterType((*UpstreamMetadata)(nil), "grpc.welaw.v1.UpstreamMetadata")
	proto.RegisterType((*CreateUpstreamRequest)(nil), "grpc.welaw.v1.CreateUpstreamRequest")
	proto.RegisterType((*CreateUpstreamReply)(nil), "grpc.welaw.v1.CreateUpstreamReply")
	proto.RegisterType((*GetUpstreamRequest)(nil), "grpc.welaw.v1.GetUpstreamRequest")
	proto.RegisterType((*GetUpstreamReply)(nil), "grpc.welaw.v1.GetUpstreamReply")
	proto.RegisterType((*ListUpstreamsRequest)(nil), "grpc.welaw.v1.ListUpstreamsRequest")
	proto.RegisterType((*ListUpstreamsReply)(nil), "grpc.welaw.v1.ListUpstreamsReply")
	proto.RegisterType((*UpdateUpstreamRequest)(nil), "grpc.welaw.v1.UpdateUpstreamRequest")
	proto.RegisterType((*UpdateUpstreamReply)(nil), "grpc.welaw.v1.UpdateUpstreamReply")
}

func init() { proto.RegisterFile("api/v1/upstream.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcf, 0x8e, 0x12, 0x4f,
	0x10, 0xce, 0x2c, 0x2c, 0x7f, 0x0a, 0xf8, 0x85, 0x5f, 0x03, 0x6e, 0x87, 0x68, 0xc4, 0xf1, 0x20,
	0xeb, 0x81, 0x0d, 0x6c, 0x3c, 0x79, 0x5c, 0x13, 0x2f, 0xe8, 0x61, 0x94, 0xa3, 0x21, 0xbd, 0x4c,
	0x2d, 0xe9, 0xec, 0xc0, 0x8c, 0xdd, 0xcd, 0x4e, 0x7c, 0x10, 0x9f, 0xc3, 0x17, 0xf0, 0xe1, 0x4c,
	0x17, 0xdd, 0x0c, 0x3b, 0x59, 0x2e, 0xc6, 0x5b, 0x57, 0xd5, 0x57, 0x7f, 0xbe, 0xfa, 0x2a, 0x0d,
	0x03, 0x91, 0xc9, 0xab, 0x87, 0xe9, 0xd5, 0x2e, 0xd3, 0x46, 0xa1, 0xd8, 0x4c, 0x32, 0x95, 0x9a,
	0x94, 0x75, 0xd6, 0x2a, 0x5b, 0x4d, 0x72, 0x4c, 0x44, 0x3e, 0x79, 0x98, 0x0e, 0xbb, 0x0e, 0x65,
	0x4d, 0x02, 0x0c, 0xff, 0xf7, 0x79, 0x1a, 0xd5, 0xde, 0x15, 0xfe, 0xae, 0x40, 0x63, 0xe1, 0xca,
	0xb0, 0x3e, 0x9c, 0xcb, 0x18, 0xb7, 0x86, 0x07, 0xa3, 0x60, 0xdc, 0x8c, 0xf6, 0x06, 0x7b, 0x0d,
	0x1d, 0xdf, 0x68, 0xb9, 0x15, 0x1b, 0xe4, 0x67, 0x14, 0x6d, 0x7b, 0xe7, 0x67, 0xb1, 0x41, 0x36,
	0x85, 0xfe, 0x01, 0x14, 0xa3, 0x5e, 0x29, 0x99, 0x19, 0x99, 0x6e, 0x79, 0x85, 0xb0, 0x3d, 0x1f,
	0xfb, 0x50, 0x84, 0x18, 0x83, 0x2a, 0x95, 0xab, 0x12, 0x84, 0xde, 0x6c, 0x04, 0xad, 0xe3, 0xec,
	0x73, 0x0a, 0x1d, 0xbb, 0x58, 0x17, 0x2a, 0x3b, 0x95, 0xf0, 0x1a, 0x45, 0xec, 0xd3, 0x4e, 0x8d,
	0x1b, 0x21, 0x13, 0x5e, 0xdf, 0x4f, 0x4d, 0x06, 0x7b, 0x05, 0xed, 0x18, 0xef, 0xc4, 0x2e, 0x31,
	0x4b, 0x4b, 0x97, 0x37, 0x7c, 0x29, 0xf2, 0x2d, 0x34, 0x2a, 0xf6, 0x02, 0x60, 0x8d, 0xe9, 0x72,
	0x95, 0xa6, 0x2a, 0xd6, 0xbc, 0x49, 0x80, 0xe6, 0x1a, 0xd3, 0x1b, 0x72, 0xb0, 0x4b, 0xa8, 0x1a,
	0xb1, 0xd6, 0x1c, 0x46, 0x95, 0x71, 0x6b, 0x36, 0x98, 0x3c, 0xda, 0xee, 0x64, 0x2e, 0xf2, 0xaf,
	0x62, 0x1d, 0x11, 0xc4, 0x52, 0x49, 0x44, 0xae, 0x79, 0x6b, 0x14, 0x8c, 0xcf, 0x23, 0x7a, 0xdb,
	0xb1, 0x6c, 0x63, 0xcd, 0xdb, 0xe4, 0xdc, 0x1b, 0xec, 0x3d, 0x34, 0x36, 0x68, 0x44, 0x2c, 0x8c,
	0xe0, 0x9d, 0x51, 0x30, 0x6e, 0xcd, 0x5e, 0x96, 0x0a, 0x7b, 0x35, 0x3e, 0x39, 0x58, 0x74, 0x48,
	0x20, 0xee, 0x32, 0xe6, 0xff, 0x39, 0xee, 0x32, 0x0e, 0x7f, 0x06, 0xd0, 0x2d, 0x27, 0x14, 0x0b,
	0x09, 0x8e, 0x17, 0xf2, 0x1c, 0x9a, 0xb9, 0xbc, 0x97, 0x19, 0xc6, 0x52, 0x38, 0x09, 0x0b, 0x07,
	0xe3, 0x50, 0x37, 0xb9, 0x34, 0x06, 0x95, 0x93, 0xcc, 0x9b, 0xec, 0x02, 0xea, 0x22, 0x93, 0x4b,
	0xbb, 0xf4, 0xbd, 0x52, 0x35, 0x91, 0xc9, 0x85, 0x4a, 0xd8, 0x10, 0x1a, 0x77, 0x62, 0x85, 0xb7,
	0x69, 0x7a, 0xef, 0x84, 0x3a, 0xd8, 0xe1, 0x1c, 0x06, 0x37, 0x0a, 0x85, 0x41, 0x3f, 0x5c, 0x84,
	0xdf, 0x77, 0xa8, 0x0d, 0xbb, 0x86, 0x86, 0xbf, 0x05, 0x1a, 0xaf, 0x35, 0xbb, 0x38, 0xc1, 0x3f,
	0x3a, 0x00, 0xc3, 0x37, 0xd0, 0x2b, 0x57, 0xcb, 0x92, 0x1f, 0x76, 0x1d, 0xa8, 0x94, 0x63, 0x69,
	0x9f, 0xe1, 0x5b, 0x60, 0x1f, 0xd1, 0x94, 0x7b, 0x3e, 0x79, 0xd6, 0xe1, 0xaf, 0x00, 0xba, 0x8f,
	0xc0, 0xb6, 0xe4, 0xdf, 0x8c, 0x67, 0x0f, 0x85, 0xd4, 0x3f, 0x3b, 0x75, 0x28, 0x5f, 0xd0, 0xb8,
	0xa3, 0xb8, 0xf4, 0x47, 0x51, 0x21, 0x6c, 0xaf, 0x5c, 0x5c, 0xa3, 0xf2, 0x97, 0xe2, 0xd8, 0x55,
	0x0b, 0x76, 0xcf, 0xa0, 0x3f, 0x97, 0xfa, 0x30, 0xb1, 0x76, 0xfc, 0xc2, 0x6f, 0xc0, 0x4a, 0x7e,
	0x4b, 0xe5, 0x1d, 0x34, 0xfd, 0x84, 0x9a, 0x07, 0xd4, 0xee, 0x24, 0x97, 0x02, 0xe9, 0xdb, 0x9e,
	0x15, 0x6d, 0xe7, 0x30, 0x58, 0x64, 0xf1, 0x3f, 0xd4, 0xb2, 0x5c, 0xed, 0x49, 0x2d, 0x6f, 0x6b,
	0xf4, 0x41, 0x5d, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x47, 0x4a, 0x43, 0xed, 0x04, 0x00,
	0x00,
}
