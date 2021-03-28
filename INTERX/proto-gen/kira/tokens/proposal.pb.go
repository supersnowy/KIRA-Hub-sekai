// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: kira/tokens/proposal.proto

package proto

import (
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MsgProposalUpsertTokenAlias struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Symbol      string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`      // Ticker (eg. ATOM, KEX, BTC)
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`          // Token Name (e.g. Cosmos, Kira, Bitcoin)
	Icon        string   `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`          // Graphical Symbol (url link to graphics)
	Decimals    uint32   `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"` // Integer number of max decimals
	Denoms      []string `protobuf:"bytes,5,rep,name=denoms,proto3" json:"denoms,omitempty"`      // An array of token denoms to be aliased
	Proposer    []byte   `protobuf:"bytes,6,opt,name=proposer,proto3" json:"proposer,omitempty"`
}

func (x *MsgProposalUpsertTokenAlias) Reset() {
	*x = MsgProposalUpsertTokenAlias{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kira_tokens_proposal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgProposalUpsertTokenAlias) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgProposalUpsertTokenAlias) ProtoMessage() {}

func (x *MsgProposalUpsertTokenAlias) ProtoReflect() protoreflect.Message {
	mi := &file_kira_tokens_proposal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgProposalUpsertTokenAlias.ProtoReflect.Descriptor instead.
func (*MsgProposalUpsertTokenAlias) Descriptor() ([]byte, []int) {
	return file_kira_tokens_proposal_proto_rawDescGZIP(), []int{0}
}

func (x *MsgProposalUpsertTokenAlias) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MsgProposalUpsertTokenAlias) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *MsgProposalUpsertTokenAlias) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MsgProposalUpsertTokenAlias) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *MsgProposalUpsertTokenAlias) GetDecimals() uint32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *MsgProposalUpsertTokenAlias) GetDenoms() []string {
	if x != nil {
		return x.Denoms
	}
	return nil
}

func (x *MsgProposalUpsertTokenAlias) GetProposer() []byte {
	if x != nil {
		return x.Proposer
	}
	return nil
}

type ProposalUpsertTokenAlias struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol   string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon     string   `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Decimals uint32   `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Denoms   []string `protobuf:"bytes,5,rep,name=denoms,proto3" json:"denoms,omitempty"`
}

func (x *ProposalUpsertTokenAlias) Reset() {
	*x = ProposalUpsertTokenAlias{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kira_tokens_proposal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalUpsertTokenAlias) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalUpsertTokenAlias) ProtoMessage() {}

func (x *ProposalUpsertTokenAlias) ProtoReflect() protoreflect.Message {
	mi := &file_kira_tokens_proposal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalUpsertTokenAlias.ProtoReflect.Descriptor instead.
func (*ProposalUpsertTokenAlias) Descriptor() ([]byte, []int) {
	return file_kira_tokens_proposal_proto_rawDescGZIP(), []int{1}
}

func (x *ProposalUpsertTokenAlias) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *ProposalUpsertTokenAlias) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProposalUpsertTokenAlias) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *ProposalUpsertTokenAlias) GetDecimals() uint32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *ProposalUpsertTokenAlias) GetDenoms() []string {
	if x != nil {
		return x.Denoms
	}
	return nil
}

type MsgProposalUpsertTokenRates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Denom       string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`                                 // denomination target
	Rate        string `protobuf:"bytes,2,opt,name=rate,proto3" json:"rate,omitempty"`                                   // Exchange rate in terms of KEX token
	FeePayments bool   `protobuf:"varint,3,opt,name=fee_payments,json=feePayments,proto3" json:"fee_payments,omitempty"` // Properties defining if it is enabled or disabled as fee payment methodß
	Proposer    []byte `protobuf:"bytes,4,opt,name=proposer,proto3" json:"proposer,omitempty"`
}

func (x *MsgProposalUpsertTokenRates) Reset() {
	*x = MsgProposalUpsertTokenRates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kira_tokens_proposal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgProposalUpsertTokenRates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgProposalUpsertTokenRates) ProtoMessage() {}

func (x *MsgProposalUpsertTokenRates) ProtoReflect() protoreflect.Message {
	mi := &file_kira_tokens_proposal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgProposalUpsertTokenRates.ProtoReflect.Descriptor instead.
func (*MsgProposalUpsertTokenRates) Descriptor() ([]byte, []int) {
	return file_kira_tokens_proposal_proto_rawDescGZIP(), []int{2}
}

func (x *MsgProposalUpsertTokenRates) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MsgProposalUpsertTokenRates) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *MsgProposalUpsertTokenRates) GetRate() string {
	if x != nil {
		return x.Rate
	}
	return ""
}

func (x *MsgProposalUpsertTokenRates) GetFeePayments() bool {
	if x != nil {
		return x.FeePayments
	}
	return false
}

func (x *MsgProposalUpsertTokenRates) GetProposer() []byte {
	if x != nil {
		return x.Proposer
	}
	return nil
}

type ProposalUpsertTokenRates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom       string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`                                 // denomination target
	Rate        string `protobuf:"bytes,2,opt,name=rate,proto3" json:"rate,omitempty"`                                   // Exchange rate in terms of KEX token
	FeePayments bool   `protobuf:"varint,3,opt,name=fee_payments,json=feePayments,proto3" json:"fee_payments,omitempty"` // Properties defining if it is enabled or disabled as fee payment methodß
}

