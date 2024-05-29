// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: alexandria/catalogue/v1/catalogue.proto

package cataloguev1

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

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Created    int64  `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	CreatedBy  string `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Updated    int64  `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	ModifiedBy string `protobuf:"bytes,6,opt,name=modified_by,json=modifiedBy,proto3" json:"modified_by,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_alexandria_catalogue_v1_catalogue_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Author) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Author) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Author) GetModifiedBy() string {
	if x != nil {
		return x.ModifiedBy
	}
	return ""
}

type ItemInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type        string    `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Categories  []string  `protobuf:"bytes,4,rep,name=categories,proto3" json:"categories,omitempty"`
	Tags        []string  `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Authors     []*Author `protobuf:"bytes,6,rep,name=authors,proto3" json:"authors,omitempty"`
}

func (x *ItemInput) Reset() {
	*x = ItemInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemInput) ProtoMessage() {}

func (x *ItemInput) ProtoReflect() protoreflect.Message {
	mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemInput.ProtoReflect.Descriptor instead.
func (*ItemInput) Descriptor() ([]byte, []int) {
	return file_alexandria_catalogue_v1_catalogue_proto_rawDescGZIP(), []int{1}
}

func (x *ItemInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ItemInput) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ItemInput) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ItemInput) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *ItemInput) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ItemInput) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string          `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        string          `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Categories  []string        `protobuf:"bytes,5,rep,name=categories,proto3" json:"categories,omitempty"`
	Tags        []string        `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	Authors     []*Author       `protobuf:"bytes,7,rep,name=authors,proto3" json:"authors,omitempty"`
	Created     int64           `protobuf:"varint,8,opt,name=created,proto3" json:"created,omitempty"`
	CreatedBy   string          `protobuf:"bytes,9,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Updated     int64           `protobuf:"varint,10,opt,name=updated,proto3" json:"updated,omitempty"`
	ModifiedBy  string          `protobuf:"bytes,11,opt,name=modified_by,json=modifiedBy,proto3" json:"modified_by,omitempty"`
	Shelf       *ShelfReference `protobuf:"bytes,12,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_alexandria_catalogue_v1_catalogue_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Item) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Item) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Item) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Item) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *Item) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Item) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Item) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Item) GetModifiedBy() string {
	if x != nil {
		return x.ModifiedBy
	}
	return ""
}

func (x *Item) GetShelf() *ShelfReference {
	if x != nil {
		return x.Shelf
	}
	return nil
}

type ItemReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ItemReference) Reset() {
	*x = ItemReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemReference) ProtoMessage() {}

func (x *ItemReference) ProtoReflect() protoreflect.Message {
	mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemReference.ProtoReflect.Descriptor instead.
func (*ItemReference) Descriptor() ([]byte, []int) {
	return file_alexandria_catalogue_v1_catalogue_proto_rawDescGZIP(), []int{3}
}

func (x *ItemReference) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ItemReference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ShelfInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location    string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ShelfInput) Reset() {
	*x = ShelfInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShelfInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShelfInput) ProtoMessage() {}

func (x *ShelfInput) ProtoReflect() protoreflect.Message {
	mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShelfInput.ProtoReflect.Descriptor instead.
func (*ShelfInput) Descriptor() ([]byte, []int) {
	return file_alexandria_catalogue_v1_catalogue_proto_rawDescGZIP(), []int{4}
}

func (x *ShelfInput) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ShelfInput) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Shelf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location    string         `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Description string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Created     int64          `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	CreatedBy   string         `protobuf:"bytes,5,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Updated     int64          `protobuf:"varint,6,opt,name=updated,proto3" json:"updated,omitempty"`
	ModifiedBy  string         `protobuf:"bytes,7,opt,name=modified_by,json=modifiedBy,proto3" json:"modified_by,omitempty"`
	Room        *RoomReference `protobuf:"bytes,8,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *Shelf) Reset() {
	*x = Shelf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shelf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shelf) ProtoMessage() {}

func (x *Shelf) ProtoReflect() protoreflect.Message {
	mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shelf.ProtoReflect.Descriptor instead.
func (*Shelf) Descriptor() ([]byte, []int) {
	return file_alexandria_catalogue_v1_catalogue_proto_rawDescGZIP(), []int{5}
}

func (x *Shelf) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Shelf) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Shelf) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Shelf) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Shelf) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Shelf) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Shelf) GetModifiedBy() string {
	if x != nil {
		return x.ModifiedBy
	}
	return ""
}

func (x *Shelf) GetRoom() *RoomReference {
	if x != nil {
		return x.Room
	}
	return nil
}

type ShelfReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *ShelfReference) Reset() {
	*x = ShelfReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShelfReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShelfReference) ProtoMessage() {}

func (x *ShelfReference) ProtoReflect() protoreflect.Message {
	mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShelfReference.ProtoReflect.Descriptor instead.
func (*ShelfReference) Descriptor() ([]byte, []int) {
	return file_alexandria_catalogue_v1_catalogue_proto_rawDescGZIP(), []int{6}
}

func (x *ShelfReference) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShelfReference) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type RoomInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location    string        `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Description string        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Shelves     []*ShelfInput `protobuf:"bytes,3,rep,name=shelves,proto3" json:"shelves,omitempty"`
}

