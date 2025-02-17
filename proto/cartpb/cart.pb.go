// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/cartpb/cart.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CartItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     uint64                 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      uint32                 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	mi := &file_proto_cartpb_cart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartpb_cart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_proto_cartpb_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartItem) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartItem) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ProductDetail struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity      int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductDetail) Reset() {
	*x = ProductDetail{}
	mi := &file_proto_cartpb_cart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetail) ProtoMessage() {}

func (x *ProductDetail) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartpb_cart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDetail.ProtoReflect.Descriptor instead.
func (*ProductDetail) Descriptor() ([]byte, []int) {
	return file_proto_cartpb_cart_proto_rawDescGZIP(), []int{1}
}

func (x *ProductDetail) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ProductDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductDetail) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductDetail) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type RemoveAllItemsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // Changed to uint32 for consistency
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveAllItemsRequest) Reset() {
	*x = RemoveAllItemsRequest{}
	mi := &file_proto_cartpb_cart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveAllItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAllItemsRequest) ProtoMessage() {}

func (x *RemoveAllItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartpb_cart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAllItemsRequest.ProtoReflect.Descriptor instead.
func (*RemoveAllItemsRequest) Descriptor() ([]byte, []int) {
	return file_proto_cartpb_cart_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveAllItemsRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RemoveAllItemsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveAllItemsResponse) Reset() {
	*x = RemoveAllItemsResponse{}
	mi := &file_proto_cartpb_cart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveAllItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAllItemsResponse) ProtoMessage() {}

func (x *RemoveAllItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartpb_cart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAllItemsResponse.ProtoReflect.Descriptor instead.
func (*RemoveAllItemsResponse) Descriptor() ([]byte, []int) {
	return file_proto_cartpb_cart_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveAllItemsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName      string                 `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Products      []*ProductDetail       `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
	TotalPrice    float64                `protobuf:"fixed64,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartResponse) Reset() {
	*x = CartResponse{}
	mi := &file_proto_cartpb_cart_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartResponse) ProtoMessage() {}

func (x *CartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartpb_cart_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartResponse.ProtoReflect.Descriptor instead.
func (*CartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cartpb_cart_proto_rawDescGZIP(), []int{4}
}

func (x *CartResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CartResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *CartResponse) GetProducts() []*ProductDetail {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *CartResponse) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *CartResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type AddToCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*CartItem            `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	mi := &file_proto_cartpb_cart_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartpb_cart_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartRequest.ProtoReflect.Descriptor instead.
func (*AddToCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cartpb_cart_proto_rawDescGZIP(), []int{5}
}

func (x *AddToCartRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddToCartRequest) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type AddToCartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddToCartResponse) Reset() {
	*x = AddToCartResponse{}
	mi := &file_proto_cartpb_cart_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddToCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartResponse) ProtoMessage() {}

func (x *AddToCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartpb_cart_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartResponse.ProtoReflect.Descriptor instead.
func (*AddToCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cartpb_cart_proto_rawDescGZIP(), []int{6}
}

func (x *AddToCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCartRequest) Reset() {
	*x = GetCartRequest{}
	mi := &file_proto_cartpb_cart_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartRequest) ProtoMessage() {}

func (x *GetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartpb_cart_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartRequest.ProtoReflect.Descriptor instead.
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cartpb_cart_proto_rawDescGZIP(), []int{7}
}

func (x *GetCartRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RemoveFromCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveFromCartRequest) Reset() {
	*x = RemoveFromCartRequest{}
	mi := &file_proto_cartpb_cart_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromCartRequest) ProtoMessage() {}

func (x *RemoveFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartpb_cart_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromCartRequest.ProtoReflect.Descriptor instead.
func (*RemoveFromCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cartpb_cart_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveFromCartRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type RemoveFromCartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveFromCartResponse) Reset() {
	*x = RemoveFromCartResponse{}
	mi := &file_proto_cartpb_cart_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveFromCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromCartResponse) ProtoMessage() {}

func (x *RemoveFromCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cartpb_cart_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromCartResponse.ProtoReflect.Descriptor instead.
func (*RemoveFromCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cartpb_cart_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveFromCartResponse) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *RemoveFromCartResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_cartpb_cart_proto protoreflect.FileDescriptor

var file_proto_cartpb_cart_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2f, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x45, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x74, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x30, 0x0a,
	0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x32, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x52, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2d, 0x0a, 0x11, 0x41, 0x64, 0x64,
	0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x16, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xa2, 0x02, 0x0a,
	0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x09,
	0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x41, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_proto_cartpb_cart_proto_rawDescOnce sync.Once
	file_proto_cartpb_cart_proto_rawDescData []byte
)

func file_proto_cartpb_cart_proto_rawDescGZIP() []byte {
	file_proto_cartpb_cart_proto_rawDescOnce.Do(func() {
		file_proto_cartpb_cart_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_cartpb_cart_proto_rawDesc), len(file_proto_cartpb_cart_proto_rawDesc)))
	})
	return file_proto_cartpb_cart_proto_rawDescData
}

var file_proto_cartpb_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_cartpb_cart_proto_goTypes = []any{
	(*CartItem)(nil),               // 0: proto.CartItem
	(*ProductDetail)(nil),          // 1: proto.ProductDetail
	(*RemoveAllItemsRequest)(nil),  // 2: proto.RemoveAllItemsRequest
	(*RemoveAllItemsResponse)(nil), // 3: proto.RemoveAllItemsResponse
	(*CartResponse)(nil),           // 4: proto.CartResponse
	(*AddToCartRequest)(nil),       // 5: proto.AddToCartRequest
	(*AddToCartResponse)(nil),      // 6: proto.AddToCartResponse
	(*GetCartRequest)(nil),         // 7: proto.GetCartRequest
	(*RemoveFromCartRequest)(nil),  // 8: proto.RemoveFromCartRequest
	(*RemoveFromCartResponse)(nil), // 9: proto.RemoveFromCartResponse
}
var file_proto_cartpb_cart_proto_depIdxs = []int32{
	1, // 0: proto.CartResponse.products:type_name -> proto.ProductDetail
	0, // 1: proto.AddToCartRequest.items:type_name -> proto.CartItem
	5, // 2: proto.CartService.AddToCart:input_type -> proto.AddToCartRequest
	7, // 3: proto.CartService.GetCart:input_type -> proto.GetCartRequest
	8, // 4: proto.CartService.RemoveFromCart:input_type -> proto.RemoveFromCartRequest
	2, // 5: proto.CartService.RemoveAllItems:input_type -> proto.RemoveAllItemsRequest
	6, // 6: proto.CartService.AddToCart:output_type -> proto.AddToCartResponse
	4, // 7: proto.CartService.GetCart:output_type -> proto.CartResponse
	9, // 8: proto.CartService.RemoveFromCart:output_type -> proto.RemoveFromCartResponse
	3, // 9: proto.CartService.RemoveAllItems:output_type -> proto.RemoveAllItemsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_cartpb_cart_proto_init() }
func file_proto_cartpb_cart_proto_init() {
	if File_proto_cartpb_cart_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_cartpb_cart_proto_rawDesc), len(file_proto_cartpb_cart_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cartpb_cart_proto_goTypes,
		DependencyIndexes: file_proto_cartpb_cart_proto_depIdxs,
		MessageInfos:      file_proto_cartpb_cart_proto_msgTypes,
	}.Build()
	File_proto_cartpb_cart_proto = out.File
	file_proto_cartpb_cart_proto_goTypes = nil
	file_proto_cartpb_cart_proto_depIdxs = nil
}
