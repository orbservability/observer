// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: com/orbservability/schema/v1/event_gateway_service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_com_orbservability_schema_v1_event_gateway_service_proto protoreflect.FileDescriptor

var file_com_orbservability_schema_v1_event_gateway_service_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x6f, 0x6d, 0x2e,
	0x6f, 0x72, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x78, 0x69, 0x65, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x69, 0x0a, 0x13, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0c,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x28, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x78, 0x69,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01,
	0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x72, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_com_orbservability_schema_v1_event_gateway_service_proto_goTypes = []interface{}{
	(*PixieEvent)(nil),    // 0: com.orbservability.schema.v1.PixieEvent
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_com_orbservability_schema_v1_event_gateway_service_proto_depIdxs = []int32{
	0, // 0: com.orbservability.schema.v1.EventGatewayService.StreamEvents:input_type -> com.orbservability.schema.v1.PixieEvent
	1, // 1: com.orbservability.schema.v1.EventGatewayService.StreamEvents:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_orbservability_schema_v1_event_gateway_service_proto_init() }
func file_com_orbservability_schema_v1_event_gateway_service_proto_init() {
	if File_com_orbservability_schema_v1_event_gateway_service_proto != nil {
		return
	}
	file_com_orbservability_schema_v1_pixie_event_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_orbservability_schema_v1_event_gateway_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_orbservability_schema_v1_event_gateway_service_proto_goTypes,
		DependencyIndexes: file_com_orbservability_schema_v1_event_gateway_service_proto_depIdxs,
	}.Build()
	File_com_orbservability_schema_v1_event_gateway_service_proto = out.File
	file_com_orbservability_schema_v1_event_gateway_service_proto_rawDesc = nil
	file_com_orbservability_schema_v1_event_gateway_service_proto_goTypes = nil
	file_com_orbservability_schema_v1_event_gateway_service_proto_depIdxs = nil
}