// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: flight_search.proto

package messages

import (
	reflect "reflect"
	sync "sync"

	types "github.com/chain4travel/camino-messenger-bot/proto/pb/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// # Flight Search Request Message
//
// ## Description
//
// Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, diam
// vitae aliquam tincidunt, nunc nisl ultricies nunc, quis aliquam nisl nisl
// velit. Nulla facilisi. Nulla facilisi. Nulla facilisi. Nulla facilisi.
//
// ## Diagram
//
// ![FlightSearch](http://34.36.20.191/tcm/static/diagrams/messages/flight_search.proto.dot.svg)
type FlightSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message Header
	// Contains api version, message info string, sender & receiver wallet
	// addresses
	Header                *types.Header                 `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ExternalSessionId     string                        `protobuf:"bytes,2,opt,name=external_session_id,json=externalSessionId,proto3" json:"external_session_id,omitempty"`
	Enrichment            string                        `protobuf:"bytes,3,opt,name=enrichment,proto3" json:"enrichment,omitempty"`
	Currency              string                        `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	Language              string                        `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`
	Market                string                        `protobuf:"bytes,6,opt,name=market,proto3" json:"market,omitempty"`
	IncludeOnRequest      string                        `protobuf:"bytes,7,opt,name=include_on_request,json=includeOnRequest,proto3" json:"include_on_request,omitempty"`
	IncludeCombinations   string                        `protobuf:"bytes,8,opt,name=include_combinations,json=includeCombinations,proto3" json:"include_combinations,omitempty"`
	Passengers            []*types.Traveller            `protobuf:"bytes,9,rep,name=passengers,proto3" json:"passengers,omitempty"`
	FlightSearchCriterias []*types.FlightSearchCriteria `protobuf:"bytes,10,rep,name=flight_search_criterias,json=flightSearchCriterias,proto3" json:"flight_search_criterias,omitempty"`
	FlightOfferCriterias  []*types.FlightOfferCriteria  `protobuf:"bytes,11,rep,name=flight_offer_criterias,json=flightOfferCriterias,proto3" json:"flight_offer_criterias,omitempty"`
}

func (x *FlightSearchRequest) Reset() {
	*x = FlightSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightSearchRequest) ProtoMessage() {}

func (x *FlightSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightSearchRequest.ProtoReflect.Descriptor instead.
func (*FlightSearchRequest) Descriptor() ([]byte, []int) {
	return file_flight_search_proto_rawDescGZIP(), []int{0}
}

func (x *FlightSearchRequest) GetHeader() *types.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FlightSearchRequest) GetExternalSessionId() string {
	if x != nil {
		return x.ExternalSessionId
	}
	return ""
}

func (x *FlightSearchRequest) GetEnrichment() string {
	if x != nil {
		return x.Enrichment
	}
	return ""
}

func (x *FlightSearchRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *FlightSearchRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *FlightSearchRequest) GetMarket() string {
	if x != nil {
		return x.Market
	}
	return ""
}

func (x *FlightSearchRequest) GetIncludeOnRequest() string {
	if x != nil {
		return x.IncludeOnRequest
	}
	return ""
}

func (x *FlightSearchRequest) GetIncludeCombinations() string {
	if x != nil {
		return x.IncludeCombinations
	}
	return ""
}

func (x *FlightSearchRequest) GetPassengers() []*types.Traveller {
	if x != nil {
		return x.Passengers
	}
	return nil
}

func (x *FlightSearchRequest) GetFlightSearchCriterias() []*types.FlightSearchCriteria {
	if x != nil {
		return x.FlightSearchCriterias
	}
	return nil
}

func (x *FlightSearchRequest) GetFlightOfferCriterias() []*types.FlightOfferCriteria {
	if x != nil {
		return x.FlightOfferCriterias
	}
	return nil
}

