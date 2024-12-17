// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: relay/relay.proto

package relay

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

// A request to fetch one or more blobs.
type GetBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key of the blob to fetch.
	BlobKey []byte `protobuf:"bytes,1,opt,name=blob_key,json=blobKey,proto3" json:"blob_key,omitempty"`
}

func (x *GetBlobRequest) Reset() {
	*x = GetBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlobRequest) ProtoMessage() {}

func (x *GetBlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlobRequest.ProtoReflect.Descriptor instead.
func (*GetBlobRequest) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{0}
}

func (x *GetBlobRequest) GetBlobKey() []byte {
	if x != nil {
		return x.BlobKey
	}
	return nil
}

// The reply to a GetBlobs request.
type GetBlobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The blob requested.
	Blob []byte `protobuf:"bytes,1,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (x *GetBlobReply) Reset() {
	*x = GetBlobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlobReply) ProtoMessage() {}

func (x *GetBlobReply) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlobReply.ProtoReflect.Descriptor instead.
func (*GetBlobReply) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{1}
}

func (x *GetBlobReply) GetBlob() []byte {
	if x != nil {
		return x.Blob
	}
	return nil
}

// Request chunks from blobs stored by this relay.
type GetChunksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The chunk requests. Chunks are returned in the same order as they are requested.
	ChunkRequests []*ChunkRequest `protobuf:"bytes,1,rep,name=chunk_requests,json=chunkRequests,proto3" json:"chunk_requests,omitempty"`
	// If this is an authenticated request, this should hold the ID of the operator. If this
	// is an unauthenticated request, this field should be empty. Relays may choose to reject
	// unauthenticated requests.
	OperatorId []byte `protobuf:"bytes,2,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	// If this is an authenticated request, this field will hold a BLS signature by the requester
	// on the hash of this request. Relays may choose to reject unauthenticated requests.
	//
	// The following describes the schema for computing the hash of this request
	// This algorithm is implemented in golang using relay.auth.HashGetChunksRequest().
	//
	// All integers are encoded as unsigned 4 byte big endian values.
	//
	// Perform a keccak256 hash on the following data in the following order:
	//  1. the operator id
	//  2. for each chunk request:
	//     a. if the chunk request is a request by index:
	//     i.   a one byte ASCII representation of the character "i" (aka Ox69)
	//     ii.  the blob key
	//     iii. the start index
	//     iv.  the end index
	//     b. if the chunk request is a request by range:
	//     i.   a one byte ASCII representation of the character "r" (aka Ox72)
	//     ii.  the blob key
	//     iii. each requested chunk index, in order
	OperatorSignature []byte `protobuf:"bytes,3,opt,name=operator_signature,json=operatorSignature,proto3" json:"operator_signature,omitempty"`
}

func (x *GetChunksRequest) Reset() {
	*x = GetChunksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunksRequest) ProtoMessage() {}

func (x *GetChunksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunksRequest.ProtoReflect.Descriptor instead.
func (*GetChunksRequest) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{2}
}

func (x *GetChunksRequest) GetChunkRequests() []*ChunkRequest {
	if x != nil {
		return x.ChunkRequests
	}
	return nil
}

func (x *GetChunksRequest) GetOperatorId() []byte {
	if x != nil {
		return x.OperatorId
	}
	return nil
}

func (x *GetChunksRequest) GetOperatorSignature() []byte {
	if x != nil {
		return x.OperatorSignature
	}
	return nil
}

// A request for chunks within a specific blob. Each chunk is requested individually by its index.
type ChunkRequestByIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The blob key.
	BlobKey []byte `protobuf:"bytes,1,opt,name=blob_key,json=blobKey,proto3" json:"blob_key,omitempty"`
	// The index of the chunk within the blob.
	ChunkIndices []uint32 `protobuf:"varint,2,rep,packed,name=chunk_indices,json=chunkIndices,proto3" json:"chunk_indices,omitempty"`
}

