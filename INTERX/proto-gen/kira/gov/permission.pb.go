// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kira/gov/permission.proto

package gov

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type PermValue int32

const (
	// PERMISSION_ZERO is a no-op permission.
	PermValue_PERMISSION_ZERO PermValue = 0
	// PERMISSION_SET_PERMISSIONS defines the permission that allows to Set Permissions to other actors.
	PermValue_PERMISSION_SET_PERMISSIONS PermValue = 1
	// PERMISSION_CLAIM_VALIDATOR defines the permission that allows to Claim a validator Seat.
	PermValue_PERMISSION_CLAIM_VALIDATOR PermValue = 2
	// PERMISSION_CLAIM_COUNCILOR defines the permission that allows to Claim a Councilor Seat.
	PermValue_PERMISSION_CLAIM_COUNCILOR PermValue = 3
	// PERMISSION_WHITELIST_ACCOUNT_PERMISSION_PROPOSAL defines the permission needed to create proposals for whitelisting an account permission.
	PermValue_PERMISSION_WHITELIST_ACCOUNT_PERMISSION_PROPOSAL PermValue = 4
	// PERMISSION_VOTE_WHITELIST_ACCOUNT_PERMISSION_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to whitelist account permission.
	PermValue_PERMISSION_VOTE_WHITELIST_ACCOUNT_PERMISSION_PROPOSAL PermValue = 5
	//  PERMISSION_UPSERT_TOKEN_ALIAS
	PermValue_PERMISSION_UPSERT_TOKEN_ALIAS PermValue = 6
	// PERMISSION_CHANGE_TX_FEE
	PermValue_PERMISSION_CHANGE_TX_FEE PermValue = 7
	// PERMISSION_UPSERT_TOKEN_RATE
	PermValue_PERMISSION_UPSERT_TOKEN_RATE PermValue = 8
	// PERMISSION_UPSERT_ROLE makes possible to add, modify and assign roles.
	PermValue_PERMISSION_UPSERT_ROLE PermValue = 9
	// PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermValue_PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL PermValue = 10
	// PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermValue_PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL PermValue = 11
	// PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission needed to create proposals for setting network property.
	PermValue_PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL PermValue = 12
	// PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set network property.
	PermValue_PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL PermValue = 13
	// PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to create proposals for upsert token Alias.
	PermValue_PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL PermValue = 14
	// PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to vote proposals for upsert token.
	PermValue_PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL PermValue = 15
	// PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES defines the permission needed to create proposals for setting poor network messages
	PermValue_PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES PermValue = 16
	// PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL defines the permission needed to vote proposals to set poor network messages
	PermValue_PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL PermValue = 17
	// PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL defines the permission needed to create proposals for upsert token rate.
	PermValue_PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL PermValue = 18
	// PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL defines the permission needed to vote proposals for upsert token rate.
	PermValue_PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL PermValue = 19
	// PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL defines the permission needed to create a proposal to unjail a validator.
	PermValue_PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL PermValue = 20
	// PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL defines the permission needed to vote a proposal to unjail a validator.
	PermValue_PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL PermValue = 21
	// PERMISSION_CREATE_CREATE_ROLE_PROPOSAL defines the permission needed to create a proposal to create a role.
	PermValue_PERMISSION_CREATE_CREATE_ROLE_PROPOSAL PermValue = 22
	// PERMISSION_VOTE_CREATE_ROLE_PROPOSAL defines the permission needed to vote a proposal to create a role.
	PermValue_PERMISSION_VOTE_CREATE_ROLE_PROPOSAL PermValue = 23
	// PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL defines the permission needed to create a proposal to blacklist/whitelisted tokens
	PermValue_PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL PermValue = 24
	// PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL defines the permission needed to vote on blacklist/whitelisted tokens proposal
	PermValue_PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL PermValue = 25
	// PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL defines the permission needed to create a proposal to reset whole validator rank
	PermValue_PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL PermValue = 26
	// PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL defines the permission needed to vote on reset whole validator rank proposal
	PermValue_PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL PermValue = 27
	// PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL defines the permission needed to create a proposal for software upgrade
	PermValue_PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL PermValue = 28
	// PERMISSION_SOFTWARE_UPGRADE_PROPOSAL defines the permission needed to vote on software upgrade proposal
	PermValue_PERMISSION_SOFTWARE_UPGRADE_PROPOSAL PermValue = 29
	// PERMISSION_SET_PERMISSIONS defines the permission that allows to Set ClaimValidatorPermission to other actors.
	PermValue_PERMISSION_SET_CLAIM_VALIDATOR_PERMISSION PermValue = 30
	// PERMISSION_CREATE_SET_PROPOSAL_DURATION_PROPOSAL defines the permission needed to create a proposal to set proposal duration.
	PermValue_PERMISSION_CREATE_SET_PROPOSAL_DURATION_PROPOSAL PermValue = 31
	// PERMISSION_VOTE_SET_PROPOSAL_DURATION_PROPOSAL defines the permission needed to vote a proposal to set proposal duration.
	PermValue_PERMISSION_VOTE_SET_PROPOSAL_DURATION_PROPOSAL PermValue = 32
	// PERMISSION_BLACKLIST_ACCOUNT_PERMISSION_PROPOSAL defines the permission needed to create proposals for blacklisting an account permission.
	PermValue_PERMISSION_BLACKLIST_ACCOUNT_PERMISSION_PROPOSAL PermValue = 33
	// PERMISSION_VOTE_BLACKLIST_ACCOUNT_PERMISSION_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to blacklist account permission.
	PermValue_PERMISSION_VOTE_BLACKLIST_ACCOUNT_PERMISSION_PROPOSAL PermValue = 34
	// PERMISSION_REMOVE_WHITELISTED_ACCOUNT_PERMISSION_PROPOSAL defines the permission needed to create proposals for removing whitelisted permission from an account.
	PermValue_PERMISSION_REMOVE_WHITELISTED_ACCOUNT_PERMISSION_PROPOSAL PermValue = 35
	// PERMISSION_VOTE_REMOVE_WHITELISTED_ACCOUNT_PERMISSION_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to remove a whitelisted account permission.
	PermValue_PERMISSION_VOTE_REMOVE_WHITELISTED_ACCOUNT_PERMISSION_PROPOSAL PermValue = 36
	// PERMISSION_REMOVE_BLACKLISTED_ACCOUNT_PERMISSION_PROPOSAL defines the permission needed to create proposals for removing blacklisted permission from an account.
	PermValue_PERMISSION_REMOVE_BLACKLISTED_ACCOUNT_PERMISSION_PROPOSAL PermValue = 37
	// PERMISSION_VOTE_REMOVE_BLACKLISTED_ACCOUNT_PERMISSION_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to remove a blacklisted account permission.
	PermValue_PERMISSION_VOTE_REMOVE_BLACKLISTED_ACCOUNT_PERMISSION_PROPOSAL PermValue = 38
	// PERMISSION_WHITELIST_ROLE_PERMISSION_PROPOSAL defines the permission needed to create proposals for whitelisting an role permission.
	PermValue_PERMISSION_WHITELIST_ROLE_PERMISSION_PROPOSAL PermValue = 39
	// PERMISSION_VOTE_WHITELIST_ROLE_PERMISSION_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to whitelist role permission.
	PermValue_PERMISSION_VOTE_WHITELIST_ROLE_PERMISSION_PROPOSAL PermValue = 40
	// PERMISSION_BLACKLIST_ROLE_PERMISSION_PROPOSAL defines the permission needed to create proposals for blacklisting an role permission.
	PermValue_PERMISSION_BLACKLIST_ROLE_PERMISSION_PROPOSAL PermValue = 41
	// PERMISSION_VOTE_BLACKLIST_ROLE_PERMISSION_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to blacklist role permission.
	PermValue_PERMISSION_VOTE_BLACKLIST_ROLE_PERMISSION_PROPOSAL PermValue = 42
	// PERMISSION_REMOVE_WHITELISTED_ROLE_PERMISSION_PROPOSAL defines the permission needed to create proposals for removing whitelisted permission from a role.
	PermValue_PERMISSION_REMOVE_WHITELISTED_ROLE_PERMISSION_PROPOSAL PermValue = 43
	// PERMISSION_VOTE_REMOVE_WHITELISTED_ROLE_PERMISSION_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to remove a whitelisted role permission.
	PermValue_PERMISSION_VOTE_REMOVE_WHITELISTED_ROLE_PERMISSION_PROPOSAL PermValue = 44
	// PERMISSION_REMOVE_BLACKLISTED_ROLE_PERMISSION_PROPOSAL defines the permission needed to create proposals for removing blacklisted permission from a role.
	PermValue_PERMISSION_REMOVE_BLACKLISTED_ROLE_PERMISSION_PROPOSAL PermValue = 45
	// PERMISSION_VOTE_REMOVE_BLACKLISTED_ROLE_PERMISSION_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to remove a blacklisted role permission.
	PermValue_PERMISSION_VOTE_REMOVE_BLACKLISTED_ROLE_PERMISSION_PROPOSAL PermValue = 46
	// PERMISSION_ASSIGN_ROLE_TO_ACCOUNT_PROPOSAL defines the permission needed to create proposals to assign role to an account
	PermValue_PERMISSION_ASSIGN_ROLE_TO_ACCOUNT_PROPOSAL PermValue = 47
	// PERMISSION_VOTE_ASSIGN_ROLE_TO_ACCOUNT_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to assign role to an account
	PermValue_PERMISSION_VOTE_ASSIGN_ROLE_TO_ACCOUNT_PROPOSAL PermValue = 48
	// PERMISSION_UNASSIGN_ROLE_FROM_ACCOUNT_PROPOSAL defines the permission needed to create proposals to unassign role from an account
	PermValue_PERMISSION_UNASSIGN_ROLE_FROM_ACCOUNT_PROPOSAL PermValue = 49
	// PERMISSION_VOTE_UNASSIGN_ROLE_FROM_ACCOUNT_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to unassign role from an account
	PermValue_PERMISSION_VOTE_UNASSIGN_ROLE_FROM_ACCOUNT_PROPOSAL PermValue = 50
	// PERMISSION_CREATE_REMOVE_ROLE_PROPOSAL defines the permission needed to create a proposal to remove a role.
	PermValue_PERMISSION_CREATE_REMOVE_ROLE_PROPOSAL PermValue = 51
	// PERMISSION_VOTE_REMOVE_ROLE_PROPOSAL defines the permission needed to vote a proposal to remove a role.
	PermValue_PERMISSION_VOTE_REMOVE_ROLE_PROPOSAL PermValue = 52
)

