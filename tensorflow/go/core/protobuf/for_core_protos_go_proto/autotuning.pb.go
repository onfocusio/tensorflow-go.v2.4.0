// This file defines protos that store the results of autotuning various
// operations.
//
// They are in proto format because we want to log them structured. They offer
// tremendous statistical, testing, and debugging value.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.14.0
// source: tensorflow/core/protobuf/autotuning.proto

package for_core_protos_go_proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AutotuneResult_FailureKind int32

const (
	AutotuneResult_UNKNOWN          AutotuneResult_FailureKind = 0
	AutotuneResult_REDZONE_MODIFIED AutotuneResult_FailureKind = 1
	AutotuneResult_WRONG_RESULT     AutotuneResult_FailureKind = 2
)

// Enum value maps for AutotuneResult_FailureKind.
var (
	AutotuneResult_FailureKind_name = map[int32]string{
		0: "UNKNOWN",
		1: "REDZONE_MODIFIED",
		2: "WRONG_RESULT",
	}
	AutotuneResult_FailureKind_value = map[string]int32{
		"UNKNOWN":          0,
		"REDZONE_MODIFIED": 1,
		"WRONG_RESULT":     2,
	}
)

func (x AutotuneResult_FailureKind) Enum() *AutotuneResult_FailureKind {
	p := new(AutotuneResult_FailureKind)
	*p = x
	return p
}

func (x AutotuneResult_FailureKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AutotuneResult_FailureKind) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_protobuf_autotuning_proto_enumTypes[0].Descriptor()
}

func (AutotuneResult_FailureKind) Type() protoreflect.EnumType {
	return &file_tensorflow_core_protobuf_autotuning_proto_enumTypes[0]
}

func (x AutotuneResult_FailureKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AutotuneResult_FailureKind.Descriptor instead.
func (AutotuneResult_FailureKind) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_autotuning_proto_rawDescGZIP(), []int{2, 0}
}

type CudnnVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major int32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor int32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch int32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
}

func (x *CudnnVersion) Reset() {
	*x = CudnnVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CudnnVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CudnnVersion) ProtoMessage() {}

func (x *CudnnVersion) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CudnnVersion.ProtoReflect.Descriptor instead.
func (*CudnnVersion) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_autotuning_proto_rawDescGZIP(), []int{0}
}

func (x *CudnnVersion) GetMajor() int32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *CudnnVersion) GetMinor() int32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *CudnnVersion) GetPatch() int32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

type ComputeCapability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major int32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor int32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
}

func (x *ComputeCapability) Reset() {
	*x = ComputeCapability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeCapability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeCapability) ProtoMessage() {}

func (x *ComputeCapability) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeCapability.ProtoReflect.Descriptor instead.
func (*ComputeCapability) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_autotuning_proto_rawDescGZIP(), []int{1}
}

func (x *ComputeCapability) GetMajor() int32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *ComputeCapability) GetMinor() int32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

type AutotuneResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScratchBytes int64                         `protobuf:"varint,8,opt,name=scratch_bytes,json=scratchBytes,proto3" json:"scratch_bytes,omitempty"`
	RunTime      *durationpb.Duration          `protobuf:"bytes,9,opt,name=run_time,json=runTime,proto3" json:"run_time,omitempty"`
	Failure      *AutotuneResult_FailureResult `protobuf:"bytes,7,opt,name=failure,proto3" json:"failure,omitempty"`
	// Types that are assignable to Key:
	//	*AutotuneResult_Conv
	//	*AutotuneResult_Gemm
	Key isAutotuneResult_Key `protobuf_oneof:"key"`
}

func (x *AutotuneResult) Reset() {
	*x = AutotuneResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutotuneResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutotuneResult) ProtoMessage() {}

func (x *AutotuneResult) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutotuneResult.ProtoReflect.Descriptor instead.
func (*AutotuneResult) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_autotuning_proto_rawDescGZIP(), []int{2}
}

func (x *AutotuneResult) GetScratchBytes() int64 {
	if x != nil {
		return x.ScratchBytes
	}
	return 0
}

func (x *AutotuneResult) GetRunTime() *durationpb.Duration {
	if x != nil {
		return x.RunTime
	}
	return nil
}

func (x *AutotuneResult) GetFailure() *AutotuneResult_FailureResult {
	if x != nil {
		return x.Failure
	}
	return nil
}

func (m *AutotuneResult) GetKey() isAutotuneResult_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (x *AutotuneResult) GetConv() *AutotuneResult_ConvKey {
	if x, ok := x.GetKey().(*AutotuneResult_Conv); ok {
		return x.Conv
	}
	return nil
}

