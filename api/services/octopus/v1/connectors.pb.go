// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: services/octopus/v1/connectors.proto

package octopus

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListConnectorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceId string `protobuf:"bytes,1,opt,name=workspaceId,proto3" json:"workspaceId,omitempty"`
}

func (x *ListConnectorsRequest) Reset() {
	*x = ListConnectorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_octopus_v1_connectors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConnectorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConnectorsRequest) ProtoMessage() {}

func (x *ListConnectorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_octopus_v1_connectors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConnectorsRequest.ProtoReflect.Descriptor instead.
func (*ListConnectorsRequest) Descriptor() ([]byte, []int) {
	return file_services_octopus_v1_connectors_proto_rawDescGZIP(), []int{0}
}

func (x *ListConnectorsRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type CreateConnectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceId               string `protobuf:"bytes,1,opt,name=workspaceId,proto3" json:"workspaceId,omitempty"`
	AirbyteWorkspaceId        string `protobuf:"bytes,2,opt,name=airbyteWorkspaceId,proto3" json:"airbyteWorkspaceId,omitempty"`
	AirbyteSourceDefinitionId string `protobuf:"bytes,3,opt,name=airbyteSourceDefinitionId,proto3" json:"airbyteSourceDefinitionId,omitempty"`
	AirbyteDestinationId      string `protobuf:"bytes,4,opt,name=airbyteDestinationId,proto3" json:"airbyteDestinationId,omitempty"`
}

func (x *CreateConnectorRequest) Reset() {
	*x = CreateConnectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_octopus_v1_connectors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConnectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConnectorRequest) ProtoMessage() {}

func (x *CreateConnectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_octopus_v1_connectors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConnectorRequest.ProtoReflect.Descriptor instead.
func (*CreateConnectorRequest) Descriptor() ([]byte, []int) {
	return file_services_octopus_v1_connectors_proto_rawDescGZIP(), []int{1}
}

func (x *CreateConnectorRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CreateConnectorRequest) GetAirbyteWorkspaceId() string {
	if x != nil {
		return x.AirbyteWorkspaceId
	}
	return ""
}

func (x *CreateConnectorRequest) GetAirbyteSourceDefinitionId() string {
	if x != nil {
		return x.AirbyteSourceDefinitionId
	}
	return ""
}

func (x *CreateConnectorRequest) GetAirbyteDestinationId() string {
	if x != nil {
		return x.AirbyteDestinationId
	}
	return ""
}

type MutateConnectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectorId string           `protobuf:"bytes,1,opt,name=connectorId,proto3" json:"connectorId,omitempty"`
	Secrets     *structpb.Struct `protobuf:"bytes,2,opt,name=secrets,proto3" json:"secrets,omitempty"`
	Config      *structpb.Struct `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *MutateConnectorRequest) Reset() {
	*x = MutateConnectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_octopus_v1_connectors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateConnectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateConnectorRequest) ProtoMessage() {}

func (x *MutateConnectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_octopus_v1_connectors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateConnectorRequest.ProtoReflect.Descriptor instead.
func (*MutateConnectorRequest) Descriptor() ([]byte, []int) {
	return file_services_octopus_v1_connectors_proto_rawDescGZIP(), []int{2}
}

func (x *MutateConnectorRequest) GetConnectorId() string {
	if x != nil {
		return x.ConnectorId
	}
	return ""
}

func (x *MutateConnectorRequest) GetSecrets() *structpb.Struct {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *MutateConnectorRequest) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

type MutateConnectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Logs    []string `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *MutateConnectorResponse) Reset() {
	*x = MutateConnectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_octopus_v1_connectors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateConnectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateConnectorResponse) ProtoMessage() {}

func (x *MutateConnectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_octopus_v1_connectors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateConnectorResponse.ProtoReflect.Descriptor instead.
func (*MutateConnectorResponse) Descriptor() ([]byte, []int) {
	return file_services_octopus_v1_connectors_proto_rawDescGZIP(), []int{3}
}

func (x *MutateConnectorResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MutateConnectorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MutateConnectorResponse) GetLogs() []string {
	if x != nil {
		return x.Logs
	}
	return nil
}

type GetConnectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectorId string `protobuf:"bytes,1,opt,name=connectorId,proto3" json:"connectorId,omitempty"`
}

func (x *GetConnectorRequest) Reset() {
	*x = GetConnectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_octopus_v1_connectors_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectorRequest) ProtoMessage() {}

func (x *GetConnectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_octopus_v1_connectors_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectorRequest.ProtoReflect.Descriptor instead.
func (*GetConnectorRequest) Descriptor() ([]byte, []int) {
	return file_services_octopus_v1_connectors_proto_rawDescGZIP(), []int{4}
}

func (x *GetConnectorRequest) GetConnectorId() string {
	if x != nil {
		return x.ConnectorId
	}
	return ""
}

type AuthenticateConnectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectorId string `protobuf:"bytes,1,opt,name=connectorId,proto3" json:"connectorId,omitempty"`
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *AuthenticateConnectorRequest) Reset() {
	*x = AuthenticateConnectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_octopus_v1_connectors_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateConnectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateConnectorRequest) ProtoMessage() {}

func (x *AuthenticateConnectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_octopus_v1_connectors_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateConnectorRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateConnectorRequest) Descriptor() ([]byte, []int) {
	return file_services_octopus_v1_connectors_proto_rawDescGZIP(), []int{5}
}

func (x *AuthenticateConnectorRequest) GetConnectorId() string {
	if x != nil {
		return x.ConnectorId
	}
	return ""
}

func (x *AuthenticateConnectorRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type TestConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Logs    []string `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *TestConnectionResponse) Reset() {
	*x = TestConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_octopus_v1_connectors_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestConnectionResponse) ProtoMessage() {}

func (x *TestConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_octopus_v1_connectors_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestConnectionResponse.ProtoReflect.Descriptor instead.
func (*TestConnectionResponse) Descriptor() ([]byte, []int) {
	return file_services_octopus_v1_connectors_proto_rawDescGZIP(), []int{6}
}

func (x *TestConnectionResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TestConnectionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TestConnectionResponse) GetLogs() []string {
	if x != nil {
		return x.Logs
	}
	return nil
}

type SyncConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectorId          string `protobuf:"bytes,1,opt,name=connectorId,proto3" json:"connectorId,omitempty"`
	AirbyteDestinationId string `protobuf:"bytes,2,opt,name=airbyteDestinationId,proto3" json:"airbyteDestinationId,omitempty"`
}

func (x *SyncConnectionRequest) Reset() {
	*x = SyncConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_octopus_v1_connectors_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncConnectionRequest) ProtoMessage() {}

func (x *SyncConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_octopus_v1_connectors_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncConnectionRequest.ProtoReflect.Descriptor instead.
func (*SyncConnectionRequest) Descriptor() ([]byte, []int) {
	return file_services_octopus_v1_connectors_proto_rawDescGZIP(), []int{7}
}

func (x *SyncConnectionRequest) GetConnectorId() string {
	if x != nil {
		return x.ConnectorId
	}
	return ""
}

func (x *SyncConnectionRequest) GetAirbyteDestinationId() string {
	if x != nil {
		return x.AirbyteDestinationId
	}
	return ""
}

type Connector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WorkspaceId               string                 `protobuf:"bytes,2,opt,name=workspaceId,proto3" json:"workspaceId,omitempty"`
	Name                      string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AirbyteSourceDefinitionId string                 `protobuf:"bytes,4,opt,name=airbyteSourceDefinitionId,proto3" json:"airbyteSourceDefinitionId,omitempty"`
	IsActive                  bool                   `protobuf:"varint,5,opt,name=isActive,proto3" json:"isActive,omitempty"`
	AuthScheme                AuthType               `protobuf:"varint,6,opt,name=authScheme,proto3,enum=v1.AuthType" json:"authScheme,omitempty"`
	Config                    *structpb.Struct       `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`
	UpdatedAt                 *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Connector) Reset() {
	*x = Connector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_octopus_v1_connectors_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connector) ProtoMessage() {}

func (x *Connector) ProtoReflect() protoreflect.Message {
	mi := &file_services_octopus_v1_connectors_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connector.ProtoReflect.Descriptor instead.
func (*Connector) Descriptor() ([]byte, []int) {
	return file_services_octopus_v1_connectors_proto_rawDescGZIP(), []int{8}
}

func (x *Connector) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Connector) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *Connector) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Connector) GetAirbyteSourceDefinitionId() string {
	if x != nil {
		return x.AirbyteSourceDefinitionId
	}
	return ""
}

func (x *Connector) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Connector) GetAuthScheme() AuthType {
	if x != nil {
		return x.AuthScheme
	}
	return AuthType_OAUTH2
}

func (x *Connector) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Connector) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ConnectorList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connectors []*Connector `protobuf:"bytes,1,rep,name=connectors,proto3" json:"connectors,omitempty"`
}

func (x *ConnectorList) Reset() {
	*x = ConnectorList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_octopus_v1_connectors_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectorList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorList) ProtoMessage() {}

func (x *ConnectorList) ProtoReflect() protoreflect.Message {
	mi := &file_services_octopus_v1_connectors_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorList.ProtoReflect.Descriptor instead.
func (*ConnectorList) Descriptor() ([]byte, []int) {
	return file_services_octopus_v1_connectors_proto_rawDescGZIP(), []int{9}
}

func (x *ConnectorList) GetConnectors() []*Connector {
	if x != nil {
		return x.Connectors
	}
	return nil
}

var File_services_octopus_v1_connectors_proto protoreflect.FileDescriptor

var file_services_octopus_v1_connectors_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x63, 0x74, 0x6f, 0x70,
	0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f,
	0x63, 0x74, 0x6f, 0x70, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x6f, 0x63, 0x74, 0x6f, 0x70, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0xdc, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x19, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x9e, 0x01, 0x0a, 0x16, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x5f, 0x0a, 0x17, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x37, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x22, 0x54, 0x0a, 0x1c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x5e, 0x0a, 0x16, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x6d, 0x0a, 0x15, 0x53, 0x79, 0x6e, 0x63,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xc4, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x19, 0x61,
	0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19,
	0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3e,
	0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x32, 0x88,
	0x04, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x0f, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x6e, 0x0a, 0x1a, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19,
	0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x3a, 0x01, 0x2a, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x74, 0x6c, 0x65, 0x72, 0x68, 0x71,
	0x2f, 0x62, 0x75, 0x74, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x63, 0x74, 0x6f, 0x70, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x6f, 0x63, 0x74, 0x6f, 0x70, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_octopus_v1_connectors_proto_rawDescOnce sync.Once
	file_services_octopus_v1_connectors_proto_rawDescData = file_services_octopus_v1_connectors_proto_rawDesc
)

func file_services_octopus_v1_connectors_proto_rawDescGZIP() []byte {
	file_services_octopus_v1_connectors_proto_rawDescOnce.Do(func() {
		file_services_octopus_v1_connectors_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_octopus_v1_connectors_proto_rawDescData)
	})
	return file_services_octopus_v1_connectors_proto_rawDescData
}

var file_services_octopus_v1_connectors_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_services_octopus_v1_connectors_proto_goTypes = []interface{}{
	(*ListConnectorsRequest)(nil),        // 0: v1.ListConnectorsRequest
	(*CreateConnectorRequest)(nil),       // 1: v1.CreateConnectorRequest
	(*MutateConnectorRequest)(nil),       // 2: v1.MutateConnectorRequest
	(*MutateConnectorResponse)(nil),      // 3: v1.MutateConnectorResponse
	(*GetConnectorRequest)(nil),          // 4: v1.GetConnectorRequest
	(*AuthenticateConnectorRequest)(nil), // 5: v1.AuthenticateConnectorRequest
	(*TestConnectionResponse)(nil),       // 6: v1.TestConnectionResponse
	(*SyncConnectionRequest)(nil),        // 7: v1.SyncConnectionRequest
	(*Connector)(nil),                    // 8: v1.Connector
	(*ConnectorList)(nil),                // 9: v1.ConnectorList
	(*structpb.Struct)(nil),              // 10: google.protobuf.Struct
	(AuthType)(0),                        // 11: v1.AuthType
	(*timestamppb.Timestamp)(nil),        // 12: google.protobuf.Timestamp
}
var file_services_octopus_v1_connectors_proto_depIdxs = []int32{
	10, // 0: v1.MutateConnectorRequest.secrets:type_name -> google.protobuf.Struct
	10, // 1: v1.MutateConnectorRequest.config:type_name -> google.protobuf.Struct
	11, // 2: v1.Connector.authScheme:type_name -> v1.AuthType
	10, // 3: v1.Connector.config:type_name -> google.protobuf.Struct
	12, // 4: v1.Connector.updatedAt:type_name -> google.protobuf.Timestamp
	8,  // 5: v1.ConnectorList.connectors:type_name -> v1.Connector
	0,  // 6: v1.ConnectorsService.ListConnectors:input_type -> v1.ListConnectorsRequest
	1,  // 7: v1.ConnectorsService.CreateConnector:input_type -> v1.CreateConnectorRequest
	2,  // 8: v1.ConnectorsService.MutateConnector:input_type -> v1.MutateConnectorRequest
	4,  // 9: v1.ConnectorsService.GetConnector:input_type -> v1.GetConnectorRequest
	5,  // 10: v1.ConnectorsService.AuthenticateOAuthConnector:input_type -> v1.AuthenticateConnectorRequest
	9,  // 11: v1.ConnectorsService.ListConnectors:output_type -> v1.ConnectorList
	8,  // 12: v1.ConnectorsService.CreateConnector:output_type -> v1.Connector
	3,  // 13: v1.ConnectorsService.MutateConnector:output_type -> v1.MutateConnectorResponse
	8,  // 14: v1.ConnectorsService.GetConnector:output_type -> v1.Connector
	8,  // 15: v1.ConnectorsService.AuthenticateOAuthConnector:output_type -> v1.Connector
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_services_octopus_v1_connectors_proto_init() }
func file_services_octopus_v1_connectors_proto_init() {
	if File_services_octopus_v1_connectors_proto != nil {
		return
	}
	file_services_octopus_v1_requests_proto_init()
	file_services_octopus_v1_responses_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_services_octopus_v1_connectors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConnectorsRequest); i {
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
		file_services_octopus_v1_connectors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConnectorRequest); i {
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
		file_services_octopus_v1_connectors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateConnectorRequest); i {
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
		file_services_octopus_v1_connectors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateConnectorResponse); i {
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
		file_services_octopus_v1_connectors_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectorRequest); i {
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
		file_services_octopus_v1_connectors_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateConnectorRequest); i {
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
		file_services_octopus_v1_connectors_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestConnectionResponse); i {
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
		file_services_octopus_v1_connectors_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncConnectionRequest); i {
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
		file_services_octopus_v1_connectors_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connector); i {
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
		file_services_octopus_v1_connectors_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectorList); i {
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
			RawDescriptor: file_services_octopus_v1_connectors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_octopus_v1_connectors_proto_goTypes,
		DependencyIndexes: file_services_octopus_v1_connectors_proto_depIdxs,
		MessageInfos:      file_services_octopus_v1_connectors_proto_msgTypes,
	}.Build()
	File_services_octopus_v1_connectors_proto = out.File
	file_services_octopus_v1_connectors_proto_rawDesc = nil
	file_services_octopus_v1_connectors_proto_goTypes = nil
	file_services_octopus_v1_connectors_proto_depIdxs = nil
}
