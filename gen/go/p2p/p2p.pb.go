// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2p/p2p.proto

package p2ppb

import (
	fmt "fmt"
	coredocument "github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type AccessType int32

const (
	AccessType_ACCESS_TYPE_INVALID                   AccessType = 0
	AccessType_ACCESS_TYPE_REQUESTER_VERIFICATION    AccessType = 1
	AccessType_ACCESS_TYPE_NFT_OWNER_VERIFICATION    AccessType = 2
	AccessType_ACCESS_TYPE_ACCESS_TOKEN_VERIFICATION AccessType = 3
)

var AccessType_name = map[int32]string{
	0: "ACCESS_TYPE_INVALID",
	1: "ACCESS_TYPE_REQUESTER_VERIFICATION",
	2: "ACCESS_TYPE_NFT_OWNER_VERIFICATION",
	3: "ACCESS_TYPE_ACCESS_TOKEN_VERIFICATION",
}

var AccessType_value = map[string]int32{
	"ACCESS_TYPE_INVALID":                   0,
	"ACCESS_TYPE_REQUESTER_VERIFICATION":    1,
	"ACCESS_TYPE_NFT_OWNER_VERIFICATION":    2,
	"ACCESS_TYPE_ACCESS_TOKEN_VERIFICATION": 3,
}

func (x AccessType) String() string {
	return proto.EnumName(AccessType_name, int32(x))
}

func (AccessType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5630330cdeb1c744, []int{0}
}

type Header struct {
	NetworkIdentifier uint32 `protobuf:"varint,1,opt,name=network_identifier,json=networkIdentifier,proto3" json:"network_identifier,omitempty"`
	NodeVersion       string `protobuf:"bytes,2,opt,name=node_version,json=nodeVersion,proto3" json:"node_version,omitempty"`
	SenderId          []byte `protobuf:"bytes,3,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// Body message type
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// sent timestamp
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_5630330cdeb1c744, []int{0}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetNetworkIdentifier() uint32 {
	if m != nil {
		return m.NetworkIdentifier
	}
	return 0
}

func (m *Header) GetNodeVersion() string {
	if m != nil {
		return m.NodeVersion
	}
	return ""
}

func (m *Header) GetSenderId() []byte {
	if m != nil {
		return m.SenderId
	}
	return nil
}

func (m *Header) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Header) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Envelope struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_5630330cdeb1c744, []int{1}
}

func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (m *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(m, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Envelope) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type SignatureRequest struct {
	Document             *coredocument.CoreDocument `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SignatureRequest) Reset()         { *m = SignatureRequest{} }
func (m *SignatureRequest) String() string { return proto.CompactTextString(m) }
func (*SignatureRequest) ProtoMessage()    {}
func (*SignatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5630330cdeb1c744, []int{2}
}

func (m *SignatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureRequest.Unmarshal(m, b)
}
func (m *SignatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureRequest.Marshal(b, m, deterministic)
}
func (m *SignatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureRequest.Merge(m, src)
}
func (m *SignatureRequest) XXX_Size() int {
	return xxx_messageInfo_SignatureRequest.Size(m)
}
func (m *SignatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureRequest proto.InternalMessageInfo

func (m *SignatureRequest) GetDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

type SignatureResponse struct {
	Signature            *coredocument.Signature `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SignatureResponse) Reset()         { *m = SignatureResponse{} }
func (m *SignatureResponse) String() string { return proto.CompactTextString(m) }
func (*SignatureResponse) ProtoMessage()    {}
func (*SignatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5630330cdeb1c744, []int{3}
}

func (m *SignatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureResponse.Unmarshal(m, b)
}
func (m *SignatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureResponse.Marshal(b, m, deterministic)
}
func (m *SignatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureResponse.Merge(m, src)
}
func (m *SignatureResponse) XXX_Size() int {
	return xxx_messageInfo_SignatureResponse.Size(m)
}
func (m *SignatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureResponse proto.InternalMessageInfo

func (m *SignatureResponse) GetSignature() *coredocument.Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AnchorDocumentRequest struct {
	Document             *coredocument.CoreDocument `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AnchorDocumentRequest) Reset()         { *m = AnchorDocumentRequest{} }
func (m *AnchorDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*AnchorDocumentRequest) ProtoMessage()    {}
func (*AnchorDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5630330cdeb1c744, []int{4}
}

func (m *AnchorDocumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnchorDocumentRequest.Unmarshal(m, b)
}
func (m *AnchorDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnchorDocumentRequest.Marshal(b, m, deterministic)
}
func (m *AnchorDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnchorDocumentRequest.Merge(m, src)
}
func (m *AnchorDocumentRequest) XXX_Size() int {
	return xxx_messageInfo_AnchorDocumentRequest.Size(m)
}
func (m *AnchorDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnchorDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnchorDocumentRequest proto.InternalMessageInfo

func (m *AnchorDocumentRequest) GetDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

type AnchorDocumentResponse struct {
	Accepted             bool     `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnchorDocumentResponse) Reset()         { *m = AnchorDocumentResponse{} }
func (m *AnchorDocumentResponse) String() string { return proto.CompactTextString(m) }
func (*AnchorDocumentResponse) ProtoMessage()    {}
func (*AnchorDocumentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5630330cdeb1c744, []int{5}
}

func (m *AnchorDocumentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnchorDocumentResponse.Unmarshal(m, b)
}
func (m *AnchorDocumentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnchorDocumentResponse.Marshal(b, m, deterministic)
}
func (m *AnchorDocumentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnchorDocumentResponse.Merge(m, src)
}
func (m *AnchorDocumentResponse) XXX_Size() int {
	return xxx_messageInfo_AnchorDocumentResponse.Size(m)
}
func (m *AnchorDocumentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnchorDocumentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnchorDocumentResponse proto.InternalMessageInfo

func (m *AnchorDocumentResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

type GetDocumentRequest struct {
	DocumentIdentifier   []byte              `protobuf:"bytes,1,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	AccessType           AccessType          `protobuf:"varint,2,opt,name=access_type,json=accessType,proto3,enum=p2p.AccessType" json:"access_type,omitempty"`
	NftRegistryAddress   []byte              `protobuf:"bytes,3,opt,name=nft_registry_address,json=nftRegistryAddress,proto3" json:"nft_registry_address,omitempty"`
	NftTokenId           []byte              `protobuf:"bytes,4,opt,name=nft_token_id,json=nftTokenId,proto3" json:"nft_token_id,omitempty"`
	AccessTokenRequest   *AccessTokenRequest `protobuf:"bytes,5,opt,name=access_token_request,json=accessTokenRequest,proto3" json:"access_token_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetDocumentRequest) Reset()         { *m = GetDocumentRequest{} }
func (m *GetDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*GetDocumentRequest) ProtoMessage()    {}
func (*GetDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5630330cdeb1c744, []int{6}
}

func (m *GetDocumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentRequest.Unmarshal(m, b)
}
func (m *GetDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentRequest.Marshal(b, m, deterministic)
}
func (m *GetDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentRequest.Merge(m, src)
}
func (m *GetDocumentRequest) XXX_Size() int {
	return xxx_messageInfo_GetDocumentRequest.Size(m)
}
func (m *GetDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentRequest proto.InternalMessageInfo

func (m *GetDocumentRequest) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *GetDocumentRequest) GetAccessType() AccessType {
	if m != nil {
		return m.AccessType
	}
	return AccessType_ACCESS_TYPE_INVALID
}

func (m *GetDocumentRequest) GetNftRegistryAddress() []byte {
	if m != nil {
		return m.NftRegistryAddress
	}
	return nil
}

func (m *GetDocumentRequest) GetNftTokenId() []byte {
	if m != nil {
		return m.NftTokenId
	}
	return nil
}

func (m *GetDocumentRequest) GetAccessTokenRequest() *AccessTokenRequest {
	if m != nil {
		return m.AccessTokenRequest
	}
	return nil
}

type AccessTokenRequest struct {
	DelegatingDocumentIdentifier []byte   `protobuf:"bytes,1,opt,name=delegating_document_identifier,json=delegatingDocumentIdentifier,proto3" json:"delegating_document_identifier,omitempty"`
	AccessTokenId                []byte   `protobuf:"bytes,2,opt,name=access_token_id,json=accessTokenId,proto3" json:"access_token_id,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *AccessTokenRequest) Reset()         { *m = AccessTokenRequest{} }
func (m *AccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*AccessTokenRequest) ProtoMessage()    {}
func (*AccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5630330cdeb1c744, []int{7}
}

func (m *AccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessTokenRequest.Unmarshal(m, b)
}
func (m *AccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessTokenRequest.Marshal(b, m, deterministic)
}
func (m *AccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTokenRequest.Merge(m, src)
}
func (m *AccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_AccessTokenRequest.Size(m)
}
func (m *AccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTokenRequest proto.InternalMessageInfo

func (m *AccessTokenRequest) GetDelegatingDocumentIdentifier() []byte {
	if m != nil {
		return m.DelegatingDocumentIdentifier
	}
	return nil
}

func (m *AccessTokenRequest) GetAccessTokenId() []byte {
	if m != nil {
		return m.AccessTokenId
	}
	return nil
}

type GetDocumentResponse struct {
	Document             *coredocument.CoreDocument `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetDocumentResponse) Reset()         { *m = GetDocumentResponse{} }
func (m *GetDocumentResponse) String() string { return proto.CompactTextString(m) }
func (*GetDocumentResponse) ProtoMessage()    {}
func (*GetDocumentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5630330cdeb1c744, []int{8}
}

func (m *GetDocumentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentResponse.Unmarshal(m, b)
}
func (m *GetDocumentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentResponse.Marshal(b, m, deterministic)
}
func (m *GetDocumentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentResponse.Merge(m, src)
}
func (m *GetDocumentResponse) XXX_Size() int {
	return xxx_messageInfo_GetDocumentResponse.Size(m)
}
func (m *GetDocumentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentResponse proto.InternalMessageInfo

func (m *GetDocumentResponse) GetDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

func init() {
	proto.RegisterEnum("p2p.AccessType", AccessType_name, AccessType_value)
	proto.RegisterType((*Header)(nil), "p2p.Header")
	proto.RegisterType((*Envelope)(nil), "p2p.Envelope")
	proto.RegisterType((*SignatureRequest)(nil), "p2p.SignatureRequest")
	proto.RegisterType((*SignatureResponse)(nil), "p2p.SignatureResponse")
	proto.RegisterType((*AnchorDocumentRequest)(nil), "p2p.AnchorDocumentRequest")
	proto.RegisterType((*AnchorDocumentResponse)(nil), "p2p.AnchorDocumentResponse")
	proto.RegisterType((*GetDocumentRequest)(nil), "p2p.GetDocumentRequest")
	proto.RegisterType((*AccessTokenRequest)(nil), "p2p.AccessTokenRequest")
	proto.RegisterType((*GetDocumentResponse)(nil), "p2p.GetDocumentResponse")
}

func init() { proto.RegisterFile("p2p/p2p.proto", fileDescriptor_5630330cdeb1c744) }

var fileDescriptor_5630330cdeb1c744 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x73, 0xd2, 0x40,
	0x14, 0x35, 0xfd, 0x12, 0x2e, 0xd4, 0xb6, 0xdb, 0x6a, 0x19, 0xec, 0x58, 0x8c, 0x63, 0x07, 0x9d,
	0x11, 0x3a, 0xf1, 0x63, 0x7c, 0xa5, 0x90, 0x6a, 0xaa, 0x02, 0x2e, 0x58, 0x47, 0x5f, 0x32, 0x81,
	0xbd, 0xd0, 0x4c, 0xcb, 0xee, 0x9a, 0x2c, 0x75, 0x78, 0xf5, 0x47, 0xf8, 0xec, 0x9f, 0xf1, 0x7f,
	0x39, 0x59, 0x12, 0x08, 0x74, 0x7c, 0xd1, 0xb7, 0xdd, 0x7b, 0xce, 0x9e, 0x9c, 0x7b, 0xef, 0x99,
	0xc0, 0xa6, 0xb4, 0x64, 0x55, 0x5a, 0xb2, 0x22, 0x03, 0xa1, 0x04, 0x59, 0x95, 0x96, 0x2c, 0x1e,
	0xf6, 0x45, 0x80, 0x4c, 0xf4, 0xc7, 0x23, 0xe4, 0xaa, 0x9a, 0xbe, 0x4c, 0x59, 0xc5, 0xc3, 0xa1,
	0x10, 0xc3, 0x2b, 0xac, 0xea, 0x5b, 0x6f, 0x3c, 0xa8, 0x2a, 0x7f, 0x84, 0xa1, 0xf2, 0x46, 0xb1,
	0x8c, 0xf9, 0xdb, 0x80, 0x8d, 0xb7, 0xe8, 0x31, 0x0c, 0xc8, 0x33, 0x20, 0x1c, 0xd5, 0x77, 0x11,
	0x5c, 0xba, 0x3e, 0x43, 0xae, 0xfc, 0x81, 0x8f, 0x41, 0xc1, 0x28, 0x19, 0xe5, 0x4d, 0xba, 0x13,
	0x23, 0xce, 0x0c, 0x20, 0x0f, 0x21, 0xcf, 0x05, 0x43, 0xf7, 0x1a, 0x83, 0xd0, 0x17, 0xbc, 0xb0,
	0x52, 0x32, 0xca, 0x59, 0x9a, 0x8b, 0x6a, 0xe7, 0xd3, 0x12, 0xb9, 0x0f, 0xd9, 0x10, 0x39, 0xc3,
	0xc0, 0xf5, 0x59, 0x61, 0xb5, 0x64, 0x94, 0xf3, 0x34, 0x33, 0x2d, 0x38, 0x8c, 0x10, 0x58, 0x53,
	0x13, 0x89, 0x85, 0x35, 0xfd, 0x4e, 0x9f, 0xc9, 0x6b, 0xc8, 0xce, 0x0c, 0x16, 0xd6, 0x4b, 0x46,
	0x39, 0x67, 0x15, 0x2b, 0xd3, 0x16, 0x2a, 0x49, 0x0b, 0x95, 0x6e, 0xc2, 0xa0, 0x73, 0xb2, 0x59,
	0x87, 0x8c, 0xcd, 0xaf, 0xf1, 0x4a, 0x48, 0x24, 0x8f, 0x60, 0xe3, 0x42, 0xb7, 0xa4, 0xcd, 0xe7,
	0xac, 0x5c, 0x25, 0x1a, 0xdb, 0xb4, 0x4b, 0x1a, 0x43, 0xd1, 0xe7, 0x7b, 0x82, 0x4d, 0xb4, 0xed,
	0x3c, 0xd5, 0x67, 0xf3, 0x0c, 0xb6, 0x3b, 0xfe, 0x90, 0x7b, 0x6a, 0x1c, 0x20, 0xc5, 0x6f, 0x63,
	0x0c, 0x15, 0x79, 0x05, 0x99, 0x64, 0xa6, 0xb1, 0x5c, 0xb1, 0xb2, 0x30, 0xe8, 0xba, 0x08, 0xb0,
	0x11, 0x5f, 0xe8, 0x8c, 0x6b, 0x9e, 0xc1, 0x4e, 0x4a, 0x2b, 0x94, 0x82, 0x87, 0x48, 0x5e, 0x42,
	0x36, 0x4c, 0x8a, 0xb1, 0xda, 0xfe, 0xa2, 0xda, 0xfc, 0xcd, 0x9c, 0x69, 0xb6, 0xe0, 0x6e, 0x8d,
	0xf7, 0x2f, 0x44, 0x30, 0xfb, 0xce, 0x7f, 0x9a, 0x7b, 0x01, 0xf7, 0x96, 0x05, 0x63, 0x87, 0x45,
	0xc8, 0x78, 0xfd, 0x3e, 0x4a, 0x85, 0x4c, 0x2b, 0x66, 0xe8, 0xec, 0x6e, 0xfe, 0x5c, 0x01, 0xf2,
	0x06, 0xd5, 0xb2, 0x89, 0x2a, 0xec, 0x26, 0xc2, 0xcb, 0xc1, 0xc9, 0x53, 0x92, 0x40, 0xa9, 0xe4,
	0x1c, 0x43, 0x2e, 0xd2, 0x0c, 0x43, 0x57, 0x07, 0x20, 0xda, 0xc0, 0x1d, 0x6b, 0x4b, 0x2f, 0xa9,
	0xa6, 0xeb, 0xdd, 0x89, 0x44, 0x0a, 0xde, 0xec, 0x4c, 0x8e, 0x61, 0x8f, 0x0f, 0x94, 0x1b, 0xe0,
	0xd0, 0x0f, 0x55, 0x30, 0x71, 0x3d, 0xc6, 0x02, 0x0c, 0xc3, 0x38, 0x53, 0x84, 0x0f, 0x14, 0x8d,
	0xa1, 0xda, 0x14, 0x21, 0x25, 0xc8, 0x47, 0x2f, 0x94, 0xb8, 0x44, 0x1e, 0xa5, 0x6f, 0x4d, 0x33,
	0x81, 0x0f, 0x54, 0x37, 0x2a, 0x39, 0x8c, 0x38, 0xb0, 0x97, 0xb8, 0xd0, 0xa4, 0x60, 0xda, 0x4e,
	0x1c, 0xbb, 0xfd, 0xb4, 0x9d, 0x08, 0x8f, 0xbb, 0xa5, 0xc4, 0xbb, 0x51, 0x33, 0x7f, 0x18, 0x40,
	0x6e, 0x52, 0x49, 0x03, 0x1e, 0x30, 0xbc, 0xc2, 0xa1, 0xa7, 0x7c, 0x3e, 0x74, 0xff, 0x3e, 0xa3,
	0x83, 0x39, 0xab, 0x71, 0x73, 0x5a, 0x47, 0xb0, 0xb5, 0xe0, 0xd3, 0x67, 0x71, 0x66, 0x37, 0x53,
	0x4e, 0x1c, 0x66, 0x7e, 0x80, 0xdd, 0x85, 0xe5, 0xc4, 0x0b, 0xfd, 0xc7, 0x88, 0x3c, 0xfd, 0x65,
	0x00, 0xcc, 0xb7, 0x41, 0xf6, 0x61, 0xb7, 0x56, 0xaf, 0xdb, 0x9d, 0x8e, 0xdb, 0xfd, 0xd2, 0xb6,
	0x5d, 0xa7, 0x79, 0x5e, 0x7b, 0xef, 0x34, 0xb6, 0x6f, 0x91, 0x23, 0x30, 0xd3, 0x00, 0xb5, 0x3f,
	0x7e, 0xb2, 0x3b, 0x5d, 0x9b, 0xba, 0xe7, 0x36, 0x75, 0x4e, 0x9d, 0x7a, 0xad, 0xeb, 0xb4, 0x9a,
	0xdb, 0xc6, 0x32, 0xaf, 0x79, 0xda, 0x75, 0x5b, 0x9f, 0x9b, 0xcb, 0xbc, 0x15, 0xf2, 0x04, 0x1e,
	0xa7, 0x79, 0xc9, 0xb9, 0xf5, 0xce, 0x6e, 0x2e, 0x52, 0x57, 0x4f, 0x0e, 0xe0, 0x76, 0x5f, 0x8c,
	0xa2, 0x45, 0x9d, 0x64, 0xda, 0x96, 0x6c, 0x47, 0x3f, 0x88, 0xb6, 0xf1, 0x75, 0x5d, 0x5a, 0x52,
	0xf6, 0x7a, 0x1b, 0xfa, 0x87, 0xf1, 0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xc5, 0xb7,
	0xdc, 0x38, 0x05, 0x00, 0x00,
}