func (x *AutotuneResult) GetGemm() *AutotuneResult_GemmKey {
	if x, ok := x.GetKey().(*AutotuneResult_Gemm); ok {
		return x.Gemm
	}
	return nil
}

type isAutotuneResult_Key interface {
	isAutotuneResult_Key()
}

type AutotuneResult_Conv struct {
	Conv *AutotuneResult_ConvKey `protobuf:"bytes,5,opt,name=conv,proto3,oneof"`
}

type AutotuneResult_Gemm struct {
	Gemm *AutotuneResult_GemmKey `protobuf:"bytes,6,opt,name=gemm,proto3,oneof"`
}

func (*AutotuneResult_Conv) isAutotuneResult_Key() {}

func (*AutotuneResult_Gemm) isAutotuneResult_Key() {}

type AutotuningLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instr *anypb.Any `protobuf:"bytes,1,opt,name=instr,proto3" json:"instr,omitempty"`
	// Records all auto-tuning results per algorithm.
	Results           []*AutotuneResult  `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	CudnnVersion      *CudnnVersion      `protobuf:"bytes,3,opt,name=cudnn_version,json=cudnnVersion,proto3" json:"cudnn_version,omitempty"`
	ComputeCapability *ComputeCapability `protobuf:"bytes,4,opt,name=compute_capability,json=computeCapability,proto3" json:"compute_capability,omitempty"`
	// stream_executor::DeviceDescription::pci_bus_id.
	DevicePciBusId string `protobuf:"bytes,5,opt,name=device_pci_bus_id,json=devicePciBusId,proto3" json:"device_pci_bus_id,omitempty"`
	BlasVersion    string `protobuf:"bytes,6,opt,name=blas_version,json=blasVersion,proto3" json:"blas_version,omitempty"`
}

func (x *AutotuningLog) Reset() {
	*x = AutotuningLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutotuningLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutotuningLog) ProtoMessage() {}

func (x *AutotuningLog) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutotuningLog.ProtoReflect.Descriptor instead.
func (*AutotuningLog) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_autotuning_proto_rawDescGZIP(), []int{3}
}

func (x *AutotuningLog) GetInstr() *anypb.Any {
	if x != nil {
		return x.Instr
	}
	return nil
}

func (x *AutotuningLog) GetResults() []*AutotuneResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *AutotuningLog) GetCudnnVersion() *CudnnVersion {
	if x != nil {
		return x.CudnnVersion
	}
	return nil
}

func (x *AutotuningLog) GetComputeCapability() *ComputeCapability {
	if x != nil {
		return x.ComputeCapability
	}
	return nil
}

func (x *AutotuningLog) GetDevicePciBusId() string {
	if x != nil {
		return x.DevicePciBusId
	}
	return ""
}

func (x *AutotuningLog) GetBlasVersion() string {
	if x != nil {
		return x.BlasVersion
	}
	return ""
}

type AutotuneResult_FailureResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind AutotuneResult_FailureKind `protobuf:"varint,1,opt,name=kind,proto3,enum=tensorflow.AutotuneResult_FailureKind" json:"kind,omitempty"`
	Msg  string                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// For failure_kind == WRONG_RESULT, this field indicates the reference
	// configuration that we compared against.
	//
	// Note that the reference algorithm isn't always correct.  However,
	// empirically it's more correct, as it's "algo 0", less fancy than the
	// compared one.
	//
	// Types that are assignable to Key:
	//	*AutotuneResult_FailureResult_ReferenceConv
	//	*AutotuneResult_FailureResult_ReferenceGemm
	Key           isAutotuneResult_FailureResult_Key `protobuf_oneof:"key"`
	BufferAddress int64                              `protobuf:"varint,13,opt,name=buffer_address,json=bufferAddress,proto3" json:"buffer_address,omitempty"`
}

func (x *AutotuneResult_FailureResult) Reset() {
	*x = AutotuneResult_FailureResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutotuneResult_FailureResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutotuneResult_FailureResult) ProtoMessage() {}

func (x *AutotuneResult_FailureResult) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutotuneResult_FailureResult.ProtoReflect.Descriptor instead.
func (*AutotuneResult_FailureResult) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_autotuning_proto_rawDescGZIP(), []int{2, 0}
}

func (x *AutotuneResult_FailureResult) GetKind() AutotuneResult_FailureKind {
	if x != nil {
		return x.Kind
	}
	return AutotuneResult_UNKNOWN
}

func (x *AutotuneResult_FailureResult) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (m *AutotuneResult_FailureResult) GetKey() isAutotuneResult_FailureResult_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (x *AutotuneResult_FailureResult) GetReferenceConv() *AutotuneResult_ConvKey {
	if x, ok := x.GetKey().(*AutotuneResult_FailureResult_ReferenceConv); ok {
		return x.ReferenceConv
	}
	return nil
}

func (x *AutotuneResult_FailureResult) GetReferenceGemm() *AutotuneResult_GemmKey {
	if x, ok := x.GetKey().(*AutotuneResult_FailureResult_ReferenceGemm); ok {
		return x.ReferenceGemm
	}
	return nil
}

func (x *AutotuneResult_FailureResult) GetBufferAddress() int64 {
	if x != nil {
		return x.BufferAddress
	}
	return 0
}

type isAutotuneResult_FailureResult_Key interface {
	isAutotuneResult_FailureResult_Key()
}

type AutotuneResult_FailureResult_ReferenceConv struct {
	ReferenceConv *AutotuneResult_ConvKey `protobuf:"bytes,11,opt,name=reference_conv,json=referenceConv,proto3,oneof"`
}

type AutotuneResult_FailureResult_ReferenceGemm struct {
	ReferenceGemm *AutotuneResult_GemmKey `protobuf:"bytes,12,opt,name=reference_gemm,json=referenceGemm,proto3,oneof"`
}

func (*AutotuneResult_FailureResult_ReferenceConv) isAutotuneResult_FailureResult_Key() {}

func (*AutotuneResult_FailureResult_ReferenceGemm) isAutotuneResult_FailureResult_Key() {}

type AutotuneResult_ConvKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm        int64 `protobuf:"varint,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	TensorOpsEnabled bool  `protobuf:"varint,2,opt,name=tensor_ops_enabled,json=tensorOpsEnabled,proto3" json:"tensor_ops_enabled,omitempty"`
}

