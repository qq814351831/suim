// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: v1/group.proto

package v1

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

//群组操作
type GroupHandleType int32

const (
	GroupHandleType_GROUP_UNKNOWN_HANDLE           GroupHandleType = 0  //未知
	GroupHandleType_GROUP_FOUND                    GroupHandleType = 1  //创建群
	GroupHandleType_GROUP_DISSOLVE                 GroupHandleType = 2  //解散群
	GroupHandleType_GROUP_INVITATION_FRIEND        GroupHandleType = 3  //邀请好友进群
	GroupHandleType_GROUP_MODIFICATION_NAME        GroupHandleType = 4  //修改群名称
	GroupHandleType_GROUP_EDIT_PROCLAMATION        GroupHandleType = 5  //编辑群公告
	GroupHandleType_GROUP_LOOK_PROCLAMATION        GroupHandleType = 6  //查看群公告
	GroupHandleType_GROUP_REMARK                   GroupHandleType = 7  //群名备注
	GroupHandleType_GROUP_NO_DISTURB               GroupHandleType = 8  //群消息免打扰
	GroupHandleType_GROUP_QUIT                     GroupHandleType = 9  //退出群
	GroupHandleType_GROUP_TOP                      GroupHandleType = 10 //置顶群
	GroupHandleType_GROUP_MODIFICATION_MY_NICKNAME GroupHandleType = 11 //修改我在群的昵称
	GroupHandleType_GROUP_SHOW_ALL_NICKNAME        GroupHandleType = 12 //显示群成员昵称
	GroupHandleType_GROUP_LIST                     GroupHandleType = 13 //群成员列表
	GroupHandleType_GROUP_ADD_FRIEND               GroupHandleType = 14 //添加群好友
	GroupHandleType_GROUP_CHAT_RECORD              GroupHandleType = 15 //查找群聊天记录
)

// Enum value maps for GroupHandleType.
var (
	GroupHandleType_name = map[int32]string{
		0:  "GROUP_UNKNOWN_HANDLE",
		1:  "GROUP_FOUND",
		2:  "GROUP_DISSOLVE",
		3:  "GROUP_INVITATION_FRIEND",
		4:  "GROUP_MODIFICATION_NAME",
		5:  "GROUP_EDIT_PROCLAMATION",
		6:  "GROUP_LOOK_PROCLAMATION",
		7:  "GROUP_REMARK",
		8:  "GROUP_NO_DISTURB",
		9:  "GROUP_QUIT",
		10: "GROUP_TOP",
		11: "GROUP_MODIFICATION_MY_NICKNAME",
		12: "GROUP_SHOW_ALL_NICKNAME",
		13: "GROUP_LIST",
		14: "GROUP_ADD_FRIEND",
		15: "GROUP_CHAT_RECORD",
	}
	GroupHandleType_value = map[string]int32{
		"GROUP_UNKNOWN_HANDLE":           0,
		"GROUP_FOUND":                    1,
		"GROUP_DISSOLVE":                 2,
		"GROUP_INVITATION_FRIEND":        3,
		"GROUP_MODIFICATION_NAME":        4,
		"GROUP_EDIT_PROCLAMATION":        5,
		"GROUP_LOOK_PROCLAMATION":        6,
		"GROUP_REMARK":                   7,
		"GROUP_NO_DISTURB":               8,
		"GROUP_QUIT":                     9,
		"GROUP_TOP":                      10,
		"GROUP_MODIFICATION_MY_NICKNAME": 11,
		"GROUP_SHOW_ALL_NICKNAME":        12,
		"GROUP_LIST":                     13,
		"GROUP_ADD_FRIEND":               14,
		"GROUP_CHAT_RECORD":              15,
	}
)

func (x GroupHandleType) Enum() *GroupHandleType {
	p := new(GroupHandleType)
	*p = x
	return p
}

func (x GroupHandleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupHandleType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_group_proto_enumTypes[0].Descriptor()
}

func (GroupHandleType) Type() protoreflect.EnumType {
	return &file_v1_group_proto_enumTypes[0]
}

