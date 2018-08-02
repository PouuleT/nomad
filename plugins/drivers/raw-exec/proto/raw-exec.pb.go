// Code generated by protoc-gen-go. DO NOT EDIT.
// source: raw-exec.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Port struct {
	Label                string   `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Value                uint32   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_raw_exec_da16efce2a45e7f7, []int{0}
}
func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (dst *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(dst, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Port) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Networks struct {
	Device               string   `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Cidr                 string   `protobuf:"bytes,2,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Mbits                uint64   `protobuf:"varint,4,opt,name=mbits,proto3" json:"mbits,omitempty"`
	ReservedPorts        *Port    `protobuf:"bytes,5,opt,name=reserved_ports,json=reservedPorts,proto3" json:"reserved_ports,omitempty"`
	DynamicPorts         *Port    `protobuf:"bytes,6,opt,name=dynamic_ports,json=dynamicPorts,proto3" json:"dynamic_ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Networks) Reset()         { *m = Networks{} }
func (m *Networks) String() string { return proto.CompactTextString(m) }
func (*Networks) ProtoMessage()    {}
func (*Networks) Descriptor() ([]byte, []int) {
	return fileDescriptor_raw_exec_da16efce2a45e7f7, []int{1}
}
func (m *Networks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Networks.Unmarshal(m, b)
}
func (m *Networks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Networks.Marshal(b, m, deterministic)
}
func (dst *Networks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Networks.Merge(dst, src)
}
func (m *Networks) XXX_Size() int {
	return xxx_messageInfo_Networks.Size(m)
}
func (m *Networks) XXX_DiscardUnknown() {
	xxx_messageInfo_Networks.DiscardUnknown(m)
}

var xxx_messageInfo_Networks proto.InternalMessageInfo

func (m *Networks) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *Networks) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

func (m *Networks) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Networks) GetMbits() uint64 {
	if m != nil {
		return m.Mbits
	}
	return 0
}

func (m *Networks) GetReservedPorts() *Port {
	if m != nil {
		return m.ReservedPorts
	}
	return nil
}

func (m *Networks) GetDynamicPorts() *Port {
	if m != nil {
		return m.DynamicPorts
	}
	return nil
}