func (x *ChunkRequestByIndex) Reset() {
	*x = ChunkRequestByIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkRequestByIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkRequestByIndex) ProtoMessage() {}

func (x *ChunkRequestByIndex) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkRequestByIndex.ProtoReflect.Descriptor instead.
func (*ChunkRequestByIndex) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{3}
}

func (x *ChunkRequestByIndex) GetBlobKey() []byte {
	if x != nil {
		return x.BlobKey
	}
	return nil
}

func (x *ChunkRequestByIndex) GetChunkIndices() []uint32 {
	if x != nil {
		return x.ChunkIndices
	}
	return nil
}

// A request for chunks within a specific blob. Each chunk is requested a range of indices.
type ChunkRequestByRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The blob key.
	BlobKey []byte `protobuf:"bytes,1,opt,name=blob_key,json=blobKey,proto3" json:"blob_key,omitempty"`
	// The first index to start fetching chunks from.
	StartIndex uint32 `protobuf:"varint,2,opt,name=start_index,json=startIndex,proto3" json:"start_index,omitempty"`
	// One past the last index to fetch chunks from. Similar semantics to golang slices.
	EndIndex uint32 `protobuf:"varint,3,opt,name=end_index,json=endIndex,proto3" json:"end_index,omitempty"`
}

func (x *ChunkRequestByRange) Reset() {
	*x = ChunkRequestByRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkRequestByRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkRequestByRange) ProtoMessage() {}

func (x *ChunkRequestByRange) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkRequestByRange.ProtoReflect.Descriptor instead.
func (*ChunkRequestByRange) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{4}
}

func (x *ChunkRequestByRange) GetBlobKey() []byte {
	if x != nil {
		return x.BlobKey
	}
	return nil
}

func (x *ChunkRequestByRange) GetStartIndex() uint32 {
	if x != nil {
		return x.StartIndex
	}
	return 0
}

func (x *ChunkRequestByRange) GetEndIndex() uint32 {
	if x != nil {
		return x.EndIndex
	}
	return 0
}

// A request for chunks within a specific blob. Requests are fulfilled in all-or-nothing fashion. If any of the
// requested chunks are not found or are unable to be fetched, the entire request will fail.
type ChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*ChunkRequest_ByIndex
	//	*ChunkRequest_ByRange
	Request isChunkRequest_Request `protobuf_oneof:"request"`
}

func (x *ChunkRequest) Reset() {
	*x = ChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkRequest) ProtoMessage() {}

func (x *ChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkRequest.ProtoReflect.Descriptor instead.
func (*ChunkRequest) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{5}
}

func (m *ChunkRequest) GetRequest() isChunkRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *ChunkRequest) GetByIndex() *ChunkRequestByIndex {
	if x, ok := x.GetRequest().(*ChunkRequest_ByIndex); ok {
		return x.ByIndex
	}
	return nil
}

func (x *ChunkRequest) GetByRange() *ChunkRequestByRange {
	if x, ok := x.GetRequest().(*ChunkRequest_ByRange); ok {
		return x.ByRange
	}
	return nil
}

type isChunkRequest_Request interface {
	isChunkRequest_Request()
}

type ChunkRequest_ByIndex struct {
	// Request chunks by their individual indices.
	ByIndex *ChunkRequestByIndex `protobuf:"bytes,1,opt,name=by_index,json=byIndex,proto3,oneof"`
}

type ChunkRequest_ByRange struct {
	// Request chunks by a range of indices.
	ByRange *ChunkRequestByRange `protobuf:"bytes,2,opt,name=by_range,json=byRange,proto3,oneof"`
}

func (*ChunkRequest_ByIndex) isChunkRequest_Request() {}

func (*ChunkRequest_ByRange) isChunkRequest_Request() {}

// The reply to a GetChunks request.
type GetChunksReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The chunks requested. The order of these chunks will be the same as the order of the requested chunks.
	// data is the raw data of the bundle (i.e. serialized byte array of the frames)
	Data [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetChunksReply) Reset() {
	*x = GetChunksReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunksReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunksReply) ProtoMessage() {}

