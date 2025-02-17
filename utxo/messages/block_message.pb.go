// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.3
// source: utxo/block_message.proto

package utxo_messages

import (
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

type BlockMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *BlockHeader   `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=Transactions,proto3" json:"Transactions,omitempty"`
}

func (x *BlockMessage) Reset() {
	*x = BlockMessage{}
	mi := &file_utxo_block_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlockMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockMessage) ProtoMessage() {}

func (x *BlockMessage) ProtoReflect() protoreflect.Message {
	mi := &file_utxo_block_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockMessage.ProtoReflect.Descriptor instead.
func (*BlockMessage) Descriptor() ([]byte, []int) {
	return file_utxo_block_message_proto_rawDescGZIP(), []int{0}
}

func (x *BlockMessage) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BlockMessage) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type BlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash          []byte  `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Confirmations int64   `protobuf:"varint,2,opt,name=Confirmations,proto3" json:"Confirmations,omitempty"`
	Size          int32   `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	StrippedSize  int32   `protobuf:"varint,4,opt,name=StrippedSize,proto3" json:"StrippedSize,omitempty"`
	Weight        int32   `protobuf:"varint,5,opt,name=Weight,proto3" json:"Weight,omitempty"`
	Height        int64   `protobuf:"varint,6,opt,name=Height,proto3" json:"Height,omitempty"`
	Version       int32   `protobuf:"varint,7,opt,name=Version,proto3" json:"Version,omitempty"`
	VersionHex    []byte  `protobuf:"bytes,8,opt,name=VersionHex,proto3" json:"VersionHex,omitempty"`
	MerkleRoot    []byte  `protobuf:"bytes,9,opt,name=MerkleRoot,proto3" json:"MerkleRoot,omitempty"`
	Time          int64   `protobuf:"varint,10,opt,name=Time,proto3" json:"Time,omitempty"`
	MedianTime    int64   `protobuf:"varint,11,opt,name=MedianTime,proto3" json:"MedianTime,omitempty"`
	Nonce         uint32  `protobuf:"varint,12,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Bits          []byte  `protobuf:"bytes,13,opt,name=Bits,proto3" json:"Bits,omitempty"`
	Difficulty    float64 `protobuf:"fixed64,14,opt,name=Difficulty,proto3" json:"Difficulty,omitempty"`
	Chainwork     []byte  `protobuf:"bytes,15,opt,name=Chainwork,proto3" json:"Chainwork,omitempty"`
	NTx           uint32  `protobuf:"varint,16,opt,name=NTx,proto3" json:"NTx,omitempty"`
	PreviousHash  []byte  `protobuf:"bytes,17,opt,name=PreviousHash,proto3" json:"PreviousHash,omitempty"`
	NextHash      []byte  `protobuf:"bytes,18,opt,name=NextHash,proto3" json:"NextHash,omitempty"`
}

func (x *BlockHeader) Reset() {
	*x = BlockHeader{}
	mi := &file_utxo_block_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeader) ProtoMessage() {}

func (x *BlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_utxo_block_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeader.ProtoReflect.Descriptor instead.
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return file_utxo_block_message_proto_rawDescGZIP(), []int{1}
}

func (x *BlockHeader) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *BlockHeader) GetConfirmations() int64 {
	if x != nil {
		return x.Confirmations
	}
	return 0
}

func (x *BlockHeader) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *BlockHeader) GetStrippedSize() int32 {
	if x != nil {
		return x.StrippedSize
	}
	return 0
}

func (x *BlockHeader) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *BlockHeader) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockHeader) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *BlockHeader) GetVersionHex() []byte {
	if x != nil {
		return x.VersionHex
	}
	return nil
}

func (x *BlockHeader) GetMerkleRoot() []byte {
	if x != nil {
		return x.MerkleRoot
	}
	return nil
}

func (x *BlockHeader) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *BlockHeader) GetMedianTime() int64 {
	if x != nil {
		return x.MedianTime
	}
	return 0
}

func (x *BlockHeader) GetNonce() uint32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *BlockHeader) GetBits() []byte {
	if x != nil {
		return x.Bits
	}
	return nil
}

func (x *BlockHeader) GetDifficulty() float64 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *BlockHeader) GetChainwork() []byte {
	if x != nil {
		return x.Chainwork
	}
	return nil
}

func (x *BlockHeader) GetNTx() uint32 {
	if x != nil {
		return x.NTx
	}
	return 0
}

func (x *BlockHeader) GetPreviousHash() []byte {
	if x != nil {
		return x.PreviousHash
	}
	return nil
}

func (x *BlockHeader) GetNextHash() []byte {
	if x != nil {
		return x.NextHash
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *TransactionHeader   `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Inputs  []*TransactionInput  `protobuf:"bytes,2,rep,name=Inputs,proto3" json:"Inputs,omitempty"`
	Outputs []*TransactionOutput `protobuf:"bytes,3,rep,name=Outputs,proto3" json:"Outputs,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	mi := &file_utxo_block_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_utxo_block_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_utxo_block_message_proto_rawDescGZIP(), []int{2}
}

func (x *Transaction) GetHeader() *TransactionHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Transaction) GetInputs() []*TransactionInput {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *Transaction) GetOutputs() []*TransactionOutput {
	if x != nil {
		return x.Outputs
	}
	return nil
}

type TransactionHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hex      []byte `protobuf:"bytes,1,opt,name=Hex,proto3" json:"Hex,omitempty"`
	Id       []byte `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	Hash     []byte `protobuf:"bytes,3,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Size     int32  `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
	VSize    int32  `protobuf:"varint,5,opt,name=VSize,proto3" json:"VSize,omitempty"`
	Weight   int32  `protobuf:"varint,6,opt,name=Weight,proto3" json:"Weight,omitempty"`
	Version  uint32 `protobuf:"varint,7,opt,name=Version,proto3" json:"Version,omitempty"`
	LockTime uint32 `protobuf:"varint,8,opt,name=LockTime,proto3" json:"LockTime,omitempty"`
}