type Resources struct {
	Cpu                  uint64    `protobuf:"varint,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	MemoryMb             uint64    `protobuf:"varint,2,opt,name=memory_mb,json=memoryMb,proto3" json:"memory_mb,omitempty"`
	DiskMb               uint64    `protobuf:"varint,3,opt,name=disk_mb,json=diskMb,proto3" json:"disk_mb,omitempty"`
	Iops                 uint64    `protobuf:"varint,4,opt,name=iops,proto3" json:"iops,omitempty"`
	Networks             *Networks `protobuf:"bytes,5,opt,name=networks,proto3" json:"networks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_raw_exec_da16efce2a45e7f7, []int{2}
}
func (m *Resources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources.Unmarshal(m, b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
}
func (dst *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(dst, src)
}
func (m *Resources) XXX_Size() int {
	return xxx_messageInfo_Resources.Size(m)
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func (m *Resources) GetCpu() uint64 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *Resources) GetMemoryMb() uint64 {
	if m != nil {
		return m.MemoryMb
	}
	return 0
}

func (m *Resources) GetDiskMb() uint64 {
	if m != nil {
		return m.DiskMb
	}
	return 0
}

func (m *Resources) GetIops() uint64 {
	if m != nil {
		return m.Iops
	}
	return 0
}

func (m *Resources) GetNetworks() *Networks {
	if m != nil {
		return m.Networks
	}
	return nil
}

type LogConfig struct {
	MaxFiles             uint32   `protobuf:"varint,1,opt,name=max_files,json=maxFiles,proto3" json:"max_files,omitempty"`
	MaxFileSizeMb        uint64   `protobuf:"varint,2,opt,name=max_file_size_mb,json=maxFileSizeMb,proto3" json:"max_file_size_mb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogConfig) Reset()         { *m = LogConfig{} }
func (m *LogConfig) String() string { return proto.CompactTextString(m) }
func (*LogConfig) ProtoMessage()    {}
func (*LogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_raw_exec_da16efce2a45e7f7, []int{3}
}
func (m *LogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogConfig.Unmarshal(m, b)
}
func (m *LogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogConfig.Marshal(b, m, deterministic)
}
func (dst *LogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogConfig.Merge(dst, src)
}
func (m *LogConfig) XXX_Size() int {
	return xxx_messageInfo_LogConfig.Size(m)
}
func (m *LogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LogConfig proto.InternalMessageInfo

func (m *LogConfig) GetMaxFiles() uint32 {
	if m != nil {
		return m.MaxFiles
	}
	return 0
}

func (m *LogConfig) GetMaxFileSizeMb() uint64 {
	if m != nil {
		return m.MaxFileSizeMb
	}
	return 0
}

type TaskInfo struct {
	Resources *Resources `protobuf:"bytes,1,opt,name=resources,proto3" json:"resources,omitempty"`
	LogConfig *LogConfig `protobuf:"bytes,2,opt,name=log_config,json=logConfig,proto3" json:"log_config,omitempty"`
	// Struct config = 3;
	// example: https://gist.github.com/jsmouret/2bc876e8def6c63410556350eca3e43d
	Config               *_struct.Struct `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TaskInfo) Reset()         { *m = TaskInfo{} }
func (m *TaskInfo) String() string { return proto.CompactTextString(m) }
func (*TaskInfo) ProtoMessage()    {}
func (*TaskInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_raw_exec_da16efce2a45e7f7, []int{4}
}
func (m *TaskInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskInfo.Unmarshal(m, b)
}
func (m *TaskInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskInfo.Marshal(b, m, deterministic)
}
func (dst *TaskInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInfo.Merge(dst, src)
}
func (m *TaskInfo) XXX_Size() int {
	return xxx_messageInfo_TaskInfo.Size(m)
}
func (m *TaskInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TaskInfo proto.InternalMessageInfo

func (m *TaskInfo) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *TaskInfo) GetLogConfig() *LogConfig {
	if m != nil {
		return m.LogConfig
	}
	return nil
}

func (m *TaskInfo) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

type TaskDir struct {
	Directory            string   `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
	SharedAllocDir       string   `protobuf:"bytes,2,opt,name=shared_alloc_dir,json=sharedAllocDir,proto3" json:"shared_alloc_dir,omitempty"`
	SharedTaskDir        string   `protobuf:"bytes,3,opt,name=shared_task_dir,json=sharedTaskDir,proto3" json:"shared_task_dir,omitempty"`
	LocalDir             string   `protobuf:"bytes,4,opt,name=local_dir,json=localDir,proto3" json:"local_dir,omitempty"`
	LogDir               string   `protobuf:"bytes,5,opt,name=log_dir,json=logDir,proto3" json:"log_dir,omitempty"`
	SecretsDir           string   `protobuf:"bytes,6,opt,name=secrets_dir,json=secretsDir,proto3" json:"secrets_dir,omitempty"`
	LogLevel             string   `protobuf:"bytes,7,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskDir) Reset()         { *m = TaskDir{} }
func (m *TaskDir) String() string { return proto.CompactTextString(m) }
func (*TaskDir) ProtoMessage()    {}
func (*TaskDir) Descriptor() ([]byte, []int) {
	return fileDescriptor_raw_exec_da16efce2a45e7f7, []int{5}
}
func (m *TaskDir) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskDir.Unmarshal(m, b)
}
func (m *TaskDir) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskDir.Marshal(b, m, deterministic)
}
func (dst *TaskDir) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskDir.Merge(dst, src)
}
func (m *TaskDir) XXX_Size() int {
	return xxx_messageInfo_TaskDir.Size(m)
}
func (m *TaskDir) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskDir.DiscardUnknown(m)
}

var xxx_messageInfo_TaskDir proto.InternalMessageInfo

func (m *TaskDir) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *TaskDir) GetSharedAllocDir() string {
	if m != nil {
		return m.SharedAllocDir
	}
	return ""
}

