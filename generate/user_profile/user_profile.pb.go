// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: user_profile.proto

package user_profile

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

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_user_profile_proto_rawDescGZIP(), []int{0}
}

type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	FullName  string `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Bio       string `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
	Role      string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	Location  string `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	AvatarUrl string `protobuf:"bytes,6,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Website   string `protobuf:"bytes,7,opt,name=website,proto3" json:"website,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_user_profile_proto_rawDescGZIP(), []int{1}
}

func (x *UserProfile) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserProfile) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UserProfile) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *UserProfile) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserProfile) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UserProfile) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *UserProfile) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

type UsersProfiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersProfile []*UserProfile `protobuf:"bytes,1,rep,name=users_profile,json=usersProfile,proto3" json:"users_profile,omitempty"`
}

func (x *UsersProfiles) Reset() {
	*x = UsersProfiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersProfiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersProfiles) ProtoMessage() {}

func (x *UsersProfiles) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersProfiles.ProtoReflect.Descriptor instead.
func (*UsersProfiles) Descriptor() ([]byte, []int) {
	return file_user_profile_proto_rawDescGZIP(), []int{2}
}

func (x *UsersProfiles) GetUsersProfile() []*UserProfile {
	if x != nil {
		return x.UsersProfile
	}
	return nil
}

type FilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Arr   []string `protobuf:"bytes,2,rep,name=arr,proto3" json:"arr,omitempty"`
}

func (x *FilterRequest) Reset() {
	*x = FilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterRequest) ProtoMessage() {}

func (x *FilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterRequest.ProtoReflect.Descriptor instead.
func (*FilterRequest) Descriptor() ([]byte, []int) {
	return file_user_profile_proto_rawDescGZIP(), []int{3}
}

func (x *FilterRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *FilterRequest) GetArr() []string {
	if x != nil {
		return x.Arr
	}
	return nil
}

type UserProfileId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserProfileId) Reset() {
	*x = UserProfileId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfileId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfileId) ProtoMessage() {}

func (x *UserProfileId) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfileId.ProtoReflect.Descriptor instead.
func (*UserProfileId) Descriptor() ([]byte, []int) {
	return file_user_profile_proto_rawDescGZIP(), []int{4}
}

func (x *UserProfileId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_user_profile_proto protoreflect.FileDescriptor

var file_user_profile_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0xbc, 0x01, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x22, 0x4f, 0x0a, 0x0d, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x37, 0x0a, 0x0d, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x72, 0x72, 0x22, 0x1f, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x32, 0xbe, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3f, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x41,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_profile_proto_rawDescOnce sync.Once
	file_user_profile_proto_rawDescData = file_user_profile_proto_rawDesc
)

func file_user_profile_proto_rawDescGZIP() []byte {
	file_user_profile_proto_rawDescOnce.Do(func() {
		file_user_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_profile_proto_rawDescData)
	})
	return file_user_profile_proto_rawDescData
}

var file_user_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_profile_proto_goTypes = []interface{}{
	(*Void)(nil),          // 0: user_profile.Void
	(*UserProfile)(nil),   // 1: user_profile.UserProfile
	(*UsersProfiles)(nil), // 2: user_profile.UsersProfiles
	(*FilterRequest)(nil), // 3: user_profile.FilterRequest
	(*UserProfileId)(nil), // 4: user_profile.UserProfileId
}
var file_user_profile_proto_depIdxs = []int32{
	1, // 0: user_profile.UsersProfiles.users_profile:type_name -> user_profile.UserProfile
	1, // 1: user_profile.UserService.Create:input_type -> user_profile.UserProfile
	3, // 2: user_profile.UserService.Get:input_type -> user_profile.FilterRequest
	4, // 3: user_profile.UserService.GetByID:input_type -> user_profile.UserProfileId
	1, // 4: user_profile.UserService.Update:input_type -> user_profile.UserProfile
	4, // 5: user_profile.UserService.Delete:input_type -> user_profile.UserProfileId
	0, // 6: user_profile.UserService.Create:output_type -> user_profile.Void
	2, // 7: user_profile.UserService.Get:output_type -> user_profile.UsersProfiles
	1, // 8: user_profile.UserService.GetByID:output_type -> user_profile.UserProfile
	0, // 9: user_profile.UserService.Update:output_type -> user_profile.Void
	0, // 10: user_profile.UserService.Delete:output_type -> user_profile.Void
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_profile_proto_init() }
func file_user_profile_proto_init() {
	if File_user_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_user_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfile); i {
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
		file_user_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersProfiles); i {
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
		file_user_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterRequest); i {
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
		file_user_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfileId); i {
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
			RawDescriptor: file_user_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_profile_proto_goTypes,
		DependencyIndexes: file_user_profile_proto_depIdxs,
		MessageInfos:      file_user_profile_proto_msgTypes,
	}.Build()
	File_user_profile_proto = out.File
	file_user_profile_proto_rawDesc = nil
	file_user_profile_proto_goTypes = nil
	file_user_profile_proto_depIdxs = nil
}