var PermValue_name = map[int32]string{
	0:  "PERMISSION_ZERO",
	1:  "PERMISSION_SET_PERMISSIONS",
	2:  "PERMISSION_CLAIM_VALIDATOR",
	3:  "PERMISSION_CLAIM_COUNCILOR",
	4:  "PERMISSION_WHITELIST_ACCOUNT_PERMISSION_PROPOSAL",
	5:  "PERMISSION_VOTE_WHITELIST_ACCOUNT_PERMISSION_PROPOSAL",
	6:  "PERMISSION_UPSERT_TOKEN_ALIAS",
	7:  "PERMISSION_CHANGE_TX_FEE",
	8:  "PERMISSION_UPSERT_TOKEN_RATE",
	9:  "PERMISSION_UPSERT_ROLE",
	10: "PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL",
	11: "PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL",
	12: "PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL",
	13: "PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL",
	14: "PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL",
	15: "PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL",
	16: "PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES",
	17: "PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL",
	18: "PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL",
	19: "PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL",
	20: "PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL",
	21: "PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL",
	22: "PERMISSION_CREATE_CREATE_ROLE_PROPOSAL",
	23: "PERMISSION_VOTE_CREATE_ROLE_PROPOSAL",
	24: "PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL",
	25: "PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL",
	26: "PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL",
	27: "PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL",
	28: "PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL",
	29: "PERMISSION_SOFTWARE_UPGRADE_PROPOSAL",
	30: "PERMISSION_SET_CLAIM_VALIDATOR_PERMISSION",
	31: "PERMISSION_CREATE_SET_PROPOSAL_DURATION_PROPOSAL",
	32: "PERMISSION_VOTE_SET_PROPOSAL_DURATION_PROPOSAL",
	33: "PERMISSION_BLACKLIST_ACCOUNT_PERMISSION_PROPOSAL",
	34: "PERMISSION_VOTE_BLACKLIST_ACCOUNT_PERMISSION_PROPOSAL",
	35: "PERMISSION_REMOVE_WHITELISTED_ACCOUNT_PERMISSION_PROPOSAL",
	36: "PERMISSION_VOTE_REMOVE_WHITELISTED_ACCOUNT_PERMISSION_PROPOSAL",
	37: "PERMISSION_REMOVE_BLACKLISTED_ACCOUNT_PERMISSION_PROPOSAL",
	38: "PERMISSION_VOTE_REMOVE_BLACKLISTED_ACCOUNT_PERMISSION_PROPOSAL",
	39: "PERMISSION_WHITELIST_ROLE_PERMISSION_PROPOSAL",
	40: "PERMISSION_VOTE_WHITELIST_ROLE_PERMISSION_PROPOSAL",
	41: "PERMISSION_BLACKLIST_ROLE_PERMISSION_PROPOSAL",
	42: "PERMISSION_VOTE_BLACKLIST_ROLE_PERMISSION_PROPOSAL",
	43: "PERMISSION_REMOVE_WHITELISTED_ROLE_PERMISSION_PROPOSAL",
	44: "PERMISSION_VOTE_REMOVE_WHITELISTED_ROLE_PERMISSION_PROPOSAL",
	45: "PERMISSION_REMOVE_BLACKLISTED_ROLE_PERMISSION_PROPOSAL",
	46: "PERMISSION_VOTE_REMOVE_BLACKLISTED_ROLE_PERMISSION_PROPOSAL",
	47: "PERMISSION_ASSIGN_ROLE_TO_ACCOUNT_PROPOSAL",
	48: "PERMISSION_VOTE_ASSIGN_ROLE_TO_ACCOUNT_PROPOSAL",
	49: "PERMISSION_UNASSIGN_ROLE_FROM_ACCOUNT_PROPOSAL",
	50: "PERMISSION_VOTE_UNASSIGN_ROLE_FROM_ACCOUNT_PROPOSAL",
	51: "PERMISSION_CREATE_REMOVE_ROLE_PROPOSAL",
	52: "PERMISSION_VOTE_REMOVE_ROLE_PROPOSAL",
}