func (m *TaskDir) GetSharedTaskDir() string {
	if m != nil {
		return m.SharedTaskDir
	}
	return ""
}

func (m *TaskDir) GetLocalDir() string {
	if m != nil {
		return m.LocalDir
	}
	return ""
}

func (m *TaskDir) GetLogDir() string {
	if m != nil {
		return m.LogDir
	}
	return ""
}

func (m *TaskDir) GetSecretsDir() string {
	if m != nil {
		return m.SecretsDir
	}
	return ""
}

func (m *TaskDir) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

type TaskEnv struct {
	NodeAttrs            map[string]string `protobuf:"bytes,1,rep,name=node_attrs,json=nodeAttrs,proto3" json:"node_attrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EnvMap               map[string]string `protobuf:"bytes,2,rep,name=env_map,json=envMap,proto3" json:"env_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskEnv) Reset()         { *m = TaskEnv{} }
func (m *TaskEnv) String() string { return proto.CompactTextString(m) }
func (*TaskEnv) ProtoMessage()    {}
func (*TaskEnv) Descriptor() ([]byte, []int) {
	return fileDescriptor_raw_exec_da16efce2a45e7f7, []int{6}
}
func (m *TaskEnv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskEnv.Unmarshal(m, b)
}
func (m *TaskEnv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskEnv.Marshal(b, m, deterministic)
}
func (dst *TaskEnv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskEnv.Merge(dst, src)
}
func (m *TaskEnv) XXX_Size() int {
	return xxx_messageInfo_TaskEnv.Size(m)
}
func (m *TaskEnv) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskEnv.DiscardUnknown(m)
}

var xxx_messageInfo_TaskEnv proto.InternalMessageInfo

func (m *TaskEnv) GetNodeAttrs() map[string]string {
	if m != nil {
		return m.NodeAttrs
	}
	return nil
}

func (m *TaskEnv) GetEnvMap() map[string]string {
	if m != nil {
		return m.EnvMap
	}
	return nil
}

