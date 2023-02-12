// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: stonk/v1/stonk_service.proto

package stonkv1

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

type GetStonkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stonk string `protobuf:"bytes,1,opt,name=stonk,proto3" json:"stonk,omitempty"`
}

func (x *GetStonkRequest) Reset() {
	*x = GetStonkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stonk_v1_stonk_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStonkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStonkRequest) ProtoMessage() {}

func (x *GetStonkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stonk_v1_stonk_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStonkRequest.ProtoReflect.Descriptor instead.
func (*GetStonkRequest) Descriptor() ([]byte, []int) {
	return file_stonk_v1_stonk_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetStonkRequest) GetStonk() string {
	if x != nil {
		return x.Stonk
	}
	return ""
}

type GetStonkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stonk    string `protobuf:"bytes,1,opt,name=stonk,proto3" json:"stonk,omitempty"`
	Buyornot string `protobuf:"bytes,2,opt,name=buyornot,proto3" json:"buyornot,omitempty"`
}

func (x *GetStonkResponse) Reset() {
	*x = GetStonkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stonk_v1_stonk_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStonkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStonkResponse) ProtoMessage() {}

func (x *GetStonkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stonk_v1_stonk_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStonkResponse.ProtoReflect.Descriptor instead.
func (*GetStonkResponse) Descriptor() ([]byte, []int) {
	return file_stonk_v1_stonk_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetStonkResponse) GetStonk() string {
	if x != nil {
		return x.Stonk
	}
	return ""
}

func (x *GetStonkResponse) GetBuyornot() string {
	if x != nil {
		return x.Buyornot
	}
	return ""
}

type GainersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GainersRequest) Reset() {
	*x = GainersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stonk_v1_stonk_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GainersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GainersRequest) ProtoMessage() {}

func (x *GainersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stonk_v1_stonk_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GainersRequest.ProtoReflect.Descriptor instead.
func (*GainersRequest) Descriptor() ([]byte, []int) {
	return file_stonk_v1_stonk_service_proto_rawDescGZIP(), []int{2}
}

type GainersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol        string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Change        string `protobuf:"bytes,4,opt,name=change,proto3" json:"change,omitempty"`
	PercentChange string `protobuf:"bytes,5,opt,name=percent_change,json=percentChange,proto3" json:"percent_change,omitempty"`
	Volume        string `protobuf:"bytes,6,opt,name=volume,proto3" json:"volume,omitempty"`
	AvgVolume     string `protobuf:"bytes,7,opt,name=avg_volume,json=avgVolume,proto3" json:"avg_volume,omitempty"`
	MarketCap     string `protobuf:"bytes,8,opt,name=market_cap,json=marketCap,proto3" json:"market_cap,omitempty"`
	PeRatio       string `protobuf:"bytes,9,opt,name=pe_ratio,json=peRatio,proto3" json:"pe_ratio,omitempty"`
	Week52High    string `protobuf:"bytes,10,opt,name=week52_high,json=week52High,proto3" json:"week52_high,omitempty"`
	Unused        string `protobuf:"bytes,11,opt,name=unused,proto3" json:"unused,omitempty"`
}

func (x *GainersResponse) Reset() {
	*x = GainersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stonk_v1_stonk_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GainersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GainersResponse) ProtoMessage() {}

func (x *GainersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stonk_v1_stonk_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GainersResponse.ProtoReflect.Descriptor instead.
func (*GainersResponse) Descriptor() ([]byte, []int) {
	return file_stonk_v1_stonk_service_proto_rawDescGZIP(), []int{3}
}

func (x *GainersResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *GainersResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GainersResponse) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *GainersResponse) GetChange() string {
	if x != nil {
		return x.Change
	}
	return ""
}

func (x *GainersResponse) GetPercentChange() string {
	if x != nil {
		return x.PercentChange
	}
	return ""
}

func (x *GainersResponse) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

func (x *GainersResponse) GetAvgVolume() string {
	if x != nil {
		return x.AvgVolume
	}
	return ""
}

func (x *GainersResponse) GetMarketCap() string {
	if x != nil {
		return x.MarketCap
	}
	return ""
}

func (x *GainersResponse) GetPeRatio() string {
	if x != nil {
		return x.PeRatio
	}
	return ""
}

func (x *GainersResponse) GetWeek52High() string {
	if x != nil {
		return x.Week52High
	}
	return ""
}

func (x *GainersResponse) GetUnused() string {
	if x != nil {
		return x.Unused
	}
	return ""
}

type LoosersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoosersRequest) Reset() {
	*x = LoosersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stonk_v1_stonk_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoosersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoosersRequest) ProtoMessage() {}

