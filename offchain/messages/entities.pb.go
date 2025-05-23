// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.3
// source: offchain/entities.proto

package entities_messages

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

type ProviderDataMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider *Provider `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Data     []*Data   `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ProviderDataMessage) Reset() {
	*x = ProviderDataMessage{}
	mi := &file_offchain_entities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProviderDataMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderDataMessage) ProtoMessage() {}

func (x *ProviderDataMessage) ProtoReflect() protoreflect.Message {
	mi := &file_offchain_entities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderDataMessage.ProtoReflect.Descriptor instead.
func (*ProviderDataMessage) Descriptor() ([]byte, []int) {
	return file_offchain_entities_proto_rawDescGZIP(), []int{0}
}

func (x *ProviderDataMessage) GetProvider() *Provider {
	if x != nil {
		return x.Provider
	}
	return nil
}

func (x *ProviderDataMessage) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  *AddressWithBlockchain `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Response []byte                 `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_offchain_entities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_offchain_entities_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_offchain_entities_proto_rawDescGZIP(), []int{1}
}

func (x *Data) GetAddress() *AddressWithBlockchain {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Data) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

type AddressWithBlockchain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blockchain string `protobuf:"bytes,1,opt,name=blockchain,proto3" json:"blockchain,omitempty"`
	Address    string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AddressWithBlockchain) Reset() {
	*x = AddressWithBlockchain{}
	mi := &file_offchain_entities_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressWithBlockchain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressWithBlockchain) ProtoMessage() {}

func (x *AddressWithBlockchain) ProtoReflect() protoreflect.Message {
	mi := &file_offchain_entities_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressWithBlockchain.ProtoReflect.Descriptor instead.
func (*AddressWithBlockchain) Descriptor() ([]byte, []int) {
	return file_offchain_entities_proto_rawDescGZIP(), []int{2}
}

func (x *AddressWithBlockchain) GetBlockchain() string {
	if x != nil {
		return x.Blockchain
	}
	return ""
}

func (x *AddressWithBlockchain) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag  string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Provider) Reset() {
	*x = Provider{}
	mi := &file_offchain_entities_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provider) ProtoMessage() {}

func (x *Provider) ProtoReflect() protoreflect.Message {
	mi := &file_offchain_entities_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provider.ProtoReflect.Descriptor instead.
func (*Provider) Descriptor() ([]byte, []int) {
	return file_offchain_entities_proto_rawDescGZIP(), []int{3}
}

func (x *Provider) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Provider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Provider) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_offchain_entities_proto protoreflect.FileDescriptor

var file_offchain_entities_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x7b, 0x0a, 0x13,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x66, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x42, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x57, 0x69,
	0x74, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x51, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x57, 0x69, 0x74, 0x68,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x42, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_offchain_entities_proto_rawDescOnce sync.Once
	file_offchain_entities_proto_rawDescData = file_offchain_entities_proto_rawDesc
)

func file_offchain_entities_proto_rawDescGZIP() []byte {
	file_offchain_entities_proto_rawDescOnce.Do(func() {
		file_offchain_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_offchain_entities_proto_rawDescData)
	})
	return file_offchain_entities_proto_rawDescData
}

var file_offchain_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_offchain_entities_proto_goTypes = []any{
	(*ProviderDataMessage)(nil),   // 0: entities_messages.ProviderDataMessage
	(*Data)(nil),                  // 1: entities_messages.Data
	(*AddressWithBlockchain)(nil), // 2: entities_messages.AddressWithBlockchain
	(*Provider)(nil),              // 3: entities_messages.Provider
}
var file_offchain_entities_proto_depIdxs = []int32{
	3, // 0: entities_messages.ProviderDataMessage.provider:type_name -> entities_messages.Provider
	1, // 1: entities_messages.ProviderDataMessage.data:type_name -> entities_messages.Data
	2, // 2: entities_messages.Data.address:type_name -> entities_messages.AddressWithBlockchain
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_offchain_entities_proto_init() }
func file_offchain_entities_proto_init() {
	if File_offchain_entities_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_offchain_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_offchain_entities_proto_goTypes,
		DependencyIndexes: file_offchain_entities_proto_depIdxs,
		MessageInfos:      file_offchain_entities_proto_msgTypes,
	}.Build()
	File_offchain_entities_proto = out.File
	file_offchain_entities_proto_rawDesc = nil
	file_offchain_entities_proto_goTypes = nil
	file_offchain_entities_proto_depIdxs = nil
}
