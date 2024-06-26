// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: bbq/intake/v1/intake_service.proto

package intakev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type RecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reading []*Reading `protobuf:"bytes,1,rep,name=reading,proto3" json:"reading,omitempty"`
}

func (x *RecordRequest) Reset() {
	*x = RecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbq_intake_v1_intake_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordRequest) ProtoMessage() {}

func (x *RecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bbq_intake_v1_intake_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordRequest.ProtoReflect.Descriptor instead.
func (*RecordRequest) Descriptor() ([]byte, []int) {
	return file_bbq_intake_v1_intake_service_proto_rawDescGZIP(), []int{0}
}

func (x *RecordRequest) GetReading() []*Reading {
	if x != nil {
		return x.Reading
	}
	return nil
}

type RecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId  string                 `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	RecordedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=recorded_at,json=recordedAt,proto3" json:"recorded_at,omitempty"`
}

func (x *RecordResponse) Reset() {
	*x = RecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbq_intake_v1_intake_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordResponse) ProtoMessage() {}

func (x *RecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bbq_intake_v1_intake_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordResponse.ProtoReflect.Descriptor instead.
func (*RecordResponse) Descriptor() ([]byte, []int) {
	return file_bbq_intake_v1_intake_service_proto_rawDescGZIP(), []int{1}
}

func (x *RecordResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *RecordResponse) GetRecordedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RecordedAt
	}
	return nil
}

type SessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description  string  `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	DeviceName   *string `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3,oneof" json:"device_name,omitempty"`
	SensorName   *string `protobuf:"bytes,3,opt,name=sensor_name,json=sensorName,proto3,oneof" json:"sensor_name,omitempty"`
	SubjectId    *string `protobuf:"bytes,4,opt,name=subject_id,json=subjectId,proto3,oneof" json:"subject_id,omitempty"`
	DesiredState *string `protobuf:"bytes,5,opt,name=desired_state,json=desiredState,proto3,oneof" json:"desired_state,omitempty"`
}

func (x *SessionRequest) Reset() {
	*x = SessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbq_intake_v1_intake_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionRequest) ProtoMessage() {}

func (x *SessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bbq_intake_v1_intake_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionRequest.ProtoReflect.Descriptor instead.
func (*SessionRequest) Descriptor() ([]byte, []int) {
	return file_bbq_intake_v1_intake_service_proto_rawDescGZIP(), []int{2}
}

func (x *SessionRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SessionRequest) GetDeviceName() string {
	if x != nil && x.DeviceName != nil {
		return *x.DeviceName
	}
	return ""
}

func (x *SessionRequest) GetSensorName() string {
	if x != nil && x.SensorName != nil {
		return *x.SensorName
	}
	return ""
}

func (x *SessionRequest) GetSubjectId() string {
	if x != nil && x.SubjectId != nil {
		return *x.SubjectId
	}
	return ""
}

func (x *SessionRequest) GetDesiredState() string {
	if x != nil && x.DesiredState != nil {
		return *x.DesiredState
	}
	return ""
}

type SessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *SessionResponse) Reset() {
	*x = SessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbq_intake_v1_intake_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionResponse) ProtoMessage() {}