func (x GroupHandleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupHandleType.Descriptor instead.
func (GroupHandleType) EnumDescriptor() ([]byte, []int) {
	return file_v1_group_proto_rawDescGZIP(), []int{0}
}

type GroupPackageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type GroupHandleType `protobuf:"varint,1,opt,name=type,proto3,enum=GroupHandleType" json:"type,omitempty"`
	Data []byte          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GroupPackageData) Reset() {
	*x = GroupPackageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupPackageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupPackageData) ProtoMessage() {}

func (x *GroupPackageData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupPackageData.ProtoReflect.Descriptor instead.
func (*GroupPackageData) Descriptor() ([]byte, []int) {
	return file_v1_group_proto_rawDescGZIP(), []int{0}
}

func (x *GroupPackageData) GetType() GroupHandleType {
	if x != nil {
		return x.Type
	}
	return GroupHandleType_GROUP_UNKNOWN_HANDLE
}

func (x *GroupPackageData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_v1_group_proto protoreflect.FileDescriptor

var file_v1_group_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4c, 0x0a, 0x10, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x89,
	0x03, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a,
	0x0e, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x44, 0x49, 0x53, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x10,
	0x02, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x1b,
	0x0a, 0x17, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x4c, 0x41,
	0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x5f, 0x4c, 0x4f, 0x4f, 0x4b, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x4c, 0x41, 0x4d, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x52,
	0x45, 0x4d, 0x41, 0x52, 0x4b, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x5f, 0x4e, 0x4f, 0x5f, 0x44, 0x49, 0x53, 0x54, 0x55, 0x52, 0x42, 0x10, 0x08, 0x12, 0x0e, 0x0a,
	0x0a, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x51, 0x55, 0x49, 0x54, 0x10, 0x09, 0x12, 0x0d, 0x0a,
	0x09, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x54, 0x4f, 0x50, 0x10, 0x0a, 0x12, 0x22, 0x0a, 0x1e,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4d, 0x59, 0x5f, 0x4e, 0x49, 0x43, 0x4b, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0b,
	0x12, 0x1b, 0x0a, 0x17, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x53, 0x48, 0x4f, 0x57, 0x5f, 0x41,
	0x4c, 0x4c, 0x5f, 0x4e, 0x49, 0x43, 0x4b, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0c, 0x12, 0x0e, 0x0a,
	0x0a, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0d, 0x12, 0x14, 0x0a,
	0x10, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e,
	0x44, 0x10, 0x0e, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x48, 0x41,
	0x54, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x10, 0x0f, 0x42, 0x1e, 0x5a, 0x1c, 0x73, 0x75,
	0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_v1_group_proto_rawDescOnce sync.Once
	file_v1_group_proto_rawDescData = file_v1_group_proto_rawDesc
)

func file_v1_group_proto_rawDescGZIP() []byte {
	file_v1_group_proto_rawDescOnce.Do(func() {
		file_v1_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_group_proto_rawDescData)
	})
	return file_v1_group_proto_rawDescData
}

var file_v1_group_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_group_proto_goTypes = []interface{}{
	(GroupHandleType)(0),     // 0: GroupHandleType
	(*GroupPackageData)(nil), // 1: GroupPackageData
}
var file_v1_group_proto_depIdxs = []int32{
	0, // 0: GroupPackageData.type:type_name -> GroupHandleType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_group_proto_init() }
func file_v1_group_proto_init() {
	if File_v1_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupPackageData); i {
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
			RawDescriptor: file_v1_group_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_group_proto_goTypes,
		DependencyIndexes: file_v1_group_proto_depIdxs,
		EnumInfos:         file_v1_group_proto_enumTypes,
		MessageInfos:      file_v1_group_proto_msgTypes,
	}.Build()
	File_v1_group_proto = out.File
	file_v1_group_proto_rawDesc = nil
	file_v1_group_proto_goTypes = nil
	file_v1_group_proto_depIdxs = nil
}