func (x *GetChunksReply) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunksReply.ProtoReflect.Descriptor instead.
func (*GetChunksReply) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{6}
}

func (x *GetChunksReply) GetData() [][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_relay_relay_proto protoreflect.FileDescriptor

var file_relay_relay_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x22, 0x2b, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x62, 0x6c, 0x6f, 0x62, 0x4b, 0x65, 0x79, 0x22, 0x22, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x22, 0x9e, 0x01, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3a, 0x0a, 0x0e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a,
	0x12, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x55, 0x0a, 0x13,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x69,
	0x63, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x13, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c,
	0x6f, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c,
	0x6f, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x22, 0x8b, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x62, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x48, 0x00, 0x52, 0x07, 0x62, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x37, 0x0a,
	0x08, 0x62, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x62,
	0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x7f, 0x0a, 0x05, 0x52, 0x65, 0x6c, 0x61, 0x79,
	0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x15, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x79, 0x72, 0x2d, 0x4c, 0x61, 0x62, 0x73,
	0x2f, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relay_relay_proto_rawDescOnce sync.Once
	file_relay_relay_proto_rawDescData = file_relay_relay_proto_rawDesc
)

func file_relay_relay_proto_rawDescGZIP() []byte {
	file_relay_relay_proto_rawDescOnce.Do(func() {
		file_relay_relay_proto_rawDescData = protoimpl.X.CompressGZIP(file_relay_relay_proto_rawDescData)
	})
	return file_relay_relay_proto_rawDescData
}

var file_relay_relay_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_relay_relay_proto_goTypes = []interface{}{
	(*GetBlobRequest)(nil),      // 0: relay.GetBlobRequest
	(*GetBlobReply)(nil),        // 1: relay.GetBlobReply
	(*GetChunksRequest)(nil),    // 2: relay.GetChunksRequest
	(*ChunkRequestByIndex)(nil), // 3: relay.ChunkRequestByIndex
	(*ChunkRequestByRange)(nil), // 4: relay.ChunkRequestByRange
	(*ChunkRequest)(nil),        // 5: relay.ChunkRequest
	(*GetChunksReply)(nil),      // 6: relay.GetChunksReply
}
var file_relay_relay_proto_depIdxs = []int32{
	5, // 0: relay.GetChunksRequest.chunk_requests:type_name -> relay.ChunkRequest
	3, // 1: relay.ChunkRequest.by_index:type_name -> relay.ChunkRequestByIndex
	4, // 2: relay.ChunkRequest.by_range:type_name -> relay.ChunkRequestByRange
	0, // 3: relay.Relay.GetBlob:input_type -> relay.GetBlobRequest
	2, // 4: relay.Relay.GetChunks:input_type -> relay.GetChunksRequest
	1, // 5: relay.Relay.GetBlob:output_type -> relay.GetBlobReply
	6, // 6: relay.Relay.GetChunks:output_type -> relay.GetChunksReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_relay_relay_proto_init() }
func file_relay_relay_proto_init() {
	if File_relay_relay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relay_relay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlobRequest); i {
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
		file_relay_relay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlobReply); i {
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
		file_relay_relay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChunksRequest); i {
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
		file_relay_relay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkRequestByIndex); i {
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
		file_relay_relay_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkRequestByRange); i {
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
		file_relay_relay_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkRequest); i {
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
		file_relay_relay_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChunksReply); i {
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
	file_relay_relay_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ChunkRequest_ByIndex)(nil),
		(*ChunkRequest_ByRange)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relay_relay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relay_relay_proto_goTypes,
		DependencyIndexes: file_relay_relay_proto_depIdxs,
		MessageInfos:      file_relay_relay_proto_msgTypes,
	}.Build()
	File_relay_relay_proto = out.File
	file_relay_relay_proto_rawDesc = nil
	file_relay_relay_proto_goTypes = nil
	file_relay_relay_proto_depIdxs = nil
}
