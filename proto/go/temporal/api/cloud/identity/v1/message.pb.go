// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: temporal/api/cloud/identity/v1/message.proto

package identityv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccountAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The role on the account, should be one of [admin, developer, read]
	// admin - gives full access the account, including users and namespaces
	// developer - gives access to create namespaces on the account
	// read - gives read only access to the account
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *AccountAccess) Reset() {
	*x = AccountAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_cloud_identity_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountAccess) ProtoMessage() {}

func (x *AccountAccess) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_identity_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountAccess.ProtoReflect.Descriptor instead.
func (*AccountAccess) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_identity_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *AccountAccess) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type NamespaceAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The permission to the namespace, should be one of [admin, write, read]
	// admin - gives full access to the namespace, including assigning namespace access to other users
	// write - gives write access to the namespace configuration and workflows within the namespace
	// read - gives read only access to the namespace configuration and workflows within the namespace
	Permission string `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *NamespaceAccess) Reset() {
	*x = NamespaceAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_cloud_identity_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceAccess) ProtoMessage() {}

func (x *NamespaceAccess) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_identity_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceAccess.ProtoReflect.Descriptor instead.
func (*NamespaceAccess) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_identity_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *NamespaceAccess) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

type Access struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account access
	AccountAccess *AccountAccess `protobuf:"bytes,1,opt,name=account_access,json=accountAccess,proto3" json:"account_access,omitempty"`
	// The map of namespace accesses
	// The key is the namespace name and the value is the access to the namespace
	NamespaceAccesses map[string]*NamespaceAccess `protobuf:"bytes,2,rep,name=namespace_accesses,json=namespaceAccesses,proto3" json:"namespace_accesses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Access) Reset() {
	*x = Access{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_cloud_identity_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Access) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Access) ProtoMessage() {}

func (x *Access) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_identity_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Access.ProtoReflect.Descriptor instead.
func (*Access) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_identity_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *Access) GetAccountAccess() *AccountAccess {
	if x != nil {
		return x.AccountAccess
	}
	return nil
}

func (x *Access) GetNamespaceAccesses() map[string]*NamespaceAccess {
	if x != nil {
		return x.NamespaceAccesses
	}
	return nil
}

type UserSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The email address associated to the user
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// The access to assigned to the user
	Access *Access `protobuf:"bytes,2,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *UserSpec) Reset() {
	*x = UserSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_cloud_identity_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSpec) ProtoMessage() {}

func (x *UserSpec) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_identity_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSpec.ProtoReflect.Descriptor instead.
func (*UserSpec) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_identity_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *UserSpec) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserSpec) GetAccess() *Access {
	if x != nil {
		return x.Access
	}
	return nil
}

type Invitation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The date and time when the user was created
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// The date and time when the invitation expires or has expired
	ExpiredTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expired_time,json=expiredTime,proto3" json:"expired_time,omitempty"`
}

func (x *Invitation) Reset() {
	*x = Invitation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_cloud_identity_v1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invitation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invitation) ProtoMessage() {}

