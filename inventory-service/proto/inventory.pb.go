// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: proto/inventory.proto

package inventorypb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type CreateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price         float64                `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Category      string                 `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Stock         int32                  `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	mi := &file_proto_inventory_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateProductRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type UpdateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Category      string                 `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Stock         int32                  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProductRequest) Reset() {
	*x = UpdateProductRequest{}
	mi := &file_proto_inventory_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRequest) ProtoMessage() {}

func (x *UpdateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateProductRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *UpdateProductRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type GetProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	mi := &file_proto_inventory_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductRequest) Reset() {
	*x = DeleteProductRequest{}
	mi := &file_proto_inventory_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductRequest) ProtoMessage() {}

func (x *DeleteProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Category      string                 `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Stock         int32                  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	mi := &file_proto_inventory_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *ProductResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ProductResponse) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type ListProductsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsRequest) Reset() {
	*x = ListProductsRequest{}
	mi := &file_proto_inventory_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsRequest) ProtoMessage() {}

func (x *ListProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsRequest.ProtoReflect.Descriptor instead.
func (*ListProductsRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{5}
}

type ListProductsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*ProductResponse     `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsResponse) Reset() {
	*x = ListProductsResponse{}
	mi := &file_proto_inventory_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResponse) ProtoMessage() {}

func (x *ListProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResponse.ProtoReflect.Descriptor instead.
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{6}
}

func (x *ListProductsResponse) GetProducts() []*ProductResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

type CreateCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCategoryRequest) Reset() {
	*x = CreateCategoryRequest{}
	mi := &file_proto_inventory_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryRequest) ProtoMessage() {}

func (x *CreateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{7}
}

func (x *CreateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCategoryRequest) Reset() {
	*x = UpdateCategoryRequest{}
	mi := &file_proto_inventory_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCategoryRequest) ProtoMessage() {}

func (x *UpdateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCategoryRequest) Reset() {
	*x = GetCategoryRequest{}
	mi := &file_proto_inventory_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryRequest) ProtoMessage() {}

func (x *GetCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{9}
}

func (x *GetCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCategoryRequest) Reset() {
	*x = DeleteCategoryRequest{}
	mi := &file_proto_inventory_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryRequest) ProtoMessage() {}

func (x *DeleteCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CategoryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryResponse) Reset() {
	*x = CategoryResponse{}
	mi := &file_proto_inventory_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryResponse) ProtoMessage() {}

func (x *CategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryResponse.ProtoReflect.Descriptor instead.
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{11}
}

func (x *CategoryResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CategoryResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListCategoriesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCategoriesRequest) Reset() {
	*x = ListCategoriesRequest{}
	mi := &file_proto_inventory_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoriesRequest) ProtoMessage() {}

