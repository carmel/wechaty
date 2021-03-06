// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: tag.proto

package puppet

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type TagContactAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ContactId string `protobuf:"bytes,2,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
}

func (x *TagContactAddRequest) Reset() {
	*x = TagContactAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagContactAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagContactAddRequest) ProtoMessage() {}

func (x *TagContactAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagContactAddRequest.ProtoReflect.Descriptor instead.
func (*TagContactAddRequest) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{0}
}

func (x *TagContactAddRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TagContactAddRequest) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

type TagContactAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TagContactAddResponse) Reset() {
	*x = TagContactAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagContactAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagContactAddResponse) ProtoMessage() {}

func (x *TagContactAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagContactAddResponse.ProtoReflect.Descriptor instead.
func (*TagContactAddResponse) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{1}
}

type TagContactRemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ContactId string `protobuf:"bytes,2,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
}

func (x *TagContactRemoveRequest) Reset() {
	*x = TagContactRemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagContactRemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagContactRemoveRequest) ProtoMessage() {}

func (x *TagContactRemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagContactRemoveRequest.ProtoReflect.Descriptor instead.
func (*TagContactRemoveRequest) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{2}
}

func (x *TagContactRemoveRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TagContactRemoveRequest) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

type TagContactRemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TagContactRemoveResponse) Reset() {
	*x = TagContactRemoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagContactRemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagContactRemoveResponse) ProtoMessage() {}

func (x *TagContactRemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagContactRemoveResponse.ProtoReflect.Descriptor instead.
func (*TagContactRemoveResponse) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{3}
}

type TagContactDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TagContactDeleteRequest) Reset() {
	*x = TagContactDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagContactDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagContactDeleteRequest) ProtoMessage() {}

func (x *TagContactDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagContactDeleteRequest.ProtoReflect.Descriptor instead.
func (*TagContactDeleteRequest) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{4}
}

func (x *TagContactDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TagContactDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TagContactDeleteResponse) Reset() {
	*x = TagContactDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagContactDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagContactDeleteResponse) ProtoMessage() {}

func (x *TagContactDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagContactDeleteResponse.ProtoReflect.Descriptor instead.
func (*TagContactDeleteResponse) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{5}
}

type TagContactListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactId *wrappers.StringValue `protobuf:"bytes,1,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
}

func (x *TagContactListRequest) Reset() {
	*x = TagContactListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagContactListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagContactListRequest) ProtoMessage() {}

func (x *TagContactListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagContactListRequest.ProtoReflect.Descriptor instead.
func (*TagContactListRequest) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{6}
}

func (x *TagContactListRequest) GetContactId() *wrappers.StringValue {
	if x != nil {
		return x.ContactId
	}
	return nil
}

type TagContactListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *TagContactListResponse) Reset() {
	*x = TagContactListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagContactListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagContactListResponse) ProtoMessage() {}

func (x *TagContactListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagContactListResponse.ProtoReflect.Descriptor instead.
func (*TagContactListResponse) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{7}
}

func (x *TagContactListResponse) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_tag_proto protoreflect.FileDescriptor

var file_tag_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x65, 0x63,
	0x68, 0x61, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x14, 0x54,
	0x61, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x48, 0x0a, 0x17, 0x54,
	0x61, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x29, 0x0a, 0x17, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x18,
	0x54, 0x61, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x0a, 0x15, 0x54, 0x61, 0x67, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x22, 0x2a,
	0x0a, 0x16, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x42, 0x4a, 0x0a, 0x1d, 0x69, 0x6f,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f,
	0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tag_proto_rawDescOnce sync.Once
	file_tag_proto_rawDescData = file_tag_proto_rawDesc
)

func file_tag_proto_rawDescGZIP() []byte {
	file_tag_proto_rawDescOnce.Do(func() {
		file_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_tag_proto_rawDescData)
	})
	return file_tag_proto_rawDescData
}

var file_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tag_proto_goTypes = []interface{}{
	(*TagContactAddRequest)(nil),     // 0: wechaty.puppet.TagContactAddRequest
	(*TagContactAddResponse)(nil),    // 1: wechaty.puppet.TagContactAddResponse
	(*TagContactRemoveRequest)(nil),  // 2: wechaty.puppet.TagContactRemoveRequest
	(*TagContactRemoveResponse)(nil), // 3: wechaty.puppet.TagContactRemoveResponse
	(*TagContactDeleteRequest)(nil),  // 4: wechaty.puppet.TagContactDeleteRequest
	(*TagContactDeleteResponse)(nil), // 5: wechaty.puppet.TagContactDeleteResponse
	(*TagContactListRequest)(nil),    // 6: wechaty.puppet.TagContactListRequest
	(*TagContactListResponse)(nil),   // 7: wechaty.puppet.TagContactListResponse
	(*wrappers.StringValue)(nil),     // 8: google.protobuf.StringValue
}
var file_tag_proto_depIdxs = []int32{
	8, // 0: wechaty.puppet.TagContactListRequest.contact_id:type_name -> google.protobuf.StringValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tag_proto_init() }
func file_tag_proto_init() {
	if File_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagContactAddRequest); i {
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
		file_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagContactAddResponse); i {
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
		file_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagContactRemoveRequest); i {
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
		file_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagContactRemoveResponse); i {
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
		file_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagContactDeleteRequest); i {
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
		file_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagContactDeleteResponse); i {
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
		file_tag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagContactListRequest); i {
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
		file_tag_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagContactListResponse); i {
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
			RawDescriptor: file_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tag_proto_goTypes,
		DependencyIndexes: file_tag_proto_depIdxs,
		MessageInfos:      file_tag_proto_msgTypes,
	}.Build()
	File_tag_proto = out.File
	file_tag_proto_rawDesc = nil
	file_tag_proto_goTypes = nil
	file_tag_proto_depIdxs = nil
}