func (x *AutotuneResult_ConvKey) Reset() {
	*x = AutotuneResult_ConvKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutotuneResult_ConvKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutotuneResult_ConvKey) ProtoMessage() {}

func (x *AutotuneResult_ConvKey) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutotuneResult_ConvKey.ProtoReflect.Descriptor instead.
func (*AutotuneResult_ConvKey) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_autotuning_proto_rawDescGZIP(), []int{2, 1}
}

func (x *AutotuneResult_ConvKey) GetAlgorithm() int64 {
	if x != nil {
		return x.Algorithm
	}
	return 0
}

func (x *AutotuneResult_ConvKey) GetTensorOpsEnabled() bool {
	if x != nil {
		return x.TensorOpsEnabled
	}
	return false
}

type AutotuneResult_GemmKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm int64 `protobuf:"varint,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
}

func (x *AutotuneResult_GemmKey) Reset() {
	*x = AutotuneResult_GemmKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutotuneResult_GemmKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutotuneResult_GemmKey) ProtoMessage() {}

func (x *AutotuneResult_GemmKey) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_autotuning_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutotuneResult_GemmKey.ProtoReflect.Descriptor instead.
func (*AutotuneResult_GemmKey) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_autotuning_proto_rawDescGZIP(), []int{2, 2}
}

func (x *AutotuneResult_GemmKey) GetAlgorithm() int64 {
	if x != nil {
		return x.Algorithm
	}
	return 0
}

var File_tensorflow_core_protobuf_autotuning_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_autotuning_proto_rawDesc = []byte{
	0x0a, 0x29, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x74,
	0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0c, 0x43, 0x75, 0x64, 0x6e, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x22, 0x3f, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x22, 0x96, 0x06, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x6f, 0x74, 0x75,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x63, 0x72, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x73, 0x63, 0x72, 0x61, 0x74, 0x63, 0x68, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a,
	0x08, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x76, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x6e,
	0x76, 0x12, 0x38, 0x0a, 0x04, 0x67, 0x65, 0x6d, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x75, 0x74,
	0x6f, 0x74, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x47, 0x65, 0x6d, 0x6d,
	0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x04, 0x67, 0x65, 0x6d, 0x6d, 0x1a, 0xa5, 0x02, 0x0a, 0x0d,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3a, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x4b, 0x0a, 0x0e, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x41, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x76, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x12, 0x4b, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x67, 0x65, 0x6d, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x47, 0x65, 0x6d,
	0x6d, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x47, 0x65, 0x6d, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x05, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x1a, 0x55, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x76, 0x4b, 0x65, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x2c, 0x0a, 0x12,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x6f, 0x70, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x4f, 0x70, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x1a, 0x27, 0x0a, 0x07, 0x47, 0x65,
	0x6d, 0x6d, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x22, 0x42, 0x0a, 0x0b, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x4b, 0x69,
	0x6e, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x52, 0x45, 0x44, 0x5a, 0x4f, 0x4e, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x02, 0x42, 0x05, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xcc,
	0x02, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67,
	0x12, 0x2a, 0x0a, 0x05, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x12, 0x34, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x74,
	0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0d, 0x63, 0x75, 0x64, 0x6e, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x75, 0x64, 0x6e, 0x6e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x75, 0x64, 0x6e, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x4c, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x11, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x29, 0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x63, 0x69, 0x5f, 0x62, 0x75,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x63, 0x69, 0x42, 0x75, 0x73, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c,
	0x61, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x62, 0x6c, 0x61, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x57, 0x5a,
	0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f,
	0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_autotuning_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_autotuning_proto_rawDescData = file_tensorflow_core_protobuf_autotuning_proto_rawDesc
)

func file_tensorflow_core_protobuf_autotuning_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_autotuning_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_autotuning_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_autotuning_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_autotuning_proto_rawDescData
}

var file_tensorflow_core_protobuf_autotuning_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_protobuf_autotuning_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tensorflow_core_protobuf_autotuning_proto_goTypes = []interface{}{
	(AutotuneResult_FailureKind)(0),      // 0: tensorflow.AutotuneResult.FailureKind
	(*CudnnVersion)(nil),                 // 1: tensorflow.CudnnVersion
	(*ComputeCapability)(nil),            // 2: tensorflow.ComputeCapability
	(*AutotuneResult)(nil),               // 3: tensorflow.AutotuneResult
	(*AutotuningLog)(nil),                // 4: tensorflow.AutotuningLog
	(*AutotuneResult_FailureResult)(nil), // 5: tensorflow.AutotuneResult.FailureResult
	(*AutotuneResult_ConvKey)(nil),       // 6: tensorflow.AutotuneResult.ConvKey
	(*AutotuneResult_GemmKey)(nil),       // 7: tensorflow.AutotuneResult.GemmKey
	(*durationpb.Duration)(nil),          // 8: google.protobuf.Duration
	(*anypb.Any)(nil),                    // 9: google.protobuf.Any
}
var file_tensorflow_core_protobuf_autotuning_proto_depIdxs = []int32{
	8,  // 0: tensorflow.AutotuneResult.run_time:type_name -> google.protobuf.Duration
	5,  // 1: tensorflow.AutotuneResult.failure:type_name -> tensorflow.AutotuneResult.FailureResult
	6,  // 2: tensorflow.AutotuneResult.conv:type_name -> tensorflow.AutotuneResult.ConvKey
	7,  // 3: tensorflow.AutotuneResult.gemm:type_name -> tensorflow.AutotuneResult.GemmKey
	9,  // 4: tensorflow.AutotuningLog.instr:type_name -> google.protobuf.Any
	3,  // 5: tensorflow.AutotuningLog.results:type_name -> tensorflow.AutotuneResult
	1,  // 6: tensorflow.AutotuningLog.cudnn_version:type_name -> tensorflow.CudnnVersion
	2,  // 7: tensorflow.AutotuningLog.compute_capability:type_name -> tensorflow.ComputeCapability
	0,  // 8: tensorflow.AutotuneResult.FailureResult.kind:type_name -> tensorflow.AutotuneResult.FailureKind
	6,  // 9: tensorflow.AutotuneResult.FailureResult.reference_conv:type_name -> tensorflow.AutotuneResult.ConvKey
	7,  // 10: tensorflow.AutotuneResult.FailureResult.reference_gemm:type_name -> tensorflow.AutotuneResult.GemmKey
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_autotuning_proto_init() }
func file_tensorflow_core_protobuf_autotuning_proto_init() {
	if File_tensorflow_core_protobuf_autotuning_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_autotuning_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CudnnVersion); i {
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
		file_tensorflow_core_protobuf_autotuning_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeCapability); i {
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
		file_tensorflow_core_protobuf_autotuning_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutotuneResult); i {
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
		file_tensorflow_core_protobuf_autotuning_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutotuningLog); i {
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
		file_tensorflow_core_protobuf_autotuning_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutotuneResult_FailureResult); i {
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
		file_tensorflow_core_protobuf_autotuning_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutotuneResult_ConvKey); i {
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
		file_tensorflow_core_protobuf_autotuning_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutotuneResult_GemmKey); i {
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
	file_tensorflow_core_protobuf_autotuning_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AutotuneResult_Conv)(nil),
		(*AutotuneResult_Gemm)(nil),
	}
	file_tensorflow_core_protobuf_autotuning_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*AutotuneResult_FailureResult_ReferenceConv)(nil),
		(*AutotuneResult_FailureResult_ReferenceGemm)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_protobuf_autotuning_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_autotuning_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_autotuning_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_protobuf_autotuning_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_protobuf_autotuning_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_autotuning_proto = out.File
	file_tensorflow_core_protobuf_autotuning_proto_rawDesc = nil
	file_tensorflow_core_protobuf_autotuning_proto_goTypes = nil
	file_tensorflow_core_protobuf_autotuning_proto_depIdxs = nil
}