func (x *ListCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoriesRequest.ProtoReflect.Descriptor instead.
func (*ListCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{12}
}

type ListCategoriesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Categories    []*CategoryResponse    `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCategoriesResponse) Reset() {
	*x = ListCategoriesResponse{}
	mi := &file_proto_inventory_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoriesResponse) ProtoMessage() {}

func (x *ListCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoriesResponse.ProtoReflect.Descriptor instead.
func (*ListCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{13}
}

func (x *ListCategoriesResponse) GetCategories() []*CategoryResponse {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_proto_inventory_proto protoreflect.FileDescriptor

const file_proto_inventory_proto_rawDesc = "" +
	"\n" +
	"\x15proto/inventory.proto\x12\tinventory\x1a\x1bgoogle/protobuf/empty.proto\"r\n" +
	"\x14CreateProductRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05price\x18\x02 \x01(\x01R\x05price\x12\x1a\n" +
	"\bcategory\x18\x03 \x01(\tR\bcategory\x12\x14\n" +
	"\x05stock\x18\x04 \x01(\x05R\x05stock\"\x82\x01\n" +
	"\x14UpdateProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x01R\x05price\x12\x1a\n" +
	"\bcategory\x18\x04 \x01(\tR\bcategory\x12\x14\n" +
	"\x05stock\x18\x05 \x01(\x05R\x05stock\"#\n" +
	"\x11GetProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"&\n" +
	"\x14DeleteProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"}\n" +
	"\x0fProductResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x01R\x05price\x12\x1a\n" +
	"\bcategory\x18\x04 \x01(\tR\bcategory\x12\x14\n" +
	"\x05stock\x18\x05 \x01(\x05R\x05stock\"\x15\n" +
	"\x13ListProductsRequest\"N\n" +
	"\x14ListProductsResponse\x126\n" +
	"\bproducts\x18\x01 \x03(\v2\x1a.inventory.ProductResponseR\bproducts\"+\n" +
	"\x15CreateCategoryRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\";\n" +
	"\x15UpdateCategoryRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\"$\n" +
	"\x12GetCategoryRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"'\n" +
	"\x15DeleteCategoryRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"6\n" +
	"\x10CategoryResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\"\x17\n" +
	"\x15ListCategoriesRequest\"U\n" +
	"\x16ListCategoriesResponse\x12;\n" +
	"\n" +
	"categories\x18\x01 \x03(\v2\x1b.inventory.CategoryResponseR\n" +
	"categories2\xa9\x06\n" +
	"\x10InventoryService\x12L\n" +
	"\rCreateProduct\x12\x1f.inventory.CreateProductRequest\x1a\x1a.inventory.ProductResponse\x12J\n" +
	"\x0eGetProductByID\x12\x1c.inventory.GetProductRequest\x1a\x1a.inventory.ProductResponse\x12L\n" +
	"\rUpdateProduct\x12\x1f.inventory.UpdateProductRequest\x1a\x1a.inventory.ProductResponse\x12H\n" +
	"\rDeleteProduct\x12\x1f.inventory.DeleteProductRequest\x1a\x16.google.protobuf.Empty\x12O\n" +
	"\fListProducts\x12\x1e.inventory.ListProductsRequest\x1a\x1f.inventory.ListProductsResponse\x12O\n" +
	"\x0eCreateCategory\x12 .inventory.CreateCategoryRequest\x1a\x1b.inventory.CategoryResponse\x12M\n" +
	"\x0fGetCategoryByID\x12\x1d.inventory.GetCategoryRequest\x1a\x1b.inventory.CategoryResponse\x12O\n" +
	"\x0eUpdateCategory\x12 .inventory.UpdateCategoryRequest\x1a\x1b.inventory.CategoryResponse\x12J\n" +
	"\x0eDeleteCategory\x12 .inventory.DeleteCategoryRequest\x1a\x16.google.protobuf.Empty\x12U\n" +
	"\x0eListCategories\x12 .inventory.ListCategoriesRequest\x1a!.inventory.ListCategoriesResponseBMZKgithub.com/Neroframe/ecommerce-platform/inventory-service/proto;inventorypbb\x06proto3"

var (
	file_proto_inventory_proto_rawDescOnce sync.Once
	file_proto_inventory_proto_rawDescData []byte
)

func file_proto_inventory_proto_rawDescGZIP() []byte {
	file_proto_inventory_proto_rawDescOnce.Do(func() {
		file_proto_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_inventory_proto_rawDesc), len(file_proto_inventory_proto_rawDesc)))
	})
	return file_proto_inventory_proto_rawDescData
}

var file_proto_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_inventory_proto_goTypes = []any{
	(*CreateProductRequest)(nil),   // 0: inventory.CreateProductRequest
	(*UpdateProductRequest)(nil),   // 1: inventory.UpdateProductRequest
	(*GetProductRequest)(nil),      // 2: inventory.GetProductRequest
	(*DeleteProductRequest)(nil),   // 3: inventory.DeleteProductRequest
	(*ProductResponse)(nil),        // 4: inventory.ProductResponse
	(*ListProductsRequest)(nil),    // 5: inventory.ListProductsRequest
	(*ListProductsResponse)(nil),   // 6: inventory.ListProductsResponse
	(*CreateCategoryRequest)(nil),  // 7: inventory.CreateCategoryRequest
	(*UpdateCategoryRequest)(nil),  // 8: inventory.UpdateCategoryRequest
	(*GetCategoryRequest)(nil),     // 9: inventory.GetCategoryRequest
	(*DeleteCategoryRequest)(nil),  // 10: inventory.DeleteCategoryRequest
	(*CategoryResponse)(nil),       // 11: inventory.CategoryResponse
	(*ListCategoriesRequest)(nil),  // 12: inventory.ListCategoriesRequest
	(*ListCategoriesResponse)(nil), // 13: inventory.ListCategoriesResponse
	(*empty.Empty)(nil),            // 14: google.protobuf.Empty
}
var file_proto_inventory_proto_depIdxs = []int32{
	4,  // 0: inventory.ListProductsResponse.products:type_name -> inventory.ProductResponse
	11, // 1: inventory.ListCategoriesResponse.categories:type_name -> inventory.CategoryResponse
	0,  // 2: inventory.InventoryService.CreateProduct:input_type -> inventory.CreateProductRequest
	2,  // 3: inventory.InventoryService.GetProductByID:input_type -> inventory.GetProductRequest
	1,  // 4: inventory.InventoryService.UpdateProduct:input_type -> inventory.UpdateProductRequest
	3,  // 5: inventory.InventoryService.DeleteProduct:input_type -> inventory.DeleteProductRequest
	5,  // 6: inventory.InventoryService.ListProducts:input_type -> inventory.ListProductsRequest
	7,  // 7: inventory.InventoryService.CreateCategory:input_type -> inventory.CreateCategoryRequest
	9,  // 8: inventory.InventoryService.GetCategoryByID:input_type -> inventory.GetCategoryRequest
	8,  // 9: inventory.InventoryService.UpdateCategory:input_type -> inventory.UpdateCategoryRequest
	10, // 10: inventory.InventoryService.DeleteCategory:input_type -> inventory.DeleteCategoryRequest
	12, // 11: inventory.InventoryService.ListCategories:input_type -> inventory.ListCategoriesRequest
	4,  // 12: inventory.InventoryService.CreateProduct:output_type -> inventory.ProductResponse
	4,  // 13: inventory.InventoryService.GetProductByID:output_type -> inventory.ProductResponse
	4,  // 14: inventory.InventoryService.UpdateProduct:output_type -> inventory.ProductResponse
	14, // 15: inventory.InventoryService.DeleteProduct:output_type -> google.protobuf.Empty
	6,  // 16: inventory.InventoryService.ListProducts:output_type -> inventory.ListProductsResponse
	11, // 17: inventory.InventoryService.CreateCategory:output_type -> inventory.CategoryResponse
	11, // 18: inventory.InventoryService.GetCategoryByID:output_type -> inventory.CategoryResponse
	11, // 19: inventory.InventoryService.UpdateCategory:output_type -> inventory.CategoryResponse
	14, // 20: inventory.InventoryService.DeleteCategory:output_type -> google.protobuf.Empty
	13, // 21: inventory.InventoryService.ListCategories:output_type -> inventory.ListCategoriesResponse
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_inventory_proto_init() }
func file_proto_inventory_proto_init() {
	if File_proto_inventory_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_inventory_proto_rawDesc), len(file_proto_inventory_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_inventory_proto_goTypes,
		DependencyIndexes: file_proto_inventory_proto_depIdxs,
		MessageInfos:      file_proto_inventory_proto_msgTypes,
	}.Build()
	File_proto_inventory_proto = out.File
	file_proto_inventory_proto_goTypes = nil
	file_proto_inventory_proto_depIdxs = nil
}
