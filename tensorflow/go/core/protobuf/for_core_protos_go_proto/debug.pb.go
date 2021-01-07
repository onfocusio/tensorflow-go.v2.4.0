// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.14.0
// source: tensorflow/core/protobuf/debug.proto

package for_core_protos_go_proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Option for watching a node in TensorFlow Debugger (tfdbg).
type DebugTensorWatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the node to watch.
	// Use "*" for wildcard. But note: currently, regex is not supported in
	// general.
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Output slot to watch.
	// The semantics of output_slot == -1 is that all outputs of the node
	// will be watched (i.e., a wildcard).
	// Other negative values of output_slot are invalid and will lead to
	// errors currently.
	OutputSlot int32 `protobuf:"varint,2,opt,name=output_slot,json=outputSlot,proto3" json:"output_slot,omitempty"`
	// Name(s) of the debugging op(s).
	// One or more than one probes on a tensor.
	// e.g., {"DebugIdentity", "DebugNanCount"}
	DebugOps []string `protobuf:"bytes,3,rep,name=debug_ops,json=debugOps,proto3" json:"debug_ops,omitempty"`
	// URL(s) for debug targets(s).
	//
	// Supported URL formats are:
	//   - file:///foo/tfdbg_dump: Writes out Event content to file
	//     /foo/tfdbg_dump.  Assumes all directories can be created if they don't
	//     already exist.
	//   - grpc://localhost:11011: Sends an RPC request to an EventListener
	//     service running at localhost:11011 with the event.
	//   - memcbk:///event_key: Routes tensors to clients using the
	//     callback registered with the DebugCallbackRegistry for event_key.
	//
	// Each debug op listed in debug_ops will publish its output tensor (debug
	// signal) to all URLs in debug_urls.
	//
	// N.B. Session::Run() supports concurrent invocations of the same inputs
	// (feed keys), outputs and target nodes. If such concurrent invocations
	// are to be debugged, the callers of Session::Run() must use distinct
	// debug_urls to make sure that the streamed or dumped events do not overlap
	// among the invocations.
	// TODO(cais): More visible documentation of this in g3docs.
	DebugUrls []string `protobuf:"bytes,4,rep,name=debug_urls,json=debugUrls,proto3" json:"debug_urls,omitempty"`
	// Do not error out if debug op creation fails (e.g., due to dtype
	// incompatibility). Instead, just log the failure.
	TolerateDebugOpCreationFailures bool `protobuf:"varint,5,opt,name=tolerate_debug_op_creation_failures,json=tolerateDebugOpCreationFailures,proto3" json:"tolerate_debug_op_creation_failures,omitempty"`
}

func (x *DebugTensorWatch) Reset() {
	*x = DebugTensorWatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_debug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugTensorWatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugTensorWatch) ProtoMessage() {}

func (x *DebugTensorWatch) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_debug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugTensorWatch.ProtoReflect.Descriptor instead.
func (*DebugTensorWatch) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_debug_proto_rawDescGZIP(), []int{0}
}

func (x *DebugTensorWatch) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *DebugTensorWatch) GetOutputSlot() int32 {
	if x != nil {
		return x.OutputSlot
	}
	return 0
}

func (x *DebugTensorWatch) GetDebugOps() []string {
	if x != nil {
		return x.DebugOps
	}
	return nil
}

func (x *DebugTensorWatch) GetDebugUrls() []string {
	if x != nil {
		return x.DebugUrls
	}
	return nil
}

func (x *DebugTensorWatch) GetTolerateDebugOpCreationFailures() bool {
	if x != nil {
		return x.TolerateDebugOpCreationFailures
	}
	return false
}

// Options for initializing DebuggerState in TensorFlow Debugger (tfdbg).
type DebugOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Debugging options
	DebugTensorWatchOpts []*DebugTensorWatch `protobuf:"bytes,4,rep,name=debug_tensor_watch_opts,json=debugTensorWatchOpts,proto3" json:"debug_tensor_watch_opts,omitempty"`
	// Caller-specified global step count.
	// Note that this is distinct from the session run count and the executor
	// step count.
	GlobalStep int64 `protobuf:"varint,10,opt,name=global_step,json=globalStep,proto3" json:"global_step,omitempty"`
	// Whether the total disk usage of tfdbg is to be reset to zero
	// in this Session.run call. This is used by wrappers and hooks
	// such as the local CLI ones to indicate that the dumped tensors
	// are cleaned up from the disk after each Session.run.
	ResetDiskByteUsage bool `protobuf:"varint,11,opt,name=reset_disk_byte_usage,json=resetDiskByteUsage,proto3" json:"reset_disk_byte_usage,omitempty"`
}

func (x *DebugOptions) Reset() {
	*x = DebugOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_debug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugOptions) ProtoMessage() {}

func (x *DebugOptions) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_debug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugOptions.ProtoReflect.Descriptor instead.
func (*DebugOptions) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_debug_proto_rawDescGZIP(), []int{1}
}

func (x *DebugOptions) GetDebugTensorWatchOpts() []*DebugTensorWatch {
	if x != nil {
		return x.DebugTensorWatchOpts
	}
	return nil
}

func (x *DebugOptions) GetGlobalStep() int64 {
	if x != nil {
		return x.GlobalStep
	}
	return 0
}