func (x *RoomInput) Reset() {
	*x = RoomInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomInput) ProtoMessage() {}

func (x *RoomInput) ProtoReflect() protoreflect.Message {
	mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomInput.ProtoReflect.Descriptor instead.
func (*RoomInput) Descriptor() ([]byte, []int) {
	return file_alexandria_catalogue_v1_catalogue_proto_rawDescGZIP(), []int{7}
}

func (x *RoomInput) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *RoomInput) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RoomInput) GetShelves() []*ShelfInput {
	if x != nil {
		return x.Shelves
	}
	return nil
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location    string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Shelves     []*Shelf `protobuf:"bytes,4,rep,name=shelves,proto3" json:"shelves,omitempty"`
	Created     int64    `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	CreatedBy   string   `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Updated     int64    `protobuf:"varint,7,opt,name=updated,proto3" json:"updated,omitempty"`
	ModifiedBy  string   `protobuf:"bytes,8,opt,name=modified_by,json=modifiedBy,proto3" json:"modified_by,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[8]
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
	return file_alexandria_catalogue_v1_catalogue_proto_rawDescGZIP(), []int{8}
}

func (x *Room) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Room) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Room) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Room) GetShelves() []*Shelf {
	if x != nil {
		return x.Shelves
	}
	return nil
}

func (x *Room) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Room) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Room) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Room) GetModifiedBy() string {
	if x != nil {
		return x.ModifiedBy
	}
	return ""
}

type RoomReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *RoomReference) Reset() {
	*x = RoomReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomReference) ProtoMessage() {}

func (x *RoomReference) ProtoReflect() protoreflect.Message {
	mi := &file_alexandria_catalogue_v1_catalogue_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomReference.ProtoReflect.Descriptor instead.
func (*RoomReference) Descriptor() ([]byte, []int) {
	return file_alexandria_catalogue_v1_catalogue_proto_rawDescGZIP(), []int{9}
}

func (x *RoomReference) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoomReference) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

var File_alexandria_catalogue_v1_catalogue_proto protoreflect.FileDescriptor

var file_alexandria_catalogue_v1_catalogue_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x72, 0x69, 0x61, 0x2f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x6c, 0x65, 0x78, 0x61,
	0x6e, 0x64, 0x72, 0x69, 0x61, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0xa0, 0x01, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x42, 0x79, 0x22, 0xc4, 0x01, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x39, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x72, 0x69, 0x61, 0x2e,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x22, 0x82, 0x03, 0x0a,
	0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x72, 0x69,
	0x61, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x61, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x72, 0x69, 0x61, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c,
	0x66, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x22, 0x33, 0x0a, 0x0d, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x0a, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x85, 0x02, 0x0a, 0x05, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x79, 0x12, 0x3a,
	0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61,
	0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x72, 0x69, 0x61, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x3c, 0x0a, 0x0e, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x09, 0x52, 0x6f, 0x6f,
	0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x07, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x72,
	0x69, 0x61, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x68, 0x65, 0x6c, 0x66, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x07, 0x73, 0x68, 0x65, 0x6c,
	0x76, 0x65, 0x73, 0x22, 0x82, 0x02, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x73, 0x68,
	0x65, 0x6c, 0x76, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x6c,
	0x65, 0x78, 0x61, 0x6e, 0x64, 0x72, 0x69, 0x61, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x07, 0x73, 0x68, 0x65,
	0x6c, 0x76, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x79, 0x22, 0x3b, 0x0a, 0x0d, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xed, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c,
	0x65, 0x78, 0x61, 0x6e, 0x64, 0x72, 0x69, 0x61, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x75, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x73, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x72, 0x69, 0x61,
	0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa,
	0x02, 0x17, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x72, 0x69, 0x61, 0x2e, 0x43, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x41, 0x6c, 0x65, 0x78,
	0x61, 0x6e, 0x64, 0x72, 0x69, 0x61, 0x5c, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x72, 0x69, 0x61,
	0x5c, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x41, 0x6c, 0x65, 0x78,
	0x61, 0x6e, 0x64, 0x72, 0x69, 0x61, 0x3a, 0x3a, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alexandria_catalogue_v1_catalogue_proto_rawDescOnce sync.Once
	file_alexandria_catalogue_v1_catalogue_proto_rawDescData = file_alexandria_catalogue_v1_catalogue_proto_rawDesc
)

func file_alexandria_catalogue_v1_catalogue_proto_rawDescGZIP() []byte {
	file_alexandria_catalogue_v1_catalogue_proto_rawDescOnce.Do(func() {
		file_alexandria_catalogue_v1_catalogue_proto_rawDescData = protoimpl.X.CompressGZIP(file_alexandria_catalogue_v1_catalogue_proto_rawDescData)
	})
	return file_alexandria_catalogue_v1_catalogue_proto_rawDescData
}

var file_alexandria_catalogue_v1_catalogue_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_alexandria_catalogue_v1_catalogue_proto_goTypes = []interface{}{
	(*Author)(nil),         // 0: alexandria.catalogue.v1.Author
	(*ItemInput)(nil),      // 1: alexandria.catalogue.v1.ItemInput
	(*Item)(nil),           // 2: alexandria.catalogue.v1.Item
	(*ItemReference)(nil),  // 3: alexandria.catalogue.v1.ItemReference
	(*ShelfInput)(nil),     // 4: alexandria.catalogue.v1.ShelfInput
	(*Shelf)(nil),          // 5: alexandria.catalogue.v1.Shelf
	(*ShelfReference)(nil), // 6: alexandria.catalogue.v1.ShelfReference
	(*RoomInput)(nil),      // 7: alexandria.catalogue.v1.RoomInput
	(*Room)(nil),           // 8: alexandria.catalogue.v1.Room
	(*RoomReference)(nil),  // 9: alexandria.catalogue.v1.RoomReference
}
var file_alexandria_catalogue_v1_catalogue_proto_depIdxs = []int32{
	0, // 0: alexandria.catalogue.v1.ItemInput.authors:type_name -> alexandria.catalogue.v1.Author
	0, // 1: alexandria.catalogue.v1.Item.authors:type_name -> alexandria.catalogue.v1.Author
	6, // 2: alexandria.catalogue.v1.Item.shelf:type_name -> alexandria.catalogue.v1.ShelfReference
	9, // 3: alexandria.catalogue.v1.Shelf.room:type_name -> alexandria.catalogue.v1.RoomReference
	4, // 4: alexandria.catalogue.v1.RoomInput.shelves:type_name -> alexandria.catalogue.v1.ShelfInput
	5, // 5: alexandria.catalogue.v1.Room.shelves:type_name -> alexandria.catalogue.v1.Shelf
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_alexandria_catalogue_v1_catalogue_proto_init() }
func file_alexandria_catalogue_v1_catalogue_proto_init() {
	if File_alexandria_catalogue_v1_catalogue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alexandria_catalogue_v1_catalogue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_alexandria_catalogue_v1_catalogue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemInput); i {
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
		file_alexandria_catalogue_v1_catalogue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_alexandria_catalogue_v1_catalogue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemReference); i {
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
		file_alexandria_catalogue_v1_catalogue_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShelfInput); i {
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
		file_alexandria_catalogue_v1_catalogue_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shelf); i {
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
		file_alexandria_catalogue_v1_catalogue_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShelfReference); i {
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
		file_alexandria_catalogue_v1_catalogue_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomInput); i {
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
		file_alexandria_catalogue_v1_catalogue_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_alexandria_catalogue_v1_catalogue_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomReference); i {
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
			RawDescriptor: file_alexandria_catalogue_v1_catalogue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alexandria_catalogue_v1_catalogue_proto_goTypes,
		DependencyIndexes: file_alexandria_catalogue_v1_catalogue_proto_depIdxs,
		MessageInfos:      file_alexandria_catalogue_v1_catalogue_proto_msgTypes,
	}.Build()
	File_alexandria_catalogue_v1_catalogue_proto = out.File
	file_alexandria_catalogue_v1_catalogue_proto_rawDesc = nil
	file_alexandria_catalogue_v1_catalogue_proto_goTypes = nil
	file_alexandria_catalogue_v1_catalogue_proto_depIdxs = nil
}