func (x *TransactionHeader) Reset() {
	*x = TransactionHeader{}
	mi := &file_utxo_block_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionHeader) ProtoMessage() {}

func (x *TransactionHeader) ProtoReflect() protoreflect.Message {
	mi := &file_utxo_block_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionHeader.ProtoReflect.Descriptor instead.
func (*TransactionHeader) Descriptor() ([]byte, []int) {
	return file_utxo_block_message_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionHeader) GetHex() []byte {
	if x != nil {
		return x.Hex
	}
	return nil
}

func (x *TransactionHeader) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *TransactionHeader) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *TransactionHeader) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *TransactionHeader) GetVSize() int32 {
	if x != nil {
		return x.VSize
	}
	return 0
}

func (x *TransactionHeader) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *TransactionHeader) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *TransactionHeader) GetLockTime() uint32 {
	if x != nil {
		return x.LockTime
	}
	return 0
}

type TransactionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coinbase  []byte     `protobuf:"bytes,1,opt,name=Coinbase,proto3" json:"Coinbase,omitempty"`
	TxId      []byte     `protobuf:"bytes,2,opt,name=TxId,proto3" json:"TxId,omitempty"`
	VOut      uint32     `protobuf:"varint,3,opt,name=VOut,proto3" json:"VOut,omitempty"`
	ScriptSig *ScriptSig `protobuf:"bytes,4,opt,name=ScriptSig,proto3" json:"ScriptSig,omitempty"`
	Sequence  uint32     `protobuf:"varint,5,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	Witness   []string   `protobuf:"bytes,6,rep,name=Witness,proto3" json:"Witness,omitempty"`
}

func (x *TransactionInput) Reset() {
	*x = TransactionInput{}
	mi := &file_utxo_block_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionInput) ProtoMessage() {}

func (x *TransactionInput) ProtoReflect() protoreflect.Message {
	mi := &file_utxo_block_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionInput.ProtoReflect.Descriptor instead.
func (*TransactionInput) Descriptor() ([]byte, []int) {
	return file_utxo_block_message_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionInput) GetCoinbase() []byte {
	if x != nil {
		return x.Coinbase
	}
	return nil
}

func (x *TransactionInput) GetTxId() []byte {
	if x != nil {
		return x.TxId
	}
	return nil
}

func (x *TransactionInput) GetVOut() uint32 {
	if x != nil {
		return x.VOut
	}
	return 0
}

func (x *TransactionInput) GetScriptSig() *ScriptSig {
	if x != nil {
		return x.ScriptSig
	}
	return nil
}

func (x *TransactionInput) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *TransactionInput) GetWitness() []string {
	if x != nil {
		return x.Witness
	}
	return nil
}

type ScriptSig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Asm string `protobuf:"bytes,1,opt,name=Asm,proto3" json:"Asm,omitempty"`
	Hex []byte `protobuf:"bytes,2,opt,name=Hex,proto3" json:"Hex,omitempty"`
}

func (x *ScriptSig) Reset() {
	*x = ScriptSig{}
	mi := &file_utxo_block_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScriptSig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScriptSig) ProtoMessage() {}

func (x *ScriptSig) ProtoReflect() protoreflect.Message {
	mi := &file_utxo_block_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScriptSig.ProtoReflect.Descriptor instead.
func (*ScriptSig) Descriptor() ([]byte, []int) {
	return file_utxo_block_message_proto_rawDescGZIP(), []int{5}
}

func (x *ScriptSig) GetAsm() string {
	if x != nil {
		return x.Asm
	}
	return ""
}

func (x *ScriptSig) GetHex() []byte {
	if x != nil {
		return x.Hex
	}
	return nil
}

type TransactionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value        float64       `protobuf:"fixed64,1,opt,name=Value,proto3" json:"Value,omitempty"`
	N            uint32        `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
	ScriptPubKey *ScriptPubKey `protobuf:"bytes,3,opt,name=ScriptPubKey,proto3" json:"ScriptPubKey,omitempty"`
}

