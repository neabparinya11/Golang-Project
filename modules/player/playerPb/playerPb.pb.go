// Version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0--rc2
// source: modules/player/playerPb/playerPb.proto

package Golang_Project

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

// Structure
type PlayerProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	RoleCode int32  `protobuf:"varint,4,opt,name=roleCode,proto3" json:"roleCode,omitempty"`
	CreateAt string `protobuf:"bytes,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt string `protobuf:"bytes,6,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
}

func (x *PlayerProfile) Reset() {
	*x = PlayerProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerProfile) ProtoMessage() {}

func (x *PlayerProfile) ProtoReflect() protoreflect.Message {
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerProfile.ProtoReflect.Descriptor instead.
func (*PlayerProfile) Descriptor() ([]byte, []int) {
	return file_modules_player_playerPb_playerPb_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlayerProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PlayerProfile) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PlayerProfile) GetRoleCode() int32 {
	if x != nil {
		return x.RoleCode
	}
	return 0
}

func (x *PlayerProfile) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *PlayerProfile) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

type CreadentialSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreadentialSearchRequest) Reset() {
	*x = CreadentialSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreadentialSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreadentialSearchRequest) ProtoMessage() {}

func (x *CreadentialSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreadentialSearchRequest.ProtoReflect.Descriptor instead.
func (*CreadentialSearchRequest) Descriptor() ([]byte, []int) {
	return file_modules_player_playerPb_playerPb_proto_rawDescGZIP(), []int{1}
}

func (x *CreadentialSearchRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreadentialSearchRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type FindOnePlayerProfileToRefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *FindOnePlayerProfileToRefreshRequest) Reset() {
	*x = FindOnePlayerProfileToRefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOnePlayerProfileToRefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOnePlayerProfileToRefreshRequest) ProtoMessage() {}

func (x *FindOnePlayerProfileToRefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOnePlayerProfileToRefreshRequest.ProtoReflect.Descriptor instead.
func (*FindOnePlayerProfileToRefreshRequest) Descriptor() ([]byte, []int) {
	return file_modules_player_playerPb_playerPb_proto_rawDescGZIP(), []int{2}
}

func (x *FindOnePlayerProfileToRefreshRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GetPlayerSavingAccoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *GetPlayerSavingAccoutRequest) Reset() {
	*x = GetPlayerSavingAccoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerSavingAccoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerSavingAccoutRequest) ProtoMessage() {}

func (x *GetPlayerSavingAccoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerSavingAccoutRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerSavingAccoutRequest) Descriptor() ([]byte, []int) {
	return file_modules_player_playerPb_playerPb_proto_rawDescGZIP(), []int{3}
}

func (x *GetPlayerSavingAccoutRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GetPlayerSavingAccoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string  `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Balance  float64 `protobuf:"fixed64,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *GetPlayerSavingAccoutResponse) Reset() {
	*x = GetPlayerSavingAccoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerSavingAccoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerSavingAccoutResponse) ProtoMessage() {}

func (x *GetPlayerSavingAccoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_player_playerPb_playerPb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerSavingAccoutResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerSavingAccoutResponse) Descriptor() ([]byte, []int) {
	return file_modules_player_playerPb_playerPb_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlayerSavingAccoutResponse) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *GetPlayerSavingAccoutResponse) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

var File_modules_player_playerPb_playerPb_proto protoreflect.FileDescriptor

var file_modules_player_playerPb_playerPb_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x62, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x50, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x72, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x22, 0x4c, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x42, 0x0a, 0x24, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x55, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x61, 0x76,
	0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x83, 0x02, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x56, 0x0a,
	0x1d, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x25,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x56, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x74, 0x12, 0x1d,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x61, 0x62,
	0x70, 0x61, 0x72, 0x69, 0x6e, 0x79, 0x61, 0x31, 0x31, 0x2f, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_player_playerPb_playerPb_proto_rawDescOnce sync.Once
	file_modules_player_playerPb_playerPb_proto_rawDescData = file_modules_player_playerPb_playerPb_proto_rawDesc
)

func file_modules_player_playerPb_playerPb_proto_rawDescGZIP() []byte {
	file_modules_player_playerPb_playerPb_proto_rawDescOnce.Do(func() {
		file_modules_player_playerPb_playerPb_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_player_playerPb_playerPb_proto_rawDescData)
	})
	return file_modules_player_playerPb_playerPb_proto_rawDescData
}

var file_modules_player_playerPb_playerPb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_modules_player_playerPb_playerPb_proto_goTypes = []interface{}{
	(*PlayerProfile)(nil),                        // 0: PlayerProfile
	(*CreadentialSearchRequest)(nil),             // 1: CreadentialSearchRequest
	(*FindOnePlayerProfileToRefreshRequest)(nil), // 2: FindOnePlayerProfileToRefreshRequest
	(*GetPlayerSavingAccoutRequest)(nil),         // 3: GetPlayerSavingAccoutRequest
	(*GetPlayerSavingAccoutResponse)(nil),        // 4: GetPlayerSavingAccoutResponse
}
var file_modules_player_playerPb_playerPb_proto_depIdxs = []int32{
	1, // 0: PlayerGrpcService.CreadentialSearch:input_type -> CreadentialSearchRequest
	2, // 1: PlayerGrpcService.FindOnePlayerProfileToRefresh:input_type -> FindOnePlayerProfileToRefreshRequest
	3, // 2: PlayerGrpcService.GetPlayerSavingAccout:input_type -> GetPlayerSavingAccoutRequest
	0, // 3: PlayerGrpcService.CreadentialSearch:output_type -> PlayerProfile
	0, // 4: PlayerGrpcService.FindOnePlayerProfileToRefresh:output_type -> PlayerProfile
	4, // 5: PlayerGrpcService.GetPlayerSavingAccout:output_type -> GetPlayerSavingAccoutResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modules_player_playerPb_playerPb_proto_init() }
func file_modules_player_playerPb_playerPb_proto_init() {
	if File_modules_player_playerPb_playerPb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_player_playerPb_playerPb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerProfile); i {
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
		file_modules_player_playerPb_playerPb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreadentialSearchRequest); i {
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
		file_modules_player_playerPb_playerPb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOnePlayerProfileToRefreshRequest); i {
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
		file_modules_player_playerPb_playerPb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerSavingAccoutRequest); i {
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
		file_modules_player_playerPb_playerPb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerSavingAccoutResponse); i {
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
			RawDescriptor: file_modules_player_playerPb_playerPb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_player_playerPb_playerPb_proto_goTypes,
		DependencyIndexes: file_modules_player_playerPb_playerPb_proto_depIdxs,
		MessageInfos:      file_modules_player_playerPb_playerPb_proto_msgTypes,
	}.Build()
	File_modules_player_playerPb_playerPb_proto = out.File
	file_modules_player_playerPb_playerPb_proto_rawDesc = nil
	file_modules_player_playerPb_playerPb_proto_goTypes = nil
	file_modules_player_playerPb_playerPb_proto_depIdxs = nil
}