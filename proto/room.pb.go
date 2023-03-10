// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/room.proto

package proto

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

type GetRoomsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetRoomsRequest) Reset() {
	*x = GetRoomsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomsRequest) ProtoMessage() {}

func (x *GetRoomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomsRequest.ProtoReflect.Descriptor instead.
func (*GetRoomsRequest) Descriptor() ([]byte, []int) {
	return file_proto_room_proto_rawDescGZIP(), []int{0}
}

func (x *GetRoomsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *GetRoomsResponse) Reset() {
	*x = GetRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomsResponse) ProtoMessage() {}

func (x *GetRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomsResponse.ProtoReflect.Descriptor instead.
func (*GetRoomsResponse) Descriptor() ([]byte, []int) {
	return file_proto_room_proto_rawDescGZIP(), []int{1}
}

func (x *GetRoomsResponse) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID        string `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	UserCreatorId string `protobuf:"bytes,2,opt,name=userCreatorId,proto3" json:"userCreatorId,omitempty"`
	Date          string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Place         string `protobuf:"bytes,4,opt,name=place,proto3" json:"place,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_proto_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_proto_room_proto_rawDescGZIP(), []int{2}
}

func (x *Room) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

func (x *Room) GetUserCreatorId() string {
	if x != nil {
		return x.UserCreatorId
	}
	return ""
}

func (x *Room) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Room) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

type GetUsersRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID string `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
}

func (x *GetUsersRoomRequest) Reset() {
	*x = GetUsersRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRoomRequest) ProtoMessage() {}

func (x *GetUsersRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRoomRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRoomRequest) Descriptor() ([]byte, []int) {
	return file_proto_room_proto_rawDescGZIP(), []int{3}
}

func (x *GetUsersRoomRequest) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

type GetUsersRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUsersRoomResponse) Reset() {
	*x = GetUsersRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRoomResponse) ProtoMessage() {}

func (x *GetUsersRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRoomResponse.ProtoReflect.Descriptor instead.
func (*GetUsersRoomResponse) Descriptor() ([]byte, []int) {
	return file_proto_room_proto_rawDescGZIP(), []int{4}
}

func (x *GetUsersRoomResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_room_proto_msgTypes[5]
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
	return file_proto_room_proto_rawDescGZIP(), []int{5}
}

func (x *User) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_proto_room_proto protoreflect.FileDescriptor

var file_proto_room_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x22, 0x34, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x6e, 0x0a, 0x04, 0x52, 0x6f, 0x6f,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x22, 0x38, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x20, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x40, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x32, 0x8c, 0x01, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x3b, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x15, 0x2e, 0x72, 0x6f, 0x6f, 0x6d,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x19, 0x2e, 0x72, 0x6f, 0x6f,
	0x6d, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x49, 0x76, 0x61, 0x6e, 0x56, 0x6f, 0x6a, 0x6e, 0x69, 0x63, 0x2f, 0x62, 0x61, 0x6e,
	0x64, 0x45, 0x46, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_room_proto_rawDescOnce sync.Once
	file_proto_room_proto_rawDescData = file_proto_room_proto_rawDesc
)

func file_proto_room_proto_rawDescGZIP() []byte {
	file_proto_room_proto_rawDescOnce.Do(func() {
		file_proto_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_room_proto_rawDescData)
	})
	return file_proto_room_proto_rawDescData
}

var file_proto_room_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_room_proto_goTypes = []interface{}{
	(*GetRoomsRequest)(nil),      // 0: room.GetRoomsRequest
	(*GetRoomsResponse)(nil),     // 1: room.GetRoomsResponse
	(*Room)(nil),                 // 2: room.Room
	(*GetUsersRoomRequest)(nil),  // 3: room.GetUsersRoomRequest
	(*GetUsersRoomResponse)(nil), // 4: room.GetUsersRoomResponse
	(*User)(nil),                 // 5: room.User
}
var file_proto_room_proto_depIdxs = []int32{
	2, // 0: room.GetRoomsResponse.rooms:type_name -> room.Room
	5, // 1: room.GetUsersRoomResponse.users:type_name -> room.User
	0, // 2: room.room.GetRooms:input_type -> room.GetRoomsRequest
	3, // 3: room.room.GetRoomsUser:input_type -> room.GetUsersRoomRequest
	1, // 4: room.room.GetRooms:output_type -> room.GetRoomsResponse
	4, // 5: room.room.GetRoomsUser:output_type -> room.GetUsersRoomResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_room_proto_init() }
func file_proto_room_proto_init() {
	if File_proto_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomsRequest); i {
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
		file_proto_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomsResponse); i {
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
		file_proto_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_proto_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRoomRequest); i {
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
		file_proto_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRoomResponse); i {
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
		file_proto_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proto_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_room_proto_goTypes,
		DependencyIndexes: file_proto_room_proto_depIdxs,
		MessageInfos:      file_proto_room_proto_msgTypes,
	}.Build()
	File_proto_room_proto = out.File
	file_proto_room_proto_rawDesc = nil
	file_proto_room_proto_goTypes = nil
	file_proto_room_proto_depIdxs = nil
}