func (x *SessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bbq_intake_v1_intake_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionResponse.ProtoReflect.Descriptor instead.
func (*SessionResponse) Descriptor() ([]byte, []int) {
	return file_bbq_intake_v1_intake_service_proto_rawDescGZIP(), []int{3}
}

func (x *SessionResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type CreateSensorAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorAlert *SensorAlert       `protobuf:"bytes,1,opt,name=sensor_alert,json=sensorAlert,proto3" json:"sensor_alert,omitempty"`
	Rules       []*SensorAlertRule `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *CreateSensorAlertRequest) Reset() {
	*x = CreateSensorAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbq_intake_v1_intake_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSensorAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSensorAlertRequest) ProtoMessage() {}

func (x *CreateSensorAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bbq_intake_v1_intake_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSensorAlertRequest.ProtoReflect.Descriptor instead.
func (*CreateSensorAlertRequest) Descriptor() ([]byte, []int) {
	return file_bbq_intake_v1_intake_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateSensorAlertRequest) GetSensorAlert() *SensorAlert {
	if x != nil {
		return x.SensorAlert
	}
	return nil
}

func (x *CreateSensorAlertRequest) GetRules() []*SensorAlertRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type CreateSensorAlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorAlert *SensorAlert `protobuf:"bytes,1,opt,name=sensor_alert,json=sensorAlert,proto3" json:"sensor_alert,omitempty"`
}

func (x *CreateSensorAlertResponse) Reset() {
	*x = CreateSensorAlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbq_intake_v1_intake_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSensorAlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSensorAlertResponse) ProtoMessage() {}

func (x *CreateSensorAlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bbq_intake_v1_intake_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSensorAlertResponse.ProtoReflect.Descriptor instead.
func (*CreateSensorAlertResponse) Descriptor() ([]byte, []int) {
	return file_bbq_intake_v1_intake_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSensorAlertResponse) GetSensorAlert() *SensorAlert {
	if x != nil {
		return x.SensorAlert
	}
	return nil
}

var File_bbq_intake_v1_intake_service_proto protoreflect.FileDescriptor

var file_bbq_intake_v1_intake_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x62, 0x62, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x62, 0x62, 0x71, 0x2e, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x62, 0x62, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x62, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a,
	0x0d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x07, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x62, 0x62, 0x71, 0x2e, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x22, 0x6c, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8d,
	0x02, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x0a, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x22, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0c, 0x64, 0x65,
	0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e,
	0x5f, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x43,
	0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x62, 0x71, 0x2e, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x8f, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x62, 0x71, 0x2e, 0x69, 0x6e, 0x74,
	0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12,
	0x34, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x62, 0x62, 0x71, 0x2e, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x5a, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x62, 0x71, 0x2e, 0x69,
	0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x32, 0xa4, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x2e,
	0x62, 0x62, 0x71, 0x2e, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x62,
	0x71, 0x2e, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x62, 0x62, 0x71, 0x2e, 0x69, 0x6e,
	0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x62, 0x71, 0x2e, 0x69, 0x6e, 0x74,
	0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb9, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d,
	0x2e, 0x62, 0x62, 0x71, 0x2e, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x12,
	0x49, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x73, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x62, 0x71, 0x2f, 0x69, 0x6e,
	0x74, 0x61, 0x6b, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x65, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x42, 0x49, 0x58, 0xaa, 0x02, 0x0d, 0x42, 0x62, 0x71, 0x2e, 0x49, 0x6e, 0x74,
	0x61, 0x6b, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x42, 0x62, 0x71, 0x5c, 0x49, 0x6e, 0x74,
	0x61, 0x6b, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x42, 0x62, 0x71, 0x5c, 0x49, 0x6e, 0x74,
	0x61, 0x6b, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0f, 0x42, 0x62, 0x71, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x61, 0x6b, 0x65,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bbq_intake_v1_intake_service_proto_rawDescOnce sync.Once
	file_bbq_intake_v1_intake_service_proto_rawDescData = file_bbq_intake_v1_intake_service_proto_rawDesc
)

func file_bbq_intake_v1_intake_service_proto_rawDescGZIP() []byte {
	file_bbq_intake_v1_intake_service_proto_rawDescOnce.Do(func() {
		file_bbq_intake_v1_intake_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_bbq_intake_v1_intake_service_proto_rawDescData)
	})
	return file_bbq_intake_v1_intake_service_proto_rawDescData
}

var file_bbq_intake_v1_intake_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bbq_intake_v1_intake_service_proto_goTypes = []interface{}{
	(*RecordRequest)(nil),             // 0: bbq.intake.v1.RecordRequest
	(*RecordResponse)(nil),            // 1: bbq.intake.v1.RecordResponse
	(*SessionRequest)(nil),            // 2: bbq.intake.v1.SessionRequest
	(*SessionResponse)(nil),           // 3: bbq.intake.v1.SessionResponse
	(*CreateSensorAlertRequest)(nil),  // 4: bbq.intake.v1.CreateSensorAlertRequest
	(*CreateSensorAlertResponse)(nil), // 5: bbq.intake.v1.CreateSensorAlertResponse
	(*Reading)(nil),                   // 6: bbq.intake.v1.Reading
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
	(*Session)(nil),                   // 8: bbq.intake.v1.Session
	(*SensorAlert)(nil),               // 9: bbq.intake.v1.SensorAlert
	(*SensorAlertRule)(nil),           // 10: bbq.intake.v1.SensorAlertRule
}
var file_bbq_intake_v1_intake_service_proto_depIdxs = []int32{
	6,  // 0: bbq.intake.v1.RecordRequest.reading:type_name -> bbq.intake.v1.Reading
	7,  // 1: bbq.intake.v1.RecordResponse.recorded_at:type_name -> google.protobuf.Timestamp
	8,  // 2: bbq.intake.v1.SessionResponse.session:type_name -> bbq.intake.v1.Session
	9,  // 3: bbq.intake.v1.CreateSensorAlertRequest.sensor_alert:type_name -> bbq.intake.v1.SensorAlert
	10, // 4: bbq.intake.v1.CreateSensorAlertRequest.rules:type_name -> bbq.intake.v1.SensorAlertRule
	9,  // 5: bbq.intake.v1.CreateSensorAlertResponse.sensor_alert:type_name -> bbq.intake.v1.SensorAlert
	0,  // 6: bbq.intake.v1.IntakeService.Record:input_type -> bbq.intake.v1.RecordRequest
	2,  // 7: bbq.intake.v1.IntakeService.Session:input_type -> bbq.intake.v1.SessionRequest
	1,  // 8: bbq.intake.v1.IntakeService.Record:output_type -> bbq.intake.v1.RecordResponse
	3,  // 9: bbq.intake.v1.IntakeService.Session:output_type -> bbq.intake.v1.SessionResponse
	8,  // [8:10] is the sub-list for method output_type
	6,  // [6:8] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_bbq_intake_v1_intake_service_proto_init() }
func file_bbq_intake_v1_intake_service_proto_init() {
	if File_bbq_intake_v1_intake_service_proto != nil {
		return
	}
	file_bbq_intake_v1_bbq_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bbq_intake_v1_intake_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordRequest); i {
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
		file_bbq_intake_v1_intake_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordResponse); i {
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
		file_bbq_intake_v1_intake_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionRequest); i {
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
		file_bbq_intake_v1_intake_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionResponse); i {
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
		file_bbq_intake_v1_intake_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSensorAlertRequest); i {
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
		file_bbq_intake_v1_intake_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSensorAlertResponse); i {
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
	file_bbq_intake_v1_intake_service_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bbq_intake_v1_intake_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bbq_intake_v1_intake_service_proto_goTypes,
		DependencyIndexes: file_bbq_intake_v1_intake_service_proto_depIdxs,
		MessageInfos:      file_bbq_intake_v1_intake_service_proto_msgTypes,
	}.Build()
	File_bbq_intake_v1_intake_service_proto = out.File
	file_bbq_intake_v1_intake_service_proto_rawDesc = nil
	file_bbq_intake_v1_intake_service_proto_goTypes = nil
	file_bbq_intake_v1_intake_service_proto_depIdxs = nil
}