// # Flight Search Response Message
//
// ## Description
//
// Vitae aliquam tincidunt, nunc nisl ultricies nunc, quis aliquam nisl nisl
// velit. Nulla facilisi. Nulla facilisi. Nulla facilisi. Nulla facilisi.
// Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, diam
//
// ## Diagram
//
// ![FlightSearch](http://34.36.20.191/tcm/static/diagrams/messages/flight_search.proto.dot.svg)
type FlightSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message Header
	// Contains api version, message info string, sender & receiver wallet
	// addresses
	Header            *types.Header               `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Context           string                      `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	Errors            string                      `protobuf:"bytes,3,opt,name=errors,proto3" json:"errors,omitempty"`
	Warnings          string                      `protobuf:"bytes,4,opt,name=warnings,proto3" json:"warnings,omitempty"`
	SupplierCode      string                      `protobuf:"bytes,5,opt,name=supplier_code,json=supplierCode,proto3" json:"supplier_code,omitempty"`
	ExternalSessionId string                      `protobuf:"bytes,6,opt,name=external_session_id,json=externalSessionId,proto3" json:"external_session_id,omitempty"`
	SearchId          string                      `protobuf:"bytes,7,opt,name=search_id,json=searchId,proto3" json:"search_id,omitempty"`
	Options           []*types.FlightSearchOption `protobuf:"bytes,8,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *FlightSearchResponse) Reset() {
	*x = FlightSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightSearchResponse) ProtoMessage() {}

func (x *FlightSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightSearchResponse.ProtoReflect.Descriptor instead.
func (*FlightSearchResponse) Descriptor() ([]byte, []int) {
	return file_flight_search_proto_rawDescGZIP(), []int{1}
}

func (x *FlightSearchResponse) GetHeader() *types.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FlightSearchResponse) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *FlightSearchResponse) GetErrors() string {
	if x != nil {
		return x.Errors
	}
	return ""
}

func (x *FlightSearchResponse) GetWarnings() string {
	if x != nil {
		return x.Warnings
	}
	return ""
}

func (x *FlightSearchResponse) GetSupplierCode() string {
	if x != nil {
		return x.SupplierCode
	}
	return ""
}

func (x *FlightSearchResponse) GetExternalSessionId() string {
	if x != nil {
		return x.ExternalSessionId
	}
	return ""
}

func (x *FlightSearchResponse) GetSearchId() string {
	if x != nil {
		return x.SearchId
	}
	return ""
}

func (x *FlightSearchResponse) GetOptions() []*types.FlightSearchOption {
	if x != nil {
		return x.Options
	}
	return nil
}

var File_flight_search_proto protoreflect.FileDescriptor

var file_flight_search_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a,
	0x17, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x63, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x04, 0x0a, 0x13, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x4f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x14, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x63, 0x6f,
	0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x0a, 0x70, 0x61, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x12, 0x53, 0x0a, 0x17, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x15, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x73, 0x12, 0x50, 0x0a, 0x16,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x14, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x73, 0x22, 0xb2,
	0x02, 0x0a, 0x14, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x64, 0x12, 0x33,
	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x32, 0x60, 0x0a, 0x13, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x34, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c,
	0x2f, 0x63, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x2d, 0x62, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_flight_search_proto_rawDescOnce sync.Once
	file_flight_search_proto_rawDescData = file_flight_search_proto_rawDesc
)

func file_flight_search_proto_rawDescGZIP() []byte {
	file_flight_search_proto_rawDescOnce.Do(func() {
		file_flight_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_flight_search_proto_rawDescData)
	})
	return file_flight_search_proto_rawDescData
}

var file_flight_search_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flight_search_proto_goTypes = []interface{}{
	(*FlightSearchRequest)(nil),        // 0: messages.FlightSearchRequest
	(*FlightSearchResponse)(nil),       // 1: messages.FlightSearchResponse
	(*types.Header)(nil),               // 2: types.Header
	(*types.Traveller)(nil),            // 3: types.Traveller
	(*types.FlightSearchCriteria)(nil), // 4: types.FlightSearchCriteria
	(*types.FlightOfferCriteria)(nil),  // 5: types.FlightOfferCriteria
	(*types.FlightSearchOption)(nil),   // 6: types.FlightSearchOption
}
var file_flight_search_proto_depIdxs = []int32{
	2, // 0: messages.FlightSearchRequest.header:type_name -> types.Header
	3, // 1: messages.FlightSearchRequest.passengers:type_name -> types.Traveller
	4, // 2: messages.FlightSearchRequest.flight_search_criterias:type_name -> types.FlightSearchCriteria
	5, // 3: messages.FlightSearchRequest.flight_offer_criterias:type_name -> types.FlightOfferCriteria
	2, // 4: messages.FlightSearchResponse.header:type_name -> types.Header
	6, // 5: messages.FlightSearchResponse.options:type_name -> types.FlightSearchOption
	0, // 6: messages.FlightSearchService.Search:input_type -> messages.FlightSearchRequest
	1, // 7: messages.FlightSearchService.Search:output_type -> messages.FlightSearchResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_flight_search_proto_init() }
func file_flight_search_proto_init() {
	if File_flight_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flight_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightSearchRequest); i {
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
		file_flight_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightSearchResponse); i {
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
			RawDescriptor: file_flight_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flight_search_proto_goTypes,
		DependencyIndexes: file_flight_search_proto_depIdxs,
		MessageInfos:      file_flight_search_proto_msgTypes,
	}.Build()
	File_flight_search_proto = out.File
	file_flight_search_proto_rawDesc = nil
	file_flight_search_proto_goTypes = nil
	file_flight_search_proto_depIdxs = nil
}