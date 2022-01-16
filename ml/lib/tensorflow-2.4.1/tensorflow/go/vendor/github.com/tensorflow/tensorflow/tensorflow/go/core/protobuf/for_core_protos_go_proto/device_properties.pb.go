// Copyright 2017 The TensorFlow Authors. All Rights Reserved.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//==============================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.0
// source: tensorflow/core/protobuf/device_properties.proto

package for_core_protos_go_proto

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

type DeviceProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device type (CPU, GPU, ...)
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Vendor (Intel, nvidia, ...)
	Vendor string `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// Model (Haswell, K40, ...)
	Model string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	// Core Frequency in Mhz
	Frequency int64 `protobuf:"varint,4,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Number of cores
	NumCores int64 `protobuf:"varint,5,opt,name=num_cores,json=numCores,proto3" json:"num_cores,omitempty"`
	// Version of the tools and libraries used with this device (e.g. gcc 4.9,
	// cudnn 5.1)
	Environment map[string]string `protobuf:"bytes,6,rep,name=environment,proto3" json:"environment,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Number of registers per core.
	NumRegisters int64 `protobuf:"varint,7,opt,name=num_registers,json=numRegisters,proto3" json:"num_registers,omitempty"`
	// L1 cache size in bytes
	L1CacheSize int64 `protobuf:"varint,8,opt,name=l1_cache_size,json=l1CacheSize,proto3" json:"l1_cache_size,omitempty"`
	// L2 cache size in bytes
	L2CacheSize int64 `protobuf:"varint,9,opt,name=l2_cache_size,json=l2CacheSize,proto3" json:"l2_cache_size,omitempty"`
	// L3 cache size in bytes
	L3CacheSize int64 `protobuf:"varint,10,opt,name=l3_cache_size,json=l3CacheSize,proto3" json:"l3_cache_size,omitempty"`
	// Shared memory size per multiprocessor in bytes. This field is
	// applicable to GPUs only.
	SharedMemorySizePerMultiprocessor int64 `protobuf:"varint,11,opt,name=shared_memory_size_per_multiprocessor,json=sharedMemorySizePerMultiprocessor,proto3" json:"shared_memory_size_per_multiprocessor,omitempty"`
	// Memory size in bytes
	MemorySize int64 `protobuf:"varint,12,opt,name=memory_size,json=memorySize,proto3" json:"memory_size,omitempty"`
	// Memory bandwidth in KB/s
	Bandwidth int64 `protobuf:"varint,13,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
}

func (x *DeviceProperties) Reset() {
	*x = DeviceProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_device_properties_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceProperties) ProtoMessage() {}

func (x *DeviceProperties) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_device_properties_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceProperties.ProtoReflect.Descriptor instead.
func (*DeviceProperties) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_device_properties_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceProperties) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DeviceProperties) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *DeviceProperties) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *DeviceProperties) GetFrequency() int64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *DeviceProperties) GetNumCores() int64 {
	if x != nil {
		return x.NumCores
	}
	return 0
}

func (x *DeviceProperties) GetEnvironment() map[string]string {
	if x != nil {
		return x.Environment
	}
	return nil
}

func (x *DeviceProperties) GetNumRegisters() int64 {
	if x != nil {
		return x.NumRegisters
	}
	return 0
}

func (x *DeviceProperties) GetL1CacheSize() int64 {
	if x != nil {
		return x.L1CacheSize
	}
	return 0
}

func (x *DeviceProperties) GetL2CacheSize() int64 {
	if x != nil {
		return x.L2CacheSize
	}
	return 0
}

func (x *DeviceProperties) GetL3CacheSize() int64 {
	if x != nil {
		return x.L3CacheSize
	}
	return 0
}

func (x *DeviceProperties) GetSharedMemorySizePerMultiprocessor() int64 {
	if x != nil {
		return x.SharedMemorySizePerMultiprocessor
	}
	return 0
}

func (x *DeviceProperties) GetMemorySize() int64 {
	if x != nil {
		return x.MemorySize
	}
	return 0
}

func (x *DeviceProperties) GetBandwidth() int64 {
	if x != nil {
		return x.Bandwidth
	}
	return 0
}

type NamedDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Properties *DeviceProperties `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *NamedDevice) Reset() {
	*x = NamedDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_device_properties_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamedDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedDevice) ProtoMessage() {}

func (x *NamedDevice) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_device_properties_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedDevice.ProtoReflect.Descriptor instead.
func (*NamedDevice) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_device_properties_proto_rawDescGZIP(), []int{1}
}

func (x *NamedDevice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamedDevice) GetProperties() *DeviceProperties {
	if x != nil {
		return x.Properties
	}
	return nil
}

var File_tensorflow_core_protobuf_device_properties_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_device_properties_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xc2,
	0x04, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x72, 0x65, 0x73,
	0x12, 0x4f, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x31, 0x5f, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c,
	0x31, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x32,
	0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6c, 0x32, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22,
	0x0a, 0x0d, 0x6c, 0x33, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x33, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x50, 0x0a, 0x25, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x21, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53,
	0x69, 0x7a, 0x65, 0x50, 0x65, 0x72, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x1a, 0x3e, 0x0a, 0x10, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x5f, 0x0a, 0x0b, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x42, 0x72, 0x42, 0x16, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x55,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_device_properties_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_device_properties_proto_rawDescData = file_tensorflow_core_protobuf_device_properties_proto_rawDesc
)

func file_tensorflow_core_protobuf_device_properties_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_device_properties_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_device_properties_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_device_properties_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_device_properties_proto_rawDescData
}

var file_tensorflow_core_protobuf_device_properties_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_core_protobuf_device_properties_proto_goTypes = []interface{}{
	(*DeviceProperties)(nil), // 0: tensorflow.DeviceProperties
	(*NamedDevice)(nil),      // 1: tensorflow.NamedDevice
	nil,                      // 2: tensorflow.DeviceProperties.EnvironmentEntry
}
var file_tensorflow_core_protobuf_device_properties_proto_depIdxs = []int32{
	2, // 0: tensorflow.DeviceProperties.environment:type_name -> tensorflow.DeviceProperties.EnvironmentEntry
	0, // 1: tensorflow.NamedDevice.properties:type_name -> tensorflow.DeviceProperties
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_device_properties_proto_init() }
func file_tensorflow_core_protobuf_device_properties_proto_init() {
	if File_tensorflow_core_protobuf_device_properties_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_device_properties_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceProperties); i {
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
		file_tensorflow_core_protobuf_device_properties_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamedDevice); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_device_properties_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_device_properties_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_device_properties_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_device_properties_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_device_properties_proto = out.File
	file_tensorflow_core_protobuf_device_properties_proto_rawDesc = nil
	file_tensorflow_core_protobuf_device_properties_proto_goTypes = nil
	file_tensorflow_core_protobuf_device_properties_proto_depIdxs = nil
}