func (x *LoosersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stonk_v1_stonk_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoosersRequest.ProtoReflect.Descriptor instead.
func (*LoosersRequest) Descriptor() ([]byte, []int) {
	return file_stonk_v1_stonk_service_proto_rawDescGZIP(), []int{4}
}

type LoosersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol        string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Change        string `protobuf:"bytes,4,opt,name=change,proto3" json:"change,omitempty"`
	PercentChange string `protobuf:"bytes,5,opt,name=percent_change,json=percentChange,proto3" json:"percent_change,omitempty"`
	Volume        string `protobuf:"bytes,6,opt,name=volume,proto3" json:"volume,omitempty"`
	AvgVolume     string `protobuf:"bytes,7,opt,name=avg_volume,json=avgVolume,proto3" json:"avg_volume,omitempty"`
	MarketCap     string `protobuf:"bytes,8,opt,name=market_cap,json=marketCap,proto3" json:"market_cap,omitempty"`
	PeRatio       string `protobuf:"bytes,9,opt,name=pe_ratio,json=peRatio,proto3" json:"pe_ratio,omitempty"`
	Week52High    string `protobuf:"bytes,10,opt,name=week52_high,json=week52High,proto3" json:"week52_high,omitempty"`
	Unused        string `protobuf:"bytes,11,opt,name=unused,proto3" json:"unused,omitempty"`
}

func (x *LoosersResponse) Reset() {
	*x = LoosersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stonk_v1_stonk_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoosersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoosersResponse) ProtoMessage() {}

func (x *LoosersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stonk_v1_stonk_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoosersResponse.ProtoReflect.Descriptor instead.
func (*LoosersResponse) Descriptor() ([]byte, []int) {
	return file_stonk_v1_stonk_service_proto_rawDescGZIP(), []int{5}
}

func (x *LoosersResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *LoosersResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoosersResponse) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *LoosersResponse) GetChange() string {
	if x != nil {
		return x.Change
	}
	return ""
}

func (x *LoosersResponse) GetPercentChange() string {
	if x != nil {
		return x.PercentChange
	}
	return ""
}

func (x *LoosersResponse) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

func (x *LoosersResponse) GetAvgVolume() string {
	if x != nil {
		return x.AvgVolume
	}
	return ""
}

func (x *LoosersResponse) GetMarketCap() string {
	if x != nil {
		return x.MarketCap
	}
	return ""
}

func (x *LoosersResponse) GetPeRatio() string {
	if x != nil {
		return x.PeRatio
	}
	return ""
}

func (x *LoosersResponse) GetWeek52High() string {
	if x != nil {
		return x.Week52High
	}
	return ""
}

func (x *LoosersResponse) GetUnused() string {
	if x != nil {
		return x.Unused
	}
	return ""
}

type CompareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompareRequest) Reset() {
	*x = CompareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stonk_v1_stonk_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareRequest) ProtoMessage() {}

func (x *CompareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stonk_v1_stonk_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareRequest.ProtoReflect.Descriptor instead.
func (*CompareRequest) Descriptor() ([]byte, []int) {
	return file_stonk_v1_stonk_service_proto_rawDescGZIP(), []int{6}
}

type CompareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strat     string `protobuf:"bytes,1,opt,name=strat,proto3" json:"strat,omitempty"`
	NoStrat   string `protobuf:"bytes,2,opt,name=no_strat,json=noStrat,proto3" json:"no_strat,omitempty"`
	GraphLink string `protobuf:"bytes,3,opt,name=graph_link,json=graphLink,proto3" json:"graph_link,omitempty"`
}

func (x *CompareResponse) Reset() {
	*x = CompareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stonk_v1_stonk_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareResponse) ProtoMessage() {}