var PermValue_value = map[string]int32{
	"PERMISSION_ZERO":                                                0,
	"PERMISSION_SET_PERMISSIONS":                                     1,
	"PERMISSION_CLAIM_VALIDATOR":                                     2,
	"PERMISSION_CLAIM_COUNCILOR":                                     3,
	"PERMISSION_WHITELIST_ACCOUNT_PERMISSION_PROPOSAL":               4,
	"PERMISSION_VOTE_WHITELIST_ACCOUNT_PERMISSION_PROPOSAL":          5,
	"PERMISSION_UPSERT_TOKEN_ALIAS":                                  6,
	"PERMISSION_CHANGE_TX_FEE":                                       7,
	"PERMISSION_UPSERT_TOKEN_RATE":                                   8,
	"PERMISSION_UPSERT_ROLE":                                         9,
	"PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL":                10,
	"PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL":                  11,
	"PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL":                12,
	"PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL":                  13,
	"PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL":                  14,
	"PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL":                    15,
	"PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES":                    16,
	"PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL":             17,
	"PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL":                   18,
	"PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL":                     19,
	"PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL":                    20,
	"PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL":                      21,
	"PERMISSION_CREATE_CREATE_ROLE_PROPOSAL":                         22,
	"PERMISSION_VOTE_CREATE_ROLE_PROPOSAL":                           23,
	"PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL":           24,
	"PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL":             25,
	"PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL":          26,
	"PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL":            27,
	"PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL":                    28,
	"PERMISSION_SOFTWARE_UPGRADE_PROPOSAL":                           29,
	"PERMISSION_SET_CLAIM_VALIDATOR_PERMISSION":                      30,
	"PERMISSION_CREATE_SET_PROPOSAL_DURATION_PROPOSAL":               31,
	"PERMISSION_VOTE_SET_PROPOSAL_DURATION_PROPOSAL":                 32,
	"PERMISSION_BLACKLIST_ACCOUNT_PERMISSION_PROPOSAL":               33,
	"PERMISSION_VOTE_BLACKLIST_ACCOUNT_PERMISSION_PROPOSAL":          34,
	"PERMISSION_REMOVE_WHITELISTED_ACCOUNT_PERMISSION_PROPOSAL":      35,
	"PERMISSION_VOTE_REMOVE_WHITELISTED_ACCOUNT_PERMISSION_PROPOSAL": 36,
	"PERMISSION_REMOVE_BLACKLISTED_ACCOUNT_PERMISSION_PROPOSAL":      37,
	"PERMISSION_VOTE_REMOVE_BLACKLISTED_ACCOUNT_PERMISSION_PROPOSAL": 38,
	"PERMISSION_WHITELIST_ROLE_PERMISSION_PROPOSAL":                  39,
	"PERMISSION_VOTE_WHITELIST_ROLE_PERMISSION_PROPOSAL":             40,
	"PERMISSION_BLACKLIST_ROLE_PERMISSION_PROPOSAL":                  41,
	"PERMISSION_VOTE_BLACKLIST_ROLE_PERMISSION_PROPOSAL":             42,
	"PERMISSION_REMOVE_WHITELISTED_ROLE_PERMISSION_PROPOSAL":         43,
	"PERMISSION_VOTE_REMOVE_WHITELISTED_ROLE_PERMISSION_PROPOSAL":    44,
	"PERMISSION_REMOVE_BLACKLISTED_ROLE_PERMISSION_PROPOSAL":         45,
	"PERMISSION_VOTE_REMOVE_BLACKLISTED_ROLE_PERMISSION_PROPOSAL":    46,
	"PERMISSION_ASSIGN_ROLE_TO_ACCOUNT_PROPOSAL":                     47,
	"PERMISSION_VOTE_ASSIGN_ROLE_TO_ACCOUNT_PROPOSAL":                48,
	"PERMISSION_UNASSIGN_ROLE_FROM_ACCOUNT_PROPOSAL":                 49,
	"PERMISSION_VOTE_UNASSIGN_ROLE_FROM_ACCOUNT_PROPOSAL":            50,
	"PERMISSION_CREATE_REMOVE_ROLE_PROPOSAL":                         51,
	"PERMISSION_VOTE_REMOVE_ROLE_PROPOSAL":                           52,
}