func (x *Invitation) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_identity_v1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invitation.ProtoReflect.Descriptor instead.
func (*Invitation) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_identity_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *Invitation) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Invitation) GetExpiredTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredTime
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the user
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The current version of the user specification
	// The next update operation will have to include this version
	ResourceVersion string `protobuf:"bytes,2,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	// The user specification
	Spec *UserSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	// The current state of the user
	State string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	// The id of the async operation that is creating/updating/deleting the user, if any
	AsyncOperationId string `protobuf:"bytes,5,opt,name=async_operation_id,json=asyncOperationId,proto3" json:"async_operation_id,omitempty"`
	// The details of the open invitation sent to the user, if any
	Invitation *Invitation `protobuf:"bytes,6,opt,name=invitation,proto3" json:"invitation,omitempty"`
	// The date and time when the user was created
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// The date and time when the user was last modified
	// Will not be set if the user has never been modified.
	LastModifiedTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=last_modified_time,json=lastModifiedTime,proto3" json:"last_modified_time,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_cloud_identity_v1_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_identity_v1_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_identity_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

func (x *User) GetSpec() *UserSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *User) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *User) GetAsyncOperationId() string {
	if x != nil {
		return x.AsyncOperationId
	}
	return ""
}

func (x *User) GetInvitation() *Invitation {
	if x != nil {
		return x.Invitation
	}
	return nil
}

func (x *User) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *User) GetLastModifiedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastModifiedTime
	}
	return nil
}

var File_temporal_api_cloud_identity_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_cloud_identity_v1_message_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x23, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x22, 0x31, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xc3, 0x02, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x54, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x6c, 0x0a, 0x12, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x11, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x1a, 0x75, 0x0a, 0x16, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x45, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x60, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x3e, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x8a, 0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d,
	0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a,
	0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x98, 0x03, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x3c, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x4a, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d,
	0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x48, 0x0a,
	0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0xba, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x69,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x54, 0x41, 0x43, 0x49,
	0xaa, 0x02, 0x1e, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x1e, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x2a, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x5c, 0x41, 0x70,
	0x69, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x22, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_cloud_identity_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_cloud_identity_v1_message_proto_rawDescData = file_temporal_api_cloud_identity_v1_message_proto_rawDesc
)

func file_temporal_api_cloud_identity_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_cloud_identity_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_cloud_identity_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_cloud_identity_v1_message_proto_rawDescData)
	})
	return file_temporal_api_cloud_identity_v1_message_proto_rawDescData
}

var file_temporal_api_cloud_identity_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_temporal_api_cloud_identity_v1_message_proto_goTypes = []interface{}{
	(*AccountAccess)(nil),         // 0: temporal.api.cloud.identity.v1.AccountAccess
	(*NamespaceAccess)(nil),       // 1: temporal.api.cloud.identity.v1.NamespaceAccess
	(*Access)(nil),                // 2: temporal.api.cloud.identity.v1.Access
	(*UserSpec)(nil),              // 3: temporal.api.cloud.identity.v1.UserSpec
	(*Invitation)(nil),            // 4: temporal.api.cloud.identity.v1.Invitation
	(*User)(nil),                  // 5: temporal.api.cloud.identity.v1.User
	nil,                           // 6: temporal.api.cloud.identity.v1.Access.NamespaceAccessesEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_temporal_api_cloud_identity_v1_message_proto_depIdxs = []int32{
	0,  // 0: temporal.api.cloud.identity.v1.Access.account_access:type_name -> temporal.api.cloud.identity.v1.AccountAccess
	6,  // 1: temporal.api.cloud.identity.v1.Access.namespace_accesses:type_name -> temporal.api.cloud.identity.v1.Access.NamespaceAccessesEntry
	2,  // 2: temporal.api.cloud.identity.v1.UserSpec.access:type_name -> temporal.api.cloud.identity.v1.Access
	7,  // 3: temporal.api.cloud.identity.v1.Invitation.created_time:type_name -> google.protobuf.Timestamp
	7,  // 4: temporal.api.cloud.identity.v1.Invitation.expired_time:type_name -> google.protobuf.Timestamp
	3,  // 5: temporal.api.cloud.identity.v1.User.spec:type_name -> temporal.api.cloud.identity.v1.UserSpec
	4,  // 6: temporal.api.cloud.identity.v1.User.invitation:type_name -> temporal.api.cloud.identity.v1.Invitation
	7,  // 7: temporal.api.cloud.identity.v1.User.created_time:type_name -> google.protobuf.Timestamp
	7,  // 8: temporal.api.cloud.identity.v1.User.last_modified_time:type_name -> google.protobuf.Timestamp
	1,  // 9: temporal.api.cloud.identity.v1.Access.NamespaceAccessesEntry.value:type_name -> temporal.api.cloud.identity.v1.NamespaceAccess
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_temporal_api_cloud_identity_v1_message_proto_init() }
func file_temporal_api_cloud_identity_v1_message_proto_init() {
	if File_temporal_api_cloud_identity_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_api_cloud_identity_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountAccess); i {
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
		file_temporal_api_cloud_identity_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceAccess); i {
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
		file_temporal_api_cloud_identity_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Access); i {
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
		file_temporal_api_cloud_identity_v1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSpec); i {
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
		file_temporal_api_cloud_identity_v1_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invitation); i {
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
		file_temporal_api_cloud_identity_v1_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_temporal_api_cloud_identity_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_cloud_identity_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_cloud_identity_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_cloud_identity_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_cloud_identity_v1_message_proto = out.File
	file_temporal_api_cloud_identity_v1_message_proto_rawDesc = nil
	file_temporal_api_cloud_identity_v1_message_proto_goTypes = nil
	file_temporal_api_cloud_identity_v1_message_proto_depIdxs = nil
}