func (x *ProposalUpsertTokenRates) Reset() {
	*x = ProposalUpsertTokenRates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kira_tokens_proposal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalUpsertTokenRates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalUpsertTokenRates) ProtoMessage() {}

func (x *ProposalUpsertTokenRates) ProtoReflect() protoreflect.Message {
	mi := &file_kira_tokens_proposal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalUpsertTokenRates.ProtoReflect.Descriptor instead.
func (*ProposalUpsertTokenRates) Descriptor() ([]byte, []int) {
	return file_kira_tokens_proposal_proto_rawDescGZIP(), []int{3}
}

func (x *ProposalUpsertTokenRates) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *ProposalUpsertTokenRates) GetRate() string {
	if x != nil {
		return x.Rate
	}
	return ""
}

func (x *ProposalUpsertTokenRates) GetFeePayments() bool {
	if x != nil {
		return x.FeePayments
	}
	return false
}

var File_kira_tokens_proposal_proto protoreflect.FileDescriptor

var file_kira_tokens_proposal_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6b, 0x69, 0x72, 0x61, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6b, 0x69,
	0x72, 0x61, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x02, 0x0a, 0x1b, 0x4d,
	0x73, 0x67, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x6e, 0x6f,
	0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x73,
	0x12, 0x60, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x44, 0xfa, 0xde, 0x1f, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0xf2, 0xde, 0x1f, 0x0f, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x22, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x72, 0x22, 0x9f, 0x01, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x65, 0x6e, 0x6f, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x6e,
	0x6f, 0x6d, 0x73, 0x3a, 0x0f, 0xd2, 0xb4, 0x2d, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0xe8, 0xa0, 0x1f, 0x01, 0x22, 0xad, 0x02, 0x0a, 0x1b, 0x4d, 0x73, 0x67, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x51, 0x0a, 0x04,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3d, 0xf2, 0xde, 0x1f, 0x0b,
	0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x72, 0x61, 0x74, 0x65, 0x22, 0xda, 0xde, 0x1f, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x63, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x66, 0x65, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x60, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x42, 0x44, 0xfa, 0xde, 0x1f, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0xf2, 0xde, 0x1f, 0x0f, 0x79, 0x61, 0x6d, 0x6c, 0x3a,
	0x22, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x22, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x72, 0x22, 0xb7, 0x01, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x51, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3d, 0xf2, 0xde, 0x1f, 0x0b, 0x79, 0x61, 0x6d, 0x6c, 0x3a,
	0x22, 0x72, 0x61, 0x74, 0x65, 0x22, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63,
	0xc8, 0xde, 0x1f, 0x00, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x65,
	0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x66, 0x65, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x0f, 0xd2,
	0xb4, 0x2d, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0xe8, 0xa0, 0x1f, 0x01, 0x42, 0x28,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x69, 0x72,
	0x61, 0x43, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x6b, 0x61, 0x69, 0x2f, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x58, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kira_tokens_proposal_proto_rawDescOnce sync.Once
	file_kira_tokens_proposal_proto_rawDescData = file_kira_tokens_proposal_proto_rawDesc
)

func file_kira_tokens_proposal_proto_rawDescGZIP() []byte {
	file_kira_tokens_proposal_proto_rawDescOnce.Do(func() {
		file_kira_tokens_proposal_proto_rawDescData = protoimpl.X.CompressGZIP(file_kira_tokens_proposal_proto_rawDescData)
	})
	return file_kira_tokens_proposal_proto_rawDescData
}

var file_kira_tokens_proposal_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kira_tokens_proposal_proto_goTypes = []interface{}{
	(*MsgProposalUpsertTokenAlias)(nil), // 0: kira.tokens.MsgProposalUpsertTokenAlias
	(*ProposalUpsertTokenAlias)(nil),    // 1: kira.tokens.ProposalUpsertTokenAlias
	(*MsgProposalUpsertTokenRates)(nil), // 2: kira.tokens.MsgProposalUpsertTokenRates
	(*ProposalUpsertTokenRates)(nil),    // 3: kira.tokens.ProposalUpsertTokenRates
}
var file_kira_tokens_proposal_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kira_tokens_proposal_proto_init() }
func file_kira_tokens_proposal_proto_init() {
	if File_kira_tokens_proposal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kira_tokens_proposal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgProposalUpsertTokenAlias); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kira_tokens_proposal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalUpsertTokenAlias); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kira_tokens_proposal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgProposalUpsertTokenRates); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kira_tokens_proposal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalUpsertTokenRates); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kira_tokens_proposal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kira_tokens_proposal_proto_goTypes,
		DependencyIndexes: file_kira_tokens_proposal_proto_depIdxs,
		MessageInfos:      file_kira_tokens_proposal_proto_msgTypes,
	}.Build()
	File_kira_tokens_proposal_proto = out.File
	file_kira_tokens_proposal_proto_rawDesc = nil
	file_kira_tokens_proposal_proto_goTypes = nil
	file_kira_tokens_proposal_proto_depIdxs = nil
}
