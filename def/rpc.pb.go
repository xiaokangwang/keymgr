package def

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

type GetPrimaryKeyHashReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPrimaryKeyHashReq) Reset() {
	*x = GetPrimaryKeyHashReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrimaryKeyHashReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrimaryKeyHashReq) ProtoMessage() {}

func (x *GetPrimaryKeyHashReq) ProtoReflect() protoreflect.Message {
	mi := &file_def_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrimaryKeyHashReq.ProtoReflect.Descriptor instead.
func (*GetPrimaryKeyHashReq) Descriptor() ([]byte, []int) {
	return file_def_rpc_proto_rawDescGZIP(), []int{0}
}

type GetPrimaryKeyHashResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimaryKey []byte `protobuf:"bytes,1,opt,name=primaryKey,proto3" json:"primaryKey,omitempty"`
}

func (x *GetPrimaryKeyHashResp) Reset() {
	*x = GetPrimaryKeyHashResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrimaryKeyHashResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrimaryKeyHashResp) ProtoMessage() {}

func (x *GetPrimaryKeyHashResp) ProtoReflect() protoreflect.Message {
	mi := &file_def_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrimaryKeyHashResp.ProtoReflect.Descriptor instead.
func (*GetPrimaryKeyHashResp) Descriptor() ([]byte, []int) {
	return file_def_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetPrimaryKeyHashResp) GetPrimaryKey() []byte {
	if x != nil {
		return x.PrimaryKey
	}
	return nil
}

type CreateECDSAKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeySeed string `protobuf:"bytes,1,opt,name=keySeed,proto3" json:"keySeed,omitempty"`
}

func (x *CreateECDSAKeyReq) Reset() {
	*x = CreateECDSAKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateECDSAKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateECDSAKeyReq) ProtoMessage() {}

func (x *CreateECDSAKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_def_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateECDSAKeyReq.ProtoReflect.Descriptor instead.
func (*CreateECDSAKeyReq) Descriptor() ([]byte, []int) {
	return file_def_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *CreateECDSAKeyReq) GetKeySeed() string {
	if x != nil {
		return x.KeySeed
	}
	return ""
}

type CreateECDSAKeyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicBlob   []byte `protobuf:"bytes,2,opt,name=publicBlob,proto3" json:"publicBlob,omitempty"`
	SshPublicKey string `protobuf:"bytes,3,opt,name=ssh_public_key,json=sshPublicKey,proto3" json:"ssh_public_key,omitempty"`
}

func (x *CreateECDSAKeyResp) Reset() {
	*x = CreateECDSAKeyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateECDSAKeyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateECDSAKeyResp) ProtoMessage() {}

