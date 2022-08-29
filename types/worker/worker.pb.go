// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: protos/worker.proto

package worker

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

type LifeCircleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId   string `protobuf:"bytes,1,opt,name=workerId,proto3" json:"workerId,omitempty"`
	Register   bool   `protobuf:"varint,2,opt,name=register,proto3" json:"register,omitempty"`
	Alivecheck bool   `protobuf:"varint,3,opt,name=alivecheck,proto3" json:"alivecheck,omitempty"`
}

func (x *LifeCircleReq) Reset() {
	*x = LifeCircleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LifeCircleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifeCircleReq) ProtoMessage() {}

func (x *LifeCircleReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifeCircleReq.ProtoReflect.Descriptor instead.
func (*LifeCircleReq) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{0}
}

func (x *LifeCircleReq) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *LifeCircleReq) GetRegister() bool {
	if x != nil {
		return x.Register
	}
	return false
}

func (x *LifeCircleReq) GetAlivecheck() bool {
	if x != nil {
		return x.Alivecheck
	}
	return false
}

type LifeCircleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Alive  bool   `protobuf:"varint,2,opt,name=alive,proto3" json:"alive,omitempty"`
}

func (x *LifeCircleResp) Reset() {
	*x = LifeCircleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LifeCircleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifeCircleResp) ProtoMessage() {}

func (x *LifeCircleResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifeCircleResp.ProtoReflect.Descriptor instead.
func (*LifeCircleResp) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{1}
}

func (x *LifeCircleResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LifeCircleResp) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

type StatusSyncReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fetch bool `protobuf:"varint,1,opt,name=fetch,proto3" json:"fetch,omitempty"`
}

func (x *StatusSyncReq) Reset() {
	*x = StatusSyncReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusSyncReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusSyncReq) ProtoMessage() {}

func (x *StatusSyncReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusSyncReq.ProtoReflect.Descriptor instead.
func (*StatusSyncReq) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{2}
}

func (x *StatusSyncReq) GetFetch() bool {
	if x != nil {
		return x.Fetch
	}
	return false
}

type StatusSyncResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusSyncResp) Reset() {
	*x = StatusSyncResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusSyncResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusSyncResp) ProtoMessage() {}

func (x *StatusSyncResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusSyncResp.ProtoReflect.Descriptor instead.
func (*StatusSyncResp) Descriptor() ([]byte, []int) {
	return file_protos_worker_proto_rawDescGZIP(), []int{3}
}

func (x *StatusSyncResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_protos_worker_proto protoreflect.FileDescriptor

var file_protos_worker_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x22, 0x67, 0x0a,
	0x0d, 0x4c, 0x69, 0x66, 0x65, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c, 0x69, 0x76,
	0x65, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x3e, 0x0a, 0x0e, 0x4c, 0x69, 0x66, 0x65, 0x43, 0x69,
	0x72, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x65, 0x74, 0x63, 0x68, 0x22, 0x28, 0x0a,
	0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x8a, 0x01, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0a, 0x4c, 0x69, 0x66, 0x65, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65,
	0x12, 0x15, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x43, 0x69,
	0x72, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x66, 0x65, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x79, 0x6e,
	0x63, 0x12, 0x15, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_worker_proto_rawDescOnce sync.Once
	file_protos_worker_proto_rawDescData = file_protos_worker_proto_rawDesc
)

func file_protos_worker_proto_rawDescGZIP() []byte {
	file_protos_worker_proto_rawDescOnce.Do(func() {
		file_protos_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_worker_proto_rawDescData)
	})
	return file_protos_worker_proto_rawDescData
}

var file_protos_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_worker_proto_goTypes = []interface{}{
	(*LifeCircleReq)(nil),  // 0: worker.LifeCircleReq
	(*LifeCircleResp)(nil), // 1: worker.LifeCircleResp
	(*StatusSyncReq)(nil),  // 2: worker.StatusSyncReq
	(*StatusSyncResp)(nil), // 3: worker.StatusSyncResp
}
var file_protos_worker_proto_depIdxs = []int32{
	0, // 0: worker.Worker.LifeCircle:input_type -> worker.LifeCircleReq
	2, // 1: worker.Worker.StatusSync:input_type -> worker.StatusSyncReq
	1, // 2: worker.Worker.LifeCircle:output_type -> worker.LifeCircleResp
	3, // 3: worker.Worker.StatusSync:output_type -> worker.StatusSyncResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_worker_proto_init() }
func file_protos_worker_proto_init() {
	if File_protos_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LifeCircleReq); i {
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
		file_protos_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LifeCircleResp); i {
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
		file_protos_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusSyncReq); i {
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
		file_protos_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusSyncResp); i {
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
			RawDescriptor: file_protos_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_worker_proto_goTypes,
		DependencyIndexes: file_protos_worker_proto_depIdxs,
		MessageInfos:      file_protos_worker_proto_msgTypes,
	}.Build()
	File_protos_worker_proto = out.File
	file_protos_worker_proto_rawDesc = nil
	file_protos_worker_proto_goTypes = nil
	file_protos_worker_proto_depIdxs = nil
}
