// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/mint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	types "github.com/cometbft/cometbft/proto/tendermint/types"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Minter represents the minting state.
type Minter struct {
	// current annual inflation rate
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation"`
	// current annual expected provisions
<<<<<<< HEAD
	AnnualProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=annual_provisions,json=annualProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"annual_provisions"`
=======
	AnnualProvisions cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=annual_provisions,json=annualProvisions,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"annual_provisions"`
	BlockHeader      types.Header                `protobuf:"bytes,3,opt,name=block_header,json=blockHeader,proto3" json:"block_header"`
>>>>>>> 9676c65... all error resolved
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df116d183c1e223, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

func (m *Minter) GetBlockHeader() types.Header {
	if m != nil {
		return m.BlockHeader
	}
	return types.Header{}
}

// Params defines the parameters for the x/mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// maximum annual change in inflation rate
	InflationRateChange github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=inflation_rate_change,json=inflationRateChange,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_rate_change"`
	// maximum inflation rate
	InflationMax github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=inflation_max,json=inflationMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_max"`
	// minimum inflation rate
	InflationMin github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=inflation_min,json=inflationMin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_min"`
	// goal of percent bonded atoms
	GoalBonded github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=goal_bonded,json=goalBonded,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"goal_bonded"`
	// expected blocks per year
	MintedAmountPerBlock github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=minted_amount_per_block,json=mintedAmountPerBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"minted_amount_per_block" yaml:"minted_amount_per_block"`
	// expected blocks per year
	BlocksPerYear uint64 `protobuf:"varint,7,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty" yaml:"blocks_per_year"`
	// expected blocks per year
	MaxMintableAmount uint64 `protobuf:"varint,8,opt,name=max_mintable_amount,json=maxMintableAmount,proto3" json:"max_mintable_amount,omitempty" yaml:"max_mintable_amount"`
	// Yearly reduction to apply to minted amount
	YearlyReduction github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=yearly_reduction,json=yearlyReduction,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"yearly_reduction" yaml:"yearly_reduction"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df116d183c1e223, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func (m *Params) GetMaxMintableAmount() uint64 {
	if m != nil {
		return m.MaxMintableAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*Minter)(nil), "cosmos.mint.v1beta1.Minter")
	proto.RegisterType((*Params)(nil), "cosmos.mint.v1beta1.Params")
}

func init() { proto.RegisterFile("cosmos/mint/v1beta1/mint.proto", fileDescriptor_2df116d183c1e223) }