type ExecContext struct {
	TaskDir              *TaskDir `protobuf:"bytes,1,opt,name=task_dir,json=taskDir,proto3" json:"task_dir,omitempty"`
	TaskEnv              *TaskEnv `protobuf:"bytes,2,opt,name=task_env,json=taskEnv,proto3" json:"task_env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecContext) Reset()         { *m = ExecContext{} }
func (m *ExecContext) String() string { return proto.CompactTextString(m) }
func (*ExecContext) ProtoMessage()    {}
func (*ExecContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_raw_exec_da16efce2a45e7f7, []int{7}
}
func (m *ExecContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecContext.Unmarshal(m, b)
}
func (m *ExecContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecContext.Marshal(b, m, deterministic)
}
func (dst *ExecContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecContext.Merge(dst, src)
}
func (m *ExecContext) XXX_Size() int {
	return xxx_messageInfo_ExecContext.Size(m)
}
func (m *ExecContext) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecContext.DiscardUnknown(m)
}

var xxx_messageInfo_ExecContext proto.InternalMessageInfo

func (m *ExecContext) GetTaskDir() *TaskDir {
	if m != nil {
		return m.TaskDir
	}
	return nil
}

func (m *ExecContext) GetTaskEnv() *TaskEnv {
	if m != nil {
		return m.TaskEnv
	}
	return nil
}

type StartRequest struct {
	ExecContext          *ExecContext `protobuf:"bytes,1,opt,name=exec_context,json=execContext,proto3" json:"exec_context,omitempty"`
	TaskInfo             *TaskInfo    `protobuf:"bytes,2,opt,name=task_info,json=taskInfo,proto3" json:"task_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_raw_exec_da16efce2a45e7f7, []int{8}
}
func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (dst *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(dst, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

func (m *StartRequest) GetExecContext() *ExecContext {
	if m != nil {
		return m.ExecContext
	}
	return nil
}

func (m *StartRequest) GetTaskInfo() *TaskInfo {
	if m != nil {
		return m.TaskInfo
	}
	return nil
}

type DriverNetwork struct {
	PortMap              map[string]uint32 `protobuf:"bytes,1,rep,name=port_map,json=portMap,proto3" json:"port_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Ip                   string            `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	AutoAdvertise        bool              `protobuf:"varint,3,opt,name=auto_advertise,json=autoAdvertise,proto3" json:"auto_advertise,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DriverNetwork) Reset()         { *m = DriverNetwork{} }
func (m *DriverNetwork) String() string { return proto.CompactTextString(m) }
func (*DriverNetwork) ProtoMessage()    {}
func (*DriverNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_raw_exec_da16efce2a45e7f7, []int{9}
}
func (m *DriverNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverNetwork.Unmarshal(m, b)
}
func (m *DriverNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverNetwork.Marshal(b, m, deterministic)
}
func (dst *DriverNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverNetwork.Merge(dst, src)
}
func (m *DriverNetwork) XXX_Size() int {
	return xxx_messageInfo_DriverNetwork.Size(m)
}
func (m *DriverNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_DriverNetwork proto.InternalMessageInfo

func (m *DriverNetwork) GetPortMap() map[string]uint32 {
	if m != nil {
		return m.PortMap
	}
	return nil
}

func (m *DriverNetwork) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *DriverNetwork) GetAutoAdvertise() bool {
	if m != nil {
		return m.AutoAdvertise
	}
	return false
}

type StartResponse struct {
	TaskId               string         `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	DriverNetwork        *DriverNetwork `protobuf:"bytes,2,opt,name=driver_network,json=driverNetwork,proto3" json:"driver_network,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_raw_exec_da16efce2a45e7f7, []int{10}
}
func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (dst *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(dst, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

func (m *StartResponse) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *StartResponse) GetDriverNetwork() *DriverNetwork {
	if m != nil {
		return m.DriverNetwork
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_raw_exec_da16efce2a45e7f7, []int{11}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Port)(nil), "proto.Port")
	proto.RegisterType((*Networks)(nil), "proto.Networks")
	proto.RegisterType((*Resources)(nil), "proto.Resources")
	proto.RegisterType((*LogConfig)(nil), "proto.LogConfig")
	proto.RegisterType((*TaskInfo)(nil), "proto.TaskInfo")
	proto.RegisterType((*TaskDir)(nil), "proto.TaskDir")
	proto.RegisterType((*TaskEnv)(nil), "proto.TaskEnv")
	proto.RegisterMapType((map[string]string)(nil), "proto.TaskEnv.EnvMapEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.TaskEnv.NodeAttrsEntry")
	proto.RegisterType((*ExecContext)(nil), "proto.ExecContext")
	proto.RegisterType((*StartRequest)(nil), "proto.StartRequest")
	proto.RegisterType((*DriverNetwork)(nil), "proto.DriverNetwork")
	proto.RegisterMapType((map[string]uint32)(nil), "proto.DriverNetwork.PortMapEntry")
	proto.RegisterType((*StartResponse)(nil), "proto.StartResponse")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RawExecClient is the client API for RawExec service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RawExecClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
}

type rawExecClient struct {
	cc *grpc.ClientConn
}

func NewRawExecClient(cc *grpc.ClientConn) RawExecClient {
	return &rawExecClient{cc}
}

func (c *rawExecClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/proto.RawExec/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RawExecServer is the server API for RawExec service.
type RawExecServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
}

func RegisterRawExecServer(s *grpc.Server, srv RawExecServer) {
	s.RegisterService(&_RawExec_serviceDesc, srv)
}

func _RawExec_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RawExecServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RawExec/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawExecServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RawExec_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RawExec",
	HandlerType: (*RawExecServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _RawExec_Start_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raw-exec.proto",
}

func init() { proto.RegisterFile("raw-exec.proto", fileDescriptor_raw_exec_da16efce2a45e7f7) }

var fileDescriptor_raw_exec_da16efce2a45e7f7 = []byte{
	// 879 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x97, 0xf3, 0xdf, 0x93, 0x3a, 0xad, 0x96, 0x8a, 0x46, 0xb9, 0x43, 0x14, 0x4b, 0x40, 0x10,
	0x90, 0xa2, 0x9c, 0x90, 0xe0, 0x28, 0x0f, 0xd5, 0x35, 0x48, 0x48, 0xd7, 0x13, 0x6c, 0x79, 0xb7,
	0x36, 0xf6, 0x24, 0xac, 0xe2, 0x78, 0xcd, 0xee, 0xc6, 0x4d, 0xfa, 0x39, 0x78, 0xe4, 0xc3, 0x20,
	0xbe, 0x0f, 0x8f, 0xbc, 0xa3, 0xfd, 0xe3, 0x34, 0x01, 0x24, 0x74, 0x4f, 0xde, 0x99, 0xf9, 0xcd,
	0xcc, 0x6f, 0xfe, 0x78, 0x60, 0x20, 0xd9, 0xc3, 0xe7, 0xb8, 0xc5, 0x74, 0x52, 0x4a, 0xa1, 0x05,
	0x69, 0xdb, 0xcf, 0xe8, 0xf9, 0x52, 0x88, 0x65, 0x8e, 0x57, 0x56, 0x9a, 0x6f, 0x16, 0x57, 0x4a,
	0xcb, 0x4d, 0xaa, 0x1d, 0x28, 0x9e, 0x42, 0xeb, 0x07, 0x21, 0x35, 0x39, 0x87, 0x76, 0xce, 0xe6,
	0x98, 0x0f, 0x83, 0xcb, 0x60, 0x1c, 0x52, 0x27, 0x18, 0x6d, 0xc5, 0xf2, 0x0d, 0x0e, 0x1b, 0x97,
	0xc1, 0x38, 0xa2, 0x4e, 0x88, 0xff, 0x08, 0xa0, 0xf7, 0x06, 0xf5, 0x83, 0x90, 0x2b, 0x45, 0xde,
	0x85, 0x4e, 0x86, 0x15, 0x4f, 0xd1, 0x7b, 0x7a, 0x89, 0x10, 0x68, 0xa5, 0x3c, 0x93, 0xd6, 0x33,
	0xa4, 0xf6, 0x4d, 0x06, 0xd0, 0xe0, 0xe5, 0xb0, 0x69, 0x35, 0x0d, 0x5e, 0x9a, 0xf0, 0xeb, 0x39,
	0xd7, 0x6a, 0xd8, 0xba, 0x0c, 0xc6, 0x2d, 0xea, 0x04, 0x32, 0x85, 0x81, 0x44, 0x85, 0xb2, 0xc2,
	0x2c, 0x29, 0x85, 0xd4, 0x6a, 0xd8, 0xbe, 0x0c, 0xc6, 0xfd, 0x69, 0xdf, 0x51, 0x9e, 0x18, 0xbe,
	0x34, 0xaa, 0x21, 0x46, 0x52, 0xe4, 0x0b, 0x88, 0xb2, 0x5d, 0xc1, 0xd6, 0x3c, 0xf5, 0x2e, 0x9d,
	0x7f, 0xbb, 0x9c, 0x78, 0x84, 0xf5, 0x88, 0x7f, 0x0d, 0x20, 0xa4, 0xa8, 0xc4, 0x46, 0xa6, 0xa8,
	0xc8, 0x19, 0x34, 0xd3, 0x72, 0x63, 0x4b, 0x68, 0x51, 0xf3, 0x24, 0xcf, 0x20, 0x5c, 0xe3, 0x5a,
	0xc8, 0x5d, 0xb2, 0x9e, 0xdb, 0x22, 0x5a, 0xb4, 0xe7, 0x14, 0x77, 0x73, 0x72, 0x01, 0xdd, 0x8c,
	0xab, 0x95, 0x31, 0x35, 0xad, 0xa9, 0x63, 0xc4, 0xbb, 0xb9, 0xa9, 0x9a, 0x8b, 0xb2, 0x2e, 0xc8,
	0xbe, 0xc9, 0xa7, 0xd0, 0x2b, 0x7c, 0xb7, 0x7c, 0x25, 0xa7, 0x9e, 0x56, 0xdd, 0x44, 0xba, 0x07,
	0xc4, 0x3f, 0x42, 0xf8, 0x5a, 0x2c, 0x5f, 0x89, 0x62, 0xc1, 0x97, 0x96, 0x03, 0xdb, 0x26, 0x0b,
	0x9e, 0xa3, 0xb2, 0xdc, 0x22, 0xda, 0x5b, 0xb3, 0xed, 0x77, 0x46, 0x26, 0x1f, 0xc3, 0x59, 0x6d,
	0x4c, 0x14, 0x7f, 0xc4, 0x27, 0x9e, 0x91, 0xc7, 0xdc, 0xf3, 0x47, 0xbc, 0x9b, 0xc7, 0xbf, 0x05,
	0xd0, 0xfb, 0x89, 0xa9, 0xd5, 0xf7, 0xc5, 0x42, 0x90, 0x09, 0x84, 0xb2, 0xae, 0xda, 0x86, 0xec,
	0x4f, 0xcf, 0x3c, 0x9b, 0x7d, 0x37, 0xe8, 0x13, 0x84, 0x5c, 0x01, 0xe4, 0x62, 0x99, 0xa4, 0x96,
	0x90, 0x8d, 0xff, 0xe4, 0xb0, 0x27, 0x4a, 0xc3, 0x7c, 0xcf, 0xf9, 0x0a, 0x3a, 0x1e, 0xdc, 0xb4,
	0xe0, 0x8b, 0x89, 0xdb, 0xbf, 0x49, 0xbd, 0x7f, 0x93, 0x7b, 0xbb, 0x7f, 0xd4, 0xc3, 0xe2, 0x3f,
	0x03, 0xe8, 0x1a, 0x7a, 0xb7, 0x5c, 0x92, 0xe7, 0x10, 0x66, 0x5c, 0x62, 0xaa, 0x85, 0xdc, 0xf9,
	0x7d, 0x7a, 0x52, 0x90, 0x31, 0x9c, 0xa9, 0x9f, 0x99, 0xc4, 0x2c, 0x61, 0x79, 0x2e, 0xd2, 0x24,
	0xe3, 0xf5, 0x7a, 0x0d, 0x9c, 0xfe, 0xc6, 0xa8, 0x4d, 0x9c, 0x8f, 0xe0, 0xd4, 0x23, 0x35, 0x53,
	0x2b, 0x0b, 0x74, 0x5b, 0x17, 0x39, 0x75, 0x9d, 0xef, 0x19, 0x84, 0xb9, 0x48, 0x59, 0x6e, 0x11,
	0x2d, 0x8b, 0xe8, 0x59, 0x85, 0x31, 0x5e, 0x40, 0xd7, 0x94, 0x6e, 0x4c, 0x6d, 0xb7, 0xda, 0xb9,
	0x58, 0x1a, 0xc3, 0xfb, 0xd0, 0x57, 0x98, 0x4a, 0xd4, 0xca, 0x1a, 0x3b, 0xd6, 0x08, 0x5e, 0xb5,
	0x0f, 0xbb, 0x4c, 0x72, 0xac, 0x30, 0x1f, 0x76, 0xeb, 0xb0, 0xcb, 0xd7, 0x46, 0x8e, 0xff, 0xf2,
	0xf5, 0xce, 0x8a, 0x8a, 0x5c, 0x03, 0x14, 0x22, 0xc3, 0x84, 0x69, 0x2d, 0xcd, 0x38, 0x9a, 0xe3,
	0xfe, 0xf4, 0x3d, 0xdf, 0x5d, 0x8f, 0x99, 0xbc, 0x11, 0x19, 0xde, 0x18, 0xfb, 0xac, 0xd0, 0x72,
	0x47, 0xc3, 0xa2, 0x96, 0xc9, 0x0b, 0xe8, 0x62, 0x51, 0x25, 0x6b, 0x56, 0x0e, 0x1b, 0xd6, 0x75,
	0xf4, 0x0f, 0xd7, 0x59, 0x51, 0xdd, 0xb1, 0xd2, 0xf9, 0x75, 0xd0, 0x0a, 0xa3, 0x6b, 0x18, 0x1c,
	0x47, 0x34, 0xbb, 0xbf, 0xc2, 0xba, 0xdd, 0xe6, 0x79, 0xfc, 0xdb, 0x87, 0xfe, 0xb7, 0x7f, 0xd9,
	0xf8, 0x2a, 0x18, 0x7d, 0x0d, 0xfd, 0x83, 0xa0, 0x6f, 0xe3, 0x1a, 0xa7, 0xd0, 0x9f, 0x6d, 0x31,
	0x7d, 0x25, 0x0a, 0x8d, 0x5b, 0x4d, 0x3e, 0x81, 0xde, 0x7e, 0x36, 0x6e, 0x0f, 0x07, 0x07, 0xec,
	0x6f, 0xb9, 0xa4, 0x5d, 0xed, 0xa7, 0x54, 0x43, 0xb1, 0xa8, 0xfc, 0x06, 0x0e, 0x8e, 0x0b, 0x75,
	0xd0, 0x59, 0x51, 0xc5, 0x0a, 0x4e, 0xee, 0x35, 0x93, 0x9a, 0xe2, 0x2f, 0x1b, 0x54, 0x9a, 0x7c,
	0x09, 0x27, 0xe6, 0x22, 0x9a, 0xfd, 0x35, 0x59, 0x7d, 0x26, 0xe2, 0xdd, 0x0f, 0xf8, 0xd0, 0x3e,
	0x1e, 0x90, 0xfb, 0x0c, 0x42, 0x9b, 0x91, 0x17, 0x0b, 0xe1, 0x53, 0x9e, 0x1e, 0xa4, 0x34, 0x7f,
	0x12, 0xb5, 0x9c, 0xcc, 0x2b, 0xfe, 0x3d, 0x80, 0xe8, 0x56, 0xf2, 0x0a, 0xa5, 0xff, 0xa1, 0xc9,
	0x35, 0xf4, 0xcc, 0x19, 0xb2, 0xa3, 0x71, 0x53, 0xfd, 0xc0, 0xbb, 0x1f, 0xe1, 0xec, 0x5d, 0xda,
	0x4f, 0xa8, 0x5b, 0x3a, 0xc9, 0x9f, 0xc9, 0xc6, 0xfe, 0x4c, 0x7e, 0x08, 0x03, 0xb6, 0xd1, 0x22,
	0x61, 0x59, 0x85, 0x52, 0x73, 0x85, 0x76, 0x99, 0x7b, 0x34, 0x32, 0xda, 0x9b, 0x5a, 0x39, 0x7a,
	0x09, 0x27, 0x87, 0xf1, 0xfe, 0x6f, 0x38, 0xd1, 0xe1, 0x70, 0x10, 0x22, 0xdf, 0x37, 0x55, 0x8a,
	0x42, 0xa1, 0x59, 0x7e, 0xd7, 0x81, 0xac, 0xbe, 0xeb, 0xb6, 0xdc, 0x8c, 0x7c, 0x03, 0x83, 0xcc,
	0xd6, 0x90, 0xf8, 0x9b, 0xe5, 0xfb, 0x73, 0xfe, 0x5f, 0x05, 0xd2, 0x28, 0x3b, 0x14, 0xe3, 0x2e,
	0xb4, 0x67, 0xeb, 0x52, 0xef, 0xa6, 0xdf, 0x42, 0x97, 0xb2, 0x07, 0xd3, 0x7f, 0x32, 0x85, 0xb6,
	0x4d, 0x4d, 0xde, 0xf1, 0x11, 0x0e, 0x07, 0x38, 0x3a, 0x3f, 0x56, 0x3a, 0x76, 0xf3, 0x8e, 0x55,
	0xbe, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x61, 0x04, 0x9c, 0xf3, 0x06, 0x00, 0x00,
}