func (x *CompareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stonk_v1_stonk_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareResponse.ProtoReflect.Descriptor instead.
func (*CompareResponse) Descriptor() ([]byte, []int) {
	return file_stonk_v1_stonk_service_proto_rawDescGZIP(), []int{7}
}

func (x *CompareResponse) GetStrat() string {
	if x != nil {
		return x.Strat
	}
	return ""
}

func (x *CompareResponse) GetNoStrat() string {
	if x != nil {
		return x.NoStrat
	}
	return ""
}

func (x *CompareResponse) GetGraphLink() string {
	if x != nil {
		return x.GraphLink
	}
	return ""
}

var File_stonk_v1_stonk_service_proto protoreflect.FileDescriptor

var file_stonk_v1_stonk_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x74, 0x6f, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x6e, 0x6b,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x74, 0x6f, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x6f, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x6f, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x6e,
	0x6b, 0x22, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x62,
	0x75, 0x79, 0x6f, 0x72, 0x6e, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62,
	0x75, 0x79, 0x6f, 0x72, 0x6e, 0x6f, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xbc, 0x02, 0x0a, 0x0f, 0x47, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x67, 0x5f, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x67, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f,
	0x63, 0x61, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x43, 0x61, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12,
	0x1f, 0x0a, 0x0b, 0x77, 0x65, 0x65, 0x6b, 0x35, 0x32, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x65, 0x6b, 0x35, 0x32, 0x48, 0x69, 0x67, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x6f, 0x6f, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xbc, 0x02, 0x0a, 0x0f, 0x4c,
	0x6f, 0x6f, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x67, 0x5f, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x67,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x5f, 0x63, 0x61, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x43, 0x61, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f,
	0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x65, 0x65, 0x6b, 0x35, 0x32, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x65, 0x6b, 0x35, 0x32, 0x48, 0x69, 0x67,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x61, 0x0a, 0x0f, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x72, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x53, 0x74, 0x72, 0x61, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x61, 0x70, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x32, 0x99,
	0x02, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x43, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x6e, 0x6b, 0x12, 0x19, 0x2e, 0x73, 0x74,
	0x6f, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x47, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12,
	0x18, 0x2e, 0x73, 0x74, 0x6f, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x6e,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x4c, 0x6f, 0x6f, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x6f,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74,
	0x6f, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x6f, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x74, 0x6f, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x65, 0x74, 0x7a, 0x79, 0x31,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x62, 0x6f, 0x74, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x74, 0x6f, 0x6e, 0x6b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stonk_v1_stonk_service_proto_rawDescOnce sync.Once
	file_stonk_v1_stonk_service_proto_rawDescData = file_stonk_v1_stonk_service_proto_rawDesc
)

func file_stonk_v1_stonk_service_proto_rawDescGZIP() []byte {
	file_stonk_v1_stonk_service_proto_rawDescOnce.Do(func() {
		file_stonk_v1_stonk_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stonk_v1_stonk_service_proto_rawDescData)
	})
	return file_stonk_v1_stonk_service_proto_rawDescData
}

var file_stonk_v1_stonk_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_stonk_v1_stonk_service_proto_goTypes = []interface{}{
	(*GetStonkRequest)(nil),  // 0: stonk.v1.GetStonkRequest
	(*GetStonkResponse)(nil), // 1: stonk.v1.GetStonkResponse
	(*GainersRequest)(nil),   // 2: stonk.v1.GainersRequest
	(*GainersResponse)(nil),  // 3: stonk.v1.GainersResponse
	(*LoosersRequest)(nil),   // 4: stonk.v1.LoosersRequest
	(*LoosersResponse)(nil),  // 5: stonk.v1.LoosersResponse
	(*CompareRequest)(nil),   // 6: stonk.v1.CompareRequest
	(*CompareResponse)(nil),  // 7: stonk.v1.CompareResponse
}
var file_stonk_v1_stonk_service_proto_depIdxs = []int32{
	0, // 0: stonk.v1.StonkService.GetStonk:input_type -> stonk.v1.GetStonkRequest
	2, // 1: stonk.v1.StonkService.Gainers:input_type -> stonk.v1.GainersRequest
	4, // 2: stonk.v1.StonkService.Loosers:input_type -> stonk.v1.LoosersRequest
	6, // 3: stonk.v1.StonkService.Compare:input_type -> stonk.v1.CompareRequest
	1, // 4: stonk.v1.StonkService.GetStonk:output_type -> stonk.v1.GetStonkResponse
	3, // 5: stonk.v1.StonkService.Gainers:output_type -> stonk.v1.GainersResponse
	5, // 6: stonk.v1.StonkService.Loosers:output_type -> stonk.v1.LoosersResponse
	7, // 7: stonk.v1.StonkService.Compare:output_type -> stonk.v1.CompareResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stonk_v1_stonk_service_proto_init() }
func file_stonk_v1_stonk_service_proto_init() {
	if File_stonk_v1_stonk_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stonk_v1_stonk_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStonkRequest); i {
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
		file_stonk_v1_stonk_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStonkResponse); i {
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
		file_stonk_v1_stonk_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GainersRequest); i {
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
		file_stonk_v1_stonk_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GainersResponse); i {
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
		file_stonk_v1_stonk_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoosersRequest); i {
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
		file_stonk_v1_stonk_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoosersResponse); i {
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
		file_stonk_v1_stonk_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareRequest); i {
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
		file_stonk_v1_stonk_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareResponse); i {
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
			RawDescriptor: file_stonk_v1_stonk_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stonk_v1_stonk_service_proto_goTypes,
		DependencyIndexes: file_stonk_v1_stonk_service_proto_depIdxs,
		MessageInfos:      file_stonk_v1_stonk_service_proto_msgTypes,
	}.Build()
	File_stonk_v1_stonk_service_proto = out.File
	file_stonk_v1_stonk_service_proto_rawDesc = nil
	file_stonk_v1_stonk_service_proto_goTypes = nil
	file_stonk_v1_stonk_service_proto_depIdxs = nil
}