func (x *DebugOptions) GetResetDiskByteUsage() bool {
	if x != nil {
		return x.ResetDiskByteUsage
	}
	return false
}

type DebuggedSourceFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The host name on which a source code file is located.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Path to the source code file.
	FilePath string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// The timestamp at which the source code file is last modified.
	LastModified int64 `protobuf:"varint,3,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	// Byte size of the file.
	Bytes int64 `protobuf:"varint,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
	// Line-by-line content of the source code file.
	Lines []string `protobuf:"bytes,5,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *DebuggedSourceFile) Reset() {
	*x = DebuggedSourceFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_debug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebuggedSourceFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebuggedSourceFile) ProtoMessage() {}

func (x *DebuggedSourceFile) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_debug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebuggedSourceFile.ProtoReflect.Descriptor instead.
func (*DebuggedSourceFile) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_debug_proto_rawDescGZIP(), []int{2}
}

func (x *DebuggedSourceFile) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *DebuggedSourceFile) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *DebuggedSourceFile) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

func (x *DebuggedSourceFile) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *DebuggedSourceFile) GetLines() []string {
	if x != nil {
		return x.Lines
	}
	return nil
}

type DebuggedSourceFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A collection of source code files.
	SourceFiles []*DebuggedSourceFile `protobuf:"bytes,1,rep,name=source_files,json=sourceFiles,proto3" json:"source_files,omitempty"`
}

func (x *DebuggedSourceFiles) Reset() {
	*x = DebuggedSourceFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_debug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebuggedSourceFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebuggedSourceFiles) ProtoMessage() {}

func (x *DebuggedSourceFiles) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_debug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebuggedSourceFiles.ProtoReflect.Descriptor instead.
func (*DebuggedSourceFiles) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_debug_proto_rawDescGZIP(), []int{3}
}

func (x *DebuggedSourceFiles) GetSourceFiles() []*DebuggedSourceFile {
	if x != nil {
		return x.SourceFiles
	}
	return nil
}

var File_tensorflow_core_protobuf_debug_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_debug_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x22, 0xda, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x62, 0x75, 0x67, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x73,
	0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x6f,
	0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x62, 0x75, 0x67, 0x4f,
	0x70, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x62, 0x75, 0x67, 0x55, 0x72, 0x6c,
	0x73, 0x12, 0x4c, 0x0a, 0x23, 0x74, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65,
	0x62, 0x75, 0x67, 0x5f, 0x6f, 0x70, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1f,
	0x74, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x65, 0x62, 0x75, 0x67, 0x4f, 0x70, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x22,
	0xb7, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x62, 0x75, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x53, 0x0a, 0x17, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x14, 0x64, 0x65, 0x62, 0x75, 0x67, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f,
	0x73, 0x74, 0x65, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x12, 0x31, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f,
	0x64, 0x69, 0x73, 0x6b, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x72, 0x65, 0x73, 0x65, 0x74, 0x44, 0x69, 0x73, 0x6b,
	0x42, 0x79, 0x74, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x12, 0x44, 0x65,
	0x62, 0x75, 0x67, 0x67, 0x65, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x22, 0x58, 0x0a, 0x13, 0x44, 0x65, 0x62, 0x75, 0x67, 0x67, 0x65, 0x64, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x65, 0x62,
	0x75, 0x67, 0x67, 0x65, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x83, 0x01, 0x0a,
	0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0b, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8,
	0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_debug_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_debug_proto_rawDescData = file_tensorflow_core_protobuf_debug_proto_rawDesc
)

func file_tensorflow_core_protobuf_debug_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_debug_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_debug_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_debug_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_debug_proto_rawDescData
}

var file_tensorflow_core_protobuf_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_core_protobuf_debug_proto_goTypes = []interface{}{
	(*DebugTensorWatch)(nil),    // 0: tensorflow.DebugTensorWatch
	(*DebugOptions)(nil),        // 1: tensorflow.DebugOptions
	(*DebuggedSourceFile)(nil),  // 2: tensorflow.DebuggedSourceFile
	(*DebuggedSourceFiles)(nil), // 3: tensorflow.DebuggedSourceFiles
}
var file_tensorflow_core_protobuf_debug_proto_depIdxs = []int32{
	0, // 0: tensorflow.DebugOptions.debug_tensor_watch_opts:type_name -> tensorflow.DebugTensorWatch
	2, // 1: tensorflow.DebuggedSourceFiles.source_files:type_name -> tensorflow.DebuggedSourceFile
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_debug_proto_init() }
func file_tensorflow_core_protobuf_debug_proto_init() {
	if File_tensorflow_core_protobuf_debug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_debug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugTensorWatch); i {
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
		file_tensorflow_core_protobuf_debug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugOptions); i {
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
		file_tensorflow_core_protobuf_debug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebuggedSourceFile); i {
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
		file_tensorflow_core_protobuf_debug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebuggedSourceFiles); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_debug_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_debug_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_debug_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_debug_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_debug_proto = out.File
	file_tensorflow_core_protobuf_debug_proto_rawDesc = nil
	file_tensorflow_core_protobuf_debug_proto_goTypes = nil
	file_tensorflow_core_protobuf_debug_proto_depIdxs = nil
}