func (x *CreateECDSAKeyResp) ProtoReflect() protoreflect.Message {
	mi := &file_def_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateECDSAKeyResp.ProtoReflect.Descriptor instead.
func (*CreateECDSAKeyResp) Descriptor() ([]byte, []int) {
	return file_def_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *CreateECDSAKeyResp) GetPublicBlob() []byte {
	if x != nil {
		return x.PublicBlob
	}
	return nil
}

func (x *CreateECDSAKeyResp) GetSshPublicKey() string {
	if x != nil {
		return x.SshPublicKey
	}
	return ""
}

type SignWithECDSAKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeySeed string `protobuf:"bytes,1,opt,name=keySeed,proto3" json:"keySeed,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SignWithECDSAKeyReq) Reset() {
	*x = SignWithECDSAKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignWithECDSAKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignWithECDSAKeyReq) ProtoMessage() {}

func (x *SignWithECDSAKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_def_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignWithECDSAKeyReq.ProtoReflect.Descriptor instead.
func (*SignWithECDSAKeyReq) Descriptor() ([]byte, []int) {
	return file_def_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *SignWithECDSAKeyReq) GetKeySeed() string {
	if x != nil {
		return x.KeySeed
	}
	return ""
}

func (x *SignWithECDSAKeyReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SignWithECDSAKeyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignWithECDSAKeyResp) Reset() {
	*x = SignWithECDSAKeyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignWithECDSAKeyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignWithECDSAKeyResp) ProtoMessage() {}

func (x *SignWithECDSAKeyResp) ProtoReflect() protoreflect.Message {
	mi := &file_def_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignWithECDSAKeyResp.ProtoReflect.Descriptor instead.
func (*SignWithECDSAKeyResp) Descriptor() ([]byte, []int) {
	return file_def_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *SignWithECDSAKeyResp) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_def_rpc_proto protoreflect.FileDescriptor

var file_def_rpc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x66, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x6f, 0x72, 0x67, 0x2e, 0x6b, 0x6b, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x65, 0x79, 0x6d, 0x67,
	0x72, 0x2e, 0x64, 0x65, 0x66, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x22, 0x37, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x48, 0x61,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x2d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x43, 0x44, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6b,
	0x65, 0x79, 0x53, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65,
	0x79, 0x53, 0x65, 0x65, 0x64, 0x22, 0x5a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x43, 0x44, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x42, 0x6c, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x24, 0x0a, 0x0e, 0x73,
	0x73, 0x68, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x22, 0x43, 0x0a, 0x13, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x45, 0x43, 0x44,
	0x53, 0x41, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x53,
	0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x53, 0x65,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x34, 0x0a, 0x14, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x69,
	0x74, 0x68, 0x45, 0x43, 0x44, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0xcd, 0x02, 0x0a,
	0x0d, 0x54, 0x50, 0x4d, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x2a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6b, 0x6b, 0x64, 0x65, 0x76, 0x2e,
	0x6b, 0x65, 0x79, 0x6d, 0x67, 0x72, 0x2e, 0x64, 0x65, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a,
	0x2b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6b, 0x6b, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x65, 0x79, 0x6d,
	0x67, 0x72, 0x2e, 0x64, 0x65, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x63, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x43, 0x44, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x12, 0x27,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6b, 0x6b, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x65, 0x79, 0x6d, 0x67,
	0x72, 0x2e, 0x64, 0x65, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x43, 0x44, 0x53,
	0x41, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6b, 0x6b,
	0x64, 0x65, 0x76, 0x2e, 0x6b, 0x65, 0x79, 0x6d, 0x67, 0x72, 0x2e, 0x64, 0x65, 0x66, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x43, 0x44, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x69, 0x0a, 0x10, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x45, 0x43, 0x44,
	0x53, 0x41, 0x4b, 0x65, 0x79, 0x12, 0x29, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6b, 0x6b, 0x64, 0x65,
	0x76, 0x2e, 0x6b, 0x65, 0x79, 0x6d, 0x67, 0x72, 0x2e, 0x64, 0x65, 0x66, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x57, 0x69, 0x74, 0x68, 0x45, 0x43, 0x44, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x1a, 0x2a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6b, 0x6b, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x65, 0x79,
	0x6d, 0x67, 0x72, 0x2e, 0x64, 0x65, 0x66, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x69, 0x74, 0x68,
	0x45, 0x43, 0x44, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x42, 0x24, 0x5a, 0x22,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61, 0x6f, 0x6b,
	0x61, 0x6e, 0x67, 0x77, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x6d, 0x67, 0x72, 0x2f, 0x64,
	0x65, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_def_rpc_proto_rawDescOnce sync.Once
	file_def_rpc_proto_rawDescData = file_def_rpc_proto_rawDesc
)

func file_def_rpc_proto_rawDescGZIP() []byte {
	file_def_rpc_proto_rawDescOnce.Do(func() {
		file_def_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_def_rpc_proto_rawDescData)
	})
	return file_def_rpc_proto_rawDescData
}

var file_def_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_def_rpc_proto_goTypes = []interface{}{
	(*GetPrimaryKeyHashReq)(nil),  // 0: org.kkdev.keymgr.def.GetPrimaryKeyHashReq
	(*GetPrimaryKeyHashResp)(nil), // 1: org.kkdev.keymgr.def.GetPrimaryKeyHashResp
	(*CreateECDSAKeyReq)(nil),     // 2: org.kkdev.keymgr.def.CreateECDSAKeyReq
	(*CreateECDSAKeyResp)(nil),    // 3: org.kkdev.keymgr.def.CreateECDSAKeyResp
	(*SignWithECDSAKeyReq)(nil),   // 4: org.kkdev.keymgr.def.SignWithECDSAKeyReq
	(*SignWithECDSAKeyResp)(nil),  // 5: org.kkdev.keymgr.def.SignWithECDSAKeyResp
}
var file_def_rpc_proto_depIdxs = []int32{
	0, // 0: org.kkdev.keymgr.def.TPMKeyService.GetPrimaryKeyHash:input_type -> org.kkdev.keymgr.def.GetPrimaryKeyHashReq
	2, // 1: org.kkdev.keymgr.def.TPMKeyService.CreateECDSAKey:input_type -> org.kkdev.keymgr.def.CreateECDSAKeyReq
	4, // 2: org.kkdev.keymgr.def.TPMKeyService.SignWithECDSAKey:input_type -> org.kkdev.keymgr.def.SignWithECDSAKeyReq
	1, // 3: org.kkdev.keymgr.def.TPMKeyService.GetPrimaryKeyHash:output_type -> org.kkdev.keymgr.def.GetPrimaryKeyHashResp
	3, // 4: org.kkdev.keymgr.def.TPMKeyService.CreateECDSAKey:output_type -> org.kkdev.keymgr.def.CreateECDSAKeyResp
	5, // 5: org.kkdev.keymgr.def.TPMKeyService.SignWithECDSAKey:output_type -> org.kkdev.keymgr.def.SignWithECDSAKeyResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_def_rpc_proto_init() }
func file_def_rpc_proto_init() {
	if File_def_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_def_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrimaryKeyHashReq); i {
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
		file_def_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrimaryKeyHashResp); i {
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
		file_def_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateECDSAKeyReq); i {
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
		file_def_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateECDSAKeyResp); i {
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
		file_def_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignWithECDSAKeyReq); i {
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
		file_def_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignWithECDSAKeyResp); i {
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
			RawDescriptor: file_def_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_def_rpc_proto_goTypes,
		DependencyIndexes: file_def_rpc_proto_depIdxs,
		MessageInfos:      file_def_rpc_proto_msgTypes,
	}.Build()
	File_def_rpc_proto = out.File
	file_def_rpc_proto_rawDesc = nil
	file_def_rpc_proto_goTypes = nil
	file_def_rpc_proto_depIdxs = nil
}
