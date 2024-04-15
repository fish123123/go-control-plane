// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.4
// source: envoy/extensions/filters/http/composite/v3/composite.proto

package compositev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// :ref:`Composite filter <config_http_filters_composite>` config. The composite filter config
// allows delegating filter handling to another filter as determined by matching on the request
// headers. This makes it possible to use different filters or filter configurations based on the
// incoming request.
//
// This is intended to be used with
// :ref:`ExtensionWithMatcher <envoy_v3_api_msg_extensions.common.matching.v3.ExtensionWithMatcher>`
// where a match tree is specified that indicates (via
// :ref:`ExecuteFilterAction <envoy_v3_api_msg_extensions.filters.http.composite.v3.ExecuteFilterAction>`)
// which filter configuration to create and delegate to.
type Composite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Composite) Reset() {
	*x = Composite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_composite_v3_composite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Composite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Composite) ProtoMessage() {}

func (x *Composite) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_composite_v3_composite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Composite.ProtoReflect.Descriptor instead.
func (*Composite) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDescGZIP(), []int{0}
}

// Configuration for an extension configuration discovery service with name.
type DynamicConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the extension configuration. It also serves as a resource name in ExtensionConfigDS.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Configuration source specifier for an extension configuration discovery
	// service. In case of a failure and without the default configuration,
	// 500(Internal Server Error) will be returned.
	ConfigDiscovery *v3.ExtensionConfigSource `protobuf:"bytes,2,opt,name=config_discovery,json=configDiscovery,proto3" json:"config_discovery,omitempty"`
}

func (x *DynamicConfig) Reset() {
	*x = DynamicConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_composite_v3_composite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicConfig) ProtoMessage() {}

func (x *DynamicConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_composite_v3_composite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicConfig.ProtoReflect.Descriptor instead.
func (*DynamicConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDescGZIP(), []int{1}
}

func (x *DynamicConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DynamicConfig) GetConfigDiscovery() *v3.ExtensionConfigSource {
	if x != nil {
		return x.ConfigDiscovery
	}
	return nil
}

// Composite match action (see :ref:`matching docs <arch_overview_matching_api>` for more info on match actions).
// This specifies the filter configuration of the filter that the composite filter should delegate filter interactions to.
type ExecuteFilterAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Filter specific configuration which depends on the filter being
	// instantiated. See the supported filters for further documentation.
	// Only one of “typed_config“ or “dynamic_config“ can be set.
	// [#extension-category: envoy.filters.http]
	TypedConfig *v3.TypedExtensionConfig `protobuf:"bytes,1,opt,name=typed_config,json=typedConfig,proto3" json:"typed_config,omitempty"`
	// Dynamic configuration of filter obtained via extension configuration discovery service.
	// Only one of “typed_config“ or “dynamic_config“ can be set.
	DynamicConfig *DynamicConfig `protobuf:"bytes,2,opt,name=dynamic_config,json=dynamicConfig,proto3" json:"dynamic_config,omitempty"`
}

func (x *ExecuteFilterAction) Reset() {
	*x = ExecuteFilterAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_composite_v3_composite_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteFilterAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteFilterAction) ProtoMessage() {}

func (x *ExecuteFilterAction) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_composite_v3_composite_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteFilterAction.ProtoReflect.Descriptor instead.
func (*ExecuteFilterAction) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDescGZIP(), []int{2}
}

func (x *ExecuteFilterAction) GetTypedConfig() *v3.TypedExtensionConfig {
	if x != nil {
		return x.TypedConfig
	}
	return nil
}

func (x *ExecuteFilterAction) GetDynamicConfig() *DynamicConfig {
	if x != nil {
		return x.DynamicConfig
	}
	return nil
}

var File_envoy_extensions_filters_http_composite_v3_composite_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x0b, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x22, 0x84, 0x01,
	0x0a, 0x0d, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x56, 0x0a, 0x10,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x22, 0xf0, 0x01, 0x0a, 0x13, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a, 0x0c,
	0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x13,
	0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x0d, 0x12, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x75, 0x0a, 0x0e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x13, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x0d, 0x12, 0x0b, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xb3, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x0a, 0x38, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x2f, 0x76,
	0x33, 0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x76, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDescData = file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDesc
)

func file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDescData
}

var file_envoy_extensions_filters_http_composite_v3_composite_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_filters_http_composite_v3_composite_proto_goTypes = []interface{}{
	(*Composite)(nil),                // 0: envoy.extensions.filters.http.composite.v3.Composite
	(*DynamicConfig)(nil),            // 1: envoy.extensions.filters.http.composite.v3.DynamicConfig
	(*ExecuteFilterAction)(nil),      // 2: envoy.extensions.filters.http.composite.v3.ExecuteFilterAction
	(*v3.ExtensionConfigSource)(nil), // 3: envoy.config.core.v3.ExtensionConfigSource
	(*v3.TypedExtensionConfig)(nil),  // 4: envoy.config.core.v3.TypedExtensionConfig
}
var file_envoy_extensions_filters_http_composite_v3_composite_proto_depIdxs = []int32{
	3, // 0: envoy.extensions.filters.http.composite.v3.DynamicConfig.config_discovery:type_name -> envoy.config.core.v3.ExtensionConfigSource
	4, // 1: envoy.extensions.filters.http.composite.v3.ExecuteFilterAction.typed_config:type_name -> envoy.config.core.v3.TypedExtensionConfig
	1, // 2: envoy.extensions.filters.http.composite.v3.ExecuteFilterAction.dynamic_config:type_name -> envoy.extensions.filters.http.composite.v3.DynamicConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_composite_v3_composite_proto_init() }
func file_envoy_extensions_filters_http_composite_v3_composite_proto_init() {
	if File_envoy_extensions_filters_http_composite_v3_composite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_composite_v3_composite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Composite); i {
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
		file_envoy_extensions_filters_http_composite_v3_composite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicConfig); i {
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
		file_envoy_extensions_filters_http_composite_v3_composite_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteFilterAction); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_composite_v3_composite_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_composite_v3_composite_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_composite_v3_composite_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_composite_v3_composite_proto = out.File
	file_envoy_extensions_filters_http_composite_v3_composite_proto_rawDesc = nil
	file_envoy_extensions_filters_http_composite_v3_composite_proto_goTypes = nil
	file_envoy_extensions_filters_http_composite_v3_composite_proto_depIdxs = nil
}