func (x *TransactionOutput) Reset() {
	*x = TransactionOutput{}
	mi := &file_utxo_block_message_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionOutput) ProtoMessage() {}

func (x *TransactionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_utxo_block_message_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionOutput.ProtoReflect.Descriptor instead.
func (*TransactionOutput) Descriptor() ([]byte, []int) {
	return file_utxo_block_message_proto_rawDescGZIP(), []int{6}
}

func (x *TransactionOutput) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TransactionOutput) GetN() uint32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *TransactionOutput) GetScriptPubKey() *ScriptPubKey {
	if x != nil {
		return x.ScriptPubKey
	}
	return nil
}

type ScriptPubKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Asm     string `protobuf:"bytes,1,opt,name=Asm,proto3" json:"Asm,omitempty"`
	Desc    string `protobuf:"bytes,2,opt,name=Desc,proto3" json:"Desc,omitempty"`
	Hex     []byte `protobuf:"bytes,3,opt,name=Hex,proto3" json:"Hex,omitempty"`
	Type    string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Address string `protobuf:"bytes,5,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *ScriptPubKey) Reset() {
	*x = ScriptPubKey{}
	mi := &file_utxo_block_message_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScriptPubKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScriptPubKey) ProtoMessage() {}

func (x *ScriptPubKey) ProtoReflect() protoreflect.Message {
	mi := &file_utxo_block_message_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScriptPubKey.ProtoReflect.Descriptor instead.
func (*ScriptPubKey) Descriptor() ([]byte, []int) {
	return file_utxo_block_message_proto_rawDescGZIP(), []int{7}
}

func (x *ScriptPubKey) GetAsm() string {
	if x != nil {
		return x.Asm
	}
	return ""
}

func (x *ScriptPubKey) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ScriptPubKey) GetHex() []byte {
	if x != nil {
		return x.Hex
	}
	return nil
}

func (x *ScriptPubKey) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ScriptPubKey) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_utxo_block_message_proto protoreflect.FileDescriptor

var file_utxo_block_message_proto_rawDesc = []byte{
	0x0a, 0x18, 0x75, 0x74, 0x78, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x74, 0x78, 0x6f,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x75, 0x74, 0x78,
	0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3e,
	0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xf7,
	0x03, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x53, 0x74, 0x72, 0x69, 0x70, 0x70, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x53, 0x74, 0x72, 0x69, 0x70, 0x70, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x65,
	0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4e,
	0x6f, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x69, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x42, 0x69, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x44, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x54, 0x78, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x4e, 0x54, 0x78, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x48, 0x61, 0x73, 0x68, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08,
	0x4e, 0x65, 0x78, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x4e, 0x65, 0x78, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22, 0xbc, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x37, 0x0a, 0x06, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x52, 0x06, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x75,
	0x74, 0x78, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22, 0xc1, 0x01, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x48, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x48, 0x65, 0x78, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xc4, 0x01, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x78, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x54, 0x78, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x56, 0x4f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x56, 0x4f, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69,
	0x67, 0x52, 0x09, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x69, 0x74, 0x6e,
	0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x57, 0x69, 0x74, 0x6e, 0x65,
	0x73, 0x73, 0x22, 0x2f, 0x0a, 0x09, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x41, 0x73, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x73,
	0x6d, 0x12, 0x10, 0x0a, 0x03, 0x48, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x48, 0x65, 0x78, 0x22, 0x78, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0c,
	0x0a, 0x01, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x4e, 0x12, 0x3f, 0x0a, 0x0c,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52,
	0x0c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x22, 0x74, 0x0a,
	0x0c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x41, 0x73, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x73, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44,
	0x65, 0x73, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x48, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x48, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_utxo_block_message_proto_rawDescOnce sync.Once
	file_utxo_block_message_proto_rawDescData = file_utxo_block_message_proto_rawDesc
)

func file_utxo_block_message_proto_rawDescGZIP() []byte {
	file_utxo_block_message_proto_rawDescOnce.Do(func() {
		file_utxo_block_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_utxo_block_message_proto_rawDescData)
	})
	return file_utxo_block_message_proto_rawDescData
}

var file_utxo_block_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_utxo_block_message_proto_goTypes = []any{
	(*BlockMessage)(nil),      // 0: utxo_messages.BlockMessage
	(*BlockHeader)(nil),       // 1: utxo_messages.BlockHeader
	(*Transaction)(nil),       // 2: utxo_messages.Transaction
	(*TransactionHeader)(nil), // 3: utxo_messages.TransactionHeader
	(*TransactionInput)(nil),  // 4: utxo_messages.TransactionInput
	(*ScriptSig)(nil),         // 5: utxo_messages.ScriptSig
	(*TransactionOutput)(nil), // 6: utxo_messages.TransactionOutput
	(*ScriptPubKey)(nil),      // 7: utxo_messages.ScriptPubKey
}
var file_utxo_block_message_proto_depIdxs = []int32{
	1, // 0: utxo_messages.BlockMessage.Header:type_name -> utxo_messages.BlockHeader
	2, // 1: utxo_messages.BlockMessage.Transactions:type_name -> utxo_messages.Transaction
	3, // 2: utxo_messages.Transaction.Header:type_name -> utxo_messages.TransactionHeader
	4, // 3: utxo_messages.Transaction.Inputs:type_name -> utxo_messages.TransactionInput
	6, // 4: utxo_messages.Transaction.Outputs:type_name -> utxo_messages.TransactionOutput
	5, // 5: utxo_messages.TransactionInput.ScriptSig:type_name -> utxo_messages.ScriptSig
	7, // 6: utxo_messages.TransactionOutput.ScriptPubKey:type_name -> utxo_messages.ScriptPubKey
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_utxo_block_message_proto_init() }
func file_utxo_block_message_proto_init() {
	if File_utxo_block_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_utxo_block_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_utxo_block_message_proto_goTypes,
		DependencyIndexes: file_utxo_block_message_proto_depIdxs,
		MessageInfos:      file_utxo_block_message_proto_msgTypes,
	}.Build()
	File_utxo_block_message_proto = out.File
	file_utxo_block_message_proto_rawDesc = nil
	file_utxo_block_message_proto_goTypes = nil
	file_utxo_block_message_proto_depIdxs = nil
}