var fileDescriptor_2df116d183c1e223 = []byte{
<<<<<<< HEAD
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x13, 0xad, 0x81, 0x8e, 0x2e, 0xba, 0xb3, 0x0a, 0x71, 0xc1, 0x74, 0xdd, 0xc3, 0xb2,
	0x0a, 0x9b, 0x50, 0xbc, 0x89, 0xa7, 0xb6, 0xd7, 0x42, 0xc9, 0xcd, 0x82, 0x84, 0x37, 0xc9, 0x98,
	0x0e, 0x4d, 0x66, 0xc2, 0xcc, 0xb4, 0xb4, 0x1f, 0x41, 0x4f, 0x1e, 0x3d, 0xfa, 0x11, 0x3c, 0xf8,
	0x21, 0x7a, 0xb3, 0x78, 0x12, 0x0f, 0x45, 0xda, 0x83, 0x5f, 0x43, 0x32, 0x13, 0x52, 0xf1, 0xb0,
	0xa7, 0x5c, 0xf2, 0xe7, 0x79, 0xde, 0xfc, 0x9e, 0x27, 0x21, 0x2f, 0xf2, 0x12, 0x2e, 0x0b, 0x2e,
	0x83, 0x82, 0x32, 0x15, 0x2c, 0xfb, 0x31, 0x51, 0xd0, 0xd7, 0x37, 0x7e, 0x29, 0xb8, 0xe2, 0xf8,
	0xcc, 0xf8, 0xbe, 0x96, 0x6a, 0xff, 0xfc, 0x71, 0xc6, 0x33, 0xae, 0xfd, 0xa0, 0xba, 0x32, 0xa3,
	0xe7, 0x4f, 0xcd, 0x68, 0x64, 0x8c, 0xfa, 0x39, 0x63, 0x9d, 0x42, 0x41, 0x19, 0x0f, 0xf4, 0xd1,
	0x48, 0x97, 0xdf, 0x6d, 0xe4, 0x8c, 0x29, 0x53, 0x44, 0xe0, 0x29, 0xea, 0x52, 0xf6, 0x3e, 0x07,
	0x45, 0x39, 0x73, 0xed, 0x0b, 0xfb, 0xba, 0x3b, 0x78, 0xb3, 0xd9, 0xf5, 0xac, 0x5f, 0xbb, 0xde,
	0x55, 0x46, 0xd5, 0x6c, 0x11, 0xfb, 0x09, 0x2f, 0x6a, 0x62, 0x7d, 0xba, 0x91, 0xe9, 0x3c, 0x50,
	0xeb, 0x92, 0x48, 0x7f, 0x44, 0x92, 0x1f, 0xdf, 0x6e, 0x50, 0x1d, 0x38, 0x22, 0x49, 0x78, 0xc4,
	0x61, 0x8a, 0x4e, 0x81, 0xb1, 0x05, 0xe4, 0x55, 0xad, 0x25, 0x95, 0x94, 0x33, 0xe9, 0xde, 0x69,
	0x21, 0xe3, 0x91, 0xc1, 0x4e, 0x1a, 0xea, 0xe5, 0x87, 0x0e, 0x72, 0x26, 0x20, 0xa0, 0x90, 0xf8,
	0x19, 0x42, 0xd5, 0x07, 0x8b, 0x52, 0xc2, 0x78, 0x61, 0x5e, 0x29, 0xec, 0x56, 0xca, 0xa8, 0x12,
	0x70, 0x89, 0x9e, 0x34, 0x0d, 0x23, 0x01, 0x8a, 0x44, 0xc9, 0x0c, 0x58, 0x46, 0x5a, 0x29, 0x76,
	0xd6, 0xa0, 0x43, 0x50, 0x64, 0xa8, 0xc1, 0x18, 0xd0, 0xc9, 0x31, 0xb1, 0x80, 0x95, 0x7b, 0xb7,
	0x85, 0xa4, 0x07, 0x0d, 0x72, 0x0c, 0xab, 0xff, 0x22, 0x28, 0x73, 0x3b, 0xed, 0x46, 0x50, 0x86,
	0xdf, 0xa1, 0xfb, 0x19, 0x87, 0x3c, 0x8a, 0x39, 0x4b, 0x49, 0xea, 0xde, 0x6b, 0x21, 0x00, 0x55,
	0xc0, 0x81, 0xe6, 0xe1, 0x2b, 0xf4, 0x30, 0xce, 0x79, 0x32, 0x97, 0x51, 0x49, 0x44, 0xb4, 0x26,
	0x20, 0x5c, 0xe7, 0xc2, 0xbe, 0xee, 0x84, 0x27, 0x46, 0x9e, 0x10, 0xf1, 0x96, 0x80, 0x78, 0xfd,
	0xfc, 0xf3, 0x97, 0x9e, 0xf5, 0xf1, 0xcf, 0xd7, 0x97, 0xee, 0x3f, 0x01, 0x2b, 0xb3, 0x42, 0xe6,
	0x07, 0x18, 0x0c, 0x37, 0x7b, 0xcf, 0xde, 0xee, 0x3d, 0xfb, 0xf7, 0xde, 0xb3, 0x3f, 0x1d, 0x3c,
	0x6b, 0x7b, 0xf0, 0xac, 0x9f, 0x07, 0xcf, 0x9a, 0xbe, 0xb8, 0xb5, 0x66, 0x4d, 0xd1, 0x6d, 0x63,
	0x47, 0x6f, 0xca, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x18, 0xb9, 0xae, 0xa4, 0x03,
	0x00, 0x00,
=======
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x63, 0x28, 0x81, 0x5c, 0x5b, 0xb5, 0xbd, 0x16, 0x6a, 0x02, 0x75, 0x2a, 0x0f, 0xa8,
	0x54, 0xaa, 0xa3, 0x82, 0xc4, 0xd0, 0xad, 0x6e, 0x06, 0x90, 0x28, 0x44, 0x5e, 0x10, 0x20, 0x61,
	0x9d, 0xed, 0x87, 0x63, 0xea, 0xbb, 0xab, 0xce, 0x97, 0x2a, 0xf9, 0x04, 0x48, 0x4c, 0x7c, 0x0c,
	0xc6, 0x0e, 0x7c, 0x88, 0x8e, 0x15, 0x13, 0x30, 0x44, 0xa8, 0x1d, 0xba, 0x67, 0x66, 0x40, 0xbe,
	0x3b, 0x12, 0x14, 0x04, 0x52, 0x0b, 0x8b, 0x93, 0xfc, 0xff, 0xef, 0x7e, 0xff, 0xf7, 0xce, 0xb9,
	0x43, 0x4e, 0xcc, 0x0b, 0xca, 0x8b, 0x26, 0xcd, 0x98, 0x6c, 0x1e, 0x6c, 0x46, 0x20, 0xc9, 0xa6,
	0xfa, 0xe1, 0xed, 0x0b, 0x2e, 0x39, 0x5e, 0xd4, 0xbe, 0xa7, 0x24, 0xe3, 0xd7, 0x97, 0x52, 0x9e,
	0x72, 0xe5, 0x37, 0xcb, 0x6f, 0xba, 0xb4, 0x7e, 0x53, 0x97, 0x86, 0xda, 0x30, 0xeb, 0xb4, 0xb5,
	0x40, 0x68, 0xc6, 0x78, 0x53, 0x3d, 0x8d, 0x74, 0x5b, 0x02, 0x4b, 0x40, 0xa8, 0x5c, 0xd9, 0xdf,
	0x87, 0x42, 0x3f, 0xb5, 0xeb, 0x7e, 0xb7, 0x50, 0x75, 0x37, 0x63, 0x12, 0x04, 0x7e, 0x8a, 0x6a,
	0x19, 0x7b, 0x9d, 0x13, 0x99, 0x71, 0x66, 0x5b, 0xab, 0xd6, 0x5a, 0xcd, 0xdf, 0x3c, 0x1a, 0x34,
	0x2a, 0x5f, 0x07, 0x8d, 0x5b, 0x3a, 0xa4, 0x48, 0xf6, 0xbc, 0x8c, 0x37, 0x29, 0x91, 0x1d, 0xef,
	0x31, 0xa4, 0x24, 0xee, 0xb7, 0x20, 0xfe, 0xf4, 0x71, 0x03, 0x99, 0x1e, 0x5a, 0x10, 0x07, 0x63,
	0x06, 0x7e, 0x85, 0x16, 0x08, 0x63, 0x5d, 0x92, 0x97, 0x9d, 0x1e, 0x64, 0x45, 0xc6, 0x59, 0x61,
	0x5f, 0xba, 0x28, 0x78, 0x5e, 0xb3, 0xda, 0x23, 0x14, 0xde, 0x46, 0x33, 0x51, 0xce, 0xe3, 0xbd,
	0xb0, 0x03, 0x24, 0x01, 0x61, 0x5f, 0x5e, 0xb5, 0xd6, 0xa6, 0xef, 0xd9, 0xde, 0x78, 0x60, 0x4f,
	0x8f, 0xfa, 0x50, 0xf9, 0xfe, 0x54, 0x19, 0x1a, 0x4c, 0xab, 0x35, 0x5a, 0x72, 0xbf, 0x54, 0x51,
	0xb5, 0x4d, 0x04, 0xa1, 0x05, 0x5e, 0x41, 0xa8, 0x5c, 0x12, 0x26, 0xc0, 0x38, 0xd5, 0xf3, 0x07,
	0xb5, 0x52, 0x69, 0x95, 0x02, 0x7e, 0x83, 0xae, 0x8f, 0x26, 0x0b, 0x05, 0x91, 0x10, 0xc6, 0x1d,
	0xc2, 0x52, 0x30, 0x03, 0x3d, 0x38, 0xf7, 0x40, 0x1f, 0xce, 0x0e, 0xd7, 0xad, 0x60, 0x71, 0x04,
	0x0d, 0x88, 0x84, 0x1d, 0x85, 0xc4, 0x2f, 0xd1, 0xec, 0x38, 0x8b, 0x92, 0x9e, 0x9a, 0xec, 0xe2,
	0x19, 0x33, 0x23, 0xd8, 0x2e, 0xe9, 0x4d, 0xc0, 0x33, 0x66, 0x4f, 0xfd, 0x2f, 0x78, 0xc6, 0xf0,
	0x33, 0x34, 0x9d, 0x72, 0x92, 0x87, 0x11, 0x67, 0x09, 0x24, 0xf6, 0x95, 0x7f, 0x42, 0xa3, 0x12,
	0xe5, 0x2b, 0x12, 0x7e, 0x6b, 0xa1, 0xe5, 0xf2, 0x65, 0x40, 0x12, 0x12, 0xca, 0xbb, 0x4c, 0x86,
	0xfb, 0x20, 0x42, 0xf5, 0x2a, 0xed, 0xaa, 0x4a, 0x69, 0x9b, 0x94, 0x3b, 0x69, 0x26, 0x3b, 0xdd,
	0xc8, 0x8b, 0x39, 0x35, 0x67, 0xc3, 0x7c, 0x6c, 0x14, 0xc9, 0x9e, 0xf9, 0xef, 0xb7, 0x20, 0x1e,
	0x0e, 0x1a, 0x4e, 0x9f, 0xd0, 0x7c, 0xcb, 0xfd, 0x03, 0xd6, 0x0d, 0x96, 0xb4, 0xb3, 0xad, 0x8c,
	0x36, 0x08, 0xbf, 0x94, 0xb1, 0x8f, 0xe6, 0x94, 0x5f, 0xa8, 0xd2, 0x3e, 0x10, 0x61, 0x5f, 0x5d,
	0xb5, 0xd6, 0xa6, 0xfc, 0xfa, 0x70, 0xd0, 0xb8, 0xa1, 0x91, 0x13, 0x05, 0x6e, 0x30, 0xab, 0x95,
	0x36, 0x88, 0xe7, 0x40, 0x04, 0x7e, 0x82, 0x16, 0x29, 0xe9, 0x95, 0xbb, 0x2f, 0x49, 0x94, 0x83,
	0xc9, 0xb6, 0xaf, 0x29, 0x8e, 0x33, 0x1c, 0x34, 0xea, 0xa6, 0xb5, 0xdf, 0x8b, 0xdc, 0x60, 0x81,
	0x92, 0xde, 0xae, 0x11, 0x75, 0x6f, 0x58, 0xa2, 0xf9, 0x32, 0x27, 0xef, 0x87, 0x02, 0x92, 0x6e,
	0xac, 0x4e, 0x70, 0x4d, 0xed, 0xca, 0xa3, 0x73, 0xef, 0xca, 0xb2, 0x8e, 0x9e, 0xe4, 0xb9, 0xc1,
	0x9c, 0x96, 0x82, 0x9f, 0xca, 0xd6, 0xca, 0xbb, 0xb3, 0xc3, 0x75, 0xfb, 0x17, 0x4c, 0x4f, 0xdf,
	0x6e, 0xfa, 0x40, 0xf9, 0x3b, 0x47, 0x27, 0x8e, 0x75, 0x7c, 0xe2, 0x58, 0xdf, 0x4e, 0x1c, 0xeb,
	0xfd, 0xa9, 0x53, 0x39, 0x3e, 0x75, 0x2a, 0x9f, 0x4f, 0x9d, 0xca, 0x8b, 0xbb, 0x7f, 0x6d, 0xc6,
	0x50, 0x54, 0x4f, 0x51, 0x55, 0x5d, 0x53, 0xf7, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x36, 0xe4,
	0xb7, 0x07, 0x3f, 0x05, 0x00, 0x00,
>>>>>>> 9676c65... all error resolved
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BlockHeader.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.YearlyReduction.Size()
		i -= size
		if _, err := m.YearlyReduction.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.MaxMintableAmount != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.MaxMintableAmount))
		i--
		dAtA[i] = 0x40
	}
	if m.BlocksPerYear != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x38
	}
	{
		size := m.MintedAmountPerBlock.Size()
		i -= size
		if _, err := m.MintedAmountPerBlock.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.GoalBonded.Size()
		i -= size
		if _, err := m.GoalBonded.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.InflationMin.Size()
		i -= size
		if _, err := m.InflationMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InflationMax.Size()
		i -= size
		if _, err := m.InflationMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.InflationRateChange.Size()
		i -= size
		if _, err := m.InflationRateChange.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.BlockHeader.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.InflationRateChange.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.InflationMax.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.InflationMin.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.GoalBonded.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.MintedAmountPerBlock.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.BlocksPerYear != 0 {
		n += 1 + sovMint(uint64(m.BlocksPerYear))
	}
	if m.MaxMintableAmount != 0 {
		n += 1 + sovMint(uint64(m.MaxMintableAmount))
	}
	l = m.YearlyReduction.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRateChange", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRateChange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoalBonded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GoalBonded.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintedAmountPerBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MintedAmountPerBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMintableAmount", wireType)
			}
			m.MaxMintableAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMintableAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field YearlyReduction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.YearlyReduction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMint
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMint
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