func (x PermValue) String() string {
	return proto.EnumName(PermValue_name, int32(x))
}

func (PermValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_214168f8815c1062, []int{0}
}

func init() {
	proto.RegisterEnum("kira.gov.PermValue", PermValue_name, PermValue_value)
}

func init() {
	proto.RegisterFile("kira/gov/permission.proto", fileDescriptor_214168f8815c1062)
}

var fileDescriptor_214168f8815c1062 = []byte{
	// 1321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x98, 0xef, 0x72, 0xd3, 0x46,
	0x17, 0xc6, 0x79, 0xdf, 0x52, 0x0a, 0x5b, 0x0a, 0xae, 0xa0, 0x01, 0xb6, 0xfc, 0x59, 0x20, 0x09,
	0x24, 0x21, 0x16, 0x24, 0xc0, 0x0c, 0xc3, 0x4c, 0x67, 0x84, 0xad, 0x24, 0x6a, 0x6c, 0xcb, 0xb3,
	0x92, 0x63, 0x48, 0xcb, 0xa8, 0x22, 0x59, 0x1c, 0xd5, 0x8e, 0x37, 0x95, 0x14, 0x68, 0x7b, 0x05,
	0x9d, 0xbd, 0x87, 0xfd, 0xd4, 0x4f, 0xbd, 0xcb, 0x8e, 0x64, 0x5b, 0xab, 0x95, 0x25, 0x5b, 0xe9,
	0xa7, 0xc4, 0xb6, 0xf6, 0xf7, 0x9c, 0xf3, 0xec, 0x39, 0xbb, 0xc7, 0x06, 0xb7, 0xfa, 0x9e, 0xef,
	0xaa, 0x3d, 0xfa, 0x49, 0x3d, 0x21, 0xfe, 0xb1, 0x17, 0x04, 0x1e, 0x1d, 0x56, 0x4f, 0x7c, 0x1a,
	0x52, 0xe5, 0x62, 0xf4, 0x51, 0xb5, 0x47, 0x3f, 0xc1, 0xeb, 0x3d, 0xda, 0xa3, 0xf1, 0x9b, 0x6a,
	0xf4, 0xdf, 0xe8, 0xf3, 0xd5, 0x7f, 0x96, 0xc0, 0xa5, 0x36, 0xf1, 0x8f, 0xf7, 0xdc, 0xc1, 0x29,
	0x51, 0xee, 0x83, 0xab, 0x6d, 0x1d, 0x37, 0x0d, 0xcb, 0x32, 0xcc, 0x96, 0xb3, 0xaf, 0x63, 0xb3,
	0x72, 0x0e, 0x5e, 0x66, 0x1c, 0x5d, 0x8c, 0x9e, 0xd9, 0x27, 0x3e, 0x55, 0x5e, 0x02, 0x98, 0x7a,
	0xc4, 0xd2, 0x6d, 0x47, 0xbc, 0xb4, 0x2a, 0xff, 0x83, 0x0b, 0x8c, 0x23, 0x25, 0x7a, 0xda, 0x22,
	0x61, 0x3b, 0x89, 0x26, 0xc8, 0xac, 0xab, 0x35, 0x34, 0xa3, 0xe9, 0xec, 0x69, 0x0d, 0xa3, 0xae,
	0xd9, 0x26, 0xae, 0xfc, 0x5f, 0xac, 0xab, 0x0d, 0x5c, 0x2f, 0x0a, 0xc7, 0x3b, 0x74, 0x43, 0xea,
	0xe7, 0xae, 0xab, 0x99, 0x9d, 0x56, 0xcd, 0x68, 0x98, 0xb8, 0xf2, 0x45, 0x66, 0x5d, 0x8d, 0x9e,
	0x0e, 0x0f, 0xbc, 0x01, 0xf5, 0x95, 0x5f, 0xc0, 0xd3, 0xd4, 0xba, 0xee, 0x8e, 0x61, 0xeb, 0x0d,
	0xc3, 0xb2, 0x1d, 0xad, 0x16, 0xad, 0x4e, 0x47, 0xed, 0xb4, 0xb1, 0xd9, 0x36, 0x2d, 0xad, 0x51,
	0x39, 0x0f, 0x57, 0x19, 0x47, 0xcb, 0x11, 0xad, 0x7b, 0xe4, 0x85, 0x64, 0xe0, 0x05, 0xa1, 0x76,
	0x70, 0x40, 0x4f, 0x87, 0xa9, 0x54, 0xda, 0x3e, 0x3d, 0xa1, 0x81, 0x3b, 0x50, 0x3c, 0xf0, 0x22,
	0x05, 0xd9, 0x33, 0x6d, 0xbd, 0xa4, 0xcc, 0x97, 0xb0, 0xca, 0x38, 0x5a, 0x8d, 0x6d, 0xa7, 0x21,
	0x29, 0x21, 0xf5, 0x1a, 0xdc, 0x49, 0x81, 0x3a, 0x6d, 0x4b, 0xc7, 0xb6, 0x63, 0x9b, 0xbb, 0x7a,
	0xcb, 0xd1, 0x1a, 0x86, 0x66, 0x55, 0x2e, 0xc0, 0x9b, 0x8c, 0xa3, 0xeb, 0xd1, 0xd2, 0xce, 0x49,
	0x40, 0xfc, 0xd0, 0xa6, 0x7d, 0x32, 0xd4, 0x06, 0x9e, 0x1b, 0x28, 0xcf, 0xc0, 0xcd, 0xb4, 0x83,
	0x3b, 0x5a, 0x6b, 0x5b, 0x77, 0xec, 0xb7, 0xce, 0x96, 0xae, 0x57, 0xbe, 0x82, 0xd7, 0x18, 0x47,
	0x57, 0x63, 0xff, 0x8e, 0xdc, 0x61, 0x8f, 0xd8, 0xbf, 0x6f, 0x11, 0xa2, 0xbc, 0x02, 0xb7, 0x8b,
	0xf4, 0xb0, 0x66, 0xeb, 0x95, 0x8b, 0xf0, 0x06, 0xe3, 0xe8, 0x5a, 0x46, 0x0e, 0xbb, 0x21, 0x51,
	0xaa, 0x60, 0x61, 0x7a, 0x29, 0x36, 0x1b, 0x7a, 0xe5, 0x12, 0x54, 0x18, 0x47, 0x57, 0xc4, 0x22,
	0x4c, 0x07, 0x44, 0x79, 0x0f, 0xd4, 0x74, 0x74, 0x58, 0xd7, 0x6c, 0x7d, 0xb2, 0xac, 0xae, 0xd9,
	0x9a, 0x83, 0xf5, 0x6d, 0xc3, 0xb2, 0xf1, 0x3b, 0xe1, 0x1f, 0x80, 0x8f, 0x19, 0x47, 0x8b, 0x71,
	0xd0, 0x3e, 0x71, 0x43, 0x32, 0xc2, 0xd5, 0xdd, 0xd0, 0xc5, 0xa4, 0xe7, 0x05, 0xa1, 0xff, 0x47,
	0xe2, 0xdc, 0x3b, 0xb0, 0x9e, 0xdd, 0xa4, 0xd9, 0xf0, 0xaf, 0xe1, 0x32, 0xe3, 0xe8, 0xc1, 0x64,
	0x73, 0x66, 0xa0, 0x73, 0x23, 0x8f, 0x1a, 0xa2, 0xa5, 0xdb, 0x5d, 0x13, 0xef, 0xc6, 0x4c, 0x1d,
	0xdb, 0x29, 0xf8, 0xe5, 0x6c, 0xe4, 0x16, 0x09, 0x5b, 0x24, 0xfc, 0x4c, 0xfd, 0x7e, 0x84, 0x25,
	0x7e, 0x38, 0x33, 0xf2, 0xd9, 0xf0, 0x6f, 0xe4, 0xc8, 0x4b, 0xa3, 0x65, 0xcf, 0x53, 0x55, 0x25,
	0xd0, 0x57, 0x04, 0x3a, 0xed, 0xb8, 0x28, 0xb2, 0x04, 0xdd, 0x01, 0x6b, 0x05, 0x7e, 0xe7, 0x82,
	0xaf, 0xc2, 0x45, 0xc6, 0x11, 0x92, 0xdd, 0xce, 0xc1, 0xbe, 0x97, 0xb0, 0x29, 0xaf, 0xdb, 0xa6,
	0x89, 0x13, 0x4f, 0x9a, 0xba, 0x65, 0x69, 0xdb, 0xba, 0x55, 0xa9, 0xc0, 0x27, 0x8c, 0xa3, 0xc7,
	0x92, 0xcf, 0x6d, 0x4a, 0xfd, 0xb1, 0x21, 0x4d, 0x12, 0x04, 0x6e, 0x8f, 0x08, 0xfc, 0x07, 0xb0,
	0x91, 0xe7, 0x75, 0x2e, 0x5c, 0x04, 0xff, 0xad, 0x38, 0x2e, 0xc6, 0x86, 0xcf, 0xd2, 0xe8, 0x82,
	0x27, 0x73, 0x4c, 0x8f, 0x5a, 0x4b, 0xd0, 0x15, 0xb8, 0xc4, 0x38, 0xba, 0x9f, 0xeb, 0x79, 0xd4,
	0x69, 0x09, 0xd8, 0x02, 0xab, 0x33, 0x2d, 0x97, 0xb1, 0xd7, 0xe0, 0x43, 0xc6, 0xd1, 0xbd, 0x1c,
	0xc7, 0x25, 0xe8, 0x5e, 0x9e, 0xe1, 0x9d, 0xd6, 0x8f, 0x9a, 0xd1, 0x10, 0xc7, 0xb6, 0xa0, 0x5e,
	0x9f, 0x0a, 0x76, 0xf8, 0xab, 0xeb, 0x0d, 0x92, 0x63, 0x3c, 0xe1, 0x62, 0xb0, 0x32, 0x15, 0x6c,
	0x21, 0xf5, 0xbb, 0x4c, 0xac, 0x05, 0xcc, 0x2d, 0xb0, 0x3c, 0x1d, 0xeb, 0xf8, 0x4f, 0x74, 0xf2,
	0x08, 0xe0, 0x02, 0x84, 0x8c, 0xa3, 0x05, 0x11, 0x66, 0x74, 0x04, 0x25, 0x9c, 0x1d, 0xb0, 0x98,
	0x8d, 0x2d, 0x97, 0x72, 0x03, 0xde, 0x65, 0x1c, 0xc1, 0x49, 0x58, 0x39, 0xa4, 0x8f, 0xe0, 0xf9,
	0x74, 0x44, 0xf1, 0x6e, 0x58, 0xa3, 0x3b, 0xc2, 0x79, 0xd3, 0xd0, 0x6a, 0xbb, 0x93, 0xc3, 0x38,
	0x21, 0xdf, 0xcc, 0xd6, 0x6d, 0xbc, 0x31, 0x41, 0x7c, 0x43, 0xbc, 0x19, 0xb8, 0x07, 0xfd, 0xd1,
	0x21, 0x3d, 0xab, 0x6e, 0x4b, 0xa8, 0xdc, 0x92, 0xeb, 0x76, 0x8e, 0xc6, 0x91, 0x74, 0xcd, 0x4d,
	0x0c, 0xd1, 0xa3, 0xfe, 0xe8, 0xee, 0x44, 0xbe, 0x88, 0x8d, 0xc3, 0x5a, 0x6b, 0x57, 0xc8, 0x40,
	0xb8, 0xce, 0x38, 0x5a, 0x49, 0x99, 0x4d, 0x02, 0x12, 0x76, 0x8f, 0xe8, 0x80, 0x24, 0x7b, 0x88,
	0xdd, 0x61, 0x3f, 0x51, 0x3a, 0x04, 0x9b, 0xd9, 0x6c, 0xca, 0xe8, 0x7c, 0x0f, 0xd7, 0x18, 0x47,
	0x8f, 0x26, 0xe9, 0xcc, 0x53, 0xc9, 0xad, 0x6c, 0xcb, 0xdc, 0xb2, 0xbb, 0x1a, 0x8e, 0x3a, 0x67,
	0x1b, 0x6b, 0xf5, 0x94, 0x59, 0xb7, 0xb3, 0x95, 0x6d, 0xd1, 0x8f, 0xe1, 0x67, 0xd7, 0x27, 0x9d,
	0x93, 0x9e, 0xef, 0x1e, 0x0a, 0x9f, 0x9a, 0x52, 0xf5, 0x14, 0x03, 0xef, 0xc8, 0x45, 0x5d, 0x84,
	0x93, 0x1b, 0x25, 0x72, 0x21, 0x33, 0x33, 0xa5, 0x46, 0x8b, 0xca, 0x5d, 0xc1, 0xb4, 0x48, 0x28,
	0x4f, 0x50, 0x62, 0x9c, 0x50, 0x1c, 0x69, 0x26, 0x4a, 0x9f, 0xa2, 0xe3, 0xe0, 0x9c, 0x7a, 0x07,
	0x6b, 0xb6, 0x34, 0xac, 0xdc, 0x83, 0x2b, 0x8c, 0xa3, 0x25, 0xf9, 0x28, 0x1d, 0x07, 0x59, 0x3f,
	0xf5, 0xdd, 0x30, 0x3d, 0xa7, 0xfc, 0x04, 0xaa, 0xb9, 0xe7, 0x68, 0x31, 0x1e, 0xc1, 0x47, 0x8c,
	0xa3, 0x87, 0xe9, 0x33, 0xb4, 0x08, 0x2e, 0x4f, 0x74, 0x71, 0x69, 0xcf, 0x1d, 0xb5, 0xee, 0x8b,
	0x52, 0x8f, 0x6b, 0xfb, 0xcc, 0x13, 0x5d, 0x39, 0x99, 0x07, 0xf2, 0x44, 0x57, 0x42, 0xea, 0x37,
	0xf0, 0x2a, 0x05, 0xc2, 0x7a, 0xd3, 0xdc, 0x4b, 0x8d, 0x8f, 0x7a, 0x7d, 0xa6, 0xdc, 0x43, 0xb8,
	0xc1, 0x38, 0xaa, 0x46, 0x58, 0x4c, 0x8e, 0xe9, 0x27, 0x31, 0x42, 0x92, 0xc3, 0x62, 0xc9, 0x3f,
	0xc1, 0x0f, 0xd3, 0xed, 0x75, 0x26, 0xdd, 0x45, 0xf8, 0x92, 0x71, 0xb4, 0x21, 0x3a, 0xad, 0xb4,
	0x76, 0x6e, 0xba, 0x89, 0xb7, 0x73, 0x64, 0x97, 0xb2, 0xe9, 0x26, 0xfe, 0xfe, 0xb7, 0x74, 0xcb,
	0xea, 0x2e, 0xe7, 0xa5, 0x5b, 0x4a, 0x7b, 0x5f, 0x1a, 0xb0, 0xc4, 0xb7, 0x82, 0xd1, 0x3d, 0x92,
	0x23, 0xf5, 0x48, 0xb4, 0x41, 0xe2, 0x67, 0x7c, 0xa3, 0x4c, 0xb3, 0x0f, 0xa6, 0xcf, 0xfc, 0x12,
	0x02, 0x8f, 0xe5, 0x43, 0x72, 0x9e, 0x88, 0x9c, 0x80, 0x68, 0x82, 0x42, 0xfe, 0x8a, 0x48, 0x20,
	0x71, 0xa8, 0x7c, 0x02, 0x25, 0x04, 0x56, 0xe5, 0x04, 0xe6, 0x89, 0xf4, 0xc1, 0xcb, 0xd9, 0xfd,
	0x55, 0x28, 0xb4, 0x06, 0x55, 0xc6, 0xd1, 0x5a, 0x6e, 0x73, 0x15, 0x88, 0x85, 0xe0, 0x75, 0x89,
	0xce, 0x2a, 0x54, 0x7c, 0x02, 0x37, 0x19, 0x47, 0x6a, 0x61, 0x5b, 0x9d, 0x25, 0xc5, 0x74, 0x6d,
	0x17, 0x0a, 0xae, 0x67, 0x53, 0x4c, 0x15, 0xf5, 0x99, 0x53, 0x2c, 0xa5, 0x58, 0xcd, 0x4b, 0x71,
	0xbe, 0xaa, 0x3c, 0xda, 0x6a, 0x96, 0x65, 0x6c, 0xb7, 0x46, 0x22, 0xb6, 0x29, 0x5a, 0x77, 0x22,
	0xa2, 0x8a, 0x5b, 0x50, 0x0b, 0x02, 0xaf, 0x37, 0x8c, 0xa8, 0x36, 0x9d, 0xb4, 0xe9, 0x04, 0xfa,
	0xb3, 0xf4, 0xbd, 0x2d, 0x4e, 0x65, 0x1e, 0xf9, 0xa9, 0x7c, 0x4b, 0xcd, 0xa2, 0xcb, 0x57, 0x60,
	0xa7, 0x95, 0x46, 0x6f, 0x61, 0xb3, 0x39, 0x0d, 0x7f, 0x26, 0xe0, 0x9d, 0xa1, 0x9b, 0xa0, 0xb7,
	0x7c, 0x7a, 0x9c, 0x85, 0xe7, 0x4c, 0x48, 0x65, 0x14, 0x36, 0xe4, 0xde, 0x99, 0xa7, 0x92, 0x3b,
	0x4f, 0x8f, 0x77, 0x5b, 0x9e, 0x84, 0x37, 0xc5, 0x3c, 0x3d, 0xda, 0xd2, 0x79, 0xf3, 0x74, 0x2e,
	0xe5, 0xb9, 0x3c, 0x4f, 0x4f, 0x93, 0xe0, 0xf9, 0xbf, 0xfe, 0xbe, 0x7b, 0xee, 0xcd, 0x8b, 0xfd,
	0xcd, 0x9e, 0x17, 0x1e, 0x9d, 0x7e, 0xa8, 0x1e, 0xd0, 0x63, 0x75, 0xd7, 0xf3, 0xdd, 0x1a, 0xf5,
	0x89, 0x1a, 0x90, 0xbe, 0xeb, 0xa9, 0x46, 0xcb, 0xd6, 0xf1, 0x5b, 0x35, 0xfe, 0x55, 0x6b, 0xbd,
	0x47, 0x86, 0xea, 0xe4, 0x37, 0xb1, 0x0f, 0x17, 0xe2, 0xf7, 0x36, 0xff, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x0d, 0x1f, 0xeb, 0x7e, 0x26, 0x13, 0x00, 0x00,
}
